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
	re := regexp.MustCompile("\\s+")
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
