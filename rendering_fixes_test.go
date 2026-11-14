package rephtml

import (
	"strings"
	"testing"
)

func TestZeroAndNegativeNumericAttributesAreEmitted(t *testing.T) {
	tests := []struct {
		name    string
		element Element
		want    string
	}{
		{"ol start 0", NewOl().Start(0), `<ol start="0"></ol>`},
		{"ol start negative", NewOl().Start(-3), `<ol start="-3"></ol>`},
		{"ol start unset", NewOl(), `<ol></ol>`},
		{"li value 0", NewLi().Value(0), `<li value="0"></li>`},
		{"li value negative", NewLi().Value(-1), `<li value="-1"></li>`},
		{"li value unset", NewLi(), `<li></li>`},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.element.HTML(); got != tt.want {
				t.Errorf("HTML() = %q, want %q", got, tt.want)
			}
		})
	}
}

func TestBdoDir(t *testing.T) {
	got := NewBdo().Dir("rtl").Text("abc").HTML()
	if want := `<bdo dir="rtl">abc</bdo>`; got != want {
		t.Errorf("HTML() = %q, want %q", got, want)
	}
	if got, want := NewBdo().Text("abc").HTML(), `<bdo>abc</bdo>`; got != want {
		t.Errorf("HTML() without dir = %q, want %q", got, want)
	}
}

func TestLinkAttributesAreLowerCase(t *testing.T) {
	got := NewLink().Rel("stylesheet").Href("a.css").CrossOrigin("anonymous").ReferrerPolicy("no-referrer").HTML()
	want := `<link rel="stylesheet" href="a.css" crossorigin="anonymous" referrerpolicy="no-referrer">`
	if got != want {
		t.Errorf("HTML() = %q, want %q", got, want)
	}
}

func TestInlineStyleSkipsEmptyDeclarations(t *testing.T) {
	tests := []struct {
		name  string
		style StyleMap
		want  string
	}{
		{"empty value skipped", StyleMap{"color": "red", "margin": ""}, `<div style="color: red;"></div>`},
		{"empty property skipped", StyleMap{"": "red", "color": "blue"}, `<div style="color: blue;"></div>`},
		{"all empty writes no attribute", StyleMap{"color": ""}, `<div></div>`},
		{"sorted and spaced", StyleMap{"b": "2", "a": "1"}, `<div style="a: 1; b: 2;"></div>`},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewDiv().Style(tt.style).HTML(); got != tt.want {
				t.Errorf("HTML() = %q, want %q", got, tt.want)
			}
		})
	}
}

func TestTableCopiesCallerSlices(t *testing.T) {
	headers := []string{"A", "B"}
	row := []string{"1", "2"}
	rows := [][]string{{"3", "4"}}
	more := [][]string{{"5", "6"}}

	table := NewTable().Headers(headers).AddRow(row).Rows(rows).AddRows(more).AddRow(row)
	want := table.HTML()

	// Mutating everything the caller handed over must not change the output.
	headers[0], row[0], rows[0][0], more[0][0] = "X", "X", "X", "X"
	rows[0] = []string{"X"}

	if got := table.HTML(); got != want {
		t.Errorf("table changed after the caller mutated its slices:\n got %q\nwant %q", got, want)
	}
}

func TestTableLegacyCellsAreEscaped(t *testing.T) {
	got := NewTable().Headers([]string{"<h>"}).AddRow([]string{"a & b", "<td>"}).HTML()
	want := `<table><tr><th>&lt;h&gt;</th></tr><tr><td>a &amp; b</td><td>&lt;td&gt;</td></tr></table>`
	if got != want {
		t.Errorf("HTML() = %q, want %q", got, want)
	}
}

func TestTimeDatetimeErrorQuotesValue(t *testing.T) {
	err := NewTime().Datetime("not a date\n").Err()
	if err == nil || !strings.Contains(err.Error(), `"not a date\n"`) {
		t.Errorf("Err() = %v, want the value quoted", err)
	}
}

func TestHtmlFileAddIgnoresNilChildren(t *testing.T) {
	var nilElement Element
	doc := NewHtmlFile().Add(nilElement).Add(nilDiv()).Add(NewDiv())
	if got, want := doc.contents, 1; len(got) != want {
		t.Errorf("len(contents) = %d, want %d", len(got), want)
	}
}
