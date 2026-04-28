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

// preparedElement is the internal rendering contract used by buffer-backed
// elements. Concrete element structs satisfy both Element and preparedElement:
// public callers use Render/HTML, while the package uses Prepare/Bytes behind
// those methods.
type preparedElement interface {
	Prepare()
	Bytes() []byte
}
