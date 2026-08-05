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
// public callers use Render/HTML, while the package uses prepare/rawBytes
// behind those methods.
//
// Every element declares its own prepare; rawBytes comes from the embedded
// base. Because the contract is unexported, rendering into a parent's buffer
// can skip the defensive copy that Render owes its callers.
type preparedElement interface {
	prepare()
	rawBytes() []byte
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
