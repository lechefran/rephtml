package rephtml

import "testing"

// Elements do not agree on where the style attribute goes relative to other
// attributes: a handful of inline elements emit it first, most emit it last,
// and the table types emit it mid-list. Attribute order is semantically
// irrelevant in HTML, but it is observable in rendered output, so these tests
// pin the current placement for every group. Without them a refactor that
// normalises placement changes the output of dozens of elements while the rest
// of the suite stays green.

func assertHTML(t *testing.T, name string, e Element, want string) {
	t.Helper()
	if got := e.HTML(); got != want {
		t.Errorf("%s HTML() = %q, want %q", name, got, want)
	}
}

// TestStyleAttributeWrittenFirst covers the inline elements that emit style
// before their other attributes.
func TestStyleAttributeWrittenFirst(t *testing.T) {
	style := StyleMap{"color": "red", "margin": "0"}

	tests := []struct {
		name    string
		element Element
		want    string
	}{
		{"Anchor", NewAnchor().Link("/x").Text("go").Style(style),
			`<a style="color: red; margin: 0;" href="/x">go</a>`},
		{"Abbr", NewAbbr().Title("t").Text("a").Style(style),
			`<abbr style="color: red; margin: 0;" title="t">a</abbr>`},
		{"Q", NewQ().Cite("/c").Text("q").Style(style),
			`<q style="color: red; margin: 0;" cite="/c">q</q>`},
		{"Data", NewData().Value("42").Text("d").Style(style),
			`<data style="color: red; margin: 0;" value="42">d</data>`},
		{"Dfn", NewDfn().Title("t").Text("f").Style(style),
			`<dfn style="color: red; margin: 0;" title="t">f</dfn>`},
	}

	for _, tt := range tests {
		assertHTML(t, tt.name, tt.element, tt.want)
	}
}

// TestStyleAttributeWrittenLast covers the majority convention, where style is
// emitted after every other attribute.
func TestStyleAttributeWrittenLast(t *testing.T) {
	style := StyleMap{"color": "red", "margin": "0"}

	tests := []struct {
		name    string
		element Element
		want    string
	}{
		{"Form", NewForm().Action("/x").Method("post").Name("n").Style(style),
			`<form action="/x" method="post" name="n" style="color: red; margin: 0;"></form>`},
		{"Img", NewImg().Src("/i.png").Alt("alt").Width("10").Height("20").Style(style),
			`<img src="/i.png" alt="alt" width="10" height="20" style="color: red; margin: 0;">`},
		{"Input", NewInput().Type("text").Name("n").Value("v").Required(true).Disabled(true).Style(style),
			`<input type="text" name="n" value="v" required disabled style="color: red; margin: 0;">`},
		// Time is the lone inline element that follows the style-last convention.
		{"Time", NewTime().Datetime("2026-01-02").Text("then").Style(style),
			`<time datetime="2026-01-02" style="color: red; margin: 0;">then</time>`},
	}

	for _, tt := range tests {
		assertHTML(t, tt.name, tt.element, tt.want)
	}
}

// TestStyleAttributeWrittenMidList covers the table types, which emit
// id and class before style and their sizing attributes after it.
func TestStyleAttributeWrittenMidList(t *testing.T) {
	style := StyleMap{"color": "red", "margin": "0"}

	tests := []struct {
		name    string
		element Element
		want    string
	}{
		{"Td", NewTd().Id("i").AddClass("c").Colspan(2).Styles(style),
			`<td id="i" class="c" style="color: red; margin: 0;" colspan="2"></td>`},
		{"Th", NewTh().Id("i").AddClass("c").Colspan(3).Styles(style),
			`<th id="i" class="c" style="color: red; margin: 0;" colspan="3"></th>`},
		{"Col", NewCol().Id("i").AddClass("c").Span(2).Styles(style),
			`<col id="i" class="c" style="color: red; margin: 0;" span="2">`},
		{"Colgroup", NewColgroup().Id("i").AddClass("c").Span(2).Styles(style),
			`<colgroup id="i" class="c" style="color: red; margin: 0;" span="2"></colgroup>`},
		{"Table", NewTable().Id("i").AddClass("c").Styles(style),
			`<table id="i" class="c" style="color: red; margin: 0;"></table>`},
	}

	for _, tt := range tests {
		assertHTML(t, tt.name, tt.element, tt.want)
	}
}

// TestCSSRulesRenderRepeatably renders each CSS rule type three times. The rest
// of the suite renders these once, so a Prepare that failed to reset its buffer
// would go unnoticed; here it shows up as doubled output.
func TestCSSRulesRenderRepeatably(t *testing.T) {
	tests := []struct {
		name    string
		element Element
	}{
		{"Style", NewStyle("body").AddStyle("color", "red")},
		{"StyleRule", NewStyleRule("a").AddStyle("color", "blue")},
		{"CharsetRule", NewCharsetRule("utf-8")},
		{"ImportRule", NewImportRule("/a.css").Condition("screen")},
		{"FontFaceRule", NewFontFaceRule().AddStyle("font-family", "X")},
		{"RawCSSRule", NewRawCSSRule("a { color: red; }")},
		{"FontFeatureValuesRule", NewFontFeatureValuesRule("X").Text("@swash { a: 1; }")},
		{"MediaRule", NewMediaRule("screen").AddRule(NewStyleRule("a").AddStyle("color", "red"))},
		{"KeyframesRule", NewKeyframesRule("spin").AddFrame("from", StyleMap{"opacity": "0"})},
		{"KeyframeBlock", NewKeyframeBlock("to").AddStyle("opacity", "1")},
	}

	for _, tt := range tests {
		first := tt.element.HTML()
		if second := tt.element.HTML(); second != first {
			t.Errorf("%s second HTML() = %q, want %q", tt.name, second, first)
		}
		if third := string(tt.element.Render()); third != first {
			t.Errorf("%s Render() = %q, want %q", tt.name, third, first)
		}
	}
}

// TestNestedTreeRendersRepeatably renders a nested tree twice. Containers
// re-render their children on every pass, so a buffer that is not reset
// somewhere in the tree compounds visibly here.
func TestNestedTreeRendersRepeatably(t *testing.T) {
	tree := NewDiv().AddStyles(StyleMap{"color": "red"}).
		Add(NewSection().Add(NewH1().Text("Title")).Add(NewP().Text("Body & <text>"))).
		Add(NewUl().Add(NewLi().Add(NewAnchor().Link("/a").Text("one"))))

	want := `<div style="color: red;"><section><h1>Title</h1><p>Body &amp; &lt;text&gt;</p></section>` +
		`<ul><li><a href="/a">one</a></li></ul></div>`

	assertHTML(t, "tree", tree, want)
	assertHTML(t, "tree second render", tree, want)
}
