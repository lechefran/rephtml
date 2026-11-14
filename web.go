package rephtml

import "bytes"

// Slot represents the HTML <slot> element, a placeholder inside a web component that you fill with your own markup.
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

// Name sets the name attribute.
func (s *Slot) Name(n string) *Slot {
	s.name = n
	return s
}

// renderTo writes the Slot component's HTML to buf.
func (s *Slot) renderTo(buf *bytes.Buffer) {
	tg := startTag(buf, "slot")
	tg.attr("name", s.name)
	tg.styleAttr(s.style)
	tg.children(s.contents)
}

// Template represents the HTML <template> element, markup that is not rendered when the page loads but can be cloned by script.
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

// Id sets the id attribute.
func (t *Template) Id(i string) *Template {
	t.id = i
	return t
}

// renderTo writes the Template component's HTML to buf.
func (t *Template) renderTo(buf *bytes.Buffer) {
	tg := startTag(buf, "template")
	tg.attr("id", t.id)
	tg.styleAttr(t.style)
	tg.children(t.contents)
}
