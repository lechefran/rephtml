package rephtml

import "testing"

// TestTextReplaces checks that Text sets the text on every kind of text
// element, replacing earlier text while keeping child elements in place.
func TestTextReplaces(t *testing.T) {
	tests := []struct {
		name string
		got  string
		want string
	}{
		{
			name: "content element replaces earlier text",
			got:  NewLabel().Text("a").Text("b").HTML(),
			want: `<label>b</label>`,
		},
		{
			name: "content element replaces AddText text",
			got:  NewButton().AddText("a").AddText("b").Text("c").HTML(),
			want: `<button>c</button>`,
		},
		{
			name: "children keep their place and text takes the first text slot",
			got: NewLabel().Text("a").Add(NewB().Text("x")).AddText("y").
				Text("<z>").HTML(),
			want: `<label>&lt;z&gt;<b>x</b></label>`,
		},
		{
			name: "text goes after children when there was none",
			got:  NewLabel().Add(NewInput().Type("checkbox")).Text("Agree").HTML(),
			want: `<label><input type="checkbox">Agree</label>`,
		},
		{
			name: "empty text clears the text and keeps children",
			got:  NewLabel().Text("a").Add(NewB().Text("x")).Text("").HTML(),
			want: `<label><b>x</b></label>`,
		},
		{
			name: "AddText after Text appends",
			got:  NewOption().Text("a").AddText("b").HTML(),
			want: `<option>ab</option>`,
		},
		{
			name: "StyleElement replaces text and keeps rules",
			got: NewStyleElement().Text("a{}").AddRule(NewRawCSSRule("b{}")).
				AddText("c{}").Text("d{}").HTML(),
			want: "<style>d{}\nb{}\n</style>",
		},
		{
			name: "StyleElement text stays defused",
			got:  NewStyleElement().Text("a{}").Text("</style>").HTML(),
			want: `<style>\3c /style></style>`,
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

// TestReplaceTextDoesNotLeakDroppedChildren checks that the slots freed by
// dropped text children are cleared, so the backing array keeps no reference.
func TestReplaceTextDoesNotLeakDroppedChildren(t *testing.T) {
	contents := []Element{escapedText("a"), escapedText("b"), escapedText("c")}
	got := replaceText(contents, escapedText("d"))
	if len(got) != 1 || got[0] != escapedText("d") {
		t.Fatalf("got %v, want [d]", got)
	}
	for i, e := range contents[1:] {
		if e != nil {
			t.Fatalf("slot %d still holds %v", i+1, e)
		}
	}
}
