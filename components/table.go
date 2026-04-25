package rephtml

import (
	"bytes"
)

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

func NewTable() *Table {
	return &Table{
		style: make(map[string]string),
	}
}

func (t *Table) AddClass(s string) *Table {
	t.class = append(t.class, s)
	return t
}

func (t *Table) AddClasses(s []string) *Table {
	t.class = append(t.class, s...)
	return t
}

func (t *Table) Class(s []string) *Table {
	t.class = append(t.class, s...)
	return t
}

func (t *Table) AddHeader(s string) *Table {
	t.headers = append(t.headers, s)
	return t
}

func (t *Table) AddHeaders(s []string) *Table {
	t.headers = append(t.headers, s...)
	return t
}

func (t *Table) Headers(s []string) *Table {
	t.headers = s
	return t
}

func (t *Table) AddId(s string) *Table {
	t.id = s
	return t
}

func (t *Table) Id(s string) *Table {
	t.id = s
	return t
}

func (t *Table) AddRow(s []string) *Table {
	t.rows = append(t.rows, s)
	return t
}

func (t *Table) AddRows(s [][]string) *Table {
	t.rows = append(t.rows, s...)
	return t
}

func (t *Table) Rows(s [][]string) *Table {
	t.rows = s
	return t
}

func (t *Table) AddStyle(k, v string) *Table {
	t.style[k] = v
	return t
}

func (t *Table) AddStyles(m map[string]string) *Table {
	for k, v := range m {
		t.style[k] = v
	}
	return t
}

func (t *Table) Styles(m map[string]string) *Table {
	t.style = m
	return t
}

func (t *Table) AddCaption(c *Caption) *Table {
	t.caption = c
	return t
}

func (t *Table) AddThead(th *Thead) *Table {
	t.thead = th
	return t
}

func (t *Table) AddTbody(tb *Tbody) *Table {
	t.tbody = tb
	return t
}

func (t *Table) AddTfoot(tf *Tfoot) *Table {
	t.tfoot = tf
	return t
}

func (t *Table) AddTr(tr *Tr) *Table {
	t.trs = append(t.trs, tr)
	return t
}

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

func (t *Table) Bytes() []byte {
	return cloneBytes(t.buf.Bytes())
}

// IsBodyElement implements BodyElement interface
func (t *Table) IsBodyElement() {}

type Thead struct {
	buf      bytes.Buffer
	class    []string
	id       string
	style    map[string]string
	contents []Element
}

func NewThead() *Thead {
	return &Thead{
		style: make(map[string]string),
	}
}

func (th *Thead) AddClass(s string) *Thead {
	th.class = append(th.class, s)
	return th
}

func (th *Thead) AddClasses(s []string) *Thead {
	th.class = append(th.class, s...)
	return th
}

func (th *Thead) Class(s []string) *Thead {
	th.class = append(th.class, s...)
	return th
}

func (th *Thead) AddId(s string) *Thead {
	th.id = s
	return th
}

func (th *Thead) Id(s string) *Thead {
	th.id = s
	return th
}

func (th *Thead) AddStyle(k, v string) *Thead {
	th.style[k] = v
	return th
}

func (th *Thead) AddStyles(m map[string]string) *Thead {
	for k, v := range m {
		th.style[k] = v
	}
	return th
}

func (th *Thead) Styles(m map[string]string) *Thead {
	th.style = m
	return th
}

func (th *Thead) AddTr(tr *Tr) *Thead {
	th.contents = appendElement(th.contents, tr)
	return th
}

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

func (th *Thead) Bytes() []byte {
	return cloneBytes(th.buf.Bytes())
}

// IsBodyElement implements BodyElement interface
func (th *Thead) IsBodyElement() {}

type Tbody struct {
	buf      bytes.Buffer
	class    []string
	id       string
	style    map[string]string
	contents []Element
}

func NewTbody() *Tbody {
	return &Tbody{
		style: make(map[string]string),
	}
}

func (tb *Tbody) AddClass(s string) *Tbody {
	tb.class = append(tb.class, s)
	return tb
}

func (tb *Tbody) AddClasses(s []string) *Tbody {
	tb.class = append(tb.class, s...)
	return tb
}

func (tb *Tbody) Class(s []string) *Tbody {
	tb.class = append(tb.class, s...)
	return tb
}

func (tb *Tbody) AddId(s string) *Tbody {
	tb.id = s
	return tb
}

func (tb *Tbody) Id(s string) *Tbody {
	tb.id = s
	return tb
}

func (tb *Tbody) AddStyle(k, v string) *Tbody {
	tb.style[k] = v
	return tb
}

func (tb *Tbody) AddStyles(m map[string]string) *Tbody {
	for k, v := range m {
		tb.style[k] = v
	}
	return tb
}

func (tb *Tbody) Styles(m map[string]string) *Tbody {
	tb.style = m
	return tb
}

func (tb *Tbody) AddTr(tr *Tr) *Tbody {
	tb.contents = appendElement(tb.contents, tr)
	return tb
}

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

func (tb *Tbody) Bytes() []byte {
	return cloneBytes(tb.buf.Bytes())
}

// IsBodyElement implements BodyElement interface
func (tb *Tbody) IsBodyElement() {}

type Tfoot struct {
	buf      bytes.Buffer
	class    []string
	id       string
	style    map[string]string
	contents []Element
}

func NewTfoot() *Tfoot {
	return &Tfoot{
		style: make(map[string]string),
	}
}

func (tf *Tfoot) AddClass(s string) *Tfoot {
	tf.class = append(tf.class, s)
	return tf
}

func (tf *Tfoot) AddClasses(s []string) *Tfoot {
	tf.class = append(tf.class, s...)
	return tf
}

func (tf *Tfoot) Class(s []string) *Tfoot {
	tf.class = append(tf.class, s...)
	return tf
}

func (tf *Tfoot) AddId(s string) *Tfoot {
	tf.id = s
	return tf
}

func (tf *Tfoot) Id(s string) *Tfoot {
	tf.id = s
	return tf
}

func (tf *Tfoot) AddStyle(k, v string) *Tfoot {
	tf.style[k] = v
	return tf
}

func (tf *Tfoot) AddStyles(m map[string]string) *Tfoot {
	for k, v := range m {
		tf.style[k] = v
	}
	return tf
}

func (tf *Tfoot) Styles(m map[string]string) *Tfoot {
	tf.style = m
	return tf
}

func (tf *Tfoot) AddTr(tr *Tr) *Tfoot {
	tf.contents = appendElement(tf.contents, tr)
	return tf
}

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

func (tf *Tfoot) Bytes() []byte {
	return cloneBytes(tf.buf.Bytes())
}

// IsBodyElement implements BodyElement interface
func (tf *Tfoot) IsBodyElement() {}

type Caption struct {
	buf   bytes.Buffer
	class []string
	id    string
	style map[string]string
	text  string
}

func NewCaption() *Caption {
	return &Caption{
		style: make(map[string]string),
	}
}

func (c *Caption) AddClass(s string) *Caption {
	c.class = append(c.class, s)
	return c
}

func (c *Caption) AddClasses(s []string) *Caption {
	c.class = append(c.class, s...)
	return c
}

func (c *Caption) Class(s []string) *Caption {
	c.class = append(c.class, s...)
	return c
}

func (c *Caption) AddId(s string) *Caption {
	c.id = s
	return c
}

func (c *Caption) Id(s string) *Caption {
	c.id = s
	return c
}

func (c *Caption) AddStyle(k, v string) *Caption {
	c.style[k] = v
	return c
}

func (c *Caption) AddStyles(m map[string]string) *Caption {
	for k, v := range m {
		c.style[k] = v
	}
	return c
}

func (c *Caption) Styles(m map[string]string) *Caption {
	c.style = m
	return c
}

func (c *Caption) Text(text string) *Caption {
	c.text = text
	return c
}

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

func (c *Caption) Bytes() []byte {
	return cloneBytes(c.buf.Bytes())
}

// IsBodyElement implements BodyElement interface
func (c *Caption) IsBodyElement() {}

type Col struct {
	buf   bytes.Buffer
	class []string
	id    string
	style map[string]string
	span  int
}

func NewCol() *Col {
	return &Col{
		style: make(map[string]string),
	}
}

func (col *Col) AddClass(s string) *Col {
	col.class = append(col.class, s)
	return col
}

func (col *Col) AddClasses(s []string) *Col {
	col.class = append(col.class, s...)
	return col
}

func (col *Col) Class(s []string) *Col {
	col.class = append(col.class, s...)
	return col
}

func (col *Col) AddId(s string) *Col {
	col.id = s
	return col
}

func (col *Col) Id(s string) *Col {
	col.id = s
	return col
}

func (col *Col) AddStyle(k, v string) *Col {
	col.style[k] = v
	return col
}

func (col *Col) AddStyles(m map[string]string) *Col {
	for k, v := range m {
		col.style[k] = v
	}
	return col
}

func (col *Col) Styles(m map[string]string) *Col {
	col.style = m
	return col
}

func (col *Col) Span(s int) *Col {
	col.span = s
	return col
}

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

func (col *Col) Bytes() []byte {
	return cloneBytes(col.buf.Bytes())
}

// IsBodyElement implements BodyElement interface
func (col *Col) IsBodyElement() {}

type Colgroup struct {
	buf      bytes.Buffer
	class    []string
	id       string
	style    map[string]string
	contents []Element
	span     int
}

func NewColgroup() *Colgroup {
	return &Colgroup{
		style: make(map[string]string),
	}
}

func (cg *Colgroup) AddClass(s string) *Colgroup {
	cg.class = append(cg.class, s)
	return cg
}

func (cg *Colgroup) AddClasses(s []string) *Colgroup {
	cg.class = append(cg.class, s...)
	return cg
}

func (cg *Colgroup) Class(s []string) *Colgroup {
	cg.class = append(cg.class, s...)
	return cg
}

func (cg *Colgroup) AddId(s string) *Colgroup {
	cg.id = s
	return cg
}

func (cg *Colgroup) Id(s string) *Colgroup {
	cg.id = s
	return cg
}

func (cg *Colgroup) AddStyle(k, v string) *Colgroup {
	cg.style[k] = v
	return cg
}

func (cg *Colgroup) AddStyles(m map[string]string) *Colgroup {
	for k, v := range m {
		cg.style[k] = v
	}
	return cg
}

func (cg *Colgroup) Styles(m map[string]string) *Colgroup {
	cg.style = m
	return cg
}

func (cg *Colgroup) Add(e Element) *Colgroup {
	cg.contents = appendElement(cg.contents, e)
	return cg
}

func (cg *Colgroup) Span(s int) *Colgroup {
	cg.span = s
	return cg
}

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

func (cg *Colgroup) Bytes() []byte {
	return cloneBytes(cg.buf.Bytes())
}

// IsBodyElement implements BodyElement interface
func (cg *Colgroup) IsBodyElement() {}

type Tr struct {
	buf      bytes.Buffer
	class    []string
	id       string
	style    map[string]string
	contents []Element
}

func NewTr() *Tr {
	return &Tr{
		style: make(map[string]string),
	}
}

func (tr *Tr) AddClass(s string) *Tr {
	tr.class = append(tr.class, s)
	return tr
}

func (tr *Tr) AddClasses(s []string) *Tr {
	tr.class = append(tr.class, s...)
	return tr
}

func (tr *Tr) Class(s []string) *Tr {
	tr.class = append(tr.class, s...)
	return tr
}

func (tr *Tr) AddId(s string) *Tr {
	tr.id = s
	return tr
}

func (tr *Tr) Id(s string) *Tr {
	tr.id = s
	return tr
}

func (tr *Tr) AddStyle(k, v string) *Tr {
	tr.style[k] = v
	return tr
}

func (tr *Tr) AddStyles(m map[string]string) *Tr {
	for k, v := range m {
		tr.style[k] = v
	}
	return tr
}

func (tr *Tr) Styles(m map[string]string) *Tr {
	tr.style = m
	return tr
}

func (tr *Tr) AddTh(th *Th) *Tr {
	tr.contents = appendElement(tr.contents, th)
	return tr
}

func (tr *Tr) AddTd(td *Td) *Tr {
	tr.contents = appendElement(tr.contents, td)
	return tr
}

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

func (tr *Tr) Bytes() []byte {
	return cloneBytes(tr.buf.Bytes())
}

// IsBodyElement implements BodyElement interface
func (tr *Tr) IsBodyElement() {}

type Td struct {
	buf      bytes.Buffer
	class    []string
	id       string
	style    map[string]string
	contents []Element
	colspan  int
	rowspan  int
}

func NewTd() *Td {
	return &Td{
		style: make(map[string]string),
	}
}

func (td *Td) AddClass(s string) *Td {
	td.class = append(td.class, s)
	return td
}

func (td *Td) AddClasses(s []string) *Td {
	td.class = append(td.class, s...)
	return td
}

func (td *Td) Class(s []string) *Td {
	td.class = append(td.class, s...)
	return td
}

func (td *Td) AddId(s string) *Td {
	td.id = s
	return td
}

func (td *Td) Id(s string) *Td {
	td.id = s
	return td
}

func (td *Td) AddStyle(k, v string) *Td {
	td.style[k] = v
	return td
}

func (td *Td) AddStyles(m map[string]string) *Td {
	for k, v := range m {
		td.style[k] = v
	}
	return td
}

func (td *Td) Styles(m map[string]string) *Td {
	td.style = m
	return td
}

func (td *Td) Add(e Element) *Td {
	td.contents = appendElement(td.contents, e)
	return td
}

func (td *Td) Colspan(c int) *Td {
	td.colspan = c
	return td
}

func (td *Td) Rowspan(r int) *Td {
	td.rowspan = r
	return td
}

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

func (td *Td) Bytes() []byte {
	return cloneBytes(td.buf.Bytes())
}

// IsBodyElement implements BodyElement interface
func (td *Td) IsBodyElement() {}

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

func NewTh() *Th {
	return &Th{
		style: make(map[string]string),
	}
}

func (th *Th) AddClass(s string) *Th {
	th.class = append(th.class, s)
	return th
}

func (th *Th) AddClasses(s []string) *Th {
	th.class = append(th.class, s...)
	return th
}

func (th *Th) Class(s []string) *Th {
	th.class = append(th.class, s...)
	return th
}

func (th *Th) AddId(s string) *Th {
	th.id = s
	return th
}

func (th *Th) Id(s string) *Th {
	th.id = s
	return th
}

func (th *Th) AddStyle(k, v string) *Th {
	th.style[k] = v
	return th
}

func (th *Th) AddStyles(m map[string]string) *Th {
	for k, v := range m {
		th.style[k] = v
	}
	return th
}

func (th *Th) Styles(m map[string]string) *Th {
	th.style = m
	return th
}

func (th *Th) Add(e Element) *Th {
	th.contents = appendElement(th.contents, e)
	return th
}

func (th *Th) Colspan(c int) *Th {
	th.colspan = c
	return th
}

func (th *Th) Rowspan(r int) *Th {
	th.rowspan = r
	return th
}

func (th *Th) Scope(s string) *Th {
	th.scope = s
	return th
}

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

func (th *Th) Bytes() []byte {
	return cloneBytes(th.buf.Bytes())
}

// IsBodyElement implements BodyElement interface
func (th *Th) IsBodyElement() {}
