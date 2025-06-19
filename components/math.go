package rephtml

import "bytes"

type Math struct {
	buf      bytes.Buffer
	style    map[string]string
	contents [][]byte
	display  string
	xmlns    string
}

func NewMath() *Math {
	return &Math{
		style: make(map[string]string),
		xmlns: "http://www.w3.org/1998/Math/MathML",
	}
}

func (m *Math) AddStyle(k, v string) *Math {
	m.style[k] = v
	return m
}

func (m *Math) AddStyles(ms map[string]string) *Math {
	for k, v := range ms {
		m.style[k] = v
	}
	return m
}

func (m *Math) Style(ms map[string]string) *Math {
	m.style = ms
	return m
}

func (m *Math) Add(e Elements) *Math {
	m.contents = append(m.contents, e.Bytes())
	return m
}

func (m *Math) Display(display string) *Math {
	m.display = display
	return m
}

func (m *Math) Xmlns(xmlns string) *Math {
	m.xmlns = xmlns
	return m
}

func (m *Math) Bytes() []byte {
	return m.buf.Bytes()
}

func (m *Math) Prepare() {
	m.buf.WriteString("<math")
	if m.xmlns != "" {
		m.buf.WriteString(" xmlns=\"" + m.xmlns + "\"")
	}
	if m.display != "" {
		m.buf.WriteString(" display=\"" + m.display + "\"")
	}
	if len(m.style) != 0 {
		idx := 0
		m.buf.WriteString(" style=\"")
		for k, v := range m.style {
			m.buf.WriteString(k + ": " + v + ";")
			if idx != len(m.style)-1 {
				m.buf.WriteByte(' ')
			}
			idx++
		}
		m.buf.WriteString("\"")
	}
	m.buf.WriteByte('>')

	for _, content := range m.contents {
		m.buf.Write(content)
	}
	m.buf.WriteString("</math>")
}
