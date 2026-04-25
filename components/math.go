package rephtml

import "bytes"

type Math struct {
	buf      bytes.Buffer
	style    map[string]string
	contents []Element
	display  string
	xmlns    string
}

func NewMath() *Math {
	return &Math{
		style: make(map[string]string),
		xmlns: "http://www.w3.org/1998/Math/MathML",
	}
}

func (m *Math) AddStyle(k, v string) *Math {
	m.style[k] = v
	return m
}

func (m *Math) AddStyles(ms map[string]string) *Math {
	for k, v := range ms {
		m.style[k] = v
	}
	return m
}

func (m *Math) Style(ms map[string]string) *Math {
	m.style = ms
	return m
}

func (m *Math) Add(e Element) *Math {
	m.contents = appendElement(m.contents, e)
	return m
}

func (m *Math) Display(display string) *Math {
	m.display = display
	return m
}

func (m *Math) Xmlns(xmlns string) *Math {
	m.xmlns = xmlns
	return m
}

func (m *Math) Bytes() []byte {
	return cloneBytes(m.buf.Bytes())
}

// IsBodyElement implements BodyElement interface
func (m *Math) IsBodyElement() {}

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

func NewSvg() *Svg {
	return &Svg{
		style: make(map[string]string),
		xmlns: "http://www.w3.org/2000/svg",
	}
}

func (s *Svg) AddStyle(k, v string) *Svg {
	s.style[k] = v
	return s
}

func (s *Svg) AddStyles(m map[string]string) *Svg {
	for k, v := range m {
		s.style[k] = v
	}
	return s
}

func (s *Svg) Style(m map[string]string) *Svg {
	s.style = m
	return s
}

func (s *Svg) Add(e Element) *Svg {
	s.contents = appendElement(s.contents, e)
	return s
}

func (s *Svg) Width(width string) *Svg {
	s.width = width
	return s
}

func (s *Svg) Height(height string) *Svg {
	s.height = height
	return s
}

func (s *Svg) ViewBox(viewBox string) *Svg {
	s.viewBox = viewBox
	return s
}

func (s *Svg) Xmlns(xmlns string) *Svg {
	s.xmlns = xmlns
	return s
}

func (s *Svg) Version(version string) *Svg {
	s.version = version
	return s
}

func (s *Svg) BaseProfile(baseProfile string) *Svg {
	s.baseProfile = baseProfile
	return s
}

func (s *Svg) PreserveAspectRatio(preserveAspectRatio string) *Svg {
	s.preserveAspectRatio = preserveAspectRatio
	return s
}

func (s *Svg) Bytes() []byte {
	return cloneBytes(s.buf.Bytes())
}

// IsBodyElement implements BodyElement interface
func (s *Svg) IsBodyElement() {}

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
