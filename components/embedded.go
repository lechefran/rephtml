package rephtml

import "bytes"

type Embed struct {
	buf    bytes.Buffer
	style  map[string]string
	src    string
	embedType string
	width  string
	height string
}

func NewEmbed() *Embed {
	return &Embed{
		style: make(map[string]string),
	}
}

func (e *Embed) AddStyle(k, v string) *Embed {
	e.style[k] = v
	return e
}

func (e *Embed) AddStyles(m map[string]string) *Embed {
	for k, v := range m {
		e.style[k] = v
	}
	return e
}

func (e *Embed) Style(m map[string]string) *Embed {
	e.style = m
	return e
}

func (e *Embed) Src(src string) *Embed {
	e.src = src
	return e
}

func (e *Embed) Type(embedType string) *Embed {
	e.embedType = embedType
	return e
}

func (e *Embed) Width(width string) *Embed {
	e.width = width
	return e
}

func (e *Embed) Height(height string) *Embed {
	e.height = height
	return e
}

func (e *Embed) Bytes() []byte {
	return e.buf.Bytes()
}

func (e *Embed) Prepare() {
	e.buf.WriteString("<embed")
	if e.src != "" {
		e.buf.WriteString(" src=\"" + e.src + "\"")
	}
	if e.embedType != "" {
		e.buf.WriteString(" type=\"" + e.embedType + "\"")
	}
	if e.width != "" {
		e.buf.WriteString(" width=\"" + e.width + "\"")
	}
	if e.height != "" {
		e.buf.WriteString(" height=\"" + e.height + "\"")
	}
	if len(e.style) != 0 {
		idx := 0
		e.buf.WriteString(" style=\"")
		for k, v := range e.style {
			e.buf.WriteString(k + ": " + v + ";")
			if idx != len(e.style)-1 {
				e.buf.WriteByte(' ')
			}
			idx++
		}
		e.buf.WriteString("\"")
	}
	e.buf.WriteString(">")
}
