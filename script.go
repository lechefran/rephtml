package rephtml

import "bytes"

// Canvas represents the Canvas component or supporting type.
type Canvas struct {
	bodyElement
	contentNode[*Canvas]
	width  string
	height string
}

// NewCanvas creates a new Canvas component.
func NewCanvas() *Canvas {
	v := &Canvas{}
	v.init(v)
	return v
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

// renderTo writes the Canvas component's HTML to buf.
func (c *Canvas) renderTo(buf *bytes.Buffer) {
	tg := startTag(buf, "canvas")
	tg.attr("width", c.width)
	tg.attr("height", c.height)
	tg.styleAttr(c.style)
	tg.children(c.contents)
}

// Noscript represents the Noscript component or supporting type.
type Noscript struct {
	headElement
	bodyElement
	contentNode[*Noscript]
}

// NewNoscript creates a new Noscript component.
func NewNoscript() *Noscript {
	v := &Noscript{}
	v.init(v)
	return v
}

// renderTo writes the Noscript component's HTML to buf.
func (n *Noscript) renderTo(buf *bytes.Buffer) {
	tg := startTag(buf, "noscript")
	tg.styleAttr(n.style)
	tg.children(n.contents)
}

// Script represents the Script component or supporting type.
type Script struct {
	headElement
	bodyElement
	textNode[*Script]
	src            string
	scriptType     string
	async          bool
	deferScript    bool
	crossOrigin    string
	integrity      string
	noModule       bool
	referrerPolicy string
}

// NewScript creates a new Script component.
func NewScript() *Script {
	v := &Script{}
	v.init(v)
	return v
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

// renderTo writes the Script component's HTML to buf.
func (s *Script) renderTo(buf *bytes.Buffer) {
	tg := startTag(buf, "script")
	tg.urlAttr("src", s.src)
	tg.attr("type", s.scriptType)
	tg.boolAttr("async", s.async)
	tg.boolAttr("defer", s.deferScript)
	tg.attr("crossorigin", s.crossOrigin)
	tg.attr("integrity", s.integrity)
	tg.boolAttr("nomodule", s.noModule)
	tg.attr("referrerpolicy", s.referrerPolicy)
	tg.styleAttr(s.style)
	// Script content is JavaScript, not HTML, so it is written unescaped.
	tg.raw(s.text)
}
