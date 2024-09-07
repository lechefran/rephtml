package rephtml

/*
HTML Elements Interface

Holds HTML element information that can be
appeneded to the HTML document's body
*/
type Elements interface {
	Bytes() []byte
	Prepare()
}
