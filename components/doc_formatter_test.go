package rephtml

import (
	"strings"
	"testing"
)

func TestFormatHTMLKeepsComparisonTextLiteral(t *testing.T) {
	input := []byte(`<html><body><p>Keep 1 < 2 and 3 > 2</p></body></html>`)
	got := string(formatHTML(input))
	want := "<html>\n" +
		"\t<body>\n" +
		"\t\t<p>Keep 1 < 2 and 3 > 2</p>\n" +
		"\t</body>\n" +
		"</html>\n"

	if got != want {
		t.Fatalf("unexpected formatted html:\ngot:\n%s\nwant:\n%s", got, want)
	}
}

func TestFormatHTMLAllowsGreaterThanInsideQuotedAttributes(t *testing.T) {
	input := []byte(`<html><body><div data-rule="score > 10"><span>ok</span></div></body></html>`)
	got := string(formatHTML(input))
	want := "<html>\n" +
		"\t<body>\n" +
		"\t\t<div data-rule=\"score > 10\">\n" +
		"\t\t\t<span>ok</span>\n" +
		"\t\t</div>\n" +
		"\t</body>\n" +
		"</html>\n"

	if got != want {
		t.Fatalf("unexpected formatted html:\ngot:\n%s\nwant:\n%s", got, want)
	}
}

func TestFormatHTMLPreservesRawTextElementBody(t *testing.T) {
	body := "  first\n\tsecond < third > fourth\n"
	input := []byte(`<html><body><pre>` + body + `</pre></body></html>`)
	got := string(formatHTML(input))

	start := strings.Index(got, "<pre>")
	end := strings.Index(got, "</pre>")
	if start == -1 || end == -1 {
		t.Fatalf("formatted html is missing pre tags:\n%s", got)
	}

	raw := got[start+len("<pre>") : end]
	if raw != body {
		t.Fatalf("pre body was changed:\ngot  %q\nwant %q\nfull output:\n%s", raw, body, got)
	}
}

func TestFormatHTMLIndentsStyleBody(t *testing.T) {
	input := []byte("<html><head><style>\nbody {\n\tcolor: red;\n}\n</style></head></html>")
	got := string(formatHTML(input))
	want := "<html>\n" +
		"\t<head>\n" +
		"\t\t<style>\n" +
		"\t\t\tbody {\n" +
		"\t\t\t\tcolor: red;\n" +
		"\t\t\t}\n" +
		"\t\t</style>\n" +
		"\t</head>\n" +
		"</html>\n"

	if got != want {
		t.Fatalf("unexpected formatted style block:\ngot:\n%s\nwant:\n%s", got, want)
	}
}
