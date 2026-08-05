package rephtml

import "bytes"

// Del represents the HTML del element for marking deleted text
type Del struct {
	bodyElement
	contentNode[*Del]
	cite     string
	datetime string
}

// NewDel creates a new Del element
func NewDel() *Del {
	v := &Del{}
	v.init(v)
	return v
}

// Prepare builds the HTML for the del element
func (d *Del) renderTo(buf *bytes.Buffer) {
	tg := startTag(buf, "del")
	tg.urlAttr("cite", d.cite)
	tg.attr("datetime", d.datetime)
	tg.styleAttr(d.style)
	tg.children(d.contents)
}

// Cite sets the cite attribute
func (d *Del) Cite(url string) *Del {
	d.cite = url
	return d
}

// Datetime sets the datetime attribute
func (d *Del) Datetime(datetime string) *Del {
	d.datetime = datetime
	return d
}

// Ins represents the HTML ins element for marking inserted text
type Ins struct {
	bodyElement
	contentNode[*Ins]
	cite     string
	datetime string
}

// NewIns creates a new Ins element
func NewIns() *Ins {
	v := &Ins{}
	v.init(v)
	return v
}

// Prepare builds the HTML for the ins element
func (i *Ins) renderTo(buf *bytes.Buffer) {
	tg := startTag(buf, "ins")
	tg.urlAttr("cite", i.cite)
	tg.attr("datetime", i.datetime)
	tg.styleAttr(i.style)
	tg.children(i.contents)
}

// Cite sets the cite attribute
func (i *Ins) Cite(url string) *Ins {
	i.cite = url
	return i
}

// Datetime sets the datetime attribute
func (i *Ins) Datetime(datetime string) *Ins {
	i.datetime = datetime
	return i
}
