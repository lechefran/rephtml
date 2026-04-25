package rephtml

import "bytes"

type Canvas struct {
	buf      bytes.Buffer
	style    map[string]string
	contents []Element
	width    string
	height   string
}

func NewCanvas() *Canvas {
	return &Canvas{
		style: make(map[string]string),
	}
}

func (c *Canvas) AddStyle(k, v string) *Canvas {
	c.style[k] = v
	return c
}

func (c *Canvas) AddStyles(m map[string]string) *Canvas {
	for k, v := range m {
		c.style[k] = v
	}
	return c
}

func (c *Canvas) Style(m map[string]string) *Canvas {
	c.style = m
	return c
}

func (c *Canvas) Add(e Element) *Canvas {
	c.contents = appendElement(c.contents, e)
	return c
}

func (c *Canvas) Width(width string) *Canvas {
	c.width = width
	return c
}

func (c *Canvas) Height(height string) *Canvas {
	c.height = height
	return c
}

func (c *Canvas) Bytes() []byte {
	return cloneBytes(c.buf.Bytes())
}

// IsBodyElement implements BodyElement interface
func (c *Canvas) IsBodyElement() {}

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

type Noscript struct {
	buf      bytes.Buffer
	style    map[string]string
	contents []Element
}

func NewNoscript() *Noscript {
	return &Noscript{
		style: make(map[string]string),
	}
}

// IsHeadElement implements HeadElement interface
func (n *Noscript) IsHeadElement() {}

// IsBodyElement implements BodyElement interface
func (n *Noscript) IsBodyElement() {}

func (n *Noscript) AddStyle(k, v string) *Noscript {
	n.style[k] = v
	return n
}

func (n *Noscript) AddStyles(m map[string]string) *Noscript {
	for k, v := range m {
		n.style[k] = v
	}
	return n
}

func (n *Noscript) Style(m map[string]string) *Noscript {
	n.style = m
	return n
}

func (n *Noscript) Add(e Element) *Noscript {
	n.contents = appendElement(n.contents, e)
	return n
}

func (n *Noscript) Bytes() []byte {
	return cloneBytes(n.buf.Bytes())
}

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

func NewScript() *Script {
	return &Script{
		style: make(map[string]string),
	}
}

// IsHeadElement implements HeadElement interface
func (s *Script) IsHeadElement() {}

// IsBodyElement implements BodyElement interface
func (s *Script) IsBodyElement() {}

func (s *Script) AddStyle(k, v string) *Script {
	s.style[k] = v
	return s
}

func (s *Script) AddStyles(m map[string]string) *Script {
	for k, v := range m {
		s.style[k] = v
	}
	return s
}

func (s *Script) Style(m map[string]string) *Script {
	s.style = m
	return s
}

func (s *Script) Src(src string) *Script {
	s.src = src
	return s
}

func (s *Script) Type(scriptType string) *Script {
	s.scriptType = scriptType
	return s
}

func (s *Script) Async(async bool) *Script {
	s.async = async
	return s
}

func (s *Script) Defer(deferScript bool) *Script {
	s.deferScript = deferScript
	return s
}

func (s *Script) Crossorigin(crossorigin string) *Script {
	s.crossOrigin = crossorigin
	return s
}

func (s *Script) Integrity(integrity string) *Script {
	s.integrity = integrity
	return s
}

func (s *Script) Nomodule(nomodule bool) *Script {
	s.noModule = nomodule
	return s
}

func (s *Script) Referrerpolicy(referrerpolicy string) *Script {
	s.referrerPolicy = referrerpolicy
	return s
}

func (s *Script) Text(text string) *Script {
	s.text = text
	return s
}

func (s *Script) Bytes() []byte {
	return cloneBytes(s.buf.Bytes())
}

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
