package rephtml

import "bytes"

// Table represents the HTML <table> element, tabular data in rows and columns.
type Table struct {
	bodyElement
	tabular[*Table]
	headers []string
	rows    [][]string
	caption *Caption
	thead   *Thead
	tbody   *Tbody
	tfoot   *Tfoot
	trs     []*Tr
}

// NewTable creates a new Table component.
func NewTable() *Table {
	v := &Table{}
	v.init(v)
	return v
}

// AddHeader appends one header cell to the table's first row.
func (t *Table) AddHeader(s string) *Table {
	t.headers = append(t.headers, s)
	return t
}

// AddHeaders appends several header cells to the table's first row.
func (t *Table) AddHeaders(s []string) *Table {
	t.headers = append(t.headers, s...)
	return t
}

// Headers sets the header cells of the table's first row, copying the caller's
// slice.
func (t *Table) Headers(s []string) *Table {
	t.headers = cloneStrings(s)
	return t
}

// AddRow appends one row of cells, copying the caller's slice.
func (t *Table) AddRow(s []string) *Table {
	t.rows = append(t.rows, cloneStrings(s))
	return t
}

// AddRows appends several rows of cells, copying the caller's slices.
func (t *Table) AddRows(s [][]string) *Table {
	t.rows = append(t.rows, cloneRows(s)...)
	return t
}

// Rows replaces the rows of cells, copying the caller's slices.
func (t *Table) Rows(s [][]string) *Table {
	t.rows = cloneRows(s)
	return t
}

// AddCaption sets the table's caption, replacing any earlier one.
func (t *Table) AddCaption(c *Caption) *Table {
	t.caption = c
	return t
}

// AddThead sets the table's thead, replacing any earlier one.
func (t *Table) AddThead(th *Thead) *Table {
	t.thead = th
	return t
}

// AddTbody sets the table's tbody, replacing any earlier one.
func (t *Table) AddTbody(tb *Tbody) *Table {
	t.tbody = tb
	return t
}

// AddTfoot sets the table's tfoot, replacing any earlier one.
func (t *Table) AddTfoot(tf *Tfoot) *Table {
	t.tfoot = tf
	return t
}

// AddTr appends one row directly to the table, ignoring a nil row.
func (t *Table) AddTr(tr *Tr) *Table {
	if tr != nil {
		t.trs = append(t.trs, tr)
	}
	return t
}

// renderTo writes the Table component's HTML to buf.
func (t *Table) renderTo(buf *bytes.Buffer) {
	tg := startTag(buf, "table")
	t.attrs(tg)
	tg.open()

	// write the structured children if present
	if t.caption != nil {
		t.caption.renderTo(buf)
	}
	if t.thead != nil {
		t.thead.renderTo(buf)
	}
	if t.tbody != nil {
		t.tbody.renderTo(buf)
	}
	if t.tfoot != nil {
		t.tfoot.renderTo(buf)
	}

	// write direct tr elements if present
	for _, tr := range t.trs {
		tr.renderTo(buf)
	}

	// write legacy header and rows if no structured elements are used
	if t.thead == nil && t.tbody == nil && t.tfoot == nil && len(t.trs) == 0 {
		// write header
		if len(t.headers) > 0 {
			buf.WriteString("<tr>")
			for _, h := range t.headers {
				startTag(buf, "th").text(h)
			}
			buf.WriteString("</tr>")
		}

		// write rows
		for _, row := range t.rows {
			buf.WriteString("<tr>")
			for _, cell := range row {
				startTag(buf, "td").text(cell)
			}
			buf.WriteString("</tr>")
		}
	}

	tg.end()
}

// Thead represents the HTML <thead> element, the header rows of a table.
type Thead struct {
	bodyElement
	tabularNode[*Thead]
}

// NewThead creates a new Thead component.
func NewThead() *Thead {
	v := &Thead{}
	v.init(v)
	return v
}

// AddTr appends a row to the table head.
func (th *Thead) AddTr(tr *Tr) *Thead {
	th.contents = appendElement(th.contents, tr)
	return th
}

// renderTo writes the Thead component's HTML to buf.
func (th *Thead) renderTo(buf *bytes.Buffer) {
	tg := startTag(buf, "thead")
	th.attrs(tg)
	tg.children(th.contents)
}

// Tbody represents the HTML <tbody> element, the body rows of a table.
type Tbody struct {
	bodyElement
	tabularNode[*Tbody]
}

// NewTbody creates a new Tbody component.
func NewTbody() *Tbody {
	v := &Tbody{}
	v.init(v)
	return v
}

// AddTr appends a row to the table body.
func (tb *Tbody) AddTr(tr *Tr) *Tbody {
	tb.contents = appendElement(tb.contents, tr)
	return tb
}

// renderTo writes the Tbody component's HTML to buf.
func (tb *Tbody) renderTo(buf *bytes.Buffer) {
	tg := startTag(buf, "tbody")
	tb.attrs(tg)
	tg.children(tb.contents)
}

// Tfoot represents the HTML <tfoot> element, the footer rows of a table.
type Tfoot struct {
	bodyElement
	tabularNode[*Tfoot]
}

// NewTfoot creates a new Tfoot component.
func NewTfoot() *Tfoot {
	v := &Tfoot{}
	v.init(v)
	return v
}

// AddTr appends a row to the table foot.
func (tf *Tfoot) AddTr(tr *Tr) *Tfoot {
	tf.contents = appendElement(tf.contents, tr)
	return tf
}

// renderTo writes the Tfoot component's HTML to buf.
func (tf *Tfoot) renderTo(buf *bytes.Buffer) {
	tg := startTag(buf, "tfoot")
	tf.attrs(tg)
	tg.children(tf.contents)
}

// Caption represents the HTML <caption> element, the title of a table.
type Caption struct {
	bodyElement
	tabular[*Caption]
	text string
}

// NewCaption creates a new Caption component.
func NewCaption() *Caption {
	v := &Caption{}
	v.init(v)
	return v
}

// Text sets the caption text, replacing any earlier text.
func (c *Caption) Text(text string) *Caption {
	c.text = text
	return c
}

// AddText appends to the caption text.
func (c *Caption) AddText(text string) *Caption {
	c.text += text
	return c
}

// renderTo writes the Caption component's HTML to buf.
func (c *Caption) renderTo(buf *bytes.Buffer) {
	tg := startTag(buf, "caption")
	c.attrs(tg)
	tg.text(c.text)
}

// Col represents the HTML <col> element, a column in a colgroup, used to share properties across a column.
type Col struct {
	bodyElement
	tabular[*Col]
	span int
}

// NewCol creates a new Col component.
func NewCol() *Col {
	v := &Col{}
	v.init(v)
	return v
}

// Span sets the span attribute.
func (col *Col) Span(s int) *Col {
	col.span = s
	return col
}

// renderTo writes the Col component's HTML to buf.
func (col *Col) renderTo(buf *bytes.Buffer) {
	tg := startTag(buf, "col")
	col.attrs(tg)
	tg.intAttr("span", col.span)
	tg.void()
}

// Colgroup represents the HTML <colgroup> element, a group of columns in a table.
type Colgroup struct {
	bodyElement
	tabularNode[*Colgroup]
	span int
}

// NewColgroup creates a new Colgroup component.
func NewColgroup() *Colgroup {
	v := &Colgroup{}
	v.init(v)
	return v
}

// Span sets the span attribute.
func (cg *Colgroup) Span(s int) *Colgroup {
	cg.span = s
	return cg
}

// renderTo writes the Colgroup component's HTML to buf.
func (cg *Colgroup) renderTo(buf *bytes.Buffer) {
	tg := startTag(buf, "colgroup")
	cg.attrs(tg)
	tg.intAttr("span", cg.span)
	tg.children(cg.contents)
}

// Tr represents the HTML <tr> element, a row of cells in a table.
type Tr struct {
	bodyElement
	tabularNode[*Tr]
}

// NewTr creates a new Tr component.
func NewTr() *Tr {
	v := &Tr{}
	v.init(v)
	return v
}

// AddTh appends a header cell to the row.
func (tr *Tr) AddTh(th *Th) *Tr {
	tr.contents = appendElement(tr.contents, th)
	return tr
}

// AddTd appends a data cell to the row.
func (tr *Tr) AddTd(td *Td) *Tr {
	tr.contents = appendElement(tr.contents, td)
	return tr
}

// renderTo writes the Tr component's HTML to buf.
func (tr *Tr) renderTo(buf *bytes.Buffer) {
	tg := startTag(buf, "tr")
	tr.attrs(tg)
	tg.children(tr.contents)
}

// Td represents the HTML <td> element, a data cell in a table.
type Td struct {
	bodyElement
	tabularNode[*Td]
	colspan int
	rowspan int
}

// NewTd creates a new Td component.
func NewTd() *Td {
	v := &Td{}
	v.init(v)
	return v
}

// Colspan sets the colspan attribute.
func (td *Td) Colspan(c int) *Td {
	td.colspan = c
	return td
}

// Rowspan sets the rowspan attribute.
func (td *Td) Rowspan(r int) *Td {
	td.rowspan = r
	return td
}

// renderTo writes the Td component's HTML to buf.
func (td *Td) renderTo(buf *bytes.Buffer) {
	tg := startTag(buf, "td")
	td.attrs(tg)
	tg.intAttr("colspan", td.colspan)
	tg.intAttr("rowspan", td.rowspan)
	tg.children(td.contents)
}

// Th represents the HTML <th> element, a header cell in a table.
type Th struct {
	bodyElement
	tabularNode[*Th]
	colspan int
	rowspan int
	scope   string
}

// NewTh creates a new Th component.
func NewTh() *Th {
	v := &Th{}
	v.init(v)
	return v
}

// Colspan sets the colspan attribute.
func (th *Th) Colspan(c int) *Th {
	th.colspan = c
	return th
}

// Rowspan sets the rowspan attribute.
func (th *Th) Rowspan(r int) *Th {
	th.rowspan = r
	return th
}

// Scope sets the scope attribute.
func (th *Th) Scope(s string) *Th {
	th.scope = s
	return th
}

// renderTo writes the Th component's HTML to buf.
func (th *Th) renderTo(buf *bytes.Buffer) {
	tg := startTag(buf, "th")
	th.attrs(tg)
	tg.intAttr("colspan", th.colspan)
	tg.intAttr("rowspan", th.rowspan)
	tg.attr("scope", th.scope)
	tg.children(th.contents)
}
