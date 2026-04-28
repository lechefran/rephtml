package rephtml

import "testing"

// TestStyleElementAddRuleBuildsUnwrappedCSS provides TestStyleElementAddRuleBuildsUnwrappedCSS behavior for the package.
func TestStyleElementAddRuleBuildsUnwrappedCSS(t *testing.T) {
	rule := NewStyleRule("body")
	rule.Props = StyleMap{
		"color":       "#111827",
		"font-family": "Arial, sans-serif",
	}

	style := NewStyleElement().AddRule(rule)
	style.Prepare()

	want := "<style>\n" +
		"body {\n" +
		"\tcolor: #111827;\n" +
		"\tfont-family: Arial, sans-serif;\n" +
		"}\n" +
		"</style>"
	if got := string(style.Bytes()); got != want {
		t.Fatalf("unexpected style element:\ngot  %q\nwant %q", got, want)
	}
}

func TestStyleUsesStyleMapDeclarations(t *testing.T) {
	style := &Style{
		Tags: []string{".card"},
		Props: StyleMap{
			"background": "#fff",
			"padding":    "1rem",
		},
	}
	style.Prepare()

	want := "<style>\n" +
		".card {\n" +
		"\tbackground: #fff;\n" +
		"\tpadding: 1rem;\n" +
		"}\n" +
		"</style>"
	if got := string(style.Bytes()); got != want {
		t.Fatalf("unexpected style output:\ngot  %q\nwant %q", got, want)
	}
}

// TestStyleElementAddUnwrapsStyle provides TestStyleElementAddUnwrapsStyle behavior for the package.
func TestStyleElementAddUnwrapsStyle(t *testing.T) {
	rule := NewStyle("body")
	rule.Props = StyleMap{
		"color": "#111827",
	}

	style := NewStyleElement().Add(rule)
	style.Prepare()

	want := "<style>\n" +
		"body {\n" +
		"\tcolor: #111827;\n" +
		"}\n" +
		"</style>"
	if got := string(style.Bytes()); got != want {
		t.Fatalf("unexpected style element:\ngot  %q\nwant %q", got, want)
	}
}

func TestStyleElementBuildsStatementAtRules(t *testing.T) {
	style := NewStyleElement().
		AddRule(NewCharsetRule("utf-8")).
		AddRule(NewImportRule("/base.css").Condition("screen"))
	style.Prepare()

	want := "<style>\n" +
		"@charset \"utf-8\";\n" +
		"\n" +
		"@import url(\"/base.css\") screen;\n" +
		"</style>"
	if got := string(style.Bytes()); got != want {
		t.Fatalf("unexpected at-rule style element:\ngot  %q\nwant %q", got, want)
	}
}

func TestFontFaceRuleUsesFontDescriptors(t *testing.T) {
	fontFace := NewFontFaceRule()
	fontFace.Props = StyleMap{
		"font-family": "Report",
		"font-weight": "700",
		"src":         `url("/fonts/report.woff2") format("woff2")`,
	}
	fontFace.Prepare()

	want := "\n@font-face {\n" +
		"\tfont-family: Report;\n" +
		"\tfont-weight: 700;\n" +
		"\tsrc: url(\"/fonts/report.woff2\") format(\"woff2\");\n" +
		"}\n"
	if got := string(fontFace.Bytes()); got != want {
		t.Fatalf("unexpected font-face rule:\ngot  %q\nwant %q", got, want)
	}
}

func TestRawCSSRuleRendersCustomCSS(t *testing.T) {
	rule := NewRawCSSRule("@layer base;")
	rule.Prepare()

	want := "\n@layer base;\n"
	if got := string(rule.Bytes()); got != want {
		t.Fatalf("unexpected raw CSS rule:\ngot  %q\nwant %q", got, want)
	}
}

func TestFontFeatureValuesRuleRendersRawBlockContent(t *testing.T) {
	rule := NewFontFeatureValuesRule("Report").Text("@styleset {\n\tswash: 1;\n}")
	rule.Prepare()

	want := "\n@font-feature-values Report {\n" +
		"\t@styleset {\n" +
		"\t\tswash: 1;\n" +
		"\t}\n" +
		"}\n"
	if got := string(rule.Bytes()); got != want {
		t.Fatalf("unexpected font-feature-values rule:\ngot  %q\nwant %q", got, want)
	}
}

func TestMediaRuleNestsStyleRules(t *testing.T) {
	rule := NewStyleRule(".card")
	rule.Props = StyleMap{"padding": "2rem"}

	media := NewMediaRule("(min-width: 800px)").AddRule(rule)
	media.Prepare()

	want := "\n@media (min-width: 800px) {\n" +
		"\t.card {\n" +
		"\t\tpadding: 2rem;\n" +
		"\t}\n" +
		"}\n"
	if got := string(media.Bytes()); got != want {
		t.Fatalf("unexpected media rule:\ngot  %q\nwant %q", got, want)
	}
}

func TestKeyframesRuleBuildsFrames(t *testing.T) {
	keyframes := NewKeyframesRule("fade").
		AddFrame("from", StyleMap{"opacity": "0"}).
		AddFrame("to", StyleMap{"opacity": "1"})
	keyframes.Prepare()

	want := "\n@keyframes fade {\n" +
		"\tfrom {\n" +
		"\t\topacity: 0;\n" +
		"\t}\n" +
		"\tto {\n" +
		"\t\topacity: 1;\n" +
		"\t}\n" +
		"}\n"
	if got := string(keyframes.Bytes()); got != want {
		t.Fatalf("unexpected keyframes rule:\ngot  %q\nwant %q", got, want)
	}
}

func TestStyleMapSupportsCustomProperties(t *testing.T) {
	rule := NewStyleRule(":root").Style(StyleMap{
		"--brand-color": "#2563eb",
		"color":         "var(--brand-color)",
	})
	rule.Prepare()

	want := "\n:root {\n" +
		"\t--brand-color: #2563eb;\n" +
		"\tcolor: var(--brand-color);\n" +
		"}\n"
	if got := string(rule.Bytes()); got != want {
		t.Fatalf("unexpected style map rule:\ngot  %q\nwant %q", got, want)
	}
}
