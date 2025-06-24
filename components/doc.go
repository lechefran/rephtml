package rephtml

import (
	"bytes"
	"log"
	"os"
)

const tab = "\t"

type HtmlFile struct {
	buf               bytes.Buffer
	head, body, style []byte
	lang              string
	title             string
	ttrack            int // tab tracker
	base              Base
	options           Options
}

func NewHtmlFile() *HtmlFile {
	return &HtmlFile{
		ttrack: 1,
	}
}

func (h *HtmlFile) Lang(lang string) *HtmlFile {
	h.lang = lang
	return h
}

func (h *HtmlFile) Title(title string) *HtmlFile {
	h.title = title
	return h
}

func (h *HtmlFile) AddOptions(opts Options) *HtmlFile {
	h.options = opts
	return h
}

// Element struct functions

func (h *HtmlFile) AddToHead(e Element) *HtmlFile {
	h.head = append(h.head, e.Bytes()...)
	return h
}

func (h *HtmlFile) AddToBody(e Element) *HtmlFile {
	h.body = append(h.body, e.Bytes()...)
	return h
}

func (h *HtmlFile) Bytes() []byte {
	return h.buf.Bytes()
}

func (h *HtmlFile) Prepare() *HtmlFile {
	return h
}

func (h *HtmlFile) WriteToFile(path string) {
	if len(h.head) == 0 && len(h.body) == 0 && len(h.style) == 0 {
		log.Print("No values were appended to the HTML File. " + path + " will not be created")
	}
	if _, err := os.Stat(path); os.IsExist(err) {
		if err := os.Remove(path); err != nil {
			log.Fatal(err)
		}
		file, err := os.Create(path)
		if err != nil {
			log.Fatal(err)
		}
		defer func(file *os.File) {
			if err := file.Close(); err != nil {
				log.Fatal(err)
			}
		}(file)
		if _, err := file.Write(h.buf.Bytes()); err != nil {
			log.Fatal(err)
		}
	} else {
		file, err := os.Create(path)
		if err != nil {
			log.Fatal(err)
		}
		defer func(file *os.File) {
			if err := file.Close(); err != nil {
				log.Fatal(err)
			}
		}(file)
		if _, err := file.Write(h.buf.Bytes()); err != nil {
			log.Fatal(err)
		}
	}
}

// Head represents the HTML head element for document metadata
type Head struct {
	buf      bytes.Buffer
	style    map[string]string
	contents [][]byte
	ttrack   int
}

// NewHead creates a new Head element
func NewHead() *Head {
	return &Head{
		style: make(map[string]string),
	}
}

// Bytes returns the buffer contents
func (h *Head) Bytes() []byte {
	return h.buf.Bytes()
}

// Prepare builds the HTML for the head element
func (h *Head) Prepare() {
	h.buf.WriteString("<head")

	if len(h.style) != 0 {
		idx := 0
		h.buf.WriteString(" style=\"")
		for k, v := range h.style {
			h.buf.WriteString(k + ": " + v + ";")
			if idx != len(h.style)-1 {
				h.buf.WriteByte(' ')
			}
			idx++
		}
		h.buf.WriteString("\"")
	}

	h.buf.WriteByte('>')

	for _, content := range h.contents {
		h.buf.Write(content)
	}

	h.buf.WriteString("</head>")
}

// Add adds content to the head element
func (h *Head) Add(e Element) *Head {
	if e != nil {
		e.Prepare()
		h.contents = append(h.contents, e.Bytes())
	}
	return h
}

// AddStyle adds a single CSS property
func (h *Head) AddStyle(k, v string) *Head {
	h.style[k] = v
	return h
}

// AddStyles adds multiple CSS properties
func (h *Head) AddStyles(m map[string]string) *Head {
	for k, v := range m {
		h.style[k] = v
	}
	return h
}

// Style replaces all styles
func (h *Head) Style(m map[string]string) *Head {
	h.style = make(map[string]string)
	for k, v := range m {
		h.style[k] = v
	}
	return h
}

// Body represents the HTML body element for document content
type Body struct {
	buf      bytes.Buffer
	style    map[string]string
	contents [][]byte
	ttrack   int
	onLoad   string
	onUnload string
}

// NewBody creates a new Body element
func NewBody() *Body {
	return &Body{
		style: make(map[string]string),
	}
}

// Bytes returns the buffer contents
func (b *Body) Bytes() []byte {
	return b.buf.Bytes()
}

// Prepare builds the HTML for the body element
func (b *Body) Prepare() {
	b.buf.WriteString("<body")

	if b.onLoad != "" {
		b.buf.WriteString(" onload=\"" + b.onLoad + "\"")
	}

	if b.onUnload != "" {
		b.buf.WriteString(" onunload=\"" + b.onUnload + "\"")
	}

	if len(b.style) != 0 {
		idx := 0
		b.buf.WriteString(" style=\"")
		for k, v := range b.style {
			b.buf.WriteString(k + ": " + v + ";")
			if idx != len(b.style)-1 {
				b.buf.WriteByte(' ')
			}
			idx++
		}
		b.buf.WriteString("\"")
	}

	b.buf.WriteByte('>')

	for _, content := range b.contents {
		b.buf.Write(content)
	}

	b.buf.WriteString("</body>")
}

// Add adds content to the body element
func (b *Body) Add(e Element) *Body {
	if e != nil {
		e.Prepare()
		b.contents = append(b.contents, e.Bytes())
	}
	return b
}

// OnLoad sets the onload attribute
func (b *Body) OnLoad(onLoad string) *Body {
	b.onLoad = onLoad
	return b
}

// OnUnload sets the onunload attribute
func (b *Body) OnUnload(onUnload string) *Body {
	b.onUnload = onUnload
	return b
}

// AddStyle adds a single CSS property
func (b *Body) AddStyle(k, v string) *Body {
	b.style[k] = v
	return b
}

// AddStyles adds multiple CSS properties
func (b *Body) AddStyles(m map[string]string) *Body {
	for k, v := range m {
		b.style[k] = v
	}
	return b
}

// Style replaces all styles
func (b *Body) Style(m map[string]string) *Body {
	b.style = make(map[string]string)
	for k, v := range m {
		b.style[k] = v
	}
	return b
}

// Title represents the HTML title element for document title
type Title struct {
	buf      bytes.Buffer
	style    map[string]string
	contents [][]byte
	ttrack   int
}

// NewTitle creates a new Title element
func NewTitle() *Title {
	return &Title{
		style: make(map[string]string),
	}
}

// Bytes returns the buffer contents
func (t *Title) Bytes() []byte {
	return t.buf.Bytes()
}

// Prepare builds the HTML for the title element
func (t *Title) Prepare() {
	t.buf.WriteString("<title")

	if len(t.style) != 0 {
		idx := 0
		t.buf.WriteString(" style=\"")
		for k, v := range t.style {
			t.buf.WriteString(k + ": " + v + ";")
			if idx != len(t.style)-1 {
				t.buf.WriteByte(' ')
			}
			idx++
		}
		t.buf.WriteString("\"")
	}

	t.buf.WriteByte('>')

	for _, content := range t.contents {
		t.buf.Write(content)
	}

	t.buf.WriteString("</title>")
}

// Add adds content to the title element
func (t *Title) Add(e Element) *Title {
	if e != nil {
		e.Prepare()
		t.contents = append(t.contents, e.Bytes())
	}
	return t
}

// Text adds text content to the title element
func (t *Title) Text(text string) *Title {
	t.contents = append(t.contents, []byte(text))
	return t
}

// AddStyle adds a single CSS property
func (t *Title) AddStyle(k, v string) *Title {
	t.style[k] = v
	return t
}

// AddStyles adds multiple CSS properties
func (t *Title) AddStyles(m map[string]string) *Title {
	for k, v := range m {
		t.style[k] = v
	}
	return t
}

// Style replaces all styles
func (t *Title) Style(m map[string]string) *Title {
	t.style = make(map[string]string)
	for k, v := range m {
		t.style[k] = v
	}
	return t
}

// Base represents the HTML base element for document base URL
type Base struct {
	buf    bytes.Buffer
	style  map[string]string
	href   string
	target string
}

// NewBase creates a new Base element
func NewBase() *Base {
	return &Base{
		style: make(map[string]string),
	}
}

// Bytes returns the buffer contents
func (b *Base) Bytes() []byte {
	return b.buf.Bytes()
}

// Prepare builds the HTML for the base element
func (b *Base) Prepare() {
	b.buf.WriteString("<base")

	if b.href != "" {
		b.buf.WriteString(" href=\"" + b.href + "\"")
	}

	if b.target != "" {
		b.buf.WriteString(" target=\"" + b.target + "\"")
	}

	if len(b.style) != 0 {
		idx := 0
		b.buf.WriteString(" style=\"")
		for k, v := range b.style {
			b.buf.WriteString(k + ": " + v + ";")
			if idx != len(b.style)-1 {
				b.buf.WriteByte(' ')
			}
			idx++
		}
		b.buf.WriteString("\"")
	}

	b.buf.WriteString(">")
}

// Href sets the href attribute
func (b *Base) Href(href string) *Base {
	b.href = href
	return b
}

// Target sets the target attribute
func (b *Base) Target(target string) *Base {
	b.target = target
	return b
}

// AddStyle adds a single CSS property
func (b *Base) AddStyle(k, v string) *Base {
	b.style[k] = v
	return b
}

// AddStyles adds multiple CSS properties
func (b *Base) AddStyles(m map[string]string) *Base {
	for k, v := range m {
		b.style[k] = v
	}
	return b
}

// Style replaces all styles
func (b *Base) Style(m map[string]string) *Base {
	b.style = make(map[string]string)
	for k, v := range m {
		b.style[k] = v
	}
	return b
}

// Link represents the HTML link element for external resources
type Link struct {
	buf            bytes.Buffer
	style          map[string]string
	rel            string
	href           string
	linkType       string
	media          string
	sizes          string
	crossOrigin    string
	integrity      string
	referrerPolicy string
	hreflang       string
}

// NewLink creates a new Link element
func NewLink() *Link {
	return &Link{
		style: make(map[string]string),
	}
}

// Bytes returns the buffer contents
func (l *Link) Bytes() []byte {
	return l.buf.Bytes()
}

// Prepare builds the HTML for the link element
func (l *Link) Prepare() {
	l.buf.WriteString("<link")
	if l.rel != "" {
		l.buf.WriteString(" rel=\"" + l.rel + "\"")
	}
	if l.href != "" {
		l.buf.WriteString(" href=\"" + l.href + "\"")
	}
	if l.linkType != "" {
		l.buf.WriteString(" type=\"" + l.linkType + "\"")
	}
	if l.media != "" {
		l.buf.WriteString(" media=\"" + l.media + "\"")
	}
	if l.sizes != "" {
		l.buf.WriteString(" sizes=\"" + l.sizes + "\"")
	}
	if l.crossOrigin != "" {
		l.buf.WriteString(" crossOrigin=\"" + l.crossOrigin + "\"")
	}
	if l.integrity != "" {
		l.buf.WriteString(" integrity=\"" + l.integrity + "\"")
	}
	if l.referrerPolicy != "" {
		l.buf.WriteString(" referrerPolicy=\"" + l.referrerPolicy + "\"")
	}
	if l.hreflang != "" {
		l.buf.WriteString(" hreflang=\"" + l.hreflang + "\"")
	}

	if len(l.style) != 0 {
		idx := 0
		l.buf.WriteString(" style=\"")
		for k, v := range l.style {
			l.buf.WriteString(k + ": " + v + ";")
			if idx != len(l.style)-1 {
				l.buf.WriteByte(' ')
			}
			idx++
		}
		l.buf.WriteString("\"")
	}
	l.buf.WriteString(">")
}

// Rel sets the rel attribute
func (l *Link) Rel(rel string) *Link {
	l.rel = rel
	return l
}

// Href sets the href attribute
func (l *Link) Href(href string) *Link {
	l.href = href
	return l
}

// Type sets the type attribute
func (l *Link) Type(linkType string) *Link {
	l.linkType = linkType
	return l
}

// Media sets the media attribute
func (l *Link) Media(media string) *Link {
	l.media = media
	return l
}

// Sizes sets the sizes attribute
func (l *Link) Sizes(sizes string) *Link {
	l.sizes = sizes
	return l
}

// CrossOrigin Cross Origin sets the crossOrigin attribute
func (l *Link) CrossOrigin(crossOrigin string) *Link {
	l.crossOrigin = crossOrigin
	return l
}

// Integrity sets the integrity attribute
func (l *Link) Integrity(integrity string) *Link {
	l.integrity = integrity
	return l
}

// ReferrerPolicy Referrer Policy sets the referrerPolicy attribute
func (l *Link) ReferrerPolicy(referrerPolicy string) *Link {
	l.referrerPolicy = referrerPolicy
	return l
}

// Hreflang sets the hreflang attribute
func (l *Link) Hreflang(hreflang string) *Link {
	l.hreflang = hreflang
	return l
}

// AddStyle adds a single CSS property
func (l *Link) AddStyle(k, v string) *Link {
	l.style[k] = v
	return l
}

// AddStyles adds multiple CSS properties
func (l *Link) AddStyles(m map[string]string) *Link {
	for k, v := range m {
		l.style[k] = v
	}
	return l
}

// Style replaces all styles
func (l *Link) Style(m map[string]string) *Link {
	l.style = make(map[string]string)
	for k, v := range m {
		l.style[k] = v
	}
	return l
}

// Meta represents the HTML meta element for metadata
type Meta struct {
	buf       bytes.Buffer
	style     map[string]string
	name      string
	content   string
	charset   string
	property  string
	httpEquiv string
	scheme    string
}

// NewMeta creates a new Meta element
func NewMeta() *Meta {
	return &Meta{
		style: make(map[string]string),
	}
}

// Bytes returns the buffer contents
func (m *Meta) Bytes() []byte {
	return m.buf.Bytes()
}

// Prepare builds the HTML for the meta element
func (m *Meta) Prepare() {
	m.buf.WriteString("<meta")

	if m.name != "" {
		m.buf.WriteString(" name=\"" + m.name + "\"")
	}

	if m.content != "" {
		m.buf.WriteString(" content=\"" + m.content + "\"")
	}

	if m.charset != "" {
		m.buf.WriteString(" charset=\"" + m.charset + "\"")
	}

	if m.property != "" {
		m.buf.WriteString(" property=\"" + m.property + "\"")
	}

	if m.httpEquiv != "" {
		m.buf.WriteString(" http-equiv=\"" + m.httpEquiv + "\"")
	}

	if m.scheme != "" {
		m.buf.WriteString(" scheme=\"" + m.scheme + "\"")
	}

	if len(m.style) != 0 {
		idx := 0
		m.buf.WriteString(" style=\"")
		for k, v := range m.style {
			m.buf.WriteString(k + ": " + v + ";")
			if idx != len(m.style)-1 {
				m.buf.WriteByte(' ')
			}
			idx++
		}
		m.buf.WriteString("\"")
	}

	m.buf.WriteString(">")
}

// Name sets the name attribute
func (m *Meta) Name(name string) *Meta {
	m.name = name
	return m
}

// Content sets the content attribute
func (m *Meta) Content(content string) *Meta {
	m.content = content
	return m
}

// Charset sets the charset attribute
func (m *Meta) Charset(charset string) *Meta {
	m.charset = charset
	return m
}

// Property sets the property attribute
func (m *Meta) Property(property string) *Meta {
	m.property = property
	return m
}

// HttpEquiv sets the http-equiv attribute
func (m *Meta) HttpEquiv(httpequiv string) *Meta {
	m.httpEquiv = httpequiv
	return m
}

// Scheme sets the scheme attribute
func (m *Meta) Scheme(scheme string) *Meta {
	m.scheme = scheme
	return m
}

// AddStyle adds a single CSS property
func (m *Meta) AddStyle(k, v string) *Meta {
	m.style[k] = v
	return m
}

// AddStyles adds multiple CSS properties
func (m *Meta) AddStyles(m2 map[string]string) *Meta {
	for k, v := range m2 {
		m.style[k] = v
	}
	return m
}

// Style replaces all styles
func (m *Meta) Style(m2 map[string]string) *Meta {
	m.style = make(map[string]string)
	for k, v := range m2 {
		m.style[k] = v
	}
	return m
}

// StyleElement represents the HTML style element for CSS styles
type StyleElement struct {
	buf       bytes.Buffer
	style     map[string]string
	contents  [][]byte
	ttrack    int
	styleType string
	media     string
}

// NewStyleElement creates a new StyleElement element
func NewStyleElement() *StyleElement {
	return &StyleElement{
		style: make(map[string]string),
	}
}

// Bytes returns the buffer contents
func (s *StyleElement) Bytes() []byte {
	return s.buf.Bytes()
}

// Prepare builds the HTML for the style element
func (s *StyleElement) Prepare() {
	s.buf.WriteString("<style")

	if s.styleType != "" {
		s.buf.WriteString(" type=\"" + s.styleType + "\"")
	}

	if s.media != "" {
		s.buf.WriteString(" media=\"" + s.media + "\"")
	}

	if len(s.style) != 0 {
		idx := 0
		s.buf.WriteString(" style=\"")
		for k, v := range s.style {
			s.buf.WriteString(k + ": " + v + ";")
			if idx != len(s.style)-1 {
				s.buf.WriteByte(' ')
			}
			idx++
		}
		s.buf.WriteString("\"")
	}

	s.buf.WriteByte('>')

	for _, content := range s.contents {
		s.buf.Write(content)
	}

	s.buf.WriteString("</style>")
}

// Add adds content to the style element
func (s *StyleElement) Add(e Element) *StyleElement {
	if e != nil {
		e.Prepare()
		s.contents = append(s.contents, e.Bytes())
	}
	return s
}

// Text adds CSS text content to the style element
func (s *StyleElement) Text(text string) *StyleElement {
	s.contents = append(s.contents, []byte(text))
	return s
}

// Type sets the type attribute
func (s *StyleElement) Type(styleType string) *StyleElement {
	s.styleType = styleType
	return s
}

// Media sets the media attribute
func (s *StyleElement) Media(media string) *StyleElement {
	s.media = media
	return s
}

// AddStyle adds a single CSS property
func (s *StyleElement) AddStyle(k, v string) *StyleElement {
	s.style[k] = v
	return s
}

// AddStyles adds multiple CSS properties
func (s *StyleElement) AddStyles(m map[string]string) *StyleElement {
	for k, v := range m {
		s.style[k] = v
	}
	return s
}

// Style replaces all styles
func (s *StyleElement) Style(m map[string]string) *StyleElement {
	s.style = make(map[string]string)
	for k, v := range m {
		s.style[k] = v
	}
	return s
}
