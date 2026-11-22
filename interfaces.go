package rephtml

import "bytes"

// Element is renderable HTML content that can be appended to other elements.
//
// Render and HTML build the output on each call and the element keeps no render
// state, so there is no separate preparation step and a built element can be
// rendered from several goroutines at once.
type Element interface {
	// Render returns the element and its children as HTML bytes.
	Render() []byte
	// HTML returns the element and its children as an HTML string.
	HTML() string
}

// HeadElement is HTML content that can be appended to a document head.
//
// The elements that may appear in a head, such as [Title], [Meta], [Link],
// [Base], [StyleElement] and [Script], implement it. [HtmlFile.AddToHead]
// rejects anything else and records the error on [HtmlFile.Err].
type HeadElement interface {
	Element
	// IsHeadElement marks the element as valid head content. It does nothing
	// and exists only so the compiler can tell head elements from the rest.
	IsHeadElement()
}

// BodyElement is HTML content that can be appended to a document body.
//
// Every element that may appear in a body, such as [Div], [P] and [Table],
// implements it. [HtmlFile.AddToBody] rejects anything else and records the
// error on [HtmlFile.Err]. A few elements, such as [Script] and [StyleElement],
// are valid in both the head and the body.
type BodyElement interface {
	Element
	// IsBodyElement marks the element as valid body content. It does nothing
	// and exists only so the compiler can tell body elements from the rest.
	IsBodyElement()
}

// preparedElement is the internal rendering contract. Concrete element structs
// satisfy both Element and preparedElement: public callers use Render/HTML,
// which allocate a buffer and hand it to renderTo.
//
// Elements write to the buffer they are given and never to storage of their
// own. That is what makes a built tree safe to render from several goroutines
// at once, and it lets a child write straight into its parent's buffer instead
// of being serialised separately and copied in.
type preparedElement interface {
	renderTo(buf *bytes.Buffer)
}

// selfElement constrains the generic element bases in element.go.
//
// It embeds comparable so a base can tell whether it was handed the concrete
// element pointer. That only happens in a New* constructor, so the check
// distinguishes a properly constructed element from a bare composite literal.
type selfElement interface {
	comparable
	preparedElement
}
