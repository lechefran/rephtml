package tests

import (
	"os"
	"path/filepath"
	"testing"

	rephtml "github.com/lechefran/rephtml/components"
)

func TestWriteToFileFormatsDocument(t *testing.T) {
	html := rephtml.NewHtmlFile().Lang("en")
	html.AddToHead(rephtml.NewTitle().Text("Readable Report"))
	html.AddToBody(rephtml.NewH1().Text("Summary"))

	path := filepath.Join(t.TempDir(), "report.html")
	html.WriteToFile(path)

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

func TestAddToHeadAndBodyPrepareDocumentSections(t *testing.T) {
	html := rephtml.NewHtmlFile()
	html.AddToHead(rephtml.NewMeta().Charset("utf-8"))
	html.AddToBody(rephtml.NewP().Text("Hello"))
	html.Prepare()

	want := `<html><head><meta charset="utf-8"></head><body><p>Hello</p></body></html>`
	if got := string(html.Bytes()); got != want {
		t.Fatalf("unexpected compact html:\ngot  %q\nwant %q", got, want)
	}
}

func TestAddToHeadFlattensHeadContents(t *testing.T) {
	html := rephtml.NewHtmlFile()
	html.AddToHead(rephtml.NewHead().Add(rephtml.NewTitle().Text("Nested Head")))
	html.Prepare()

	want := `<html><head><title>Nested Head</title></head></html>`
	if got := string(html.Bytes()); got != want {
		t.Fatalf("unexpected html:\ngot  %q\nwant %q", got, want)
	}
}

func TestAddToBodyFlattensBodyContents(t *testing.T) {
	html := rephtml.NewHtmlFile()
	html.AddToBody(rephtml.NewBody().Add(rephtml.NewP().Text("Nested Body")))
	html.Prepare()

	want := `<html><body><p>Nested Body</p></body></html>`
	if got := string(html.Bytes()); got != want {
		t.Fatalf("unexpected html:\ngot  %q\nwant %q", got, want)
	}
}
