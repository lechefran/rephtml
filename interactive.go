package rephtml

import "bytes"

// Details represents the HTML details element for creating a disclosure widget
type Details struct {
	bodyElement
	contentNode[*Details]
	open bool
}

// NewDetails creates a new Details element
func NewDetails() *Details {
	v := &Details{}
	v.init(v)
	return v
}

// renderTo writes the details element's HTML to buf.
func (d *Details) renderTo(buf *bytes.Buffer) {
	tg := startTag(buf, "details")
	tg.boolAttr("open", d.open)
	d.attrs(tg)
	tg.children(d.contents)
}

// Open sets the open attribute and returns the element.
func (d *Details) Open(open bool) *Details {
	d.open = open
	return d
}

// Dialog represents the HTML dialog element for creating modal dialogs
type Dialog struct {
	bodyElement
	contentNode[*Dialog]
	open bool
}

// NewDialog creates a new Dialog element
func NewDialog() *Dialog {
	v := &Dialog{}
	v.init(v)
	return v
}

// renderTo writes the dialog element's HTML to buf.
func (d *Dialog) renderTo(buf *bytes.Buffer) {
	tg := startTag(buf, "dialog")
	tg.boolAttr("open", d.open)
	d.attrs(tg)
	tg.children(d.contents)
}

// Open sets the open attribute and returns the element.
func (d *Dialog) Open(open bool) *Dialog {
	d.open = open
	return d
}

// Summary represents the HTML summary element for details disclosure summary
type Summary struct {
	bodyElement
	contentNode[*Summary]
}

// NewSummary creates a new Summary element
func NewSummary() *Summary {
	v := &Summary{}
	v.init(v)
	return v
}

// renderTo writes the summary element's HTML to buf.
func (s *Summary) renderTo(buf *bytes.Buffer) {
	tg := startTag(buf, "summary")
	s.attrs(tg)
	tg.children(s.contents)
}
