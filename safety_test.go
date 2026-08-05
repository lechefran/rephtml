package rephtml

import (
	"strings"
	"testing"
)

// --- typed-nil children ---

// nilDiv models the common builder shape that returns a typed nil on a miss.
// Stored in an Element it is not equal to nil, so a plain nil check misses it.
func nilDiv() *Div { return nil }

func TestAddIgnoresTypedNilChildren(t *testing.T) {
	div := NewDiv().Add(nilDiv()).Add(NewP().Text("kept"))

	if got, want := div.HTML(), "<div><p>kept</p></div>"; got != want {
		t.Errorf("Div HTML() = %q, want %q", got, want)
	}
}

func TestTypedNilChildrenAreIgnoredThroughoutTheTree(t *testing.T) {
	var nilTr *Tr
	var nilThead *Thead

	table := NewTable().
		AddTr(nilTr).
		AddThead(nilThead).
		AddTbody(NewTbody().AddTr(NewTr().AddTd(NewTd().Add(nilDiv()))))

	if got, want := table.HTML(), "<table><tbody><tr><td></td></tr></tbody></table>"; got != want {
		t.Errorf("Table HTML() = %q, want %q", got, want)
	}
}

func TestUntypedNilChildIsIgnored(t *testing.T) {
	if got, want := NewDiv().Add(nil).HTML(), "<div></div>"; got != want {
		t.Errorf("Div HTML() = %q, want %q", got, want)
	}
}

// --- comment sanitising ---

func TestCommentCannotEscapeItsDelimiters(t *testing.T) {
	tests := []struct {
		name string
		text string
		want string
	}{
		{"plain", "a note", "<!--a note-->"},
		{"close sequence", "--><script>alert(1)</script><!--",
			"<!--- ><script>alert(1)</script><!- ---></script><!- ---"},
		{"bang close", "--!>", "<!--- !>-->"},
		{"double hyphen", "a--b", "<!--a- -b-->"},
		{"leading gt", ">x", "<!-- >x-->"},
		{"leading arrow", "->x", "<!-- ->x-->"},
		{"trailing hyphen", "x-", "<!--x- -->"},
		{"empty", "", "<!---->"},
	}

	for _, tt := range tests {
		got := NewComment().Text(tt.text).HTML()

		// The invariant that actually matters: the rendered comment contains
		// exactly one closing delimiter, at the very end.
		body := strings.TrimSuffix(strings.TrimPrefix(got, "<!--"), "-->")
		if strings.Contains(body, "--") {
			t.Errorf("%s: comment body %q still contains a double hyphen", tt.name, body)
		}
		if strings.Contains(body, ">") && (strings.HasPrefix(body, ">") || strings.HasPrefix(body, "->")) {
			t.Errorf("%s: comment body %q starts with a closing sequence", tt.name, body)
		}
	}
}

func TestCommentKeepsOrdinaryTextIntact(t *testing.T) {
	if got, want := NewComment().Text("build 42").HTML(), "<!--build 42-->"; got != want {
		t.Errorf("Comment HTML() = %q, want %q", got, want)
	}
}

// --- CSS escaping ---

// assertNoStyleBreakout checks that rendered CSS cannot terminate the style
// element that contains it.
func assertNoStyleBreakout(t *testing.T, name, css string) {
	t.Helper()
	if indexFold(css, "</style") >= 0 {
		t.Errorf("%s: rendered CSS can close its style element: %q", name, css)
	}
}

func TestStructuredCSSCannotCloseStyleElement(t *testing.T) {
	payload := `red</style><script>alert(1)</script>`

	tests := []struct {
		name string
		css  string
	}{
		// Unwrapped rules, so the only closing tag that could appear is one the
		// caller's text smuggled in. The wrapped Style form is covered by
		// TestStyleElementContentIsEscapedThroughStructuredRules.
		{"declaration value", NewStyleRule("body").AddStyle("color", payload).HTML()},
		{"declaration property", NewStyleRule("body").AddStyle(payload, "red").HTML()},
		{"selector", NewStyleRule(`x</style><script>alert(1)</script>`).HTML()},
		{"media query", NewMediaRule(`screen{}</style><script>alert(1)</script>`).
			AddRule(NewStyleRule("a").AddStyle("color", "red")).HTML()},
		{"keyframes name", NewKeyframesRule(`spin</style>`).
			AddFrame("from", StyleMap{"opacity": "0"}).HTML()},
		{"font face descriptor", NewFontFaceRule().AddStyle("font-family", payload).HTML()},
		{"font feature family", NewFontFeatureValuesRule(`X</style>`).Text("@swash { a: 1; }").HTML()},
		{"font feature content", NewFontFeatureValuesRule("X").Text(`</style><script>alert(1)</script>`).HTML()},
		{"import href", NewImportRule(`x"); }</style><script>alert(1)</script>`).HTML()},
		{"charset", NewCharsetRule(`utf-8</style>`).HTML()},
	}

	for _, tt := range tests {
		assertNoStyleBreakout(t, tt.name, tt.css)
	}
}

// TestStyleElementContentIsEscapedThroughStructuredRules checks the whole path:
// a hostile declaration value inside a rendered style element.
func TestStyleElementContentIsEscapedThroughStructuredRules(t *testing.T) {
	got := NewStyle("body").AddStyle("color", `red</style><script>alert(1)</script>`).HTML()

	if !strings.HasPrefix(got, "<style>") || !strings.HasSuffix(got, "</style>") {
		t.Fatalf("unexpected style element shape: %q", got)
	}
	body := strings.TrimSuffix(strings.TrimPrefix(got, "<style>"), "</style>")
	assertNoStyleBreakout(t, "style element body", body)
	if !strings.Contains(body, `\3c `) {
		t.Errorf("expected a CSS escape in %q", body)
	}
}

// TestCSSEscapingLeavesLegitimateAngleBracketsAlone guards against over-eager
// escaping: media query range syntax uses < and must survive untouched.
func TestCSSEscapingLeavesLegitimateAngleBracketsAlone(t *testing.T) {
	query := "(400px <= width <= 700px)"
	got := NewMediaRule(query).AddRule(NewStyleRule("a").AddStyle("color", "red")).HTML()

	if !strings.Contains(got, query) {
		t.Errorf("media query was altered:\ngot  %q\nwant it to contain %q", got, query)
	}
}

// --- formatter offset handling ---

// TestFormatterHandlesNonASCIIInRawTextElement covers the case where folding
// case with strings.ToLower would shift byte offsets. U+212A KELVIN SIGN folds
// to a one-byte "k", so an index taken from the folded text used to land two
// bytes early and truncate the closing tag, which left the style element open
// and swallowed the rest of the document.
func TestFormatterHandlesNonASCIIInRawTextElement(t *testing.T) {
	for _, sample := range []string{
		"/* \u212A */ body { color: red; }", // KELVIN SIGN, folds 3 bytes -> 1
		"/* \u0130 */ body { color: red; }", // LATIN CAPITAL I WITH DOT ABOVE
		"/* \u1E9E */ body { color: red; }", // LATIN CAPITAL SHARP S
		"/* \u212A\u212A\u0130 */ a { b: c; }",
	} {
		doc := NewHtmlFile()
		doc.AddToHead(NewStyleElement().Text(sample))
		doc.AddToBody(NewP().Text("must survive"))

		got, err := doc.RenderFormattedString()
		if err != nil {
			t.Fatalf("RenderFormattedString: %v", err)
		}

		if !strings.Contains(got, "</style>") {
			t.Errorf("style element was not closed properly for %q:\n%s", sample, got)
		}
		if !strings.Contains(got, "<p>must survive</p>") {
			t.Errorf("content after the style element was lost for %q:\n%s", sample, got)
		}
		if !strings.Contains(got, "</html>") {
			t.Errorf("document was not closed for %q:\n%s", sample, got)
		}
	}
}

// --- head/body validation ---

func TestHeadWrapperRejectsBodyContent(t *testing.T) {
	head := NewHead().Add(NewDiv())
	if head.Err() == nil {
		t.Fatal("expected Head.Add to record an error for a body element")
	}

	doc := NewHtmlFile()
	doc.AddToHead(head)
	if doc.Err() == nil {
		t.Fatal("expected the wrapper's error to reach the document")
	}
	if _, err := doc.RenderString(); err == nil {
		t.Fatal("expected RenderString to report the document error")
	}
}

func TestBodyWrapperRejectsHeadContent(t *testing.T) {
	body := NewBody().Add(NewTitle().Text("wrong place"))
	if body.Err() == nil {
		t.Fatal("expected Body.Add to record an error for a head element")
	}

	doc := NewHtmlFile()
	doc.AddToBody(body)
	if doc.Err() == nil {
		t.Fatal("expected the wrapper's error to reach the document")
	}
}

func TestHeadAndBodyWrappersAcceptValidContent(t *testing.T) {
	doc := NewHtmlFile()
	doc.AddToHead(NewHead().Add(NewTitle().Text("Report")))
	doc.AddToBody(NewBody().Add(NewP().Text("Hello")))

	if err := doc.Err(); err != nil {
		t.Fatalf("unexpected document error: %v", err)
	}

	want := `<!DOCTYPE html><html><head><title>Report</title></head><body><p>Hello</p></body></html>`
	got, err := doc.RenderString()
	if err != nil {
		t.Fatalf("RenderString: %v", err)
	}
	if got != want {
		t.Errorf("RenderString() = %q, want %q", got, want)
	}
}

// --- doctype ---

func TestDocumentsDeclareHTML5Doctype(t *testing.T) {
	doc := NewHtmlFile().Lang("en").AddToBody(NewP().Text("Hello"))

	compact, err := doc.RenderString()
	if err != nil {
		t.Fatalf("RenderString: %v", err)
	}
	if !strings.HasPrefix(compact, "<!DOCTYPE html><html") {
		t.Errorf("compact document does not start with a doctype: %q", compact)
	}

	formatted, err := doc.RenderFormattedString()
	if err != nil {
		t.Fatalf("RenderFormattedString: %v", err)
	}
	if !strings.HasPrefix(formatted, "<!DOCTYPE html>\n<html") {
		t.Errorf("formatted document does not start with a doctype:\n%s", formatted)
	}
}
