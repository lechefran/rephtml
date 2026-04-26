package rephtml

import "bytes"

// Canvas represents the Canvas component or supporting type.
type Canvas struct {
	buf      bytes.Buffer
	style    map[string]string
	contents []Element
	width    string
	height   string
}

// NewCanvas creates a new Canvas component.
func NewCanvas() *Canvas {
	return &Canvas{
		style: make(map[string]string),
	}
}

// AddStyle adds one inline CSS declaration to the Canvas component.
func (c *Canvas) AddStyle(k, v string) *Canvas {
	c.style[k] = v
	return c
}

// AddStyles adds multiple inline CSS declarations to the Canvas component.
func (c *Canvas) AddStyles(m map[string]string) *Canvas {
	for k, v := range m {
		c.style[k] = v
	}
	return c
}

// Style replaces the inline CSS declarations on the Canvas component.
func (c *Canvas) Style(m map[string]string) *Canvas {
	c.style = m
	return c
}

// Add appends child content to the Canvas component.
func (c *Canvas) Add(e Element) *Canvas {
	c.contents = appendElement(c.contents, e)
	return c
}

// Width sets the width value on the Canvas component.
func (c *Canvas) Width(width string) *Canvas {
	c.width = width
	return c
}

// Height sets the height value on the Canvas component.
func (c *Canvas) Height(height string) *Canvas {
	c.height = height
	return c
}

// Bytes returns a defensive copy of the rendered Canvas bytes.
func (c *Canvas) Bytes() []byte {
	return cloneBytes(c.buf.Bytes())
}

// IsBodyElement implements BodyElement interface
func (c *Canvas) IsBodyElement() {}

// Prepare renders the Canvas component into its internal buffer.
func (c *Canvas) Prepare() {
	c.buf.Reset()
	c.buf.WriteString("<canvas")
	if c.width != "" {
		writeAttr(&c.buf, "width", c.width)
	}
	if c.height != "" {
		writeAttr(&c.buf, "height", c.height)
	}
	if len(c.style) != 0 {
		parseStyle(&c.buf, c.style)
	}
	c.buf.WriteByte('>')

	writeElements(&c.buf, c.contents)
	c.buf.WriteString("</canvas>")
}

// Noscript represents the Noscript component or supporting type.
type Noscript struct {
	buf      bytes.Buffer
	style    map[string]string
	contents []Element
}

// NewNoscript creates a new Noscript component.
func NewNoscript() *Noscript {
	return &Noscript{
		style: make(map[string]string),
	}
}

// IsHeadElement implements HeadElement interface
func (n *Noscript) IsHeadElement() {}

// IsBodyElement implements BodyElement interface
func (n *Noscript) IsBodyElement() {}

// AddStyle adds one inline CSS declaration to the Noscript component.
func (n *Noscript) AddStyle(k, v string) *Noscript {
	n.style[k] = v
	return n
}

// AddStyles adds multiple inline CSS declarations to the Noscript component.
func (n *Noscript) AddStyles(m map[string]string) *Noscript {
	for k, v := range m {
		n.style[k] = v
	}
	return n
}

// Style replaces the inline CSS declarations on the Noscript component.
func (n *Noscript) Style(m map[string]string) *Noscript {
	n.style = m
	return n
}

// Add appends child content to the Noscript component.
func (n *Noscript) Add(e Element) *Noscript {
	n.contents = appendElement(n.contents, e)
	return n
}

// Bytes returns a defensive copy of the rendered Noscript bytes.
func (n *Noscript) Bytes() []byte {
	return cloneBytes(n.buf.Bytes())
}

// Prepare renders the Noscript component into its internal buffer.
func (n *Noscript) Prepare() {
	n.buf.Reset()
	n.buf.WriteString("<noscript")
	if len(n.style) != 0 {
		parseStyle(&n.buf, n.style)
	}
	n.buf.WriteByte('>')

	writeElements(&n.buf, n.contents)
	n.buf.WriteString("</noscript>")
}

// Script represents the Script component or supporting type.
type Script struct {
	buf            bytes.Buffer
	style          map[string]string
	src            string
	scriptType     string
	async          bool
	deferScript    bool
	crossOrigin    string
	integrity      string
	noModule       bool
	referrerPolicy string
	text           string
}

// NewScript creates a new Script component.
func NewScript() *Script {
	return &Script{
		style: make(map[string]string),
	}
}

// IsHeadElement implements HeadElement interface
func (s *Script) IsHeadElement() {}

// IsBodyElement implements BodyElement interface
func (s *Script) IsBodyElement() {}

// AddStyle adds one inline CSS declaration to the Script component.
func (s *Script) AddStyle(k, v string) *Script {
	s.style[k] = v
	return s
}

// AddStyles adds multiple inline CSS declarations to the Script component.
func (s *Script) AddStyles(m map[string]string) *Script {
	for k, v := range m {
		s.style[k] = v
	}
	return s
}

// Style replaces the inline CSS declarations on the Script component.
func (s *Script) Style(m map[string]string) *Script {
	s.style = m
	return s
}

// Src sets the src value on the Script component.
func (s *Script) Src(src string) *Script {
	s.src = src
	return s
}

// Type sets the type value on the Script component.
func (s *Script) Type(scriptType string) *Script {
	s.scriptType = scriptType
	return s
}

// Async sets the async value on the Script component.
func (s *Script) Async(async bool) *Script {
	s.async = async
	return s
}

// Defer sets the defer value on the Script component.
func (s *Script) Defer(deferScript bool) *Script {
	s.deferScript = deferScript
	return s
}

// Crossorigin sets the crossorigin value on the Script component.
func (s *Script) Crossorigin(crossorigin string) *Script {
	s.crossOrigin = crossorigin
	return s
}

// Integrity sets the integrity value on the Script component.
func (s *Script) Integrity(integrity string) *Script {
	s.integrity = integrity
	return s
}

// Nomodule sets the nomodule value on the Script component.
func (s *Script) Nomodule(nomodule bool) *Script {
	s.noModule = nomodule
	return s
}

// Referrerpolicy sets the referrerpolicy value on the Script component.
func (s *Script) Referrerpolicy(referrerpolicy string) *Script {
	s.referrerPolicy = referrerpolicy
	return s
}

// Text sets or appends text content on the Script component.
func (s *Script) Text(text string) *Script {
	s.text = text
	return s
}

// Bytes returns a defensive copy of the rendered Script bytes.
func (s *Script) Bytes() []byte {
	return cloneBytes(s.buf.Bytes())
}

// Prepare renders the Script component into its internal buffer.
func (s *Script) Prepare() {
	s.buf.Reset()
	s.buf.WriteString("<script")
	if s.src != "" {
		writeAttr(&s.buf, "src", s.src)
	}
	if s.scriptType != "" {
		writeAttr(&s.buf, "type", s.scriptType)
	}
	if s.async {
		s.buf.WriteString(" async")
	}
	if s.deferScript {
		s.buf.WriteString(" defer")
	}
	if s.crossOrigin != "" {
		writeAttr(&s.buf, "crossorigin", s.crossOrigin)
	}
	if s.integrity != "" {
		writeAttr(&s.buf, "integrity", s.integrity)
	}
	if s.noModule {
		s.buf.WriteString(" nomodule")
	}
	if s.referrerPolicy != "" {
		writeAttr(&s.buf, "referrerpolicy", s.referrerPolicy)
	}
	if len(s.style) != 0 {
		parseStyle(&s.buf, s.style)
	}
	s.buf.WriteByte('>')

	if s.text != "" {
		s.buf.WriteString(s.text)
	}
	s.buf.WriteString("</script>")
}
