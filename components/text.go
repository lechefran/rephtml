package rephtml

import "bytes"

type P struct {
	buf   bytes.Buffer
	style map[string]string
	text  string
}

func NewP() *P {
	return &P{
		style: make(map[string]string),
	}
}

func (p *P) AddStyle(k, v string) *P {
	p.style[k] = v
	return p
}

func (p *P) AddStyles(m map[string]string) *P {
	for k, v := range m {
		p.style[k] = v
	}
	return p
}

func (p *P) Style(m map[string]string) *P {
	p.style = m
	return p
}

func (p *P) Text(s string) *P {
	p.text = s
	return p
}

func (p *P) Bytes() []byte {
	return p.buf.Bytes()
}

func (p *P) Prepare() *P {
	if len(p.style) != 0 {
		idx := 0
		p.buf.WriteString("<p style=\"")
		for k, v := range p.style {
			p.buf.WriteString(k + ": " + v + ";")
			if idx != len(p.style)-1 {
				if idx != len(p.style)-1 {
					p.buf.WriteString(" ")
				}
			}
			idx++
		}
		p.buf.WriteString("\">" + p.text + "</p>")
	} else {
		p.buf.WriteString("<p>" + p.text + "</p>")
	}
	return p
}
