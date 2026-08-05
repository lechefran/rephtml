package rephtml

import (
	"bytes"
	"strings"
)

// P represents the P component or supporting type.
type P struct {
	bodyElement
	textNode[*P]
}

// NewP creates a new P component.
func NewP() *P {
	v := &P{}
	v.init(v)
	return v
}

// renderTo writes the P component's HTML to buf.
func (p *P) renderTo(buf *bytes.Buffer) {
	tg := startTag(buf, "p")
	tg.styleAttr(p.style)
	tg.text(p.text)
}

// Comment represents an HTML comment.
//
// Comment text is sanitised rather than escaped. Character references are not
// decoded inside a comment, so escaping would both fail to stop the text from
// closing the comment early and leave the entities visible in the output.
type Comment struct {
	text string
}

// NewComment creates a new Comment component.
func NewComment() *Comment {
	return &Comment{}
}

// Text sets the comment's text content.
func (c *Comment) Text(s string) *Comment {
	c.text = s
	return c
}

// Render returns the comment's HTML bytes.
func (c *Comment) Render() []byte {
	return renderElement(c)
}

// HTML returns the comment's HTML as a string.
func (c *Comment) HTML() string {
	return htmlElement(c)
}

// IsBodyElement implements BodyElement.
func (c *Comment) IsBodyElement() {}

// renderTo writes the Comment component's HTML to buf.
func (c *Comment) renderTo(buf *bytes.Buffer) {
	buf.WriteString("<!--")
	buf.WriteString(sanitizeComment(c.text))
	buf.WriteString("-->")
}

// sanitizeComment makes text safe to place inside an HTML comment.
//
// A comment ends at the first "-->", and "--!>" closes it too, so any run of
// two hyphens is broken up with a space. Comment data also may not begin with
// ">" or "->", which a leading space prevents. The text is otherwise left
// alone, since nothing else inside a comment is interpreted.
func sanitizeComment(text string) string {
	if text == "" {
		return ""
	}

	var b strings.Builder
	b.Grow(len(text) + 2)

	if text[0] == '>' || strings.HasPrefix(text, "->") {
		b.WriteByte(' ')
	}

	for i := 0; i < len(text); i++ {
		if text[i] == '-' && i+1 < len(text) && text[i+1] == '-' {
			b.WriteString("- ")
			continue
		}
		b.WriteByte(text[i])
	}

	// A trailing hyphen would pair with the closing "-->" delimiter.
	if strings.HasSuffix(b.String(), "-") {
		b.WriteByte(' ')
	}
	return b.String()
}

// Hr represents the Hr component or supporting type.
type Hr struct {
	bodyElement
	node[*Hr]
}

// NewHr creates a new Hr component.
func NewHr() *Hr {
	v := &Hr{}
	v.init(v)
	return v
}

// renderTo writes the Hr component's HTML to buf.
func (h *Hr) renderTo(buf *bytes.Buffer) {
	tg := startTag(buf, "hr")
	tg.styleAttr(h.style)
	tg.void()
}

// Pre represents the Pre component or supporting type.
type Pre struct {
	bodyElement
	textNode[*Pre]
}

// NewPre creates a new Pre component.
func NewPre() *Pre {
	v := &Pre{}
	v.init(v)
	return v
}

// renderTo writes the Pre component's HTML to buf.
func (p *Pre) renderTo(buf *bytes.Buffer) {
	tg := startTag(buf, "pre")
	tg.styleAttr(p.style)
	tg.text(p.text)
}

// Blockquote represents the Blockquote component or supporting type.
type Blockquote struct {
	bodyElement
	textNode[*Blockquote]
	cite string
}

// NewBlockquote creates a new Blockquote component.
func NewBlockquote() *Blockquote {
	v := &Blockquote{}
	v.init(v)
	return v
}

// Cite sets the cite value on the Blockquote component.
func (b *Blockquote) Cite(c string) *Blockquote {
	b.cite = c
	return b
}

// renderTo writes the Blockquote component's HTML to buf.
func (b *Blockquote) renderTo(buf *bytes.Buffer) {
	tg := startTag(buf, "blockquote")
	tg.urlAttr("cite", b.cite)
	tg.styleAttr(b.style)
	tg.text(b.text)
}

// Menu represents the Menu component or supporting type.
type Menu struct {
	bodyElement
	contentNode[*Menu]
	menuType string
	label    string
}

// NewMenu creates a new Menu component.
func NewMenu() *Menu {
	v := &Menu{}
	v.init(v)
	return v
}

// Type sets the type value on the Menu component.
func (m *Menu) Type(t string) *Menu {
	m.menuType = t
	return m
}

// Label sets the label value on the Menu component.
func (m *Menu) Label(l string) *Menu {
	m.label = l
	return m
}

// renderTo writes the Menu component's HTML to buf.
func (m *Menu) renderTo(buf *bytes.Buffer) {
	tg := startTag(buf, "menu")
	tg.attr("type", m.menuType)
	tg.attr("label", m.label)
	tg.styleAttr(m.style)
	tg.children(m.contents)
}

// Ol represents the Ol component or supporting type.
type Ol struct {
	bodyElement
	contentNode[*Ol]
	start    int
	listType string
	reversed bool
}

// NewOl creates a new Ol component.
func NewOl() *Ol {
	v := &Ol{}
	v.init(v)
	return v
}

// Start sets the start value on the Ol component.
func (o *Ol) Start(s int) *Ol {
	o.start = s
	return o
}

// Type sets the type value on the Ol component.
func (o *Ol) Type(t string) *Ol {
	o.listType = t
	return o
}

// Reversed sets the reversed value on the Ol component.
func (o *Ol) Reversed(r bool) *Ol {
	o.reversed = r
	return o
}

// renderTo writes the Ol component's HTML to buf.
func (o *Ol) renderTo(buf *bytes.Buffer) {
	tg := startTag(buf, "ol")
	tg.intAttr("start", o.start)
	tg.attr("type", o.listType)
	tg.boolAttr("reversed", o.reversed)
	tg.styleAttr(o.style)
	tg.children(o.contents)
}

// Ul represents the Ul component or supporting type.
type Ul struct {
	bodyElement
	contentNode[*Ul]
}

// NewUl creates a new Ul component.
func NewUl() *Ul {
	v := &Ul{}
	v.init(v)
	return v
}

// renderTo writes the Ul component's HTML to buf.
func (u *Ul) renderTo(buf *bytes.Buffer) {
	tg := startTag(buf, "ul")
	tg.styleAttr(u.style)
	tg.children(u.contents)
}

// Li represents the Li component or supporting type.
type Li struct {
	bodyElement
	contentNode[*Li]
	value int
}

// NewLi creates a new Li component.
func NewLi() *Li {
	v := &Li{}
	v.init(v)
	return v
}

// Value sets the value value on the Li component.
func (l *Li) Value(v int) *Li {
	l.value = v
	return l
}

// renderTo writes the Li component's HTML to buf.
func (l *Li) renderTo(buf *bytes.Buffer) {
	tg := startTag(buf, "li")
	tg.intAttr("value", l.value)
	tg.styleAttr(l.style)
	tg.children(l.contents)
}

// Dl represents the Dl component or supporting type.
type Dl struct {
	bodyElement
	contentNode[*Dl]
}

// NewDl creates a new Dl component.
func NewDl() *Dl {
	v := &Dl{}
	v.init(v)
	return v
}

// renderTo writes the Dl component's HTML to buf.
func (d *Dl) renderTo(buf *bytes.Buffer) {
	tg := startTag(buf, "dl")
	tg.styleAttr(d.style)
	tg.children(d.contents)
}

// Dt represents the Dt component or supporting type.
type Dt struct {
	bodyElement
	contentNode[*Dt]
}

// NewDt creates a new Dt component.
func NewDt() *Dt {
	v := &Dt{}
	v.init(v)
	return v
}

// renderTo writes the Dt component's HTML to buf.
func (d *Dt) renderTo(buf *bytes.Buffer) {
	tg := startTag(buf, "dt")
	tg.styleAttr(d.style)
	tg.children(d.contents)
}

// Dd represents the Dd component or supporting type.
type Dd struct {
	bodyElement
	contentNode[*Dd]
}

// NewDd creates a new Dd component.
func NewDd() *Dd {
	v := &Dd{}
	v.init(v)
	return v
}

// renderTo writes the Dd component's HTML to buf.
func (d *Dd) renderTo(buf *bytes.Buffer) {
	tg := startTag(buf, "dd")
	tg.styleAttr(d.style)
	tg.children(d.contents)
}

// Figure represents the Figure component or supporting type.
type Figure struct {
	bodyElement
	contentNode[*Figure]
}

// NewFigure creates a new Figure component.
func NewFigure() *Figure {
	v := &Figure{}
	v.init(v)
	return v
}

// renderTo writes the Figure component's HTML to buf.
func (f *Figure) renderTo(buf *bytes.Buffer) {
	tg := startTag(buf, "figure")
	tg.styleAttr(f.style)
	tg.children(f.contents)
}

// Figcaption represents the Figcaption component or supporting type.
type Figcaption struct {
	bodyElement
	contentNode[*Figcaption]
}

// NewFigcaption creates a new Figcaption component.
func NewFigcaption() *Figcaption {
	v := &Figcaption{}
	v.init(v)
	return v
}

// renderTo writes the Figcaption component's HTML to buf.
func (f *Figcaption) renderTo(buf *bytes.Buffer) {
	tg := startTag(buf, "figcaption")
	tg.styleAttr(f.style)
	tg.children(f.contents)
}

// Search represents the Search component or supporting type.
type Search struct {
	bodyElement
	contentNode[*Search]
}

// NewSearch creates a new Search component.
func NewSearch() *Search {
	v := &Search{}
	v.init(v)
	return v
}

// renderTo writes the Search component's HTML to buf.
func (s *Search) renderTo(buf *bytes.Buffer) {
	tg := startTag(buf, "search")
	tg.styleAttr(s.style)
	tg.children(s.contents)
}
