package rephtml

// Slot represents the Slot component or supporting type.
type Slot struct {
	bodyElement
	contentNode[*Slot]
	name string
}

// NewSlot creates a new Slot component.
func NewSlot() *Slot {
	v := &Slot{}
	v.init(v)
	return v
}

// Name sets the name value on the Slot component.
func (s *Slot) Name(n string) *Slot {
	s.name = n
	return s
}

// prepare renders the Slot component into its internal buffer.
func (s *Slot) prepare() {
	tg := openTag(&s.buf, "slot")
	tg.attr("name", s.name)
	tg.styleAttr(s.style)
	tg.children(s.contents)
}

// Template represents the Template component or supporting type.
type Template struct {
	headElement
	bodyElement
	contentNode[*Template]
	id string
}

// NewTemplate creates a new Template component.
func NewTemplate() *Template {
	v := &Template{}
	v.init(v)
	return v
}

// Id sets the id value on the Template component.
func (t *Template) Id(i string) *Template {
	t.id = i
	return t
}

// prepare renders the Template component into its internal buffer.
func (t *Template) prepare() {
	tg := openTag(&t.buf, "template")
	tg.attr("id", t.id)
	tg.styleAttr(t.style)
	tg.children(t.contents)
}
