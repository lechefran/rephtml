package rephtml

// Element Interface
/*
Holds HTML element information that can be
appended to the HTML document's body
*/
type Element interface {
	Bytes() []byte
	Prepare()
}
