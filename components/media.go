package rephtml

import "bytes"

// Area represents the Area component or supporting type.
type Area struct {
	buf    bytes.Buffer
	style  map[string]string
	alt    string
	coords string
	href   string
	shape  string
	target string
}

// NewArea creates a new Area component.
func NewArea() *Area {
	return &Area{
		style: make(map[string]string),
	}
}

// AddStyle adds one inline CSS declaration to the Area component.
func (a *Area) AddStyle(k, v string) *Area {
	a.style[k] = v
	return a
}

// AddStyles adds multiple inline CSS declarations to the Area component.
func (a *Area) AddStyles(m map[string]string) *Area {
	for k, v := range m {
		a.style[k] = v
	}
	return a
}

// Style replaces the inline CSS declarations on the Area component.
func (a *Area) Style(m map[string]string) *Area {
	a.style = m
	return a
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

// Bytes returns a defensive copy of the rendered Area bytes.
func (a *Area) Bytes() []byte {
	return cloneBytes(a.buf.Bytes())
}

// IsBodyElement implements BodyElement interface
func (a *Area) IsBodyElement() {}

// Prepare renders the Area component into its internal buffer.
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

// Img represents the Img component or supporting type.
type Img struct {
	buf    bytes.Buffer
	style  map[string]string
	src    string
	alt    string
	width  string
	height string
	title  string
}

// NewImg creates a new Img component.
func NewImg() *Img {
	return &Img{
		style: make(map[string]string),
	}
}

// AddStyle adds one inline CSS declaration to the Img component.
func (i *Img) AddStyle(k, v string) *Img {
	i.style[k] = v
	return i
}

// AddStyles adds multiple inline CSS declarations to the Img component.
func (i *Img) AddStyles(m map[string]string) *Img {
	for k, v := range m {
		i.style[k] = v
	}
	return i
}

// Style replaces the inline CSS declarations on the Img component.
func (i *Img) Style(m map[string]string) *Img {
	i.style = m
	return i
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

// Bytes returns a defensive copy of the rendered Img bytes.
func (i *Img) Bytes() []byte {
	return cloneBytes(i.buf.Bytes())
}

// IsBodyElement implements BodyElement interface
func (i *Img) IsBodyElement() {}

// Prepare renders the Img component into its internal buffer.
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

// Audio represents the Audio component or supporting type.
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

// NewAudio creates a new Audio component.
func NewAudio() *Audio {
	return &Audio{
		style: make(map[string]string),
	}
}

// AddStyle adds one inline CSS declaration to the Audio component.
func (a *Audio) AddStyle(k, v string) *Audio {
	a.style[k] = v
	return a
}

// AddStyles adds multiple inline CSS declarations to the Audio component.
func (a *Audio) AddStyles(m map[string]string) *Audio {
	for k, v := range m {
		a.style[k] = v
	}
	return a
}

// Style replaces the inline CSS declarations on the Audio component.
func (a *Audio) Style(m map[string]string) *Audio {
	a.style = m
	return a
}

// Add appends child content to the Audio component.
func (a *Audio) Add(e Element) *Audio {
	a.contents = appendElement(a.contents, e)
	return a
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

// Bytes returns a defensive copy of the rendered Audio bytes.
func (a *Audio) Bytes() []byte {
	return cloneBytes(a.buf.Bytes())
}

// IsBodyElement implements BodyElement interface
func (a *Audio) IsBodyElement() {}

// Prepare renders the Audio component into its internal buffer.
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

// Track represents the Track component or supporting type.
type Track struct {
	buf        bytes.Buffer
	style      map[string]string
	src        string
	kind       string
	srclang    string
	label      string
	defaultVal bool
}

// NewTrack creates a new Track component.
func NewTrack() *Track {
	return &Track{
		style: make(map[string]string),
	}
}

// AddStyle adds one inline CSS declaration to the Track component.
func (t *Track) AddStyle(k, v string) *Track {
	t.style[k] = v
	return t
}

// AddStyles adds multiple inline CSS declarations to the Track component.
func (t *Track) AddStyles(m map[string]string) *Track {
	for k, v := range m {
		t.style[k] = v
	}
	return t
}

// Style replaces the inline CSS declarations on the Track component.
func (t *Track) Style(m map[string]string) *Track {
	t.style = m
	return t
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

// Bytes returns a defensive copy of the rendered Track bytes.
func (t *Track) Bytes() []byte {
	return cloneBytes(t.buf.Bytes())
}

// IsBodyElement implements BodyElement interface
func (t *Track) IsBodyElement() {}

// Prepare renders the Track component into its internal buffer.
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

// Map represents the Map component or supporting type.
type Map struct {
	buf      bytes.Buffer
	style    map[string]string
	contents []Element
	name     string
}

// NewMap creates a new Map component.
func NewMap() *Map {
	return &Map{
		style: make(map[string]string),
	}
}

// AddStyle adds one inline CSS declaration to the Map component.
func (m *Map) AddStyle(k, v string) *Map {
	m.style[k] = v
	return m
}

// AddStyles adds multiple inline CSS declarations to the Map component.
func (m *Map) AddStyles(ms map[string]string) *Map {
	for k, v := range ms {
		m.style[k] = v
	}
	return m
}

// Style replaces the inline CSS declarations on the Map component.
func (m *Map) Style(ms map[string]string) *Map {
	m.style = ms
	return m
}

// Add appends child content to the Map component.
func (m *Map) Add(e Element) *Map {
	m.contents = appendElement(m.contents, e)
	return m
}

// Name sets the name value on the Map component.
func (m *Map) Name(name string) *Map {
	m.name = name
	return m
}

// Bytes returns a defensive copy of the rendered Map bytes.
func (m *Map) Bytes() []byte {
	return cloneBytes(m.buf.Bytes())
}

// IsBodyElement implements BodyElement interface
func (m *Map) IsBodyElement() {}

// Prepare renders the Map component into its internal buffer.
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

// Video represents the Video component or supporting type.
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

// NewVideo creates a new Video component.
func NewVideo() *Video {
	return &Video{
		style: make(map[string]string),
	}
}

// AddStyle adds one inline CSS declaration to the Video component.
func (v *Video) AddStyle(k, val string) *Video {
	v.style[k] = val
	return v
}

// AddStyles adds multiple inline CSS declarations to the Video component.
func (v *Video) AddStyles(m map[string]string) *Video {
	for k, val := range m {
		v.style[k] = val
	}
	return v
}

// Style replaces the inline CSS declarations on the Video component.
func (v *Video) Style(m map[string]string) *Video {
	v.style = m
	return v
}

// Add appends child content to the Video component.
func (v *Video) Add(e Element) *Video {
	v.contents = appendElement(v.contents, e)
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

// Bytes returns a defensive copy of the rendered Video bytes.
func (v *Video) Bytes() []byte {
	return cloneBytes(v.buf.Bytes())
}

// IsBodyElement implements BodyElement interface
func (v *Video) IsBodyElement() {}

// Prepare renders the Video component into its internal buffer.
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
