package rephtml

import "bytes"

type Slot struct {
	buf      bytes.Buffer
	style    map[string]string
	contents []Element
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
	s.contents = appendElement(s.contents, e)
	return s
}

func (s *Slot) Name(n string) *Slot {
	s.name = n
	return s
}

func (s *Slot) Bytes() []byte {
	return cloneBytes(s.buf.Bytes())
}

// IsBodyElement implements BodyElement interface
func (s *Slot) IsBodyElement() {}

func (s *Slot) Prepare() {
	s.buf.Reset()
	s.buf.WriteString("<slot")
	if s.name != "" {
		s.buf.WriteString(" name=\"" + s.name + "\"")
	}
	if len(s.style) != 0 {
		parseStyle(&s.buf, s.style)
	}
	s.buf.WriteByte('>')

	writeElements(&s.buf, s.contents)
	s.buf.WriteString("</slot>")
}

type Template struct {
	buf      bytes.Buffer
	style    map[string]string
	contents []Element
	id       string
}

func NewTemplate() *Template {
	return &Template{
		style: make(map[string]string),
	}
}

// IsHeadElement implements HeadElement interface
func (t *Template) IsHeadElement() {}

// IsBodyElement implements BodyElement interface
func (t *Template) IsBodyElement() {}

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
	t.contents = appendElement(t.contents, e)
	return t
}

func (t *Template) Id(i string) *Template {
	t.id = i
	return t
}

func (t *Template) Bytes() []byte {
	return cloneBytes(t.buf.Bytes())
}

func (t *Template) Prepare() {
	t.buf.Reset()
	t.buf.WriteString("<template")
	if t.id != "" {
		t.buf.WriteString(" id=\"" + t.id + "\"")
	}
	if len(t.style) != 0 {
		parseStyle(&t.buf, t.style)
	}
	t.buf.WriteByte('>')

	writeElements(&t.buf, t.contents)
	t.buf.WriteString("</template>")
}
