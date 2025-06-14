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

func (t *Table) Prepare() {
	// see if table has id, class, and style tags to add
	t.buf.WriteString("<table")
	if t.id != "" {
		t.buf.WriteString(" id=\"" + t.id + "\"")
	}
	if len(t.class) != 0 {
		t.buf.WriteString(" class=\"")
		for i := 0; i < len(t.class); i++ {
			t.buf.WriteString(t.class[i])
			if i != len(t.class)-1 {
				t.buf.WriteString(" ")
			}
		}
		t.buf.WriteString("\"")
	}
	if len(t.style) != 0 {
		idx := 0
		t.buf.WriteString(" style=\"")
		for k, v := range t.style {
			t.buf.WriteString(k + ": " + v + ";")
			if idx != len(t.style)-1 {
				t.buf.WriteString(" ")
			}
			idx++
		}
		t.buf.WriteString("\"")
	}
	t.buf.WriteByte('>')

	// write header
	t.buf.WriteString("<tr>")
	for _, h := range t.headers {
		t.buf.WriteString("<th>" + h + "</th>")
	}
	t.buf.WriteString("</tr>")

	// write rows
	for i := 0; i < len(t.rows); i++ {
		t.buf.WriteString("<tr>")
		for j := 0; j < len(t.rows[i]); j++ {
			t.buf.WriteString("<td>" + t.rows[i][j] + "</td>")
		}
		t.buf.WriteString("</tr>")
	}
	t.buf.WriteString("</table>")
}

func (t *Table) Bytes() []byte {
	return t.buf.Bytes()
}

type Thead struct {
	buf      bytes.Buffer
	class    []string
	id       string
	style    map[string]string
	contents [][]byte
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

func (th *Thead) Add(e Elements) *Thead {
	th.contents = append(th.contents, e.Bytes())
	return th
}

func (th *Thead) Prepare() {
	th.buf.WriteString("<thead")
	if th.id != "" {
		th.buf.WriteString(" id=\"" + th.id + "\"")
	}
	if len(th.class) != 0 {
		th.buf.WriteString(" class=\"")
		for i := 0; i < len(th.class); i++ {
			th.buf.WriteString(th.class[i])
			if i != len(th.class)-1 {
				th.buf.WriteString(" ")
			}
		}
		th.buf.WriteString("\"")
	}
	if len(th.style) != 0 {
		idx := 0
		th.buf.WriteString(" style=\"")
		for k, v := range th.style {
			th.buf.WriteString(k + ": " + v + ";")
			if idx != len(th.style)-1 {
				th.buf.WriteString(" ")
			}
			idx++
		}
		th.buf.WriteString("\"")
	}
	th.buf.WriteByte('>')

	for _, c := range th.contents {
		th.buf.Write(c)
	}
	th.buf.WriteString("</thead>")
}

func (th *Thead) Bytes() []byte {
	return th.buf.Bytes()
}

type Tbody struct {
	buf      bytes.Buffer
	class    []string
	id       string
	style    map[string]string
	contents [][]byte
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

func (tb *Tbody) Add(e Elements) *Tbody {
	tb.contents = append(tb.contents, e.Bytes())
	return tb
}

func (tb *Tbody) Prepare() {
	tb.buf.WriteString("<tbody")
	if tb.id != "" {
		tb.buf.WriteString(" id=\"" + tb.id + "\"")
	}
	if len(tb.class) != 0 {
		tb.buf.WriteString(" class=\"")
		for i := 0; i < len(tb.class); i++ {
			tb.buf.WriteString(tb.class[i])
			if i != len(tb.class)-1 {
				tb.buf.WriteString(" ")
			}
		}
		tb.buf.WriteString("\"")
	}
	if len(tb.style) != 0 {
		idx := 0
		tb.buf.WriteString(" style=\"")
		for k, v := range tb.style {
			tb.buf.WriteString(k + ": " + v + ";")
			if idx != len(tb.style)-1 {
				tb.buf.WriteString(" ")
			}
			idx++
		}
		tb.buf.WriteString("\"")
	}
	tb.buf.WriteByte('>')

	for _, c := range tb.contents {
		tb.buf.Write(c)
	}
	tb.buf.WriteString("</tbody>")
}

func (tb *Tbody) Bytes() []byte {
	return tb.buf.Bytes()
}

type Tfoot struct {
	buf      bytes.Buffer
	class    []string
	id       string
	style    map[string]string
	contents [][]byte
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

func (tf *Tfoot) Add(e Elements) *Tfoot {
	tf.contents = append(tf.contents, e.Bytes())
	return tf
}

func (tf *Tfoot) Prepare() {
	tf.buf.WriteString("<tfoot")
	if tf.id != "" {
		tf.buf.WriteString(" id=\"" + tf.id + "\"")
	}
	if len(tf.class) != 0 {
		tf.buf.WriteString(" class=\"")
		for i := 0; i < len(tf.class); i++ {
			tf.buf.WriteString(tf.class[i])
			if i != len(tf.class)-1 {
				tf.buf.WriteString(" ")
			}
		}
		tf.buf.WriteString("\"")
	}
	if len(tf.style) != 0 {
		idx := 0
		tf.buf.WriteString(" style=\"")
		for k, v := range tf.style {
			tf.buf.WriteString(k + ": " + v + ";")
			if idx != len(tf.style)-1 {
				tf.buf.WriteString(" ")
			}
			idx++
		}
		tf.buf.WriteString("\"")
	}
	tf.buf.WriteByte('>')

	for _, c := range tf.contents {
		tf.buf.Write(c)
	}
	tf.buf.WriteString("</tfoot>")
}

func (tf *Tfoot) Bytes() []byte {
	return tf.buf.Bytes()
}
