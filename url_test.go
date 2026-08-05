package rephtml

import (
	"strings"
	"testing"
)

// Escaping a URL attribute value stops it breaking out of its quotes but says
// nothing about what following the link does. These tests cover the scheme
// filter that handles the rest.

func TestUnsafeURLSchemesAreBlocked(t *testing.T) {
	unsafe := []struct {
		name string
		url  string
	}{
		{"javascript", "javascript:alert(1)"},
		{"mixed case", "JaVaScRiPt:alert(1)"},
		{"uppercase", "JAVASCRIPT:alert(1)"},
		{"leading space", "   javascript:alert(1)"},
		{"leading control", "\x01\x02javascript:alert(1)"},
		{"embedded tab", "java\tscript:alert(1)"},
		{"embedded newline", "java\nscript:alert(1)"},
		{"embedded carriage return", "java\rscript:alert(1)"},
		{"embedded nul", "java\x00script:alert(1)"},
		{"tab before colon", "javascript\t:alert(1)"},
		{"vbscript", "vbscript:msgbox(1)"},
		{"livescript", "livescript:alert(1)"},
		{"mocha", "mocha:alert(1)"},
		{"data html", "data:text/html,<script>alert(1)</script>"},
		{"data html base64", "data:text/html;base64,PHNjcmlwdD4="},
		{"data svg", "data:image/svg+xml,<svg onload=alert(1)>"},
		{"data xhtml", "data:application/xhtml+xml,<html/>"},
		{"data no media type", "data:,hello"},
	}

	for _, tt := range unsafe {
		if IsSafeURL(tt.url) {
			t.Errorf("%s: IsSafeURL(%q) = true, want false", tt.name, tt.url)
		}
		got := NewAnchor().Link(tt.url).Text("click").HTML()
		if !strings.Contains(got, BlockedURL) {
			t.Errorf("%s: rendered %q, want the blocked placeholder", tt.name, got)
		}
	}
}

func TestSafeURLsPassThrough(t *testing.T) {
	safe := []struct {
		name string
		url  string
	}{
		{"absolute path", "/reports/q4"},
		{"relative path", "reports/q4.html"},
		{"fragment", "#section-2"},
		{"query", "?page=2"},
		{"protocol relative", "//cdn.example.com/app.js"},
		{"http", "http://example.com/a"},
		{"https", "https://example.com/a?b=c#d"},
		{"mailto", "mailto:someone@example.com"},
		{"tel", "tel:+15551234567"},
		{"sms", "sms:+15551234567"},
		{"ftp", "ftp://files.example.com/x"},
		{"blob", "blob:https://example.com/uuid"},
		{"app deep link", "zoommtg://zoom.us/join?confno=1"},
		{"colon in path", "/a/b:c/d"},
		{"relative with colon", "foo/bar:baz"},
		{"data png", "data:image/png;base64,iVBORw0KGgo="},
		{"data gif", "data:image/gif;base64,R0lGOD"},
		{"data audio", "data:audio/mpeg;base64,SUQz"},
		{"data font", "data:font/woff2;base64,d09GMg"},
	}

	for _, tt := range safe {
		if !IsSafeURL(tt.url) {
			t.Errorf("%s: IsSafeURL(%q) = false, want true", tt.name, tt.url)
		}
		got := NewAnchor().Link(tt.url).Text("click").HTML()
		if strings.Contains(got, BlockedURL) {
			t.Errorf("%s: %q was blocked but should be allowed, rendered %q", tt.name, tt.url, got)
		}
	}
}

// TestURLFilteringAppliesToEveryURLAttribute walks the URL-bearing attributes
// rather than just href, since a javascript: URL is equally live in several of
// them.
func TestURLFilteringAppliesToEveryURLAttribute(t *testing.T) {
	const payload = "javascript:alert(1)"

	tests := []struct {
		name string
		got  string
	}{
		{"Anchor href", NewAnchor().Link(payload).HTML()},
		{"Area href", NewArea().Href(payload).HTML()},
		{"Base href", NewBase().Href(payload).HTML()},
		{"Link href", NewLink().Href(payload).HTML()},
		{"Img src", NewImg().Src(payload).HTML()},
		{"Script src", NewScript().Src(payload).HTML()},
		{"Iframe src", NewIframe().Src(payload).HTML()},
		{"Embed src", NewEmbed().Src(payload).HTML()},
		{"Audio src", NewAudio().Src(payload).HTML()},
		{"Video src", NewVideo().Src(payload).HTML()},
		{"Video poster", NewVideo().Poster(payload).HTML()},
		{"Track src", NewTrack().Src(payload).HTML()},
		{"Source src", NewSource().Src(payload).HTML()},
		{"Object data", NewObject().Data(payload).HTML()},
		{"Form action", NewForm().Action(payload).HTML()},
		{"Button formaction", NewButton().FormAction(payload).HTML()},
		{"Blockquote cite", NewBlockquote().Cite(payload).HTML()},
		{"Q cite", NewQ().Cite(payload).HTML()},
		{"Del cite", NewDel().Cite(payload).HTML()},
		{"Ins cite", NewIns().Cite(payload).HTML()},
	}

	for _, tt := range tests {
		if strings.Contains(tt.got, "javascript:") {
			t.Errorf("%s: unsafe URL reached the output: %s", tt.name, tt.got)
		}
		if !strings.Contains(tt.got, BlockedURL) {
			t.Errorf("%s: expected the blocked placeholder, got %s", tt.name, tt.got)
		}
	}
}

func TestDocumentManifestIsFiltered(t *testing.T) {
	doc := NewHtmlFile().Manifest("javascript:alert(1)")
	got, err := doc.RenderString()
	if err != nil {
		t.Fatalf("RenderString: %v", err)
	}
	if strings.Contains(got, "javascript:") {
		t.Errorf("unsafe manifest reached the output: %s", got)
	}
}

// TestSrcsetFiltersEachCandidate covers the list form, where the URL has to be
// picked out of each comma-separated candidate.
func TestSrcsetFiltersEachCandidate(t *testing.T) {
	got := NewSource().Srcset("/a.png 1x, javascript:alert(1) 2x, /c.png 3x").HTML()

	if strings.Contains(got, "javascript:") {
		t.Errorf("unsafe srcset candidate reached the output: %s", got)
	}
	for _, keep := range []string{"/a.png 1x", "/c.png 3x", BlockedURL + " 2x"} {
		if !strings.Contains(got, keep) {
			t.Errorf("expected %q in %s", keep, got)
		}
	}
}

func TestSrcsetLeavesSafeListsUntouched(t *testing.T) {
	srcset := "/a.png 1x, /b.png 2x"
	got := NewSource().Srcset(srcset).HTML()

	if !strings.Contains(got, srcset) {
		t.Errorf("safe srcset was altered:\ngot  %s\nwant it to contain %q", got, srcset)
	}
}

// TestBlockedURLIsVisible pins that a rejected URL is replaced rather than
// dropped, so the problem shows up in the output instead of silently changing
// what a link points at.
func TestBlockedURLIsVisible(t *testing.T) {
	got := NewAnchor().Link("javascript:alert(1)").Text("Click me").HTML()
	want := `<a href="#rephtml-blocked-url">Click me</a>`

	if got != want {
		t.Errorf("Anchor HTML() = %q, want %q", got, want)
	}
}

func TestEmptyURLAttributesAreStillOmitted(t *testing.T) {
	if got, want := NewAnchor().Text("x").HTML(), "<a>x</a>"; got != want {
		t.Errorf("Anchor HTML() = %q, want %q", got, want)
	}
	if got, want := NewImg().HTML(), "<img>"; got != want {
		t.Errorf("Img HTML() = %q, want %q", got, want)
	}
}
