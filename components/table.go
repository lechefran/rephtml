package rephtml

import (
	"bytes"
)

// Table represents the Table component or supporting type.
type Table struct {
	buf     bytes.Buffer
	class   []string
	headers []string
	id      string
	rows    [][]string
	style   map[string]string // rewrite to make use of CssProps?
	caption *Caption
	thead   *Thead
	tbody   *Tbody
	tfoot   *Tfoot
	trs     []*Tr
}

// NewTable creates a new Table component.
func NewTable() *Table {
	return &Table{
		style: make(map[string]string),
	}
}

// AddClass sets the addclass value on the Table component.
func (t *Table) AddClass(s string) *Table {
	t.class = append(t.class, s)
	return t
}

// AddClasses sets the addclasses value on the Table component.
func (t *Table) AddClasses(s []string) *Table {
	t.class = append(t.class, s...)
	return t
}

// Class sets the class value on the Table component.
func (t *Table) Class(s []string) *Table {
	t.class = cloneStrings(s)
	return t
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

// AddId sets the addid value on the Table component.
func (t *Table) AddId(s string) *Table {
	t.id = s
	return t
}

// Id sets the id value on the Table component.
func (t *Table) Id(s string) *Table {
	t.id = s
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

// AddStyle adds one inline CSS declaration to the Table component.
func (t *Table) AddStyle(k, v string) *Table {
	t.style[k] = v
	return t
}

// AddStyles adds multiple inline CSS declarations to the Table component.
func (t *Table) AddStyles(m map[string]string) *Table {
	for k, v := range m {
		t.style[k] = v
	}
	return t
}

// Styles replaces the inline CSS declarations on the Table component.
func (t *Table) Styles(m map[string]string) *Table {
	t.style = cloneStyleMap(m)
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

// AddTr sets the addtr value on the Table component.
func (t *Table) AddTr(tr *Tr) *Table {
	t.trs = append(t.trs, tr)
	return t
}

// Prepare renders the Table component into its internal buffer.
func (t *Table) Prepare() {
	t.buf.Reset()
	// see if table has id, class, and style tags to add
	t.buf.WriteString("<table")
	if t.id != "" {
		writeAttr(&t.buf, "id", t.id)
	}
	if len(t.class) != 0 {
		writeClassAttr(&t.buf, t.class)
	}
	if len(t.style) != 0 {
		parseStyle(&t.buf, t.style)
	}
	t.buf.WriteByte('>')

	// write caption if present
	if t.caption != nil {
		t.caption.Prepare()
		t.buf.Write(t.caption.Bytes())
	}

	// write thead if present
	if t.thead != nil {
		t.thead.Prepare()
		t.buf.Write(t.thead.Bytes())
	}

	// write tbody if present
	if t.tbody != nil {
		t.tbody.Prepare()
		t.buf.Write(t.tbody.Bytes())
	}

	// write tfoot if present
	if t.tfoot != nil {
		t.tfoot.Prepare()
		t.buf.Write(t.tfoot.Bytes())
	}

	// write direct tr elements if present
	for _, tr := range t.trs {
		tr.Prepare()
		t.buf.Write(tr.Bytes())
	}

	// write legacy header and rows if no structured elements are used
	if t.thead == nil && t.tbody == nil && t.tfoot == nil && len(t.trs) == 0 {
		// write header
		if len(t.headers) > 0 {
			t.buf.WriteString("<tr>")
			for _, h := range t.headers {
				t.buf.WriteString("<th>" + escapeText(h) + "</th>")
			}
			t.buf.WriteString("</tr>")
		}

		// write rows
		for i := 0; i < len(t.rows); i++ {
			t.buf.WriteString("<tr>")
			for j := 0; j < len(t.rows[i]); j++ {
				t.buf.WriteString("<td>" + escapeText(t.rows[i][j]) + "</td>")
			}
			t.buf.WriteString("</tr>")
		}
	}

	t.buf.WriteString("</table>")
}

// Bytes returns a defensive copy of the rendered Table bytes.
func (t *Table) Bytes() []byte {
	return cloneBytes(t.buf.Bytes())
}

// IsBodyElement implements BodyElement interface
func (t *Table) IsBodyElement() {}

// Thead represents the Thead component or supporting type.
type Thead struct {
	buf      bytes.Buffer
	class    []string
	id       string
	style    map[string]string
	contents []Element
}

// NewThead creates a new Thead component.
func NewThead() *Thead {
	return &Thead{
		style: make(map[string]string),
	}
}

// AddClass sets the addclass value on the Thead component.
func (th *Thead) AddClass(s string) *Thead {
	th.class = append(th.class, s)
	return th
}

// AddClasses sets the addclasses value on the Thead component.
func (th *Thead) AddClasses(s []string) *Thead {
	th.class = append(th.class, s...)
	return th
}

// Class sets the class value on the Thead component.
func (th *Thead) Class(s []string) *Thead {
	th.class = cloneStrings(s)
	return th
}

// AddId sets the addid value on the Thead component.
func (th *Thead) AddId(s string) *Thead {
	th.id = s
	return th
}

// Id sets the id value on the Thead component.
func (th *Thead) Id(s string) *Thead {
	th.id = s
	return th
}

// AddStyle adds one inline CSS declaration to the Thead component.
func (th *Thead) AddStyle(k, v string) *Thead {
	th.style[k] = v
	return th
}

// AddStyles adds multiple inline CSS declarations to the Thead component.
func (th *Thead) AddStyles(m map[string]string) *Thead {
	for k, v := range m {
		th.style[k] = v
	}
	return th
}

// Styles replaces the inline CSS declarations on the Thead component.
func (th *Thead) Styles(m map[string]string) *Thead {
	th.style = cloneStyleMap(m)
	return th
}

// AddTr sets the addtr value on the Thead component.
func (th *Thead) AddTr(tr *Tr) *Thead {
	th.contents = appendElement(th.contents, tr)
	return th
}

// Prepare renders the Thead component into its internal buffer.
func (th *Thead) Prepare() {
	th.buf.Reset()
	th.buf.WriteString("<thead")
	if th.id != "" {
		writeAttr(&th.buf, "id", th.id)
	}
	if len(th.class) != 0 {
		writeClassAttr(&th.buf, th.class)
	}
	if len(th.style) != 0 {
		parseStyle(&th.buf, th.style)
	}
	th.buf.WriteByte('>')

	writeElements(&th.buf, th.contents)
	th.buf.WriteString("</thead>")
}

// Bytes returns a defensive copy of the rendered Thead bytes.
func (th *Thead) Bytes() []byte {
	return cloneBytes(th.buf.Bytes())
}

// IsBodyElement implements BodyElement interface
func (th *Thead) IsBodyElement() {}

// Tbody represents the Tbody component or supporting type.
type Tbody struct {
	buf      bytes.Buffer
	class    []string
	id       string
	style    map[string]string
	contents []Element
}

// NewTbody creates a new Tbody component.
func NewTbody() *Tbody {
	return &Tbody{
		style: make(map[string]string),
	}
}

// AddClass sets the addclass value on the Tbody component.
func (tb *Tbody) AddClass(s string) *Tbody {
	tb.class = append(tb.class, s)
	return tb
}

// AddClasses sets the addclasses value on the Tbody component.
func (tb *Tbody) AddClasses(s []string) *Tbody {
	tb.class = append(tb.class, s...)
	return tb
}

// Class sets the class value on the Tbody component.
func (tb *Tbody) Class(s []string) *Tbody {
	tb.class = cloneStrings(s)
	return tb
}

// AddId sets the addid value on the Tbody component.
func (tb *Tbody) AddId(s string) *Tbody {
	tb.id = s
	return tb
}

// Id sets the id value on the Tbody component.
func (tb *Tbody) Id(s string) *Tbody {
	tb.id = s
	return tb
}

// AddStyle adds one inline CSS declaration to the Tbody component.
func (tb *Tbody) AddStyle(k, v string) *Tbody {
	tb.style[k] = v
	return tb
}

// AddStyles adds multiple inline CSS declarations to the Tbody component.
func (tb *Tbody) AddStyles(m map[string]string) *Tbody {
	for k, v := range m {
		tb.style[k] = v
	}
	return tb
}

// Styles replaces the inline CSS declarations on the Tbody component.
func (tb *Tbody) Styles(m map[string]string) *Tbody {
	tb.style = cloneStyleMap(m)
	return tb
}

// AddTr sets the addtr value on the Tbody component.
func (tb *Tbody) AddTr(tr *Tr) *Tbody {
	tb.contents = appendElement(tb.contents, tr)
	return tb
}

// Prepare renders the Tbody component into its internal buffer.
func (tb *Tbody) Prepare() {
	tb.buf.Reset()
	tb.buf.WriteString("<tbody")
	if tb.id != "" {
		writeAttr(&tb.buf, "id", tb.id)
	}
	if len(tb.class) != 0 {
		writeClassAttr(&tb.buf, tb.class)
	}
	if len(tb.style) != 0 {
		parseStyle(&tb.buf, tb.style)
	}
	tb.buf.WriteByte('>')

	writeElements(&tb.buf, tb.contents)
	tb.buf.WriteString("</tbody>")
}

// Bytes returns a defensive copy of the rendered Tbody bytes.
func (tb *Tbody) Bytes() []byte {
	return cloneBytes(tb.buf.Bytes())
}

// IsBodyElement implements BodyElement interface
func (tb *Tbody) IsBodyElement() {}

// Tfoot represents the Tfoot component or supporting type.
type Tfoot struct {
	buf      bytes.Buffer
	class    []string
	id       string
	style    map[string]string
	contents []Element
}

// NewTfoot creates a new Tfoot component.
func NewTfoot() *Tfoot {
	return &Tfoot{
		style: make(map[string]string),
	}
}

// AddClass sets the addclass value on the Tfoot component.
func (tf *Tfoot) AddClass(s string) *Tfoot {
	tf.class = append(tf.class, s)
	return tf
}

// AddClasses sets the addclasses value on the Tfoot component.
func (tf *Tfoot) AddClasses(s []string) *Tfoot {
	tf.class = append(tf.class, s...)
	return tf
}

// Class sets the class value on the Tfoot component.
func (tf *Tfoot) Class(s []string) *Tfoot {
	tf.class = cloneStrings(s)
	return tf
}

// AddId sets the addid value on the Tfoot component.
func (tf *Tfoot) AddId(s string) *Tfoot {
	tf.id = s
	return tf
}

// Id sets the id value on the Tfoot component.
func (tf *Tfoot) Id(s string) *Tfoot {
	tf.id = s
	return tf
}

// AddStyle adds one inline CSS declaration to the Tfoot component.
func (tf *Tfoot) AddStyle(k, v string) *Tfoot {
	tf.style[k] = v
	return tf
}

// AddStyles adds multiple inline CSS declarations to the Tfoot component.
func (tf *Tfoot) AddStyles(m map[string]string) *Tfoot {
	for k, v := range m {
		tf.style[k] = v
	}
	return tf
}

// Styles replaces the inline CSS declarations on the Tfoot component.
func (tf *Tfoot) Styles(m map[string]string) *Tfoot {
	tf.style = cloneStyleMap(m)
	return tf
}

// AddTr sets the addtr value on the Tfoot component.
func (tf *Tfoot) AddTr(tr *Tr) *Tfoot {
	tf.contents = appendElement(tf.contents, tr)
	return tf
}

// Prepare renders the Tfoot component into its internal buffer.
func (tf *Tfoot) Prepare() {
	tf.buf.Reset()
	tf.buf.WriteString("<tfoot")
	if tf.id != "" {
		writeAttr(&tf.buf, "id", tf.id)
	}
	if len(tf.class) != 0 {
		writeClassAttr(&tf.buf, tf.class)
	}
	if len(tf.style) != 0 {
		parseStyle(&tf.buf, tf.style)
	}
	tf.buf.WriteByte('>')

	writeElements(&tf.buf, tf.contents)
	tf.buf.WriteString("</tfoot>")
}

// Bytes returns a defensive copy of the rendered Tfoot bytes.
func (tf *Tfoot) Bytes() []byte {
	return cloneBytes(tf.buf.Bytes())
}

// IsBodyElement implements BodyElement interface
func (tf *Tfoot) IsBodyElement() {}

// Caption represents the Caption component or supporting type.
type Caption struct {
	buf   bytes.Buffer
	class []string
	id    string
	style map[string]string
	text  string
}

// NewCaption creates a new Caption component.
func NewCaption() *Caption {
	return &Caption{
		style: make(map[string]string),
	}
}

// AddClass sets the addclass value on the Caption component.
func (c *Caption) AddClass(s string) *Caption {
	c.class = append(c.class, s)
	return c
}

// AddClasses sets the addclasses value on the Caption component.
func (c *Caption) AddClasses(s []string) *Caption {
	c.class = append(c.class, s...)
	return c
}

// Class sets the class value on the Caption component.
func (c *Caption) Class(s []string) *Caption {
	c.class = cloneStrings(s)
	return c
}

// AddId sets the addid value on the Caption component.
func (c *Caption) AddId(s string) *Caption {
	c.id = s
	return c
}

// Id sets the id value on the Caption component.
func (c *Caption) Id(s string) *Caption {
	c.id = s
	return c
}

// AddStyle adds one inline CSS declaration to the Caption component.
func (c *Caption) AddStyle(k, v string) *Caption {
	c.style[k] = v
	return c
}

// AddStyles adds multiple inline CSS declarations to the Caption component.
func (c *Caption) AddStyles(m map[string]string) *Caption {
	for k, v := range m {
		c.style[k] = v
	}
	return c
}

// Styles replaces the inline CSS declarations on the Caption component.
func (c *Caption) Styles(m map[string]string) *Caption {
	c.style = cloneStyleMap(m)
	return c
}

// Text sets or appends text content on the Caption component.
func (c *Caption) Text(text string) *Caption {
	c.text = text
	return c
}

// Prepare renders the Caption component into its internal buffer.
func (c *Caption) Prepare() {
	c.buf.Reset()
	c.buf.WriteString("<caption")
	if c.id != "" {
		writeAttr(&c.buf, "id", c.id)
	}
	if len(c.class) != 0 {
		writeClassAttr(&c.buf, c.class)
	}
	if len(c.style) != 0 {
		parseStyle(&c.buf, c.style)
	}
	c.buf.WriteByte('>')

	c.buf.WriteString(escapeText(c.text))
	c.buf.WriteString("</caption>")
}

// Bytes returns a defensive copy of the rendered Caption bytes.
func (c *Caption) Bytes() []byte {
	return cloneBytes(c.buf.Bytes())
}

// IsBodyElement implements BodyElement interface
func (c *Caption) IsBodyElement() {}

// Col represents the Col component or supporting type.
type Col struct {
	buf   bytes.Buffer
	class []string
	id    string
	style map[string]string
	span  int
}

// NewCol creates a new Col component.
func NewCol() *Col {
	return &Col{
		style: make(map[string]string),
	}
}

// AddClass sets the addclass value on the Col component.
func (col *Col) AddClass(s string) *Col {
	col.class = append(col.class, s)
	return col
}

// AddClasses sets the addclasses value on the Col component.
func (col *Col) AddClasses(s []string) *Col {
	col.class = append(col.class, s...)
	return col
}

// Class sets the class value on the Col component.
func (col *Col) Class(s []string) *Col {
	col.class = cloneStrings(s)
	return col
}

// AddId sets the addid value on the Col component.
func (col *Col) AddId(s string) *Col {
	col.id = s
	return col
}

// Id sets the id value on the Col component.
func (col *Col) Id(s string) *Col {
	col.id = s
	return col
}

// AddStyle adds one inline CSS declaration to the Col component.
func (col *Col) AddStyle(k, v string) *Col {
	col.style[k] = v
	return col
}

// AddStyles adds multiple inline CSS declarations to the Col component.
func (col *Col) AddStyles(m map[string]string) *Col {
	for k, v := range m {
		col.style[k] = v
	}
	return col
}

// Styles replaces the inline CSS declarations on the Col component.
func (col *Col) Styles(m map[string]string) *Col {
	col.style = cloneStyleMap(m)
	return col
}

// Span sets the span value on the Col component.
func (col *Col) Span(s int) *Col {
	col.span = s
	return col
}

// Prepare renders the Col component into its internal buffer.
func (col *Col) Prepare() {
	col.buf.Reset()
	col.buf.WriteString("<col")
	if col.id != "" {
		writeAttr(&col.buf, "id", col.id)
	}
	if len(col.class) != 0 {
		writeClassAttr(&col.buf, col.class)
	}
	if len(col.style) != 0 {
		parseStyle(&col.buf, col.style)
	}
	if col.span > 0 {
		writeIntAttr(&col.buf, "span", col.span)
	}
	col.buf.WriteString(">")
}

// Bytes returns a defensive copy of the rendered Col bytes.
func (col *Col) Bytes() []byte {
	return cloneBytes(col.buf.Bytes())
}

// IsBodyElement implements BodyElement interface
func (col *Col) IsBodyElement() {}

// Colgroup represents the Colgroup component or supporting type.
type Colgroup struct {
	buf      bytes.Buffer
	class    []string
	id       string
	style    map[string]string
	contents []Element
	span     int
}

// NewColgroup creates a new Colgroup component.
func NewColgroup() *Colgroup {
	return &Colgroup{
		style: make(map[string]string),
	}
}

// AddClass sets the addclass value on the Colgroup component.
func (cg *Colgroup) AddClass(s string) *Colgroup {
	cg.class = append(cg.class, s)
	return cg
}

// AddClasses sets the addclasses value on the Colgroup component.
func (cg *Colgroup) AddClasses(s []string) *Colgroup {
	cg.class = append(cg.class, s...)
	return cg
}

// Class sets the class value on the Colgroup component.
func (cg *Colgroup) Class(s []string) *Colgroup {
	cg.class = cloneStrings(s)
	return cg
}

// AddId sets the addid value on the Colgroup component.
func (cg *Colgroup) AddId(s string) *Colgroup {
	cg.id = s
	return cg
}

// Id sets the id value on the Colgroup component.
func (cg *Colgroup) Id(s string) *Colgroup {
	cg.id = s
	return cg
}

// AddStyle adds one inline CSS declaration to the Colgroup component.
func (cg *Colgroup) AddStyle(k, v string) *Colgroup {
	cg.style[k] = v
	return cg
}

// AddStyles adds multiple inline CSS declarations to the Colgroup component.
func (cg *Colgroup) AddStyles(m map[string]string) *Colgroup {
	for k, v := range m {
		cg.style[k] = v
	}
	return cg
}

// Styles replaces the inline CSS declarations on the Colgroup component.
func (cg *Colgroup) Styles(m map[string]string) *Colgroup {
	cg.style = cloneStyleMap(m)
	return cg
}

// Add appends child content to the Colgroup component.
func (cg *Colgroup) Add(e Element) *Colgroup {
	cg.contents = appendElement(cg.contents, e)
	return cg
}

// Span sets the span value on the Colgroup component.
func (cg *Colgroup) Span(s int) *Colgroup {
	cg.span = s
	return cg
}

// Prepare renders the Colgroup component into its internal buffer.
func (cg *Colgroup) Prepare() {
	cg.buf.Reset()
	cg.buf.WriteString("<colgroup")
	if cg.id != "" {
		writeAttr(&cg.buf, "id", cg.id)
	}
	if len(cg.class) != 0 {
		writeClassAttr(&cg.buf, cg.class)
	}
	if len(cg.style) != 0 {
		parseStyle(&cg.buf, cg.style)
	}
	if cg.span > 0 {
		writeIntAttr(&cg.buf, "span", cg.span)
	}
	cg.buf.WriteByte('>')

	writeElements(&cg.buf, cg.contents)
	cg.buf.WriteString("</colgroup>")
}

// Bytes returns a defensive copy of the rendered Colgroup bytes.
func (cg *Colgroup) Bytes() []byte {
	return cloneBytes(cg.buf.Bytes())
}

// IsBodyElement implements BodyElement interface
func (cg *Colgroup) IsBodyElement() {}

// Tr represents the Tr component or supporting type.
type Tr struct {
	buf      bytes.Buffer
	class    []string
	id       string
	style    map[string]string
	contents []Element
}

// NewTr creates a new Tr component.
func NewTr() *Tr {
	return &Tr{
		style: make(map[string]string),
	}
}

// AddClass sets the addclass value on the Tr component.
func (tr *Tr) AddClass(s string) *Tr {
	tr.class = append(tr.class, s)
	return tr
}

// AddClasses sets the addclasses value on the Tr component.
func (tr *Tr) AddClasses(s []string) *Tr {
	tr.class = append(tr.class, s...)
	return tr
}

// Class sets the class value on the Tr component.
func (tr *Tr) Class(s []string) *Tr {
	tr.class = cloneStrings(s)
	return tr
}

// AddId sets the addid value on the Tr component.
func (tr *Tr) AddId(s string) *Tr {
	tr.id = s
	return tr
}

// Id sets the id value on the Tr component.
func (tr *Tr) Id(s string) *Tr {
	tr.id = s
	return tr
}

// AddStyle adds one inline CSS declaration to the Tr component.
func (tr *Tr) AddStyle(k, v string) *Tr {
	tr.style[k] = v
	return tr
}

// AddStyles adds multiple inline CSS declarations to the Tr component.
func (tr *Tr) AddStyles(m map[string]string) *Tr {
	for k, v := range m {
		tr.style[k] = v
	}
	return tr
}

// Styles replaces the inline CSS declarations on the Tr component.
func (tr *Tr) Styles(m map[string]string) *Tr {
	tr.style = cloneStyleMap(m)
	return tr
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

// Prepare renders the Tr component into its internal buffer.
func (tr *Tr) Prepare() {
	tr.buf.Reset()
	tr.buf.WriteString("<tr")
	if tr.id != "" {
		writeAttr(&tr.buf, "id", tr.id)
	}
	if len(tr.class) != 0 {
		writeClassAttr(&tr.buf, tr.class)
	}
	if len(tr.style) != 0 {
		parseStyle(&tr.buf, tr.style)
	}
	tr.buf.WriteByte('>')

	writeElements(&tr.buf, tr.contents)
	tr.buf.WriteString("</tr>")
}

// Bytes returns a defensive copy of the rendered Tr bytes.
func (tr *Tr) Bytes() []byte {
	return cloneBytes(tr.buf.Bytes())
}

// IsBodyElement implements BodyElement interface
func (tr *Tr) IsBodyElement() {}

// Td represents the Td component or supporting type.
type Td struct {
	buf      bytes.Buffer
	class    []string
	id       string
	style    map[string]string
	contents []Element
	colspan  int
	rowspan  int
}

// NewTd creates a new Td component.
func NewTd() *Td {
	return &Td{
		style: make(map[string]string),
	}
}

// AddClass sets the addclass value on the Td component.
func (td *Td) AddClass(s string) *Td {
	td.class = append(td.class, s)
	return td
}

// AddClasses sets the addclasses value on the Td component.
func (td *Td) AddClasses(s []string) *Td {
	td.class = append(td.class, s...)
	return td
}

// Class sets the class value on the Td component.
func (td *Td) Class(s []string) *Td {
	td.class = cloneStrings(s)
	return td
}

// AddId sets the addid value on the Td component.
func (td *Td) AddId(s string) *Td {
	td.id = s
	return td
}

// Id sets the id value on the Td component.
func (td *Td) Id(s string) *Td {
	td.id = s
	return td
}

// AddStyle adds one inline CSS declaration to the Td component.
func (td *Td) AddStyle(k, v string) *Td {
	td.style[k] = v
	return td
}

// AddStyles adds multiple inline CSS declarations to the Td component.
func (td *Td) AddStyles(m map[string]string) *Td {
	for k, v := range m {
		td.style[k] = v
	}
	return td
}

// Styles replaces the inline CSS declarations on the Td component.
func (td *Td) Styles(m map[string]string) *Td {
	td.style = cloneStyleMap(m)
	return td
}

// Add appends child content to the Td component.
func (td *Td) Add(e Element) *Td {
	td.contents = appendElement(td.contents, e)
	return td
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

// Prepare renders the Td component into its internal buffer.
func (td *Td) Prepare() {
	td.buf.Reset()
	td.buf.WriteString("<td")
	if td.id != "" {
		writeAttr(&td.buf, "id", td.id)
	}
	if len(td.class) != 0 {
		writeClassAttr(&td.buf, td.class)
	}
	if len(td.style) != 0 {
		parseStyle(&td.buf, td.style)
	}
	if td.colspan > 0 {
		writeIntAttr(&td.buf, "colspan", td.colspan)
	}
	if td.rowspan > 0 {
		writeIntAttr(&td.buf, "rowspan", td.rowspan)
	}
	td.buf.WriteByte('>')

	writeElements(&td.buf, td.contents)
	td.buf.WriteString("</td>")
}

// Bytes returns a defensive copy of the rendered Td bytes.
func (td *Td) Bytes() []byte {
	return cloneBytes(td.buf.Bytes())
}

// IsBodyElement implements BodyElement interface
func (td *Td) IsBodyElement() {}

// Th represents the Th component or supporting type.
type Th struct {
	buf      bytes.Buffer
	class    []string
	id       string
	style    map[string]string
	contents []Element
	colspan  int
	rowspan  int
	scope    string
}

// NewTh creates a new Th component.
func NewTh() *Th {
	return &Th{
		style: make(map[string]string),
	}
}

// AddClass sets the addclass value on the Th component.
func (th *Th) AddClass(s string) *Th {
	th.class = append(th.class, s)
	return th
}

// AddClasses sets the addclasses value on the Th component.
func (th *Th) AddClasses(s []string) *Th {
	th.class = append(th.class, s...)
	return th
}

// Class sets the class value on the Th component.
func (th *Th) Class(s []string) *Th {
	th.class = cloneStrings(s)
	return th
}

// AddId sets the addid value on the Th component.
func (th *Th) AddId(s string) *Th {
	th.id = s
	return th
}

// Id sets the id value on the Th component.
func (th *Th) Id(s string) *Th {
	th.id = s
	return th
}

// AddStyle adds one inline CSS declaration to the Th component.
func (th *Th) AddStyle(k, v string) *Th {
	th.style[k] = v
	return th
}

// AddStyles adds multiple inline CSS declarations to the Th component.
func (th *Th) AddStyles(m map[string]string) *Th {
	for k, v := range m {
		th.style[k] = v
	}
	return th
}

// Styles replaces the inline CSS declarations on the Th component.
func (th *Th) Styles(m map[string]string) *Th {
	th.style = cloneStyleMap(m)
	return th
}

// Add appends child content to the Th component.
func (th *Th) Add(e Element) *Th {
	th.contents = appendElement(th.contents, e)
	return th
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

// Prepare renders the Th component into its internal buffer.
func (th *Th) Prepare() {
	th.buf.Reset()
	th.buf.WriteString("<th")
	if th.id != "" {
		writeAttr(&th.buf, "id", th.id)
	}
	if len(th.class) != 0 {
		writeClassAttr(&th.buf, th.class)
	}
	if len(th.style) != 0 {
		parseStyle(&th.buf, th.style)
	}
	if th.colspan > 0 {
		writeIntAttr(&th.buf, "colspan", th.colspan)
	}
	if th.rowspan > 0 {
		writeIntAttr(&th.buf, "rowspan", th.rowspan)
	}
	if th.scope != "" {
		writeAttr(&th.buf, "scope", th.scope)
	}
	th.buf.WriteByte('>')

	writeElements(&th.buf, th.contents)
	th.buf.WriteString("</th>")
}

// Bytes returns a defensive copy of the rendered Th bytes.
func (th *Th) Bytes() []byte {
	return cloneBytes(th.buf.Bytes())
}

// IsBodyElement implements BodyElement interface
func (th *Th) IsBodyElement() {}
