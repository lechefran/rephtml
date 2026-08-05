package rephtml

import "bytes"

// Table represents the Table component or supporting type.
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

// AddHeader sets the addheader value on the Table component.
func (t *Table) AddHeader(s string) *Table {
	t.headers = append(t.headers, s)
	return t
}

// AddHeaders sets the addheaders value on the Table component.
func (t *Table) AddHeaders(s []string) *Table {
	t.headers = append(t.headers, s...)
	return t
}

// Headers sets the headers value on the Table component.
func (t *Table) Headers(s []string) *Table {
	t.headers = s
	return t
}

// AddRow sets the addrow value on the Table component.
func (t *Table) AddRow(s []string) *Table {
	t.rows = append(t.rows, s)
	return t
}

// AddRows sets the addrows value on the Table component.
func (t *Table) AddRows(s [][]string) *Table {
	t.rows = append(t.rows, s...)
	return t
}

// Rows sets the rows value on the Table component.
func (t *Table) Rows(s [][]string) *Table {
	t.rows = s
	return t
}

// AddCaption sets the addcaption value on the Table component.
func (t *Table) AddCaption(c *Caption) *Table {
	t.caption = c
	return t
}

// AddThead sets the addthead value on the Table component.
func (t *Table) AddThead(th *Thead) *Table {
	t.thead = th
	return t
}

// AddTbody sets the addtbody value on the Table component.
func (t *Table) AddTbody(tb *Tbody) *Table {
	t.tbody = tb
	return t
}

// AddTfoot sets the addtfoot value on the Table component.
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
	tg.attr("id", t.id)
	tg.classAttr(t.class)
	tg.styleAttr(t.style)
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
				buf.WriteString("<th>" + escapeText(h) + "</th>")
			}
			buf.WriteString("</tr>")
		}

		// write rows
		for i := 0; i < len(t.rows); i++ {
			buf.WriteString("<tr>")
			for j := 0; j < len(t.rows[i]); j++ {
				buf.WriteString("<td>" + escapeText(t.rows[i][j]) + "</td>")
			}
			buf.WriteString("</tr>")
		}
	}

	tg.end()
}

// Thead represents the Thead component or supporting type.
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

// AddTr sets the addtr value on the Thead component.
func (th *Thead) AddTr(tr *Tr) *Thead {
	th.contents = appendElement(th.contents, tr)
	return th
}

// renderTo writes the Thead component's HTML to buf.
func (th *Thead) renderTo(buf *bytes.Buffer) {
	tg := startTag(buf, "thead")
	tg.attr("id", th.id)
	tg.classAttr(th.class)
	tg.styleAttr(th.style)
	tg.children(th.contents)
}

// Tbody represents the Tbody component or supporting type.
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

// AddTr sets the addtr value on the Tbody component.
func (tb *Tbody) AddTr(tr *Tr) *Tbody {
	tb.contents = appendElement(tb.contents, tr)
	return tb
}

// renderTo writes the Tbody component's HTML to buf.
func (tb *Tbody) renderTo(buf *bytes.Buffer) {
	tg := startTag(buf, "tbody")
	tg.attr("id", tb.id)
	tg.classAttr(tb.class)
	tg.styleAttr(tb.style)
	tg.children(tb.contents)
}

// Tfoot represents the Tfoot component or supporting type.
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

// AddTr sets the addtr value on the Tfoot component.
func (tf *Tfoot) AddTr(tr *Tr) *Tfoot {
	tf.contents = appendElement(tf.contents, tr)
	return tf
}

// renderTo writes the Tfoot component's HTML to buf.
func (tf *Tfoot) renderTo(buf *bytes.Buffer) {
	tg := startTag(buf, "tfoot")
	tg.attr("id", tf.id)
	tg.classAttr(tf.class)
	tg.styleAttr(tf.style)
	tg.children(tf.contents)
}

// Caption represents the Caption component or supporting type.
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

// Text sets or appends text content on the Caption component.
func (c *Caption) Text(text string) *Caption {
	c.text = text
	return c
}

// renderTo writes the Caption component's HTML to buf.
func (c *Caption) renderTo(buf *bytes.Buffer) {
	tg := startTag(buf, "caption")
	tg.attr("id", c.id)
	tg.classAttr(c.class)
	tg.styleAttr(c.style)
	tg.text(c.text)
}

// Col represents the Col component or supporting type.
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

// Span sets the span value on the Col component.
func (col *Col) Span(s int) *Col {
	col.span = s
	return col
}

// renderTo writes the Col component's HTML to buf.
func (col *Col) renderTo(buf *bytes.Buffer) {
	tg := startTag(buf, "col")
	tg.attr("id", col.id)
	tg.classAttr(col.class)
	tg.styleAttr(col.style)
	tg.intAttr("span", col.span)
	tg.void()
}

// Colgroup represents the Colgroup component or supporting type.
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

// Span sets the span value on the Colgroup component.
func (cg *Colgroup) Span(s int) *Colgroup {
	cg.span = s
	return cg
}

// renderTo writes the Colgroup component's HTML to buf.
func (cg *Colgroup) renderTo(buf *bytes.Buffer) {
	tg := startTag(buf, "colgroup")
	tg.attr("id", cg.id)
	tg.classAttr(cg.class)
	tg.styleAttr(cg.style)
	tg.intAttr("span", cg.span)
	tg.children(cg.contents)
}

// Tr represents the Tr component or supporting type.
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

// AddTh sets the addth value on the Tr component.
func (tr *Tr) AddTh(th *Th) *Tr {
	tr.contents = appendElement(tr.contents, th)
	return tr
}

// AddTd sets the addtd value on the Tr component.
func (tr *Tr) AddTd(td *Td) *Tr {
	tr.contents = appendElement(tr.contents, td)
	return tr
}

// renderTo writes the Tr component's HTML to buf.
func (tr *Tr) renderTo(buf *bytes.Buffer) {
	tg := startTag(buf, "tr")
	tg.attr("id", tr.id)
	tg.classAttr(tr.class)
	tg.styleAttr(tr.style)
	tg.children(tr.contents)
}

// Td represents the Td component or supporting type.
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

// Colspan sets the colspan value on the Td component.
func (td *Td) Colspan(c int) *Td {
	td.colspan = c
	return td
}

// Rowspan sets the rowspan value on the Td component.
func (td *Td) Rowspan(r int) *Td {
	td.rowspan = r
	return td
}

// renderTo writes the Td component's HTML to buf.
func (td *Td) renderTo(buf *bytes.Buffer) {
	tg := startTag(buf, "td")
	tg.attr("id", td.id)
	tg.classAttr(td.class)
	tg.styleAttr(td.style)
	tg.intAttr("colspan", td.colspan)
	tg.intAttr("rowspan", td.rowspan)
	tg.children(td.contents)
}

// Th represents the Th component or supporting type.
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

// Colspan sets the colspan value on the Th component.
func (th *Th) Colspan(c int) *Th {
	th.colspan = c
	return th
}

// Rowspan sets the rowspan value on the Th component.
func (th *Th) Rowspan(r int) *Th {
	th.rowspan = r
	return th
}

// Scope sets the scope value on the Th component.
func (th *Th) Scope(s string) *Th {
	th.scope = s
	return th
}

// renderTo writes the Th component's HTML to buf.
func (th *Th) renderTo(buf *bytes.Buffer) {
	tg := startTag(buf, "th")
	tg.attr("id", th.id)
	tg.classAttr(th.class)
	tg.styleAttr(th.style)
	tg.intAttr("colspan", th.colspan)
	tg.intAttr("rowspan", th.rowspan)
	tg.attr("scope", th.scope)
	tg.children(th.contents)
}
