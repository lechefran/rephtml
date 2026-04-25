package rephtml

import "bytes"

type Header struct {
	buf      bytes.Buffer
	style    map[string]string
	contents []Element
}

func NewHeader() *Header {
	return &Header{
		style: make(map[string]string),
	}
}

func (h *Header) AddStyle(k, v string) *Header {
	h.style[k] = v
	return h
}

func (h *Header) AddStyles(m map[string]string) *Header {
	for k, v := range m {
		h.style[k] = v
	}
	return h
}

func (h *Header) Style(m map[string]string) *Header {
	h.style = m
	return h
}

func (h *Header) Add(e Element) *Header {
	h.contents = appendElement(h.contents, e)
	return h
}

func (h *Header) Bytes() []byte {
	return cloneBytes(h.buf.Bytes())
}

// IsBodyElement implements BodyElement interface
func (h *Header) IsBodyElement() {}

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

type Nav struct {
	buf      bytes.Buffer
	style    map[string]string
	contents []Element
	role     string
}

func NewNav() *Nav {
	return &Nav{
		style: make(map[string]string),
	}
}

func (n *Nav) AddStyle(k, v string) *Nav {
	n.style[k] = v
	return n
}

func (n *Nav) AddStyles(m map[string]string) *Nav {
	for k, v := range m {
		n.style[k] = v
	}
	return n
}

func (n *Nav) Style(m map[string]string) *Nav {
	n.style = m
	return n
}

func (n *Nav) Add(e Element) *Nav {
	n.contents = appendElement(n.contents, e)
	return n
}

func (n *Nav) Role(r string) *Nav {
	n.role = r
	return n
}

func (n *Nav) Bytes() []byte {
	return cloneBytes(n.buf.Bytes())
}

// IsBodyElement implements BodyElement interface
func (n *Nav) IsBodyElement() {}

func (n *Nav) Prepare() {
	n.buf.Reset()
	n.buf.WriteString("<nav")
	if n.role != "" {
		n.buf.WriteString(" role=\"" + n.role + "\"")
	}
	if len(n.style) != 0 {
		parseStyle(&n.buf, n.style)
	}
	n.buf.WriteByte('>')

	writeElements(&n.buf, n.contents)
	n.buf.WriteString("</nav>")
}

type Main struct {
	buf      bytes.Buffer
	style    map[string]string
	contents []Element
}

func NewMain() *Main {
	return &Main{
		style: make(map[string]string),
	}
}

func (m *Main) AddStyle(k, v string) *Main {
	m.style[k] = v
	return m
}

func (m *Main) AddStyles(mp map[string]string) *Main {
	for k, v := range mp {
		m.style[k] = v
	}
	return m
}

func (m *Main) Style(mp map[string]string) *Main {
	m.style = mp
	return m
}

func (m *Main) Add(e Element) *Main {
	m.contents = appendElement(m.contents, e)
	return m
}

func (m *Main) Bytes() []byte {
	return cloneBytes(m.buf.Bytes())
}

// IsBodyElement implements BodyElement interface
func (m *Main) IsBodyElement() {}

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

type Section struct {
	buf       bytes.Buffer
	style     map[string]string
	contents  []Element
	ariaLabel string
}

func NewSection() *Section {
	return &Section{
		style: make(map[string]string),
	}
}

func (s *Section) AddStyle(k, v string) *Section {
	s.style[k] = v
	return s
}

func (s *Section) AddStyles(m map[string]string) *Section {
	for k, v := range m {
		s.style[k] = v
	}
	return s
}

func (s *Section) Style(m map[string]string) *Section {
	s.style = m
	return s
}

func (s *Section) Add(e Element) *Section {
	s.contents = appendElement(s.contents, e)
	return s
}

func (s *Section) AriaLabel(label string) *Section {
	s.ariaLabel = label
	return s
}

func (s *Section) Bytes() []byte {
	return cloneBytes(s.buf.Bytes())
}

// IsBodyElement implements BodyElement interface
func (s *Section) IsBodyElement() {}

func (s *Section) Prepare() {
	s.buf.Reset()
	s.buf.WriteString("<section")
	if s.ariaLabel != "" {
		s.buf.WriteString(" aria-label=\"" + s.ariaLabel + "\"")
	}
	if len(s.style) != 0 {
		parseStyle(&s.buf, s.style)
	}
	s.buf.WriteByte('>')

	writeElements(&s.buf, s.contents)
	s.buf.WriteString("</section>")
}

type Article struct {
	buf      bytes.Buffer
	style    map[string]string
	contents []Element
}

func NewArticle() *Article {
	return &Article{
		style: make(map[string]string),
	}
}

func (a *Article) AddStyle(k, v string) *Article {
	a.style[k] = v
	return a
}

func (a *Article) AddStyles(mp map[string]string) *Article {
	for k, v := range mp {
		a.style[k] = v
	}
	return a
}

func (a *Article) Style(mp map[string]string) *Article {
	a.style = mp
	return a
}

func (a *Article) Add(e Element) *Article {
	a.contents = appendElement(a.contents, e)
	return a
}

func (a *Article) Bytes() []byte {
	return cloneBytes(a.buf.Bytes())
}

// IsBodyElement implements BodyElement interface
func (a *Article) IsBodyElement() {}

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

type Aside struct {
	buf      bytes.Buffer
	style    map[string]string
	contents []Element
}

func NewAside() *Aside {
	return &Aside{
		style: make(map[string]string),
	}
}

func (as *Aside) AddStyle(k, v string) *Aside {
	as.style[k] = v
	return as
}

func (as *Aside) AddStyles(mp map[string]string) *Aside {
	for k, v := range mp {
		as.style[k] = v
	}
	return as
}

func (as *Aside) Style(mp map[string]string) *Aside {
	as.style = mp
	return as
}

func (as *Aside) Add(e Element) *Aside {
	as.contents = appendElement(as.contents, e)
	return as
}

func (as *Aside) Bytes() []byte {
	return cloneBytes(as.buf.Bytes())
}

// IsBodyElement implements BodyElement interface
func (as *Aside) IsBodyElement() {}

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

type Footer struct {
	buf      bytes.Buffer
	style    map[string]string
	contents []Element
}

func NewFooter() *Footer {
	return &Footer{
		style: make(map[string]string),
	}
}

func (f *Footer) AddStyle(k, v string) *Footer {
	f.style[k] = v
	return f
}

func (f *Footer) AddStyles(mp map[string]string) *Footer {
	for k, v := range mp {
		f.style[k] = v
	}
	return f
}

func (f *Footer) Style(mp map[string]string) *Footer {
	f.style = mp
	return f
}

func (f *Footer) Add(e Element) *Footer {
	f.contents = appendElement(f.contents, e)
	return f
}

func (f *Footer) Bytes() []byte {
	return cloneBytes(f.buf.Bytes())
}

// IsBodyElement implements BodyElement interface
func (f *Footer) IsBodyElement() {}

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

type Address struct {
	buf      bytes.Buffer
	style    map[string]string
	contents []Element
}

func NewAddress() *Address {
	return &Address{
		style: make(map[string]string),
	}
}

func (ad *Address) AddStyle(k, v string) *Address {
	ad.style[k] = v
	return ad
}

func (ad *Address) AddStyles(mp map[string]string) *Address {
	for k, v := range mp {
		ad.style[k] = v
	}
	return ad
}

func (ad *Address) Style(mp map[string]string) *Address {
	ad.style = mp
	return ad
}

func (ad *Address) Add(e Element) *Address {
	ad.contents = appendElement(ad.contents, e)
	return ad
}

func (ad *Address) Bytes() []byte {
	return cloneBytes(ad.buf.Bytes())
}

// IsBodyElement implements BodyElement interface
func (ad *Address) IsBodyElement() {}

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
