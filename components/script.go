package rephtml

import "bytes"

type Canvas struct {
	buf      bytes.Buffer
	style    map[string]string
	contents [][]byte
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

func (c *Canvas) Add(e Elements) *Canvas {
	c.contents = append(c.contents, e.Bytes())
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
	return c.buf.Bytes()
}

func (c *Canvas) Prepare() {
	c.buf.WriteString("<canvas")
	if c.width != "" {
		c.buf.WriteString(" width=\"" + c.width + "\"")
	}
	if c.height != "" {
		c.buf.WriteString(" height=\"" + c.height + "\"")
	}
	if len(c.style) != 0 {
		idx := 0
		c.buf.WriteString(" style=\"")
		for k, v := range c.style {
			c.buf.WriteString(k + ": " + v + ";")
			if idx != len(c.style)-1 {
				c.buf.WriteByte(' ')
			}
			idx++
		}
		c.buf.WriteString("\"")
	}
	c.buf.WriteByte('>')

	for _, content := range c.contents {
		c.buf.Write(content)
	}
	c.buf.WriteString("</canvas>")
}

type Noscript struct {
	buf      bytes.Buffer
	style    map[string]string
	contents [][]byte
}

func NewNoscript() *Noscript {
	return &Noscript{
		style: make(map[string]string),
	}
}

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

func (n *Noscript) Add(e Elements) *Noscript {
	n.contents = append(n.contents, e.Bytes())
	return n
}

func (n *Noscript) Bytes() []byte {
	return n.buf.Bytes()
}

func (n *Noscript) Prepare() {
	if len(n.style) != 0 {
		idx := 0
		n.buf.WriteString("<noscript style=\"")
		for k, v := range n.style {
			n.buf.WriteString(k + ": " + v + ";")
			if idx != len(n.style)-1 {
				n.buf.WriteByte(' ')
			}
			idx++
		}
		n.buf.WriteString("\">")
	} else {
		n.buf.WriteString("<noscript>")
	}

	for _, content := range n.contents {
		n.buf.Write(content)
	}
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
	return s.buf.Bytes()
}

func (s *Script) Prepare() {
	s.buf.WriteString("<script")
	if s.src != "" {
		s.buf.WriteString(" src=\"" + s.src + "\"")
	}
	if s.scriptType != "" {
		s.buf.WriteString(" type=\"" + s.scriptType + "\"")
	}
	if s.async {
		s.buf.WriteString(" async")
	}
	if s.deferScript {
		s.buf.WriteString(" defer")
	}
	if s.crossOrigin != "" {
		s.buf.WriteString(" crossorigin=\"" + s.crossOrigin + "\"")
	}
	if s.integrity != "" {
		s.buf.WriteString(" integrity=\"" + s.integrity + "\"")
	}
	if s.noModule {
		s.buf.WriteString(" nomodule")
	}
	if s.referrerPolicy != "" {
		s.buf.WriteString(" referrerpolicy=\"" + s.referrerPolicy + "\"")
	}
	if len(s.style) != 0 {
		idx := 0
		s.buf.WriteString(" style=\"")
		for k, v := range s.style {
			s.buf.WriteString(k + ": " + v + ";")
			if idx != len(s.style)-1 {
				s.buf.WriteByte(' ')
			}
			idx++
		}
		s.buf.WriteString("\"")
	}
	s.buf.WriteByte('>')

	if s.text != "" {
		s.buf.WriteString(s.text)
	}
	s.buf.WriteString("</script>")
}
