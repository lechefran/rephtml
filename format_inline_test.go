package rephtml

import "testing"

// The tests in this file cover inline content in formatted output. Inline
// elements and the text around them are whitespace-sensitive: a line break
// between two of them renders as a space. The formatter keeps each run of
// inline content on one line, exactly as rendered, so formatting never changes
// how a page displays.

func TestFormatterKeepsInlineContentOnOneLine(t *testing.T) {
	tests := []struct {
		name string
		in   string
		want []string
	}{
		{
			name: "adjacent inline elements gain no space",
			in:   `<div><b>a</b><i>b</i></div>`,
			want: []string{"<div>", "\t<b>a</b><i>b</i>", "</div>"},
		},
		{
			name: "whitespace between inline elements is kept",
			in:   `<div><b>a</b> <i>b</i></div>`,
			want: []string{"<div>", "\t<b>a</b> <i>b</i>", "</div>"},
		},
		{
			name: "text around an inline element stays attached",
			in:   `<div>Hello <b>x</b>.</div>`,
			want: []string{"<div>", "\tHello <b>x</b>.", "</div>"},
		},
		{
			name: "nested inline elements stay together",
			in:   `<div><a href="/"><b>x</b></a></div>`,
			want: []string{"<div>", "\t<a href=\"/\"><b>x</b></a>", "</div>"},
		},
		{
			name: "a line break stays inline",
			in:   `<p>a<br>b</p>`,
			want: []string{"<p>", "\ta<br>b", "</p>"},
		},
		{
			name: "a comment between inline elements stays inline",
			in:   `<div><b>a</b><!--c--><i>b</i></div>`,
			want: []string{"<div>", "\t<b>a</b><!--c--><i>b</i>", "</div>"},
		},
		{
			name: "a textarea stays inline in its label",
			in:   `<div><label>Bio <textarea>x</textarea></label></div>`,
			want: []string{"<div>", "\t<label>Bio <textarea>x</textarea></label>", "</div>"},
		},
		{
			name: "whitespace at the edges of an inline run is dropped",
			in:   `<div> <b>x</b> </div>`,
			want: []string{"<div>", "\t<b>x</b>", "</div>"},
		},
		{
			name: "whitespace between block elements is still dropped",
			in:   `<div><p>a</p> <p>b</p></div>`,
			want: []string{"<div>", "\t<p>a</p>", "\t<p>b</p>", "</div>"},
		},
		{
			name: "a single inline element is unchanged",
			in:   `<div><span>x</span></div>`,
			want: []string{"<div>", "\t<span>x</span>", "</div>"},
		},
		{
			name: "block elements inside an inline element keep their own lines",
			in:   `<div><a href="/"><div>x</div></a></div>`,
			want: []string{"<div>", "\t<a href=\"/\">", "\t\t<div>x</div>", "\t</a>", "</div>"},
		},
		{
			name: "an inline run that closes an earlier element is outdented",
			in:   `<div><a><p>x</p>t<b>y</b></a></div>`,
			want: []string{"<div>", "\t<a>", "\t\t<p>x</p>", "\tt<b>y</b></a>", "</div>"},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := formatRenderedForTest([]byte(tt.in))
			if want := formattedLines(tt.want...); got != want {
				t.Fatalf("formatted %s:\ngot:\n%s\nwant:\n%s", tt.in, got, want)
			}
		})
	}
}

func TestRenderFormattedKeepsInlineSpacing(t *testing.T) {
	doc := NewHtmlFile().AddToBody(NewFieldset().Add(
		NewLegend().Text("Contact ").Add(NewStrong().Text("required")).AddText(" fields"),
	))

	got, err := doc.RenderFormattedString()
	if err != nil {
		t.Fatal(err)
	}
	want := formattedLines(
		"<!DOCTYPE html>",
		"<html>",
		"\t<body>",
		"\t\t<fieldset>",
		"\t\t\t<legend>",
		"\t\t\t\tContact <strong>required</strong> fields",
		"\t\t\t</legend>",
		"\t\t</fieldset>",
		"\t</body>",
		"</html>",
	)
	if got != want {
		t.Fatalf("got:\n%s\nwant:\n%s", got, want)
	}
}
