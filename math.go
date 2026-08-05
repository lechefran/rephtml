package rephtml

// Math represents a MathML math element.
type Math struct {
	bodyElement
	contentNode[*Math]
	display string
	xmlns   string
}

// NewMath creates a MathML math element with the default MathML namespace.
func NewMath() *Math {
	v := &Math{}
	v.init(v)
	v.xmlns = "http://www.w3.org/1998/Math/MathML"
	return v
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

// Prepare renders the math element into its internal buffer.
func (m *Math) prepare() {
	tg := openTag(&m.buf, "math")
	tg.attr("xmlns", m.xmlns)
	tg.attr("display", m.display)
	tg.styleAttr(m.style)
	tg.children(m.contents)
}

// Svg represents an inline SVG element.
type Svg struct {
	bodyElement
	contentNode[*Svg]
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
	v := &Svg{}
	v.init(v)
	v.xmlns = "http://www.w3.org/2000/svg"
	return v
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

// Prepare renders the SVG element into its internal buffer.
func (s *Svg) prepare() {
	tg := openTag(&s.buf, "svg")
	tg.attr("xmlns", s.xmlns)
	tg.attr("width", s.width)
	tg.attr("height", s.height)
	tg.attr("viewBox", s.viewBox)
	tg.attr("version", s.version)
	tg.attr("baseProfile", s.baseProfile)
	tg.attr("preserveAspectRatio", s.preserveAspectRatio)
	tg.styleAttr(s.style)
	tg.children(s.contents)
}
