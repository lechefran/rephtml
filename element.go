package rephtml

import (
	"bytes"
	"maps"
)

// The types in this file hold the machinery every element shares: an inline
// style map, child content, and the Render and HTML entry points. Elements
// embed the tier they need and declare only what makes them different, which is
// their attributes and their renderTo method.
//
// Each tier is parameterised by the concrete element pointer type, named Self.
// That is what lets a promoted setter return the element's own type, so
// NewDiv().AddStyle("color", "red").Add(child) still type-checks as a *Div.
// The constructor is responsible for handing the base that pointer, which is
// what init does.

// base holds the concrete element pointer.
//
// Render and HTML dispatch through self, so an element's own renderTo method is
// what runs even though the entry points live here. The element itself holds no
// render state: the buffer belongs to whoever started the render.
type base[Self selfElement] struct {
	self Self
}

// init records the concrete element pointer. Constructors must call it before
// returning, or the element has no way to reach its own renderTo method.
func (b *base[Self]) init(self Self) {
	b.self = self
}

// renderSelf writes the concrete element's HTML to buf.
//
// Elements reach their own renderTo through the pointer recorded by init, so an
// element built as a bare composite literal instead of through its constructor
// has nothing to dispatch to. Reporting that directly is far more useful than
// the nil dereference it would otherwise become.
func (b *base[Self]) renderSelf(buf *bytes.Buffer) {
	var unbound Self
	if b.self == unbound {
		panic("rephtml: element was not created with its New* constructor")
	}
	b.self.renderTo(buf)
}

// Render returns the element's HTML bytes.
//
// The returned slice is the caller's own copy, so it may be retained or
// modified without affecting the element or any later render.
func (b *base[Self]) Render() []byte {
	buf := getBuffer()
	defer putBuffer(buf)
	b.renderSelf(buf)
	return cloneBytes(buf.Bytes())
}

// HTML returns the element's HTML as a string.
func (b *base[Self]) HTML() string {
	buf := getBuffer()
	defer putBuffer(buf)
	b.renderSelf(buf)
	return buf.String()
}

// node adds the global id, class and inline style attributes.
type node[Self selfElement] struct {
	base[Self]
	globalAttrs
}

// globalAttrs holds the attributes every element shares.
//
// It is a plain struct rather than part of the generic node so that attrs and
// classStyle compile to ordinary, inlinable methods. Every element calls one of
// them on each render.
type globalAttrs struct {
	style StyleMap
	class []string
	id    string
}

// attrs writes the id, class and style attributes, in that order, omitting
// any that are unset. Elements call it where their inline style belongs.
//
// The check is kept separate from the writing so that it inlines: most
// elements set none of these, and then rendering them costs no call.
func (a *globalAttrs) attrs(tg tag) {
	if a.id != "" || len(a.class) != 0 || len(a.style) != 0 {
		a.writeAttrs(tg, true)
	}
}

// classStyle writes the class and style attributes, for an element that writes
// its id attribute at a position of its own.
func (a *globalAttrs) classStyle(tg tag) {
	if len(a.class) != 0 || len(a.style) != 0 {
		a.writeAttrs(tg, false)
	}
}

// writeAttrs writes the attributes for attrs and classStyle.
func (a *globalAttrs) writeAttrs(tg tag, withID bool) {
	if withID {
		tg.attr("id", a.id)
	}
	tg.classAttr(a.class)
	tg.styleAttr(a.style)
}

// init records the concrete element pointer and prepares the style map.
func (n *node[Self]) init(self Self) {
	n.self = self
	n.style = make(StyleMap)
}

// AddStyle adds one inline CSS declaration.
func (n *node[Self]) AddStyle(k, v string) Self {
	n.style[k] = v
	return n.self
}

// AddStyles adds multiple inline CSS declarations.
func (n *node[Self]) AddStyles(m StyleMap) Self {
	maps.Copy(n.style, m)
	return n.self
}

// Style replaces the inline CSS declarations, copying the caller's map.
func (n *node[Self]) Style(m StyleMap) Self {
	n.style = cloneStyleMap(m)
	return n.self
}

// AddClass appends one class name.
func (n *node[Self]) AddClass(s string) Self {
	n.class = append(n.class, s)
	return n.self
}

// AddClasses appends multiple class names.
func (n *node[Self]) AddClasses(s []string) Self {
	n.class = append(n.class, s...)
	return n.self
}

// Class replaces the class names, copying the caller's slice.
func (n *node[Self]) Class(s []string) Self {
	n.class = cloneStrings(s)
	return n.self
}

// Id sets the id attribute.
func (n *node[Self]) Id(s string) Self {
	n.id = s
	return n.self
}

// contentNode adds child elements.
type contentNode[Self selfElement] struct {
	node[Self]
	contents []Element
}

// Add appends child content, ignoring nil children.
func (n *contentNode[Self]) Add(e Element) Self {
	n.contents = appendElement(n.contents, e)
	return n.self
}

// collectErrors appends the errors recorded anywhere among the children.
func (n *contentNode[Self]) collectErrors(errs []error) []error {
	return collectChildErrors(errs, n.contents)
}

// errorReporter is implemented by elements that can record an error while
// being built, or that hold children which can. HtmlFile.Err walks the document
// through it, so an error recorded deep in the tree, such as an invalid Time
// datetime, stops the document from rendering instead of being lost.
type errorReporter interface {
	collectErrors(errs []error) []error
}

// collectChildErrors appends the errors recorded by each child and its
// descendants, in document order. Elements from outside the package are not
// inspected.
func collectChildErrors(errs []error, contents []Element) []error {
	for _, e := range contents {
		if r, ok := e.(errorReporter); ok {
			errs = r.collectErrors(errs)
		}
	}
	return errs
}

// appendError appends err when it is not nil.
func appendError(errs []error, err error) []error {
	if err != nil {
		errs = append(errs, err)
	}
	return errs
}

// textContent is a content element whose text is added as escaped child
// content, so text and child elements can be interleaved.
type textContent[Self selfElement] struct {
	contentNode[Self]
}

// Text sets the escaped text content, replacing any text set before. Child
// elements are kept: the text takes the place of the first earlier text, or
// goes after the children when there was none. Use AddText to append.
func (n *textContent[Self]) Text(s string) Self {
	n.contents = replaceText(n.contents, escapedText(s))
	return n.self
}

// AddText appends escaped text content after any existing content.
func (n *textContent[Self]) AddText(s string) Self {
	n.contents = appendElement(n.contents, escapedText(s))
	return n.self
}

// textNode adds a single escaped text body.
//
// Elements that are nothing more than a tag around text record their tag name
// with initTag and use the shared renderTo below. Elements with attributes of
// their own call init and define their own renderTo, which takes precedence.
type textNode[Self selfElement] struct {
	node[Self]
	name string
	text string
}

// initTag records the concrete element pointer and the tag name that the
// shared renderTo writes.
func (n *textNode[Self]) initTag(self Self, name string) {
	n.init(self)
	n.name = name
}

// renderTo writes the element as its tag, inline style and escaped text.
func (n *textNode[Self]) renderTo(buf *bytes.Buffer) {
	tg := startTag(buf, n.name)
	n.attrs(tg)
	tg.text(n.text)
}

// Text replaces the text content.
func (n *textNode[Self]) Text(s string) Self {
	n.text = s
	return n.self
}

// AddText appends to the text content.
func (n *textNode[Self]) AddText(s string) Self {
	n.text += s
	return n.self
}

// bodyElement marks an element as valid document body content.
//
// It is a mixin rather than part of the tier chain because an element can be
// valid in both the head and the body: Script, Noscript, Template and
// StyleElement embed both markers.
type bodyElement struct{}

// IsBodyElement implements BodyElement.
func (bodyElement) IsBodyElement() {}

// headElement marks an element as valid document head content.
type headElement struct{}

// IsHeadElement implements HeadElement.
func (headElement) IsHeadElement() {}
