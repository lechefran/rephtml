package rephtml

import "bytes"

type Area struct {
	buf    bytes.Buffer
	style  map[string]string
	alt    string
	coords string
	href   string
	shape  string
	target string
}

func NewArea() *Area {
	return &Area{
		style: make(map[string]string),
	}
}

func (a *Area) AddStyle(k, v string) *Area {
	a.style[k] = v
	return a
}

func (a *Area) AddStyles(m map[string]string) *Area {
	for k, v := range m {
		a.style[k] = v
	}
	return a
}

func (a *Area) Style(m map[string]string) *Area {
	a.style = m
	return a
}

func (a *Area) Alt(alt string) *Area {
	a.alt = alt
	return a
}

func (a *Area) Coords(coords string) *Area {
	a.coords = coords
	return a
}

func (a *Area) Href(href string) *Area {
	a.href = href
	return a
}

func (a *Area) Shape(shape string) *Area {
	a.shape = shape
	return a
}

func (a *Area) Target(target string) *Area {
	a.target = target
	return a
}

func (a *Area) Bytes() []byte {
	return cloneBytes(a.buf.Bytes())
}

// IsBodyElement implements BodyElement interface
func (a *Area) IsBodyElement() {}

func (a *Area) Prepare() {
	a.buf.Reset()
	a.buf.WriteString("<area")
	if a.alt != "" {
		writeAttr(&a.buf, "alt", a.alt)
	}
	if a.coords != "" {
		writeAttr(&a.buf, "coords", a.coords)
	}
	if a.href != "" {
		writeAttr(&a.buf, "href", a.href)
	}
	if a.shape != "" {
		writeAttr(&a.buf, "shape", a.shape)
	}
	if a.target != "" {
		writeAttr(&a.buf, "target", a.target)
	}
	if len(a.style) != 0 {
		parseStyle(&a.buf, a.style)
	}
	a.buf.WriteString(">")
}

type Img struct {
	buf    bytes.Buffer
	style  map[string]string
	src    string
	alt    string
	width  string
	height string
	title  string
}

func NewImg() *Img {
	return &Img{
		style: make(map[string]string),
	}
}

func (i *Img) AddStyle(k, v string) *Img {
	i.style[k] = v
	return i
}

func (i *Img) AddStyles(m map[string]string) *Img {
	for k, v := range m {
		i.style[k] = v
	}
	return i
}

func (i *Img) Style(m map[string]string) *Img {
	i.style = m
	return i
}

func (i *Img) Src(src string) *Img {
	i.src = src
	return i
}

func (i *Img) Alt(alt string) *Img {
	i.alt = alt
	return i
}

func (i *Img) Width(width string) *Img {
	i.width = width
	return i
}

func (i *Img) Height(height string) *Img {
	i.height = height
	return i
}

func (i *Img) Title(title string) *Img {
	i.title = title
	return i
}

func (i *Img) Bytes() []byte {
	return cloneBytes(i.buf.Bytes())
}

// IsBodyElement implements BodyElement interface
func (i *Img) IsBodyElement() {}

func (i *Img) Prepare() {
	i.buf.Reset()
	i.buf.WriteString("<img")
	if i.src != "" {
		writeAttr(&i.buf, "src", i.src)
	}
	if i.alt != "" {
		writeAttr(&i.buf, "alt", i.alt)
	}
	if i.width != "" {
		writeAttr(&i.buf, "width", i.width)
	}
	if i.height != "" {
		writeAttr(&i.buf, "height", i.height)
	}
	if i.title != "" {
		writeAttr(&i.buf, "title", i.title)
	}
	if len(i.style) != 0 {
		parseStyle(&i.buf, i.style)
	}
	i.buf.WriteString(">")
}

type Audio struct {
	buf      bytes.Buffer
	style    map[string]string
	contents []Element
	src      string
	controls bool
	autoplay bool
	loop     bool
	muted    bool
	preload  string
}

func NewAudio() *Audio {
	return &Audio{
		style: make(map[string]string),
	}
}

func (a *Audio) AddStyle(k, v string) *Audio {
	a.style[k] = v
	return a
}

func (a *Audio) AddStyles(m map[string]string) *Audio {
	for k, v := range m {
		a.style[k] = v
	}
	return a
}

func (a *Audio) Style(m map[string]string) *Audio {
	a.style = m
	return a
}

func (a *Audio) Add(e Element) *Audio {
	a.contents = appendElement(a.contents, e)
	return a
}

func (a *Audio) Src(src string) *Audio {
	a.src = src
	return a
}

func (a *Audio) Controls(controls bool) *Audio {
	a.controls = controls
	return a
}

func (a *Audio) Autoplay(autoplay bool) *Audio {
	a.autoplay = autoplay
	return a
}

func (a *Audio) Loop(loop bool) *Audio {
	a.loop = loop
	return a
}

func (a *Audio) Muted(muted bool) *Audio {
	a.muted = muted
	return a
}

func (a *Audio) Preload(preload string) *Audio {
	a.preload = preload
	return a
}

func (a *Audio) Bytes() []byte {
	return cloneBytes(a.buf.Bytes())
}

// IsBodyElement implements BodyElement interface
func (a *Audio) IsBodyElement() {}

func (a *Audio) Prepare() {
	a.buf.Reset()
	a.buf.WriteString("<audio")
	if a.src != "" {
		writeAttr(&a.buf, "src", a.src)
	}
	if a.controls {
		a.buf.WriteString(" controls")
	}
	if a.autoplay {
		a.buf.WriteString(" autoplay")
	}
	if a.loop {
		a.buf.WriteString(" loop")
	}
	if a.muted {
		a.buf.WriteString(" muted")
	}
	if a.preload != "" {
		writeAttr(&a.buf, "preload", a.preload)
	}
	if len(a.style) != 0 {
		parseStyle(&a.buf, a.style)
	}
	a.buf.WriteByte('>')

	writeElements(&a.buf, a.contents)
	a.buf.WriteString("</audio>")
}

type Track struct {
	buf        bytes.Buffer
	style      map[string]string
	src        string
	kind       string
	srclang    string
	label      string
	defaultVal bool
}

func NewTrack() *Track {
	return &Track{
		style: make(map[string]string),
	}
}

func (t *Track) AddStyle(k, v string) *Track {
	t.style[k] = v
	return t
}

func (t *Track) AddStyles(m map[string]string) *Track {
	for k, v := range m {
		t.style[k] = v
	}
	return t
}

func (t *Track) Style(m map[string]string) *Track {
	t.style = m
	return t
}

func (t *Track) Src(src string) *Track {
	t.src = src
	return t
}

func (t *Track) Kind(kind string) *Track {
	t.kind = kind
	return t
}

func (t *Track) Srclang(srclang string) *Track {
	t.srclang = srclang
	return t
}

func (t *Track) Label(label string) *Track {
	t.label = label
	return t
}

func (t *Track) Default(def bool) *Track {
	t.defaultVal = def
	return t
}

func (t *Track) Bytes() []byte {
	return cloneBytes(t.buf.Bytes())
}

// IsBodyElement implements BodyElement interface
func (t *Track) IsBodyElement() {}

func (t *Track) Prepare() {
	t.buf.Reset()
	t.buf.WriteString("<track")
	if t.src != "" {
		writeAttr(&t.buf, "src", t.src)
	}
	if t.kind != "" {
		writeAttr(&t.buf, "kind", t.kind)
	}
	if t.srclang != "" {
		writeAttr(&t.buf, "srclang", t.srclang)
	}
	if t.label != "" {
		writeAttr(&t.buf, "label", t.label)
	}
	if t.defaultVal {
		t.buf.WriteString(" default")
	}
	if len(t.style) != 0 {
		parseStyle(&t.buf, t.style)
	}
	t.buf.WriteString(">")
}

type Map struct {
	buf      bytes.Buffer
	style    map[string]string
	contents []Element
	name     string
}

func NewMap() *Map {
	return &Map{
		style: make(map[string]string),
	}
}

func (m *Map) AddStyle(k, v string) *Map {
	m.style[k] = v
	return m
}

func (m *Map) AddStyles(ms map[string]string) *Map {
	for k, v := range ms {
		m.style[k] = v
	}
	return m
}

func (m *Map) Style(ms map[string]string) *Map {
	m.style = ms
	return m
}

func (m *Map) Add(e Element) *Map {
	m.contents = appendElement(m.contents, e)
	return m
}

func (m *Map) Name(name string) *Map {
	m.name = name
	return m
}

func (m *Map) Bytes() []byte {
	return cloneBytes(m.buf.Bytes())
}

// IsBodyElement implements BodyElement interface
func (m *Map) IsBodyElement() {}

func (m *Map) Prepare() {
	m.buf.Reset()
	m.buf.WriteString("<map")
	if m.name != "" {
		writeAttr(&m.buf, "name", m.name)
	}
	if len(m.style) != 0 {
		parseStyle(&m.buf, m.style)
	}
	m.buf.WriteByte('>')

	writeElements(&m.buf, m.contents)
	m.buf.WriteString("</map>")
}

type Video struct {
	buf      bytes.Buffer
	style    map[string]string
	contents []Element
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

func NewVideo() *Video {
	return &Video{
		style: make(map[string]string),
	}
}

func (v *Video) AddStyle(k, val string) *Video {
	v.style[k] = val
	return v
}

func (v *Video) AddStyles(m map[string]string) *Video {
	for k, val := range m {
		v.style[k] = val
	}
	return v
}

func (v *Video) Style(m map[string]string) *Video {
	v.style = m
	return v
}

func (v *Video) Add(e Element) *Video {
	v.contents = appendElement(v.contents, e)
	return v
}

func (v *Video) Src(src string) *Video {
	v.src = src
	return v
}

func (v *Video) Controls(controls bool) *Video {
	v.controls = controls
	return v
}

func (v *Video) Autoplay(autoplay bool) *Video {
	v.autoplay = autoplay
	return v
}

func (v *Video) Loop(loop bool) *Video {
	v.loop = loop
	return v
}

func (v *Video) Muted(muted bool) *Video {
	v.muted = muted
	return v
}

func (v *Video) Preload(preload string) *Video {
	v.preload = preload
	return v
}

func (v *Video) Width(width string) *Video {
	v.width = width
	return v
}

func (v *Video) Height(height string) *Video {
	v.height = height
	return v
}

func (v *Video) Poster(poster string) *Video {
	v.poster = poster
	return v
}

func (v *Video) Bytes() []byte {
	return cloneBytes(v.buf.Bytes())
}

// IsBodyElement implements BodyElement interface
func (v *Video) IsBodyElement() {}

func (v *Video) Prepare() {
	v.buf.Reset()
	v.buf.WriteString("<video")
	if v.src != "" {
		writeAttr(&v.buf, "src", v.src)
	}
	if v.controls {
		v.buf.WriteString(" controls")
	}
	if v.autoplay {
		v.buf.WriteString(" autoplay")
	}
	if v.loop {
		v.buf.WriteString(" loop")
	}
	if v.muted {
		v.buf.WriteString(" muted")
	}
	if v.preload != "" {
		writeAttr(&v.buf, "preload", v.preload)
	}
	if v.width != "" {
		writeAttr(&v.buf, "width", v.width)
	}
	if v.height != "" {
		writeAttr(&v.buf, "height", v.height)
	}
	if v.poster != "" {
		writeAttr(&v.buf, "poster", v.poster)
	}
	if len(v.style) != 0 {
		parseStyle(&v.buf, v.style)
	}
	v.buf.WriteByte('>')

	writeElements(&v.buf, v.contents)
	v.buf.WriteString("</video>")
}
