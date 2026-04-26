package rephtml

import "bytes"

// P represents the P component or supporting type.
type P struct {
	buf   bytes.Buffer
	style map[string]string
	text  string
}

// NewP creates a new P component.
func NewP() *P {
	return &P{
		style: make(map[string]string),
	}
}

// AddStyle adds one inline CSS declaration to the P component.
func (p *P) AddStyle(k, v string) *P {
	p.style[k] = v
	return p
}

// AddStyles adds multiple inline CSS declarations to the P component.
func (p *P) AddStyles(m map[string]string) *P {
	for k, v := range m {
		p.style[k] = v
	}
	return p
}

// Style replaces the inline CSS declarations on the P component.
func (p *P) Style(m map[string]string) *P {
	p.style = cloneStyleMap(m)
	return p
}

// Text sets or appends text content on the P component.
func (p *P) Text(s string) *P {
	p.text = s
	return p
}

// Bytes returns a defensive copy of the rendered P bytes.
func (p *P) Bytes() []byte {
	return cloneBytes(p.buf.Bytes())
}

// IsBodyElement implements BodyElement interface
func (p *P) IsBodyElement() {}

// Prepare renders the P component into its internal buffer.
func (p *P) Prepare() {
	p.buf.Reset()
	p.buf.WriteString("<p")
	if len(p.style) != 0 {
		parseStyle(&p.buf, p.style)
	}
	p.buf.WriteString(">" + escapeText(p.text) + "</p>")
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

// Bytes returns a defensive copy of the rendered Comment bytes.
func (c *Comment) Bytes() []byte {
	return cloneBytes(c.buf.Bytes())
}

// IsBodyElement implements BodyElement interface
func (c *Comment) IsBodyElement() {}

// Prepare renders the Comment component into its internal buffer.
func (c *Comment) Prepare() {
	c.buf.Reset()
	c.buf.WriteString("<!--" + c.text + "-->")
}

// Hr represents the Hr component or supporting type.
type Hr struct {
	buf   bytes.Buffer
	style map[string]string
}

// NewHr creates a new Hr component.
func NewHr() *Hr {
	return &Hr{
		style: make(map[string]string),
	}
}

// AddStyle adds one inline CSS declaration to the Hr component.
func (h *Hr) AddStyle(k, v string) *Hr {
	h.style[k] = v
	return h
}

// AddStyles adds multiple inline CSS declarations to the Hr component.
func (h *Hr) AddStyles(m map[string]string) *Hr {
	for k, v := range m {
		h.style[k] = v
	}
	return h
}

// Style replaces the inline CSS declarations on the Hr component.
func (h *Hr) Style(m map[string]string) *Hr {
	h.style = cloneStyleMap(m)
	return h
}

// Bytes returns a defensive copy of the rendered Hr bytes.
func (h *Hr) Bytes() []byte {
	return cloneBytes(h.buf.Bytes())
}

// IsBodyElement implements BodyElement interface
func (h *Hr) IsBodyElement() {}

// Prepare renders the Hr component into its internal buffer.
func (h *Hr) Prepare() {
	h.buf.Reset()
	h.buf.WriteString("<hr")
	if len(h.style) != 0 {
		parseStyle(&h.buf, h.style)
	}
	h.buf.WriteByte('>')
}

// Pre represents the Pre component or supporting type.
type Pre struct {
	buf   bytes.Buffer
	style map[string]string
	text  string
}

// NewPre creates a new Pre component.
func NewPre() *Pre {
	return &Pre{
		style: make(map[string]string),
	}
}

// AddStyle adds one inline CSS declaration to the Pre component.
func (p *Pre) AddStyle(k, v string) *Pre {
	p.style[k] = v
	return p
}

// AddStyles adds multiple inline CSS declarations to the Pre component.
func (p *Pre) AddStyles(m map[string]string) *Pre {
	for k, v := range m {
		p.style[k] = v
	}
	return p
}

// Style replaces the inline CSS declarations on the Pre component.
func (p *Pre) Style(m map[string]string) *Pre {
	p.style = cloneStyleMap(m)
	return p
}

// Text sets or appends text content on the Pre component.
func (p *Pre) Text(str string) *Pre {
	p.text = str
	return p
}

// Bytes returns a defensive copy of the rendered Pre bytes.
func (p *Pre) Bytes() []byte {
	return cloneBytes(p.buf.Bytes())
}

// IsBodyElement implements BodyElement interface
func (p *Pre) IsBodyElement() {}

// Prepare renders the Pre component into its internal buffer.
func (p *Pre) Prepare() {
	p.buf.Reset()
	p.buf.WriteString("<pre")
	if len(p.style) != 0 {
		parseStyle(&p.buf, p.style)
	}
	p.buf.WriteString(">" + escapeText(p.text) + "</pre>")
}

// Blockquote represents the Blockquote component or supporting type.
type Blockquote struct {
	buf   bytes.Buffer
	style map[string]string
	text  string
	cite  string
}

// NewBlockquote creates a new Blockquote component.
func NewBlockquote() *Blockquote {
	return &Blockquote{
		style: make(map[string]string),
	}
}

// AddStyle adds one inline CSS declaration to the Blockquote component.
func (b *Blockquote) AddStyle(k, v string) *Blockquote {
	b.style[k] = v
	return b
}

// AddStyles adds multiple inline CSS declarations to the Blockquote component.
func (b *Blockquote) AddStyles(m map[string]string) *Blockquote {
	for k, v := range m {
		b.style[k] = v
	}
	return b
}

// Style replaces the inline CSS declarations on the Blockquote component.
func (b *Blockquote) Style(m map[string]string) *Blockquote {
	b.style = cloneStyleMap(m)
	return b
}

// Text sets or appends text content on the Blockquote component.
func (b *Blockquote) Text(str string) *Blockquote {
	b.text = str
	return b
}

// Cite sets the cite value on the Blockquote component.
func (b *Blockquote) Cite(c string) *Blockquote {
	b.cite = c
	return b
}

// Bytes returns a defensive copy of the rendered Blockquote bytes.
func (b *Blockquote) Bytes() []byte {
	return cloneBytes(b.buf.Bytes())
}

// IsBodyElement implements BodyElement interface
func (b *Blockquote) IsBodyElement() {}

// Prepare renders the Blockquote component into its internal buffer.
func (b *Blockquote) Prepare() {
	b.buf.Reset()
	b.buf.WriteString("<blockquote")
	if b.cite != "" {
		writeAttr(&b.buf, "cite", b.cite)
	}
	if len(b.style) != 0 {
		parseStyle(&b.buf, b.style)
	}
	b.buf.WriteString(">" + escapeText(b.text) + "</blockquote>")
}

// Menu represents the Menu component or supporting type.
type Menu struct {
	buf      bytes.Buffer
	style    map[string]string
	contents []Element
	menuType string
	label    string
}

// NewMenu creates a new Menu component.
func NewMenu() *Menu {
	return &Menu{
		style: make(map[string]string),
	}
}

// AddStyle adds one inline CSS declaration to the Menu component.
func (m *Menu) AddStyle(k, v string) *Menu {
	m.style[k] = v
	return m
}

// AddStyles adds multiple inline CSS declarations to the Menu component.
func (m *Menu) AddStyles(ms map[string]string) *Menu {
	for k, v := range ms {
		m.style[k] = v
	}
	return m
}

// Style replaces the inline CSS declarations on the Menu component.
func (m *Menu) Style(ms map[string]string) *Menu {
	m.style = cloneStyleMap(ms)
	return m
}

// Add appends child content to the Menu component.
func (m *Menu) Add(e Element) *Menu {
	m.contents = appendElement(m.contents, e)
	return m
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

// Bytes returns a defensive copy of the rendered Menu bytes.
func (m *Menu) Bytes() []byte {
	return cloneBytes(m.buf.Bytes())
}

// IsBodyElement implements BodyElement interface
func (m *Menu) IsBodyElement() {}

// Prepare renders the Menu component into its internal buffer.
func (m *Menu) Prepare() {
	m.buf.Reset()
	m.buf.WriteString("<menu")
	if m.menuType != "" {
		writeAttr(&m.buf, "type", m.menuType)
	}
	if m.label != "" {
		writeAttr(&m.buf, "label", m.label)
	}
	if len(m.style) != 0 {
		parseStyle(&m.buf, m.style)
	}
	m.buf.WriteByte('>')

	writeElements(&m.buf, m.contents)
	m.buf.WriteString("</menu>")
}

// Ol represents the Ol component or supporting type.
type Ol struct {
	buf      bytes.Buffer
	style    map[string]string
	contents []Element
	start    int
	listType string
	reversed bool
}

// NewOl creates a new Ol component.
func NewOl() *Ol {
	return &Ol{
		style: make(map[string]string),
	}
}

// AddStyle adds one inline CSS declaration to the Ol component.
func (o *Ol) AddStyle(k, v string) *Ol {
	o.style[k] = v
	return o
}

// AddStyles adds multiple inline CSS declarations to the Ol component.
func (o *Ol) AddStyles(m map[string]string) *Ol {
	for k, v := range m {
		o.style[k] = v
	}
	return o
}

// Style replaces the inline CSS declarations on the Ol component.
func (o *Ol) Style(m map[string]string) *Ol {
	o.style = cloneStyleMap(m)
	return o
}

// Add appends child content to the Ol component.
func (o *Ol) Add(e Element) *Ol {
	o.contents = appendElement(o.contents, e)
	return o
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

// Bytes returns a defensive copy of the rendered Ol bytes.
func (o *Ol) Bytes() []byte {
	return cloneBytes(o.buf.Bytes())
}

// IsBodyElement implements BodyElement interface
func (o *Ol) IsBodyElement() {}

// Prepare renders the Ol component into its internal buffer.
func (o *Ol) Prepare() {
	o.buf.Reset()
	o.buf.WriteString("<ol")
	if o.start > 0 {
		writeIntAttr(&o.buf, "start", o.start)
	}
	if o.listType != "" {
		writeAttr(&o.buf, "type", o.listType)
	}
	if o.reversed {
		o.buf.WriteString(" reversed")
	}
	if len(o.style) != 0 {
		parseStyle(&o.buf, o.style)
	}
	o.buf.WriteByte('>')

	writeElements(&o.buf, o.contents)
	o.buf.WriteString("</ol>")
}

// Ul represents the Ul component or supporting type.
type Ul struct {
	buf      bytes.Buffer
	style    map[string]string
	contents []Element
}

// NewUl creates a new Ul component.
func NewUl() *Ul {
	return &Ul{
		style: make(map[string]string),
	}
}

// AddStyle adds one inline CSS declaration to the Ul component.
func (u *Ul) AddStyle(k, v string) *Ul {
	u.style[k] = v
	return u
}

// AddStyles adds multiple inline CSS declarations to the Ul component.
func (u *Ul) AddStyles(m map[string]string) *Ul {
	for k, v := range m {
		u.style[k] = v
	}
	return u
}

// Style replaces the inline CSS declarations on the Ul component.
func (u *Ul) Style(m map[string]string) *Ul {
	u.style = cloneStyleMap(m)
	return u
}

// Add appends child content to the Ul component.
func (u *Ul) Add(e Element) *Ul {
	u.contents = appendElement(u.contents, e)
	return u
}

// Bytes returns a defensive copy of the rendered Ul bytes.
func (u *Ul) Bytes() []byte {
	return cloneBytes(u.buf.Bytes())
}

// IsBodyElement implements BodyElement interface
func (u *Ul) IsBodyElement() {}

// Prepare renders the Ul component into its internal buffer.
func (u *Ul) Prepare() {
	u.buf.Reset()
	u.buf.WriteString("<ul")
	if len(u.style) != 0 {
		parseStyle(&u.buf, u.style)
	}
	u.buf.WriteByte('>')

	writeElements(&u.buf, u.contents)
	u.buf.WriteString("</ul>")
}

// Li represents the Li component or supporting type.
type Li struct {
	buf      bytes.Buffer
	style    map[string]string
	contents []Element
	value    int
}

// NewLi creates a new Li component.
func NewLi() *Li {
	return &Li{
		style: make(map[string]string),
	}
}

// AddStyle adds one inline CSS declaration to the Li component.
func (l *Li) AddStyle(k, v string) *Li {
	l.style[k] = v
	return l
}

// AddStyles adds multiple inline CSS declarations to the Li component.
func (l *Li) AddStyles(m map[string]string) *Li {
	for k, v := range m {
		l.style[k] = v
	}
	return l
}

// Style replaces the inline CSS declarations on the Li component.
func (l *Li) Style(m map[string]string) *Li {
	l.style = cloneStyleMap(m)
	return l
}

// Add appends child content to the Li component.
func (l *Li) Add(e Element) *Li {
	l.contents = appendElement(l.contents, e)
	return l
}

// Value sets the value value on the Li component.
func (l *Li) Value(v int) *Li {
	l.value = v
	return l
}

// Bytes returns a defensive copy of the rendered Li bytes.
func (l *Li) Bytes() []byte {
	return cloneBytes(l.buf.Bytes())
}

// IsBodyElement implements BodyElement interface
func (l *Li) IsBodyElement() {}

// Prepare renders the Li component into its internal buffer.
func (l *Li) Prepare() {
	l.buf.Reset()
	l.buf.WriteString("<li")
	if l.value > 0 {
		writeIntAttr(&l.buf, "value", l.value)
	}
	if len(l.style) != 0 {
		parseStyle(&l.buf, l.style)
	}
	l.buf.WriteByte('>')

	writeElements(&l.buf, l.contents)
	l.buf.WriteString("</li>")
}

// Dl represents the Dl component or supporting type.
type Dl struct {
	buf      bytes.Buffer
	style    map[string]string
	contents []Element
}

// NewDl creates a new Dl component.
func NewDl() *Dl {
	return &Dl{
		style: make(map[string]string),
	}
}

// AddStyle adds one inline CSS declaration to the Dl component.
func (d *Dl) AddStyle(k, v string) *Dl {
	d.style[k] = v
	return d
}

// AddStyles adds multiple inline CSS declarations to the Dl component.
func (d *Dl) AddStyles(m map[string]string) *Dl {
	for k, v := range m {
		d.style[k] = v
	}
	return d
}

// Style replaces the inline CSS declarations on the Dl component.
func (d *Dl) Style(m map[string]string) *Dl {
	d.style = cloneStyleMap(m)
	return d
}

// Add appends child content to the Dl component.
func (d *Dl) Add(e Element) *Dl {
	d.contents = appendElement(d.contents, e)
	return d
}

// Bytes returns a defensive copy of the rendered Dl bytes.
func (d *Dl) Bytes() []byte {
	return cloneBytes(d.buf.Bytes())
}

// IsBodyElement implements BodyElement interface
func (d *Dl) IsBodyElement() {}

// Prepare renders the Dl component into its internal buffer.
func (d *Dl) Prepare() {
	d.buf.Reset()
	d.buf.WriteString("<dl")
	if len(d.style) != 0 {
		parseStyle(&d.buf, d.style)
	}
	d.buf.WriteByte('>')

	writeElements(&d.buf, d.contents)
	d.buf.WriteString("</dl>")
}

// Dt represents the Dt component or supporting type.
type Dt struct {
	buf      bytes.Buffer
	style    map[string]string
	contents []Element
}

// NewDt creates a new Dt component.
func NewDt() *Dt {
	return &Dt{
		style: make(map[string]string),
	}
}

// AddStyle adds one inline CSS declaration to the Dt component.
func (d *Dt) AddStyle(k, v string) *Dt {
	d.style[k] = v
	return d
}

// AddStyles adds multiple inline CSS declarations to the Dt component.
func (d *Dt) AddStyles(m map[string]string) *Dt {
	for k, v := range m {
		d.style[k] = v
	}
	return d
}

// Style replaces the inline CSS declarations on the Dt component.
func (d *Dt) Style(m map[string]string) *Dt {
	d.style = cloneStyleMap(m)
	return d
}

// Add appends child content to the Dt component.
func (d *Dt) Add(e Element) *Dt {
	d.contents = appendElement(d.contents, e)
	return d
}

// Bytes returns a defensive copy of the rendered Dt bytes.
func (d *Dt) Bytes() []byte {
	return cloneBytes(d.buf.Bytes())
}

// IsBodyElement implements BodyElement interface
func (d *Dt) IsBodyElement() {}

// Prepare renders the Dt component into its internal buffer.
func (d *Dt) Prepare() {
	d.buf.Reset()
	d.buf.WriteString("<dt")
	if len(d.style) != 0 {
		parseStyle(&d.buf, d.style)
	}
	d.buf.WriteByte('>')

	writeElements(&d.buf, d.contents)
	d.buf.WriteString("</dt>")
}

// Dd represents the Dd component or supporting type.
type Dd struct {
	buf      bytes.Buffer
	style    map[string]string
	contents []Element
}

// NewDd creates a new Dd component.
func NewDd() *Dd {
	return &Dd{
		style: make(map[string]string),
	}
}

// AddStyle adds one inline CSS declaration to the Dd component.
func (d *Dd) AddStyle(k, v string) *Dd {
	d.style[k] = v
	return d
}

// AddStyles adds multiple inline CSS declarations to the Dd component.
func (d *Dd) AddStyles(m map[string]string) *Dd {
	for k, v := range m {
		d.style[k] = v
	}
	return d
}

// Style replaces the inline CSS declarations on the Dd component.
func (d *Dd) Style(m map[string]string) *Dd {
	d.style = cloneStyleMap(m)
	return d
}

// Add appends child content to the Dd component.
func (d *Dd) Add(e Element) *Dd {
	d.contents = appendElement(d.contents, e)
	return d
}

// Bytes returns a defensive copy of the rendered Dd bytes.
func (d *Dd) Bytes() []byte {
	return cloneBytes(d.buf.Bytes())
}

// IsBodyElement implements BodyElement interface
func (d *Dd) IsBodyElement() {}

// Prepare renders the Dd component into its internal buffer.
func (d *Dd) Prepare() {
	d.buf.Reset()
	d.buf.WriteString("<dd")
	if len(d.style) != 0 {
		parseStyle(&d.buf, d.style)
	}
	d.buf.WriteByte('>')

	writeElements(&d.buf, d.contents)
	d.buf.WriteString("</dd>")
}

// Figure represents the Figure component or supporting type.
type Figure struct {
	buf      bytes.Buffer
	style    map[string]string
	contents []Element
}

// NewFigure creates a new Figure component.
func NewFigure() *Figure {
	return &Figure{
		style: make(map[string]string),
	}
}

// AddStyle adds one inline CSS declaration to the Figure component.
func (f *Figure) AddStyle(k, v string) *Figure {
	f.style[k] = v
	return f
}

// AddStyles adds multiple inline CSS declarations to the Figure component.
func (f *Figure) AddStyles(m map[string]string) *Figure {
	for k, v := range m {
		f.style[k] = v
	}
	return f
}

// Style replaces the inline CSS declarations on the Figure component.
func (f *Figure) Style(m map[string]string) *Figure {
	f.style = cloneStyleMap(m)
	return f
}

// Add appends child content to the Figure component.
func (f *Figure) Add(e Element) *Figure {
	f.contents = appendElement(f.contents, e)
	return f
}

// Bytes returns a defensive copy of the rendered Figure bytes.
func (f *Figure) Bytes() []byte {
	return cloneBytes(f.buf.Bytes())
}

// IsBodyElement implements BodyElement interface
func (f *Figure) IsBodyElement() {}

// Prepare renders the Figure component into its internal buffer.
func (f *Figure) Prepare() {
	f.buf.Reset()
	f.buf.WriteString("<figure")
	if len(f.style) != 0 {
		parseStyle(&f.buf, f.style)
	}
	f.buf.WriteByte('>')

	writeElements(&f.buf, f.contents)
	f.buf.WriteString("</figure>")
}

// Figcaption represents the Figcaption component or supporting type.
type Figcaption struct {
	buf      bytes.Buffer
	style    map[string]string
	contents []Element
}

// NewFigcaption creates a new Figcaption component.
func NewFigcaption() *Figcaption {
	return &Figcaption{
		style: make(map[string]string),
	}
}

// AddStyle adds one inline CSS declaration to the Figcaption component.
func (f *Figcaption) AddStyle(k, v string) *Figcaption {
	f.style[k] = v
	return f
}

// AddStyles adds multiple inline CSS declarations to the Figcaption component.
func (f *Figcaption) AddStyles(m map[string]string) *Figcaption {
	for k, v := range m {
		f.style[k] = v
	}
	return f
}

// Style replaces the inline CSS declarations on the Figcaption component.
func (f *Figcaption) Style(m map[string]string) *Figcaption {
	f.style = cloneStyleMap(m)
	return f
}

// Add appends child content to the Figcaption component.
func (f *Figcaption) Add(e Element) *Figcaption {
	f.contents = appendElement(f.contents, e)
	return f
}

// Bytes returns a defensive copy of the rendered Figcaption bytes.
func (f *Figcaption) Bytes() []byte {
	return cloneBytes(f.buf.Bytes())
}

// IsBodyElement implements BodyElement interface
func (f *Figcaption) IsBodyElement() {}

// Prepare renders the Figcaption component into its internal buffer.
func (f *Figcaption) Prepare() {
	f.buf.Reset()
	f.buf.WriteString("<figcaption")
	if len(f.style) != 0 {
		parseStyle(&f.buf, f.style)
	}
	f.buf.WriteByte('>')

	writeElements(&f.buf, f.contents)
	f.buf.WriteString("</figcaption>")
}

// Search represents the Search component or supporting type.
type Search struct {
	buf      bytes.Buffer
	style    map[string]string
	contents []Element
}

// NewSearch creates a new Search component.
func NewSearch() *Search {
	return &Search{
		style: make(map[string]string),
	}
}

// AddStyle adds one inline CSS declaration to the Search component.
func (s *Search) AddStyle(k, v string) *Search {
	s.style[k] = v
	return s
}

// AddStyles adds multiple inline CSS declarations to the Search component.
func (s *Search) AddStyles(m map[string]string) *Search {
	for k, v := range m {
		s.style[k] = v
	}
	return s
}

// Style replaces the inline CSS declarations on the Search component.
func (s *Search) Style(m map[string]string) *Search {
	s.style = cloneStyleMap(m)
	return s
}

// Add appends child content to the Search component.
func (s *Search) Add(e Element) *Search {
	s.contents = appendElement(s.contents, e)
	return s
}

// Bytes returns a defensive copy of the rendered Search bytes.
func (s *Search) Bytes() []byte {
	return cloneBytes(s.buf.Bytes())
}

// IsBodyElement implements BodyElement interface
func (s *Search) IsBodyElement() {}

// Prepare renders the Search component into its internal buffer.
func (s *Search) Prepare() {
	s.buf.Reset()
	s.buf.WriteString("<search")
	if len(s.style) != 0 {
		parseStyle(&s.buf, s.style)
	}
	s.buf.WriteByte('>')

	writeElements(&s.buf, s.contents)
	s.buf.WriteString("</search>")
}
