package rephtml

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

// prepare renders the Div component into its internal buffer.
func (d *Div) prepare() {
	tg := openTag(&d.buf, "div")
	tg.styleAttr(d.style)
	tg.children(d.contents)
}
