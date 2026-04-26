package rephtml

import "bytes"

// Hgroup represents the Hgroup component or supporting type.
type Hgroup struct {
	buf      bytes.Buffer
	style    map[string]string
	contents []Element
}

// NewHgroup creates a new Hgroup component.
func NewHgroup() *Hgroup {
	return &Hgroup{
		style: make(map[string]string),
	}
}

// AddStyle adds one inline CSS declaration to the Hgroup component.
func (h *Hgroup) AddStyle(k, v string) *Hgroup {
	h.style[k] = v
	return h
}

// AddStyles adds multiple inline CSS declarations to the Hgroup component.
func (h *Hgroup) AddStyles(m map[string]string) *Hgroup {
	for k, v := range m {
		h.style[k] = v
	}
	return h
}

// Style replaces the inline CSS declarations on the Hgroup component.
func (h *Hgroup) Style(m map[string]string) *Hgroup {
	h.style = cloneStyleMap(m)
	return h
}

// Add appends child content to the Hgroup component.
func (h *Hgroup) Add(e Element) *Hgroup {
	h.contents = appendElement(h.contents, e)
	return h
}

// Bytes returns a defensive copy of the rendered Hgroup bytes.
func (h *Hgroup) Bytes() []byte {
	return cloneBytes(h.buf.Bytes())
}

// IsBodyElement implements BodyElement interface
func (h *Hgroup) IsBodyElement() {}

// Prepare renders the Hgroup component into its internal buffer.
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

// H1 represents the H1 component or supporting type.
type H1 struct {
	buf   bytes.Buffer
	style map[string]string
	text  string
}

// NewH1 creates a new H1 component.
func NewH1() *H1 {
	return &H1{
		style: make(map[string]string),
	}
}

// AddStyle adds one inline CSS declaration to the H1 component.
func (h *H1) AddStyle(k, v string) *H1 {
	h.style[k] = v
	return h
}

// Style replaces the inline CSS declarations on the H1 component.
func (h *H1) Style(m map[string]string) *H1 {
	h.style = cloneStyleMap(m)
	return h
}

// Text sets or appends text content on the H1 component.
func (h *H1) Text(s string) *H1 {
	h.text = s
	return h
}

// Bytes returns a defensive copy of the rendered H1 bytes.
func (h *H1) Bytes() []byte {
	return cloneBytes(h.buf.Bytes())
}

// IsBodyElement implements BodyElement interface
func (h *H1) IsBodyElement() {}

// Prepare renders the H1 component into its internal buffer.
func (h *H1) Prepare() {
	h.buf.Reset()
	h.buf.WriteString("<h1")
	if len(h.style) != 0 {
		parseStyle(&h.buf, h.style)
	}
	h.buf.WriteString(">" + escapeText(h.text) + "</h1>")
}

// H2 represents the H2 component or supporting type.
type H2 struct {
	buf   bytes.Buffer
	style map[string]string
	text  string
}

// NewH2 creates a new H2 component.
func NewH2() *H2 {
	return &H2{
		style: make(map[string]string),
	}
}

// AddStyle adds one inline CSS declaration to the H2 component.
func (h *H2) AddStyle(k, v string) *H2 {
	h.style[k] = v
	return h
}

// Style replaces the inline CSS declarations on the H2 component.
func (h *H2) Style(m map[string]string) *H2 {
	h.style = cloneStyleMap(m)
	return h
}

// Text sets or appends text content on the H2 component.
func (h *H2) Text(s string) *H2 {
	h.text = s
	return h
}

// Bytes returns a defensive copy of the rendered H2 bytes.
func (h *H2) Bytes() []byte {
	return cloneBytes(h.buf.Bytes())
}

// IsBodyElement implements BodyElement interface
func (h *H2) IsBodyElement() {}

// Prepare renders the H2 component into its internal buffer.
func (h *H2) Prepare() {
	h.buf.Reset()
	h.buf.WriteString("<h2")
	if len(h.style) != 0 {
		parseStyle(&h.buf, h.style)
	}
	h.buf.WriteString(">" + escapeText(h.text) + "</h2>")
}

// H3 represents the H3 component or supporting type.
type H3 struct {
	buf   bytes.Buffer
	style map[string]string
	text  string
}

// NewH3 creates a new H3 component.
func NewH3() *H3 {
	return &H3{
		style: make(map[string]string),
	}
}

// AddStyle adds one inline CSS declaration to the H3 component.
func (h *H3) AddStyle(k, v string) *H3 {
	h.style[k] = v
	return h
}

// Style replaces the inline CSS declarations on the H3 component.
func (h *H3) Style(m map[string]string) *H3 {
	h.style = cloneStyleMap(m)
	return h
}

// Text sets or appends text content on the H3 component.
func (h *H3) Text(s string) *H3 {
	h.text = s
	return h
}

// Bytes returns a defensive copy of the rendered H3 bytes.
func (h *H3) Bytes() []byte {
	return cloneBytes(h.buf.Bytes())
}

// IsBodyElement implements BodyElement interface
func (h *H3) IsBodyElement() {}

// Prepare renders the H3 component into its internal buffer.
func (h *H3) Prepare() {
	h.buf.Reset()
	h.buf.WriteString("<h3")
	if len(h.style) != 0 {
		parseStyle(&h.buf, h.style)
	}
	h.buf.WriteString(">" + escapeText(h.text) + "</h3>")
}

// H4 represents the H4 component or supporting type.
type H4 struct {
	buf   bytes.Buffer
	style map[string]string
	text  string
}

// NewH4 creates a new H4 component.
func NewH4() *H4 {
	return &H4{
		style: make(map[string]string),
	}
}

// AddStyle adds one inline CSS declaration to the H4 component.
func (h *H4) AddStyle(k, v string) *H4 {
	h.style[k] = v
	return h
}

// Style replaces the inline CSS declarations on the H4 component.
func (h *H4) Style(m map[string]string) *H4 {
	h.style = cloneStyleMap(m)
	return h
}

// Text sets or appends text content on the H4 component.
func (h *H4) Text(s string) *H4 {
	h.text = s
	return h
}

// Bytes returns a defensive copy of the rendered H4 bytes.
func (h *H4) Bytes() []byte {
	return cloneBytes(h.buf.Bytes())
}

// IsBodyElement implements BodyElement interface
func (h *H4) IsBodyElement() {}

// Prepare renders the H4 component into its internal buffer.
func (h *H4) Prepare() {
	h.buf.Reset()
	h.buf.WriteString("<h4")
	if len(h.style) != 0 {
		parseStyle(&h.buf, h.style)
	}
	h.buf.WriteString(">" + escapeText(h.text) + "</h4>")
}

// H5 represents the H5 component or supporting type.
type H5 struct {
	buf   bytes.Buffer
	style map[string]string
	text  string
}

// NewH5 creates a new H5 component.
func NewH5() *H5 {
	return &H5{
		style: make(map[string]string),
	}
}

// AddStyle adds one inline CSS declaration to the H5 component.
func (h *H5) AddStyle(k, v string) *H5 {
	h.style[k] = v
	return h
}

// Style replaces the inline CSS declarations on the H5 component.
func (h *H5) Style(m map[string]string) *H5 {
	h.style = cloneStyleMap(m)
	return h
}

// Text sets or appends text content on the H5 component.
func (h *H5) Text(s string) *H5 {
	h.text = s
	return h
}

// Bytes returns a defensive copy of the rendered H5 bytes.
func (h *H5) Bytes() []byte {
	return cloneBytes(h.buf.Bytes())
}

// IsBodyElement implements BodyElement interface
func (h *H5) IsBodyElement() {}

// Prepare renders the H5 component into its internal buffer.
func (h *H5) Prepare() {
	h.buf.Reset()
	h.buf.WriteString("<h5")
	if len(h.style) != 0 {
		parseStyle(&h.buf, h.style)
	}
	h.buf.WriteString(">" + escapeText(h.text) + "</h5>")
}

// H6 represents the H6 component or supporting type.
type H6 struct {
	buf   bytes.Buffer
	style map[string]string
	text  string
}

// NewH6 creates a new H6 component.
func NewH6() *H6 {
	return &H6{
		style: make(map[string]string),
	}
}

// AddStyle adds one inline CSS declaration to the H6 component.
func (h *H6) AddStyle(k, v string) *H6 {
	h.style[k] = v
	return h
}

// Style replaces the inline CSS declarations on the H6 component.
func (h *H6) Style(m map[string]string) *H6 {
	h.style = cloneStyleMap(m)
	return h
}

// Text sets or appends text content on the H6 component.
func (h *H6) Text(s string) *H6 {
	h.text = s
	return h
}

// Bytes returns a defensive copy of the rendered H6 bytes.
func (h *H6) Bytes() []byte {
	return cloneBytes(h.buf.Bytes())
}

// IsBodyElement implements BodyElement interface
func (h *H6) IsBodyElement() {}

// Prepare renders the H6 component into its internal buffer.
func (h *H6) Prepare() {
	h.buf.Reset()
	h.buf.WriteString("<h6")
	if len(h.style) != 0 {
		parseStyle(&h.buf, h.style)
	}
	h.buf.WriteString(">" + escapeText(h.text) + "</h6>")
}
