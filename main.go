package main

import (
	"bytes"
	"log"
	"os"
	"strings"
)

const indent = "\n\t"
const newspace = "\n"
const tab = "\t"

// write custom html interface
// write css styling attributes struct (settings)

type HtmlDiv struct {
	buf     bytes.Buffer
	content [][]string
}

func NewDiv() *HtmlDiv {
	return &HtmlDiv{}
}

func (d *HtmlDiv) ReadBuffer(b *bytes.Buffer) {
	b.Read(d.buf.Bytes())
}

type HtmlFile struct {
	buf         bytes.Buffer
	head        []byte
	style, body [][]byte
}

func NewHtmlFile() *HtmlFile {
	return &HtmlFile{}
}

func (h *HtmlFile) Body(s string) *HtmlFile {
	// h.body = append(h.body, s)
	return h
}

func (h *HtmlFile) Div() *HtmlFile {
	return h
}

func (h *HtmlFile) H1(s string) *HtmlFile {
	h.body = append(h.body, []byte("<h1>"+s+"</h1>"))
	return h
}

func (h *HtmlFile) H2(s string) *HtmlFile {
	h.body = append(h.body, []byte("<h2>"+s+"</h2>"))
	return h
}

func (h *HtmlFile) H3(s string) *HtmlFile {
	h.body = append(h.body, []byte("<h3>"+s+"</h3>"))
	return h
}

func (h *HtmlFile) H4(s string) *HtmlFile {
	h.body = append(h.body, []byte("<h4>"+s+"</h4>"))
	return h
}

func (h *HtmlFile) H5(s string) *HtmlFile {
	h.body = append(h.body, []byte("<h5>"+s+"</h5>"))
	return h
}

func (h *HtmlFile) H6(s string) *HtmlFile {
	h.body = append(h.body, []byte("<h6>"+s+"</h6>"))
	return h
}

func (h *HtmlFile) P(s string) *HtmlFile {
	h.body = append(h.body, []byte("<p>"+s+"</p>"))
	return h
}

func (h *HtmlFile) Prepare() *HtmlFile {
	h.buf.WriteString("<html>")
	h.buf.WriteString(indent + "<header>")
	h.buf.WriteString(indent + "<style>")
	for _, s := range h.style {
		h.buf.Write([]byte(indent))
		h.buf.Write([]byte(tab))
		h.buf.Write(s)
	}
	h.buf.WriteString(indent + "</style>")
	h.buf.WriteString(indent + "</header>")
	h.buf.WriteString(indent + "<body>")
	// can definitely be improved
	for _, s := range h.body {
		h.buf.Write([]byte(indent))
		h.buf.Write([]byte(tab))
		h.buf.Write(s)
	}
	h.buf.WriteString(indent + "</body>" + newspace)
	h.buf.WriteString("</html>")
	return h
}

func (h *HtmlFile) Style(s string) *HtmlFile {
	h.style = append(h.style, []byte(s))
	return h
}

func (h *HtmlFile) Table(harr []string, darr [][]string) *HtmlFile {
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

func (h *HtmlFile) WriteToFile(s string) {
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

// write a table struct

func main() {
	html := NewHtmlFile()
	html.Style(`h1, h2, h3, h4, h5, h6, p {
			font-family: Arial;
			text-align: center;
		}`)
	html.Style(`th, td {
			font-family: Arial;
			padding: 10px;
			text-align: center;
		}`)
	html.Style(`table {
		margin-left: auto;
		margin-right: auto;
		width: 80%;
		}`)
	html.H1("Test")
	html.Table([]string{"header1", "header2", "header3"}, [][]string{{"a1", "a2", "a3"}, {"b1", "b2", "b3"}})
	html.P("Test paragraph for testing purposes")
	html.Prepare()
	html.WriteToFile("report.html")
}
