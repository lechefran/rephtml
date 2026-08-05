package rephtml

// Details represents the HTML details element for creating a disclosure widget
type Details struct {
	bodyElement
	contentNode[*Details]
	open bool
}

// NewDetails creates a new Details element
func NewDetails() *Details {
	v := &Details{}
	v.init(v)
	return v
}

// Prepare builds the HTML for the details element
func (d *Details) prepare() {
	tg := openTag(&d.buf, "details")
	tg.boolAttr("open", d.open)
	tg.styleAttr(d.style)
	tg.children(d.contents)
}

// Open sets the open attribute
func (d *Details) Open(open bool) *Details {
	d.open = open
	return d
}

// Dialog represents the HTML dialog element for creating modal dialogs
type Dialog struct {
	bodyElement
	contentNode[*Dialog]
	open bool
}

// NewDialog creates a new Dialog element
func NewDialog() *Dialog {
	v := &Dialog{}
	v.init(v)
	return v
}

// Prepare builds the HTML for the dialog element
func (d *Dialog) prepare() {
	tg := openTag(&d.buf, "dialog")
	tg.boolAttr("open", d.open)
	tg.styleAttr(d.style)
	tg.children(d.contents)
}

// Open sets the open attribute
func (d *Dialog) Open(open bool) *Dialog {
	d.open = open
	return d
}

// Summary represents the HTML summary element for details disclosure summary
type Summary struct {
	bodyElement
	contentNode[*Summary]
}

// NewSummary creates a new Summary element
func NewSummary() *Summary {
	v := &Summary{}
	v.init(v)
	return v
}

// Prepare builds the HTML for the summary element
func (s *Summary) prepare() {
	tg := openTag(&s.buf, "summary")
	tg.styleAttr(s.style)
	tg.children(s.contents)
}
