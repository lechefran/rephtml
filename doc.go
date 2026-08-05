package rephtml

import (
	"bytes"
	"errors"
	"fmt"
	"os"
	"reflect"
	"strings"
)

const tab = "\t"

// HtmlFile represents the HTML element for the document root
type HtmlFile struct {
	style        StyleMap
	contents     []Element
	headContent  []Element
	bodyContent  []Element
	headStyle    StyleMap
	bodyStyle    StyleMap
	lang         string
	dir          string
	xmlLang      string
	xmlns        string
	manifest     string
	contextMenu  string
	bodyOnLoad   string
	bodyOnUnload string
	err          error
}

// NewHtmlFile creates a new HtmlFile element
func NewHtmlFile() *HtmlFile {
	return &HtmlFile{
		style:     make(StyleMap),
		headStyle: make(StyleMap),
		bodyStyle: make(StyleMap),
	}
}

// Err returns document structure or rendering errors recorded while building.
func (h *HtmlFile) Err() error {
	return h.err
}

// addError records an error without interrupting fluent document construction.
func (h *HtmlFile) addError(err error) {
	h.err = errors.Join(h.err, err)
}

// doctype is emitted ahead of every document. Without it browsers fall back to
// quirks mode, which changes box sizing and several inherited layout rules.
const doctype = "<!DOCTYPE html>"

// renderTo writes the document's HTML to buf.
func (h *HtmlFile) renderTo(buf *bytes.Buffer) {
	buf.WriteString(doctype)

	tg := startTag(buf, "html")
	tg.attr("lang", h.lang)
	tg.attr("dir", h.dir)
	tg.attr("xml:lang", h.xmlLang)
	tg.attr("xmlns", h.xmlns)
	tg.urlAttr("manifest", h.manifest)
	tg.attr("contextmenu", h.contextMenu)
	tg.styleAttr(h.style)
	tg.open()

	if len(h.headContent) != 0 || len(h.headStyle) != 0 {
		h.writeHeadElement(buf)
	}

	writeElements(buf, h.contents)

	if len(h.bodyContent) != 0 || len(h.bodyStyle) != 0 || h.bodyOnLoad != "" || h.bodyOnUnload != "" {
		h.writeBodyElement(buf)
	}

	tg.end()
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

	if head, ok := e.(*Head); ok {
		h.mergeHead(head)
		return h
	}

	if _, ok := e.(*Body); ok {
		h.rejectStructuralElement("head", e)
		return h
	}

	if _, ok := e.(HeadElement); !ok {
		h.addError(fmt.Errorf("cannot add %s to head element", elementName(e)))
		return h
	}

	h.headContent = appendElement(h.headContent, e)
	return h
}

// AddToBody sets the addtobody value on the HtmlFile component.
func (h *HtmlFile) AddToBody(e Element) *HtmlFile {
	if e == nil {
		return h
	}

	if body, ok := e.(*Body); ok {
		h.mergeBody(body)
		return h
	}

	if _, ok := e.(*Head); ok {
		h.rejectStructuralElement("body", e)
		return h
	}

	if _, ok := e.(BodyElement); !ok {
		h.addError(fmt.Errorf("cannot add %s to body element", elementName(e)))
		return h
	}

	h.bodyContent = appendElement(h.bodyContent, e)
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
func (h *HtmlFile) AddStyles(m StyleMap) *HtmlFile {
	for k, v := range m {
		h.style[k] = v
	}
	return h
}

// Style replaces all styles
func (h *HtmlFile) Style(m StyleMap) *HtmlFile {
	h.style = cloneStyleMap(m)
	return h
}

// WriteToFile writes the formatted HTML document to path.
func (h *HtmlFile) WriteToFile(path string) error {
	html, err := h.RenderFormatted()
	if err != nil {
		return err
	}

	file, err := os.Create(path)
	if err != nil {
		return fmt.Errorf("create html file %q: %w", path, err)
	}

	if _, err := file.Write(html); err != nil {
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

// Render returns the compact HTML document bytes.
func (h *HtmlFile) Render() ([]byte, error) {
	html, err := h.renderDocument()
	if err != nil {
		return nil, err
	}
	return []byte(html), nil
}

// renderDocument returns compact HTML generated by this package.
func (h *HtmlFile) renderDocument() (renderedDocumentHTML, error) {
	if h.err != nil {
		return nil, h.err
	}
	return renderedDocumentHTML(renderElement(h)), nil
}

// RenderString returns the compact HTML document as a string.
func (h *HtmlFile) RenderString() (string, error) {
	html, err := h.Render()
	if err != nil {
		return "", err
	}
	return string(html), nil
}

// RenderFormatted returns human-readable formatted HTML document bytes.
func (h *HtmlFile) RenderFormatted() ([]byte, error) {
	html, err := h.renderDocument()
	if err != nil {
		return nil, err
	}
	return formatRenderedDocumentHTML(html), nil
}

// RenderFormattedString returns formatted HTML as a string.
func (h *HtmlFile) RenderFormattedString() (string, error) {
	html, err := h.RenderFormatted()
	if err != nil {
		return "", err
	}
	return string(html), nil
}

// mergeHead copies a Head wrapper into the generated document head.
func (h *HtmlFile) mergeHead(head *Head) {
	if head == nil {
		return
	}
	// The wrapper validated its own content; carry any complaint into the
	// document so it surfaces where callers already check for it.
	if head.err != nil {
		h.addError(head.err)
	}
	for k, v := range head.style {
		h.headStyle[k] = v
	}
	h.headContent = append(h.headContent, head.contents...)
}

// mergeBody copies a Body wrapper into the generated document body.
func (h *HtmlFile) mergeBody(body *Body) {
	if body == nil {
		return
	}
	if body.err != nil {
		h.addError(body.err)
	}
	for k, v := range body.style {
		h.bodyStyle[k] = v
	}
	if body.onLoad != "" {
		h.bodyOnLoad = body.onLoad
	}
	if body.onUnload != "" {
		h.bodyOnUnload = body.onUnload
	}
	h.bodyContent = append(h.bodyContent, body.contents...)
}

// rejectStructuralElement rejects nested document structure elements.
func (h *HtmlFile) rejectStructuralElement(target string, e Element) {
	h.addError(fmt.Errorf("cannot add %s to %s element", elementName(e), target))
}

// writeHeadElement writes the generated head section.
func (h *HtmlFile) writeHeadElement(buf *bytes.Buffer) {
	tg := startTag(buf, "head")
	tg.styleAttr(h.headStyle)
	tg.children(h.headContent)
}

// writeBodyElement writes the generated body section.
func (h *HtmlFile) writeBodyElement(buf *bytes.Buffer) {
	tg := startTag(buf, "body")
	tg.attr("onload", h.bodyOnLoad)
	tg.attr("onunload", h.bodyOnUnload)
	tg.styleAttr(h.bodyStyle)
	tg.children(h.bodyContent)
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

// renderedDocumentHTML marks compact HTML produced by HtmlFile.Render.
//
// The formatter below is intentionally scoped to package-rendered documents.
// It is not a public, general-purpose HTML formatter.
type renderedDocumentHTML []byte

// htmlToken stores one formatter token.
type htmlToken struct {
	kind string
	name string
	text string
}

// formatRenderedDocumentHTML returns an indented representation of rephtml output.
//
// This formatter is a small tokenizer over HTML generated by this package. Keep
// it private and use it only for document pretty-printing paths such as
// WriteToFile and RenderFormatted. If this package later needs to format
// arbitrary external HTML, replace this with a real HTML tokenizer first.
func formatRenderedDocumentHTML(src renderedDocumentHTML) []byte {
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
				// The index has to point into input, so the search folds case
				// per byte. strings.ToLower would not do: it is not
				// length-preserving, so a character such as U+212A KELVIN SIGN
				// earlier in the body would shift the result and truncate the
				// closing tag, leaving the element open for the rest of the
				// document.
				closeAt := indexFold(input[rawStart:], closeTag)
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
	// Folded per byte so the index stays valid in text; see tokenizeHTML.
	closeAt := lastIndexFold(text, closeTag)
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
		return hasPrefixFold(input[idx:], "<!doctype")
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
	contentNode[*Head]
	err error
}

// NewHead creates a new Head element
func NewHead() *Head {
	v := &Head{}
	v.init(v)
	return v
}

// Add appends content to the head, recording an error for anything that does
// not belong there.
//
// This is the same check HtmlFile.AddToHead applies. Without it a wrapper is a
// way around that check, since AddToHead flattens a Head rather than inspecting
// what it holds. The error travels with the wrapper and is reported by the
// document when the wrapper is added to it.
func (h *Head) Add(e Element) *Head {
	if isNilElement(e) {
		return h
	}
	if _, ok := e.(HeadElement); !ok {
		h.err = errors.Join(h.err, fmt.Errorf("cannot add %s to head element", elementName(e)))
		return h
	}
	h.contents = appendElement(h.contents, e)
	return h
}

// Err returns content errors recorded while building the head.
func (h *Head) Err() error {
	return h.err
}

// renderTo writes the head element's HTML to buf.
func (h *Head) renderTo(buf *bytes.Buffer) {
	tg := startTag(buf, "head")
	tg.styleAttr(h.style)
	tg.children(h.contents)
}

// Body represents the HTML body element for document content
type Body struct {
	bodyElement
	contentNode[*Body]
	onLoad   string
	onUnload string
	err      error
}

// NewBody creates a new Body element
func NewBody() *Body {
	v := &Body{}
	v.init(v)
	return v
}

// Add appends content to the body, recording an error for anything that does
// not belong there.
//
// This mirrors Head.Add: AddToBody flattens a Body wrapper rather than
// inspecting its contents, so the wrapper has to apply the same check itself.
func (b *Body) Add(e Element) *Body {
	if isNilElement(e) {
		return b
	}
	if _, ok := e.(BodyElement); !ok {
		b.err = errors.Join(b.err, fmt.Errorf("cannot add %s to body element", elementName(e)))
		return b
	}
	b.contents = appendElement(b.contents, e)
	return b
}

// Err returns content errors recorded while building the body.
func (b *Body) Err() error {
	return b.err
}

// renderTo writes the body element's HTML to buf.
func (b *Body) renderTo(buf *bytes.Buffer) {
	tg := startTag(buf, "body")
	tg.attr("onload", b.onLoad)
	tg.attr("onunload", b.onUnload)
	tg.styleAttr(b.style)
	tg.children(b.contents)
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

// Title represents the HTML title element for document title
type Title struct {
	headElement
	contentNode[*Title]
}

// NewTitle creates a new Title element
func NewTitle() *Title {
	v := &Title{}
	v.init(v)
	return v
}

// Prepare builds the HTML for the title element
func (t *Title) renderTo(buf *bytes.Buffer) {
	tg := startTag(buf, "title")
	tg.styleAttr(t.style)
	tg.children(t.contents)
}

// Text adds text content to the title element
func (t *Title) Text(text string) *Title {
	t.contents = appendElement(t.contents, escapedText(text))
	return t
}

// Base represents the HTML base element for document base URL
type Base struct {
	headElement
	node[*Base]
	href   string
	target string
}

// NewBase creates a new Base element
func NewBase() *Base {
	v := &Base{}
	v.init(v)
	return v
}

// Prepare builds the HTML for the base element
func (b *Base) renderTo(buf *bytes.Buffer) {
	tg := startTag(buf, "base")
	tg.urlAttr("href", b.href)
	tg.attr("target", b.target)
	tg.styleAttr(b.style)
	tg.void()
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

// Link represents the HTML link element for external resources
type Link struct {
	headElement
	node[*Link]
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
	v := &Link{}
	v.init(v)
	return v
}

// Prepare builds the HTML for the link element
func (l *Link) renderTo(buf *bytes.Buffer) {
	tg := startTag(buf, "link")
	tg.attr("rel", l.rel)
	tg.urlAttr("href", l.href)
	tg.attr("type", l.linkType)
	tg.attr("media", l.media)
	tg.attr("sizes", l.sizes)
	tg.attr("crossOrigin", l.crossOrigin)
	tg.attr("integrity", l.integrity)
	tg.attr("referrerPolicy", l.referrerPolicy)
	tg.attr("hreflang", l.hreflang)
	tg.styleAttr(l.style)
	tg.void()
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

// Meta represents the HTML meta element for metadata
type Meta struct {
	headElement
	node[*Meta]
	name      string
	content   string
	charset   string
	property  string
	httpEquiv string
	scheme    string
}

// NewMeta creates a new Meta element
func NewMeta() *Meta {
	v := &Meta{}
	v.init(v)
	return v
}

// Prepare builds the HTML for the meta element
func (m *Meta) renderTo(buf *bytes.Buffer) {
	tg := startTag(buf, "meta")
	tg.attr("name", m.name)
	tg.attr("content", m.content)
	tg.attr("charset", m.charset)
	tg.attr("property", m.property)
	tg.attr("http-equiv", m.httpEquiv)
	tg.attr("scheme", m.scheme)
	tg.styleAttr(m.style)
	tg.void()
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

// StyleElement represents the HTML style element for CSS styles
type StyleElement struct {
	headElement
	bodyElement
	contentNode[*StyleElement]
	styleType string
	media     string
}

// NewStyleElement creates a new StyleElement element
func NewStyleElement() *StyleElement {
	v := &StyleElement{}
	v.init(v)
	return v
}

// Prepare builds the HTML for the style element
func (s *StyleElement) renderTo(buf *bytes.Buffer) {
	tg := startTag(buf, "style")
	tg.attr("type", s.styleType)
	tg.attr("media", s.media)
	tg.styleAttr(s.style)
	tg.children(s.contents)
}

// Add adds content to the style element
func (s *StyleElement) Add(e Element) *StyleElement {
	switch element := e.(type) {
	case nil:
		return s
	case CSSRule:
		return s.AddRule(element)
	case *Style:
		if element == nil {
			return s
		}
		return s.AddRule(&StyleRule{
			Props: cloneStyleMap(element.Props),
			Tags:  append([]string(nil), element.Tags...),
		})
	default:
		s.contents = appendElement(s.contents, e)
	}
	return s
}

// AddRule appends typed CSS rule content to the style element.
func (s *StyleElement) AddRule(rule CSSRule) *StyleElement {
	if !isNilElement(rule) {
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
