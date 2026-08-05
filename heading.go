package rephtml

// Hgroup represents the Hgroup component or supporting type.
type Hgroup struct {
	bodyElement
	contentNode[*Hgroup]
}

// NewHgroup creates a new Hgroup component.
func NewHgroup() *Hgroup {
	v := &Hgroup{}
	v.init(v)
	return v
}

// prepare renders the Hgroup component into its internal buffer.
func (h *Hgroup) prepare() {
	tg := openTag(&h.buf, "hgroup")
	tg.styleAttr(h.style)
	tg.children(h.contents)
}

// H1 represents the H1 component or supporting type.
type H1 struct {
	bodyElement
	textNode[*H1]
}

// NewH1 creates a new H1 component.
func NewH1() *H1 {
	v := &H1{}
	v.init(v)
	return v
}

// prepare renders the H1 component into its internal buffer.
func (h *H1) prepare() {
	tg := openTag(&h.buf, "h1")
	tg.styleAttr(h.style)
	tg.text(h.text)
}

// H2 represents the H2 component or supporting type.
type H2 struct {
	bodyElement
	textNode[*H2]
}

// NewH2 creates a new H2 component.
func NewH2() *H2 {
	v := &H2{}
	v.init(v)
	return v
}

// prepare renders the H2 component into its internal buffer.
func (h *H2) prepare() {
	tg := openTag(&h.buf, "h2")
	tg.styleAttr(h.style)
	tg.text(h.text)
}

// H3 represents the H3 component or supporting type.
type H3 struct {
	bodyElement
	textNode[*H3]
}

// NewH3 creates a new H3 component.
func NewH3() *H3 {
	v := &H3{}
	v.init(v)
	return v
}

// prepare renders the H3 component into its internal buffer.
func (h *H3) prepare() {
	tg := openTag(&h.buf, "h3")
	tg.styleAttr(h.style)
	tg.text(h.text)
}

// H4 represents the H4 component or supporting type.
type H4 struct {
	bodyElement
	textNode[*H4]
}

// NewH4 creates a new H4 component.
func NewH4() *H4 {
	v := &H4{}
	v.init(v)
	return v
}

// prepare renders the H4 component into its internal buffer.
func (h *H4) prepare() {
	tg := openTag(&h.buf, "h4")
	tg.styleAttr(h.style)
	tg.text(h.text)
}

// H5 represents the H5 component or supporting type.
type H5 struct {
	bodyElement
	textNode[*H5]
}

// NewH5 creates a new H5 component.
func NewH5() *H5 {
	v := &H5{}
	v.init(v)
	return v
}

// prepare renders the H5 component into its internal buffer.
func (h *H5) prepare() {
	tg := openTag(&h.buf, "h5")
	tg.styleAttr(h.style)
	tg.text(h.text)
}

// H6 represents the H6 component or supporting type.
type H6 struct {
	bodyElement
	textNode[*H6]
}

// NewH6 creates a new H6 component.
func NewH6() *H6 {
	v := &H6{}
	v.init(v)
	return v
}

// prepare renders the H6 component into its internal buffer.
func (h *H6) prepare() {
	tg := openTag(&h.buf, "h6")
	tg.styleAttr(h.style)
	tg.text(h.text)
}
