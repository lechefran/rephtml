// Package rephtml builds HTML documents from typed, composable Go values.
//
// Every HTML element is a struct created by a New constructor. Setters return
// the element, so a tree reads top to bottom as it is built:
//
//	page := rephtml.NewHtmlFile().Lang("en")
//	page.AddToHead(rephtml.NewTitle().Text("Sales Report"))
//	page.AddToBody(rephtml.NewMain().
//		Add(rephtml.NewH1().Text("Sales Report")).
//		Add(rephtml.NewP().Text("Quarterly regional performance")))
//
// # Chaining
//
// Setters that every element shares, such as Id, Class, AddStyle and Add, are
// promoted from an internal base type. Their signatures show a return type
// named Self. Self is always the receiver's own pointer type, so calling
// AddClass on a [*Div] returns a [*Div], and the chain keeps its concrete type.
//
// # Rendering
//
// A regular element renders itself through Render or HTML. The document root,
// [HtmlFile], returns an error as well, because adding content that does not
// belong in the head or body is recorded and reported when the document is
// rendered rather than hidden; see [HtmlFile.Err]. HtmlFile offers compact and
// formatted output as bytes ([HtmlFile.Render], [HtmlFile.RenderFormatted]), as
// strings ([HtmlFile.RenderString], [HtmlFile.RenderFormattedString]), to an
// [io.Writer] ([HtmlFile.WriteTo], [HtmlFile.WriteFormattedTo]) and to a file
// ([HtmlFile.WriteToFile]).
//
// # Escaping and URLs
//
// Text and attribute values are HTML-escaped. URL attributes are also filtered
// by scheme, so a javascript: URL is replaced with [BlockedURL]; [IsSafeURL]
// reports whether a value would be kept. Script and style text is the
// deliberate exception: it is written as raw JavaScript or CSS, with only the
// sequence that would end the element rewritten, so do not pass untrusted input
// as code.
//
// # Concurrency
//
// Rendering never mutates the element tree, so a tree that is fully built can
// be rendered from any number of goroutines. Building is not synchronized:
// finish calling setters before sharing the tree.
package rephtml
