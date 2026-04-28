package rephtml

import "bytes"

// Embed represents the Embed component or supporting type.
type Embed struct {
	buf       bytes.Buffer
	style     StyleMap
	src       string
	embedType string
	width     string
	height    string
}

// NewEmbed creates a new Embed component.
func NewEmbed() *Embed {
	return &Embed{
		style: make(StyleMap),
	}
}

// AddStyle adds one inline CSS declaration to the Embed component.
func (e *Embed) AddStyle(k, v string) *Embed {
	e.style[k] = v
	return e
}

// AddStyles adds multiple inline CSS declarations to the Embed component.
func (e *Embed) AddStyles(m StyleMap) *Embed {
	for k, v := range m {
		e.style[k] = v
	}
	return e
}

// Style replaces the inline CSS declarations on the Embed component.
func (e *Embed) Style(m StyleMap) *Embed {
	e.style = cloneStyleMap(m)
	return e
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

// Bytes returns a defensive copy of the rendered Embed bytes.
func (e *Embed) Bytes() []byte {
	return cloneBytes(e.buf.Bytes())
}

// Render returns freshly prepared Embed HTML bytes.
func (e *Embed) Render() []byte {
	return renderPrepared(e)
}

// HTML returns freshly prepared Embed HTML as a string.
func (e *Embed) HTML() string {
	return htmlPrepared(e)
}

// String returns freshly prepared Embed HTML as a string.
func (e *Embed) String() string {
	return e.HTML()
}

// IsBodyElement implements BodyElement interface
func (e *Embed) IsBodyElement() {}

// Prepare renders the Embed component into its internal buffer.
func (e *Embed) Prepare() {
	e.buf.Reset()
	e.buf.WriteString("<embed")
	if e.src != "" {
		writeAttr(&e.buf, "src", e.src)
	}
	if e.embedType != "" {
		writeAttr(&e.buf, "type", e.embedType)
	}
	if e.width != "" {
		writeAttr(&e.buf, "width", e.width)
	}
	if e.height != "" {
		writeAttr(&e.buf, "height", e.height)
	}
	if len(e.style) != 0 {
		parseStyle(&e.buf, e.style)
	}
	e.buf.WriteString(">")
}

// Iframe represents the Iframe component or supporting type.
type Iframe struct {
	buf             bytes.Buffer
	style           StyleMap
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
	return &Iframe{
		style: make(StyleMap),
	}
}

// AddStyle adds one inline CSS declaration to the Iframe component.
func (i *Iframe) AddStyle(k, v string) *Iframe {
	i.style[k] = v
	return i
}

// AddStyles adds multiple inline CSS declarations to the Iframe component.
func (i *Iframe) AddStyles(m StyleMap) *Iframe {
	for k, v := range m {
		i.style[k] = v
	}
	return i
}

// Style replaces the inline CSS declarations on the Iframe component.
func (i *Iframe) Style(m StyleMap) *Iframe {
	i.style = cloneStyleMap(m)
	return i
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

// Bytes returns a defensive copy of the rendered Iframe bytes.
func (i *Iframe) Bytes() []byte {
	return cloneBytes(i.buf.Bytes())
}

// Render returns freshly prepared Iframe HTML bytes.
func (i *Iframe) Render() []byte {
	return renderPrepared(i)
}

// HTML returns freshly prepared Iframe HTML as a string.
func (i *Iframe) HTML() string {
	return htmlPrepared(i)
}

// String returns freshly prepared Iframe HTML as a string.
func (i *Iframe) String() string {
	return i.HTML()
}

// IsBodyElement implements BodyElement interface
func (i *Iframe) IsBodyElement() {}

// Prepare renders the Iframe component into its internal buffer.
func (i *Iframe) Prepare() {
	i.buf.Reset()
	i.buf.WriteString("<iframe")
	if i.src != "" {
		writeAttr(&i.buf, "src", i.src)
	}
	if i.width != "" {
		writeAttr(&i.buf, "width", i.width)
	}
	if i.height != "" {
		writeAttr(&i.buf, "height", i.height)
	}
	if i.name != "" {
		writeAttr(&i.buf, "name", i.name)
	}
	if i.sandbox != "" {
		writeAttr(&i.buf, "sandbox", i.sandbox)
	}
	if i.allow != "" {
		writeAttr(&i.buf, "allow", i.allow)
	}
	if i.allowfullscreen {
		i.buf.WriteString(" allowfullscreen")
	}
	if i.loading != "" {
		writeAttr(&i.buf, "loading", i.loading)
	}
	if i.referrerpolicy != "" {
		writeAttr(&i.buf, "referrerpolicy", i.referrerpolicy)
	}
	if i.srcdoc != "" {
		writeAttr(&i.buf, "srcdoc", i.srcdoc)
	}
	if len(i.style) != 0 {
		parseStyle(&i.buf, i.style)
	}
	i.buf.WriteString("></iframe>")
}

// Object represents the Object component or supporting type.
type Object struct {
	buf      bytes.Buffer
	style    StyleMap
	contents []Element
	data     string
	objType  string
	width    string
	height   string
	name     string
	usemap   string
	form     string
}

// NewObject creates a new Object component.
func NewObject() *Object {
	return &Object{
		style: make(StyleMap),
	}
}

// AddStyle adds one inline CSS declaration to the Object component.
func (o *Object) AddStyle(k, v string) *Object {
	o.style[k] = v
	return o
}

// AddStyles adds multiple inline CSS declarations to the Object component.
func (o *Object) AddStyles(m StyleMap) *Object {
	for k, v := range m {
		o.style[k] = v
	}
	return o
}

// Style replaces the inline CSS declarations on the Object component.
func (o *Object) Style(m StyleMap) *Object {
	o.style = cloneStyleMap(m)
	return o
}

// Add appends child content to the Object component.
func (o *Object) Add(e Element) *Object {
	o.contents = appendElement(o.contents, e)
	return o
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

// Bytes returns a defensive copy of the rendered Object bytes.
func (o *Object) Bytes() []byte {
	return cloneBytes(o.buf.Bytes())
}

// Render returns freshly prepared Object HTML bytes.
func (o *Object) Render() []byte {
	return renderPrepared(o)
}

// HTML returns freshly prepared Object HTML as a string.
func (o *Object) HTML() string {
	return htmlPrepared(o)
}

// String returns freshly prepared Object HTML as a string.
func (o *Object) String() string {
	return o.HTML()
}

// IsBodyElement implements BodyElement interface
func (o *Object) IsBodyElement() {}

// Prepare renders the Object component into its internal buffer.
func (o *Object) Prepare() {
	o.buf.Reset()
	o.buf.WriteString("<object")
	if o.data != "" {
		writeAttr(&o.buf, "data", o.data)
	}
	if o.objType != "" {
		writeAttr(&o.buf, "type", o.objType)
	}
	if o.width != "" {
		writeAttr(&o.buf, "width", o.width)
	}
	if o.height != "" {
		writeAttr(&o.buf, "height", o.height)
	}
	if o.name != "" {
		writeAttr(&o.buf, "name", o.name)
	}
	if o.usemap != "" {
		writeAttr(&o.buf, "usemap", o.usemap)
	}
	if o.form != "" {
		writeAttr(&o.buf, "form", o.form)
	}
	if len(o.style) != 0 {
		parseStyle(&o.buf, o.style)
	}
	o.buf.WriteByte('>')

	writeElements(&o.buf, o.contents)
	o.buf.WriteString("</object>")
}

// Picture represents the Picture component or supporting type.
type Picture struct {
	buf      bytes.Buffer
	style    StyleMap
	contents []Element
}

// NewPicture creates a new Picture component.
func NewPicture() *Picture {
	return &Picture{
		style: make(StyleMap),
	}
}

// AddStyle adds one inline CSS declaration to the Picture component.
func (p *Picture) AddStyle(k, v string) *Picture {
	p.style[k] = v
	return p
}

// AddStyles adds multiple inline CSS declarations to the Picture component.
func (p *Picture) AddStyles(m StyleMap) *Picture {
	for k, v := range m {
		p.style[k] = v
	}
	return p
}

// Style replaces the inline CSS declarations on the Picture component.
func (p *Picture) Style(m StyleMap) *Picture {
	p.style = cloneStyleMap(m)
	return p
}

// Add appends child content to the Picture component.
func (p *Picture) Add(e Element) *Picture {
	p.contents = appendElement(p.contents, e)
	return p
}

// Bytes returns a defensive copy of the rendered Picture bytes.
func (p *Picture) Bytes() []byte {
	return cloneBytes(p.buf.Bytes())
}

// Render returns freshly prepared Picture HTML bytes.
func (p *Picture) Render() []byte {
	return renderPrepared(p)
}

// HTML returns freshly prepared Picture HTML as a string.
func (p *Picture) HTML() string {
	return htmlPrepared(p)
}

// String returns freshly prepared Picture HTML as a string.
func (p *Picture) String() string {
	return p.HTML()
}

// IsBodyElement implements BodyElement interface
func (p *Picture) IsBodyElement() {}

// Prepare renders the Picture component into its internal buffer.
func (p *Picture) Prepare() {
	p.buf.Reset()
	p.buf.WriteString("<picture")
	if len(p.style) != 0 {
		parseStyle(&p.buf, p.style)
	}
	p.buf.WriteByte('>')

	writeElements(&p.buf, p.contents)
	p.buf.WriteString("</picture>")
}

// Portal represents the Portal component or supporting type.
type Portal struct {
	buf            bytes.Buffer
	style          StyleMap
	src            string
	referrerpolicy string
}

// NewPortal creates a new Portal component.
func NewPortal() *Portal {
	return &Portal{
		style: make(StyleMap),
	}
}

// AddStyle adds one inline CSS declaration to the Portal component.
func (p *Portal) AddStyle(k, v string) *Portal {
	p.style[k] = v
	return p
}

// AddStyles adds multiple inline CSS declarations to the Portal component.
func (p *Portal) AddStyles(m StyleMap) *Portal {
	for k, v := range m {
		p.style[k] = v
	}
	return p
}

// Style replaces the inline CSS declarations on the Portal component.
func (p *Portal) Style(m StyleMap) *Portal {
	p.style = cloneStyleMap(m)
	return p
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

// Bytes returns a defensive copy of the rendered Portal bytes.
func (p *Portal) Bytes() []byte {
	return cloneBytes(p.buf.Bytes())
}

// Render returns freshly prepared Portal HTML bytes.
func (p *Portal) Render() []byte {
	return renderPrepared(p)
}

// HTML returns freshly prepared Portal HTML as a string.
func (p *Portal) HTML() string {
	return htmlPrepared(p)
}

// String returns freshly prepared Portal HTML as a string.
func (p *Portal) String() string {
	return p.HTML()
}

// IsBodyElement implements BodyElement interface
func (p *Portal) IsBodyElement() {}

// Prepare renders the Portal component into its internal buffer.
func (p *Portal) Prepare() {
	p.buf.Reset()
	p.buf.WriteString("<portal")
	if p.src != "" {
		writeAttr(&p.buf, "src", p.src)
	}
	if p.referrerpolicy != "" {
		writeAttr(&p.buf, "referrerpolicy", p.referrerpolicy)
	}
	if len(p.style) != 0 {
		parseStyle(&p.buf, p.style)
	}
	p.buf.WriteString("></portal>")
}

// Source represents the Source component or supporting type.
type Source struct {
	buf     bytes.Buffer
	style   StyleMap
	src     string
	srcset  string
	media   string
	sizes   string
	srcType string
}

// NewSource creates a new Source component.
func NewSource() *Source {
	return &Source{
		style: make(StyleMap),
	}
}

// AddStyle adds one inline CSS declaration to the Source component.
func (s *Source) AddStyle(k, v string) *Source {
	s.style[k] = v
	return s
}

// AddStyles adds multiple inline CSS declarations to the Source component.
func (s *Source) AddStyles(m StyleMap) *Source {
	for k, v := range m {
		s.style[k] = v
	}
	return s
}

// Style replaces the inline CSS declarations on the Source component.
func (s *Source) Style(m StyleMap) *Source {
	s.style = cloneStyleMap(m)
	return s
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

// Bytes returns a defensive copy of the rendered Source bytes.
func (s *Source) Bytes() []byte {
	return cloneBytes(s.buf.Bytes())
}

// Render returns freshly prepared Source HTML bytes.
func (s *Source) Render() []byte {
	return renderPrepared(s)
}

// HTML returns freshly prepared Source HTML as a string.
func (s *Source) HTML() string {
	return htmlPrepared(s)
}

// String returns freshly prepared Source HTML as a string.
func (s *Source) String() string {
	return s.HTML()
}

// IsBodyElement implements BodyElement interface
func (s *Source) IsBodyElement() {}

// Prepare renders the Source component into its internal buffer.
func (s *Source) Prepare() {
	s.buf.Reset()
	s.buf.WriteString("<source")
	if s.src != "" {
		writeAttr(&s.buf, "src", s.src)
	}
	if s.srcset != "" {
		writeAttr(&s.buf, "srcset", s.srcset)
	}
	if s.media != "" {
		writeAttr(&s.buf, "media", s.media)
	}
	if s.sizes != "" {
		writeAttr(&s.buf, "sizes", s.sizes)
	}
	if s.srcType != "" {
		writeAttr(&s.buf, "type", s.srcType)
	}
	if len(s.style) != 0 {
		parseStyle(&s.buf, s.style)
	}
	s.buf.WriteString(">")
}
