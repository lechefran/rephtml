package rephtml

import "bytes"

// Header represents the Header component or supporting type.
type Header struct {
	buf      bytes.Buffer
	style    map[string]string
	contents []Element
}

// NewHeader creates a new Header component.
func NewHeader() *Header {
	return &Header{
		style: make(map[string]string),
	}
}

// AddStyle adds one inline CSS declaration to the Header component.
func (h *Header) AddStyle(k, v string) *Header {
	h.style[k] = v
	return h
}

// AddStyles adds multiple inline CSS declarations to the Header component.
func (h *Header) AddStyles(m map[string]string) *Header {
	for k, v := range m {
		h.style[k] = v
	}
	return h
}

// Style replaces the inline CSS declarations on the Header component.
func (h *Header) Style(m map[string]string) *Header {
	h.style = cloneStyleMap(m)
	return h
}

// Add appends child content to the Header component.
func (h *Header) Add(e Element) *Header {
	h.contents = appendElement(h.contents, e)
	return h
}

// Bytes returns a defensive copy of the rendered Header bytes.
func (h *Header) Bytes() []byte {
	return cloneBytes(h.buf.Bytes())
}

// IsBodyElement implements BodyElement interface
func (h *Header) IsBodyElement() {}

// Prepare renders the Header component into its internal buffer.
func (h *Header) Prepare() {
	h.buf.Reset()
	h.buf.WriteString("<header")
	if len(h.style) != 0 {
		parseStyle(&h.buf, h.style)
	}
	h.buf.WriteByte('>')

	writeElements(&h.buf, h.contents)
	h.buf.WriteString("</header>")
}

// Nav represents the Nav component or supporting type.
type Nav struct {
	buf      bytes.Buffer
	style    map[string]string
	contents []Element
	role     string
}

// NewNav creates a new Nav component.
func NewNav() *Nav {
	return &Nav{
		style: make(map[string]string),
	}
}

// AddStyle adds one inline CSS declaration to the Nav component.
func (n *Nav) AddStyle(k, v string) *Nav {
	n.style[k] = v
	return n
}

// AddStyles adds multiple inline CSS declarations to the Nav component.
func (n *Nav) AddStyles(m map[string]string) *Nav {
	for k, v := range m {
		n.style[k] = v
	}
	return n
}

// Style replaces the inline CSS declarations on the Nav component.
func (n *Nav) Style(m map[string]string) *Nav {
	n.style = cloneStyleMap(m)
	return n
}

// Add appends child content to the Nav component.
func (n *Nav) Add(e Element) *Nav {
	n.contents = appendElement(n.contents, e)
	return n
}

// Role sets the role value on the Nav component.
func (n *Nav) Role(r string) *Nav {
	n.role = r
	return n
}

// Bytes returns a defensive copy of the rendered Nav bytes.
func (n *Nav) Bytes() []byte {
	return cloneBytes(n.buf.Bytes())
}

// IsBodyElement implements BodyElement interface
func (n *Nav) IsBodyElement() {}

// Prepare renders the Nav component into its internal buffer.
func (n *Nav) Prepare() {
	n.buf.Reset()
	n.buf.WriteString("<nav")
	if n.role != "" {
		writeAttr(&n.buf, "role", n.role)
	}
	if len(n.style) != 0 {
		parseStyle(&n.buf, n.style)
	}
	n.buf.WriteByte('>')

	writeElements(&n.buf, n.contents)
	n.buf.WriteString("</nav>")
}

// Main represents the Main component or supporting type.
type Main struct {
	buf      bytes.Buffer
	style    map[string]string
	contents []Element
}

// NewMain creates a new Main component.
func NewMain() *Main {
	return &Main{
		style: make(map[string]string),
	}
}

// AddStyle adds one inline CSS declaration to the Main component.
func (m *Main) AddStyle(k, v string) *Main {
	m.style[k] = v
	return m
}

// AddStyles adds multiple inline CSS declarations to the Main component.
func (m *Main) AddStyles(mp map[string]string) *Main {
	for k, v := range mp {
		m.style[k] = v
	}
	return m
}

// Style replaces the inline CSS declarations on the Main component.
func (m *Main) Style(mp map[string]string) *Main {
	m.style = cloneStyleMap(mp)
	return m
}

// Add appends child content to the Main component.
func (m *Main) Add(e Element) *Main {
	m.contents = appendElement(m.contents, e)
	return m
}

// Bytes returns a defensive copy of the rendered Main bytes.
func (m *Main) Bytes() []byte {
	return cloneBytes(m.buf.Bytes())
}

// IsBodyElement implements BodyElement interface
func (m *Main) IsBodyElement() {}

// Prepare renders the Main component into its internal buffer.
func (m *Main) Prepare() {
	m.buf.Reset()
	m.buf.WriteString("<main")
	if len(m.style) != 0 {
		parseStyle(&m.buf, m.style)
	}
	m.buf.WriteByte('>')

	writeElements(&m.buf, m.contents)
	m.buf.WriteString("</main>")
}

// Section represents the Section component or supporting type.
type Section struct {
	buf       bytes.Buffer
	style     map[string]string
	contents  []Element
	ariaLabel string
}

// NewSection creates a new Section component.
func NewSection() *Section {
	return &Section{
		style: make(map[string]string),
	}
}

// AddStyle adds one inline CSS declaration to the Section component.
func (s *Section) AddStyle(k, v string) *Section {
	s.style[k] = v
	return s
}

// AddStyles adds multiple inline CSS declarations to the Section component.
func (s *Section) AddStyles(m map[string]string) *Section {
	for k, v := range m {
		s.style[k] = v
	}
	return s
}

// Style replaces the inline CSS declarations on the Section component.
func (s *Section) Style(m map[string]string) *Section {
	s.style = cloneStyleMap(m)
	return s
}

// Add appends child content to the Section component.
func (s *Section) Add(e Element) *Section {
	s.contents = appendElement(s.contents, e)
	return s
}

// AriaLabel sets the arialabel value on the Section component.
func (s *Section) AriaLabel(label string) *Section {
	s.ariaLabel = label
	return s
}

// Bytes returns a defensive copy of the rendered Section bytes.
func (s *Section) Bytes() []byte {
	return cloneBytes(s.buf.Bytes())
}

// IsBodyElement implements BodyElement interface
func (s *Section) IsBodyElement() {}

// Prepare renders the Section component into its internal buffer.
func (s *Section) Prepare() {
	s.buf.Reset()
	s.buf.WriteString("<section")
	if s.ariaLabel != "" {
		writeAttr(&s.buf, "aria-label", s.ariaLabel)
	}
	if len(s.style) != 0 {
		parseStyle(&s.buf, s.style)
	}
	s.buf.WriteByte('>')

	writeElements(&s.buf, s.contents)
	s.buf.WriteString("</section>")
}

// Article represents the Article component or supporting type.
type Article struct {
	buf      bytes.Buffer
	style    map[string]string
	contents []Element
}

// NewArticle creates a new Article component.
func NewArticle() *Article {
	return &Article{
		style: make(map[string]string),
	}
}

// AddStyle adds one inline CSS declaration to the Article component.
func (a *Article) AddStyle(k, v string) *Article {
	a.style[k] = v
	return a
}

// AddStyles adds multiple inline CSS declarations to the Article component.
func (a *Article) AddStyles(mp map[string]string) *Article {
	for k, v := range mp {
		a.style[k] = v
	}
	return a
}

// Style replaces the inline CSS declarations on the Article component.
func (a *Article) Style(mp map[string]string) *Article {
	a.style = cloneStyleMap(mp)
	return a
}

// Add appends child content to the Article component.
func (a *Article) Add(e Element) *Article {
	a.contents = appendElement(a.contents, e)
	return a
}

// Bytes returns a defensive copy of the rendered Article bytes.
func (a *Article) Bytes() []byte {
	return cloneBytes(a.buf.Bytes())
}

// IsBodyElement implements BodyElement interface
func (a *Article) IsBodyElement() {}

// Prepare renders the Article component into its internal buffer.
func (a *Article) Prepare() {
	a.buf.Reset()
	a.buf.WriteString("<article")
	if len(a.style) != 0 {
		parseStyle(&a.buf, a.style)
	}
	a.buf.WriteByte('>')

	writeElements(&a.buf, a.contents)
	a.buf.WriteString("</article>")
}

// Aside represents the Aside component or supporting type.
type Aside struct {
	buf      bytes.Buffer
	style    map[string]string
	contents []Element
}

// NewAside creates a new Aside component.
func NewAside() *Aside {
	return &Aside{
		style: make(map[string]string),
	}
}

// AddStyle adds one inline CSS declaration to the Aside component.
func (as *Aside) AddStyle(k, v string) *Aside {
	as.style[k] = v
	return as
}

// AddStyles adds multiple inline CSS declarations to the Aside component.
func (as *Aside) AddStyles(mp map[string]string) *Aside {
	for k, v := range mp {
		as.style[k] = v
	}
	return as
}

// Style replaces the inline CSS declarations on the Aside component.
func (as *Aside) Style(mp map[string]string) *Aside {
	as.style = cloneStyleMap(mp)
	return as
}

// Add appends child content to the Aside component.
func (as *Aside) Add(e Element) *Aside {
	as.contents = appendElement(as.contents, e)
	return as
}

// Bytes returns a defensive copy of the rendered Aside bytes.
func (as *Aside) Bytes() []byte {
	return cloneBytes(as.buf.Bytes())
}

// IsBodyElement implements BodyElement interface
func (as *Aside) IsBodyElement() {}

// Prepare renders the Aside component into its internal buffer.
func (as *Aside) Prepare() {
	as.buf.Reset()
	as.buf.WriteString("<aside")
	if len(as.style) != 0 {
		parseStyle(&as.buf, as.style)
	}
	as.buf.WriteByte('>')

	writeElements(&as.buf, as.contents)
	as.buf.WriteString("</aside>")
}

// Footer represents the Footer component or supporting type.
type Footer struct {
	buf      bytes.Buffer
	style    map[string]string
	contents []Element
}

// NewFooter creates a new Footer component.
func NewFooter() *Footer {
	return &Footer{
		style: make(map[string]string),
	}
}

// AddStyle adds one inline CSS declaration to the Footer component.
func (f *Footer) AddStyle(k, v string) *Footer {
	f.style[k] = v
	return f
}

// AddStyles adds multiple inline CSS declarations to the Footer component.
func (f *Footer) AddStyles(mp map[string]string) *Footer {
	for k, v := range mp {
		f.style[k] = v
	}
	return f
}

// Style replaces the inline CSS declarations on the Footer component.
func (f *Footer) Style(mp map[string]string) *Footer {
	f.style = cloneStyleMap(mp)
	return f
}

// Add appends child content to the Footer component.
func (f *Footer) Add(e Element) *Footer {
	f.contents = appendElement(f.contents, e)
	return f
}

// Bytes returns a defensive copy of the rendered Footer bytes.
func (f *Footer) Bytes() []byte {
	return cloneBytes(f.buf.Bytes())
}

// IsBodyElement implements BodyElement interface
func (f *Footer) IsBodyElement() {}

// Prepare renders the Footer component into its internal buffer.
func (f *Footer) Prepare() {
	f.buf.Reset()
	f.buf.WriteString("<footer")
	if len(f.style) != 0 {
		parseStyle(&f.buf, f.style)
	}
	f.buf.WriteByte('>')

	writeElements(&f.buf, f.contents)
	f.buf.WriteString("</footer>")
}

// Address represents the Address component or supporting type.
type Address struct {
	buf      bytes.Buffer
	style    map[string]string
	contents []Element
}

// NewAddress creates a new Address component.
func NewAddress() *Address {
	return &Address{
		style: make(map[string]string),
	}
}

// AddStyle adds one inline CSS declaration to the Address component.
func (ad *Address) AddStyle(k, v string) *Address {
	ad.style[k] = v
	return ad
}

// AddStyles adds multiple inline CSS declarations to the Address component.
func (ad *Address) AddStyles(mp map[string]string) *Address {
	for k, v := range mp {
		ad.style[k] = v
	}
	return ad
}

// Style replaces the inline CSS declarations on the Address component.
func (ad *Address) Style(mp map[string]string) *Address {
	ad.style = cloneStyleMap(mp)
	return ad
}

// Add appends child content to the Address component.
func (ad *Address) Add(e Element) *Address {
	ad.contents = appendElement(ad.contents, e)
	return ad
}

// Bytes returns a defensive copy of the rendered Address bytes.
func (ad *Address) Bytes() []byte {
	return cloneBytes(ad.buf.Bytes())
}

// IsBodyElement implements BodyElement interface
func (ad *Address) IsBodyElement() {}

// Prepare renders the Address component into its internal buffer.
func (ad *Address) Prepare() {
	ad.buf.Reset()
	ad.buf.WriteString("<address")
	if len(ad.style) != 0 {
		parseStyle(&ad.buf, ad.style)
	}
	ad.buf.WriteByte('>')

	writeElements(&ad.buf, ad.contents)
	ad.buf.WriteString("</address>")
}
