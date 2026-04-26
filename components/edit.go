package rephtml

import "bytes"

// Del represents the HTML del element for marking deleted text
type Del struct {
	buf      bytes.Buffer
	style    map[string]string
	contents []Element
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
	return cloneBytes(d.buf.Bytes())
}

// IsBodyElement implements BodyElement interface
func (d *Del) IsBodyElement() {}

// Prepare builds the HTML for the del element
func (d *Del) Prepare() {
	d.buf.Reset()
	d.buf.WriteString("<del")
	if d.cite != "" {
		writeAttr(&d.buf, "cite", d.cite)
	}
	if d.datetime != "" {
		writeAttr(&d.buf, "datetime", d.datetime)
	}
	if len(d.style) != 0 {
		parseStyle(&d.buf, d.style)
	}
	d.buf.WriteByte('>')

	writeElements(&d.buf, d.contents)
	d.buf.WriteString("</del>")
}

// Add adds content to the del element
func (d *Del) Add(e Element) *Del {
	if e != nil {
		d.contents = appendElement(d.contents, e)
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
	d.style = cloneStyleMap(m)
	return d
}

// Ins represents the HTML ins element for marking inserted text
type Ins struct {
	buf      bytes.Buffer
	style    map[string]string
	contents []Element
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
	return cloneBytes(i.buf.Bytes())
}

// IsBodyElement implements BodyElement interface
func (i *Ins) IsBodyElement() {}

// Prepare builds the HTML for the ins element
func (i *Ins) Prepare() {
	i.buf.Reset()
	i.buf.WriteString("<ins")

	if i.cite != "" {
		writeAttr(&i.buf, "cite", i.cite)
	}

	if i.datetime != "" {
		writeAttr(&i.buf, "datetime", i.datetime)
	}

	if len(i.style) != 0 {
		parseStyle(&i.buf, i.style)
	}

	i.buf.WriteByte('>')

	writeElements(&i.buf, i.contents)

	i.buf.WriteString("</ins>")
}

// Add adds content to the ins element
func (i *Ins) Add(e Element) *Ins {
	if e != nil {
		i.contents = appendElement(i.contents, e)
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
	i.style = cloneStyleMap(m)
	return i
}
