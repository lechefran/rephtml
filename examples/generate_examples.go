package main

import (
	"log"

	rephtml "github.com/lechefran/rephtml/components"
)

func main() {
	writeReportExample()
	writeFormExample()
	writeMediaExample()
}

func writeReportExample() {
	html := rephtml.NewHtmlFile().Lang("en")

	title := rephtml.NewTitle().Text("Sales Report")
	html.AddToHead(title)
	html.AddToHead(rephtml.NewStyleElement().Text(`
body {
	font-family: Arial, sans-serif;
	color: #1f2933;
	background: #f6f8fa;
}
main {
	max-width: 880px;
	margin: 0 auto;
	padding: 32px;
}
table {
	width: 100%;
	border-collapse: collapse;
	background: #ffffff;
}
caption {
	font-weight: 700;
	padding: 12px;
	text-align: left;
}
th, td {
	border: 1px solid #d6dde5;
	padding: 10px 12px;
	text-align: left;
}
`))

	heading := rephtml.NewHgroup().
		AddStyle("margin-bottom", "24px").
		Add(rephtml.NewH1().Text("Sales Report")).
		Add(rephtml.NewP().Text("Quarterly regional performance"))

	table := rephtml.NewTable().
		AddCaption(rephtml.NewCaption().Text("Revenue by region")).
		AddStyle("box-shadow", "0 1px 3px rgba(0,0,0,0.08)").
		AddThead(rephtml.NewThead().AddTr(
			rephtml.NewTr().
				AddTh(rephtml.NewTh().Add(rephtml.NewSpan().Text("Region"))).
				AddTh(rephtml.NewTh().Add(rephtml.NewSpan().Text("Revenue"))).
				AddTh(rephtml.NewTh().Add(rephtml.NewSpan().Text("Growth"))),
		)).
		AddTbody(rephtml.NewTbody().
			AddTr(rephtml.NewTr().
				AddTd(rephtml.NewTd().Add(rephtml.NewSpan().Text("North"))).
				AddTd(rephtml.NewTd().Add(rephtml.NewSpan().Text("$125,000"))).
				AddTd(rephtml.NewTd().Add(rephtml.NewStrong().Text("12%")))).
			AddTr(rephtml.NewTr().
				AddTd(rephtml.NewTd().Add(rephtml.NewSpan().Text("South"))).
				AddTd(rephtml.NewTd().Add(rephtml.NewSpan().Text("$98,000"))).
				AddTd(rephtml.NewTd().Add(rephtml.NewStrong().Text("8%")))))

	main := rephtml.NewMain().
		Add(heading).
		Add(rephtml.NewSection().
			AriaLabel("Revenue summary").
			Add(rephtml.NewP().Text("The north region led quarterly revenue.")).
			Add(table))

	html.AddToBody(main)
	write(html, "examples/output/report.html")
}

func writeFormExample() {
	html := rephtml.NewHtmlFile().Lang("en")
	html.AddToHead(rephtml.NewTitle().Text("Search Form"))
	html.AddToHead(rephtml.NewStyleElement().Text(`
body {
	font-family: Arial, sans-serif;
	background: #eef4ff;
	color: #243b53;
}
main {
	max-width: 560px;
	margin: 48px auto;
}
form {
	display: grid;
	gap: 12px;
	background: #ffffff;
	padding: 24px;
	border: 1px solid #c9d8ee;
}
input, button {
	font: inherit;
	padding: 10px 12px;
}
button {
	background: #1d4ed8;
	color: #ffffff;
	border: 0;
}
`))

	search := rephtml.NewSearch().Add(
		rephtml.NewForm().
			AddStyle("border-radius", "6px").
			Action("/search").
			Method("get").
			Add(rephtml.NewLabel().For("query").Text("Search reports")).
			Add(rephtml.NewInput().
				AddStyle("border", "1px solid #9fb3c8").
				Id("query").
				Type("search").
				Name("q").
				Placeholder("Region or keyword")).
			Add(rephtml.NewButton().Type("submit").Text("Search")),
	)

	html.AddToBody(rephtml.NewMain().Add(search))
	write(html, "examples/output/form.html")
}

func writeMediaExample() {
	html := rephtml.NewHtmlFile().Lang("en")
	html.AddToHead(rephtml.NewTitle().Text("Media Card"))
	html.AddToHead(rephtml.NewStyleElement().Text(`
body {
	font-family: Arial, sans-serif;
	margin: 0;
	background: #fbfaf7;
	color: #2f2f2f;
}
main {
	max-width: 720px;
	margin: 40px auto;
}
figure {
	margin: 0;
}
img {
	display: block;
	max-width: 100%;
	height: auto;
}
figcaption {
	padding: 12px 0;
	color: #5f6c72;
}
`))

	figure := rephtml.NewFigure().
		AddStyle("border", "1px solid #ded8cf").
		AddStyle("padding", "16px").
		AddStyle("background", "#ffffff").
		Add(rephtml.NewPicture().
			Add(rephtml.NewSource().
				Srcset("chart-large.webp").
				Media("(min-width: 800px)").
				Type("image/webp")).
			Add(rephtml.NewImg().
				Src("chart.png").
				Alt("Bar chart showing regional revenue").
				Width("640").
				Height("360"))).
		Add(rephtml.NewFigcaption().Add(rephtml.NewSpan().Text("Regional revenue snapshot")))

	html.AddToBody(rephtml.NewMain().Add(figure))
	write(html, "examples/output/media.html")
}

func write(html *rephtml.HtmlFile, path string) {
	html.WriteToFile(path)
	log.Printf("wrote %s", path)
}
