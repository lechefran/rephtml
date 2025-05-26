package rephtml

// Elements Interface
/*
Holds HTML element information that can be
appended to the HTML document's body
*/
type Elements interface {
	Bytes() []byte
	Prepare()
}
