package rephtml

import "bytes"

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

// prepare renders the P component into its internal buffer.
func (p *P) prepare() {
	tg := openTag(&p.buf, "p")
	tg.styleAttr(p.style)
	tg.text(p.text)
}

// Comment represents the Comment component or supporting type.
type Comment struct {
	buf  bytes.Buffer
	text string
}

// NewComment creates a new Comment component.
func NewComment() *Comment {
	return &Comment{}
}

// Text sets or appends text content on the Comment component.
func (c *Comment) Text(s string) *Comment {
	c.text = s
	return c
}

// rawBytes returns the prepared Comment bytes without copying them.
func (c *Comment) rawBytes() []byte {
	return c.buf.Bytes()
}

// Render returns freshly prepared Comment HTML bytes.
func (c *Comment) Render() []byte {
	return renderPrepared(c)
}

// HTML returns freshly prepared Comment HTML as a string.
func (c *Comment) HTML() string {
	return htmlPrepared(c)
}

// prepare renders the Comment component into its internal buffer.
func (c *Comment) prepare() {
	c.buf.Reset()
	c.buf.WriteString("<!--" + c.text + "-->")
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

// prepare renders the Hr component into its internal buffer.
func (h *Hr) prepare() {
	tg := openTag(&h.buf, "hr")
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

// prepare renders the Pre component into its internal buffer.
func (p *Pre) prepare() {
	tg := openTag(&p.buf, "pre")
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

// prepare renders the Blockquote component into its internal buffer.
func (b *Blockquote) prepare() {
	tg := openTag(&b.buf, "blockquote")
	tg.attr("cite", b.cite)
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

// prepare renders the Menu component into its internal buffer.
func (m *Menu) prepare() {
	tg := openTag(&m.buf, "menu")
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

// prepare renders the Ol component into its internal buffer.
func (o *Ol) prepare() {
	tg := openTag(&o.buf, "ol")
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

// prepare renders the Ul component into its internal buffer.
func (u *Ul) prepare() {
	tg := openTag(&u.buf, "ul")
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

// prepare renders the Li component into its internal buffer.
func (l *Li) prepare() {
	tg := openTag(&l.buf, "li")
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

// prepare renders the Dl component into its internal buffer.
func (d *Dl) prepare() {
	tg := openTag(&d.buf, "dl")
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

// prepare renders the Dt component into its internal buffer.
func (d *Dt) prepare() {
	tg := openTag(&d.buf, "dt")
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

// prepare renders the Dd component into its internal buffer.
func (d *Dd) prepare() {
	tg := openTag(&d.buf, "dd")
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

// prepare renders the Figure component into its internal buffer.
func (f *Figure) prepare() {
	tg := openTag(&f.buf, "figure")
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

// prepare renders the Figcaption component into its internal buffer.
func (f *Figcaption) prepare() {
	tg := openTag(&f.buf, "figcaption")
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

// prepare renders the Search component into its internal buffer.
func (s *Search) prepare() {
	tg := openTag(&s.buf, "search")
	tg.styleAttr(s.style)
	tg.children(s.contents)
}

// IsBodyElement implements the marker interface.
func (c *Comment) IsBodyElement() {}
