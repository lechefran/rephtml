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
	return a.buf.Bytes()
}

func (a *Area) Prepare() {
	a.buf.WriteString("<area")
	if a.alt != "" {
		a.buf.WriteString(" alt=\"" + a.alt + "\"")
	}
	if a.coords != "" {
		a.buf.WriteString(" coords=\"" + a.coords + "\"")
	}
	if a.href != "" {
		a.buf.WriteString(" href=\"" + a.href + "\"")
	}
	if a.shape != "" {
		a.buf.WriteString(" shape=\"" + a.shape + "\"")
	}
	if a.target != "" {
		a.buf.WriteString(" target=\"" + a.target + "\"")
	}
	if len(a.style) != 0 {
		idx := 0
		a.buf.WriteString(" style=\"")
		for k, v := range a.style {
			a.buf.WriteString(k + ": " + v + ";")
			if idx != len(a.style)-1 {
				a.buf.WriteByte(' ')
			}
			idx++
		}
		a.buf.WriteString("\"")
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
	return i.buf.Bytes()
}

func (i *Img) Prepare() {
	i.buf.WriteString("<img")
	if i.src != "" {
		i.buf.WriteString(" src=\"" + i.src + "\"")
	}
	if i.alt != "" {
		i.buf.WriteString(" alt=\"" + i.alt + "\"")
	}
	if i.width != "" {
		i.buf.WriteString(" width=\"" + i.width + "\"")
	}
	if i.height != "" {
		i.buf.WriteString(" height=\"" + i.height + "\"")
	}
	if i.title != "" {
		i.buf.WriteString(" title=\"" + i.title + "\"")
	}
	if len(i.style) != 0 {
		idx := 0
		i.buf.WriteString(" style=\"")
		for k, v := range i.style {
			i.buf.WriteString(k + ": " + v + ";")
			if idx != len(i.style)-1 {
				i.buf.WriteByte(' ')
			}
			idx++
		}
		i.buf.WriteString("\"")
	}
	i.buf.WriteString(">")
}

type Audio struct {
	buf      bytes.Buffer
	style    map[string]string
	contents [][]byte
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

func (a *Audio) Add(e Elements) *Audio {
	a.contents = append(a.contents, e.Bytes())
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
	return a.buf.Bytes()
}

func (a *Audio) Prepare() {
	a.buf.WriteString("<audio")
	if a.src != "" {
		a.buf.WriteString(" src=\"" + a.src + "\"")
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
		a.buf.WriteString(" preload=\"" + a.preload + "\"")
	}
	if len(a.style) != 0 {
		idx := 0
		a.buf.WriteString(" style=\"")
		for k, v := range a.style {
			a.buf.WriteString(k + ": " + v + ";")
			if idx != len(a.style)-1 {
				a.buf.WriteByte(' ')
			}
			idx++
		}
		a.buf.WriteString("\"")
	}
	a.buf.WriteByte('>')

	for _, content := range a.contents {
		a.buf.Write(content)
	}
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
	return t.buf.Bytes()
}

func (t *Track) Prepare() {
	t.buf.WriteString("<track")
	if t.src != "" {
		t.buf.WriteString(" src=\"" + t.src + "\"")
	}
	if t.kind != "" {
		t.buf.WriteString(" kind=\"" + t.kind + "\"")
	}
	if t.srclang != "" {
		t.buf.WriteString(" srclang=\"" + t.srclang + "\"")
	}
	if t.label != "" {
		t.buf.WriteString(" label=\"" + t.label + "\"")
	}
	if t.defaultVal {
		t.buf.WriteString(" default")
	}
	if len(t.style) != 0 {
		idx := 0
		t.buf.WriteString(" style=\"")
		for k, v := range t.style {
			t.buf.WriteString(k + ": " + v + ";")
			if idx != len(t.style)-1 {
				t.buf.WriteByte(' ')
			}
			idx++
		}
		t.buf.WriteString("\"")
	}
	t.buf.WriteString(">")
}

type Map struct {
	buf      bytes.Buffer
	style    map[string]string
	contents [][]byte
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

func (m *Map) Add(e Elements) *Map {
	m.contents = append(m.contents, e.Bytes())
	return m
}

func (m *Map) Name(name string) *Map {
	m.name = name
	return m
}

func (m *Map) Bytes() []byte {
	return m.buf.Bytes()
}

func (m *Map) Prepare() {
	m.buf.WriteString("<map")
	if m.name != "" {
		m.buf.WriteString(" name=\"" + m.name + "\"")
	}
	if len(m.style) != 0 {
		idx := 0
		m.buf.WriteString(" style=\"")
		for k, v := range m.style {
			m.buf.WriteString(k + ": " + v + ";")
			if idx != len(m.style)-1 {
				m.buf.WriteByte(' ')
			}
			idx++
		}
		m.buf.WriteString("\"")
	}
	m.buf.WriteByte('>')

	for _, content := range m.contents {
		m.buf.Write(content)
	}
	m.buf.WriteString("</map>")
}
