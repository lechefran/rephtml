package rephtml

import "bytes"

// Slot represents the Slot component or supporting type.
type Slot struct {
	buf      bytes.Buffer
	style    StyleMap
	contents []Element
	name     string
}

// NewSlot creates a new Slot component.
func NewSlot() *Slot {
	return &Slot{
		style: make(StyleMap),
	}
}

// AddStyle adds one inline CSS declaration to the Slot component.
func (s *Slot) AddStyle(k, v string) *Slot {
	s.style[k] = v
	return s
}

// AddStyles adds multiple inline CSS declarations to the Slot component.
func (s *Slot) AddStyles(m StyleMap) *Slot {
	for k, v := range m {
		s.style[k] = v
	}
	return s
}

// Style replaces the inline CSS declarations on the Slot component.
func (s *Slot) Style(m StyleMap) *Slot {
	s.style = cloneStyleMap(m)
	return s
}

// Add appends child content to the Slot component.
func (s *Slot) Add(e Element) *Slot {
	s.contents = appendElement(s.contents, e)
	return s
}

// Name sets the name value on the Slot component.
func (s *Slot) Name(n string) *Slot {
	s.name = n
	return s
}

// Bytes returns a defensive copy of the rendered Slot bytes.
func (s *Slot) Bytes() []byte {
	return cloneBytes(s.buf.Bytes())
}

// IsBodyElement implements BodyElement interface
func (s *Slot) IsBodyElement() {}

// Prepare renders the Slot component into its internal buffer.
func (s *Slot) Prepare() {
	s.buf.Reset()
	s.buf.WriteString("<slot")
	if s.name != "" {
		writeAttr(&s.buf, "name", s.name)
	}
	if len(s.style) != 0 {
		parseStyle(&s.buf, s.style)
	}
	s.buf.WriteByte('>')

	writeElements(&s.buf, s.contents)
	s.buf.WriteString("</slot>")
}

// Template represents the Template component or supporting type.
type Template struct {
	buf      bytes.Buffer
	style    StyleMap
	contents []Element
	id       string
}

// NewTemplate creates a new Template component.
func NewTemplate() *Template {
	return &Template{
		style: make(StyleMap),
	}
}

// IsHeadElement implements HeadElement interface
func (t *Template) IsHeadElement() {}

// IsBodyElement implements BodyElement interface
func (t *Template) IsBodyElement() {}

// AddStyle adds one inline CSS declaration to the Template component.
func (t *Template) AddStyle(k, v string) *Template {
	t.style[k] = v
	return t
}

// AddStyles adds multiple inline CSS declarations to the Template component.
func (t *Template) AddStyles(m StyleMap) *Template {
	for k, v := range m {
		t.style[k] = v
	}
	return t
}

// Style replaces the inline CSS declarations on the Template component.
func (t *Template) Style(m StyleMap) *Template {
	t.style = cloneStyleMap(m)
	return t
}

// Add appends child content to the Template component.
func (t *Template) Add(e Element) *Template {
	t.contents = appendElement(t.contents, e)
	return t
}

// Id sets the id value on the Template component.
func (t *Template) Id(i string) *Template {
	t.id = i
	return t
}

// Bytes returns a defensive copy of the rendered Template bytes.
func (t *Template) Bytes() []byte {
	return cloneBytes(t.buf.Bytes())
}

// Prepare renders the Template component into its internal buffer.
func (t *Template) Prepare() {
	t.buf.Reset()
	t.buf.WriteString("<template")
	if t.id != "" {
		writeAttr(&t.buf, "id", t.id)
	}
	if len(t.style) != 0 {
		parseStyle(&t.buf, t.style)
	}
	t.buf.WriteByte('>')

	writeElements(&t.buf, t.contents)
	t.buf.WriteString("</template>")
}
