package rephtml

// Element Interface
/*
Holds HTML element information that can be
appended to the HTML document
*/
type Element interface {
	Bytes() []byte
	Prepare()
}

// HeadElement Interface
/*
Signifies that the HTML element can be
appended to the HTML head element
*/
type HeadElement interface {
	Element
	IsHeadElement()
}

// BodyElement Interface
/*
Signifies that the HTML element can be
appended to the HTML body element
*/
type BodyElement interface {
	Element
	IsBodyElement()
}
