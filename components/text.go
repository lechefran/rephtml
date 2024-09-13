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

func (p *P) Prepare() {
	if len(p.style) != 0 {
		idx := 0
		p.buf.WriteString("<p style=\"")
		for k, v := range p.style {
			p.buf.WriteString(k + ": " + v + ";")
			if idx != len(p.style)-1 {
				p.buf.WriteByte(' ')
			}
			idx++
		}
		p.buf.WriteString("\">" + p.text + "</p>")
	} else {
		p.buf.WriteString("<p>" + p.text + "</p>")
	}
}

type H1 struct {
	buf   bytes.Buffer
	style map[string]string
	text  string
}

func NewH1() *H1 {
	return &H1{
		style: make(map[string]string),
	}
}

func (h *H1) AddStyle(k, v string) *H1 {
	h.style[k] = v
	return h
}

func (h *H1) Style(m map[string]string) *H1 {
	h.style = m
	return h
}

func (h *H1) Text(s string) *H1 {
	h.text = s
	return h
}

func (h *H1) Bytes() []byte {
	return h.buf.Bytes()
}

func (h *H1) Prepare() {
	if len(h.style) != 0 {
		idx := 0
		h.buf.WriteString("<h1 style=\"")
		for k, v := range h.style {
			h.buf.WriteString(k + ": " + v + ";")
			if idx != len(h.style)-1 {
				h.buf.WriteByte(' ')
			}
			idx++
		}
		h.buf.WriteString("\">" + h.text + "</h1>")
	} else {
		h.buf.WriteString("<h1>" + h.text + "</h1>")
	}
}

type H2 struct {
	buf   bytes.Buffer
	style map[string]string
	text  string
}

func NewH2() *H2 {
	return &H2{
		style: make(map[string]string),
	}
}

func (h *H2) AddStyle(k, v string) *H2 {
	h.style[k] = v
	return h
}

func (h *H2) Style(m map[string]string) *H2 {
	h.style = m
	return h
}

func (h *H2) Text(s string) *H2 {
	h.text = s
	return h
}

func (h *H2) Bytes() []byte {
	return h.buf.Bytes()
}

func (h *H2) Prepare() {
	if len(h.style) != 0 {
		idx := 0
		h.buf.WriteString("<h2 style=\"")
		for k, v := range h.style {
			h.buf.WriteString(k + ": " + v + ";")
			if idx != len(h.style)-1 {
				h.buf.WriteByte(' ')
			}
			idx++
		}
		h.buf.WriteString("\">" + h.text + "</h2>")
	} else {
		h.buf.WriteString("<h2>" + h.text + "</h2>")
	}
}

type H3 struct {
	buf   bytes.Buffer
	style map[string]string
	text  string
}

func NewH3() *H3 {
	return &H3{
		style: make(map[string]string),
	}
}

func (h *H3) AddStyle(k, v string) *H3 {
	h.style[k] = v
	return h
}

func (h *H3) Style(m map[string]string) *H3 {
	h.style = m
	return h
}

func (h *H3) Text(s string) *H3 {
	h.text = s
	return h
}

func (h *H3) Bytes() []byte {
	return h.buf.Bytes()
}

func (h *H3) Prepare() {
	if len(h.style) != 0 {
		idx := 0
		h.buf.WriteString("<h3 style=\"")
		for k, v := range h.style {
			h.buf.WriteString(k + ": " + v + ";")
			if idx != len(h.style)-1 {
				h.buf.WriteByte(' ')
			}
			idx++
		}
		h.buf.WriteString("\">" + h.text + "</h3>")
	} else {
		h.buf.WriteString("<h3>" + h.text + "</h3>")
	}
}

type H4 struct {
	buf   bytes.Buffer
	style map[string]string
	text  string
}

func NewH4() *H4 {
	return &H4{
		style: make(map[string]string),
	}
}

func (h *H4) AddStyle(k, v string) *H4 {
	h.style[k] = v
	return h
}

func (h *H4) Style(m map[string]string) *H4 {
	h.style = m
	return h
}

func (h *H4) Text(s string) *H4 {
	h.text = s
	return h
}

func (h *H4) Bytes() []byte {
	return h.buf.Bytes()
}

func (h *H4) Prepare() {
	if len(h.style) != 0 {
		idx := 0
		h.buf.WriteString("<h4 style=\"")
		for k, v := range h.style {
			h.buf.WriteString(k + ": " + v + ";")
			if idx != len(h.style)-1 {
				h.buf.WriteByte(' ')
			}
			idx++
		}
		h.buf.WriteString("\">" + h.text + "</h4>")
	} else {
		h.buf.WriteString("<h4>" + h.text + "</h4>")
	}
}

type H5 struct {
	buf   bytes.Buffer
	style map[string]string
	text  string
}

func NewH5() *H5 {
	return &H5{
		style: make(map[string]string),
	}
}

func (h *H5) AddStyle(k, v string) *H5 {
	h.style[k] = v
	return h
}

func (h *H5) Style(m map[string]string) *H5 {
	h.style = m
	return h
}

func (h *H5) Text(s string) *H5 {
	h.text = s
	return h
}

func (h *H5) Bytes() []byte {
	return h.buf.Bytes()
}

func (h *H5) Prepare() {
	if len(h.style) != 0 {
		idx := 0
		h.buf.WriteString("<h5 style=\"")
		for k, v := range h.style {
			h.buf.WriteString(k + ": " + v + ";")
			if idx != len(h.style)-1 {
				h.buf.WriteByte(' ')
			}
			idx++
		}
		h.buf.WriteString("\">" + h.text + "</h5>")
	} else {
		h.buf.WriteString("<h5>" + h.text + "</h5>")
	}
}

type H6 struct {
	buf   bytes.Buffer
	style map[string]string
	text  string
}

func NewH6() *H6 {
	return &H6{
		style: make(map[string]string),
	}
}

func (h *H6) AddStyle(k, v string) *H6 {
	h.style[k] = v
	return h
}

func (h *H6) Style(m map[string]string) *H6 {
	h.style = m
	return h
}

func (h *H6) Text(s string) *H6 {
	h.text = s
	return h
}

func (h *H6) Bytes() []byte {
	return h.buf.Bytes()
}

func (h *H6) Prepare() {
	if len(h.style) != 0 {
		idx := 0
		h.buf.WriteString("<h6 style=\"")
		for k, v := range h.style {
			h.buf.WriteString(k + ": " + v + ";")
			if idx != len(h.style)-1 {
				h.buf.WriteByte(' ')
			}
			idx++
		}
		h.buf.WriteString("\">" + h.text + "</h6>")
	} else {
		h.buf.WriteString("<h6>" + h.text + "</h6>")
	}
}
