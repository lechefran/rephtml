package rephtml

import "bytes"

// Details represents the HTML details element for creating a disclosure widget
type Details struct {
	buf      bytes.Buffer
	style    map[string]string
	contents [][]byte
	ttrack   int
	open     bool
}

// NewDetails creates a new Details element
func NewDetails() *Details {
	return &Details{
		style: make(map[string]string),
	}
}

// Bytes returns the buffer contents
func (d *Details) Bytes() []byte {
	return d.buf.Bytes()
}

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
	
	for _, content := range d.contents {
		d.buf.Write(content)
	}
	
	d.buf.WriteString("</details>")
}

// Add adds content to the details element
func (d *Details) Add(e Element) *Details {
	if e != nil {
		e.Prepare()
		d.contents = append(d.contents, e.Bytes())
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
func (d *Details) AddStyles(m map[string]string) *Details {
	for k, v := range m {
		d.style[k] = v
	}
	return d
}

// Style replaces all styles
func (d *Details) Style(m map[string]string) *Details {
	d.style = make(map[string]string)
	for k, v := range m {
		d.style[k] = v
	}
	return d
}

// Dialog represents the HTML dialog element for creating modal dialogs
type Dialog struct {
	buf      bytes.Buffer
	style    map[string]string
	contents [][]byte
	ttrack   int
	open     bool
}

// NewDialog creates a new Dialog element
func NewDialog() *Dialog {
	return &Dialog{
		style: make(map[string]string),
	}
}

// Bytes returns the buffer contents
func (d *Dialog) Bytes() []byte {
	return d.buf.Bytes()
}

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
	
	for _, content := range d.contents {
		d.buf.Write(content)
	}
	
	d.buf.WriteString("</dialog>")
}

// Add adds content to the dialog element
func (d *Dialog) Add(e Element) *Dialog {
	if e != nil {
		e.Prepare()
		d.contents = append(d.contents, e.Bytes())
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
func (d *Dialog) AddStyles(m map[string]string) *Dialog {
	for k, v := range m {
		d.style[k] = v
	}
	return d
}

// Style replaces all styles
func (d *Dialog) Style(m map[string]string) *Dialog {
	d.style = make(map[string]string)
	for k, v := range m {
		d.style[k] = v
	}
	return d
}

// Summary represents the HTML summary element for details disclosure summary
type Summary struct {
	buf      bytes.Buffer
	style    map[string]string
	contents [][]byte
	ttrack   int
}

// NewSummary creates a new Summary element
func NewSummary() *Summary {
	return &Summary{
		style: make(map[string]string),
	}
}

// Bytes returns the buffer contents
func (s *Summary) Bytes() []byte {
	return s.buf.Bytes()
}

// Prepare builds the HTML for the summary element
func (s *Summary) Prepare() {
	s.buf.Reset()
	s.buf.WriteString("<summary")
	
	if len(s.style) != 0 {
		parseStyle(&s.buf, s.style)
	}
	
	s.buf.WriteByte('>')
	
	for _, content := range s.contents {
		s.buf.Write(content)
	}
	
	s.buf.WriteString("</summary>")
}

// Add adds content to the summary element
func (s *Summary) Add(e Element) *Summary {
	if e != nil {
		e.Prepare()
		s.contents = append(s.contents, e.Bytes())
	}
	return s
}

// AddStyle adds a single CSS property
func (s *Summary) AddStyle(k, v string) *Summary {
	s.style[k] = v
	return s
}

// AddStyles adds multiple CSS properties
func (s *Summary) AddStyles(m map[string]string) *Summary {
	for k, v := range m {
		s.style[k] = v
	}
	return s
}

// Style replaces all styles
func (s *Summary) Style(m map[string]string) *Summary {
	s.style = make(map[string]string)
	for k, v := range m {
		s.style[k] = v
	}
	return s
}
