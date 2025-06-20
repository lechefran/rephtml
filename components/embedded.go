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

type Iframe struct {
	buf             bytes.Buffer
	style           map[string]string
	src             string
	width           string
	height          string
	name            string
	sandbox         string
	allow           string
	allowfullscreen bool
	loading         string
	referrerpolicy  string
	srcdoc          string
}

func NewIframe() *Iframe {
	return &Iframe{
		style: make(map[string]string),
	}
}

func (i *Iframe) AddStyle(k, v string) *Iframe {
	i.style[k] = v
	return i
}

func (i *Iframe) AddStyles(m map[string]string) *Iframe {
	for k, v := range m {
		i.style[k] = v
	}
	return i
}

func (i *Iframe) Style(m map[string]string) *Iframe {
	i.style = m
	return i
}

func (i *Iframe) Src(src string) *Iframe {
	i.src = src
	return i
}

func (i *Iframe) Width(width string) *Iframe {
	i.width = width
	return i
}

func (i *Iframe) Height(height string) *Iframe {
	i.height = height
	return i
}

func (i *Iframe) Name(name string) *Iframe {
	i.name = name
	return i
}

func (i *Iframe) Sandbox(sandbox string) *Iframe {
	i.sandbox = sandbox
	return i
}

func (i *Iframe) Allow(allow string) *Iframe {
	i.allow = allow
	return i
}

func (i *Iframe) Allowfullscreen(allowfullscreen bool) *Iframe {
	i.allowfullscreen = allowfullscreen
	return i
}

func (i *Iframe) Loading(loading string) *Iframe {
	i.loading = loading
	return i
}

func (i *Iframe) Referrerpolicy(referrerpolicy string) *Iframe {
	i.referrerpolicy = referrerpolicy
	return i
}

func (i *Iframe) Srcdoc(srcdoc string) *Iframe {
	i.srcdoc = srcdoc
	return i
}

func (i *Iframe) Bytes() []byte {
	return i.buf.Bytes()
}

func (i *Iframe) Prepare() {
	i.buf.WriteString("<iframe")
	if i.src != "" {
		i.buf.WriteString(" src=\"" + i.src + "\"")
	}
	if i.width != "" {
		i.buf.WriteString(" width=\"" + i.width + "\"")
	}
	if i.height != "" {
		i.buf.WriteString(" height=\"" + i.height + "\"")
	}
	if i.name != "" {
		i.buf.WriteString(" name=\"" + i.name + "\"")
	}
	if i.sandbox != "" {
		i.buf.WriteString(" sandbox=\"" + i.sandbox + "\"")
	}
	if i.allow != "" {
		i.buf.WriteString(" allow=\"" + i.allow + "\"")
	}
	if i.allowfullscreen {
		i.buf.WriteString(" allowfullscreen")
	}
	if i.loading != "" {
		i.buf.WriteString(" loading=\"" + i.loading + "\"")
	}
	if i.referrerpolicy != "" {
		i.buf.WriteString(" referrerpolicy=\"" + i.referrerpolicy + "\"")
	}
	if i.srcdoc != "" {
		i.buf.WriteString(" srcdoc=\"" + i.srcdoc + "\"")
	}
	if len(i.style) != 0 {
		idx := 0
		i.buf.WriteString(" style=\"")
		for k, v := range i.style {
			i.buf.WriteString(k + ": " + v + ";")
			if idx != len(i.style)-1 {
				i.buf.WriteByte(' ')
			}
			idx++
		}
		i.buf.WriteString("\"")
	}
	i.buf.WriteString("></iframe>")
}
