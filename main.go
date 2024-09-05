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

	tableStyle := &rephtml.Style{
		Props: tableProps,
		Tags:  []string{"table"},
	}
	tableStyle.PropMap(pmap)
	tableStyle.Prepare()

	// create table
	table := rephtml.NewTable()
	table.Id("myId")
	table.Class([]string{"testclass1", "testclass2"})
	table.Headers([]string{"testhdr1", "testhdr2"})
	table.AddRow([]string{"testval1", "testhval2"})
	table.AddStyle("color", "blue")
	table.AddStyle("color", "red") // override color from blue to red
	table.Prepare()

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
	html.StyleString(`#myId {
			font-size: 60px;
		}`)
	html.Style(tableStyle)
	html.H1String("Test")
	html.TableString([]string{"hdr1", "hdr2", "hdr3"}, [][]string{{"a1", "a2", "a3"}, {"b1", "b2", "b3"}})
	html.PString("Test paragraph for testing purposes")
	html.PStringWithStyle("Test style paragraph for testing purposes", "font-size: 30px")
	html.Table(table)
	html.Prepare()
	html.WriteToFile("report.html")

	// rgx := regexp.MustCompile(`^[<a-zA-Z]+(?: [a-zA-Z0-9]+="[^"]+")+([>])*$`)
	// fmt.Println(rgx.MatchString("<table id=\"myId\">"))
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
