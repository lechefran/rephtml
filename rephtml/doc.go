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

// vv struct functions vv

func (h *HtmlFile) Style(s *StyleTag) *HtmlFile {
	h.style = append(h.style, s.Bytes())
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

func (h *HtmlFile) parseStyleBytes(b []byte) []byte {
	var fb bytes.Buffer
	open, close := b[0], b[len(b)-1]
	contents := bytes.Fields(b[open+1 : close])
	t := tabs(h.ttrack)
	for _, c := range contents {
		fb.WriteString(t)
		fb.Write(c)
		fb.WriteString(";" + newline)
	}
	return fb.Bytes()
}

func (h *HtmlFile) Prepare() *HtmlFile {
	t := tabs(h.ttrack + 1)
	h.buf.WriteString("<html>" + newline)
	h.buf.WriteString(t + "<header>" + newline)
	t = tabs(h.ttrack + 1)
	h.buf.WriteString(t + "<style>" + newline)
	t = tabs(h.ttrack + 1)
	// for _, s := range h.style {
	// 	h.parseStyleBytes(s)
	// }
	h.buf.WriteString(indent + "</style>")
	h.buf.WriteString(indent + "</header>")
	h.buf.WriteString(indent + "<body>")
	for _, s := range h.body {
		h.buf.WriteString(indent + tab)
		h.buf.Write(s)
	}
	h.buf.WriteString(indent + "</body>" + newline)
	h.buf.WriteString("</html>")
	return h
}

// change of plans- write to byte arr first, parse later
func (h *HtmlFile) StyleString(s string) *HtmlFile {
	// s = strings.ReplaceAll(s, "\t", "")
	// s = strings.ReplaceAll(s, ";", "; ")
	// open, close := strings.Index(s, "{"), strings.Index(s, "}")
	// props := s[open+1 : close-1]
	// fs := s[:open+1]
	// propsarr := strings.Fields(props)
	// for _, p := range propsarr {
	// 	fs += p
	// }
	// fs += s[close:]
	fs := strings.ReplaceAll(s, ";", "; ")
	h.style = append(h.style, []byte(fs))
	return h
}

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

func (h *HtmlFile) WriteToFile(s string) {
	if len(h.head) == 0 && len(h.body) == 0 && len(h.style) == 0 {
		log.Print("No values were appended to the HTML File. " + s + " will not be created")
	}
	if _, err := os.Stat(s); os.IsExist(err) {
		if err := os.Remove(s); err != nil {
			log.Fatal(err)
		}
		file, err := os.Create(s)
		if err != nil {
			log.Fatal(err)
		}
		defer file.Close()
		if _, err := file.Write(h.buf.Bytes()); err != nil {
			log.Fatal(err)
		}
	} else {
		file, err := os.Create(s)
		if err != nil {
			log.Fatal(err)
		}
		defer file.Close()
		if _, err := file.Write(h.buf.Bytes()); err != nil {
			log.Fatal(err)
		}
	}
}
