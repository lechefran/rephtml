package rephtml

import "bytes"

// Hgroup represents the Hgroup component or supporting type.
type Hgroup struct {
	bodyElement
	contentNode[*Hgroup]
}

// NewHgroup creates a new Hgroup component.
func NewHgroup() *Hgroup {
	v := &Hgroup{}
	v.init(v)
	return v
}

// renderTo writes the Hgroup component's HTML to buf.
func (h *Hgroup) renderTo(buf *bytes.Buffer) {
	tg := startTag(buf, "hgroup")
	tg.styleAttr(h.style)
	tg.children(h.contents)
}

// H1 represents the H1 component or supporting type.
type H1 struct {
	bodyElement
	textNode[*H1]
}

// NewH1 creates a new H1 component.
func NewH1() *H1 {
	v := &H1{}
	v.init(v)
	return v
}

// renderTo writes the H1 component's HTML to buf.
func (h *H1) renderTo(buf *bytes.Buffer) {
	tg := startTag(buf, "h1")
	tg.styleAttr(h.style)
	tg.text(h.text)
}

// H2 represents the H2 component or supporting type.
type H2 struct {
	bodyElement
	textNode[*H2]
}

// NewH2 creates a new H2 component.
func NewH2() *H2 {
	v := &H2{}
	v.init(v)
	return v
}

// renderTo writes the H2 component's HTML to buf.
func (h *H2) renderTo(buf *bytes.Buffer) {
	tg := startTag(buf, "h2")
	tg.styleAttr(h.style)
	tg.text(h.text)
}

// H3 represents the H3 component or supporting type.
type H3 struct {
	bodyElement
	textNode[*H3]
}

// NewH3 creates a new H3 component.
func NewH3() *H3 {
	v := &H3{}
	v.init(v)
	return v
}

// renderTo writes the H3 component's HTML to buf.
func (h *H3) renderTo(buf *bytes.Buffer) {
	tg := startTag(buf, "h3")
	tg.styleAttr(h.style)
	tg.text(h.text)
}

// H4 represents the H4 component or supporting type.
type H4 struct {
	bodyElement
	textNode[*H4]
}

// NewH4 creates a new H4 component.
func NewH4() *H4 {
	v := &H4{}
	v.init(v)
	return v
}

// renderTo writes the H4 component's HTML to buf.
func (h *H4) renderTo(buf *bytes.Buffer) {
	tg := startTag(buf, "h4")
	tg.styleAttr(h.style)
	tg.text(h.text)
}

// H5 represents the H5 component or supporting type.
type H5 struct {
	bodyElement
	textNode[*H5]
}

// NewH5 creates a new H5 component.
func NewH5() *H5 {
	v := &H5{}
	v.init(v)
	return v
}

// renderTo writes the H5 component's HTML to buf.
func (h *H5) renderTo(buf *bytes.Buffer) {
	tg := startTag(buf, "h5")
	tg.styleAttr(h.style)
	tg.text(h.text)
}

// H6 represents the H6 component or supporting type.
type H6 struct {
	bodyElement
	textNode[*H6]
}

// NewH6 creates a new H6 component.
func NewH6() *H6 {
	v := &H6{}
	v.init(v)
	return v
}

// renderTo writes the H6 component's HTML to buf.
func (h *H6) renderTo(buf *bytes.Buffer) {
	tg := startTag(buf, "h6")
	tg.styleAttr(h.style)
	tg.text(h.text)
}
