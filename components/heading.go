package rephtml

import "bytes"

type Hgroup struct {
	buf      bytes.Buffer
	style    map[string]string
	contents []Element
}

func NewHgroup() *Hgroup {
	return &Hgroup{
		style: make(map[string]string),
	}
}

func (h *Hgroup) AddStyle(k, v string) *Hgroup {
	h.style[k] = v
	return h
}

func (h *Hgroup) AddStyles(m map[string]string) *Hgroup {
	for k, v := range m {
		h.style[k] = v
	}
	return h
}

func (h *Hgroup) Style(m map[string]string) *Hgroup {
	h.style = m
	return h
}

func (h *Hgroup) Add(e Element) *Hgroup {
	h.contents = appendElement(h.contents, e)
	return h
}

func (h *Hgroup) Bytes() []byte {
	return cloneBytes(h.buf.Bytes())
}

// IsBodyElement implements BodyElement interface
func (h *Hgroup) IsBodyElement() {}

func (h *Hgroup) Prepare() {
	h.buf.Reset()
	h.buf.WriteString("<hgroup")
	if len(h.style) != 0 {
		parseStyle(&h.buf, h.style)
	}
	h.buf.WriteByte('>')

	writeElements(&h.buf, h.contents)
	h.buf.WriteString("</hgroup>")
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
	return cloneBytes(h.buf.Bytes())
}

// IsBodyElement implements BodyElement interface
func (h *H1) IsBodyElement() {}

func (h *H1) Prepare() {
	h.buf.Reset()
	h.buf.WriteString("<h1")
	if len(h.style) != 0 {
		parseStyle(&h.buf, h.style)
	}
	h.buf.WriteString(">" + h.text + "</h1>")
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
	return cloneBytes(h.buf.Bytes())
}

// IsBodyElement implements BodyElement interface
func (h *H2) IsBodyElement() {}

func (h *H2) Prepare() {
	h.buf.Reset()
	h.buf.WriteString("<h2")
	if len(h.style) != 0 {
		parseStyle(&h.buf, h.style)
	}
	h.buf.WriteString(">" + h.text + "</h2>")
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
	return cloneBytes(h.buf.Bytes())
}

// IsBodyElement implements BodyElement interface
func (h *H3) IsBodyElement() {}

func (h *H3) Prepare() {
	h.buf.Reset()
	h.buf.WriteString("<h3")
	if len(h.style) != 0 {
		parseStyle(&h.buf, h.style)
	}
	h.buf.WriteString(">" + h.text + "</h3>")
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
	return cloneBytes(h.buf.Bytes())
}

// IsBodyElement implements BodyElement interface
func (h *H4) IsBodyElement() {}

func (h *H4) Prepare() {
	h.buf.Reset()
	h.buf.WriteString("<h4")
	if len(h.style) != 0 {
		parseStyle(&h.buf, h.style)
	}
	h.buf.WriteString(">" + h.text + "</h4>")
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
	return cloneBytes(h.buf.Bytes())
}

// IsBodyElement implements BodyElement interface
func (h *H5) IsBodyElement() {}

func (h *H5) Prepare() {
	h.buf.Reset()
	h.buf.WriteString("<h5")
	if len(h.style) != 0 {
		parseStyle(&h.buf, h.style)
	}
	h.buf.WriteString(">" + h.text + "</h5>")
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
	return cloneBytes(h.buf.Bytes())
}

// IsBodyElement implements BodyElement interface
func (h *H6) IsBodyElement() {}

func (h *H6) Prepare() {
	h.buf.Reset()
	h.buf.WriteString("<h6")
	if len(h.style) != 0 {
		parseStyle(&h.buf, h.style)
	}
	h.buf.WriteString(">" + h.text + "</h6>")
}
