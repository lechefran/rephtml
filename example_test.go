package rephtml_test

import (
	"fmt"
	"os"

	"github.com/lechefran/rephtml"
)

func ExampleHtmlFile_WriteTo() {
	page := rephtml.NewHtmlFile().Lang("en")
	page.AddToHead(rephtml.NewTitle().Text("Hello"))
	page.AddToBody(rephtml.NewP().Text("Hello, <world>"))

	if _, err := page.WriteTo(os.Stdout); err != nil {
		fmt.Println(err)
	}
	// Output:
	// <!DOCTYPE html><html lang="en"><head><title>Hello</title></head><body><p>Hello, &lt;world&gt;</p></body></html>
}

func ExampleHtmlFile_WriteFormattedTo() {
	page := rephtml.NewHtmlFile()
	page.AddToBody(rephtml.NewDiv().Add(rephtml.NewP().Text("Hi")))

	if _, err := page.WriteFormattedTo(os.Stdout); err != nil {
		fmt.Println(err)
	}
	// Output:
	// <!DOCTYPE html>
	// <html>
	// 	<body>
	// 		<div>
	// 			<p>Hi</p>
	// 		</div>
	// 	</body>
	// </html>
}

func ExampleDiv_Id() {
	nav := rephtml.NewDiv().
		Id("menu").
		AddClass("nav").
		AddClasses([]string{"sticky", "dark"}).
		Add(rephtml.NewSpan().AddClass("item").Text("Home"))

	fmt.Println(nav.HTML())
	// Output:
	// <div id="menu" class="nav sticky dark"><span class="item">Home</span></div>
}

func ExampleLabel_AddText() {
	label := rephtml.NewLabel().
		For("email").
		AddText("Email ").
		Add(rephtml.NewStrong().Text("required")).
		AddText(" & private")

	fmt.Println(label.HTML())
	// Output:
	// <label for="email">Email <strong>required</strong> &amp; private</label>
}

func ExampleNewDiv() {
	div := rephtml.NewDiv().
		AddStyle("color", "red").
		Add(rephtml.NewP().Text("Fish & chips"))

	fmt.Println(div.HTML())
	// Output:
	// <div style="color: red;"><p>Fish &amp; chips</p></div>
}
