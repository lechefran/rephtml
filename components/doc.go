package rephtml

import (
	"bytes"
	"errors"
	"fmt"
	"log"
	"os"
	"reflect"
	"strings"
)

const tab = "\t"

// HtmlFile represents the HTML element for the document root
type HtmlFile struct {
	buf         bytes.Buffer
	style       map[string]string
	contents    []Element
	headContent []Element
	bodyContent []Element
	ttrack      int
	lang        string
	dir         string
	xmlLang     string
	xmlns       string
	manifest    string
	contextMenu string
	Opts        Options
	err         error
}

// NewHtmlFile creates a new HtmlFile element
func NewHtmlFile() *HtmlFile {
	return &HtmlFile{
		style: make(map[string]string),
	}
}

// Bytes returns the buffer contents
func (h *HtmlFile) Bytes() []byte {
	return cloneBytes(h.buf.Bytes())
}

// Err returns validation or rendering errors recorded while building the document.
func (h *HtmlFile) Err() error {
	return h.err
}

// addError records an error without interrupting fluent document construction.
func (h *HtmlFile) addError(err error) {
	h.err = errors.Join(h.err, err)
}

// Prepare builds the HTML for the html element
func (h *HtmlFile) Prepare() {
	h.buf.Reset()
	h.buf.WriteString("<html")
	if h.lang != "" {
		writeAttr(&h.buf, "lang", h.lang)
	}
	if h.dir != "" {
		writeAttr(&h.buf, "dir", h.dir)
	}
	if h.xmlLang != "" {
		writeAttr(&h.buf, "xml:lang", h.xmlLang)
	}
	if h.xmlns != "" {
		writeAttr(&h.buf, "xmlns", h.xmlns)
	}
	if h.manifest != "" {
		writeAttr(&h.buf, "manifest", h.manifest)
	}
	if h.contextMenu != "" {
		writeAttr(&h.buf, "contextmenu", h.contextMenu)
	}
	if len(h.style) != 0 {
		parseStyle(&h.buf, h.style)
	}
	h.buf.WriteByte('>')

	if len(h.headContent) != 0 {
		h.writeContentElement("head", h.headContent)
	}

	writeElements(&h.buf, h.contents)

	if len(h.bodyContent) != 0 {
		h.writeContentElement("body", h.bodyContent)
	}

	h.buf.WriteString("</html>")
}

// Add adds content to the html element
func (h *HtmlFile) Add(e Element) *HtmlFile {
	if e != nil {
		h.contents = appendElement(h.contents, e)
	}
	return h
}

// AddToHead sets the addtohead value on the HtmlFile component.
func (h *HtmlFile) AddToHead(e Element) *HtmlFile {
	if e == nil {
		return h
	}

	if head, ok := e.(*Head); ok && len(head.style) == 0 {
		h.headContent = append(h.headContent, head.contents...)
		return h
	}

	v := h.Opts.Validation
	if v == DEFAULT {
		if _, ok := e.(HeadElement); !ok {
			log.Printf("%s is not a valid head element but will be added to the head element\n",
				elementName(e))
		}
	} else if v == STRICT {
		if _, ok := e.(HeadElement); !ok {
			h.addError(fmt.Errorf("cannot add %s to head element", elementName(e)))
			return h
		}
	}

	h.headContent = appendElement(h.headContent, e)
	return h
}

// AddToBody sets the addtobody value on the HtmlFile component.
func (h *HtmlFile) AddToBody(e Element) *HtmlFile {
	if e == nil {
		return h
	}

	if body, ok := e.(*Body); ok && len(body.style) == 0 && body.onLoad == "" && body.onUnload == "" {
		h.bodyContent = append(h.bodyContent, body.contents...)
		return h
	}

	v := h.Opts.Validation
	if v == DEFAULT {
		if _, ok := e.(BodyElement); !ok {
			log.Printf("%s is not a valid body element but will be added to body element\n",
				elementName(e))
		}
	} else if v == STRICT {
		if _, ok := e.(BodyElement); !ok {
			h.addError(fmt.Errorf("cannot add %s to body element", elementName(e)))
			return h
		}
	}

	h.bodyContent = appendElement(h.bodyContent, e)
	return h
}

// AddOptions sets the addoptions value on the HtmlFile component.
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

// WriteToFile writes the formatted HTML document to path.
func (h *HtmlFile) WriteToFile(path string) error {
	if h.err != nil {
		return h.err
	}

	if len(h.contents) == 0 && len(h.headContent) == 0 && len(h.bodyContent) == 0 && len(h.style) == 0 {
		log.Print("No values were appended to the HTML File. " + path + " will not be created")
	}

	h.Prepare()

	file, err := os.Create(path)
	if err != nil {
		return fmt.Errorf("create html file %q: %w", path, err)
	}

	if _, err := file.Write(formatHTML(h.buf.Bytes())); err != nil {
		closeErr := file.Close()
		if closeErr != nil {
			return errors.Join(
				fmt.Errorf("write html file %q: %w", path, err),
				fmt.Errorf("close html file %q: %w", path, closeErr),
			)
		}
		return fmt.Errorf("write html file %q: %w", path, err)
	}

	if err := file.Close(); err != nil {
		return fmt.Errorf("close html file %q: %w", path, err)
	}

	return nil
}

// writeContentElement writes a generated head or body section.
func (h *HtmlFile) writeContentElement(tag string, contents []Element) {
	h.buf.WriteString("<" + tag + ">")
	writeElements(&h.buf, contents)
	h.buf.WriteString("</" + tag + ">")
}

// elementName returns a stable type name for error and warning messages.
func elementName(e Element) string {
	if e == nil {
		return "<nil>"
	}
	t := reflect.TypeOf(e)
	for t.Kind() == reflect.Pointer {
		t = t.Elem()
	}
	if t.Name() == "" {
		return t.String()
	}
	return t.Name()
}

// htmlToken stores one formatter token.
type htmlToken struct {
	kind string
	name string
	text string
}

// formatHTML returns an indented representation of rendered HTML bytes.
func formatHTML(src []byte) []byte {
	tokens := tokenizeHTML(string(src))
	if len(tokens) == 0 {
		return src
	}

	var buf bytes.Buffer
	indent := 0

	for i := 0; i < len(tokens); i++ {
		token := tokens[i]
		switch token.kind {
		case "raw":
			writeRawBlock(&buf, indent, token.name, token.text)
		case "start":
			if i+1 < len(tokens) && tokens[i+1].kind == "end" && tokens[i+1].name == token.name {
				writeFormattedLine(&buf, indent, token.text+tokens[i+1].text)
				i++
				continue
			}
			if i+2 < len(tokens) && tokens[i+1].kind == "text" && tokens[i+2].kind == "end" &&
				tokens[i+2].name == token.name && !isRawTextElement(token.name) && !strings.Contains(tokens[i+1].text, "\n") {
				writeFormattedLine(&buf, indent, token.text+tokens[i+1].text+tokens[i+2].text)
				i += 2
				continue
			}
			writeFormattedLine(&buf, indent, token.text)
			indent++
		case "end":
			if indent > 0 {
				indent--
			}
			writeFormattedLine(&buf, indent, token.text)
		case "void", "comment":
			writeFormattedLine(&buf, indent, token.text)
		case "text":
			writeFormattedText(&buf, indent, token.text)
		}
	}

	return buf.Bytes()
}

// tokenizeHTML converts rendered HTML into formatter tokens.
func tokenizeHTML(input string) []htmlToken {
	tokens := []htmlToken{}
	for i := 0; i < len(input); {
		if input[i] != '<' || !isHTMLTagStart(input, i) {
			next := nextHTMLTagStart(input, i+1)
			if next == -1 {
				tokens = appendTextToken(tokens, input[i:])
				break
			}
			tokens = appendTextToken(tokens, input[i:next])
			i = next
			continue
		}

		if strings.HasPrefix(input[i:], "<!--") {
			end := strings.Index(input[i:], "-->")
			if end == -1 {
				tokens = append(tokens, htmlToken{kind: "comment", text: input[i:]})
				break
			}
			end += len("-->")
			tokens = append(tokens, htmlToken{kind: "comment", text: input[i : i+end]})
			i += end
			continue
		}

		end := findTagEnd(input, i)
		if end == -1 {
			tokens = appendTextToken(tokens, input[i:])
			break
		}

		tag := input[i : end+1]
		name := tagName(tag)
		if name == "" {
			tokens = appendTextToken(tokens, tag)
			i = end + 1
			continue
		}

		switch {
		case isEndTag(tag):
			tokens = append(tokens, htmlToken{kind: "end", name: name, text: tag})
		case isVoidTag(name) || isSelfClosingTag(tag):
			tokens = append(tokens, htmlToken{kind: "void", name: name, text: tag})
		default:
			if isRawTextElement(name) {
				rawStart := end + 1
				closeTag := "</" + name + ">"
				closeAt := strings.Index(strings.ToLower(input[rawStart:]), closeTag)
				if closeAt != -1 {
					rawEnd := rawStart + closeAt
					tokens = append(tokens, htmlToken{
						kind: "raw",
						name: name,
						text: input[i : rawEnd+len(closeTag)],
					})
					i = rawEnd + len(closeTag)
					continue
				}
			}
			tokens = append(tokens, htmlToken{kind: "start", name: name, text: tag})
		}

		i = end + 1
	}
	return tokens
}

// appendTextToken appends meaningful text content to the token stream.
func appendTextToken(tokens []htmlToken, text string) []htmlToken {
	if text == "" {
		return tokens
	}
	if strings.TrimSpace(text) == "" {
		return tokens
	}
	return append(tokens, htmlToken{kind: "text", text: text})
}

// writeFormattedLine writes one indented formatter line.
func writeFormattedLine(buf *bytes.Buffer, indent int, line string) {
	buf.WriteString(tabs(indent))
	buf.WriteString(line)
	buf.WriteByte('\n')
}

// writeFormattedText writes formatted text lines.
func writeFormattedText(buf *bytes.Buffer, indent int, text string) {
	for _, line := range strings.Split(text, "\n") {
		if line == "" {
			continue
		}
		writeFormattedLine(buf, indent, line)
	}
}

// writeRawBlock writes raw-text elements without corrupting sensitive content.
func writeRawBlock(buf *bytes.Buffer, indent int, name, text string) {
	if name == "style" {
		if openTag, body, closeTag, ok := splitRawElement(text, name); ok {
			writeFormattedLine(buf, indent, openTag)
			writeIndentedRawLines(buf, indent+1, body)
			writeFormattedLine(buf, indent, closeTag)
			return
		}
	}

	buf.WriteString(tabs(indent))
	buf.WriteString(text)
	if !strings.HasSuffix(text, "\n") {
		buf.WriteByte('\n')
	}
}

// splitRawElement separates a raw-text element into open tag, body, and close tag.
func splitRawElement(text, name string) (string, string, string, bool) {
	openEnd := findTagEnd(text, 0)
	if openEnd == -1 {
		return "", "", "", false
	}

	closeTag := "</" + name + ">"
	closeAt := strings.LastIndex(strings.ToLower(text), closeTag)
	if closeAt == -1 || closeAt < openEnd {
		return "", "", "", false
	}

	return text[:openEnd+1], text[openEnd+1 : closeAt], text[closeAt:], true
}

// writeIndentedRawLines writes raw block body lines under the current indent.
func writeIndentedRawLines(buf *bytes.Buffer, indent int, text string) {
	text = strings.Trim(text, "\n")
	if text == "" {
		return
	}

	for _, line := range strings.Split(text, "\n") {
		if strings.TrimSpace(line) == "" {
			buf.WriteByte('\n')
			continue
		}
		writeFormattedLine(buf, indent, line)
	}
}

// tagName extracts the lower-case element name from an HTML tag.
func tagName(tag string) string {
	tag = strings.TrimSpace(tag)
	tag = strings.TrimPrefix(tag, "<")
	tag = strings.TrimPrefix(tag, "/")
	tag = strings.TrimSuffix(tag, ">")
	tag = strings.TrimSpace(tag)
	if tag == "" || strings.HasPrefix(tag, "!") {
		return ""
	}
	if idx := strings.IndexAny(tag, " \t\r\n/"); idx != -1 {
		tag = tag[:idx]
	}
	return strings.ToLower(tag)
}

// isEndTag reports whether a tag is a closing tag.
func isEndTag(tag string) bool {
	return strings.HasPrefix(strings.TrimSpace(tag), "</")
}

// isSelfClosingTag reports whether a tag ends with a self-closing slash.
func isSelfClosingTag(tag string) bool {
	tag = strings.TrimSpace(tag)
	tag = strings.TrimSuffix(tag, ">")
	return strings.HasSuffix(strings.TrimSpace(tag), "/")
}

// isVoidTag reports whether a tag name is an HTML void element.
func isVoidTag(tag string) bool {
	switch tag {
	case "area", "base", "br", "col", "embed", "hr", "img", "input", "link", "meta", "source", "track", "wbr":
		return true
	default:
		return false
	}
}

// isRawTextElement reports whether a tag body needs raw-text handling.
func isRawTextElement(tag string) bool {
	switch tag {
	case "script", "style", "pre", "textarea":
		return true
	default:
		return false
	}
}

// nextHTMLTagStart returns the next likely HTML tag start position.
func nextHTMLTagStart(input string, from int) int {
	for i := from; i < len(input); i++ {
		if input[i] == '<' && isHTMLTagStart(input, i) {
			return i
		}
	}
	return -1
}

// isHTMLTagStart reports whether a less-than sign starts likely HTML markup.
func isHTMLTagStart(input string, idx int) bool {
	if idx+1 >= len(input) || input[idx] != '<' {
		return false
	}
	if strings.HasPrefix(input[idx:], "<!--") {
		return true
	}

	next := input[idx+1]
	if next == '/' {
		return idx+2 < len(input) && isTagNameStart(input[idx+2])
	}
	if next == '!' {
		return strings.HasPrefix(strings.ToLower(input[idx:]), "<!doctype")
	}
	return isTagNameStart(next)
}

// isTagNameStart reports whether a byte can start an HTML tag name.
func isTagNameStart(c byte) bool {
	return (c >= 'a' && c <= 'z') || (c >= 'A' && c <= 'Z')
}

// findTagEnd returns the closing greater-than sign while respecting quotes.
func findTagEnd(input string, start int) int {
	var quote byte
	for i := start + 1; i < len(input); i++ {
		c := input[i]
		if quote != 0 {
			if c == quote {
				quote = 0
			}
			continue
		}
		switch c {
		case '\'', '"':
			quote = c
		case '>':
			return i
		}
	}
	return -1
}

// Head represents the HTML head element for document metadata
type Head struct {
	buf      bytes.Buffer
	style    map[string]string
	contents []Element
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
	return cloneBytes(h.buf.Bytes())
}

// Prepare builds the HTML for the head element
func (h *Head) Prepare() {
	h.buf.Reset()
	h.buf.WriteString("<head")
	if len(h.style) != 0 {
		parseStyle(&h.buf, h.style)
	}
	h.buf.WriteByte('>')

	writeElements(&h.buf, h.contents)
	h.buf.WriteString("</head>")
}

// Add adds content to the head element
func (h *Head) Add(e Element) *Head {
	if e != nil {
		h.contents = appendElement(h.contents, e)
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
	contents []Element
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
	return cloneBytes(b.buf.Bytes())
}

// IsBodyElement implements BodyElement interface
func (b *Body) IsBodyElement() {}

// Prepare builds the HTML for the body element
func (b *Body) Prepare() {
	b.buf.Reset()
	b.buf.WriteString("<body")
	if b.onLoad != "" {
		writeAttr(&b.buf, "onload", b.onLoad)
	}
	if b.onUnload != "" {
		writeAttr(&b.buf, "onunload", b.onUnload)
	}
	if len(b.style) != 0 {
		parseStyle(&b.buf, b.style)
	}
	b.buf.WriteByte('>')

	writeElements(&b.buf, b.contents)
	b.buf.WriteString("</body>")
}

// Add adds content to the body element
func (b *Body) Add(e Element) *Body {
	if e != nil {
		b.contents = appendElement(b.contents, e)
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
	contents []Element
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
	return cloneBytes(t.buf.Bytes())
}

// Prepare builds the HTML for the title element
func (t *Title) Prepare() {
	t.buf.Reset()
	t.buf.WriteString("<title")
	if len(t.style) != 0 {
		parseStyle(&t.buf, t.style)
	}
	t.buf.WriteByte('>')

	writeElements(&t.buf, t.contents)
	t.buf.WriteString("</title>")
}

// Add adds content to the title element
func (t *Title) Add(e Element) *Title {
	if e != nil {
		t.contents = appendElement(t.contents, e)
	}
	return t
}

// IsHeadElement implements HeadElement interface
func (t *Title) IsHeadElement() {}

// Text adds text content to the title element
func (t *Title) Text(text string) *Title {
	t.contents = appendElement(t.contents, escapedText(text))
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
	return cloneBytes(b.buf.Bytes())
}

// Prepare builds the HTML for the base element
func (b *Base) Prepare() {
	b.buf.Reset()
	b.buf.WriteString("<base")
	if b.href != "" {
		writeAttr(&b.buf, "href", b.href)
	}
	if b.target != "" {
		writeAttr(&b.buf, "target", b.target)
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
	return cloneBytes(l.buf.Bytes())
}

// Prepare builds the HTML for the link element
func (l *Link) Prepare() {
	l.buf.Reset()
	l.buf.WriteString("<link")
	if l.rel != "" {
		writeAttr(&l.buf, "rel", l.rel)
	}
	if l.href != "" {
		writeAttr(&l.buf, "href", l.href)
	}
	if l.linkType != "" {
		writeAttr(&l.buf, "type", l.linkType)
	}
	if l.media != "" {
		writeAttr(&l.buf, "media", l.media)
	}
	if l.sizes != "" {
		writeAttr(&l.buf, "sizes", l.sizes)
	}
	if l.crossOrigin != "" {
		writeAttr(&l.buf, "crossOrigin", l.crossOrigin)
	}
	if l.integrity != "" {
		writeAttr(&l.buf, "integrity", l.integrity)
	}
	if l.referrerPolicy != "" {
		writeAttr(&l.buf, "referrerPolicy", l.referrerPolicy)
	}
	if l.hreflang != "" {
		writeAttr(&l.buf, "hreflang", l.hreflang)
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
	return cloneBytes(m.buf.Bytes())
}

// Prepare builds the HTML for the meta element
func (m *Meta) Prepare() {
	m.buf.Reset()
	m.buf.WriteString("<meta")
	if m.name != "" {
		writeAttr(&m.buf, "name", m.name)
	}
	if m.content != "" {
		writeAttr(&m.buf, "content", m.content)
	}
	if m.charset != "" {
		writeAttr(&m.buf, "charset", m.charset)
	}
	if m.property != "" {
		writeAttr(&m.buf, "property", m.property)
	}
	if m.httpEquiv != "" {
		writeAttr(&m.buf, "http-equiv", m.httpEquiv)
	}
	if m.scheme != "" {
		writeAttr(&m.buf, "scheme", m.scheme)
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
	contents  []Element
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
	return cloneBytes(s.buf.Bytes())
}

// Prepare builds the HTML for the style element
func (s *StyleElement) Prepare() {
	s.buf.Reset()
	s.buf.WriteString("<style")
	if s.styleType != "" {
		writeAttr(&s.buf, "type", s.styleType)
	}
	if s.media != "" {
		writeAttr(&s.buf, "media", s.media)
	}
	if len(s.style) != 0 {
		parseStyle(&s.buf, s.style)
	}

	s.buf.WriteByte('>')

	writeElements(&s.buf, s.contents)

	s.buf.WriteString("</style>")
}

// Add adds content to the style element
func (s *StyleElement) Add(e Element) *StyleElement {
	switch element := e.(type) {
	case nil:
		return s
	case *StyleRule:
		return s.AddRule(element)
	case *Style:
		if element == nil {
			return s
		}
		return s.AddRule(&StyleRule{
			pmap:  element.pmap,
			Props: element.Props,
			Tags:  append([]string(nil), element.Tags...),
		})
	default:
		s.contents = appendElement(s.contents, e)
	}
	return s
}

// AddRule sets the addrule value on the StyleElement component.
func (s *StyleElement) AddRule(rule *StyleRule) *StyleElement {
	if rule != nil {
		s.contents = appendElement(s.contents, rule)
	}
	return s
}

// Text adds CSS text content to the style element
func (s *StyleElement) Text(text string) *StyleElement {
	s.contents = appendElement(s.contents, rawText(text))
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

// IsBodyElement implements BodyElement interface
func (s *StyleElement) IsBodyElement() {}
