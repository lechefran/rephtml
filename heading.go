package rephtml

import "bytes"

// Hgroup represents the HTML <hgroup> element, a heading together with related content such as a subheading.
type Hgroup struct {
	bodyElement
	contentNode[*Hgroup]
}

// NewHgroup returns a new, empty <hgroup> element.
func NewHgroup() *Hgroup {
	v := &Hgroup{}
	v.init(v)
	return v
}

// renderTo writes the Hgroup component's HTML to buf.
func (h *Hgroup) renderTo(buf *bytes.Buffer) {
	tg := startTag(buf, "hgroup")
	h.attrs(tg)
	tg.children(h.contents)
}

// H1 represents the HTML <h1> element, a top-level section heading.
type H1 struct {
	bodyElement
	textNode[*H1]
}

// NewH1 returns a new, empty <h1> element.
func NewH1() *H1 {
	v := &H1{}
	v.initTag(v, "h1")
	return v
}

// H2 represents the HTML <h2> element, a second-level section heading.
type H2 struct {
	bodyElement
	textNode[*H2]
}

// NewH2 returns a new, empty <h2> element.
func NewH2() *H2 {
	v := &H2{}
	v.initTag(v, "h2")
	return v
}

// H3 represents the HTML <h3> element, a third-level section heading.
type H3 struct {
	bodyElement
	textNode[*H3]
}

// NewH3 returns a new, empty <h3> element.
func NewH3() *H3 {
	v := &H3{}
	v.initTag(v, "h3")
	return v
}

// H4 represents the HTML <h4> element, a fourth-level section heading.
type H4 struct {
	bodyElement
	textNode[*H4]
}

// NewH4 returns a new, empty <h4> element.
func NewH4() *H4 {
	v := &H4{}
	v.initTag(v, "h4")
	return v
}

// H5 represents the HTML <h5> element, a fifth-level section heading.
type H5 struct {
	bodyElement
	textNode[*H5]
}

// NewH5 returns a new, empty <h5> element.
func NewH5() *H5 {
	v := &H5{}
	v.initTag(v, "h5")
	return v
}

// H6 represents the HTML <h6> element, a sixth-level section heading.
type H6 struct {
	bodyElement
	textNode[*H6]
}

// NewH6 returns a new, empty <h6> element.
func NewH6() *H6 {
	v := &H6{}
	v.initTag(v, "h6")
	return v
}
