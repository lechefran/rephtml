package rephtml

import "bytes"

// Math represents the Math component or supporting type.
type Math struct {
	buf      bytes.Buffer
	style    map[string]string
	contents []Element
	display  string
	xmlns    string
}

// NewMath creates a new Math component.
func NewMath() *Math {
	return &Math{
		style: make(map[string]string),
		xmlns: "http://www.w3.org/1998/Math/MathML",
	}
}

// AddStyle adds one inline CSS declaration to the Math component.
func (m *Math) AddStyle(k, v string) *Math {
	m.style[k] = v
	return m
}

// AddStyles adds multiple inline CSS declarations to the Math component.
func (m *Math) AddStyles(ms map[string]string) *Math {
	for k, v := range ms {
		m.style[k] = v
	}
	return m
}

// Style replaces the inline CSS declarations on the Math component.
func (m *Math) Style(ms map[string]string) *Math {
	m.style = cloneStyleMap(ms)
	return m
}

// Add appends child content to the Math component.
func (m *Math) Add(e Element) *Math {
	m.contents = appendElement(m.contents, e)
	return m
}

// Display sets the display value on the Math component.
func (m *Math) Display(display string) *Math {
	m.display = display
	return m
}

// Xmlns sets the xmlns value on the Math component.
func (m *Math) Xmlns(xmlns string) *Math {
	m.xmlns = xmlns
	return m
}

// Bytes returns a defensive copy of the rendered Math bytes.
func (m *Math) Bytes() []byte {
	return cloneBytes(m.buf.Bytes())
}

// IsBodyElement implements BodyElement interface
func (m *Math) IsBodyElement() {}

// Prepare renders the Math component into its internal buffer.
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

// Svg represents the Svg component or supporting type.
type Svg struct {
	buf                 bytes.Buffer
	style               map[string]string
	contents            []Element
	width               string
	height              string
	viewBox             string
	xmlns               string
	version             string
	baseProfile         string
	preserveAspectRatio string
}

// NewSvg creates a new Svg component.
func NewSvg() *Svg {
	return &Svg{
		style: make(map[string]string),
		xmlns: "http://www.w3.org/2000/svg",
	}
}

// AddStyle adds one inline CSS declaration to the Svg component.
func (s *Svg) AddStyle(k, v string) *Svg {
	s.style[k] = v
	return s
}

// AddStyles adds multiple inline CSS declarations to the Svg component.
func (s *Svg) AddStyles(m map[string]string) *Svg {
	for k, v := range m {
		s.style[k] = v
	}
	return s
}

// Style replaces the inline CSS declarations on the Svg component.
func (s *Svg) Style(m map[string]string) *Svg {
	s.style = cloneStyleMap(m)
	return s
}

// Add appends child content to the Svg component.
func (s *Svg) Add(e Element) *Svg {
	s.contents = appendElement(s.contents, e)
	return s
}

// Width sets the width value on the Svg component.
func (s *Svg) Width(width string) *Svg {
	s.width = width
	return s
}

// Height sets the height value on the Svg component.
func (s *Svg) Height(height string) *Svg {
	s.height = height
	return s
}

// ViewBox sets the viewbox value on the Svg component.
func (s *Svg) ViewBox(viewBox string) *Svg {
	s.viewBox = viewBox
	return s
}

// Xmlns sets the xmlns value on the Svg component.
func (s *Svg) Xmlns(xmlns string) *Svg {
	s.xmlns = xmlns
	return s
}

// Version sets the version value on the Svg component.
func (s *Svg) Version(version string) *Svg {
	s.version = version
	return s
}

// BaseProfile sets the baseprofile value on the Svg component.
func (s *Svg) BaseProfile(baseProfile string) *Svg {
	s.baseProfile = baseProfile
	return s
}

// PreserveAspectRatio sets the preserveaspectratio value on the Svg component.
func (s *Svg) PreserveAspectRatio(preserveAspectRatio string) *Svg {
	s.preserveAspectRatio = preserveAspectRatio
	return s
}

// Bytes returns a defensive copy of the rendered Svg bytes.
func (s *Svg) Bytes() []byte {
	return cloneBytes(s.buf.Bytes())
}

// IsBodyElement implements BodyElement interface
func (s *Svg) IsBodyElement() {}

// Prepare renders the Svg component into its internal buffer.
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
