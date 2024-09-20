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

	// create parapgraph styles
	pgStyles := make(map[string]string, 2)
	pgStyles["color"] = "#a0d6b4"
	pgStyles["font-size"] = "60px"

	// create parapgraphs
	pg := rephtml.NewP()
	pg.Text("Hello, World!")
	pg.AddStyle("font-size", "60px")
	pg.AddStyles(pgStyles)
	pg.Prepare()

	pg1 := rephtml.NewP()
	pg1.Text("Test paragraph in DIV")
	pg1.AddStyle("color", "#FFFFFF")
	pg1.Prepare()

	h1 := rephtml.NewH1()
	h1.Text("Paragraph struct!")
	h1.Prepare()

	// create comment
	c := rephtml.NewComment().Text("This is a test comment.")
	c.Prepare()

	// create div styles
	dStyles := make(map[string]string, 1)
	dStyles["background-color"] = "#CCCCFF"

	// create div
	d := rephtml.NewDiv().Add(pg1).Add(h1).AddStyles(dStyles)
	d.Add(c)
	d.Add(table) // to fix
	d.Prepare()

	// create another div
	d1 := rephtml.NewDiv().Add(h1).AddStyles(dStyles)
	d1.Prepare()

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
	html.P(pg)
	html.Table(table)
	html.Div(d)
	html.Div(d1)
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
