package rephtml

import "bytes"

// Div represents the HTML <div> element, a generic container for flow content with no semantic meaning.
type Div struct {
	bodyElement
	contentNode[*Div]
}

// NewDiv returns a new, empty <div> element.
func NewDiv() *Div {
	v := &Div{}
	v.init(v)
	return v
}

// renderTo writes the Div component's HTML to buf.
func (d *Div) renderTo(buf *bytes.Buffer) {
	tg := startTag(buf, "div")
	d.attrs(tg)
	tg.children(d.contents)
}
