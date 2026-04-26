package rephtml

import (
	"os"
	"path/filepath"
	"strings"
	"testing"
)

func TestWriteToFileFormatsDocument(t *testing.T) {
	html := NewHtmlFile().Lang("en")
	html.AddToHead(NewTitle().Text("Readable Report"))
	html.AddToBody(NewH1().Text("Summary"))

	path := filepath.Join(t.TempDir(), "report.html")
	if err := html.WriteToFile(path); err != nil {
		t.Fatal(err)
	}

	got, err := os.ReadFile(path)
	if err != nil {
		t.Fatal(err)
	}

	want := "<html lang=\"en\">\n" +
		"\t<head>\n" +
		"\t\t<title>Readable Report</title>\n" +
		"\t</head>\n" +
		"\t<body>\n" +
		"\t\t<h1>Summary</h1>\n" +
		"\t</body>\n" +
		"</html>\n"

	if string(got) != want {
		t.Fatalf("unexpected formatted html:\ngot:\n%s\nwant:\n%s", got, want)
	}
}

func TestWriteToFileReturnsFilesystemError(t *testing.T) {
	html := NewHtmlFile().AddToBody(NewP().Text("Hello"))

	err := html.WriteToFile(filepath.Join(t.TempDir(), "missing", "report.html"))
	if err == nil {
		t.Fatal("expected filesystem error")
	}
}

func TestStrictValidationRecordsErrorWithoutExiting(t *testing.T) {
	html := NewHtmlFile().AddOptions(Options{Validation: STRICT})
	html.AddToHead(NewP().Text("body element"))

	if html.Err() == nil {
		t.Fatal("expected strict validation error")
	}

	path := filepath.Join(t.TempDir(), "invalid.html")
	if err := html.WriteToFile(path); err == nil {
		t.Fatal("expected WriteToFile to return strict validation error")
	}
	if _, err := os.Stat(path); !os.IsNotExist(err) {
		t.Fatalf("strict validation error should not create file, stat error: %v", err)
	}
}

func TestAddToHeadAndBodyPrepareDocumentSections(t *testing.T) {
	html := NewHtmlFile()
	html.AddToHead(NewMeta().Charset("utf-8"))
	html.AddToBody(NewP().Text("Hello"))
	html.Prepare()

	want := `<html><head><meta charset="utf-8"></head><body><p>Hello</p></body></html>`
	if got := string(html.Bytes()); got != want {
		t.Fatalf("unexpected compact html:\ngot  %q\nwant %q", got, want)
	}
}

func TestAddToHeadFlattensHeadContents(t *testing.T) {
	html := NewHtmlFile()
	html.AddToHead(NewHead().Add(NewTitle().Text("Nested Head")))
	html.Prepare()

	want := `<html><head><title>Nested Head</title></head></html>`
	if got := string(html.Bytes()); got != want {
		t.Fatalf("unexpected html:\ngot  %q\nwant %q", got, want)
	}
}

func TestAddToBodyFlattensBodyContents(t *testing.T) {
	html := NewHtmlFile()
	html.AddToBody(NewBody().Add(NewP().Text("Nested Body")))
	html.Prepare()

	want := `<html><body><p>Nested Body</p></body></html>`
	if got := string(html.Bytes()); got != want {
		t.Fatalf("unexpected html:\ngot  %q\nwant %q", got, want)
	}
}

// TestFormatHTMLKeepsComparisonTextLiteral provides TestFormatHTMLKeepsComparisonTextLiteral behavior for the package.
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

// TestFormatHTMLAllowsGreaterThanInsideQuotedAttributes provides TestFormatHTMLAllowsGreaterThanInsideQuotedAttributes behavior for the package.
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

func TestFormatHTMLFormatsTableInsideDiv(t *testing.T) {
	input := []byte(`<html><body><div><table><tr><td>Alpha Beta</td></tr></table></div></body></html>`)
	got := string(formatHTML(input))
	want := "<html>\n" +
		"\t<body>\n" +
		"\t\t<div>\n" +
		"\t\t\t<table>\n" +
		"\t\t\t\t<tr>\n" +
		"\t\t\t\t\t<td>Alpha Beta</td>\n" +
		"\t\t\t\t</tr>\n" +
		"\t\t\t</table>\n" +
		"\t\t</div>\n" +
		"\t</body>\n" +
		"</html>\n"

	if got != want {
		t.Fatalf("unexpected formatted html:\ngot:\n%s\nwant:\n%s", got, want)
	}
}

// TestFormatHTMLPreservesRawTextElementBody provides TestFormatHTMLPreservesRawTextElementBody behavior for the package.
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

// TestFormatHTMLIndentsStyleBody provides TestFormatHTMLIndentsStyleBody behavior for the package.
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
