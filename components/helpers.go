package rephtml

import (
	"bytes"
	"regexp"
)

/*
Internal parsing function to handle style
*/
func parseStyle(buf *bytes.Buffer, style map[string]string) {
	idx := 0
	buf.WriteString(" style=\"")
	for k, v := range style {
		buf.WriteString(k + ": " + v + ";")
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
func strip(bytes []byte) []byte {
	re := regexp.MustCompile(`\\s+`)
	return re.ReplaceAll(bytes, nil)
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

type rawText string

func (r rawText) Bytes() []byte {
	return []byte(r)
}

func (r rawText) Prepare() {}

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
