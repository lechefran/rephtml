package rephtml

import "bytes"

// Details represents the HTML details element for creating a disclosure widget
type Details struct {
	buf      bytes.Buffer
	style    StyleMap
	contents []Element
	open     bool
}

// NewDetails creates a new Details element
func NewDetails() *Details {
	return &Details{
		style: make(StyleMap),
	}
}

// Bytes returns the buffer contents
func (d *Details) Bytes() []byte {
	return cloneBytes(d.buf.Bytes())
}

// IsBodyElement implements BodyElement interface
func (d *Details) IsBodyElement() {}

// Prepare builds the HTML for the details element
func (d *Details) Prepare() {
	d.buf.Reset()
	d.buf.WriteString("<details")

	if d.open {
		d.buf.WriteString(" open")
	}

	if len(d.style) != 0 {
		parseStyle(&d.buf, d.style)
	}

	d.buf.WriteByte('>')

	writeElements(&d.buf, d.contents)

	d.buf.WriteString("</details>")
}

// Add adds content to the details element
func (d *Details) Add(e Element) *Details {
	if e != nil {
		d.contents = appendElement(d.contents, e)
	}
	return d
}

// Open sets the open attribute
func (d *Details) Open(open bool) *Details {
	d.open = open
	return d
}

// AddStyle adds a single CSS property
func (d *Details) AddStyle(k, v string) *Details {
	d.style[k] = v
	return d
}

// AddStyles adds multiple CSS properties
func (d *Details) AddStyles(m StyleMap) *Details {
	for k, v := range m {
		d.style[k] = v
	}
	return d
}

// Style replaces all styles
func (d *Details) Style(m StyleMap) *Details {
	d.style = cloneStyleMap(m)
	return d
}

// Dialog represents the HTML dialog element for creating modal dialogs
type Dialog struct {
	buf      bytes.Buffer
	style    StyleMap
	contents []Element
	open     bool
}

// NewDialog creates a new Dialog element
func NewDialog() *Dialog {
	return &Dialog{
		style: make(StyleMap),
	}
}

// Bytes returns the buffer contents
func (d *Dialog) Bytes() []byte {
	return cloneBytes(d.buf.Bytes())
}

// IsBodyElement implements BodyElement interface
func (d *Dialog) IsBodyElement() {}

// Prepare builds the HTML for the dialog element
func (d *Dialog) Prepare() {
	d.buf.Reset()
	d.buf.WriteString("<dialog")

	if d.open {
		d.buf.WriteString(" open")
	}

	if len(d.style) != 0 {
		parseStyle(&d.buf, d.style)
	}

	d.buf.WriteByte('>')

	writeElements(&d.buf, d.contents)

	d.buf.WriteString("</dialog>")
}

// Add adds content to the dialog element
func (d *Dialog) Add(e Element) *Dialog {
	if e != nil {
		d.contents = appendElement(d.contents, e)
	}
	return d
}

// Open sets the open attribute
func (d *Dialog) Open(open bool) *Dialog {
	d.open = open
	return d
}

// AddStyle adds a single CSS property
func (d *Dialog) AddStyle(k, v string) *Dialog {
	d.style[k] = v
	return d
}

// AddStyles adds multiple CSS properties
func (d *Dialog) AddStyles(m StyleMap) *Dialog {
	for k, v := range m {
		d.style[k] = v
	}
	return d
}

// Style replaces all styles
func (d *Dialog) Style(m StyleMap) *Dialog {
	d.style = cloneStyleMap(m)
	return d
}

// Summary represents the HTML summary element for details disclosure summary
type Summary struct {
	buf      bytes.Buffer
	style    StyleMap
	contents []Element
}

// NewSummary creates a new Summary element
func NewSummary() *Summary {
	return &Summary{
		style: make(StyleMap),
	}
}

// Bytes returns the buffer contents
func (s *Summary) Bytes() []byte {
	return cloneBytes(s.buf.Bytes())
}

// IsBodyElement implements BodyElement interface
func (s *Summary) IsBodyElement() {}

// Prepare builds the HTML for the summary element
func (s *Summary) Prepare() {
	s.buf.Reset()
	s.buf.WriteString("<summary")

	if len(s.style) != 0 {
		parseStyle(&s.buf, s.style)
	}

	s.buf.WriteByte('>')

	writeElements(&s.buf, s.contents)

	s.buf.WriteString("</summary>")
}

// Add adds content to the summary element
func (s *Summary) Add(e Element) *Summary {
	if e != nil {
		s.contents = appendElement(s.contents, e)
	}
	return s
}

// AddStyle adds a single CSS property
func (s *Summary) AddStyle(k, v string) *Summary {
	s.style[k] = v
	return s
}

// AddStyles adds multiple CSS properties
func (s *Summary) AddStyles(m StyleMap) *Summary {
	for k, v := range m {
		s.style[k] = v
	}
	return s
}

// Style replaces all styles
func (s *Summary) Style(m StyleMap) *Summary {
	s.style = cloneStyleMap(m)
	return s
}
