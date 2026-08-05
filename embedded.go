package rephtml

// Embed represents the Embed component or supporting type.
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

// Src sets the src value on the Embed component.
func (e *Embed) Src(src string) *Embed {
	e.src = src
	return e
}

// Type sets the type value on the Embed component.
func (e *Embed) Type(embedType string) *Embed {
	e.embedType = embedType
	return e
}

// Width sets the width value on the Embed component.
func (e *Embed) Width(width string) *Embed {
	e.width = width
	return e
}

// Height sets the height value on the Embed component.
func (e *Embed) Height(height string) *Embed {
	e.height = height
	return e
}

// prepare renders the Embed component into its internal buffer.
func (e *Embed) prepare() {
	tg := openTag(&e.buf, "embed")
	tg.attr("src", e.src)
	tg.attr("type", e.embedType)
	tg.attr("width", e.width)
	tg.attr("height", e.height)
	tg.styleAttr(e.style)
	tg.void()
}

// Iframe represents the Iframe component or supporting type.
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

// Src sets the src value on the Iframe component.
func (i *Iframe) Src(src string) *Iframe {
	i.src = src
	return i
}

// Width sets the width value on the Iframe component.
func (i *Iframe) Width(width string) *Iframe {
	i.width = width
	return i
}

// Height sets the height value on the Iframe component.
func (i *Iframe) Height(height string) *Iframe {
	i.height = height
	return i
}

// Name sets the name value on the Iframe component.
func (i *Iframe) Name(name string) *Iframe {
	i.name = name
	return i
}

// Sandbox sets the sandbox value on the Iframe component.
func (i *Iframe) Sandbox(sandbox string) *Iframe {
	i.sandbox = sandbox
	return i
}

// Allow sets the allow value on the Iframe component.
func (i *Iframe) Allow(allow string) *Iframe {
	i.allow = allow
	return i
}

// Allowfullscreen sets the allowfullscreen value on the Iframe component.
func (i *Iframe) Allowfullscreen(allowfullscreen bool) *Iframe {
	i.allowfullscreen = allowfullscreen
	return i
}

// Loading sets the loading value on the Iframe component.
func (i *Iframe) Loading(loading string) *Iframe {
	i.loading = loading
	return i
}

// Referrerpolicy sets the referrerpolicy value on the Iframe component.
func (i *Iframe) Referrerpolicy(referrerpolicy string) *Iframe {
	i.referrerpolicy = referrerpolicy
	return i
}

// Srcdoc sets the srcdoc value on the Iframe component.
func (i *Iframe) Srcdoc(srcdoc string) *Iframe {
	i.srcdoc = srcdoc
	return i
}

// prepare renders the Iframe component into its internal buffer.
func (i *Iframe) prepare() {
	tg := openTag(&i.buf, "iframe")
	tg.attr("src", i.src)
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

// Object represents the Object component or supporting type.
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

// Data sets the data value on the Object component.
func (o *Object) Data(data string) *Object {
	o.data = data
	return o
}

// Type sets the type value on the Object component.
func (o *Object) Type(objType string) *Object {
	o.objType = objType
	return o
}

// Width sets the width value on the Object component.
func (o *Object) Width(width string) *Object {
	o.width = width
	return o
}

// Height sets the height value on the Object component.
func (o *Object) Height(height string) *Object {
	o.height = height
	return o
}

// Name sets the name value on the Object component.
func (o *Object) Name(name string) *Object {
	o.name = name
	return o
}

// Usemap sets the usemap value on the Object component.
func (o *Object) Usemap(usemap string) *Object {
	o.usemap = usemap
	return o
}

// Form sets the form value on the Object component.
func (o *Object) Form(form string) *Object {
	o.form = form
	return o
}

// prepare renders the Object component into its internal buffer.
func (o *Object) prepare() {
	tg := openTag(&o.buf, "object")
	tg.attr("data", o.data)
	tg.attr("type", o.objType)
	tg.attr("width", o.width)
	tg.attr("height", o.height)
	tg.attr("name", o.name)
	tg.attr("usemap", o.usemap)
	tg.attr("form", o.form)
	tg.styleAttr(o.style)
	tg.children(o.contents)
}

// Picture represents the Picture component or supporting type.
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

// prepare renders the Picture component into its internal buffer.
func (p *Picture) prepare() {
	tg := openTag(&p.buf, "picture")
	tg.styleAttr(p.style)
	tg.children(p.contents)
}

// Portal represents the Portal component or supporting type.
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

// Src sets the src value on the Portal component.
func (p *Portal) Src(src string) *Portal {
	p.src = src
	return p
}

// Referrerpolicy sets the referrerpolicy value on the Portal component.
func (p *Portal) Referrerpolicy(referrerpolicy string) *Portal {
	p.referrerpolicy = referrerpolicy
	return p
}

// prepare renders the Portal component into its internal buffer.
func (p *Portal) prepare() {
	tg := openTag(&p.buf, "portal")
	tg.attr("src", p.src)
	tg.attr("referrerpolicy", p.referrerpolicy)
	tg.styleAttr(p.style)
	tg.empty()
}

// Source represents the Source component or supporting type.
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

// Src sets the src value on the Source component.
func (s *Source) Src(src string) *Source {
	s.src = src
	return s
}

// Srcset sets the srcset value on the Source component.
func (s *Source) Srcset(srcset string) *Source {
	s.srcset = srcset
	return s
}

// Media sets the media value on the Source component.
func (s *Source) Media(media string) *Source {
	s.media = media
	return s
}

// Sizes sets the sizes value on the Source component.
func (s *Source) Sizes(sizes string) *Source {
	s.sizes = sizes
	return s
}

// Type sets the type value on the Source component.
func (s *Source) Type(srcType string) *Source {
	s.srcType = srcType
	return s
}

// prepare renders the Source component into its internal buffer.
func (s *Source) prepare() {
	tg := openTag(&s.buf, "source")
	tg.attr("src", s.src)
	tg.attr("srcset", s.srcset)
	tg.attr("media", s.media)
	tg.attr("sizes", s.sizes)
	tg.attr("type", s.srcType)
	tg.styleAttr(s.style)
	tg.void()
}
