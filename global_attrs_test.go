package rephtml

import (
	"reflect"
	"sort"
	"strings"
	"testing"
)

// TestEveryElementSupportsIdAndClass sets id and class through the public
// setters on every element in the formatting cases, and checks that both reach
// the output, escaped. Types that are not HTML elements with attributes have no
// Id method and are listed explicitly, so a new element that forgets them fails.
func TestEveryElementSupportsIdAndClass(t *testing.T) {
	withoutAttrs := map[string]bool{
		"Comment": true, "Style": true, "StyleRule": true, "RawCSSRule": true,
		"CharsetRule": true, "ImportRule": true, "FontFaceRule": true,
		"FontFeatureValuesRule": true, "MediaRule": true, "KeyframeBlock": true,
		"KeyframesRule": true,
	}

	var skipped []string
	for _, group := range [][]formatElementCase{
		standaloneFormatCases(),
		headFormatCases(),
		bodyFormatCases(),
		styleFormatCases(),
	} {
		for _, tc := range group {
			v := reflect.ValueOf(tc.element)
			id, class := v.MethodByName("Id"), v.MethodByName("Class")
			if !id.IsValid() || !class.IsValid() {
				skipped = append(skipped, tc.name)
				continue
			}
			t.Run(tc.name, func(t *testing.T) {
				id.Call([]reflect.Value{reflect.ValueOf(`a"b<`)})
				class.Call([]reflect.Value{reflect.ValueOf([]string{"c1", "<c2>"})})

				got := tc.element.HTML()
				for _, want := range []string{` id="a&#34;b&lt;"`, ` class="c1 &lt;c2&gt;"`} {
					if !strings.Contains(got, want) {
						t.Fatalf("%s.HTML() = %q, want it to contain %q", tc.name, got, want)
					}
				}
			})
		}
	}

	sort.Strings(skipped)
	for _, name := range skipped {
		if !withoutAttrs[name] {
			t.Errorf("%s has no Id or Class method", name)
		}
	}
}

func TestIdAndClassPosition(t *testing.T) {
	tests := []struct {
		name string
		got  string
		want string
	}{
		{
			name: "id, class and style are written together",
			got:  NewDiv().AddStyle("color", "red").AddClass("a").Id("main").HTML(),
			want: `<div id="main" class="a" style="color: red;"></div>`,
		},
		{
			name: "AddClasses appends and Class replaces",
			got:  NewSpan().AddClass("x").AddClasses([]string{"y", "z"}).Text("t").HTML(),
			want: `<span class="x y z">t</span>`,
		},
		{
			name: "attributes sit where the element writes its style",
			got:  NewAnchor().Link("/home").Id("home").Text("Home").HTML(),
			want: `<a id="home" href="/home">Home</a>`,
		},
		{
			name: "Input keeps id at its existing position",
			got:  NewInput().Type("text").Name("q").Id("q").AddClass("wide").Required(true).HTML(),
			want: `<input type="text" name="q" id="q" required class="wide">`,
		},
		{
			name: "Datalist keeps its existing layout",
			got:  NewDatalist().Id("opts").AddStyle("color", "red").HTML(),
			want: `<datalist id="opts" style="color: red;"></datalist>`,
		},
		{
			name: "Template keeps its existing layout",
			got:  NewTemplate().Id("row").HTML(),
			want: `<template id="row"></template>`,
		},
		{
			name: "table elements are unchanged",
			got:  NewTd().Id("cell").AddClass("num").AddStyle("color", "red").HTML(),
			want: `<td id="cell" class="num" style="color: red;"></td>`,
		},
		{
			name: "unset id and class add nothing",
			got:  NewDiv().Class(nil).Id("").HTML(),
			want: `<div></div>`,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if tt.got != tt.want {
				t.Fatalf("got %q, want %q", tt.got, tt.want)
			}
		})
	}
}

func TestClassCopiesCallerSlice(t *testing.T) {
	classes := []string{"a", "b"}
	div := NewDiv().Class(classes)
	classes[0] = "changed"

	if got, want := div.HTML(), `<div class="a b"></div>`; got != want {
		t.Fatalf("got %q, want %q", got, want)
	}
}

func TestMergedHeadAndBodyKeepIdAndClass(t *testing.T) {
	doc := NewHtmlFile().
		AddToHead(NewHead().Id("h").AddClass("hc")).
		AddToBody(NewBody().OnLoad("init()").Id("b").AddClass("bc").Add(NewP().Text("x")))

	got, err := doc.RenderString()
	if err != nil {
		t.Fatal(err)
	}
	want := `<!DOCTYPE html><html><head id="h" class="hc"></head>` +
		`<body onload="init()" id="b" class="bc"><p>x</p></body></html>`
	if got != want {
		t.Fatalf("got %q, want %q", got, want)
	}
}

func TestHtmlFileIdAndClass(t *testing.T) {
	classes := []string{"a", "<b>"}
	doc := NewHtmlFile().Lang("en").Id(`r"1`).Class(classes).
		AddClass("c").AddClasses([]string{"d"}).AddStyle("color", "red")
	classes[0] = "changed"

	got, err := doc.RenderString()
	if err != nil {
		t.Fatal(err)
	}
	want := `<!DOCTYPE html><html lang="en" id="r&#34;1" class="a &lt;b&gt; c d" style="color: red;"></html>`
	if got != want {
		t.Fatalf("got %q, want %q", got, want)
	}

	if got, err := NewHtmlFile().Lang("en").Class(nil).RenderString(); err != nil ||
		got != `<!DOCTYPE html><html lang="en"></html>` {
		t.Fatalf("unset id and class: got %q, %v", got, err)
	}
}

func TestAddText(t *testing.T) {
	tests := []struct {
		name string
		got  string
		want string
	}{
		{
			name: "text element appends",
			got:  NewP().Text("a").AddText(" & b").HTML(),
			want: `<p>a &amp; b</p>`,
		},
		{
			name: "Text still replaces on a text element",
			got:  NewP().AddText("a").Text("b").HTML(),
			want: `<p>b</p>`,
		},
		{
			name: "content element appends after children",
			got:  NewLabel().Text("a").Add(NewB().Text("b")).AddText("<c>").HTML(),
			want: `<label>a<b>b</b>&lt;c&gt;</label>`,
		},
		{
			name: "Caption appends",
			got:  NewCaption().Text("Q1").AddText(" totals").HTML(),
			want: `<caption>Q1 totals</caption>`,
		},
		{
			name: "Comment appends and stays sanitised",
			got:  NewComment().Text("a").AddText("--> b").HTML(),
			want: `<!--a- -> b-->`,
		},
		{
			name: "StyleElement appends and stays defused",
			got:  NewStyleElement().Text("p{}").AddText("</style>").HTML(),
			want: `<style>p{}\3c /style></style>`,
		},
		{
			name: "Script appends and stays defused",
			got:  NewScript().Text("a;").AddText("</script>").HTML(),
			want: `<script>a;<\/script></script>`,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if tt.got != tt.want {
				t.Fatalf("got %q, want %q", tt.got, tt.want)
			}
		})
	}
}
