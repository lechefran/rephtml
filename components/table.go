package rephtml

import (
	"bytes"
	"fmt"
)

type Table struct {
	buf     bytes.Buffer
	class   []string
	headers []string
	id      string
	rows    [][]string
	style   []string
}

func NewTable() *Table {
	return &Table{}
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
	t.headers = append(t.headers, s...)
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
	t.rows = append(t.rows, s...)
	return t
}

func (t *Table) AddStyle(s string) *Table {
	t.style = append(t.style, s)
	return t
}

func (t *Table) AddStyles(s []string) *Table {
	t.style = append(t.style, s...)
	return t
}

func (t *Table) Styles(s []string) *Table {
	t.style = append(t.style, s...)
	return t
}

func (t *Table) Prepare() *Table {
	// see if table has id and class tags to add
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
	return t
}

func (t *Table) Bytes() []byte {
	fmt.Println(t.buf.String())
	return t.buf.Bytes()
}
