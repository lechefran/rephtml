package rephtml

import (
	"bytes"
	"html"
	"regexp"
	"strconv"
)

/*
Internal parsing function to handle style
*/
func parseStyle(buf *bytes.Buffer, style map[string]string) {
	idx := 0
	buf.WriteString(" style=\"")
	for k, v := range style {
		buf.WriteString(escapeAttr(k) + ": " + escapeAttr(v) + ";")
		if idx != len(style)-1 {
			buf.WriteByte(' ')
		}
		idx++
	}
	buf.WriteString("\"")
}

/*
Internal parsing function to remove all spaces from a byte array
*/
func strip(b []byte) []byte {
	re := regexp.MustCompile(`\s+`)
	return re.ReplaceAll(b, nil)
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

func cloneBytes(b []byte) []byte {
	return append([]byte(nil), b...)
}

func escapeText(text string) string {
	return html.EscapeString(text)
}

func escapeAttr(value string) string {
	return html.EscapeString(value)
}

func writeAttr(buf *bytes.Buffer, name, value string) {
	buf.WriteByte(' ')
	buf.WriteString(name)
	buf.WriteString("=\"")
	buf.WriteString(escapeAttr(value))
	buf.WriteByte('"')
}

func writeIntAttr(buf *bytes.Buffer, name string, value int) {
	writeAttr(buf, name, strconv.Itoa(value))
}

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

type rawText string

func (r rawText) Bytes() []byte {
	return []byte(r)
}

func (r rawText) Prepare() {}

type escapedText string

func (e escapedText) Bytes() []byte {
	return []byte(escapeText(string(e)))
}

func (e escapedText) Prepare() {}

func appendElement(contents []Element, e Element) []Element {
	if e == nil {
		return contents
	}
	return append(contents, e)
}

func writeElement(buf *bytes.Buffer, e Element) {
	if e == nil {
		return
	}
	e.Prepare()
	buf.Write(e.Bytes())
}

func writeElements(buf *bytes.Buffer, contents []Element) {
	for _, e := range contents {
		writeElement(buf, e)
	}
}
