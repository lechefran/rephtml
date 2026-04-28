package rephtml

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
