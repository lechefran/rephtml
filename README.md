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

**Total Elements:** 116 HTML elements (excluding deprecated ones)