package rephtml

import (
	"bytes"
	"log"
	"os"
	"strings"
)

const indent = "\n\t"
const newline = "\n"
const tab = "\t"

type HtmlFile struct {
	buf         bytes.Buffer
	head        []byte
	style, body [][]byte
	ttrack      int // tab tracker
}

func NewHtmlFile() *HtmlFile {
	return &HtmlFile{
		ttrack: 1,
	}
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
Add a paragraph element to the HTML document with
a string parameter as the assigned value
*/
func (h *HtmlFile) PString(s string) *HtmlFile {
	h.body = append(h.body, []byte("<p>"+s+"</p>"))
	return h
}

/*
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
	h.buf.WriteString("<html>" + newline)
	h.buf.WriteString(t + "<header>" + newline)
	h.ttrack++
	t = tabs(h.ttrack)
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
	t = tabs(h.ttrack)
	h.buf.WriteString(t + "</header>" + newline)
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
		defer file.Close()
		if _, err := file.Write(h.buf.Bytes()); err != nil {
			log.Fatal(err)
		}
	} else {
		file, err := os.Create(path)
		if err != nil {
			log.Fatal(err)
		}
		defer file.Close()
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

	// get indexes for open an close braces
	open, close := bytes.Index(nsb, []byte("{"))+1, len(nsb)-1

	// cut array into parts: opening, contents, and closing
	openb, closeb := nsb[:open], nsb[close]
	contents := nsb[open:close]

	// add spacing between open values
	openbTmp := []byte{}
	for i := 0; i < len(openb); i++ {
		if openb[i] == ',' {
			openbTmp = append(openbTmp, openb[i])
			openbTmp = append(openbTmp, ' ')
		} else if openb[i] == '{' {
			openbTmp = append(openbTmp, ' ')
			openbTmp = append(openbTmp, openb[i])
		} else {
			openbTmp = append(openbTmp, openb[i])
		}
	}
	openbTmp = append(openbTmp, '\n')
	openb = openbTmp

	// write openb to buffer
	fb.WriteString(tabs(h.ttrack))
	fb.Write(openb)

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
	fb.WriteByte(closeb)
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
