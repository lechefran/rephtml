package rephtml

import "bytes"

// Canvas represents the HTML <canvas> element, a drawing surface for scripts.
type Canvas struct {
	bodyElement
	contentNode[*Canvas]
	width  string
	height string
}

// NewCanvas returns a new, empty <canvas> element.
func NewCanvas() *Canvas {
	v := &Canvas{}
	v.init(v)
	return v
}

// Width sets the width attribute and returns the element.
func (c *Canvas) Width(width string) *Canvas {
	c.width = width
	return c
}

// Height sets the height attribute and returns the element.
func (c *Canvas) Height(height string) *Canvas {
	c.height = height
	return c
}

// renderTo writes the Canvas component's HTML to buf.
func (c *Canvas) renderTo(buf *bytes.Buffer) {
	tg := startTag(buf, "canvas")
	tg.attr("width", c.width)
	tg.attr("height", c.height)
	c.attrs(tg)
	tg.children(c.contents)
}

// Noscript represents the HTML <noscript> element, content shown when scripting is unsupported or disabled.
type Noscript struct {
	headElement
	bodyElement
	contentNode[*Noscript]
}

// NewNoscript returns a new, empty <noscript> element.
func NewNoscript() *Noscript {
	v := &Noscript{}
	v.init(v)
	return v
}

// renderTo writes the Noscript component's HTML to buf.
func (n *Noscript) renderTo(buf *bytes.Buffer) {
	tg := startTag(buf, "noscript")
	n.attrs(tg)
	tg.children(n.contents)
}

// Script represents the HTML <script> element, embedded or referenced executable code.
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

// NewScript returns a new, empty <script> element.
func NewScript() *Script {
	v := &Script{}
	v.init(v)
	return v
}

// Src sets the src attribute and returns the element.
//
// An unsafe URL is replaced with [BlockedURL] when rendered; see [IsSafeURL].
func (s *Script) Src(src string) *Script {
	s.src = src
	return s
}

// Type sets the type attribute and returns the element.
func (s *Script) Type(scriptType string) *Script {
	s.scriptType = scriptType
	return s
}

// Async sets the async attribute and returns the element.
func (s *Script) Async(async bool) *Script {
	s.async = async
	return s
}

// Defer sets the defer attribute and returns the element.
func (s *Script) Defer(deferScript bool) *Script {
	s.deferScript = deferScript
	return s
}

// CrossOrigin sets the crossorigin attribute and returns the element.
func (s *Script) CrossOrigin(crossorigin string) *Script {
	s.crossOrigin = crossorigin
	return s
}

// Integrity sets the integrity attribute and returns the element.
func (s *Script) Integrity(integrity string) *Script {
	s.integrity = integrity
	return s
}

// Nomodule sets the nomodule attribute and returns the element.
func (s *Script) Nomodule(nomodule bool) *Script {
	s.noModule = nomodule
	return s
}

// ReferrerPolicy sets the referrerpolicy attribute and returns the element.
func (s *Script) ReferrerPolicy(referrerpolicy string) *Script {
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
	s.attrs(tg)
	// Script content is JavaScript, not HTML, so it is not entity-escaped. Only
	// the sequence that would end the element is defused.
	tg.raw(escapeScript(s.text))
}
