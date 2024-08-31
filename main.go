package main

import (
	"bytes"
	"encoding/csv"
	"fmt"
	"log"
	"os"
	"strings"

	"github.com/lechefran/rephtml/rephtml"
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
	// pmap := rephtml.NewPropMap()
	// tableProps := rephtml.CssProps{
	// 	FontFamily:  "Arial",
	// 	FontSize:    "16px",
	// 	MarginLeft:  "auto",
	// 	MarginRight: "auto",
	// 	Width:       "80%",
	// }
	// tableTag := []string{"table"}
	// tableStyle := &rephtml.StyleTag{
	// 	Props: tableProps,
	// 	Tags:  tableTag,
	// }
	// tableStyle.PropMap(pmap)
	// tableStyle.Prepare()

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
	// html.Style(tableStyle)
	html.H1String("Test")
	html.TableString([]string{"header1", "header2", "header3"}, [][]string{{"a1", "a2", "a3"}, {"b1", "b2", "b3"}})
	html.PString("Test paragraph for testing purposes")
	html.Prepare()
	html.WriteToFile("report.html")

	// html1 := &rephtml.HtmlFile{}
	// html1.Prepare()
	// html1.WriteToFile("test.html")

	// contents := ReadCsv("cssprops.csv")
	// cmap := map[string]string{}
	// for _, c := range contents {
	// 	tmp := strings.Fields(c[0])[0]
	// 	pname, vname := tmp, ""
	// 	tmp = strings.ReplaceAll(tmp, "-", "\\s")
	// 	lst := strings.Split(tmp, "\\s")
	// 	for _, l := range lst {
	// 		vname += strings.ToUpper(string(l[0])) + string(l[1:])
	// 	}
	// 	cmap[vname] = pname
	// }

	// keys := make([]string, 0, len(cmap))
	// for k := range cmap {
	// 	keys = append(keys, k)
	// }
	// sort.Strings(keys)
	// for _, k := range keys {
	// 	fmt.Printf("pmap[\"%s\"] = \"%s\"\n", k, cmap[k])
	// }

	tst := "th, td {font-family:Arial;padding:10px;text-align:center;}"
	open, close := strings.Index(tst, "{"), len(tst)-1
	openb, closeb := tst[open], tst[close]
	fmt.Println("original string: " + tst)
	fmt.Println("open: " + string(openb))
	fmt.Println("string leading to open: " + tst[:open])
	fmt.Println("close: " + string(closeb))
	tst = strings.ReplaceAll(tst, ";", "; ")
	fmt.Println("modified string after replace all: " + tst)
	tst1 := []byte(tst)
	tst2 := bytes.Fields(tst1)
	for _, t := range tst2 {
		fmt.Println(string(t))
	}
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
