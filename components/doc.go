package rephtml

import (
	"bytes"
	"log"
	"os"
	"strings"
)

const newline = "\n"
const tab = "\t"

type HtmlFile struct {
	buf         bytes.Buffer
	head        []byte
	style, body [][]byte
	lang        string
	title       string
	ttrack      int // tab tracker
	base        Base
	options     Options
}

type Base struct {
	link, target string
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

func (h *HtmlFile) Base(b Base) *HtmlFile {
	h.base = b
	return h
}

// vv element struct functions vv

func (h *HtmlFile) Div(d *Div) *HtmlFile {
	d.Tabs(h.ttrack)
	h.body = append(h.body, d.Bytes())
	return h
}

func (h *HtmlFile) Style(s *Style) *HtmlFile {
	h.style = append(h.style, s.Bytes())
	return h
}

func (h *HtmlFile) Table(t *Table) *HtmlFile {
	h.body = append(h.body, t.Bytes())
	return h
}

// vv element string functions vv

func (h *HtmlFile) H1String(s string) *HtmlFile {
	h.body = append(h.body, []byte("<h1>"+s+"</h1>"))
	return h
}

func (h *HtmlFile) H2String(s string) *HtmlFile {
	h.body = append(h.body, []byte("<h2>"+s+"</h2>"))
	return h
}

func (h *HtmlFile) H3String(s string) *HtmlFile {
	h.body = append(h.body, []byte("<h3>"+s+"</h3>"))
	return h
}

func (h *HtmlFile) H4String(s string) *HtmlFile {
	h.body = append(h.body, []byte("<h4>"+s+"</h4>"))
	return h
}

func (h *HtmlFile) H5String(s string) *HtmlFile {
	h.body = append(h.body, []byte("<h5>"+s+"</h5>"))
	return h
}

func (h *HtmlFile) H6String(s string) *HtmlFile {
	h.body = append(h.body, []byte("<h6>"+s+"</h6>"))
	return h
}

/*
PString

Add a paragraph element to the HTML document with
a string parameter as the assigned value
*/
func (h *HtmlFile) PString(s string) *HtmlFile {
	h.body = append(h.body, []byte("<p>"+s+"</p>"))
	return h
}

/*
PStringWithStyle

Add a paragraph element to the HTML document with
a string parameter as the assigned value and a style value
*/
func (h *HtmlFile) PStringWithStyle(s, style string) *HtmlFile {
	if style == "" {
		h.body = append(h.body, []byte("<p>"+s+"</p>"))
	} else {
		h.body = append(h.body, []byte("<p style=\""+style+";\">"+s+"</p>"))
	}
	return h
}

/*
P

Add a paragraph element to the HTML document with
a paragraph struct
*/
func (h *HtmlFile) P(p *P) *HtmlFile {
	h.body = append(h.body, p.Bytes())
	return h
}

func (h *HtmlFile) StyleString(s string) *HtmlFile {
	fs := strings.ReplaceAll(s, ";", "; ")
	h.style = append(h.style, []byte(fs))
	return h
}

func (h *HtmlFile) TableString(harr []string, darr [][]string) *HtmlFile {
	tbl := "<table>"

	// write headers
	tbl += "<tr>"
	for _, h := range harr {
		tbl += "<th>" + h + "</th>"
	}
	tbl += "</tr>"

	// write rows
	for _, d := range darr {
		tbl += "<tr>"
		for _, d1 := range d {
			tbl += "<td>" + d1 + "</td>"
		}
		tbl += "</tr>"
	}
	tbl += "</table>"
	h.body = append(h.body, []byte(tbl))
	return h
}

// vv general functions vv

func (h *HtmlFile) Bytes() []byte {
	return h.buf.Bytes()
}

func (h *HtmlFile) Prepare() *HtmlFile {
	t := tabs(h.ttrack)
	if h.lang != "" {
		h.buf.WriteString("<html>" + newline)
	} else {
		h.buf.WriteString("<html lang=\"" + h.lang + "\">" + newline)
	}
	h.buf.WriteString(t + "<head>" + newline)
	h.ttrack++
	t = tabs(h.ttrack)
	if h.title != "" {
		h.buf.WriteString("<title>" + h.title + "</title>" + newline)
	}
	if h.base.link != "" {
		h.buf.WriteString("<base href=\"" + h.base.link)
		if h.base.target != "" {
			h.buf.WriteString("\" target=\"" + h.base.target)
		}
		h.buf.WriteString("\">" + newline)
	}
	if len(h.style) > 0 {
		h.buf.WriteString(t + "<style>" + newline)
		h.ttrack++
		t = tabs(h.ttrack)
		for i := 0; i < len(h.style); i++ {
			if i != len(h.style)-1 {
				h.buf.Write(h.formatStyle(h.style[i]))
				h.buf.WriteString(newline)
			} else {
				h.buf.Write(h.formatStyle(h.style[i]))
			}
		}

		h.ttrack--
		t = tabs(h.ttrack)
		h.buf.WriteString(t + "</style>" + newline)
		h.ttrack--
	}
	t = tabs(h.ttrack)
	h.buf.WriteString(t + "</head>" + newline)
	h.buf.WriteString(t + "<body>" + newline)
	h.ttrack++
	t = tabs(h.ttrack)

	for i := 0; i < len(h.body); i++ {
		if bytes.Contains(h.body[i], []byte("<div")) &&
			bytes.Contains(h.body[i], []byte(">")) {
			h.buf.Write(h.formatDiv(h.body[i]))
		} else if bytes.Contains(h.body[i], []byte("<table")) &&
			bytes.Contains(h.body[i], []byte(">")) {
			h.buf.Write(h.formatTable(h.body[i]))
		} else {
			h.buf.WriteString(t)
			h.buf.Write(h.body[i])
			h.buf.WriteString(newline)
		}
	}

	h.ttrack--
	h.buf.WriteString(tabs(h.ttrack) + "</body>" + newline)
	h.buf.WriteString("</html>")
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

// vv helper functions vv

/*
Internal parsing function to format div element and its contents
*/
func (h *HtmlFile) formatDiv(b []byte) []byte {
	var fb bytes.Buffer

	// split byte array by tags
	var curr []byte
	var sarr [][]byte
	for i := 0; i < len(b)-1; i++ {
		curr = append(curr, b[i])
		if b[i] == '>' && b[i+1] == '<' {
			sarr = append(sarr, curr)
			curr = []byte{} // reset values of curr
		}
		if i+1 == len(b)-1 {
			curr = append(curr, '>')
			sarr = append(sarr, curr)
		}
	}

	for _, s := range sarr {
		if bytes.Contains(s, []byte("<div")) && bytes.Contains(s, []byte(">")) {
			fb.WriteString(tabs(h.ttrack))
			fb.Write(s)
			h.ttrack++
		} else if bytes.Equal(s, []byte("</div>")) {
			h.ttrack--
			fb.WriteByte('\n')
			fb.WriteString(tabs(h.ttrack))
			fb.Write(s)
		} else {
			fb.WriteByte('\n')
			fb.WriteString(tabs(h.ttrack))
			fb.Write(s)
		}
	}
	fb.WriteByte('\n')
	return fb.Bytes()
}

/*
Internal parsing function to format style attributes
*/
func (h *HtmlFile) formatStyle(b []byte) []byte {
	var fb bytes.Buffer
	nsb := strip(b) // remove all spaces

	// get indexes for open and close braces
	op, cls := bytes.Index(nsb, []byte("{"))+1, len(nsb)-1

	// cut array into parts: opening, contents, and closing
	opb, clsb := nsb[:op], nsb[cls]
	contents := nsb[op:cls]

	// add spacing between open values
	var opbTmp []byte
	for i := 0; i < len(opb); i++ {
		if opb[i] == ',' {
			opbTmp = append(opbTmp, opb[i])
			opbTmp = append(opbTmp, ' ')
		} else if opb[i] == '{' {
			opbTmp = append(opbTmp, ' ')
			opbTmp = append(opbTmp, opb[i])
		} else {
			opbTmp = append(opbTmp, opb[i])
		}
	}
	opbTmp = append(opbTmp, '\n')
	opb = opbTmp

	// write opb to buffer
	fb.WriteString(tabs(h.ttrack))
	fb.Write(opb)

	// next, process contents
	contentsTmp := make([]byte, 0, len(contents))
	for _, b := range contents {
		if b == ';' {
			contentsTmp = append(contentsTmp, ' ')
		} else {
			contentsTmp = append(contentsTmp, b)
		}
	}
	contents = contentsTmp

	// split contents
	carr := bytes.Fields(contents)
	carrTmp := make([][]byte, 0, len(carr))
	for _, c := range carr {
		cTmp := []byte{}
		for i := 0; i < len(c); i++ {
			if c[i] == ':' {
				cTmp = append(cTmp, c[i])
				cTmp = append(cTmp, ' ')
			} else {
				cTmp = append(cTmp, c[i])
			}
		}
		cTmp = append(cTmp, ';')
		cTmp = append(cTmp, '\n')
		carrTmp = append(carrTmp, cTmp)
	}
	carr = carrTmp

	h.ttrack++
	for _, c := range carr {
		fb.WriteString(tabs(h.ttrack))
		fb.Write(c)
	}
	h.ttrack--
	fb.WriteString(tabs(h.ttrack))
	fb.WriteByte(clsb)
	fb.WriteByte('\n')
	return fb.Bytes()
}

/*
Internal parsing function to format table elements
*/
func (h *HtmlFile) formatTable(b []byte) []byte {
	var fb bytes.Buffer

	// see if open table tag has id and class values
	var tmp bytes.Buffer
	var opent []byte
	for _, c := range b {
		if c != '>' {
			tmp.WriteByte(c)
		} else {
			tmp.WriteByte('>')
			break
		}
	}
	opent = tmp.Bytes()
	b = b[bytes.IndexByte(b, '>')+1:]
	nsb := strip(b) // remove all spaces

	// split byte array by tags
	var nsbsplit []byte
	for i := 0; i < len(nsb)-1; i++ {
		nsbsplit = append(nsbsplit, nsb[i])
		if nsb[i] == '>' && nsb[i+1] == '<' {
			nsbsplit = append(nsbsplit, ' ')
		}
		if i+1 == len(nsb)-1 {
			nsbsplit = append(nsbsplit, '>')
		}
	}
	sarr := bytes.Fields(nsbsplit)
	contents := sarr[:len(sarr)-1] // remove closing tag, and obtain contents

	// write open tag to buffer
	fb.WriteString(tabs(h.ttrack))
	fb.Write(opent)
	fb.WriteByte('\n')
	h.ttrack++

	// loop through contents
	for _, c := range contents {
		if bytes.Equal(c, []byte("<tr>")) {
			fb.WriteString(tabs(h.ttrack))
			fb.Write(c)
			fb.WriteByte('\n')
			h.ttrack++
		} else if bytes.Equal(c, []byte("</tr>")) {
			h.ttrack--
			fb.WriteString(tabs(h.ttrack))
			fb.Write(c)
			fb.WriteByte('\n')
		} else {
			fb.WriteString(tabs(h.ttrack))
			fb.Write(c)
			fb.WriteByte('\n')
		}
	}

	// write close tag to header
	h.ttrack--
	fb.WriteString(tabs(h.ttrack))
	fb.Write(sarr[len(sarr)-1])
	fb.WriteByte('\n')
	return fb.Bytes()
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

// Cross Origin sets the crossOrigin attribute
func (l *Link) CrossOrigin(crossOrigin string) *Link {
	l.crossOrigin = crossOrigin
	return l
}

// Integrity sets the integrity attribute
func (l *Link) Integrity(integrity string) *Link {
	l.integrity = integrity
	return l
}

// Referrer Policy sets the referrerPolicy attribute
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
func (s *StyleElement) Add(e Elements) *StyleElement {
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
