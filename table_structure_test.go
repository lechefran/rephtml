package rephtml

import "testing"

func TestTableStructure(t *testing.T) {
	row := func(s string) *Tr { return NewTr().AddTd(NewTd().Add(NewSpan().Text(s))) }

	tests := []struct {
		name  string
		table *Table
		want  string
	}{
		{
			name: "sections render in content-model order",
			table: NewTable().
				AddTfoot(NewTfoot().AddTr(row("f"))).
				AddTbody(NewTbody().AddTr(row("b"))).
				AddThead(NewThead().AddTr(row("h"))).
				AddColgroup(NewColgroup().Span(2)).
				AddCaption(NewCaption().Text("c")),
			want: `<table><caption>c</caption><colgroup span="2"></colgroup>` +
				`<thead><tr><td><span>h</span></td></tr></thead>` +
				`<tbody><tr><td><span>b</span></td></tr></tbody>` +
				`<tfoot><tr><td><span>f</span></td></tr></tfoot></table>`,
		},
		{
			name:  "direct rows render before tfoot",
			table: NewTable().AddTfoot(NewTfoot().AddTr(row("f"))).AddTr(row("r")),
			want: `<table><tr><td><span>r</span></td></tr>` +
				`<tfoot><tr><td><span>f</span></td></tr></tfoot></table>`,
		},
		{
			name: "colgroups append in order and nil is ignored",
			table: NewTable().
				AddColgroup(NewColgroup().Add(NewCol().Span(2))).
				AddColgroup(nil).
				AddColgroup(NewColgroup().AddClass("totals")),
			want: `<table><colgroup><col span="2"></colgroup><colgroup class="totals"></colgroup></table>`,
		},
		{
			name:  "a colgroup does not replace legacy rows",
			table: NewTable().AddColgroup(NewColgroup()).Headers([]string{"h"}).AddRow([]string{"r"}),
			want:  `<table><colgroup></colgroup><tr><th>h</th></tr><tr><td>r</td></tr></table>`,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.table.HTML(); got != tt.want {
				t.Fatalf("got %q, want %q", got, tt.want)
			}
		})
	}
}

func TestTableColgroupErrorsReachDocument(t *testing.T) {
	cg := NewColgroup().Add(NewTime().Datetime("tomorrow"))
	doc := NewHtmlFile().AddToBody(NewTable().AddColgroup(cg))

	if doc.Err() == nil {
		t.Fatal("Err() = nil, want the colgroup child's error")
	}
}
