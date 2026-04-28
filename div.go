package rephtml

import (
	"bytes"
)

// Div represents the Div component or supporting type.
type Div struct {
	buf      bytes.Buffer
	contents []Element
	style    StyleMap
}

// NewDiv creates a new Div component.
func NewDiv() *Div {
	return &Div{
		style: make(StyleMap),
	}
}

// Add appends child content to the Div component.
func (d *Div) Add(e Element) *Div {
	d.contents = appendElement(d.contents, e)
	return d
}

// AddStyles adds multiple inline CSS declarations to the Div component.
func (d *Div) AddStyles(m StyleMap) *Div {
	for k, v := range m {
		d.style[k] = v
	}
	return d
}

// Bytes returns a defensive copy of the rendered Div bytes.
func (d *Div) Bytes() []byte {
	return cloneBytes(d.buf.Bytes())
}

// Render returns freshly prepared Div HTML bytes.
func (d *Div) Render() []byte {
	return renderPrepared(d)
}

// HTML returns freshly prepared Div HTML as a string.
func (d *Div) HTML() string {
	return htmlPrepared(d)
}

// String returns freshly prepared Div HTML as a string.
func (d *Div) String() string {
	return d.HTML()
}

// IsBodyElement implements BodyElement interface
func (d *Div) IsBodyElement() {}

// Prepare renders the Div component into its internal buffer.
func (d *Div) Prepare() {
	d.buf.Reset()
	d.buf.WriteString("<div")
	if len(d.style) != 0 {
		parseStyle(&d.buf, d.style)
	}
	d.buf.WriteString(">")

	writeElements(&d.buf, d.contents)
	d.buf.WriteString("</div>")
}
