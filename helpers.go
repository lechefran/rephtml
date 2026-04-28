package rephtml

import (
	"bytes"
	"html"
	"slices"
	"strconv"
)

/*
Internal parsing function to handle style
*/
func parseStyle(buf *bytes.Buffer, style StyleMap) {
	buf.WriteString(" style=\"")

	keys := make([]string, 0, len(style))
	for k := range style {
		keys = append(keys, k)
	}
	slices.Sort(keys)

	for idx, k := range keys {
		buf.WriteString(escapeAttr(k) + ": " + escapeAttr(style[k]) + ";")
		if idx != len(keys)-1 {
			buf.WriteByte(' ')
		}
	}
	buf.WriteString("\"")
}

/*
Internal function that creates tabs based on the value of t
*/
func tabs(t int) string {
	res := ""
	for i := 0; i < t; i++ {
		res += tab
	}
	return res
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
	e.Prepare()
	return e.Bytes()
}

// htmlPrepared prepares an element and returns its rendered HTML.
func htmlPrepared(e preparedElement) string {
	return string(renderPrepared(e))
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

// rawText stores content that should be rendered without escaping.
type rawText string

// Bytes returns a defensive copy of the rendered rawText bytes.
func (r rawText) Bytes() []byte {
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

// String returns raw text as a string.
func (r rawText) String() string {
	return r.HTML()
}

// Prepare renders the rawText component into its internal buffer.
func (r rawText) Prepare() {}

// escapedText stores content that should be escaped before rendering.
type escapedText string

// Bytes returns a defensive copy of the rendered escapedText bytes.
func (e escapedText) Bytes() []byte {
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

// String returns escaped text as a string.
func (e escapedText) String() string {
	return e.HTML()
}

// Prepare renders the escapedText component into its internal buffer.
func (e escapedText) Prepare() {}

// appendElement appends a non-nil child element to a content slice.
func appendElement(contents []Element, e Element) []Element {
	if e == nil {
		return contents
	}
	return append(contents, e)
}

// writeElement renders and writes one element to the destination buffer.
func writeElement(buf *bytes.Buffer, e Element) {
	if e == nil {
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
