package rephtml

import "bytes"

type P struct {
	buf   bytes.Buffer
	style map[string]string
	text  string
}

func NewP() *P {
	return &P{
		style: make(map[string]string),
	}
}

func (p *P) AddStyle(k, v string) *P {
	p.style[k] = v
	return p
}

func (p *P) AddStyles(m map[string]string) *P {
	for k, v := range m {
		p.style[k] = v
	}
	return p
}

func (p *P) Style(m map[string]string) *P {
	p.style = m
	return p
}

func (p *P) Text(s string) *P {
	p.text = s
	return p
}

func (p *P) Bytes() []byte {
	return cloneBytes(p.buf.Bytes())
}

// IsBodyElement implements BodyElement interface
func (p *P) IsBodyElement() {}

func (p *P) Prepare() {
	p.buf.Reset()
	p.buf.WriteString("<p")
	if len(p.style) != 0 {
		parseStyle(&p.buf, p.style)
	}
	p.buf.WriteString(">" + p.text + "</p>")
}

type Comment struct {
	buf  bytes.Buffer
	text string
}

func NewComment() *Comment {
	return &Comment{}
}

func (c *Comment) Text(s string) *Comment {
	c.text = s
	return c
}

func (c *Comment) Bytes() []byte {
	return cloneBytes(c.buf.Bytes())
}

// IsBodyElement implements BodyElement interface
func (c *Comment) IsBodyElement() {}

func (c *Comment) Prepare() {
	c.buf.Reset()
	c.buf.WriteString("<!--" + c.text + "-->")
}

type Hr struct {
	buf   bytes.Buffer
	style map[string]string
}

func NewHr() *Hr {
	return &Hr{
		style: make(map[string]string),
	}
}

func (h *Hr) AddStyle(k, v string) *Hr {
	h.style[k] = v
	return h
}

func (h *Hr) AddStyles(m map[string]string) *Hr {
	for k, v := range m {
		h.style[k] = v
	}
	return h
}

func (h *Hr) Style(m map[string]string) *Hr {
	h.style = m
	return h
}

func (h *Hr) Bytes() []byte {
	return cloneBytes(h.buf.Bytes())
}

// IsBodyElement implements BodyElement interface
func (h *Hr) IsBodyElement() {}

func (h *Hr) Prepare() {
	h.buf.Reset()
	h.buf.WriteString("<hr")
	if len(h.style) != 0 {
		parseStyle(&h.buf, h.style)
	}
	h.buf.WriteByte('>')
}

type Pre struct {
	buf   bytes.Buffer
	style map[string]string
	text  string
}

func NewPre() *Pre {
	return &Pre{
		style: make(map[string]string),
	}
}

func (p *Pre) AddStyle(k, v string) *Pre {
	p.style[k] = v
	return p
}

func (p *Pre) AddStyles(m map[string]string) *Pre {
	for k, v := range m {
		p.style[k] = v
	}
	return p
}

func (p *Pre) Style(m map[string]string) *Pre {
	p.style = m
	return p
}

func (p *Pre) Text(str string) *Pre {
	p.text = str
	return p
}

func (p *Pre) Bytes() []byte {
	return cloneBytes(p.buf.Bytes())
}

// IsBodyElement implements BodyElement interface
func (p *Pre) IsBodyElement() {}

func (p *Pre) Prepare() {
	p.buf.Reset()
	p.buf.WriteString("<pre")
	if len(p.style) != 0 {
		parseStyle(&p.buf, p.style)
	}
	p.buf.WriteString(">" + p.text + "</pre>")
}

type Blockquote struct {
	buf   bytes.Buffer
	style map[string]string
	text  string
	cite  string
}

func NewBlockquote() *Blockquote {
	return &Blockquote{
		style: make(map[string]string),
	}
}

func (b *Blockquote) AddStyle(k, v string) *Blockquote {
	b.style[k] = v
	return b
}

func (b *Blockquote) AddStyles(m map[string]string) *Blockquote {
	for k, v := range m {
		b.style[k] = v
	}
	return b
}

func (b *Blockquote) Style(m map[string]string) *Blockquote {
	b.style = m
	return b
}

func (b *Blockquote) Text(str string) *Blockquote {
	b.text = str
	return b
}

func (b *Blockquote) Cite(c string) *Blockquote {
	b.cite = c
	return b
}

func (b *Blockquote) Bytes() []byte {
	return cloneBytes(b.buf.Bytes())
}

// IsBodyElement implements BodyElement interface
func (b *Blockquote) IsBodyElement() {}

func (b *Blockquote) Prepare() {
	b.buf.Reset()
	b.buf.WriteString("<blockquote")
	if b.cite != "" {
		b.buf.WriteString(" cite=\"" + b.cite + "\"")
	}
	if len(b.style) != 0 {
		parseStyle(&b.buf, b.style)
	}
	b.buf.WriteString(">" + b.text + "</blockquote>")
}

type Menu struct {
	buf      bytes.Buffer
	style    map[string]string
	contents []Element
	menuType string
	label    string
}

func NewMenu() *Menu {
	return &Menu{
		style: make(map[string]string),
	}
}

func (m *Menu) AddStyle(k, v string) *Menu {
	m.style[k] = v
	return m
}

func (m *Menu) AddStyles(ms map[string]string) *Menu {
	for k, v := range ms {
		m.style[k] = v
	}
	return m
}

func (m *Menu) Style(ms map[string]string) *Menu {
	m.style = ms
	return m
}

func (m *Menu) Add(e Element) *Menu {
	m.contents = appendElement(m.contents, e)
	return m
}

func (m *Menu) Type(t string) *Menu {
	m.menuType = t
	return m
}

func (m *Menu) Label(l string) *Menu {
	m.label = l
	return m
}

func (m *Menu) Bytes() []byte {
	return cloneBytes(m.buf.Bytes())
}

// IsBodyElement implements BodyElement interface
func (m *Menu) IsBodyElement() {}

func (m *Menu) Prepare() {
	m.buf.Reset()
	m.buf.WriteString("<menu")
	if m.menuType != "" {
		m.buf.WriteString(" type=\"" + m.menuType + "\"")
	}
	if m.label != "" {
		m.buf.WriteString(" label=\"" + m.label + "\"")
	}
	if len(m.style) != 0 {
		parseStyle(&m.buf, m.style)
	}
	m.buf.WriteByte('>')

	writeElements(&m.buf, m.contents)
	m.buf.WriteString("</menu>")
}

type Ol struct {
	buf      bytes.Buffer
	style    map[string]string
	contents []Element
	start    int
	listType string
	reversed bool
}

func NewOl() *Ol {
	return &Ol{
		style: make(map[string]string),
	}
}

func (o *Ol) AddStyle(k, v string) *Ol {
	o.style[k] = v
	return o
}

func (o *Ol) AddStyles(m map[string]string) *Ol {
	for k, v := range m {
		o.style[k] = v
	}
	return o
}

func (o *Ol) Style(m map[string]string) *Ol {
	o.style = m
	return o
}

func (o *Ol) Add(e Element) *Ol {
	o.contents = appendElement(o.contents, e)
	return o
}

func (o *Ol) Start(s int) *Ol {
	o.start = s
	return o
}

func (o *Ol) Type(t string) *Ol {
	o.listType = t
	return o
}

func (o *Ol) Reversed(r bool) *Ol {
	o.reversed = r
	return o
}

func (o *Ol) Bytes() []byte {
	return cloneBytes(o.buf.Bytes())
}

// IsBodyElement implements BodyElement interface
func (o *Ol) IsBodyElement() {}

func (o *Ol) Prepare() {
	o.buf.Reset()
	o.buf.WriteString("<ol")
	if o.start > 0 {
		o.buf.WriteString(" start=\"")
		o.buf.WriteString(string(rune(o.start + '0')))
		o.buf.WriteString("\"")
	}
	if o.listType != "" {
		o.buf.WriteString(" type=\"" + o.listType + "\"")
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

type Ul struct {
	buf      bytes.Buffer
	style    map[string]string
	contents []Element
}

func NewUl() *Ul {
	return &Ul{
		style: make(map[string]string),
	}
}

func (u *Ul) AddStyle(k, v string) *Ul {
	u.style[k] = v
	return u
}

func (u *Ul) AddStyles(m map[string]string) *Ul {
	for k, v := range m {
		u.style[k] = v
	}
	return u
}

func (u *Ul) Style(m map[string]string) *Ul {
	u.style = m
	return u
}

func (u *Ul) Add(e Element) *Ul {
	u.contents = appendElement(u.contents, e)
	return u
}

func (u *Ul) Bytes() []byte {
	return cloneBytes(u.buf.Bytes())
}

// IsBodyElement implements BodyElement interface
func (u *Ul) IsBodyElement() {}

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

type Li struct {
	buf      bytes.Buffer
	style    map[string]string
	contents []Element
	value    int
}

func NewLi() *Li {
	return &Li{
		style: make(map[string]string),
	}
}

func (l *Li) AddStyle(k, v string) *Li {
	l.style[k] = v
	return l
}

func (l *Li) AddStyles(m map[string]string) *Li {
	for k, v := range m {
		l.style[k] = v
	}
	return l
}

func (l *Li) Style(m map[string]string) *Li {
	l.style = m
	return l
}

func (l *Li) Add(e Element) *Li {
	l.contents = appendElement(l.contents, e)
	return l
}

func (l *Li) Value(v int) *Li {
	l.value = v
	return l
}

func (l *Li) Bytes() []byte {
	return cloneBytes(l.buf.Bytes())
}

// IsBodyElement implements BodyElement interface
func (l *Li) IsBodyElement() {}

func (l *Li) Prepare() {
	l.buf.Reset()
	l.buf.WriteString("<li")
	if l.value > 0 {
		l.buf.WriteString(" value=\"")
		l.buf.WriteString(string(rune(l.value + '0')))
		l.buf.WriteString("\"")
	}
	if len(l.style) != 0 {
		parseStyle(&l.buf, l.style)
	}
	l.buf.WriteByte('>')

	writeElements(&l.buf, l.contents)
	l.buf.WriteString("</li>")
}

type Dl struct {
	buf      bytes.Buffer
	style    map[string]string
	contents []Element
}

func NewDl() *Dl {
	return &Dl{
		style: make(map[string]string),
	}
}

func (d *Dl) AddStyle(k, v string) *Dl {
	d.style[k] = v
	return d
}

func (d *Dl) AddStyles(m map[string]string) *Dl {
	for k, v := range m {
		d.style[k] = v
	}
	return d
}

func (d *Dl) Style(m map[string]string) *Dl {
	d.style = m
	return d
}

func (d *Dl) Add(e Element) *Dl {
	d.contents = appendElement(d.contents, e)
	return d
}

func (d *Dl) Bytes() []byte {
	return cloneBytes(d.buf.Bytes())
}

// IsBodyElement implements BodyElement interface
func (d *Dl) IsBodyElement() {}

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

type Dt struct {
	buf      bytes.Buffer
	style    map[string]string
	contents []Element
}

func NewDt() *Dt {
	return &Dt{
		style: make(map[string]string),
	}
}

func (d *Dt) AddStyle(k, v string) *Dt {
	d.style[k] = v
	return d
}

func (d *Dt) AddStyles(m map[string]string) *Dt {
	for k, v := range m {
		d.style[k] = v
	}
	return d
}

func (d *Dt) Style(m map[string]string) *Dt {
	d.style = m
	return d
}

func (d *Dt) Add(e Element) *Dt {
	d.contents = appendElement(d.contents, e)
	return d
}

func (d *Dt) Bytes() []byte {
	return cloneBytes(d.buf.Bytes())
}

// IsBodyElement implements BodyElement interface
func (d *Dt) IsBodyElement() {}

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

type Dd struct {
	buf      bytes.Buffer
	style    map[string]string
	contents []Element
}

func NewDd() *Dd {
	return &Dd{
		style: make(map[string]string),
	}
}

func (d *Dd) AddStyle(k, v string) *Dd {
	d.style[k] = v
	return d
}

func (d *Dd) AddStyles(m map[string]string) *Dd {
	for k, v := range m {
		d.style[k] = v
	}
	return d
}

func (d *Dd) Style(m map[string]string) *Dd {
	d.style = m
	return d
}

func (d *Dd) Add(e Element) *Dd {
	d.contents = appendElement(d.contents, e)
	return d
}

func (d *Dd) Bytes() []byte {
	return cloneBytes(d.buf.Bytes())
}

// IsBodyElement implements BodyElement interface
func (d *Dd) IsBodyElement() {}

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

type Figure struct {
	buf      bytes.Buffer
	style    map[string]string
	contents []Element
}

func NewFigure() *Figure {
	return &Figure{
		style: make(map[string]string),
	}
}

func (f *Figure) AddStyle(k, v string) *Figure {
	f.style[k] = v
	return f
}

func (f *Figure) AddStyles(m map[string]string) *Figure {
	for k, v := range m {
		f.style[k] = v
	}
	return f
}

func (f *Figure) Style(m map[string]string) *Figure {
	f.style = m
	return f
}

func (f *Figure) Add(e Element) *Figure {
	f.contents = appendElement(f.contents, e)
	return f
}

func (f *Figure) Bytes() []byte {
	return cloneBytes(f.buf.Bytes())
}

// IsBodyElement implements BodyElement interface
func (f *Figure) IsBodyElement() {}

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

type Figcaption struct {
	buf      bytes.Buffer
	style    map[string]string
	contents []Element
}

func NewFigcaption() *Figcaption {
	return &Figcaption{
		style: make(map[string]string),
	}
}

func (f *Figcaption) AddStyle(k, v string) *Figcaption {
	f.style[k] = v
	return f
}

func (f *Figcaption) AddStyles(m map[string]string) *Figcaption {
	for k, v := range m {
		f.style[k] = v
	}
	return f
}

func (f *Figcaption) Style(m map[string]string) *Figcaption {
	f.style = m
	return f
}

func (f *Figcaption) Add(e Element) *Figcaption {
	f.contents = appendElement(f.contents, e)
	return f
}

func (f *Figcaption) Bytes() []byte {
	return cloneBytes(f.buf.Bytes())
}

// IsBodyElement implements BodyElement interface
func (f *Figcaption) IsBodyElement() {}

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

type Search struct {
	buf      bytes.Buffer
	style    map[string]string
	contents []Element
}

func NewSearch() *Search {
	return &Search{
		style: make(map[string]string),
	}
}

func (s *Search) AddStyle(k, v string) *Search {
	s.style[k] = v
	return s
}

func (s *Search) AddStyles(m map[string]string) *Search {
	for k, v := range m {
		s.style[k] = v
	}
	return s
}

func (s *Search) Style(m map[string]string) *Search {
	s.style = m
	return s
}

func (s *Search) Add(e Element) *Search {
	s.contents = appendElement(s.contents, e)
	return s
}

func (s *Search) Bytes() []byte {
	return cloneBytes(s.buf.Bytes())
}

// IsBodyElement implements BodyElement interface
func (s *Search) IsBodyElement() {}

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
