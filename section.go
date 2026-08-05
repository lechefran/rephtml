package rephtml

import "bytes"

// Header represents the Header component or supporting type.
type Header struct {
	bodyElement
	contentNode[*Header]
}

// NewHeader creates a new Header component.
func NewHeader() *Header {
	v := &Header{}
	v.init(v)
	return v
}

// renderTo writes the Header component's HTML to buf.
func (h *Header) renderTo(buf *bytes.Buffer) {
	tg := startTag(buf, "header")
	tg.styleAttr(h.style)
	tg.children(h.contents)
}

// Nav represents the Nav component or supporting type.
type Nav struct {
	bodyElement
	contentNode[*Nav]
	role string
}

// NewNav creates a new Nav component.
func NewNav() *Nav {
	v := &Nav{}
	v.init(v)
	return v
}

// Role sets the role value on the Nav component.
func (n *Nav) Role(r string) *Nav {
	n.role = r
	return n
}

// renderTo writes the Nav component's HTML to buf.
func (n *Nav) renderTo(buf *bytes.Buffer) {
	tg := startTag(buf, "nav")
	tg.attr("role", n.role)
	tg.styleAttr(n.style)
	tg.children(n.contents)
}

// Main represents the Main component or supporting type.
type Main struct {
	bodyElement
	contentNode[*Main]
}

// NewMain creates a new Main component.
func NewMain() *Main {
	v := &Main{}
	v.init(v)
	return v
}

// renderTo writes the Main component's HTML to buf.
func (m *Main) renderTo(buf *bytes.Buffer) {
	tg := startTag(buf, "main")
	tg.styleAttr(m.style)
	tg.children(m.contents)
}

// Section represents the Section component or supporting type.
type Section struct {
	bodyElement
	contentNode[*Section]
	ariaLabel string
}

// NewSection creates a new Section component.
func NewSection() *Section {
	v := &Section{}
	v.init(v)
	return v
}

// AriaLabel sets the arialabel value on the Section component.
func (s *Section) AriaLabel(label string) *Section {
	s.ariaLabel = label
	return s
}

// renderTo writes the Section component's HTML to buf.
func (s *Section) renderTo(buf *bytes.Buffer) {
	tg := startTag(buf, "section")
	tg.attr("aria-label", s.ariaLabel)
	tg.styleAttr(s.style)
	tg.children(s.contents)
}

// Article represents the Article component or supporting type.
type Article struct {
	bodyElement
	contentNode[*Article]
}

// NewArticle creates a new Article component.
func NewArticle() *Article {
	v := &Article{}
	v.init(v)
	return v
}

// renderTo writes the Article component's HTML to buf.
func (a *Article) renderTo(buf *bytes.Buffer) {
	tg := startTag(buf, "article")
	tg.styleAttr(a.style)
	tg.children(a.contents)
}

// Aside represents the Aside component or supporting type.
type Aside struct {
	bodyElement
	contentNode[*Aside]
}

// NewAside creates a new Aside component.
func NewAside() *Aside {
	v := &Aside{}
	v.init(v)
	return v
}

// renderTo writes the Aside component's HTML to buf.
func (as *Aside) renderTo(buf *bytes.Buffer) {
	tg := startTag(buf, "aside")
	tg.styleAttr(as.style)
	tg.children(as.contents)
}

// Footer represents the Footer component or supporting type.
type Footer struct {
	bodyElement
	contentNode[*Footer]
}

// NewFooter creates a new Footer component.
func NewFooter() *Footer {
	v := &Footer{}
	v.init(v)
	return v
}

// renderTo writes the Footer component's HTML to buf.
func (f *Footer) renderTo(buf *bytes.Buffer) {
	tg := startTag(buf, "footer")
	tg.styleAttr(f.style)
	tg.children(f.contents)
}

// Address represents the Address component or supporting type.
type Address struct {
	bodyElement
	contentNode[*Address]
}

// NewAddress creates a new Address component.
func NewAddress() *Address {
	v := &Address{}
	v.init(v)
	return v
}

// renderTo writes the Address component's HTML to buf.
func (ad *Address) renderTo(buf *bytes.Buffer) {
	tg := startTag(buf, "address")
	tg.styleAttr(ad.style)
	tg.children(ad.contents)
}
