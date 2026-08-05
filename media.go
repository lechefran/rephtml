package rephtml

import "bytes"

// Area represents the Area component or supporting type.
type Area struct {
	bodyElement
	node[*Area]
	alt    string
	coords string
	href   string
	shape  string
	target string
}

// NewArea creates a new Area component.
func NewArea() *Area {
	v := &Area{}
	v.init(v)
	return v
}

// Alt sets the alt value on the Area component.
func (a *Area) Alt(alt string) *Area {
	a.alt = alt
	return a
}

// Coords sets the coords value on the Area component.
func (a *Area) Coords(coords string) *Area {
	a.coords = coords
	return a
}

// Href sets the href value on the Area component.
func (a *Area) Href(href string) *Area {
	a.href = href
	return a
}

// Shape sets the shape value on the Area component.
func (a *Area) Shape(shape string) *Area {
	a.shape = shape
	return a
}

// Target sets the target value on the Area component.
func (a *Area) Target(target string) *Area {
	a.target = target
	return a
}

// renderTo writes the Area component's HTML to buf.
func (a *Area) renderTo(buf *bytes.Buffer) {
	tg := startTag(buf, "area")
	tg.attr("alt", a.alt)
	tg.attr("coords", a.coords)
	tg.urlAttr("href", a.href)
	tg.attr("shape", a.shape)
	tg.attr("target", a.target)
	tg.styleAttr(a.style)
	tg.void()
}

// Img represents the Img component or supporting type.
type Img struct {
	bodyElement
	node[*Img]
	src    string
	alt    string
	width  string
	height string
	title  string
}

// NewImg creates a new Img component.
func NewImg() *Img {
	v := &Img{}
	v.init(v)
	return v
}

// Src sets the src value on the Img component.
func (i *Img) Src(src string) *Img {
	i.src = src
	return i
}

// Alt sets the alt value on the Img component.
func (i *Img) Alt(alt string) *Img {
	i.alt = alt
	return i
}

// Width sets the width value on the Img component.
func (i *Img) Width(width string) *Img {
	i.width = width
	return i
}

// Height sets the height value on the Img component.
func (i *Img) Height(height string) *Img {
	i.height = height
	return i
}

// Title sets the title value on the Img component.
func (i *Img) Title(title string) *Img {
	i.title = title
	return i
}

// renderTo writes the Img component's HTML to buf.
func (i *Img) renderTo(buf *bytes.Buffer) {
	tg := startTag(buf, "img")
	tg.urlAttr("src", i.src)
	tg.attr("alt", i.alt)
	tg.attr("width", i.width)
	tg.attr("height", i.height)
	tg.attr("title", i.title)
	tg.styleAttr(i.style)
	tg.void()
}

// Audio represents the Audio component or supporting type.
type Audio struct {
	bodyElement
	contentNode[*Audio]
	src      string
	controls bool
	autoplay bool
	loop     bool
	muted    bool
	preload  string
}

// NewAudio creates a new Audio component.
func NewAudio() *Audio {
	v := &Audio{}
	v.init(v)
	return v
}

// Src sets the src value on the Audio component.
func (a *Audio) Src(src string) *Audio {
	a.src = src
	return a
}

// Controls sets the controls value on the Audio component.
func (a *Audio) Controls(controls bool) *Audio {
	a.controls = controls
	return a
}

// Autoplay sets the autoplay value on the Audio component.
func (a *Audio) Autoplay(autoplay bool) *Audio {
	a.autoplay = autoplay
	return a
}

// Loop sets the loop value on the Audio component.
func (a *Audio) Loop(loop bool) *Audio {
	a.loop = loop
	return a
}

// Muted sets the muted value on the Audio component.
func (a *Audio) Muted(muted bool) *Audio {
	a.muted = muted
	return a
}

// Preload sets the preload value on the Audio component.
func (a *Audio) Preload(preload string) *Audio {
	a.preload = preload
	return a
}

// renderTo writes the Audio component's HTML to buf.
func (a *Audio) renderTo(buf *bytes.Buffer) {
	tg := startTag(buf, "audio")
	tg.urlAttr("src", a.src)
	tg.boolAttr("controls", a.controls)
	tg.boolAttr("autoplay", a.autoplay)
	tg.boolAttr("loop", a.loop)
	tg.boolAttr("muted", a.muted)
	tg.attr("preload", a.preload)
	tg.styleAttr(a.style)
	tg.children(a.contents)
}

// Track represents the Track component or supporting type.
type Track struct {
	bodyElement
	node[*Track]
	src        string
	kind       string
	srclang    string
	label      string
	defaultVal bool
}

// NewTrack creates a new Track component.
func NewTrack() *Track {
	v := &Track{}
	v.init(v)
	return v
}

// Src sets the src value on the Track component.
func (t *Track) Src(src string) *Track {
	t.src = src
	return t
}

// Kind sets the kind value on the Track component.
func (t *Track) Kind(kind string) *Track {
	t.kind = kind
	return t
}

// Srclang sets the srclang value on the Track component.
func (t *Track) Srclang(srclang string) *Track {
	t.srclang = srclang
	return t
}

// Label sets the label value on the Track component.
func (t *Track) Label(label string) *Track {
	t.label = label
	return t
}

// Default sets the default value on the Track component.
func (t *Track) Default(def bool) *Track {
	t.defaultVal = def
	return t
}

// renderTo writes the Track component's HTML to buf.
func (t *Track) renderTo(buf *bytes.Buffer) {
	tg := startTag(buf, "track")
	tg.urlAttr("src", t.src)
	tg.attr("kind", t.kind)
	tg.attr("srclang", t.srclang)
	tg.attr("label", t.label)
	tg.boolAttr("default", t.defaultVal)
	tg.styleAttr(t.style)
	tg.void()
}

// Map represents the Map component or supporting type.
type Map struct {
	bodyElement
	contentNode[*Map]
	name string
}

// NewMap creates a new Map component.
func NewMap() *Map {
	v := &Map{}
	v.init(v)
	return v
}

// Name sets the name value on the Map component.
func (m *Map) Name(name string) *Map {
	m.name = name
	return m
}

// renderTo writes the Map component's HTML to buf.
func (m *Map) renderTo(buf *bytes.Buffer) {
	tg := startTag(buf, "map")
	tg.attr("name", m.name)
	tg.styleAttr(m.style)
	tg.children(m.contents)
}

// Video represents the Video component or supporting type.
type Video struct {
	bodyElement
	contentNode[*Video]
	src      string
	controls bool
	autoplay bool
	loop     bool
	muted    bool
	preload  string
	width    string
	height   string
	poster   string
}

// NewVideo creates a new Video component.
func NewVideo() *Video {
	v := &Video{}
	v.init(v)
	return v
}

// Src sets the src value on the Video component.
func (v *Video) Src(src string) *Video {
	v.src = src
	return v
}

// Controls sets the controls value on the Video component.
func (v *Video) Controls(controls bool) *Video {
	v.controls = controls
	return v
}

// Autoplay sets the autoplay value on the Video component.
func (v *Video) Autoplay(autoplay bool) *Video {
	v.autoplay = autoplay
	return v
}

// Loop sets the loop value on the Video component.
func (v *Video) Loop(loop bool) *Video {
	v.loop = loop
	return v
}

// Muted sets the muted value on the Video component.
func (v *Video) Muted(muted bool) *Video {
	v.muted = muted
	return v
}

// Preload sets the preload value on the Video component.
func (v *Video) Preload(preload string) *Video {
	v.preload = preload
	return v
}

// Width sets the width value on the Video component.
func (v *Video) Width(width string) *Video {
	v.width = width
	return v
}

// Height sets the height value on the Video component.
func (v *Video) Height(height string) *Video {
	v.height = height
	return v
}

// Poster sets the poster value on the Video component.
func (v *Video) Poster(poster string) *Video {
	v.poster = poster
	return v
}

// renderTo writes the Video component's HTML to buf.
func (v *Video) renderTo(buf *bytes.Buffer) {
	tg := startTag(buf, "video")
	tg.urlAttr("src", v.src)
	tg.boolAttr("controls", v.controls)
	tg.boolAttr("autoplay", v.autoplay)
	tg.boolAttr("loop", v.loop)
	tg.boolAttr("muted", v.muted)
	tg.attr("preload", v.preload)
	tg.attr("width", v.width)
	tg.attr("height", v.height)
	tg.urlAttr("poster", v.poster)
	tg.styleAttr(v.style)
	tg.children(v.contents)
}
