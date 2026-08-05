package rephtml

import (
	"bytes"
	"html"
	"strconv"
	"strings"
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

// renderPrepared prepares an element and returns a snapshot of its rendered bytes.
func renderPrepared(e preparedElement) []byte {
	if e == nil {
		return nil
	}
	e.prepare()
	return cloneBytes(e.rawBytes())
}

// htmlPrepared prepares an element and returns its rendered HTML.
func htmlPrepared(e preparedElement) string {
	if e == nil {
		return ""
	}
	e.prepare()
	return string(e.rawBytes())
}

// escapeText escapes ordinary HTML text node content.
func escapeText(text string) string {
	return html.EscapeString(text)
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

// openTag resets buf and begins an element start tag. Element Prepare methods
// start here, which is also what guarantees the buffer is cleared before a
// re-render.
func openTag(buf *bytes.Buffer, name string) tag {
	buf.Reset()
	return appendTag(buf, name)
}

// appendTag begins an element start tag without resetting buf, for elements
// written into a buffer that already holds content.
func appendTag(buf *bytes.Buffer, name string) tag {
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

// prepare is a no-op: raw text is already its own rendered form.
func (r rawText) prepare() {}

// rawBytes returns the raw text bytes.
func (r rawText) rawBytes() []byte {
	return []byte(r)
}

// Render returns raw text bytes.
func (r rawText) Render() []byte {
	return renderPrepared(r)
}

// HTML returns raw text as a string.
func (r rawText) HTML() string {
	return htmlPrepared(r)
}

// escapedText stores content that should be escaped before rendering.
type escapedText string

// prepare is a no-op: escaping happens when the bytes are read.
func (e escapedText) prepare() {}

// rawBytes returns the escaped text bytes.
func (e escapedText) rawBytes() []byte {
	return []byte(escapeText(string(e)))
}

// Render returns escaped text bytes.
func (e escapedText) Render() []byte {
	return renderPrepared(e)
}

// HTML returns escaped text as a string.
func (e escapedText) HTML() string {
	return htmlPrepared(e)
}

// appendElement appends a non-nil child element to a content slice.
func appendElement(contents []Element, e Element) []Element {
	if e == nil {
		return contents
	}
	return append(contents, e)
}

// writeElement renders and writes one element to the destination buffer.
//
// Elements in this package expose an unexported prepared contract, which lets
// a parent copy a child's bytes straight out of the child's buffer. Render
// would owe the caller a defensive copy that is thrown away immediately, so
// taking that path saves a full copy of the subtree at every level of nesting.
// Element implementations from outside the package fall back to Render.
func writeElement(buf *bytes.Buffer, e Element) {
	if e == nil {
		return
	}
	if p, ok := e.(preparedElement); ok {
		p.prepare()
		buf.Write(p.rawBytes())
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
