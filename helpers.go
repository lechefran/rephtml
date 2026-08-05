package rephtml

import (
	"bytes"
	"html"
	"reflect"
	"strconv"
	"strings"
	"sync"
)

// parseStyle writes an inline style attribute from a StyleMap.
//
// Declarations are emitted in sorted key order, since Go map iteration order is
// unspecified and rendered output has to be stable.
func parseStyle(buf *bytes.Buffer, style StyleMap) {
	buf.WriteString(" style=\"")

	keys := sortedStyleKeys(style)
	for idx, k := range keys {
		buf.WriteString(escapeAttr(k))
		buf.WriteString(": ")
		buf.WriteString(escapeAttr(style[k]))
		buf.WriteByte(';')
		if idx != len(keys)-1 {
			buf.WriteByte(' ')
		}
	}
	buf.WriteString("\"")
}

// indentCache holds pre-built indentation strings for the common nesting
// depths, so formatting a document does not rebuild them once per output line.
var indentCache = func() [33]string {
	var c [33]string
	for i := range c {
		c[i] = strings.Repeat(tab, i)
	}
	return c
}()

// tabs returns an indentation string of t tabs.
func tabs(t int) string {
	if t < 0 {
		return ""
	}
	if t < len(indentCache) {
		return indentCache[t]
	}
	return strings.Repeat(tab, t)
}

// cloneBytes returns a copy of rendered bytes so callers cannot mutate buffers.
func cloneBytes(b []byte) []byte {
	return append([]byte(nil), b...)
}

// cloneStyleMap returns a copy of CSS declarations so callers can reuse maps safely.
func cloneStyleMap(style StyleMap) StyleMap {
	clone := make(StyleMap, len(style))
	for k, v := range style {
		clone[k] = v
	}
	return clone
}

// cloneStrings returns a copy of a string slice so callers can reuse slices safely.
func cloneStrings(values []string) []string {
	return append([]string(nil), values...)
}

// bufferPool supplies the scratch buffers rendering writes into.
//
// Elements no longer keep a buffer of their own, so without pooling every
// render would grow a new one from nothing. A pooled buffer comes back already
// sized for the last document of similar shape, which is where the growth
// allocations went. Each render holds its buffer exclusively, so this stays
// safe to use from several goroutines.
var bufferPool = sync.Pool{
	New: func() any { return new(bytes.Buffer) },
}

// maxPooledBuffer caps what goes back into the pool, so one very large document
// does not pin its buffer in memory for the life of the process.
const maxPooledBuffer = 1 << 20

// getBuffer takes a reset buffer from the pool.
func getBuffer() *bytes.Buffer {
	buf := bufferPool.Get().(*bytes.Buffer)
	buf.Reset()
	return buf
}

// putBuffer returns a buffer to the pool unless it has grown unreasonably.
func putBuffer(buf *bytes.Buffer) {
	if buf.Cap() <= maxPooledBuffer {
		bufferPool.Put(buf)
	}
}

// renderElement renders an element and returns its bytes.
// It backs the Render method of the types that do not embed the generic base.
func renderElement(e preparedElement) []byte {
	if e == nil {
		return nil
	}
	buf := getBuffer()
	defer putBuffer(buf)
	e.renderTo(buf)
	// Copied out at exactly the right size, so the caller owns the result and
	// the buffer can be reused.
	return cloneBytes(buf.Bytes())
}

// htmlElement renders an element and returns its HTML.
func htmlElement(e preparedElement) string {
	if e == nil {
		return ""
	}
	buf := getBuffer()
	defer putBuffer(buf)
	e.renderTo(buf)
	return buf.String()
}

// asciiLower folds one ASCII letter to lower case.
func asciiLower(c byte) byte {
	if c >= 'A' && c <= 'Z' {
		return c + ('a' - 'A')
	}
	return c
}

// hasPrefixFold reports whether s begins with prefix, comparing ASCII letters
// case-insensitively.
//
// It compares byte by byte rather than folding whole strings, so a match is
// always exactly len(prefix) bytes of s. strings.ToLower cannot be used for
// this: it is not length-preserving, so an index into the folded string does
// not point at the same place in the original.
func hasPrefixFold(s, prefix string) bool {
	if len(s) < len(prefix) {
		return false
	}
	for i := 0; i < len(prefix); i++ {
		if asciiLower(s[i]) != asciiLower(prefix[i]) {
			return false
		}
	}
	return true
}

// indexFold returns the byte index in s of the first occurrence of substr,
// comparing ASCII letters case-insensitively, or -1 when there is none.
func indexFold(s, substr string) int {
	if substr == "" {
		return 0
	}
	for i := 0; i+len(substr) <= len(s); i++ {
		if hasPrefixFold(s[i:], substr) {
			return i
		}
	}
	return -1
}

// lastIndexFold returns the byte index in s of the last occurrence of substr,
// comparing ASCII letters case-insensitively, or -1 when there is none.
func lastIndexFold(s, substr string) int {
	if substr == "" {
		return len(s)
	}
	for i := len(s) - len(substr); i >= 0; i-- {
		if hasPrefixFold(s[i:], substr) {
			return i
		}
	}
	return -1
}

// escapeText escapes ordinary HTML text node content.
func escapeText(text string) string {
	return html.EscapeString(text)
}

// escapeCSS neutralises text that would end the style element containing it.
//
// A style element holds raw text: the HTML tokenizer does not decode character
// references there and ends the element at the first "</style". Entity escaping
// is therefore useless, and the one sequence that has to be defused is that
// one. Rewriting its "<" as the CSS character escape \3c keeps the declaration
// meaning identical while leaving nothing for the tokenizer to match.
//
// Only that sequence is touched, so a "<" used legitimately, as in the media
// query range syntax "(400px <= width)", survives untouched.
func escapeCSS(css string) string {
	const endTag = "</style"
	if indexFold(css, endTag) < 0 {
		return css
	}

	var b strings.Builder
	b.Grow(len(css) + 8)
	for i := 0; i < len(css); i++ {
		if css[i] == '<' && hasPrefixFold(css[i:], endTag) {
			// The trailing space terminates the hex escape.
			b.WriteString(`\3c `)
			continue
		}
		b.WriteByte(css[i])
	}
	return b.String()
}

// escapeAttr escapes HTML attribute values.
func escapeAttr(value string) string {
	return html.EscapeString(value)
}

// writeAttr writes one escaped name/value HTML attribute.
func writeAttr(buf *bytes.Buffer, name, value string) {
	buf.WriteByte(' ')
	buf.WriteString(name)
	buf.WriteString("=\"")
	buf.WriteString(escapeAttr(value))
	buf.WriteByte('"')
}

// writeIntAttr writes one integer HTML attribute.
func writeIntAttr(buf *bytes.Buffer, name string, value int) {
	writeAttr(buf, name, strconv.Itoa(value))
}

// writeClassAttr writes a class attribute from class names.
func writeClassAttr(buf *bytes.Buffer, classes []string) {
	buf.WriteString(" class=\"")
	for i, class := range classes {
		buf.WriteString(escapeAttr(class))
		if i != len(classes)-1 {
			buf.WriteByte(' ')
		}
	}
	buf.WriteByte('"')
}

// tag writes a single HTML element into a buffer.
//
// Element rendering is otherwise the same three lines repeated per attribute:
// test whether the field is set, then write it. tag folds that test into each
// write so a prepare method reads as a description of the element's output.
// Attributes are emitted in call order, so each element keeps control of its
// own attribute layout, including where the style attribute sits.
type tag struct {
	buf  *bytes.Buffer
	name string
}

// startTag begins an element start tag. It appends, because the buffer it is
// given usually already holds the element's ancestors and preceding siblings.
func startTag(buf *bytes.Buffer, name string) tag {
	buf.WriteByte('<')
	buf.WriteString(name)
	return tag{buf: buf, name: name}
}

// attr writes a name/value attribute, omitting it when the value is empty.
func (t tag) attr(name, value string) {
	if value != "" {
		writeAttr(t.buf, name, value)
	}
}

// urlAttr writes a URL attribute, replacing a value whose scheme could execute
// script. Escaping alone would not help here: a javascript: URL needs no
// special characters to do its work.
func (t tag) urlAttr(name, value string) {
	if value != "" {
		writeAttr(t.buf, name, safeURLValue(value))
	}
}

// srcsetAttr writes a srcset attribute, filtering each candidate URL in the list.
func (t tag) srcsetAttr(name, value string) {
	if value != "" {
		writeAttr(t.buf, name, filterSrcset(value))
	}
}

// boolAttr writes a valueless boolean attribute when it is set.
func (t tag) boolAttr(name string, on bool) {
	if on {
		t.buf.WriteByte(' ')
		t.buf.WriteString(name)
	}
}

// intAttr writes a numeric attribute, omitting it when the value is not positive.
func (t tag) intAttr(name string, value int) {
	if value > 0 {
		writeIntAttr(t.buf, name, value)
	}
}

// classAttr writes a class attribute, omitting it when there are no classes.
func (t tag) classAttr(classes []string) {
	if len(classes) != 0 {
		writeClassAttr(t.buf, classes)
	}
}

// styleAttr writes an inline style attribute, omitting it when there are no
// declarations.
func (t tag) styleAttr(style StyleMap) {
	if len(style) != 0 {
		parseStyle(t.buf, style)
	}
}

// open finishes the start tag, leaving the element open for custom content.
func (t tag) open() {
	t.buf.WriteByte('>')
}

// end writes the element's closing tag.
func (t tag) end() {
	t.buf.WriteString("</")
	t.buf.WriteString(t.name)
	t.buf.WriteByte('>')
}

// void finishes a void element, which has no content and no closing tag.
func (t tag) void() {
	t.open()
}

// empty closes an element that carries no content.
func (t tag) empty() {
	t.open()
	t.end()
}

// text closes an element around escaped text content.
func (t tag) text(s string) {
	t.open()
	t.buf.WriteString(escapeText(s))
	t.end()
}

// raw closes an element around unescaped content, for raw-text elements such
// as script whose bodies are not parsed as HTML.
func (t tag) raw(s string) {
	t.open()
	t.buf.WriteString(s)
	t.end()
}

// children closes an element around its rendered child elements.
func (t tag) children(contents []Element) {
	t.open()
	writeElements(t.buf, contents)
	t.end()
}

// rawText stores content that should be rendered without escaping.
type rawText string

// renderTo writes the raw text to buf.
func (r rawText) renderTo(buf *bytes.Buffer) {
	buf.WriteString(string(r))
}

// Render returns raw text bytes.
func (r rawText) Render() []byte {
	return renderElement(r)
}

// HTML returns raw text as a string.
func (r rawText) HTML() string {
	return htmlElement(r)
}

// escapedText stores content that should be escaped before rendering.
type escapedText string

// renderTo writes the escaped text to buf.
func (e escapedText) renderTo(buf *bytes.Buffer) {
	buf.WriteString(escapeText(string(e)))
}

// Render returns escaped text bytes.
func (e escapedText) Render() []byte {
	return renderElement(e)
}

// HTML returns escaped text as a string.
func (e escapedText) HTML() string {
	return htmlElement(e)
}

// isNilElement reports whether e is nil or holds a nil pointer.
//
// A plain e == nil is not enough. A builder that returns a typed nil, such as
// func() *Div { return nil }, produces an interface with a type but no value,
// which compares unequal to nil and then panics the moment it is rendered.
// Checking the underlying value catches that at the point it is added.
func isNilElement(e Element) bool {
	if e == nil {
		return true
	}
	switch v := reflect.ValueOf(e); v.Kind() {
	case reflect.Chan, reflect.Func, reflect.Interface, reflect.Map, reflect.Pointer, reflect.Slice:
		return v.IsNil()
	default:
		return false
	}
}

// appendElement appends a non-nil child element to a content slice.
func appendElement(contents []Element, e Element) []Element {
	if isNilElement(e) {
		return contents
	}
	return append(contents, e)
}

// writeElement renders and writes one element to the destination buffer.
//
// Elements in this package can write straight into the destination, so a
// subtree is serialised once no matter how deeply it is nested. Element
// implementations from outside the package fall back to Render, which costs
// one intermediate buffer for that element.
func writeElement(buf *bytes.Buffer, e Element) {
	if isNilElement(e) {
		return
	}
	if p, ok := e.(preparedElement); ok {
		p.renderTo(buf)
		return
	}
	buf.Write(e.Render())
}

// writeElements renders and writes all elements to the destination buffer.
func writeElements(buf *bytes.Buffer, contents []Element) {
	for _, e := range contents {
		writeElement(buf, e)
	}
}
