package rephtml

import (
	"bytes"
	"log"
	"os"
	"reflect"
)

const tab = "\t"

// HtmlFile represents the HTML element for the document root
type HtmlFile struct {
	buf         bytes.Buffer
	style       map[string]string
	contents    [][]byte
	ttrack      int
	lang        string
	dir         string
	xmlLang     string
	xmlns       string
	manifest    string
	contextMenu string
	Opts        Options
}

// NewHtmlFile creates a new HtmlFile element
func NewHtmlFile() *HtmlFile {
	return &HtmlFile{
		style: make(map[string]string),
	}
}

// Bytes returns the buffer contents
func (h *HtmlFile) Bytes() []byte {
	return h.buf.Bytes()
}

// Prepare builds the HTML for the html element
func (h *HtmlFile) Prepare() {
	h.buf.Reset()
	h.buf.WriteString("<html")
	if h.lang != "" {
		h.buf.WriteString(" lang=\"" + h.lang + "\"")
	}
	if h.dir != "" {
		h.buf.WriteString(" dir=\"" + h.dir + "\"")
	}
	if h.xmlLang != "" {
		h.buf.WriteString(" xml:lang=\"" + h.xmlLang + "\"")
	}
	if h.xmlns != "" {
		h.buf.WriteString(" xmlns=\"" + h.xmlns + "\"")
	}
	if h.manifest != "" {
		h.buf.WriteString(" manifest=\"" + h.manifest + "\"")
	}
	if h.contextMenu != "" {
		h.buf.WriteString(" contextmenu=\"" + h.contextMenu + "\"")
	}
	if len(h.style) != 0 {
		parseStyle(&h.buf, h.style)
	}
	h.buf.WriteByte('>')

	for _, content := range h.contents {
		h.buf.Write(content)
	}
	h.buf.WriteString("</html>")
}

// Add adds content to the html element
func (h *HtmlFile) Add(e Element) *HtmlFile {
	if e != nil {
		if len(e.Bytes()) == 0 {
			e.Prepare()
		}
		h.contents = append(h.contents, e.Bytes())
	}
	return h
}

func (h *HtmlFile) AddToHead(e Element) *HtmlFile {
	if e != nil {
		constitution := h.Opts.Constitution
		if constitution == DEFAULT {
			if _, ok := e.(HeadElement); !ok {
				log.Printf("%s is not a valid head element but will be added to the head element\n",
					reflect.TypeOf(e).Elem().Name())
			}
		} else if constitution == STRICT {
			if _, ok := e.(HeadElement); !ok {
				log.Fatalf("Cannot add %s to head element\n", reflect.TypeOf(e).Elem().Name())
			}
		}

		if len(e.Bytes()) == 0 {
			e.Prepare()
		}
		h.contents = append(h.contents, e.Bytes())
	}
	return h
}

func (h *HtmlFile) AddToBody(e Element) *HtmlFile {
	if e != nil {
		if len(e.Bytes()) == 0 {
			e.Prepare()
		}
		h.contents = append(h.contents, e.Bytes())
	}
	return h
}

func (h *HtmlFile) AddOptions(o Options) *HtmlFile {
	h.Opts = o
	return h
}

// Lang sets the lang attribute
func (h *HtmlFile) Lang(lang string) *HtmlFile {
	h.lang = lang
	return h
}

// Dir sets the dir attribute
func (h *HtmlFile) Dir(dir string) *HtmlFile {
	h.dir = dir
	return h
}

// XmlLang sets the xml:lang attribute
func (h *HtmlFile) XmlLang(xmlLang string) *HtmlFile {
	h.xmlLang = xmlLang
	return h
}

// Xmlns sets the xmlns attribute
func (h *HtmlFile) Xmlns(xmlns string) *HtmlFile {
	h.xmlns = xmlns
	return h
}

// Manifest sets the manifest attribute
func (h *HtmlFile) Manifest(manifest string) *HtmlFile {
	h.manifest = manifest
	return h
}

// ContextMenu sets the contextmenu attribute
func (h *HtmlFile) ContextMenu(contextMenu string) *HtmlFile {
	h.contextMenu = contextMenu
	return h
}

// AddStyle adds a single CSS property
func (h *HtmlFile) AddStyle(k, v string) *HtmlFile {
	h.style[k] = v
	return h
}

// AddStyles adds multiple CSS properties
func (h *HtmlFile) AddStyles(m map[string]string) *HtmlFile {
	for k, v := range m {
		h.style[k] = v
	}
	return h
}

// Style replaces all styles
func (h *HtmlFile) Style(m map[string]string) *HtmlFile {
	h.style = make(map[string]string)
	for k, v := range m {
		h.style[k] = v
	}
	return h
}

func (h *HtmlFile) WriteToFile(path string) {
	if len(h.contents) == 0 && len(h.style) == 0 {
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
	h.buf.Reset()
	h.buf.WriteString("<head")
	if len(h.style) != 0 {
		parseStyle(&h.buf, h.style)
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
	b.buf.Reset()
	b.buf.WriteString("<body")
	if b.onLoad != "" {
		b.buf.WriteString(" onload=\"" + b.onLoad + "\"")
	}
	if b.onUnload != "" {
		b.buf.WriteString(" onunload=\"" + b.onUnload + "\"")
	}
	if len(b.style) != 0 {
		parseStyle(&b.buf, b.style)
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
	t.buf.Reset()
	t.buf.WriteString("<title")
	if len(t.style) != 0 {
		parseStyle(&t.buf, t.style)
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

// IsHeadElement implements HeadElement interface
func (t *Title) IsHeadElement() {}

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
	b.buf.Reset()
	b.buf.WriteString("<base")
	if b.href != "" {
		b.buf.WriteString(" href=\"" + b.href + "\"")
	}
	if b.target != "" {
		b.buf.WriteString(" target=\"" + b.target + "\"")
	}
	if len(b.style) != 0 {
		parseStyle(&b.buf, b.style)
	}
	b.buf.WriteString(">")
}

// IsHeadElement implements HeadElement interface
func (b *Base) IsHeadElement() {}

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
	l.buf.Reset()
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
		parseStyle(&l.buf, l.style)
	}
	l.buf.WriteString(">")
}

// IsHeadElement implements HeadElement interface
func (l *Link) IsHeadElement() {}

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
	m.buf.Reset()
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
		parseStyle(&m.buf, m.style)
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

// IsHeadElement implements HeadElement interface
func (m *Meta) IsHeadElement() {}

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
	s.buf.Reset()
	s.buf.WriteString("<style")
	if s.styleType != "" {
		s.buf.WriteString(" type=\"" + s.styleType + "\"")
	}
	if s.media != "" {
		s.buf.WriteString(" media=\"" + s.media + "\"")
	}
	if len(s.style) != 0 {
		parseStyle(&s.buf, s.style)
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

// IsHeadElement implements HeadElement interface
func (s *StyleElement) IsHeadElement() {}
