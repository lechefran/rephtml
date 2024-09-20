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
