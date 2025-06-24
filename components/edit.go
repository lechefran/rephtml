package rephtml

import "bytes"

// Del represents the HTML del element for marking deleted text
type Del struct {
	buf      bytes.Buffer
	style    map[string]string
	contents [][]byte
	ttrack   int
	cite     string
	datetime string
}

// NewDel creates a new Del element
func NewDel() *Del {
	return &Del{
		style: make(map[string]string),
	}
}

// Bytes returns the buffer contents
func (d *Del) Bytes() []byte {
	return d.buf.Bytes()
}

// Prepare builds the HTML for the del element
func (d *Del) Prepare() {
	d.buf.WriteString("<del")
	
	if d.cite != "" {
		d.buf.WriteString(" cite=\"" + d.cite + "\"")
	}
	
	if d.datetime != "" {
		d.buf.WriteString(" datetime=\"" + d.datetime + "\"")
	}
	
	if len(d.style) != 0 {
		idx := 0
		d.buf.WriteString(" style=\"")
		for k, v := range d.style {
			d.buf.WriteString(k + ": " + v + ";")
			if idx != len(d.style)-1 {
				d.buf.WriteByte(' ')
			}
			idx++
		}
		d.buf.WriteString("\"")
	}
	
	d.buf.WriteByte('>')
	
	for _, content := range d.contents {
		d.buf.Write(content)
	}
	
	d.buf.WriteString("</del>")
}

// Add adds content to the del element
func (d *Del) Add(e Element) *Del {
	if e != nil {
		e.Prepare()
		d.contents = append(d.contents, e.Bytes())
	}
	return d
}

// Cite sets the cite attribute
func (d *Del) Cite(url string) *Del {
	d.cite = url
	return d
}

// Datetime sets the datetime attribute
func (d *Del) Datetime(datetime string) *Del {
	d.datetime = datetime
	return d
}

// AddStyle adds a single CSS property
func (d *Del) AddStyle(k, v string) *Del {
	d.style[k] = v
	return d
}

// AddStyles adds multiple CSS properties
func (d *Del) AddStyles(m map[string]string) *Del {
	for k, v := range m {
		d.style[k] = v
	}
	return d
}

// Style replaces all styles
func (d *Del) Style(m map[string]string) *Del {
	d.style = make(map[string]string)
	for k, v := range m {
		d.style[k] = v
	}
	return d
}

// Ins represents the HTML ins element for marking inserted text
type Ins struct {
	buf      bytes.Buffer
	style    map[string]string
	contents [][]byte
	ttrack   int
	cite     string
	datetime string
}

// NewIns creates a new Ins element
func NewIns() *Ins {
	return &Ins{
		style: make(map[string]string),
	}
}

// Bytes returns the buffer contents
func (i *Ins) Bytes() []byte {
	return i.buf.Bytes()
}

// Prepare builds the HTML for the ins element
func (i *Ins) Prepare() {
	i.buf.WriteString("<ins")
	
	if i.cite != "" {
		i.buf.WriteString(" cite=\"" + i.cite + "\"")
	}
	
	if i.datetime != "" {
		i.buf.WriteString(" datetime=\"" + i.datetime + "\"")
	}
	
	if len(i.style) != 0 {
		idx := 0
		i.buf.WriteString(" style=\"")
		for k, v := range i.style {
			i.buf.WriteString(k + ": " + v + ";")
			if idx != len(i.style)-1 {
				i.buf.WriteByte(' ')
			}
			idx++
		}
		i.buf.WriteString("\"")
	}
	
	i.buf.WriteByte('>')
	
	for _, content := range i.contents {
		i.buf.Write(content)
	}
	
	i.buf.WriteString("</ins>")
}

// Add adds content to the ins element
func (i *Ins) Add(e Element) *Ins {
	if e != nil {
		e.Prepare()
		i.contents = append(i.contents, e.Bytes())
	}
	return i
}

// Cite sets the cite attribute
func (i *Ins) Cite(url string) *Ins {
	i.cite = url
	return i
}

// Datetime sets the datetime attribute
func (i *Ins) Datetime(datetime string) *Ins {
	i.datetime = datetime
	return i
}

// AddStyle adds a single CSS property
func (i *Ins) AddStyle(k, v string) *Ins {
	i.style[k] = v
	return i
}

// AddStyles adds multiple CSS properties
func (i *Ins) AddStyles(m map[string]string) *Ins {
	for k, v := range m {
		i.style[k] = v
	}
	return i
}

// Style replaces all styles
func (i *Ins) Style(m map[string]string) *Ins {
	i.style = make(map[string]string)
	for k, v := range m {
		i.style[k] = v
	}
	return i
}
