package rephtml

import (
	"bytes"
	"maps"
)

// The types in this file hold the machinery every element shares: a render
// buffer, an inline style map, child content, and the render entry points.
// Elements embed the tier they need and declare only what makes them
// different, which is their attributes and their prepare method.
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

// node adds inline CSS declarations.
type node[Self selfElement] struct {
	base[Self]
	style StyleMap
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

// textNode adds a single escaped text body.
type textNode[Self selfElement] struct {
	node[Self]
	text string
}

// Text replaces the text content.
func (n *textNode[Self]) Text(s string) Self {
	n.text = s
	return n.self
}

// tabular adds the id and class attributes shared by the table elements.
type tabular[Self selfElement] struct {
	node[Self]
	class []string
	id    string
}

// AddClass appends one class name.
func (n *tabular[Self]) AddClass(s string) Self {
	n.class = append(n.class, s)
	return n.self
}

// AddClasses appends multiple class names.
func (n *tabular[Self]) AddClasses(s []string) Self {
	n.class = append(n.class, s...)
	return n.self
}

// Class replaces the class names, copying the caller's slice.
func (n *tabular[Self]) Class(s []string) Self {
	n.class = cloneStrings(s)
	return n.self
}

// Id sets the id attribute.
func (n *tabular[Self]) Id(s string) Self {
	n.id = s
	return n.self
}

// AddId sets the id attribute.
//
// Deprecated: an element has at most one id. Use Id.
func (n *tabular[Self]) AddId(s string) Self {
	return n.Id(s)
}

// Styles replaces the inline CSS declarations.
//
// Deprecated: use Style, which every other element uses for this.
func (n *tabular[Self]) Styles(m StyleMap) Self {
	return n.Style(m)
}

// tabularNode is a table element that also holds child elements.
type tabularNode[Self selfElement] struct {
	tabular[Self]
	contents []Element
}

// Add appends child content, ignoring nil children.
func (n *tabularNode[Self]) Add(e Element) Self {
	n.contents = appendElement(n.contents, e)
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
