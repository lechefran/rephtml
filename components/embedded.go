package rephtml

import "bytes"

type Embed struct {
	buf       bytes.Buffer
	style     map[string]string
	src       string
	embedType string
	width     string
	height    string
}

func NewEmbed() *Embed {
	return &Embed{
		style: make(map[string]string),
	}
}

func (e *Embed) AddStyle(k, v string) *Embed {
	e.style[k] = v
	return e
}

func (e *Embed) AddStyles(m map[string]string) *Embed {
	for k, v := range m {
		e.style[k] = v
	}
	return e
}

func (e *Embed) Style(m map[string]string) *Embed {
	e.style = m
	return e
}

func (e *Embed) Src(src string) *Embed {
	e.src = src
	return e
}

func (e *Embed) Type(embedType string) *Embed {
	e.embedType = embedType
	return e
}

func (e *Embed) Width(width string) *Embed {
	e.width = width
	return e
}

func (e *Embed) Height(height string) *Embed {
	e.height = height
	return e
}

func (e *Embed) Bytes() []byte {
	return cloneBytes(e.buf.Bytes())
}

// IsBodyElement implements BodyElement interface
func (e *Embed) IsBodyElement() {}

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

type Iframe struct {
	buf             bytes.Buffer
	style           map[string]string
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

func NewIframe() *Iframe {
	return &Iframe{
		style: make(map[string]string),
	}
}

func (i *Iframe) AddStyle(k, v string) *Iframe {
	i.style[k] = v
	return i
}

func (i *Iframe) AddStyles(m map[string]string) *Iframe {
	for k, v := range m {
		i.style[k] = v
	}
	return i
}

func (i *Iframe) Style(m map[string]string) *Iframe {
	i.style = m
	return i
}

func (i *Iframe) Src(src string) *Iframe {
	i.src = src
	return i
}

func (i *Iframe) Width(width string) *Iframe {
	i.width = width
	return i
}

func (i *Iframe) Height(height string) *Iframe {
	i.height = height
	return i
}

func (i *Iframe) Name(name string) *Iframe {
	i.name = name
	return i
}

func (i *Iframe) Sandbox(sandbox string) *Iframe {
	i.sandbox = sandbox
	return i
}

func (i *Iframe) Allow(allow string) *Iframe {
	i.allow = allow
	return i
}

func (i *Iframe) Allowfullscreen(allowfullscreen bool) *Iframe {
	i.allowfullscreen = allowfullscreen
	return i
}

func (i *Iframe) Loading(loading string) *Iframe {
	i.loading = loading
	return i
}

func (i *Iframe) Referrerpolicy(referrerpolicy string) *Iframe {
	i.referrerpolicy = referrerpolicy
	return i
}

func (i *Iframe) Srcdoc(srcdoc string) *Iframe {
	i.srcdoc = srcdoc
	return i
}

func (i *Iframe) Bytes() []byte {
	return cloneBytes(i.buf.Bytes())
}

// IsBodyElement implements BodyElement interface
func (i *Iframe) IsBodyElement() {}

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

type Object struct {
	buf      bytes.Buffer
	style    map[string]string
	contents []Element
	data     string
	objType  string
	width    string
	height   string
	name     string
	usemap   string
	form     string
}

func NewObject() *Object {
	return &Object{
		style: make(map[string]string),
	}
}

func (o *Object) AddStyle(k, v string) *Object {
	o.style[k] = v
	return o
}

func (o *Object) AddStyles(m map[string]string) *Object {
	for k, v := range m {
		o.style[k] = v
	}
	return o
}

func (o *Object) Style(m map[string]string) *Object {
	o.style = m
	return o
}

func (o *Object) Add(e Element) *Object {
	o.contents = appendElement(o.contents, e)
	return o
}

func (o *Object) Data(data string) *Object {
	o.data = data
	return o
}

func (o *Object) Type(objType string) *Object {
	o.objType = objType
	return o
}

func (o *Object) Width(width string) *Object {
	o.width = width
	return o
}

func (o *Object) Height(height string) *Object {
	o.height = height
	return o
}

func (o *Object) Name(name string) *Object {
	o.name = name
	return o
}

func (o *Object) Usemap(usemap string) *Object {
	o.usemap = usemap
	return o
}

func (o *Object) Form(form string) *Object {
	o.form = form
	return o
}

func (o *Object) Bytes() []byte {
	return cloneBytes(o.buf.Bytes())
}

// IsBodyElement implements BodyElement interface
func (o *Object) IsBodyElement() {}

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

type Picture struct {
	buf      bytes.Buffer
	style    map[string]string
	contents []Element
}

func NewPicture() *Picture {
	return &Picture{
		style: make(map[string]string),
	}
}

func (p *Picture) AddStyle(k, v string) *Picture {
	p.style[k] = v
	return p
}

func (p *Picture) AddStyles(m map[string]string) *Picture {
	for k, v := range m {
		p.style[k] = v
	}
	return p
}

func (p *Picture) Style(m map[string]string) *Picture {
	p.style = m
	return p
}

func (p *Picture) Add(e Element) *Picture {
	p.contents = appendElement(p.contents, e)
	return p
}

func (p *Picture) Bytes() []byte {
	return cloneBytes(p.buf.Bytes())
}

// IsBodyElement implements BodyElement interface
func (p *Picture) IsBodyElement() {}

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

type Portal struct {
	buf            bytes.Buffer
	style          map[string]string
	src            string
	referrerpolicy string
}

func NewPortal() *Portal {
	return &Portal{
		style: make(map[string]string),
	}
}

func (p *Portal) AddStyle(k, v string) *Portal {
	p.style[k] = v
	return p
}

func (p *Portal) AddStyles(m map[string]string) *Portal {
	for k, v := range m {
		p.style[k] = v
	}
	return p
}

func (p *Portal) Style(m map[string]string) *Portal {
	p.style = m
	return p
}

func (p *Portal) Src(src string) *Portal {
	p.src = src
	return p
}

func (p *Portal) Referrerpolicy(referrerpolicy string) *Portal {
	p.referrerpolicy = referrerpolicy
	return p
}

func (p *Portal) Bytes() []byte {
	return cloneBytes(p.buf.Bytes())
}

// IsBodyElement implements BodyElement interface
func (p *Portal) IsBodyElement() {}

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

type Source struct {
	buf     bytes.Buffer
	style   map[string]string
	src     string
	srcset  string
	media   string
	sizes   string
	srcType string
}

func NewSource() *Source {
	return &Source{
		style: make(map[string]string),
	}
}

func (s *Source) AddStyle(k, v string) *Source {
	s.style[k] = v
	return s
}

func (s *Source) AddStyles(m map[string]string) *Source {
	for k, v := range m {
		s.style[k] = v
	}
	return s
}

func (s *Source) Style(m map[string]string) *Source {
	s.style = m
	return s
}

func (s *Source) Src(src string) *Source {
	s.src = src
	return s
}

func (s *Source) Srcset(srcset string) *Source {
	s.srcset = srcset
	return s
}

func (s *Source) Media(media string) *Source {
	s.media = media
	return s
}

func (s *Source) Sizes(sizes string) *Source {
	s.sizes = sizes
	return s
}

func (s *Source) Type(srcType string) *Source {
	s.srcType = srcType
	return s
}

func (s *Source) Bytes() []byte {
	return cloneBytes(s.buf.Bytes())
}

// IsBodyElement implements BodyElement interface
func (s *Source) IsBodyElement() {}

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
