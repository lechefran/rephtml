package rephtml

import "testing"

func TestStyleElementAddRuleBuildsUnwrappedCSS(t *testing.T) {
	rule := NewStyleRule("body")
	rule.Props = CssProps{
		Color:      "#111827",
		FontFamily: "Arial, sans-serif",
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

func TestStyleUsesDefaultPropMap(t *testing.T) {
	style := &Style{
		Tags: []string{".card"},
		Props: CssProps{
			Background: "#fff",
			Padding:    "1rem",
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
