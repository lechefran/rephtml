package main

import (
	"bytes"
	"encoding/csv"
	"log"
	"os"

	rephtml "github.com/lechefran/rephtml/components"
)

const indent = "\n\t"
const newline = "\n"
const tab = "\t"

type HtmlDiv struct {
	buf     bytes.Buffer
	content [][]byte
}

func NewDiv() *HtmlDiv {
	return &HtmlDiv{}
}

func (d *HtmlDiv) ReadBuffer(b *bytes.Buffer) {
	b.Read(d.buf.Bytes())
}

func main() {
	// initialize a reusable prop map
	pmap := rephtml.NewPropMap()
	tableProps := rephtml.CssProps{
		FontFamily:  "Arial",
		FontSize:    "16px",
		MarginLeft:  "auto",
		MarginRight: "auto",
		Width:       "80%",
	}
	tableTag := []string{"table"}
	tableStyle := &rephtml.Style{
		Props: tableProps,
		Tags:  tableTag,
	}
	tableStyle.PropMap(pmap)
	tableStyle.Prepare()

	html := *rephtml.NewHtmlFile()
	html.StyleString(`h1, h2, h3, h4, h5, h6, p {
			font-family: Arial;
			text-align: center;
		}`)
	html.StyleString(`th, td {
			font-family: Arial;
			padding: 10px;
			text-align: center;
		}`)
	html.Style(tableStyle)
	html.H1String("Test")
	html.TableString([]string{"header1", "header2"}, [][]string{{"a1", "a2", "a3"}, {"b1", "b2", "b3"}})
	html.PString("Test paragraph for testing purposes")
	html.PStringWithStyle("Test style paragraph for testing purposes", "font-size: 30px")
	html.Prepare()
	html.WriteToFile("report.html")
}

func ReadCsv(path string) [][]string {
	f, err := os.Open(path)
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()
	rdr := csv.NewReader(f)
	rdr.FieldsPerRecord = -1
	res, err := rdr.ReadAll()
	if err != nil {
		log.Fatal(err)
	}
	return res
}
