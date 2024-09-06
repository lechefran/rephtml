package rephtml

import "bytes"

type Div struct {
	buf      bytes.Buffer
	contents [][]byte
	style    map[string]string
}

func NewDiv() *Div {
	return &Div{
		style: make(map[string]string),
	}
}

func (d *Div) Add(e Elements) *Div {
	d.contents = append(d.contents, e.Bytes())
	return d
}

func (d *Div) AddStyles(m map[string]string) *Div {
	for k, v := range m {
		d.style[k] = v
	}
	return d
}

func (d *Div) Bytes() []byte {
	return d.buf.Bytes()
}

func (d *Div) Prepare() *Div {
	if len(d.style) != 0 {
		idx := 0
		d.buf.WriteString("<div style=\"")
		for k, v := range d.style {
			d.buf.WriteString(k + ": " + v + ";")
			if idx != len(d.style)-1 {
				d.buf.WriteByte(' ')
			}
			idx++
		}
		d.buf.WriteString("\">")
	} else {
		d.buf.WriteString("<div>")
	}

	for _, c := range d.contents {
		d.buf.Write(c)
	}
	d.buf.WriteString("</div>")
	return d
}
