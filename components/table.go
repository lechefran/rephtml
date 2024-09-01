package rephtml

import "bytes"

type Table struct {
	buf     bytes.Buffer
	class   [][]byte
	headers [][]byte
	id      []byte
	rows    [][]byte
	style   [][]byte
}
