package rephtml

import "bytes"

type Slot struct {
	buf      bytes.Buffer
	style    map[string]string
	contents [][]byte
	name     string
}

func NewSlot() *Slot {
	return &Slot{
		style: make(map[string]string),
	}
}

func (s *Slot) AddStyle(k, v string) *Slot {
	s.style[k] = v
	return s
}

func (s *Slot) AddStyles(m map[string]string) *Slot {
	for k, v := range m {
		s.style[k] = v
	}
	return s
}

func (s *Slot) Style(m map[string]string) *Slot {
	s.style = m
	return s
}

func (s *Slot) Add(e Element) *Slot {
	s.contents = append(s.contents, e.Bytes())
	return s
}

func (s *Slot) Name(n string) *Slot {
	s.name = n
	return s
}

func (s *Slot) Bytes() []byte {
	return s.buf.Bytes()
}

func (s *Slot) Prepare() {
	s.buf.WriteString("<slot")
	if s.name != "" {
		s.buf.WriteString(" name=\"" + s.name + "\"")
	}
	if len(s.style) != 0 {
		idx := 0
		s.buf.WriteString(" style=\"")
		for k, v := range s.style {
			s.buf.WriteString(k + ": " + v + ";")
			if idx != len(s.style)-1 {
				s.buf.WriteByte(' ')
			}
			idx++
		}
		s.buf.WriteString("\"")
	}
	s.buf.WriteByte('>')

	for _, content := range s.contents {
		s.buf.Write(content)
	}
	s.buf.WriteString("</slot>")
}

type Template struct {
	buf      bytes.Buffer
	style    map[string]string
	contents [][]byte
	id       string
}

func NewTemplate() *Template {
	return &Template{
		style: make(map[string]string),
	}
}

func (t *Template) AddStyle(k, v string) *Template {
	t.style[k] = v
	return t
}

func (t *Template) AddStyles(m map[string]string) *Template {
	for k, v := range m {
		t.style[k] = v
	}
	return t
}

func (t *Template) Style(m map[string]string) *Template {
	t.style = m
	return t
}

func (t *Template) Add(e Element) *Template {
	t.contents = append(t.contents, e.Bytes())
	return t
}

func (t *Template) Id(i string) *Template {
	t.id = i
	return t
}

func (t *Template) Bytes() []byte {
	return t.buf.Bytes()
}

func (t *Template) Prepare() {
	t.buf.WriteString("<template")
	if t.id != "" {
		t.buf.WriteString(" id=\"" + t.id + "\"")
	}
	if len(t.style) != 0 {
		idx := 0
		t.buf.WriteString(" style=\"")
		for k, v := range t.style {
			t.buf.WriteString(k + ": " + v + ";")
			if idx != len(t.style)-1 {
				t.buf.WriteByte(' ')
			}
			idx++
		}
		t.buf.WriteString("\"")
	}
	t.buf.WriteByte('>')

	for _, content := range t.contents {
		t.buf.Write(content)
	}
	t.buf.WriteString("</template>")
}
