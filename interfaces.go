package rephtml

import "bytes"

// Element is renderable HTML content that can be appended to other elements.
//
// Render and HTML prepare the element internally before returning output, so
// callers do not need to call Prepare before reading rendered bytes.
type Element interface {
	Render() []byte
	HTML() string
}

// HeadElement is HTML content that can be appended to a document head.
type HeadElement interface {
	Element
	IsHeadElement()
}

// BodyElement is HTML content that can be appended to a document body.
type BodyElement interface {
	Element
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
