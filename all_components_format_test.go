package rephtml

import (
	"go/ast"
	"go/parser"
	"go/token"
	"os"
	"path/filepath"
	"sort"
	"strings"
	"testing"
)

type formatElementCase struct {
	name    string
	element Element
}

// renderDocumentString renders a document through the public API and fails the
// test if the document recorded a structure error.
func renderDocumentString(t *testing.T, html *HtmlFile) string {
	t.Helper()
	got, err := html.RenderString()
	if err != nil {
		t.Fatalf("RenderString: %v", err)
	}
	return got
}

// renderDocumentMarkup renders a document's markup directly, bypassing the
// error gate in RenderString, so tests can assert what a rejected element did
// or did not emit.
func renderDocumentMarkup(html *HtmlFile) string {
	html.prepare()
	return string(html.rawBytes())
}

func formatRenderedForTest(input []byte) string {
	return string(formatRenderedDocumentHTML(renderedDocumentHTML(input)))
}

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

func TestInvalidHeadElementRecordsDocumentError(t *testing.T) {
	html := NewHtmlFile()
	html.AddToHead(NewP().Text("body element"))

	if err := html.Err(); err == nil {
		t.Fatal("expected document error")
	}

	path := filepath.Join(t.TempDir(), "invalid.html")
	if err := html.WriteToFile(path); err == nil {
		t.Fatal("expected WriteToFile to return document error")
	}
	if _, err := os.Stat(path); !os.IsNotExist(err) {
		t.Fatalf("document error should not create file, stat error: %v", err)
	}
}

func TestAddToHeadAndBodyPrepareDocumentSections(t *testing.T) {
	html := NewHtmlFile()
	html.AddToHead(NewMeta().Charset("utf-8"))
	html.AddToBody(NewP().Text("Hello"))

	want := `<html><head><meta charset="utf-8"></head><body><p>Hello</p></body></html>`
	if got := renderDocumentString(t, html); got != want {
		t.Fatalf("unexpected compact html:\ngot  %q\nwant %q", got, want)
	}
}

func TestAddToHeadFlattensHeadContents(t *testing.T) {
	html := NewHtmlFile()
	html.AddToHead(NewHead().Add(NewTitle().Text("Nested Head")))

	want := `<html><head><title>Nested Head</title></head></html>`
	if got := renderDocumentString(t, html); got != want {
		t.Fatalf("unexpected html:\ngot  %q\nwant %q", got, want)
	}
}

func TestAddToHeadFlattensHeadWrapperAndMergesStyles(t *testing.T) {
	html := NewHtmlFile()
	html.AddToHead(NewHead().AddStyle("color", "red").Add(NewTitle().Text("Styled Head")))

	want := `<html><head style="color: red;"><title>Styled Head</title></head></html>`
	if got := renderDocumentString(t, html); got != want {
		t.Fatalf("unexpected html:\ngot  %q\nwant %q", got, want)
	}
}

func TestAddToBodyFlattensBodyContents(t *testing.T) {
	html := NewHtmlFile()
	html.AddToBody(NewBody().Add(NewP().Text("Nested Body")))

	want := `<html><body><p>Nested Body</p></body></html>`
	if got := renderDocumentString(t, html); got != want {
		t.Fatalf("unexpected html:\ngot  %q\nwant %q", got, want)
	}
}

func TestAddToBodyFlattensBodyWrapperAndMergesAttributes(t *testing.T) {
	html := NewHtmlFile()
	html.AddToBody(
		NewBody().
			OnLoad("init()").
			OnUnload("cleanup()").
			AddStyle("color", "red").
			Add(NewP().Text("Loaded Body")),
	)

	want := `<html><body onload="init()" onunload="cleanup()" style="color: red;"><p>Loaded Body</p></body></html>`
	if got := renderDocumentString(t, html); got != want {
		t.Fatalf("unexpected html:\ngot  %q\nwant %q", got, want)
	}
}

func TestAddToHeadRejectsBodyWrapper(t *testing.T) {
	html := NewHtmlFile()
	html.AddToHead(NewBody().Add(NewP().Text("Invalid Body")))
	if html.Err() == nil {
		t.Fatal("expected document error")
	}

	want := `<html></html>`
	if got := renderDocumentMarkup(html); got != want {
		t.Fatalf("unexpected html:\ngot  %q\nwant %q", got, want)
	}
}

func TestAddToBodyRejectsHeadWrapper(t *testing.T) {
	html := NewHtmlFile()
	html.AddToBody(NewHead().Add(NewTitle().Text("Invalid Head")))
	if html.Err() == nil {
		t.Fatal("expected document error")
	}

	want := `<html></html>`
	if got := renderDocumentMarkup(html); got != want {
		t.Fatalf("unexpected html:\ngot  %q\nwant %q", got, want)
	}
}

func TestRenderFormattedReturnsHTMLAndError(t *testing.T) {
	html := NewHtmlFile().Lang("en").AddToBody(NewP().Text("Hello"))

	got, err := html.RenderFormattedString()
	if err != nil {
		t.Fatal(err)
	}

	want := "<html lang=\"en\">\n" +
		"\t<body>\n" +
		"\t\t<p>Hello</p>\n" +
		"\t</body>\n" +
		"</html>\n"
	if got != want {
		t.Fatalf("unexpected formatted html:\ngot:\n%s\nwant:\n%s", got, want)
	}
}

func TestRenderFormattedReturnsPendingError(t *testing.T) {
	html := NewHtmlFile().AddToHead(NewP().Text("Invalid"))
	if got, err := html.RenderFormattedString(); err == nil || got != "" {
		t.Fatalf("expected error and empty result, got result %q error %v", got, err)
	}
}

func TestFormatRenderedDocumentHTMLKeepsComparisonTextLiteral(t *testing.T) {
	input := []byte(`<html><body><p>Keep 1 < 2 and 3 > 2</p></body></html>`)
	got := formatRenderedForTest(input)
	want := "<html>\n" +
		"\t<body>\n" +
		"\t\t<p>Keep 1 < 2 and 3 > 2</p>\n" +
		"\t</body>\n" +
		"</html>\n"

	if got != want {
		t.Fatalf("unexpected formatted html:\ngot:\n%s\nwant:\n%s", got, want)
	}
}

func TestFormatRenderedDocumentHTMLAllowsGreaterThanInsideQuotedAttributes(t *testing.T) {
	input := []byte(`<html><body><div data-rule="score > 10"><span>ok</span></div></body></html>`)
	got := formatRenderedForTest(input)
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

func TestFormatRenderedDocumentHTMLFormatsTableInsideDiv(t *testing.T) {
	input := []byte(`<html><body><div><table><tr><td>Alpha Beta</td></tr></table></div></body></html>`)
	got := formatRenderedForTest(input)
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

func TestFormatRenderedDocumentHTMLPreservesRawTextElementBody(t *testing.T) {
	body := "  first\n\tsecond < third > fourth\n"
	input := []byte(`<html><body><pre>` + body + `</pre></body></html>`)
	got := formatRenderedForTest(input)

	assertRawElementBody(t, got, "pre", body)
}

func TestFormatRenderedDocumentHTMLIndentsStyleBody(t *testing.T) {
	input := []byte("<html><head><style>\nbody {\n\tcolor: red;\n}\n</style></head></html>")
	got := formatRenderedForTest(input)
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

func TestFormatRenderedDocumentHTMLFormatsDoctypeDocument(t *testing.T) {
	input := []byte(`<!doctype html><html><body><p>Hello</p></body></html>`)
	got := formatRenderedForTest(input)
	want := "<!doctype html>\n" +
		"<html>\n" +
		"\t<body>\n" +
		"\t\t<p>Hello</p>\n" +
		"\t</body>\n" +
		"</html>\n"

	if got != want {
		t.Fatalf("unexpected formatted doctype document:\ngot:\n%s\nwant:\n%s", got, want)
	}
}

func TestFormatRenderedDocumentHTMLFormatsNestedDocumentSections(t *testing.T) {
	input := []byte(`<html lang="en"><head><meta charset="utf-8"><title>Report</title></head><body><main><section><h2>Metrics</h2><p>Fast</p></section><aside><p>Notes</p></aside></main></body></html>`)
	got := formatRenderedForTest(input)
	want := "<html lang=\"en\">\n" +
		"\t<head>\n" +
		"\t\t<meta charset=\"utf-8\">\n" +
		"\t\t<title>Report</title>\n" +
		"\t</head>\n" +
		"\t<body>\n" +
		"\t\t<main>\n" +
		"\t\t\t<section>\n" +
		"\t\t\t\t<h2>Metrics</h2>\n" +
		"\t\t\t\t<p>Fast</p>\n" +
		"\t\t\t</section>\n" +
		"\t\t\t<aside>\n" +
		"\t\t\t\t<p>Notes</p>\n" +
		"\t\t\t</aside>\n" +
		"\t\t</main>\n" +
		"\t</body>\n" +
		"</html>\n"

	if got != want {
		t.Fatalf("unexpected formatted nested document:\ngot:\n%s\nwant:\n%s", got, want)
	}
}

func TestFormatRenderedDocumentHTMLFormatsCommentsAndVoidElements(t *testing.T) {
	input := []byte(`<html><body><!--intro--><hr><img src="chart.png" alt="Chart"><p>Line<br>break</p></body></html>`)
	got := formatRenderedForTest(input)
	want := "<html>\n" +
		"\t<body>\n" +
		"\t\t<!--intro-->\n" +
		"\t\t<hr>\n" +
		"\t\t<img src=\"chart.png\" alt=\"Chart\">\n" +
		"\t\t<p>\n" +
		"\t\t\tLine\n" +
		"\t\t\t<br>\n" +
		"\t\t\tbreak\n" +
		"\t\t</p>\n" +
		"\t</body>\n" +
		"</html>\n"

	if got != want {
		t.Fatalf("unexpected formatted comments and void elements:\ngot:\n%s\nwant:\n%s", got, want)
	}
}

func TestFormatRenderedDocumentHTMLPreservesScriptRawTextElementBody(t *testing.T) {
	body := "if (score < 10 && total > 2) {\n\tconsole.log(\"ok\");\n}\n"
	input := []byte(`<html><body><script>` + body + `</script></body></html>`)
	got := formatRenderedForTest(input)

	assertRawElementBody(t, got, "script", body)
}

func TestFormatRenderedDocumentHTMLPreservesTextareaRawTextElementBody(t *testing.T) {
	body := "  first line\nsecond < third > fourth\n"
	input := []byte(`<html><body><textarea>` + body + `</textarea></body></html>`)
	got := formatRenderedForTest(input)

	assertRawElementBody(t, got, "textarea", body)
}

func TestWriteToFileFormatsNestedDocument(t *testing.T) {
	html := NewHtmlFile().Lang("en")
	html.AddToHead(NewMeta().Charset("utf-8"))
	html.AddToHead(NewTitle().Text("Report"))
	html.AddToBody(
		NewMain().Add(
			NewSection().
				Add(NewH2().Text("Summary")).
				Add(NewTable().Headers([]string{"Name", "Value"}).AddRow([]string{"Alpha", "42"})),
		),
	)

	path := filepath.Join(t.TempDir(), "nested.html")
	if err := html.WriteToFile(path); err != nil {
		t.Fatal(err)
	}

	got, err := os.ReadFile(path)
	if err != nil {
		t.Fatal(err)
	}

	want := "<html lang=\"en\">\n" +
		"\t<head>\n" +
		"\t\t<meta charset=\"utf-8\">\n" +
		"\t\t<title>Report</title>\n" +
		"\t</head>\n" +
		"\t<body>\n" +
		"\t\t<main>\n" +
		"\t\t\t<section>\n" +
		"\t\t\t\t<h2>Summary</h2>\n" +
		"\t\t\t\t<table>\n" +
		"\t\t\t\t\t<tr>\n" +
		"\t\t\t\t\t\t<th>Name</th>\n" +
		"\t\t\t\t\t\t<th>Value</th>\n" +
		"\t\t\t\t\t</tr>\n" +
		"\t\t\t\t\t<tr>\n" +
		"\t\t\t\t\t\t<td>Alpha</td>\n" +
		"\t\t\t\t\t\t<td>42</td>\n" +
		"\t\t\t\t\t</tr>\n" +
		"\t\t\t\t</table>\n" +
		"\t\t\t</section>\n" +
		"\t\t</main>\n" +
		"\t</body>\n" +
		"</html>\n"

	if string(got) != want {
		t.Fatalf("unexpected nested document file:\ngot:\n%s\nwant:\n%s", got, want)
	}
}

func TestFormattedStandaloneElementStructs(t *testing.T) {
	for _, tt := range standaloneFormatCases() {
		t.Run(tt.name, func(t *testing.T) {
			rendered := renderForFormatting(t, tt.element)
			got := formatRenderedForTest([]byte(rendered))
			want := formattedLines(rendered)
			if got != want {
				t.Fatalf("unexpected formatted %s:\ngot:\n%s\nwant:\n%s", tt.name, got, want)
			}
		})
	}
}

func TestFormattedHeadElementStructs(t *testing.T) {
	for _, tt := range headFormatCases() {
		t.Run(tt.name, func(t *testing.T) {
			rendered := renderForFormatting(t, tt.element)
			got := formatRenderedForTest([]byte("<html><head>" + rendered + "</head></html>"))
			want := formattedLines(
				"<html>",
				"\t<head>",
				"\t\t"+rendered,
				"\t</head>",
				"</html>",
			)
			if got != want {
				t.Fatalf("unexpected formatted %s:\ngot:\n%s\nwant:\n%s", tt.name, got, want)
			}
		})
	}
}

func TestFormattedBodyElementStructs(t *testing.T) {
	for _, tt := range bodyFormatCases() {
		t.Run(tt.name, func(t *testing.T) {
			rendered := renderForFormatting(t, tt.element)
			got := formatRenderedForTest([]byte("<html><body>" + rendered + "</body></html>"))
			want := formattedLines(
				"<html>",
				"\t<body>",
				"\t\t"+rendered,
				"\t</body>",
				"</html>",
			)
			if got != want {
				t.Fatalf("unexpected formatted %s:\ngot:\n%s\nwant:\n%s", tt.name, got, want)
			}
		})
	}
}

func TestFormattedStyleElementStructs(t *testing.T) {
	for _, tt := range styleFormatCases() {
		t.Run(tt.name, func(t *testing.T) {
			style := NewStyleElement().Add(tt.element)
			if tt.name == "StyleElement" || tt.name == "Style" {
				style = nil
			}

			var rendered string
			if style != nil {
				rendered = renderForFormatting(t, style)
			} else {
				rendered = renderForFormatting(t, tt.element)
			}

			got := formatRenderedForTest([]byte("<html><head>" + rendered + "</head></html>"))
			want := formattedHeadStyleElement(t, rendered)
			if got != want {
				t.Fatalf("unexpected formatted %s:\ngot:\n%s\nwant:\n%s", tt.name, got, want)
			}
		})
	}
}

func TestFormattingCasesCoverAllElementStructs(t *testing.T) {
	want := exportedElementStructNames(t)
	got := coveredFormattingStructNames()

	var missing []string
	for name := range want {
		if !got[name] {
			missing = append(missing, name)
		}
	}
	sort.Strings(missing)
	if len(missing) != 0 {
		t.Fatalf("missing formatting cases for element structs: %s", strings.Join(missing, ", "))
	}

	var unknown []string
	for name := range got {
		if !want[name] {
			unknown = append(unknown, name)
		}
	}
	sort.Strings(unknown)
	if len(unknown) != 0 {
		t.Fatalf("formatting cases reference non-element structs: %s", strings.Join(unknown, ", "))
	}
}

func standaloneFormatCases() []formatElementCase {
	return nil
}

func headFormatCases() []formatElementCase {
	return []formatElementCase{
		{name: "Head", element: NewHead()},
		{name: "Title", element: NewTitle().Text("Title")},
		{name: "Base", element: NewBase().Href("/")},
		{name: "Link", element: NewLink().Rel("stylesheet").Href("/app.css")},
		{name: "Meta", element: NewMeta().Charset("utf-8")},
	}
}

func bodyFormatCases() []formatElementCase {
	return []formatElementCase{
		{name: "Body", element: NewBody()},
		{name: "Header", element: NewHeader()},
		{name: "Nav", element: NewNav()},
		{name: "Main", element: NewMain()},
		{name: "Section", element: NewSection()},
		{name: "Article", element: NewArticle()},
		{name: "Aside", element: NewAside()},
		{name: "Footer", element: NewFooter()},
		{name: "Address", element: NewAddress()},
		{name: "Hgroup", element: NewHgroup()},
		{name: "H1", element: NewH1().Text("H1")},
		{name: "H2", element: NewH2().Text("H2")},
		{name: "H3", element: NewH3().Text("H3")},
		{name: "H4", element: NewH4().Text("H4")},
		{name: "H5", element: NewH5().Text("H5")},
		{name: "H6", element: NewH6().Text("H6")},
		{name: "P", element: NewP().Text("Paragraph")},
		{name: "Comment", element: NewComment().Text("comment")},
		{name: "Hr", element: NewHr()},
		{name: "Pre", element: NewPre().Text("preformatted")},
		{name: "Blockquote", element: NewBlockquote().Text("Quote")},
		{name: "Menu", element: NewMenu()},
		{name: "Ol", element: NewOl()},
		{name: "Ul", element: NewUl()},
		{name: "Li", element: NewLi()},
		{name: "Dl", element: NewDl()},
		{name: "Dt", element: NewDt()},
		{name: "Dd", element: NewDd()},
		{name: "Figure", element: NewFigure()},
		{name: "Figcaption", element: NewFigcaption()},
		{name: "Search", element: NewSearch()},
		{name: "Div", element: NewDiv()},
		{name: "Anchor", element: NewAnchor().Link("/docs").Text("Docs")},
		{name: "Abbr", element: NewAbbr().Text("HTML")},
		{name: "B", element: NewB().Text("Bold")},
		{name: "I", element: NewI().Text("Italic")},
		{name: "Q", element: NewQ().Text("Quote")},
		{name: "S", element: NewS().Text("Old")},
		{name: "U", element: NewU().Text("Underline")},
		{name: "Bdi", element: NewBdi().Text("Isolate")},
		{name: "Bdo", element: NewBdo().Text("Override")},
		{name: "Br", element: NewBr()},
		{name: "Cite", element: NewCite().Text("Citation")},
		{name: "Code", element: NewCode().Text("code")},
		{name: "Data", element: NewData().Value("42").Text("Answer")},
		{name: "Dfn", element: NewDfn().Text("Term")},
		{name: "Em", element: NewEm().Text("Emphasis")},
		{name: "Mark", element: NewMark().Text("Marked")},
		{name: "Ruby", element: NewRuby()},
		{name: "Rb", element: NewRb()},
		{name: "Rt", element: NewRt()},
		{name: "Rtc", element: NewRtc()},
		{name: "Rp", element: NewRp()},
		{name: "Kbd", element: NewKbd().Text("Ctrl+C")},
		{name: "Sub", element: NewSub().Text("2")},
		{name: "Sup", element: NewSup().Text("2")},
		{name: "Samp", element: NewSamp().Text("output")},
		{name: "Small", element: NewSmall().Text("small")},
		{name: "Span", element: NewSpan().Text("span")},
		{name: "Strong", element: NewStrong().Text("strong")},
		{name: "Time", element: NewTime().Datetime("2026-04-26").Text("today")},
		{name: "Var", element: NewVar().Text("x")},
		{name: "Wbr", element: NewWbr()},
		{name: "Form", element: NewForm()},
		{name: "Label", element: NewLabel()},
		{name: "Input", element: NewInput()},
		{name: "Output", element: NewOutput()},
		{name: "Fieldset", element: NewFieldset()},
		{name: "Button", element: NewButton()},
		{name: "Select", element: NewSelect()},
		{name: "Datalist", element: NewDatalist()},
		{name: "Optgroup", element: NewOptgroup()},
		{name: "Option", element: NewOption()},
		{name: "Textarea", element: NewTextarea()},
		{name: "Progress", element: NewProgress()},
		{name: "Meter", element: NewMeter()},
		{name: "Legend", element: NewLegend()},
		{name: "Table", element: NewTable()},
		{name: "Thead", element: NewThead()},
		{name: "Tbody", element: NewTbody()},
		{name: "Tfoot", element: NewTfoot()},
		{name: "Caption", element: NewCaption()},
		{name: "Col", element: NewCol()},
		{name: "Colgroup", element: NewColgroup()},
		{name: "Tr", element: NewTr()},
		{name: "Td", element: NewTd()},
		{name: "Th", element: NewTh()},
		{name: "Area", element: NewArea()},
		{name: "Img", element: NewImg()},
		{name: "Audio", element: NewAudio()},
		{name: "Track", element: NewTrack()},
		{name: "Map", element: NewMap()},
		{name: "Video", element: NewVideo()},
		{name: "Embed", element: NewEmbed()},
		{name: "Iframe", element: NewIframe()},
		{name: "Object", element: NewObject()},
		{name: "Picture", element: NewPicture()},
		{name: "Portal", element: NewPortal()},
		{name: "Source", element: NewSource()},
		{name: "Canvas", element: NewCanvas()},
		{name: "Noscript", element: NewNoscript()},
		{name: "Script", element: NewScript()},
		{name: "Math", element: NewMath()},
		{name: "Svg", element: NewSvg()},
		{name: "Slot", element: NewSlot()},
		{name: "Template", element: NewTemplate()},
		{name: "Details", element: NewDetails()},
		{name: "Dialog", element: NewDialog()},
		{name: "Summary", element: NewSummary()},
		{name: "Del", element: NewDel()},
		{name: "Ins", element: NewIns()},
	}
}

func styleFormatCases() []formatElementCase {
	style := NewStyle("body")
	style.Props = StyleMap{"color": "red"}

	styleElement := NewStyleElement().Text(".x { color: red; }")

	styleRule := NewStyleRule(".card")
	styleRule.Props = StyleMap{"padding": "1rem"}

	fontFace := NewFontFaceRule()
	fontFace.Props = StyleMap{
		"font-family": "Report",
		"src":         `url("/report.woff2") format("woff2")`,
	}

	fontFeatures := NewFontFeatureValuesRule("Report").Text("@styleset {\n\tswash: 1;\n}")

	mediaRule := NewMediaRule("(min-width: 800px)")
	mediaChild := NewStyleRule(".card")
	mediaChild.Props = StyleMap{"padding": "2rem"}
	mediaRule.AddRule(mediaChild)

	keyframeBlock := NewKeyframeBlock("from")
	keyframeBlock.Props = StyleMap{"opacity": "0"}

	keyframes := NewKeyframesRule("fade").
		AddFrame("from", StyleMap{"opacity": "0"}).
		AddFrame("to", StyleMap{"opacity": "1"})

	return []formatElementCase{
		{name: "StyleElement", element: styleElement},
		{name: "Style", element: style},
		{name: "StyleRule", element: styleRule},
		{name: "RawCSSRule", element: NewRawCSSRule("@layer base;")},
		{name: "CharsetRule", element: NewCharsetRule("utf-8")},
		{name: "ImportRule", element: NewImportRule("/base.css").Condition("screen")},
		{name: "FontFaceRule", element: fontFace},
		{name: "FontFeatureValuesRule", element: fontFeatures},
		{name: "MediaRule", element: mediaRule},
		{name: "KeyframeBlock", element: keyframeBlock},
		{name: "KeyframesRule", element: keyframes},
	}
}

func renderForFormatting(t *testing.T, element Element) string {
	t.Helper()
	if element == nil {
		t.Fatal("formatting case has nil element")
	}
	return element.HTML()
}

func assertRawElementBody(t *testing.T, formatted, tag, want string) {
	t.Helper()

	openTag := "<" + tag + ">"
	closeTag := "</" + tag + ">"
	start := strings.Index(formatted, openTag)
	end := strings.Index(formatted, closeTag)
	if start == -1 || end == -1 {
		t.Fatalf("formatted html is missing %s tags:\n%s", tag, formatted)
	}

	got := formatted[start+len(openTag) : end]
	if got != want {
		t.Fatalf("%s body was changed:\ngot  %q\nwant %q\nfull output:\n%s", tag, got, want, formatted)
	}
}

func formattedHeadStyleElement(t *testing.T, rendered string) string {
	t.Helper()
	openEnd := strings.Index(rendered, ">")
	closeStart := strings.LastIndex(rendered, "</style>")
	if openEnd == -1 || closeStart == -1 {
		t.Fatalf("rendered style element is malformed: %q", rendered)
	}

	lines := []string{
		"<html>",
		"\t<head>",
		"\t\t" + rendered[:openEnd+1],
	}
	body := strings.Trim(rendered[openEnd+1:closeStart], "\n")
	if body != "" {
		for _, line := range strings.Split(body, "\n") {
			if strings.TrimSpace(line) == "" {
				lines = append(lines, "")
				continue
			}
			lines = append(lines, "\t\t\t"+line)
		}
	}
	lines = append(lines,
		"\t\t</style>",
		"\t</head>",
		"</html>",
	)
	return formattedLines(lines...)
}

func formattedLines(lines ...string) string {
	return strings.Join(lines, "\n") + "\n"
}

func coveredFormattingStructNames() map[string]bool {
	covered := map[string]bool{}
	for _, group := range [][]formatElementCase{
		standaloneFormatCases(),
		headFormatCases(),
		bodyFormatCases(),
		styleFormatCases(),
	} {
		for _, tc := range group {
			covered[tc.name] = true
		}
	}
	return covered
}

func exportedElementStructNames(t *testing.T) map[string]bool {
	t.Helper()

	fset := token.NewFileSet()
	entries, err := os.ReadDir(".")
	if err != nil {
		t.Fatal(err)
	}

	structs := map[string]bool{}
	methods := map[string]map[string]bool{}
	for _, entry := range entries {
		name := entry.Name()
		if entry.IsDir() || !strings.HasSuffix(name, ".go") || strings.HasSuffix(name, "_test.go") {
			continue
		}

		file, err := parser.ParseFile(fset, name, nil, 0)
		if err != nil {
			t.Fatal(err)
		}
		for _, decl := range file.Decls {
			switch d := decl.(type) {
			case *ast.GenDecl:
				recordExportedStructs(d, structs)
			case *ast.FuncDecl:
				recordElementMethod(d, methods)
			}
		}
	}

	elements := map[string]bool{}
	for name := range structs {
		if methods[name]["prepare"] && !methods[name]["documentRender"] {
			elements[name] = true
		}
	}
	return elements
}

func recordExportedStructs(decl *ast.GenDecl, structs map[string]bool) {
	for _, spec := range decl.Specs {
		typeSpec, ok := spec.(*ast.TypeSpec)
		if !ok || !typeSpec.Name.IsExported() {
			continue
		}
		if _, ok := typeSpec.Type.(*ast.StructType); ok {
			structs[typeSpec.Name.Name] = true
		}
	}
}

// recordElementMethod records the declared methods that identify an element struct.
//
// prepare is the discriminator: Render, HTML and the style setters are promoted
// from the embedded generic bases in element.go and are no longer declared per
// type, but every element still declares its own prepare. HtmlFile also declares
// prepare and is not an Element, so its document-shaped Render, which returns
// ([]byte, error) rather than []byte, is recorded as a disqualifier.
func recordElementMethod(decl *ast.FuncDecl, methods map[string]map[string]bool) {
	if decl.Recv == nil {
		return
	}
	key := decl.Name.Name
	switch key {
	case "prepare":
	case "Render":
		if decl.Type.Results == nil || len(decl.Type.Results.List) == 1 {
			return
		}
		key = "documentRender"
	default:
		return
	}
	name := receiverName(decl.Recv.List[0].Type)
	if name == "" {
		// Generic receivers such as *node[Self] are shared bases, not elements.
		return
	}
	if methods[name] == nil {
		methods[name] = map[string]bool{}
	}
	methods[name][key] = true
}

func receiverName(expr ast.Expr) string {
	switch t := expr.(type) {
	case *ast.Ident:
		return t.Name
	case *ast.StarExpr:
		return receiverName(t.X)
	default:
		return ""
	}
}
