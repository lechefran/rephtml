package rephtml

import "testing"

func assertRender(t *testing.T, name string, element Element, want string) {
	t.Helper()

	element.Prepare()
	if got := string(element.Bytes()); got != want {
		t.Fatalf("%s rendered unexpected HTML:\ngot  %q\nwant %q", name, got, want)
	}
}

func TestElementInterfaceContracts(t *testing.T) {
	var _ Element = NewP()
	var _ HeadElement = NewTitle()
	var _ BodyElement = NewP()
	var _ BodyElement = (*Hgroup)(nil)
	var _ BodyElement = (*Search)(nil)
}

func TestDocumentComponentOutputFormats(t *testing.T) {
	tests := []struct {
		name    string
		element Element
		want    string
	}{
		{
			name: "HtmlFile",
			element: NewHtmlFile().
				Lang("en").
				AddToHead(NewTitle().Text("Doc")).
				AddToBody(NewP().Text("Hello")),
			want: `<html lang="en"><head><title>Doc</title></head><body><p>Hello</p></body></html>`,
		},
		{
			name:    "Head",
			element: NewHead().Add(NewTitle().Text("Doc")),
			want:    `<head><title>Doc</title></head>`,
		},
		{
			name:    "Body",
			element: NewBody().OnLoad("init()").Add(NewP().Text("Loaded")),
			want:    `<body onload="init()"><p>Loaded</p></body>`,
		},
		{
			name:    "Title",
			element: NewTitle().Text(`A&B <title>`),
			want:    `<title>A&amp;B &lt;title&gt;</title>`,
		},
		{
			name:    "Base",
			element: NewBase().Href("/docs").Target("_blank"),
			want:    `<base href="/docs" target="_blank">`,
		},
		{
			name:    "Link",
			element: NewLink().Rel("stylesheet").Href("/app.css"),
			want:    `<link rel="stylesheet" href="/app.css">`,
		},
		{
			name:    "Meta",
			element: NewMeta().Name("viewport").Content("width=device-width"),
			want:    `<meta name="viewport" content="width=device-width">`,
		},
		{
			name:    "StyleElement",
			element: NewStyleElement().Text(".x { color: red; }"),
			want:    `<style>.x { color: red; }</style>`,
		},
	}

	for _, tt := range tests {
		assertRender(t, tt.name, tt.element, tt.want)
	}
}

func TestSectionComponentOutputFormats(t *testing.T) {
	tests := []struct {
		name    string
		element Element
		want    string
	}{
		{name: "Header", element: NewHeader().Add(NewH1().Text("Title")), want: `<header><h1>Title</h1></header>`},
		{name: "Nav", element: NewNav().Role("navigation").Add(NewAnchor().Link("/").Text("Home")), want: `<nav role="navigation"><a href="/">Home</a></nav>`},
		{name: "Main", element: NewMain().Add(NewP().Text("Main")), want: `<main><p>Main</p></main>`},
		{name: "Section", element: NewSection().AriaLabel("Summary").Add(NewP().Text("Section")), want: `<section aria-label="Summary"><p>Section</p></section>`},
		{name: "Article", element: NewArticle().Add(NewP().Text("Article")), want: `<article><p>Article</p></article>`},
		{name: "Aside", element: NewAside().Add(NewP().Text("Aside")), want: `<aside><p>Aside</p></aside>`},
		{name: "Footer", element: NewFooter().Add(NewP().Text("Footer")), want: `<footer><p>Footer</p></footer>`},
		{name: "Address", element: NewAddress().Add(NewP().Text("Address")), want: `<address><p>Address</p></address>`},
	}

	for _, tt := range tests {
		assertRender(t, tt.name, tt.element, tt.want)
	}
}

func TestHeadingComponentOutputFormats(t *testing.T) {
	tests := []struct {
		name    string
		element Element
		want    string
	}{
		{name: "Hgroup", element: NewHgroup().Add(NewH1().Text("Title")).Add(NewP().Text("Subtitle")), want: `<hgroup><h1>Title</h1><p>Subtitle</p></hgroup>`},
		{name: "H1", element: NewH1().Text("One"), want: `<h1>One</h1>`},
		{name: "H2", element: NewH2().Text("Two"), want: `<h2>Two</h2>`},
		{name: "H3", element: NewH3().Text("Three"), want: `<h3>Three</h3>`},
		{name: "H4", element: NewH4().Text("Four"), want: `<h4>Four</h4>`},
		{name: "H5", element: NewH5().Text("Five"), want: `<h5>Five</h5>`},
		{name: "H6", element: NewH6().Text("Six"), want: `<h6>Six</h6>`},
	}

	for _, tt := range tests {
		assertRender(t, tt.name, tt.element, tt.want)
	}
}

func TestTextComponentOutputFormats(t *testing.T) {
	tests := []struct {
		name    string
		element Element
		want    string
	}{
		{name: "P", element: NewP().Text(`A < B`), want: `<p>A &lt; B</p>`},
		{name: "Comment", element: NewComment().Text("note"), want: `<!--note-->`},
		{name: "Hr", element: NewHr(), want: `<hr>`},
		{name: "Pre", element: NewPre().Text(`A < B`), want: `<pre>A &lt; B</pre>`},
		{name: "Blockquote", element: NewBlockquote().Cite("/source").Text("Quote"), want: `<blockquote cite="/source">Quote</blockquote>`},
		{name: "Menu", element: NewMenu().Type("toolbar").Label("Actions").Add(NewLi().Add(NewSpan().Text("Run"))), want: `<menu type="toolbar" label="Actions"><li><span>Run</span></li></menu>`},
		{name: "Ol", element: NewOl().Start(10).Type("A").Reversed(true).Add(NewLi().Add(NewSpan().Text("First"))), want: `<ol start="10" type="A" reversed><li><span>First</span></li></ol>`},
		{name: "Ul", element: NewUl().Add(NewLi().Add(NewSpan().Text("Item"))), want: `<ul><li><span>Item</span></li></ul>`},
		{name: "Li", element: NewLi().Value(2).Add(NewSpan().Text("Item")), want: `<li value="2"><span>Item</span></li>`},
		{name: "Dl", element: NewDl().Add(NewDt().Add(NewSpan().Text("Term"))).Add(NewDd().Add(NewSpan().Text("Definition"))), want: `<dl><dt><span>Term</span></dt><dd><span>Definition</span></dd></dl>`},
		{name: "Dt", element: NewDt().Add(NewSpan().Text("Term")), want: `<dt><span>Term</span></dt>`},
		{name: "Dd", element: NewDd().Add(NewSpan().Text("Definition")), want: `<dd><span>Definition</span></dd>`},
		{name: "Figure", element: NewFigure().Add(NewImg().Src("chart.png").Alt("Chart")).Add(NewFigcaption().Add(NewSpan().Text("Caption"))), want: `<figure><img src="chart.png" alt="Chart"><figcaption><span>Caption</span></figcaption></figure>`},
		{name: "Figcaption", element: NewFigcaption().Add(NewSpan().Text("Caption")), want: `<figcaption><span>Caption</span></figcaption>`},
		{name: "Search", element: NewSearch().Add(NewForm().Action("/search")), want: `<search><form action="/search"></form></search>`},
		{name: "Div", element: NewDiv().Add(NewSpan().Text("Content")), want: `<div><span>Content</span></div>`},
		{name: "DivWithTable", element: NewDiv().Add(NewTable().AddRow([]string{"Alpha Beta"})), want: `<div><table><tr><td>Alpha Beta</td></tr></table></div>`},
	}

	for _, tt := range tests {
		assertRender(t, tt.name, tt.element, tt.want)
	}
}

func TestInlineComponentOutputFormats(t *testing.T) {
	tests := []struct {
		name    string
		element Element
		want    string
	}{
		{name: "Anchor", element: NewAnchor().Link("/docs").Text("Docs"), want: `<a href="/docs">Docs</a>`},
		{name: "Abbr", element: NewAbbr().Title("HyperText Markup Language").Text("HTML"), want: `<abbr title="HyperText Markup Language">HTML</abbr>`},
		{name: "B", element: NewB().Text("Bold"), want: `<b>Bold</b>`},
		{name: "I", element: NewI().Text("Italic"), want: `<i>Italic</i>`},
		{name: "Q", element: NewQ().Cite("/quote").Text("Quote"), want: `<q cite="/quote">Quote</q>`},
		{name: "S", element: NewS().Text("Old"), want: `<s>Old</s>`},
		{name: "U", element: NewU().Text("Underline"), want: `<u>Underline</u>`},
		{name: "Bdi", element: NewBdi().Text("Isolate"), want: `<bdi>Isolate</bdi>`},
		{name: "Bdo", element: NewBdo().Text("Override"), want: `<bdo>Override</bdo>`},
		{name: "Br", element: NewBr(), want: `<br>`},
		{name: "Cite", element: NewCite().Text("Citation"), want: `<cite>Citation</cite>`},
		{name: "Code", element: NewCode().Text("<tag>"), want: `<code>&lt;tag&gt;</code>`},
		{name: "Data", element: NewData().Value("42").Text("Answer"), want: `<data value="42">Answer</data>`},
		{name: "Dfn", element: NewDfn().Title("Definition").Text("Term"), want: `<dfn title="Definition">Term</dfn>`},
		{name: "Em", element: NewEm().Text("Emphasis"), want: `<em>Emphasis</em>`},
		{name: "Mark", element: NewMark().Text("Marked"), want: `<mark>Marked</mark>`},
		{name: "Ruby", element: NewRuby().Add(NewRb().Text("漢")).Add(NewRt().Text("kan")), want: `<ruby><rb>漢</rb><rt>kan</rt></ruby>`},
		{name: "Rb", element: NewRb().Text("Base"), want: `<rb>Base</rb>`},
		{name: "Rt", element: NewRt().Text("Text"), want: `<rt>Text</rt>`},
		{name: "Rtc", element: NewRtc().Add(NewRt().Text("Text")), want: `<rtc><rt>Text</rt></rtc>`},
		{name: "Rp", element: NewRp().Text("("), want: `<rp>(</rp>`},
		{name: "Kbd", element: NewKbd().Text("Ctrl+C"), want: `<kbd>Ctrl+C</kbd>`},
		{name: "Sub", element: NewSub().Text("2"), want: `<sub>2</sub>`},
		{name: "Sup", element: NewSup().Text("2"), want: `<sup>2</sup>`},
		{name: "Samp", element: NewSamp().Text("output"), want: `<samp>output</samp>`},
		{name: "Small", element: NewSmall().Text("small"), want: `<small>small</small>`},
		{name: "Span", element: NewSpan().Text("span"), want: `<span>span</span>`},
		{name: "Strong", element: NewStrong().Text("strong"), want: `<strong>strong</strong>`},
		{name: "Time", element: NewTime().Datetime("2026-04-26").Text("today"), want: `<time datetime="2026-04-26">today</time>`},
		{name: "Var", element: NewVar().Text("x"), want: `<var>x</var>`},
		{name: "Wbr", element: NewWbr(), want: `<wbr>`},
	}

	for _, tt := range tests {
		assertRender(t, tt.name, tt.element, tt.want)
	}
}

func TestTimeDatetimeValidationRecordsError(t *testing.T) {
	timeElement := NewTime().Datetime("not a datetime").Text("invalid")
	if timeElement.Err() == nil {
		t.Fatal("expected invalid datetime error")
	}

	assertRender(t, "Time invalid datetime", timeElement, `<time>invalid</time>`)
}

func TestFormComponentOutputFormats(t *testing.T) {
	tests := []struct {
		name    string
		element Element
		want    string
	}{
		{name: "Form", element: NewForm().Action("/submit").Method("post").Add(NewInput().Name("q")), want: `<form action="/submit" method="post"><input name="q"></form>`},
		{name: "Label", element: NewLabel().For("q").Text("Query"), want: `<label for="q">Query</label>`},
		{name: "Input", element: NewInput().Type("text").Name("q").Required(true), want: `<input type="text" name="q" required>`},
		{name: "Output", element: NewOutput().For("a b").Name("sum").Text("3"), want: `<output for="a b" name="sum">3</output>`},
		{name: "Fieldset", element: NewFieldset().Name("group").Disabled(true).Add(NewLegend().Text("Legend")), want: `<fieldset name="group" disabled><legend>Legend</legend></fieldset>`},
		{name: "Button", element: NewButton().Type("submit").Name("action").Value("save").Text("Save"), want: `<button type="submit" name="action" value="save">Save</button>`},
		{name: "Select", element: NewSelect().Name("choice").Required(true).Add(NewOption().Value("a").Text("A")), want: `<select name="choice" required><option value="a">A</option></select>`},
		{name: "Datalist", element: NewDatalist().Id("choices").Add(NewOption().Value("a").Text("A")), want: `<datalist id="choices"><option value="a">A</option></datalist>`},
		{name: "Optgroup", element: NewOptgroup().Label("Group").Add(NewOption().Text("A")), want: `<optgroup label="Group"><option>A</option></optgroup>`},
		{name: "Option", element: NewOption().Value("a").Label("A").Selected(true).Text("Alpha"), want: `<option value="a" label="A" selected>Alpha</option>`},
		{name: "Textarea", element: NewTextarea().Name("notes").Rows("2").Cols("20").Text("Hello"), want: `<textarea name="notes" rows="2" cols="20">Hello</textarea>`},
		{name: "Progress", element: NewProgress().Value("50").Max("100").Text("50%"), want: `<progress value="50" max="100">50%</progress>`},
		{name: "Meter", element: NewMeter().Value("5").Min("0").Max("10").Low("3").High("8").Optimum("6").Text("5"), want: `<meter value="5" min="0" max="10" low="3" high="8" optimum="6">5</meter>`},
		{name: "Legend", element: NewLegend().Text("Legend"), want: `<legend>Legend</legend>`},
	}

	for _, tt := range tests {
		assertRender(t, tt.name, tt.element, tt.want)
	}
}

func TestTableComponentOutputFormats(t *testing.T) {
	tests := []struct {
		name    string
		element Element
		want    string
	}{
		{name: "Table", element: NewTable().Headers([]string{"A", "B"}).AddRow([]string{"1", "2"}), want: `<table><tr><th>A</th><th>B</th></tr><tr><td>1</td><td>2</td></tr></table>`},
		{name: "Thead", element: NewThead().AddTr(NewTr().AddTh(NewTh().Add(NewSpan().Text("H")))), want: `<thead><tr><th><span>H</span></th></tr></thead>`},
		{name: "Tbody", element: NewTbody().AddTr(NewTr().AddTd(NewTd().Add(NewSpan().Text("C")))), want: `<tbody><tr><td><span>C</span></td></tr></tbody>`},
		{name: "Tfoot", element: NewTfoot().AddTr(NewTr().AddTd(NewTd().Add(NewSpan().Text("F")))), want: `<tfoot><tr><td><span>F</span></td></tr></tfoot>`},
		{name: "Caption", element: NewCaption().Text("Caption"), want: `<caption>Caption</caption>`},
		{name: "Col", element: NewCol().Span(2), want: `<col span="2">`},
		{name: "Colgroup", element: NewColgroup().Span(2).Add(NewCol()), want: `<colgroup span="2"><col></colgroup>`},
		{name: "Tr", element: NewTr().AddTd(NewTd().Add(NewSpan().Text("C"))), want: `<tr><td><span>C</span></td></tr>`},
		{name: "Td", element: NewTd().Colspan(2).Rowspan(3).Add(NewSpan().Text("C")), want: `<td colspan="2" rowspan="3"><span>C</span></td>`},
		{name: "Th", element: NewTh().Colspan(2).Rowspan(3).Scope("col").Add(NewSpan().Text("H")), want: `<th colspan="2" rowspan="3" scope="col"><span>H</span></th>`},
	}

	for _, tt := range tests {
		assertRender(t, tt.name, tt.element, tt.want)
	}
}

func TestTableClassReplacesExistingClasses(t *testing.T) {
	table := NewTable().
		AddClass("legacy").
		AddClasses([]string{"extra"}).
		Class([]string{"current"}).
		AddRow([]string{"cell"})

	assertRender(t, "Table class replacement", table, `<table class="current"><tr><td>cell</td></tr></table>`)
}

func TestTableClassCopiesInputSlice(t *testing.T) {
	classes := []string{"current"}
	table := NewTable().Class(classes).AddRow([]string{"cell"})
	classes[0] = "mutated"

	assertRender(t, "Table class slice copy", table, `<table class="current"><tr><td>cell</td></tr></table>`)
}

func TestMediaComponentOutputFormats(t *testing.T) {
	tests := []struct {
		name    string
		element Element
		want    string
	}{
		{name: "Area", element: NewArea().Alt("Map").Coords("0,0,10,10").Href("/map").Shape("rect").Target("_blank"), want: `<area alt="Map" coords="0,0,10,10" href="/map" shape="rect" target="_blank">`},
		{name: "Img", element: NewImg().Src("image.png").Alt("Image").Width("640").Height("480").Title("Title"), want: `<img src="image.png" alt="Image" width="640" height="480" title="Title">`},
		{name: "Audio", element: NewAudio().Src("audio.mp3").Controls(true).Preload("metadata"), want: `<audio src="audio.mp3" controls preload="metadata"></audio>`},
		{name: "Track", element: NewTrack().Src("captions.vtt").Kind("captions").Srclang("en").Label("English").Default(true), want: `<track src="captions.vtt" kind="captions" srclang="en" label="English" default>`},
		{name: "Map", element: NewMap().Name("primary").Add(NewArea().Href("/")), want: `<map name="primary"><area href="/"></map>`},
		{name: "Video", element: NewVideo().Src("video.mp4").Controls(true).Width("640").Height("360").Poster("poster.png"), want: `<video src="video.mp4" controls width="640" height="360" poster="poster.png"></video>`},
	}

	for _, tt := range tests {
		assertRender(t, tt.name, tt.element, tt.want)
	}
}

func TestEmbeddedComponentOutputFormats(t *testing.T) {
	tests := []struct {
		name    string
		element Element
		want    string
	}{
		{name: "Embed", element: NewEmbed().Src("file.pdf").Type("application/pdf").Width("640").Height("480"), want: `<embed src="file.pdf" type="application/pdf" width="640" height="480">`},
		{name: "Iframe", element: NewIframe().Src("/frame").Width("640").Height("480").Name("frame").Allowfullscreen(true), want: `<iframe src="/frame" width="640" height="480" name="frame" allowfullscreen></iframe>`},
		{name: "Object", element: NewObject().Data("movie.swf").Type("application/x-shockwave-flash").Width("640").Height("480").Add(NewP().Text("Fallback")), want: `<object data="movie.swf" type="application/x-shockwave-flash" width="640" height="480"><p>Fallback</p></object>`},
		{name: "Picture", element: NewPicture().Add(NewSource().Srcset("large.webp").Type("image/webp")).Add(NewImg().Src("fallback.png").Alt("Fallback")), want: `<picture><source srcset="large.webp" type="image/webp"><img src="fallback.png" alt="Fallback"></picture>`},
		{name: "Portal", element: NewPortal().Src("/portal").Referrerpolicy("no-referrer"), want: `<portal src="/portal" referrerpolicy="no-referrer"></portal>`},
		{name: "Source", element: NewSource().Src("media.mp4").Srcset("media.webp").Media("(min-width: 800px)").Sizes("100vw").Type("video/mp4"), want: `<source src="media.mp4" srcset="media.webp" media="(min-width: 800px)" sizes="100vw" type="video/mp4">`},
	}

	for _, tt := range tests {
		assertRender(t, tt.name, tt.element, tt.want)
	}
}

func TestScriptMathWebAndInteractiveOutputFormats(t *testing.T) {
	tests := []struct {
		name    string
		element Element
		want    string
	}{
		{name: "Canvas", element: NewCanvas().Width("300").Height("150").Add(NewSpan().Text("Fallback")), want: `<canvas width="300" height="150"><span>Fallback</span></canvas>`},
		{name: "Noscript", element: NewNoscript().Add(NewP().Text("No script")), want: `<noscript><p>No script</p></noscript>`},
		{name: "Script", element: NewScript().Src("/app.js").Type("module").Async(true).Text(`console.log("ok");`), want: `<script src="/app.js" type="module" async>console.log("ok");</script>`},
		{name: "Math", element: NewMath().Display("block").Add(NewSpan().Text("x")), want: `<math xmlns="http://www.w3.org/1998/Math/MathML" display="block"><span>x</span></math>`},
		{name: "Svg", element: NewSvg().Width("10").Height("10").ViewBox("0 0 10 10"), want: `<svg xmlns="http://www.w3.org/2000/svg" width="10" height="10" viewBox="0 0 10 10"></svg>`},
		{name: "Slot", element: NewSlot().Name("content").Add(NewSpan().Text("Fallback")), want: `<slot name="content"><span>Fallback</span></slot>`},
		{name: "Template", element: NewTemplate().Id("row").Add(NewP().Text("Template")), want: `<template id="row"><p>Template</p></template>`},
		{name: "Details", element: NewDetails().Open(true).Add(NewSummary().Add(NewSpan().Text("More"))).Add(NewP().Text("Details")), want: `<details open><summary><span>More</span></summary><p>Details</p></details>`},
		{name: "Dialog", element: NewDialog().Open(true).Add(NewP().Text("Dialog")), want: `<dialog open><p>Dialog</p></dialog>`},
		{name: "Summary", element: NewSummary().Add(NewSpan().Text("Summary")), want: `<summary><span>Summary</span></summary>`},
		{name: "Del", element: NewDel().Cite("/changes").Datetime("2026-04-26").Add(NewSpan().Text("Removed")), want: `<del cite="/changes" datetime="2026-04-26"><span>Removed</span></del>`},
		{name: "Ins", element: NewIns().Cite("/changes").Datetime("2026-04-26").Add(NewSpan().Text("Added")), want: `<ins cite="/changes" datetime="2026-04-26"><span>Added</span></ins>`},
	}

	for _, tt := range tests {
		assertRender(t, tt.name, tt.element, tt.want)
	}
}

func TestStyleSupportOutputFormats(t *testing.T) {
	style := NewStyle("body")
	style.Props = CssProps{Color: "#111827"}
	assertRender(t, "Style", style, "<style>\nbody {\n\tcolor: #111827;\n}\n</style>")

	rule := NewStyleRule(".card")
	rule.Props = CssProps{Padding: "1rem"}
	assertRender(t, "StyleRule", rule, "\n.card {\n\tpadding: 1rem;\n}\n")

	if NewPropMap().pmap["Color"] != "color" {
		t.Fatal("PropMap does not include Color mapping")
	}
}

func TestOptionsAndStrictnessDefaults(t *testing.T) {
	options := Options{Validation: STRICT, AllowImages: true}
	if options.Validation != STRICT || !options.AllowImages {
		t.Fatalf("unexpected options values: %+v", options)
	}
	if DEFAULT != "default" || LAZY != "lazy" || STRICT != "strict" {
		t.Fatalf("unexpected strictness constants: %q %q %q", DEFAULT, LAZY, STRICT)
	}
}
