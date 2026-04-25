package rephtml

import "bytes"

type Math struct {
	buf      bytes.Buffer
	style    map[string]string
	contents [][]byte
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
	m.contents = append(m.contents, e.Bytes())
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
	return m.buf.Bytes()
}

// IsBodyElement implements BodyElement interface
func (m *Math) IsBodyElement() {}

func (m *Math) Prepare() {
	m.buf.Reset()
	m.buf.WriteString("<math")
	if m.xmlns != "" {
		m.buf.WriteString(" xmlns=\"" + m.xmlns + "\"")
	}
	if m.display != "" {
		m.buf.WriteString(" display=\"" + m.display + "\"")
	}
	if len(m.style) != 0 {
		parseStyle(&m.buf, m.style)
	}
	m.buf.WriteByte('>')

	for _, content := range m.contents {
		m.buf.Write(content)
	}
	m.buf.WriteString("</math>")
}

type Svg struct {
	buf       bytes.Buffer
	style     map[string]string
	contents  [][]byte
	width     string
	height    string
	viewBox   string
	xmlns     string
	version   string
	baseProfile string
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
	s.contents = append(s.contents, e.Bytes())
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
	return s.buf.Bytes()
}

// IsBodyElement implements BodyElement interface
func (s *Svg) IsBodyElement() {}

func (s *Svg) Prepare() {
	s.buf.Reset()
	s.buf.WriteString("<svg")
	if s.xmlns != "" {
		s.buf.WriteString(" xmlns=\"" + s.xmlns + "\"")
	}
	if s.width != "" {
		s.buf.WriteString(" width=\"" + s.width + "\"")
	}
	if s.height != "" {
		s.buf.WriteString(" height=\"" + s.height + "\"")
	}
	if s.viewBox != "" {
		s.buf.WriteString(" viewBox=\"" + s.viewBox + "\"")
	}
	if s.version != "" {
		s.buf.WriteString(" version=\"" + s.version + "\"")
	}
	if s.baseProfile != "" {
		s.buf.WriteString(" baseProfile=\"" + s.baseProfile + "\"")
	}
	if s.preserveAspectRatio != "" {
		s.buf.WriteString(" preserveAspectRatio=\"" + s.preserveAspectRatio + "\"")
	}
	if len(s.style) != 0 {
		parseStyle(&s.buf, s.style)
	}
	s.buf.WriteByte('>')

	for _, content := range s.contents {
		s.buf.Write(content)
	}
	s.buf.WriteString("</svg>")
}
