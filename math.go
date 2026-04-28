package rephtml

import "bytes"

// Math represents a MathML math element.
type Math struct {
	buf      bytes.Buffer
	style    StyleMap
	contents []Element
	display  string
	xmlns    string
}

// NewMath creates a MathML math element with the default MathML namespace.
func NewMath() *Math {
	return &Math{
		style: make(StyleMap),
		xmlns: "http://www.w3.org/1998/Math/MathML",
	}
}

// AddStyle adds one inline CSS declaration to the math element.
func (m *Math) AddStyle(k, v string) *Math {
	m.style[k] = v
	return m
}

// AddStyles adds multiple inline CSS declarations to the math element.
func (m *Math) AddStyles(ms StyleMap) *Math {
	for k, v := range ms {
		m.style[k] = v
	}
	return m
}

// Style replaces the inline CSS declarations on the math element.
func (m *Math) Style(ms StyleMap) *Math {
	m.style = cloneStyleMap(ms)
	return m
}

// Add appends child MathML content.
func (m *Math) Add(e Element) *Math {
	m.contents = appendElement(m.contents, e)
	return m
}

// Display sets the MathML display attribute.
func (m *Math) Display(display string) *Math {
	m.display = display
	return m
}

// Xmlns replaces the MathML namespace attribute.
func (m *Math) Xmlns(xmlns string) *Math {
	m.xmlns = xmlns
	return m
}

// Bytes returns a defensive copy of the rendered math bytes.
func (m *Math) Bytes() []byte {
	return cloneBytes(m.buf.Bytes())
}

// Render returns freshly prepared Math HTML bytes.
func (m *Math) Render() []byte {
	return renderPrepared(m)
}

// HTML returns freshly prepared Math HTML as a string.
func (m *Math) HTML() string {
	return htmlPrepared(m)
}

// String returns freshly prepared Math HTML as a string.
func (m *Math) String() string {
	return m.HTML()
}

// IsBodyElement implements BodyElement interface
func (m *Math) IsBodyElement() {}

// Prepare renders the math element into its internal buffer.
func (m *Math) Prepare() {
	m.buf.Reset()
	m.buf.WriteString("<math")
	if m.xmlns != "" {
		writeAttr(&m.buf, "xmlns", m.xmlns)
	}
	if m.display != "" {
		writeAttr(&m.buf, "display", m.display)
	}
	if len(m.style) != 0 {
		parseStyle(&m.buf, m.style)
	}
	m.buf.WriteByte('>')

	writeElements(&m.buf, m.contents)
	m.buf.WriteString("</math>")
}

// Svg represents an inline SVG element.
type Svg struct {
	buf                 bytes.Buffer
	style               StyleMap
	contents            []Element
	width               string
	height              string
	viewBox             string
	xmlns               string
	version             string
	baseProfile         string
	preserveAspectRatio string
}

// NewSvg creates an SVG element with the default SVG namespace.
func NewSvg() *Svg {
	return &Svg{
		style: make(StyleMap),
		xmlns: "http://www.w3.org/2000/svg",
	}
}

// AddStyle adds one inline CSS declaration to the SVG element.
func (s *Svg) AddStyle(k, v string) *Svg {
	s.style[k] = v
	return s
}

// AddStyles adds multiple inline CSS declarations to the SVG element.
func (s *Svg) AddStyles(m StyleMap) *Svg {
	for k, v := range m {
		s.style[k] = v
	}
	return s
}

// Style replaces the inline CSS declarations on the SVG element.
func (s *Svg) Style(m StyleMap) *Svg {
	s.style = cloneStyleMap(m)
	return s
}

// Add appends child SVG content.
func (s *Svg) Add(e Element) *Svg {
	s.contents = appendElement(s.contents, e)
	return s
}

// Width sets the SVG width attribute.
func (s *Svg) Width(width string) *Svg {
	s.width = width
	return s
}

// Height sets the SVG height attribute.
func (s *Svg) Height(height string) *Svg {
	s.height = height
	return s
}

// ViewBox sets the SVG viewBox attribute.
func (s *Svg) ViewBox(viewBox string) *Svg {
	s.viewBox = viewBox
	return s
}

// Xmlns replaces the SVG namespace attribute.
func (s *Svg) Xmlns(xmlns string) *Svg {
	s.xmlns = xmlns
	return s
}

// Version sets the SVG version attribute.
func (s *Svg) Version(version string) *Svg {
	s.version = version
	return s
}

// BaseProfile sets the SVG baseProfile attribute.
func (s *Svg) BaseProfile(baseProfile string) *Svg {
	s.baseProfile = baseProfile
	return s
}

// PreserveAspectRatio sets the SVG preserveAspectRatio attribute.
func (s *Svg) PreserveAspectRatio(preserveAspectRatio string) *Svg {
	s.preserveAspectRatio = preserveAspectRatio
	return s
}

// Bytes returns a defensive copy of the rendered SVG bytes.
func (s *Svg) Bytes() []byte {
	return cloneBytes(s.buf.Bytes())
}

// Render returns freshly prepared Svg HTML bytes.
func (s *Svg) Render() []byte {
	return renderPrepared(s)
}

// HTML returns freshly prepared Svg HTML as a string.
func (s *Svg) HTML() string {
	return htmlPrepared(s)
}

// String returns freshly prepared Svg HTML as a string.
func (s *Svg) String() string {
	return s.HTML()
}

// IsBodyElement implements BodyElement interface
func (s *Svg) IsBodyElement() {}

// Prepare renders the SVG element into its internal buffer.
func (s *Svg) Prepare() {
	s.buf.Reset()
	s.buf.WriteString("<svg")
	if s.xmlns != "" {
		writeAttr(&s.buf, "xmlns", s.xmlns)
	}
	if s.width != "" {
		writeAttr(&s.buf, "width", s.width)
	}
	if s.height != "" {
		writeAttr(&s.buf, "height", s.height)
	}
	if s.viewBox != "" {
		writeAttr(&s.buf, "viewBox", s.viewBox)
	}
	if s.version != "" {
		writeAttr(&s.buf, "version", s.version)
	}
	if s.baseProfile != "" {
		writeAttr(&s.buf, "baseProfile", s.baseProfile)
	}
	if s.preserveAspectRatio != "" {
		writeAttr(&s.buf, "preserveAspectRatio", s.preserveAspectRatio)
	}
	if len(s.style) != 0 {
		parseStyle(&s.buf, s.style)
	}
	s.buf.WriteByte('>')

	writeElements(&s.buf, s.contents)
	s.buf.WriteString("</svg>")
}
