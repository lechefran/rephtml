package rephtml

import (
	"bytes"
	"log"
	"os"
	"regexp"
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

// vv struct functions vv

func (h *HtmlFile) Style(s *Style) *HtmlFile {
	h.style = append(h.style, s.Bytes())
	return h
}

func (h *HtmlFile) Table(t *Table) *HtmlFile {
	h.body = append(h.body, t.Bytes())
	return h
}

// vv string functions vv

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

func (h *HtmlFile) PStringWithStyle(s, style string) *HtmlFile {
	if style == "" {
		h.body = append(h.body, []byte("<p>"+s+"</p>"))
	} else {
		h.body = append(h.body, []byte("<p style=\""+style+"\";>"+s+"</p>"))
	}
	return h
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

	// split con
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
	nsb := strip(b) // remove all spaces
	nsbsplit := []byte{}

	// split byte array by tags
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

	// obtain the open and close table tags, and contents
	opent, closet := sarr[0], sarr[len(sarr)-1]
	contents := sarr[1 : len(sarr)-2]

	// write open tag to buffer
	fb.Write(opent)
	fb.WriteByte('\n')

	// loop through contents
	var curr bytes.Buffer
	for _, c := range contents {
		if bytes.Equal(c, []byte("<tr>")) {
			curr.WriteByte('\t')
			curr.Write(c)
			curr.WriteByte('\n')
		} else if bytes.Equal(c, []byte("</tr>")) {
			curr.WriteByte('\t')
			curr.Write(c)
			curr.WriteByte('\n')
			fb.Write(curr.Bytes())
			curr.Reset()
		} else {
			curr.WriteByte('\t')
			curr.Write(c)
			curr.WriteByte('\n')
		}
	}

	// write close tag to header
	fb.Write(closet)
	fb.WriteByte('\n')
	return fb.Bytes()
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
		if i != len(h.body)-1 {
			bbytes := h.body[i]
			if bytes.Contains(bbytes, []byte("table")) {
				h.buf.Write(h.formatTable(bbytes))
			} else {
				h.buf.WriteString(t)
				h.buf.Write(bbytes)
				h.buf.WriteString(newline)
			}
		} else {
			bbytes := h.body[i]
			if bytes.Contains(bbytes, []byte("table")) {
				h.buf.Write(h.formatTable(bbytes))
			} else {
				h.buf.WriteString(t)
				h.buf.Write(bbytes)
			}
		}
	}
	h.ttrack--
	t = tabs(h.ttrack)
	h.buf.WriteString(newline + t + "</body>" + newline)
	h.buf.WriteString("</html>")
	return h
}

// change of plans- write to byte arr first, parse later
func (h *HtmlFile) StyleString(s string) *HtmlFile {
	fs := strings.ReplaceAll(s, ";", "; ")
	h.style = append(h.style, []byte(fs))
	return h
}

// todo: fix to write to byte arr first, parse later
func (h *HtmlFile) TableString(harr []string, darr [][]string) *HtmlFile {
	t := indent + tab
	tbl := "<table>"
	t += tab
	tbl += t + "<tr>"
	t += tab
	for _, h := range harr {
		tbl += t + "<th>" + h + "</th>"
	}
	t = t[:strings.LastIndex(t, tab)]
	tbl += t + "</tr>"
	for _, d := range darr {
		tbl += t + "<tr>"
		t += tab
		for _, d1 := range d {
			tbl += t + "<td>" + d1 + "</td>"
		}
		t = t[:strings.LastIndex(t, tab)]
		tbl += t + "</tr>"
	}
	t = t[:strings.LastIndex(t, tab)]
	tbl += t + "</table>"
	h.body = append(h.body, []byte(tbl))
	return h
}

func strip(bytes []byte) []byte {
	re := regexp.MustCompile("\\s+")
	return re.ReplaceAll(bytes, nil)
}

/*
Helper method that creates tabs based on the value of ttrack
*/
func tabs(t int) string {
	res := ""
	for i := 0; i < t; i++ {
		res += tab
	}
	return res
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
