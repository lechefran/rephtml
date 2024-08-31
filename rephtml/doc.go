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

/*
Internal parsing function to format style attributes
*/
func (h *HtmlFile) parseStyleBytes(b []byte) []byte {
	var fb bytes.Buffer
	open, close := bytes.Index(b, []byte("{"))+1, len(b)-1
	contents := bytes.Fields(b[open:close])
	t := tabs(h.ttrack)
	fb.WriteString(t)
	fb.Write(b[:open])
	fb.WriteString(newline)
	h.ttrack++
	t = tabs(h.ttrack)
	for _, c := range contents {
		fb.WriteString(t)
		fb.Write(c)
		fb.WriteString(newline)
	}
	h.ttrack--
	t = tabs(h.ttrack)
	fb.WriteString(t + "}" + newline)
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
			h.buf.Write(h.parseStyleBytes(h.style[i]))
			h.buf.WriteString(newline)
		} else {
			h.buf.Write(h.parseStyleBytes(h.style[i]))
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
			h.buf.WriteString(t)
			h.buf.Write(h.body[i])
			h.buf.WriteString(newline)
		} else {
			h.buf.WriteString(t)
			h.buf.Write(h.body[i])
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
