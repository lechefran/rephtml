# rephtml

[![CI](https://github.com/lechefran/rephtml/actions/workflows/ci.yml/badge.svg)](https://github.com/lechefran/rephtml/actions/workflows/ci.yml)
[![Go Reference](https://pkg.go.dev/badge/github.com/lechefran/rephtml.svg)](https://pkg.go.dev/github.com/lechefran/rephtml)
[![Latest release](https://img.shields.io/github/v/release/lechefran/rephtml)](https://github.com/lechefran/rephtml/releases/latest)
[![Go version](https://img.shields.io/github/go-mod/go-version/lechefran/rephtml)](go.mod)
[![License](https://img.shields.io/github/license/lechefran/rephtml)](LICENSE)

rephtml is a fluent Go package for building HTML documents from typed components. Regular HTML element components implement the `Element` interface and can render themselves through `Render` or `HTML` without requiring callers to manage an explicit preparation step. The document root uses error-returning render methods so document structure errors are not hidden.

The package is designed around small composable structs:

- `NewHtmlFile` creates the document root and owns the final file output.
- `AddToHead` and `AddToBody` place components into generated `<head>` and `<body>` sections, rejecting content that does not belong there.
- `Add` composes child elements into container elements, ignoring nil children.
- `HTML` and `Render` prepare regular element components internally before returning output.
- Rendering does not mutate the tree, so a built tree is safe to render concurrently.
- `Text` methods escape normal text content by default.
- Attribute setters escape attribute values by default, and URL attributes are filtered by scheme.
- `StyleElement.Text` and `Script.Text` intentionally preserve raw CSS and JavaScript content, only rewriting the `</style` and `</script` sequences that would end the element early.
- `WriteToFile` writes a human-readable formatted HTML file, declaring `<!DOCTYPE html>`. The write is atomic, so a failure never leaves a truncated file.
- `Render`, `RenderString`, `RenderFormatted`, and `RenderFormattedString` return generated HTML plus any document error.
- `WriteTo` and `WriteFormattedTo` stream a document to any `io.Writer`, such as an `http.ResponseWriter`.

The current release is **v0.2.0**. The module is still below v1, so a minor release may contain breaking changes. Deprecated names keep working until v1.0, and every release is described in the [changelog](CHANGELOG.md).

## Installation

```sh
go get github.com/lechefran/rephtml@v0.2.0
```

Use `@latest` to follow new releases. The module requires the Go version declared in [go.mod](go.mod).

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
(`Render`, `HTML`) and the common setters (`AddStyle`, `AddStyles`, `Style`,
`Id`, `AddClass`, `AddClasses`, `Class`, and `Add` on containers). The base is
parameterised by the concrete element type, so those setters return the
element's own type and chaining works as usual:

```go
div := rephtml.NewDiv().AddStyle("color", "red").Add(rephtml.NewP()) // *Div
```

Every element supports the global `id` and `class` attributes. They are written
together with the inline style, as `id`, then `class`, then `style`:

```go
rephtml.NewDiv().Id("menu").AddClass("nav").AddStyle("color", "red").HTML()
// <div id="menu" class="nav" style="color: red;"></div>
```

### Text and AddText

`AddText` always appends text. `Text` depends on the element, and v1.0 will make
it consistent:

| `Text` replaces the text | `Text` appends text, like `AddText` |
|---|---|
| inline text elements such as `Span`, `B` and `Anchor`; headings; `P`, `Pre`, `Blockquote`; `Script`; `Caption`; `Comment` | elements that also hold child elements: `Title`, `Label`, `Button`, `Option`, `Output`, `Textarea`, `Progress`, `Meter`, `Legend`; and `StyleElement` |

Use `AddText` whenever you mean to append, so your code reads the same on every
element and keeps working when `Text` changes.

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

### Attribute value types

A setter takes an `int` when the attribute is always a whole number, such as
`Td.Colspan`, `Col.Span`, `Ol.Start` and `Li.Value`. It takes a `string` when
the value can carry units, dates, decimals or keywords, such as `Img.Width`,
`Input.Min` and `Meter.Value`. The string is written as given.

A few older setters take a `string` for an attribute that is always a whole
number: `Input.Size`, `Input.MaxLength`, `Input.MinLength`, `Select.Size`,
`Textarea.Rows`, `Textarea.Cols`, `Textarea.MaxLength` and
`Textarea.MinLength`. They keep that signature until v1.0, so pass the number
as text, for example `Rows("4")`.

## Concurrency

Rendering does not mutate the element tree. An element writes into the buffer it
is given and keeps no render state of its own, so a tree built once can be
rendered from any number of goroutines:

```go
page := buildPage() // a *rephtml.HtmlFile, built once at startup

http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	if _, err := page.WriteTo(w); err != nil { // safe from concurrent requests
		log.Printf("write page: %v", err)
	}
})
```

`WriteTo` renders into a pooled buffer and copies it to the writer, so it avoids
the extra copy that `Render` makes. A document error is returned before anything
is written, so a broken page never produces a partial response.

Building is not concurrency-safe. Calling setters on an element while another
goroutine renders it is a data race, as it would be for any Go value. Finish
building before you share the tree.

## Rendering Documents

`HtmlFile` returns errors from render methods because document structure can be invalid. Every output method returns the document error before producing any output:

| Method | Output | Destination |
|---|---|---|
| `Render` | Compact | `[]byte` |
| `RenderString` | Compact | `string` |
| `RenderFormatted` | Indented, human-readable | `[]byte` |
| `RenderFormattedString` | Indented, human-readable | `string` |
| `WriteTo` | Compact | Any `io.Writer`. Implements `io.WriterTo` |
| `WriteFormattedTo` | Indented, human-readable | Any `io.Writer` |
| `WriteToFile` | Indented, human-readable | A file, written atomically |

Use `WriteTo` when serving a page or copying it somewhere, since it makes no intermediate copy. Use the formatted variants for files and debugging. Formatting puts block elements on their own indented lines but keeps each run of inline content, such as `Hello <b>x</b>.`, on one line exactly as rendered, so it never changes how a page displays. Prefer the compact form when exact bytes matter.

`WriteToFile` writes to a temporary file in the destination directory, syncs it, and renames it over `path`. A failed write therefore never leaves a truncated file, and an existing file is only replaced once the whole document is written. The file is created with mode 0644 (before the umask).

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

The document reads a wrapper when it renders, not when the wrapper is added.
Content, attributes and errors added to a `Head` or `Body` after `AddToHead` or
`AddToBody` all appear in the document. The wrapper's content renders at the
point it was added, among elements added directly. When several wrappers are
merged, the last non-empty `id`, `onload` and `onunload` win, classes accumulate,
and a later wrapper's style declaration overrides an earlier one:

```go
body := rephtml.NewBody()
html := rephtml.NewHtmlFile().AddToBody(body)
body.Add(rephtml.NewP().Text("added after merging")) // still rendered
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
| `StyleElement.Text`, `RawCSSRule` and `ImportRule` conditions | Raw CSS, by design, except that a `</style` sequence is rewritten as `\3c /style` so the text cannot close the element. Do not pass untrusted input as CSS |
| `Script.Text` | Raw JavaScript, by design, except that a `</script` sequence is rewritten as `<\/script`, which means the same thing inside a JavaScript string. Do not pass untrusted input as code |
| `Meta` with `http-equiv="refresh"` | The redirect target in `content` is filtered by scheme like a URL attribute |
| `Iframe.Srcdoc` | Escaped as an attribute, but the document inside runs. Pair untrusted content with `Sandbox` |

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

### Raw-text elements

`<style>` and `<script>` hold raw text: the browser does not decode entities inside them and ends the element at the first `</style` or `</script`, whatever the CSS or JavaScript around it says. rephtml leaves your code alone except for that one sequence, which it rewrites so the text cannot close the element:

```go
style := rephtml.NewStyleElement().Text("a{color:red}</style><script>alert(1)</script>")
fmt.Println(style.HTML())
// <style>a{color:red}\3c /style><script>alert(1)</script></style>

script := rephtml.NewScript().Text(`var tpl = "</script><b>x</b>";`)
fmt.Println(script.HTML())
// <script>var tpl = "<\/script><b>x</b>";</script>
```

Inside CSS, `\3c ` is the escape for `<`. Inside a JavaScript string, `\/` is just `/`, which is the same rewrite `encoding/json` and `html/template` use. Matching is case-insensitive, and a `<` that is not part of a closing tag, such as the `<=` in a media query, is untouched. This applies to `StyleElement.Text`, `RawCSSRule`, `ImportRule` conditions and `Script.Text`.

This is a guard against accidentally closing the element, not a way to run untrusted code safely. Do not pass untrusted input to either `Text` method.

A `<style>` element can only hold CSS, so `StyleElement.Add` accepts CSS rules and drops any other element, recording the problem on `StyleElement.Err`:

```go
style := rephtml.NewStyleElement().Add(rephtml.NewDiv())
fmt.Println(style.Err())
// cannot add Div to style element: only CSS rules are allowed
```

A few other values get the same care:

- `<meta http-equiv="refresh">` filters the redirect target in `content` by scheme, like a URL attribute: `0; url=javascript:alert(1)` is rendered as `0; url=#rephtml-blocked-url`.
- `Iframe.Srcdoc` is escaped as an attribute, but the browser runs the document inside it. Use `Sandbox` for anything untrusted.
- Inline `style` attributes skip declarations with an empty property or value, so `AddStyle("margin", "")` writes nothing instead of `margin: ;`, and no `style` attribute is written when nothing remains.

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

`WriteToFile` and `RenderFormatted` format documents generated by rephtml with indentation and return filesystem or document structure errors. That formatter is intentionally scoped to rephtml output rather than exposed as a general-purpose HTML formatter. CSS inside `<style>` blocks is also indented for readability, while whitespace-sensitive content is preserved: `<pre>` and `<textarea>` bodies, and inline elements with the text and whitespace around them.

## Examples

The examples above are self-contained and use the root package import path: `github.com/lechefran/rephtml`. Runnable examples also appear on [pkg.go.dev](https://pkg.go.dev/github.com/lechefran/rephtml#pkg-examples).

## Upgrading from v0.1.0-beta

v0.2.0 does not remove any API, so code that compiled against `v0.1.0-beta` still compiles. Some setters have new, consistently capitalised names, and the old names are deprecated. They still work and are kept until v1.0:

| Deprecated | Use instead |
|---|---|
| `Form.Autocomplete`, `Input.Autocomplete` | `AutoComplete` |
| `Form.Novalidate` | `NoValidate` |
| `Input.Maxlength`, `Input.Minlength` | `MaxLength`, `MinLength` |
| `Iframe.Allowfullscreen` | `AllowFullscreen` |
| `Iframe.Referrerpolicy`, `Portal.Referrerpolicy`, `Script.Referrerpolicy` | `ReferrerPolicy` |
| `Script.Crossorigin` | `CrossOrigin` |

`go vet` and staticcheck report uses of deprecated names, so they are easy to find.

Some rendered output changes, all of them corrections:

- `</style` in `StyleElement.Text`, `RawCSSRule` and `ImportRule` conditions, and `</script` in `Script.Text`, are rewritten so the text cannot close its element. See [Raw-text elements](#raw-text-elements).
- `<meta http-equiv="refresh">` replaces an unsafe redirect target with `BlockedURL`.
- `StyleElement.Add` drops elements that are not CSS rules and reports them through `Err`.
- `Ol.Start(0)` and `Li.Value(0)` now emit the attribute. Zero was previously treated as unset.
- Inline `style` attributes no longer include declarations with empty values.
- `Link` writes `crossorigin` and `referrerpolicy` in lower case.
- `Table.Headers`, `Rows`, `AddRow` and `AddRows` copy the slices you pass in, so changing a slice after the call no longer changes the table.
- `WriteToFile` writes atomically and creates files with mode 0644.

New in v0.2.0: `WriteTo`, `WriteFormattedTo`, `Bdo.Dir` and `StyleElement.Err`.

## Releases

Release notes for every version are in the [changelog](CHANGELOG.md) and on the [GitHub releases page](https://github.com/lechefran/rephtml/releases). Versions follow [semantic versioning](https://semver.org/), and the Go module proxy serves each tagged version, so `go get github.com/lechefran/rephtml@v0.2.0` pins an exact release.

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
- [X] `<ol>` - Ordered list (`Start` accepts zero and negative numbers)
- [X] `<ul>` - Unordered list
- [X] `<menu>` - Menu list
- [X] `<li>` - List item (`Value` accepts zero and negative numbers)
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
- [X] `<bdo>` - Bidirectional override (set the direction with `Dir`)
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
