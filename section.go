package rephtml

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

// prepare renders the Header component into its internal buffer.
func (h *Header) prepare() {
	tg := openTag(&h.buf, "header")
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

// prepare renders the Nav component into its internal buffer.
func (n *Nav) prepare() {
	tg := openTag(&n.buf, "nav")
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

// prepare renders the Main component into its internal buffer.
func (m *Main) prepare() {
	tg := openTag(&m.buf, "main")
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

// prepare renders the Section component into its internal buffer.
func (s *Section) prepare() {
	tg := openTag(&s.buf, "section")
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

// prepare renders the Article component into its internal buffer.
func (a *Article) prepare() {
	tg := openTag(&a.buf, "article")
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

// prepare renders the Aside component into its internal buffer.
func (as *Aside) prepare() {
	tg := openTag(&as.buf, "aside")
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

// prepare renders the Footer component into its internal buffer.
func (f *Footer) prepare() {
	tg := openTag(&f.buf, "footer")
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

// prepare renders the Address component into its internal buffer.
func (ad *Address) prepare() {
	tg := openTag(&ad.buf, "address")
	tg.styleAttr(ad.style)
	tg.children(ad.contents)
}
