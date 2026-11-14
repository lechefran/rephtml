package rephtml

import "bytes"

// Embed represents the HTML <embed> element, an external resource, such as a plugin or PDF, embedded at this point.
type Embed struct {
	bodyElement
	node[*Embed]
	src       string
	embedType string
	width     string
	height    string
}

// NewEmbed creates a new Embed component.
func NewEmbed() *Embed {
	v := &Embed{}
	v.init(v)
	return v
}

// Src sets the src attribute.
func (e *Embed) Src(src string) *Embed {
	e.src = src
	return e
}

// Type sets the type attribute.
func (e *Embed) Type(embedType string) *Embed {
	e.embedType = embedType
	return e
}

// Width sets the width attribute.
func (e *Embed) Width(width string) *Embed {
	e.width = width
	return e
}

// Height sets the height attribute.
func (e *Embed) Height(height string) *Embed {
	e.height = height
	return e
}

// renderTo writes the Embed component's HTML to buf.
func (e *Embed) renderTo(buf *bytes.Buffer) {
	tg := startTag(buf, "embed")
	tg.urlAttr("src", e.src)
	tg.attr("type", e.embedType)
	tg.attr("width", e.width)
	tg.attr("height", e.height)
	tg.styleAttr(e.style)
	tg.void()
}

// Iframe represents the HTML <iframe> element, a nested browsing context that embeds another HTML page.
type Iframe struct {
	bodyElement
	node[*Iframe]
	src             string
	width           string
	height          string
	name            string
	sandbox         string
	allow           string
	allowfullscreen bool
	loading         string
	referrerpolicy  string
	srcdoc          string
}

// NewIframe creates a new Iframe component.
func NewIframe() *Iframe {
	v := &Iframe{}
	v.init(v)
	return v
}

// Src sets the src attribute.
func (i *Iframe) Src(src string) *Iframe {
	i.src = src
	return i
}

// Width sets the width attribute.
func (i *Iframe) Width(width string) *Iframe {
	i.width = width
	return i
}

// Height sets the height attribute.
func (i *Iframe) Height(height string) *Iframe {
	i.height = height
	return i
}

// Name sets the name attribute.
func (i *Iframe) Name(name string) *Iframe {
	i.name = name
	return i
}

// Sandbox sets the sandbox attribute.
func (i *Iframe) Sandbox(sandbox string) *Iframe {
	i.sandbox = sandbox
	return i
}

// Allow sets the allow attribute.
func (i *Iframe) Allow(allow string) *Iframe {
	i.allow = allow
	return i
}

// AllowFullscreen sets the allowfullscreen attribute.
func (i *Iframe) AllowFullscreen(allowfullscreen bool) *Iframe {
	i.allowfullscreen = allowfullscreen
	return i
}

// Allowfullscreen sets the allowfullscreen attribute.
//
// Deprecated: use AllowFullscreen.
func (i *Iframe) Allowfullscreen(allowfullscreen bool) *Iframe {
	return i.AllowFullscreen(allowfullscreen)
}

// Loading sets the loading attribute.
func (i *Iframe) Loading(loading string) *Iframe {
	i.loading = loading
	return i
}

// ReferrerPolicy sets the referrerpolicy attribute.
func (i *Iframe) ReferrerPolicy(referrerpolicy string) *Iframe {
	i.referrerpolicy = referrerpolicy
	return i
}

// Referrerpolicy sets the referrerpolicy attribute.
//
// Deprecated: use ReferrerPolicy.
func (i *Iframe) Referrerpolicy(referrerpolicy string) *Iframe {
	return i.ReferrerPolicy(referrerpolicy)
}

// Srcdoc sets the srcdoc attribute, the HTML document the frame displays.
//
// The value is escaped as an attribute, but its content is not otherwise
// checked: the browser parses it as a full document and runs any script in it.
// Do not pass untrusted HTML unless the frame also has a restrictive Sandbox.
func (i *Iframe) Srcdoc(srcdoc string) *Iframe {
	i.srcdoc = srcdoc
	return i
}

// renderTo writes the Iframe component's HTML to buf.
func (i *Iframe) renderTo(buf *bytes.Buffer) {
	tg := startTag(buf, "iframe")
	tg.urlAttr("src", i.src)
	tg.attr("width", i.width)
	tg.attr("height", i.height)
	tg.attr("name", i.name)
	tg.attr("sandbox", i.sandbox)
	tg.attr("allow", i.allow)
	tg.boolAttr("allowfullscreen", i.allowfullscreen)
	tg.attr("loading", i.loading)
	tg.attr("referrerpolicy", i.referrerpolicy)
	tg.attr("srcdoc", i.srcdoc)
	tg.styleAttr(i.style)
	tg.empty()
}

// Object represents the HTML <object> element, an external resource treated as an image, a nested browsing context or a plugin.
type Object struct {
	bodyElement
	contentNode[*Object]
	data    string
	objType string
	width   string
	height  string
	name    string
	usemap  string
	form    string
}

// NewObject creates a new Object component.
func NewObject() *Object {
	v := &Object{}
	v.init(v)
	return v
}

// Data sets the data attribute.
func (o *Object) Data(data string) *Object {
	o.data = data
	return o
}

// Type sets the type attribute.
func (o *Object) Type(objType string) *Object {
	o.objType = objType
	return o
}

// Width sets the width attribute.
func (o *Object) Width(width string) *Object {
	o.width = width
	return o
}

// Height sets the height attribute.
func (o *Object) Height(height string) *Object {
	o.height = height
	return o
}

// Name sets the name attribute.
func (o *Object) Name(name string) *Object {
	o.name = name
	return o
}

// Usemap sets the usemap attribute.
func (o *Object) Usemap(usemap string) *Object {
	o.usemap = usemap
	return o
}

// Form sets the form attribute.
func (o *Object) Form(form string) *Object {
	o.form = form
	return o
}

// renderTo writes the Object component's HTML to buf.
func (o *Object) renderTo(buf *bytes.Buffer) {
	tg := startTag(buf, "object")
	tg.urlAttr("data", o.data)
	tg.attr("type", o.objType)
	tg.attr("width", o.width)
	tg.attr("height", o.height)
	tg.attr("name", o.name)
	tg.urlAttr("usemap", o.usemap)
	tg.attr("form", o.form)
	tg.styleAttr(o.style)
	tg.children(o.contents)
}

// Picture represents the HTML <picture> element, a container that offers alternative sources for an image.
type Picture struct {
	bodyElement
	contentNode[*Picture]
}

// NewPicture creates a new Picture component.
func NewPicture() *Picture {
	v := &Picture{}
	v.init(v)
	return v
}

// renderTo writes the Picture component's HTML to buf.
func (p *Picture) renderTo(buf *bytes.Buffer) {
	tg := startTag(buf, "picture")
	tg.styleAttr(p.style)
	tg.children(p.contents)
}

// Portal represents the HTML <portal> element, a preview of another page embedded in this one. It is obsolete and was never widely supported.
type Portal struct {
	bodyElement
	node[*Portal]
	src            string
	referrerpolicy string
}

// NewPortal creates a new Portal component.
func NewPortal() *Portal {
	v := &Portal{}
	v.init(v)
	return v
}

// Src sets the src attribute.
func (p *Portal) Src(src string) *Portal {
	p.src = src
	return p
}

// ReferrerPolicy sets the referrerpolicy attribute.
func (p *Portal) ReferrerPolicy(referrerpolicy string) *Portal {
	p.referrerpolicy = referrerpolicy
	return p
}

// Referrerpolicy sets the referrerpolicy attribute.
//
// Deprecated: use ReferrerPolicy.
func (p *Portal) Referrerpolicy(referrerpolicy string) *Portal {
	return p.ReferrerPolicy(referrerpolicy)
}

// renderTo writes the Portal component's HTML to buf.
func (p *Portal) renderTo(buf *bytes.Buffer) {
	tg := startTag(buf, "portal")
	tg.urlAttr("src", p.src)
	tg.attr("referrerpolicy", p.referrerpolicy)
	tg.styleAttr(p.style)
	tg.empty()
}

// Source represents the HTML <source> element, a media resource for a picture, audio or video element.
type Source struct {
	bodyElement
	node[*Source]
	src     string
	srcset  string
	media   string
	sizes   string
	srcType string
}

// NewSource creates a new Source component.
func NewSource() *Source {
	v := &Source{}
	v.init(v)
	return v
}

// Src sets the src attribute.
func (s *Source) Src(src string) *Source {
	s.src = src
	return s
}

// Srcset sets the srcset attribute.
func (s *Source) Srcset(srcset string) *Source {
	s.srcset = srcset
	return s
}

// Media sets the media attribute.
func (s *Source) Media(media string) *Source {
	s.media = media
	return s
}

// Sizes sets the sizes attribute.
func (s *Source) Sizes(sizes string) *Source {
	s.sizes = sizes
	return s
}

// Type sets the type attribute.
func (s *Source) Type(srcType string) *Source {
	s.srcType = srcType
	return s
}

// renderTo writes the Source component's HTML to buf.
func (s *Source) renderTo(buf *bytes.Buffer) {
	tg := startTag(buf, "source")
	tg.urlAttr("src", s.src)
	tg.srcsetAttr("srcset", s.srcset)
	tg.attr("media", s.media)
	tg.attr("sizes", s.sizes)
	tg.attr("type", s.srcType)
	tg.styleAttr(s.style)
	tg.void()
}
