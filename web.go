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

// Name sets the name attribute and returns the element.
func (s *Slot) Name(n string) *Slot {
	s.name = n
	return s
}

// renderTo writes the Slot component's HTML to buf.
func (s *Slot) renderTo(buf *bytes.Buffer) {
	tg := startTag(buf, "slot")
	tg.attr("name", s.name)
	s.attrs(tg)
	tg.children(s.contents)
}

// Template represents the HTML <template> element, markup that is not rendered when the page loads but can be cloned by script.
type Template struct {
	headElement
	bodyElement
	contentNode[*Template]
}

// NewTemplate creates a new Template component.
func NewTemplate() *Template {
	v := &Template{}
	v.init(v)
	return v
}

// renderTo writes the Template component's HTML to buf.
func (t *Template) renderTo(buf *bytes.Buffer) {
	tg := startTag(buf, "template")
	t.attrs(tg)
	tg.children(t.contents)
}
