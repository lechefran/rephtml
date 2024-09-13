package rephtml

import (
	"bytes"
)

type Div struct {
	buf      bytes.Buffer
	contents [][]byte
	ttrack   int
	style    map[string]string
}

func NewDiv() *Div {
	return &Div{
		style: make(map[string]string),
	}
}

func (d *Div) Tabs(i int) *Div {
	d.ttrack = i
	return d
}

func (d *Div) Add(e Elements) *Div {
	d.contents = append(d.contents, e.Bytes())
	return d
}

func (d *Div) AddStyles(m map[string]string) *Div {
	for k, v := range m {
		d.style[k] = v
	}
	return d
}

func (d *Div) Bytes() []byte {
	return d.buf.Bytes()
}

func (d *Div) Prepare() {
	if len(d.style) != 0 {
		idx := 0
		d.buf.WriteString("<div style=\"")
		for k, v := range d.style {
			d.buf.WriteString(k + ": " + v + ";")
			if idx != len(d.style)-1 {
				d.buf.WriteByte(' ')
			}
			idx++
		}
		d.buf.WriteString("\">")
	} else {
		d.buf.WriteString("<div>")
	}

	for _, c := range d.contents {
		// if bytes.Contains(c, []byte("<div")) && bytes.Contains(c, []byte(">")) {
		// 	d.buf.WriteByte('\n')
		// }
		if bytes.Contains(c, []byte("<table")) && bytes.Contains(c, []byte(">")) {
			d.buf.Write(d.formatTable(c))
		} else {
			d.buf.Write(c)
		}
	}
	d.buf.WriteString("</div>")
}

func (d *Div) formatDiv(b []byte) []byte {
	var fb bytes.Buffer

	// split byte array by tags
	curr := []byte{}
	sarr := [][]byte{}
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
			fb.WriteString(tabs(d.ttrack))
			fb.Write(s)
			d.ttrack++
		} else if bytes.Equal(s, []byte("</div>")) {
			d.ttrack--
			fb.WriteByte('\n')
			fb.WriteString(tabs(d.ttrack))
			fb.Write(s)
		} else {
			fb.WriteByte('\n')
			fb.WriteString(tabs(d.ttrack))
			fb.Write(s)
		}
	}
	fb.WriteByte('\n')
	return fb.Bytes()
}

func (d *Div) formatTable(b []byte) []byte {
	var fb bytes.Buffer

	// see if open table tag has id and class values
	var tmp bytes.Buffer
	opent := []byte{}
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
	nsbsplit := []byte{}
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
	fb.Write(opent)
	fb.WriteByte('\n')
	d.ttrack += 4

	// loop through contents
	for _, c := range contents {
		if bytes.Equal(c, []byte("<tr>")) {
			fb.WriteString(tabs(d.ttrack))
			fb.Write(c)
			fb.WriteByte('\n')
			d.ttrack++
		} else if bytes.Equal(c, []byte("</tr>")) {
			d.ttrack--
			fb.WriteString(tabs(d.ttrack))
			fb.Write(c)
			fb.WriteByte('\n')
		} else {
			fb.WriteString(tabs(d.ttrack))
			fb.Write(c)
			fb.WriteByte('\n')
		}
	}

	d.ttrack--
	fb.WriteString(tabs(d.ttrack))
	fb.WriteString("</table>")
	fb.WriteByte('\n')
	return fb.Bytes()
}
