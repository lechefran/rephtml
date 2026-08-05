# rephtml

rephtml is a fluent Go package for building HTML documents from typed components. Regular HTML element components implement the `Element` interface and can render themselves through `Render` or `HTML` without requiring callers to manage an explicit preparation step. The document root uses error-returning render methods so document structure errors are not hidden.

The package is designed around small composable structs:

- `NewHtmlFile` creates the document root and owns the final file output.
- `AddToHead` and `AddToBody` place components into generated `<head>` and `<body>` sections, rejecting content that does not belong there.
- `Add` composes child elements into container elements, ignoring nil children.
- `HTML` and `Render` prepare regular element components internally before returning output.
- Rendering does not mutate the tree, so a built tree is safe to render concurrently.
- `Text` methods escape normal text content by default.
- Attribute setters escape attribute values by default, and URL attributes are filtered by scheme.
- `StyleElement.Text` and `Script.Text` intentionally preserve raw CSS and JavaScript content.
- `WriteToFile` writes a human-readable formatted HTML file, declaring `<!DOCTYPE html>`.
- `Render`, `RenderString`, `RenderFormatted`, and `RenderFormattedString` return generated HTML plus any document error.

## Installation

```sh
go get github.com/lechefran/rephtml
```

## Quick Start

```go
package main

import rephtml "github.com/lechefran/rephtml"

func main() {
	html := rephtml.NewHtmlFile().Lang("en")

	html.AddToHead(rephtml.NewTitle().Text("Sales Report"))
	html.AddToHead(rephtml.NewStyleElement().Text(`
body {
	font-family: Arial, sans-serif;
	color: #1f2933;
}
main {
	max-width: 720px;
	margin: 40px auto;
}
`))

	table := rephtml.NewTable().
		Headers([]string{"Region", "Revenue", "Growth"}).
		AddRow([]string{"North", "$125,000", "12%"}).
		AddRow([]string{"South", "$98,000", "8%"})

	html.AddToBody(rephtml.NewMain().
		Add(rephtml.NewH1().Text("Sales Report")).
		Add(rephtml.NewP().Text("Quarterly regional performance")).
		Add(table))

	if err := html.WriteToFile("report.html"); err != nil {
		panic(err)
	}
}
```

## Rendering Elements

Regular elements prepare themselves when you call `HTML` or `Render`. There is no separate preparation step to remember.

```go
package main

import (
	"fmt"

	rephtml "github.com/lechefran/rephtml"
)

func main() {
	paragraph := rephtml.NewP().Text(`A&B < C`)

	fmt.Println(paragraph.HTML())
	fmt.Println(string(paragraph.Render()))
}
```

Output:

```html
<p>A&amp;B &lt; C</p>
<p>A&amp;B &lt; C</p>
```

Containers also render their current child state automatically:

```go
card := rephtml.NewSection().
	AriaLabel("Invoice summary").
	Add(rephtml.NewH2().Text("Invoice #1042")).
	Add(rephtml.NewP().Text("Payment due Friday"))

fmt.Println(card.HTML())
```

## How Elements Are Built

Elements embed a shared generic base that supplies the render entry points
(`Render`, `HTML`) and the common setters (`AddStyle`, `AddStyles`, `Style`, and
`Add` on containers). The base is parameterised by the concrete element type, so
those setters return the element's own type and chaining works as usual:

```go
div := rephtml.NewDiv().AddStyle("color", "red").Add(rephtml.NewP()) // *Div
```

Because the methods are promoted rather than declared per element, godoc renders
their result type as `Self`. The concrete result is always the element's own
pointer type, for example `*Div` above.

Elements must be created with their `New*` constructor, which is what binds the
element to its base. A zero value such as `&rephtml.Div{}` is not a usable
element.

`Add` ignores nil children, including a typed nil, so a builder that returns
`nil` on a miss composes without a guard at every call site:

```go
func badge(u *User) *rephtml.Span {
	if u == nil {
		return nil // safe to Add; nothing is emitted
	}
	return rephtml.NewSpan().Text(u.Name)
}
```

## Concurrency

Rendering does not mutate the element tree. An element writes into the buffer it
is given and keeps no render state of its own, so a tree built once can be
rendered from any number of goroutines:

```go
page := buildPage() // once, at startup

http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
	w.Write(page.Render()) // safe from concurrent requests
})
```

Building is not concurrency-safe. Calling setters on an element while another
goroutine renders it is a data race, as it would be for any Go value. Finish
building before you share the tree.

## Rendering Documents

`HtmlFile` returns errors from render methods because document structure can be invalid. Use `RenderString` for compact HTML, `RenderFormattedString` for readable output, and `WriteToFile` for files.

Documents are emitted with `<!DOCTYPE html>`, so browsers render them in standards mode.

```go
html := rephtml.NewHtmlFile().Lang("en")
html.AddToHead(rephtml.NewMeta().Charset("utf-8"))
html.AddToHead(rephtml.NewTitle().Text("Dashboard"))
html.AddToBody(
	rephtml.NewMain().
		Add(rephtml.NewH1().Text("Dashboard")).
		Add(rephtml.NewP().Text("Operational overview")),
)

compact, err := html.RenderString()
if err != nil {
	panic(err)
}

formatted, err := html.RenderFormattedString()
if err != nil {
	panic(err)
}

fmt.Println(compact)
fmt.Println(formatted)
```

`AddToHead` and `AddToBody` enforce simple document structure. For example, adding a body-only element to the head records an error that render/write methods return:

```go
html := rephtml.NewHtmlFile()
html.AddToHead(rephtml.NewP().Text("not allowed in head"))

if err := html.Err(); err != nil {
	fmt.Println(err)
}
```

The `Head` and `Body` wrappers apply the same rule. `AddToHead` flattens a `Head`
into the generated head rather than nesting it, so the wrapper checks its own
content as it is added and carries any error into the document:

```go
head := rephtml.NewHead().Add(rephtml.NewDiv()) // div is body content
fmt.Println(head.Err())                          // reported here

html := rephtml.NewHtmlFile()
html.AddToHead(head)
fmt.Println(html.Err()) // and again here, once merged
```

## Composition Examples

Build structured tables when you need explicit table sections:

```go
table := rephtml.NewTable().
	AddCaption(rephtml.NewCaption().Text("Quarterly revenue")).
	AddThead(
		rephtml.NewThead().AddTr(
			rephtml.NewTr().
				AddTh(rephtml.NewTh().Add(rephtml.NewSpan().Text("Region"))).
				AddTh(rephtml.NewTh().Add(rephtml.NewSpan().Text("Revenue"))),
		),
	).
	AddTbody(
		rephtml.NewTbody().AddTr(
			rephtml.NewTr().
				AddTd(rephtml.NewTd().Add(rephtml.NewSpan().Text("North"))).
				AddTd(rephtml.NewTd().Add(rephtml.NewSpan().Text("$125,000"))),
		),
	)

fmt.Println(table.HTML())
```

Build forms with the same composable `Add` pattern:

```go
form := rephtml.NewForm().
	Action("/search").
	Method("get").
	Add(rephtml.NewLabel().For("q").Text("Search")).
	Add(rephtml.NewInput().
		Type("search").
		Id("q").
		Name("q").
		Placeholder("Orders, customers, invoices")).
	Add(rephtml.NewButton().Type("submit").Text("Run search"))

fmt.Println(form.HTML())
```

## CSS Rules

Use `StyleElement.Text` when you already have CSS text. Use `StyleMap` for both inline styles and generated CSS rules; property names are emitted exactly as written, so custom properties and vendor prefixes work without package changes. CSS at-rules use dedicated builders such as `NewMediaRule`, `NewKeyframesRule`, `NewFontFaceRule`, `NewImportRule`, and `NewCharsetRule`.

```go
body := rephtml.NewStyleRule("body")
body.Props = rephtml.StyleMap{
	"background":  "#f7f7fb",
	"color":       "#1f2937",
	"font-family": "Arial, sans-serif",
	"margin":      "0",
}

mainWide := rephtml.NewStyleRule("main")
mainWide.Props = rephtml.StyleMap{
	"max-width": "960px",
}

html.AddToHead(rephtml.NewStyleElement().
	Type("text/css").
	AddRule(body).
	AddRule(rephtml.NewMediaRule("(min-width: 800px)").AddRule(mainWide)))
```

Inline styles use the same `StyleMap` and render in stable key order:

```go
button := rephtml.NewButton().
	Text("Save").
	Style(rephtml.StyleMap{
		"background":    "#2563eb",
		"border":        "0",
		"border-radius": "4px",
		"color":         "#fff",
		"padding":       "0.5rem 0.75rem",
	})

fmt.Println(button.HTML())
```

## Formatting and Escaping

rephtml escapes normal text and attribute values so characters like `<`, `>`, `&`, and quotes do not corrupt the generated HTML. Content is escaped for the context it lands in, which is not always HTML entity escaping:

| Content | Handling |
|---|---|
| Text nodes and attribute values | HTML entity escaped |
| Comment text | Sanitised, since character references are not decoded inside a comment. Hyphen runs are broken up so the text cannot close the comment early |
| Structured CSS: selectors, properties, values, `@media` queries, `@keyframes` names, `@import` hrefs | Escaped so the text cannot close the surrounding `<style>` element. A `<` used legitimately, as in the media query range syntax `(400px <= width)`, is left alone |
| URL attributes (`href`, `src`, `action`, …) | Escaped, then filtered by scheme so a `javascript:` URL cannot be emitted. See [URLs](#urls) |
| `StyleElement.Text` and `Script.Text` | **Raw, by design.** These are the escape hatches for hand-written CSS and JavaScript. Do not pass untrusted input to them |

For regular elements, call `HTML()` or `Render()` directly:

```go
paragraph := rephtml.NewP().Text("Hello")
fmt.Println(paragraph.HTML())
```

Raw CSS and JavaScript are preserved intentionally:

```go
style := rephtml.NewStyleElement().Text(`
.notice > strong {
	color: #b91c1c;
}
`)

script := rephtml.NewScript().Text(`console.log("ready <now>");`)
```

### URLs

Escaping keeps a URL inside its quotes but says nothing about what following it
does, so URL attributes are also filtered by scheme. A URL whose scheme executes
script is replaced with `rephtml.BlockedURL`:

```go
link := rephtml.NewAnchor().Link("javascript:alert(1)").Text("Click me")
fmt.Println(link.HTML())
// <a href="#rephtml-blocked-url">Click me</a>
```

The value is replaced rather than dropped so the problem is visible in the
output instead of quietly changing where a link points. Use `IsSafeURL` to
detect it up front and decide for yourself:

```go
if !rephtml.IsSafeURL(u) {
	return fmt.Errorf("refusing to link to %q", u)
}
```

What is blocked:

- `javascript:`, `vbscript:`, `livescript:`, `mocha:`
- `data:` URLs a browser would treat as a document, which is anything other than
  an image, audio, video or font media type. `data:image/png` is fine;
  `data:text/html` and `data:image/svg+xml` are not

Everything else passes, including relative URLs, fragments, queries,
protocol-relative URLs, and schemes like `mailto:`, `tel:`, `ftp:`, `blob:` and
application deep links. This is a blocklist rather than an allowlist because a
static site generator has good reason to emit custom schemes, while the set that
actually runs script is small and has not grown in years.

The check normalises the way a browser does before reading the scheme — leading
whitespace and control characters, and tab, newline and NUL anywhere — so
`java&#9;script:` and `  JAVASCRIPT:` are caught too.

Filtering covers `href`, `src`, `srcset`, `action`, `formaction`, `data`,
`poster`, `cite`, `manifest` and `usemap`. `srcset` is filtered per candidate,
so one bad entry does not discard the rest of the list.

`WriteToFile` and `RenderFormatted` format documents generated by rephtml with indentation and return filesystem or document structure errors. That formatter is intentionally scoped to rephtml output rather than exposed as a general-purpose HTML formatter. CSS inside `<style>` blocks is also indented for readability, while whitespace-sensitive blocks such as `<pre>` and `<textarea>` are preserved.

## Examples

The examples above are self-contained and use the root package import path: `github.com/lechefran/rephtml`.

## Supported Elements
### Document Structure
- [X] `<html>` - Root element
- [X] `<head>` - Document head
- [X] `<body>` - Document body
- [X] `<title>` - Document title
- [X] `<base>` - Base URL
- [X] `<link>` - External resource link
- [X] `<meta>` - Metadata
- [X] `<style>` - Internal CSS

### Content Sectioning
- [X] `<header>` - Header section
- [X] `<nav>` - Navigation section
- [X] `<main>` - Main content
- [X] `<section>` - Generic section
- [X] `<article>` - Article content
- [X] `<aside>` - Sidebar content
- [X] `<footer>` - Footer section
- [X] `<address>` - Contact information
- [X] `<h1>` through `<h6>` - Headings
- [X] `<hgroup>` - Heading group

### Text Content
- [X] `<div>` - Generic container
- [X] `<p>` - Paragraph
- [X] `<hr>` - Horizontal rule
- [X] `<pre>` - Preformatted text
- [X] `<blockquote>` - Block quotation
- [X] `<ol>` - Ordered list
- [X] `<ul>` - Unordered list
- [X] `<menu>` - Menu list
- [X] `<li>` - List item
- [X] `<dl>` - Description list
- [X] `<dt>` - Description term
- [X] `<dd>` - Description details
- [X] `<figure>` - Figure with caption
- [X] `<figcaption>` - Figure caption
- [X] `<search>` - Search section

### Inline Text Semantics
- [X] `<a>` - Anchor/link
- [X] `<abbr>` - Abbreviation
- [X] `<b>` - Bold text
- [X] `<bdi>` - Bidirectional isolate
- [X] `<bdo>` - Bidirectional override
- [X] `<br>` - Line break
- [X] `<cite>` - Citation
- [X] `<code>` - Code
- [X] `<data>` - Machine-readable data
- [X] `<dfn>` - Definition
- [X] `<em>` - Emphasis
- [X] `<i>` - Italic
- [X] `<kbd>` - Keyboard input
- [X] `<mark>` - Highlighted text
- [X] `<q>` - Inline quotation
- [X] `<ruby>` - Ruby annotation
- [X] `<rb>` - Ruby base
- [X] `<rt>` - Ruby text
- [X] `<rtc>` - Ruby text container
- [X] `<rp>` - Ruby parentheses
- [X] `<s>` - Strikethrough
- [X] `<samp>` - Sample output
- [X] `<small>` - Small text
- [X] `<span>` - Generic inline container
- [X] `<strong>` - Strong importance
- [X] `<sub>` - Subscript
- [X] `<sup>` - Superscript
- [X] `<time>` - Date/time
- [X] `<u>` - Underline
- [X] `<var>` - Variable
- [X] `<wbr>` - Word break opportunity

### Image and Multimedia
- [X] `<area>` - Image map area
- [X] `<audio>` - Audio content
- [X] `<img>` - Image
- [X] `<map>` - Image map
- [X] `<track>` - Media track
- [X] `<video>` - Video content

### Embedded Content
- [X] `<embed>` - External content
- [X] `<iframe>` - Inline frame
- [X] `<object>` - External object
- [X] `<picture>` - Responsive image container
- [X] `<portal>` - Portal element
- [X] `<source>` - Media source

### SVG and MathML
- [X] `<svg>` - SVG graphics
- [X] `<math>` - MathML mathematics

### Scripting
- [X] `<canvas>` - Graphics canvas
- [X] `<noscript>` - No script fallback
- [X] `<script>` - Script

### Demarcating Edits
- [X] `<del>` - Deleted text
- [X] `<ins>` - Inserted text

### Table Content
- [X] `<table>` - Table
- [X] `<caption>` - Table caption
- [X] `<colgroup>` - Column group
- [X] `<col>` - Table column
- [X] `<tbody>` - Table body
- [X] `<thead>` - Table head
- [X] `<tfoot>` - Table foot
- [X] `<tr>` - Table row
- [X] `<td>` - Table data cell
- [X] `<th>` - Table header cell

### Forms
- [X] `<form>` - Form
- [X] `<label>` - Form label
- [X] `<input>` - Form input
- [X] `<button>` - Button
- [X] `<select>` - Selection list
- [X] `<datalist>` - Data list options
- [X] `<optgroup>` - Option group
- [X] `<option>` - Option
- [X] `<textarea>` - Text area
- [X] `<output>` - Form output
- [X] `<progress>` - Progress indicator
- [X] `<meter>` - Scalar measurement
- [X] `<fieldset>` - Form field grouping
- [X] `<legend>` - Fieldset legend

### Interactive Elements
- [X] `<details>` - Disclosure widget
- [X] `<summary>` - Details summary
- [X] `<dialog>` - Dialog box

### Web Components
- [X] `<slot>` - Web component slot
- [X] `<template>` - Template element

---

**Total Elements:** 118 HTML elements (excluding deprecated ones)
