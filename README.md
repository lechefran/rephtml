# rephtml
A simple HTML file creator written in Go! rephtml supports both string and struct parameters.
## Simple Example
```
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
html.H1("Test H1 Header")
html.Table([]string{"header1", "header2", "header3"},
          [][]string{{"record1", "record2", "record3"},
          {"record4", "record5", "record6"}})
html.P("Test paragraph for testing purposes")
html.Prepare()
html.WriteToFile("report.html") // create a file named report.html in the current directory
```

## Supported Elements
### Document Structure
- [X] `<html>` - Root element
- [X] `<head>` - Document head
- [X] `<body>` - Document body
- [X] `<title>` - Document title
- [X] `<base>` - Base URL
- [ ] `<link>` - External resource link
- [ ] `<meta>` - Metadata
- [ ] `<style>` - Internal CSS

### Content Sectioning
- [ ] `<header>` - Header section
- [ ] `<nav>` - Navigation section
- [ ] `<main>` - Main content
- [ ] `<section>` - Generic section
- [ ] `<article>` - Article content
- [ ] `<aside>` - Sidebar content
- [ ] `<footer>` - Footer section
- [ ] `<address>` - Contact information
- [X] `<h1>` through `<h6>` - Headings

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
- [ ] `<area>` - Image map area
- [ ] `<audio>` - Audio content
- [ ] `<img>` - Image
- [ ] `<map>` - Image map
- [ ] `<track>` - Media track
- [ ] `<video>` - Video content

### Embedded Content
- [ ] `<embed>` - External content
- [ ] `<iframe>` - Inline frame
- [ ] `<object>` - External object
- [ ] `<picture>` - Responsive image container
- [ ] `<portal>` - Portal element
- [ ] `<source>` - Media source

### SVG and MathML
- [ ] `<svg>` - SVG graphics
- [ ] `<math>` - MathML mathematics

### Scripting
- [ ] `<canvas>` - Graphics canvas
- [ ] `<noscript>` - No script fallback
- [ ] `<script>` - Script

### Demarcating Edits
- [ ] `<del>` - Deleted text
- [ ] `<ins>` - Inserted text

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
- [ ] `<form>` - Form
- [ ] `<label>` - Form label
- [ ] `<input>` - Form input
- [ ] `<button>` - Button
- [ ] `<select>` - Selection list
- [ ] `<datalist>` - Data list options
- [ ] `<optgroup>` - Option group
- [ ] `<option>` - Option
- [ ] `<textarea>` - Text area
- [ ] `<output>` - Form output
- [ ] `<progress>` - Progress indicator
- [ ] `<meter>` - Scalar measurement
- [ ] `<fieldset>` - Form field grouping
- [ ] `<legend>` - Fieldset legend

### Interactive Elements
- [ ] `<details>` - Disclosure widget
- [ ] `<summary>` - Details summary
- [ ] `<dialog>` - Dialog box

### Web Components
- [ ] `<slot>` - Web component slot
- [ ] `<template>` - Template element

---

**Total Elements:** 121 HTML elements (excluding deprecated ones)