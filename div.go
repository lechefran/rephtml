package rephtml

import "bytes"

// Div represents the Div component or supporting type.
type Div struct {
	bodyElement
	contentNode[*Div]
}

// NewDiv creates a new Div component.
func NewDiv() *Div {
	v := &Div{}
	v.init(v)
	return v
}

// renderTo writes the Div component's HTML to buf.
func (d *Div) renderTo(buf *bytes.Buffer) {
	tg := startTag(buf, "div")
	tg.styleAttr(d.style)
	tg.children(d.contents)
}
