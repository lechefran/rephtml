package rephtml

import (
	"bytes"
	"slices"
	"strconv"
	"strings"
)

// StyleMap holds CSS declarations as property/value pairs.
//
// Property names are emitted exactly as provided, so standard declarations,
// custom properties such as --brand-color, and vendor-prefixed properties all
// use the same representation.
type StyleMap map[string]string

// NewStyleMap returns an empty StyleMap.
func NewStyleMap() StyleMap {
	return StyleMap{}
}

// Style renders a complete style element for one selector rule.
type Style struct {
	Props StyleMap
	Tags  []string
}

// NewStyle creates a style element for one CSS selector rule.
func NewStyle(tags ...string) *Style {
	return &Style{
		Props: StyleMap{},
		Tags:  append([]string(nil), tags...),
	}
}

// Render returns freshly prepared Style HTML bytes.
func (s *Style) Render() []byte {
	return renderElement(s)
}

// HTML returns freshly prepared Style HTML as a string.
func (s *Style) HTML() string {
	return htmlElement(s)
}

// renderTo writes the Style component's HTML to buf.
func (s *Style) renderTo(buf *bytes.Buffer) {
	buf.WriteString("<style>")
	buf.WriteString(formatStyleRule(s.Tags, s.Props))
	buf.WriteString("</style>")
}

// AddStyle adds one CSS declaration to the style rule.
func (s *Style) AddStyle(k, v string) *Style {
	if s.Props == nil {
		s.Props = StyleMap{}
	}
	s.Props[k] = v
	return s
}

// AddStyles adds CSS declarations to the style rule.
func (s *Style) AddStyles(styles StyleMap) *Style {
	if s.Props == nil {
		s.Props = StyleMap{}
	}
	for k, v := range styles {
		s.Props[k] = v
	}
	return s
}

// Style replaces CSS declarations on the style rule.
func (s *Style) Style(styles StyleMap) *Style {
	s.Props = cloneStyleMap(styles)
	return s
}

// CSSRule represents renderable CSS that can be placed inside a style element.
type CSSRule interface {
	Element
	IsCSSRule()
}

// StyleRule renders one CSS selector rule without wrapping it in a style tag.
type StyleRule struct {
	Props StyleMap
	Tags  []string
}

// NewStyleRule creates one unwrapped selector rule.
func NewStyleRule(tags ...string) *StyleRule {
	return &StyleRule{
		Props: StyleMap{},
		Tags:  append([]string(nil), tags...),
	}
}

// Render returns freshly prepared StyleRule HTML bytes.
func (s *StyleRule) Render() []byte {
	return renderElement(s)
}

// HTML returns freshly prepared StyleRule HTML as a string.
func (s *StyleRule) HTML() string {
	return htmlElement(s)
}

// renderTo writes the StyleRule component's HTML to buf.
func (s *StyleRule) renderTo(buf *bytes.Buffer) {
	buf.WriteString(formatStyleRule(s.Tags, s.Props))
}

// AddStyle adds one CSS declaration to the rule.
func (s *StyleRule) AddStyle(k, v string) *StyleRule {
	if s.Props == nil {
		s.Props = StyleMap{}
	}
	s.Props[k] = v
	return s
}

// AddStyles adds CSS declarations to the rule.
func (s *StyleRule) AddStyles(styles StyleMap) *StyleRule {
	if s.Props == nil {
		s.Props = StyleMap{}
	}
	for k, v := range styles {
		s.Props[k] = v
	}
	return s
}

// Style replaces CSS declarations on the rule.
func (s *StyleRule) Style(styles StyleMap) *StyleRule {
	s.Props = cloneStyleMap(styles)
	return s
}

// IsCSSRule marks StyleRule as CSS content for style elements.
func (s *StyleRule) IsCSSRule() {}

// RawCSSRule represents raw CSS content for custom or unsupported rules.
type RawCSSRule struct {
	css string
}

// NewRawCSSRule creates a raw CSS rule.
func NewRawCSSRule(css string) *RawCSSRule {
	return &RawCSSRule{css: css}
}

// Text replaces the raw CSS content.
func (r *RawCSSRule) Text(css string) *RawCSSRule {
	r.css = css
	return r
}

// Render returns freshly prepared RawCSSRule HTML bytes.
func (r *RawCSSRule) Render() []byte {
	return renderElement(r)
}

// HTML returns freshly prepared RawCSSRule HTML as a string.
func (r *RawCSSRule) HTML() string {
	return htmlElement(r)
}

// renderTo writes the RawCSSRule component's HTML to buf.
func (r *RawCSSRule) renderTo(buf *bytes.Buffer) {
	css := strings.Trim(r.css, "\n")
	if css == "" {
		return
	}
	buf.WriteByte('\n')
	buf.WriteString(css)
	buf.WriteByte('\n')
}

// IsCSSRule marks RawCSSRule as CSS content for style elements.
func (r *RawCSSRule) IsCSSRule() {}

// CharsetRule represents a CSS @charset rule.
type CharsetRule struct {
	Charset string
}

// NewCharsetRule creates a new @charset rule.
func NewCharsetRule(charset string) *CharsetRule {
	return &CharsetRule{Charset: charset}
}

// Render returns freshly prepared CharsetRule HTML bytes.
func (c *CharsetRule) Render() []byte {
	return renderElement(c)
}

// HTML returns freshly prepared CharsetRule HTML as a string.
func (c *CharsetRule) HTML() string {
	return htmlElement(c)
}

// renderTo writes the CharsetRule component's HTML to buf.
func (c *CharsetRule) renderTo(buf *bytes.Buffer) {
	if c.Charset == "" {
		return
	}
	buf.WriteString("\n@charset ")
	buf.WriteString(quoteCSSString(c.Charset))
	buf.WriteString(";\n")
}

// IsCSSRule marks CharsetRule as CSS content for style elements.
func (c *CharsetRule) IsCSSRule() {}

// ImportRule represents a CSS @import rule.
type ImportRule struct {
	Href       string
	Conditions []string
}

// NewImportRule creates a new @import rule.
func NewImportRule(href string) *ImportRule {
	return &ImportRule{Href: href}
}

// Condition appends an import condition such as media, layer, or supports.
func (i *ImportRule) Condition(condition string) *ImportRule {
	i.Conditions = append(i.Conditions, condition)
	return i
}

// ConditionsList replaces import conditions.
func (i *ImportRule) ConditionsList(conditions []string) *ImportRule {
	i.Conditions = cloneStrings(conditions)
	return i
}

// Render returns freshly prepared ImportRule HTML bytes.
func (i *ImportRule) Render() []byte {
	return renderElement(i)
}

// HTML returns freshly prepared ImportRule HTML as a string.
func (i *ImportRule) HTML() string {
	return htmlElement(i)
}

// renderTo writes the ImportRule component's HTML to buf.
func (i *ImportRule) renderTo(buf *bytes.Buffer) {
	if i.Href == "" {
		return
	}
	buf.WriteString("\n@import ")
	buf.WriteString(formatImportHref(i.Href))
	if len(i.Conditions) != 0 {
		buf.WriteByte(' ')
		buf.WriteString(strings.Join(i.Conditions, " "))
	}
	buf.WriteString(";\n")
}

// IsCSSRule marks ImportRule as CSS content for style elements.
func (i *ImportRule) IsCSSRule() {}

// FontFaceRule represents a CSS @font-face declaration block.
type FontFaceRule struct {
	Props StyleMap
}

// NewFontFaceRule creates a new @font-face rule.
func NewFontFaceRule() *FontFaceRule {
	return &FontFaceRule{Props: StyleMap{}}
}

// Render returns freshly prepared FontFaceRule HTML bytes.
func (f *FontFaceRule) Render() []byte {
	return renderElement(f)
}

// HTML returns freshly prepared FontFaceRule HTML as a string.
func (f *FontFaceRule) HTML() string {
	return htmlElement(f)
}

// renderTo writes the FontFaceRule component's HTML to buf.
func (f *FontFaceRule) renderTo(buf *bytes.Buffer) {
	buf.WriteString(formatDeclarationBlock("@font-face", "", f.Props))
}

// AddStyle adds one font descriptor to the rule.
func (f *FontFaceRule) AddStyle(k, v string) *FontFaceRule {
	if f.Props == nil {
		f.Props = StyleMap{}
	}
	f.Props[k] = v
	return f
}

// AddStyles adds font descriptors to the rule.
func (f *FontFaceRule) AddStyles(styles StyleMap) *FontFaceRule {
	if f.Props == nil {
		f.Props = StyleMap{}
	}
	for k, v := range styles {
		f.Props[k] = v
	}
	return f
}

// Style replaces font descriptors on the rule.
func (f *FontFaceRule) Style(styles StyleMap) *FontFaceRule {
	f.Props = cloneStyleMap(styles)
	return f
}

// IsCSSRule marks FontFaceRule as CSS content for style elements.
func (f *FontFaceRule) IsCSSRule() {}

// FontFeatureValuesRule represents a CSS @font-feature-values block.
type FontFeatureValuesRule struct {
	Family  string
	Content string
}

// NewFontFeatureValuesRule creates a new @font-feature-values rule.
func NewFontFeatureValuesRule(family string) *FontFeatureValuesRule {
	return &FontFeatureValuesRule{Family: family}
}

// Text replaces the raw @font-feature-values block content.
func (f *FontFeatureValuesRule) Text(content string) *FontFeatureValuesRule {
	f.Content = content
	return f
}

// Render returns freshly prepared FontFeatureValuesRule HTML bytes.
func (f *FontFeatureValuesRule) Render() []byte {
	return renderElement(f)
}

// HTML returns freshly prepared FontFeatureValuesRule HTML as a string.
func (f *FontFeatureValuesRule) HTML() string {
	return htmlElement(f)
}

// renderTo writes the FontFeatureValuesRule component's HTML to buf.
func (f *FontFeatureValuesRule) renderTo(buf *bytes.Buffer) {
	buf.WriteString("\n@font-feature-values")
	if f.Family != "" {
		buf.WriteByte(' ')
		buf.WriteString(escapeCSS(f.Family))
	}
	buf.WriteString(" {\n")
	writeIndentedCSS(buf, tab, escapeCSS(f.Content))
	buf.WriteString("}\n")
}

// IsCSSRule marks FontFeatureValuesRule as CSS content for style elements.
func (f *FontFeatureValuesRule) IsCSSRule() {}

// MediaRule represents a CSS @media grouping rule.
type MediaRule struct {
	Query string
	Rules []CSSRule
}

// NewMediaRule creates a new @media rule.
func NewMediaRule(query string) *MediaRule {
	return &MediaRule{Query: query}
}

// AddRule appends a CSS rule inside the media block.
func (m *MediaRule) AddRule(rule CSSRule) *MediaRule {
	if !isNilElement(rule) {
		m.Rules = append(m.Rules, rule)
	}
	return m
}

// Add appends CSS content inside the media block.
func (m *MediaRule) Add(e Element) *MediaRule {
	switch element := e.(type) {
	case nil:
		return m
	case CSSRule:
		return m.AddRule(element)
	case *Style:
		if element == nil {
			return m
		}
		return m.AddRule(&StyleRule{
			Props: cloneStyleMap(element.Props),
			Tags:  append([]string(nil), element.Tags...),
		})
	default:
		return m
	}
}

// Render returns freshly prepared MediaRule HTML bytes.
func (m *MediaRule) Render() []byte {
	return renderElement(m)
}

// HTML returns freshly prepared MediaRule HTML as a string.
func (m *MediaRule) HTML() string {
	return htmlElement(m)
}

// renderTo writes the MediaRule component's HTML to buf.
func (m *MediaRule) renderTo(buf *bytes.Buffer) {
	buf.WriteString(formatGroupingAtRule("@media", m.Query, m.Rules))
}

// IsCSSRule marks MediaRule as CSS content for style elements.
func (m *MediaRule) IsCSSRule() {}

// KeyframeBlock represents one keyframe selector block inside @keyframes.
type KeyframeBlock struct {
	Selector string
	Props    StyleMap
}

// NewKeyframeBlock creates a new keyframe selector block.
func NewKeyframeBlock(selector string) *KeyframeBlock {
	return &KeyframeBlock{
		Props:    StyleMap{},
		Selector: selector,
	}
}

// Render returns freshly prepared KeyframeBlock HTML bytes.
func (k *KeyframeBlock) Render() []byte {
	return renderElement(k)
}

// HTML returns freshly prepared KeyframeBlock HTML as a string.
func (k *KeyframeBlock) HTML() string {
	return htmlElement(k)
}

// renderTo writes the KeyframeBlock component's HTML to buf.
func (k *KeyframeBlock) renderTo(buf *bytes.Buffer) {
	buf.WriteString(formatStyleRule([]string{k.Selector}, k.Props))
}

// AddStyle adds one CSS declaration to the keyframe block.
func (k *KeyframeBlock) AddStyle(prop, value string) *KeyframeBlock {
	if k.Props == nil {
		k.Props = StyleMap{}
	}
	k.Props[prop] = value
	return k
}

// AddStyles adds CSS declarations to the keyframe block.
func (k *KeyframeBlock) AddStyles(styles StyleMap) *KeyframeBlock {
	if k.Props == nil {
		k.Props = StyleMap{}
	}
	for prop, value := range styles {
		k.Props[prop] = value
	}
	return k
}

// Style replaces CSS declarations on the keyframe block.
func (k *KeyframeBlock) Style(styles StyleMap) *KeyframeBlock {
	k.Props = cloneStyleMap(styles)
	return k
}

// KeyframesRule represents a CSS @keyframes rule.
type KeyframesRule struct {
	Name   string
	Frames []*KeyframeBlock
}

// NewKeyframesRule creates a new @keyframes rule.
func NewKeyframesRule(name string) *KeyframesRule {
	return &KeyframesRule{Name: name}
}

// AddFrame appends one keyframe selector block.
func (k *KeyframesRule) AddFrame(selector string, props StyleMap) *KeyframesRule {
	frame := NewKeyframeBlock(selector).Style(props)
	return k.AddBlock(frame)
}

// AddBlock appends one keyframe block.
func (k *KeyframesRule) AddBlock(frame *KeyframeBlock) *KeyframesRule {
	if frame != nil {
		k.Frames = append(k.Frames, frame)
	}
	return k
}

// Render returns freshly prepared KeyframesRule HTML bytes.
func (k *KeyframesRule) Render() []byte {
	return renderElement(k)
}

// HTML returns freshly prepared KeyframesRule HTML as a string.
func (k *KeyframesRule) HTML() string {
	return htmlElement(k)
}

// renderTo writes the KeyframesRule component's HTML to buf.
func (k *KeyframesRule) renderTo(buf *bytes.Buffer) {
	buf.WriteString("\n@keyframes ")
	buf.WriteString(escapeCSS(k.Name))
	buf.WriteString(" {\n")
	for _, frame := range k.Frames {
		if frame == nil {
			continue
		}
		writeIndentedCSS(buf, tab, frame.HTML())
	}
	buf.WriteString("}\n")
}

// IsCSSRule marks KeyframesRule as CSS content for style elements.
func (k *KeyframesRule) IsCSSRule() {}

// formatStyleRule renders one CSS rule from selectors and declarations.
func formatStyleRule(tags []string, props StyleMap) string {
	var buf bytes.Buffer
	buf.WriteByte('\n')
	writeCSSStyleRule(&buf, "", tags, props)
	return buf.String()
}

// formatDeclarationBlock renders an at-rule with a declaration block.
func formatDeclarationBlock(name, prelude string, props StyleMap) string {
	var buf bytes.Buffer
	buf.WriteByte('\n')
	buf.WriteString(name)
	if prelude != "" {
		buf.WriteByte(' ')
		buf.WriteString(escapeCSS(prelude))
	}
	buf.WriteString(" {\n")
	writeCSSDeclarations(&buf, tab, props)
	buf.WriteString("}\n")
	return buf.String()
}

// formatGroupingAtRule renders an at-rule containing nested CSS rules.
func formatGroupingAtRule(name, prelude string, rules []CSSRule) string {
	var buf bytes.Buffer
	buf.WriteByte('\n')
	buf.WriteString(name)
	if prelude != "" {
		buf.WriteByte(' ')
		buf.WriteString(escapeCSS(prelude))
	}
	buf.WriteString(" {\n")
	for _, rule := range rules {
		if isNilElement(rule) {
			continue
		}
		writeIndentedCSS(&buf, tab, rule.HTML())
	}
	buf.WriteString("}\n")
	return buf.String()
}

// writeCSSStyleRule writes a CSS selector rule at the given indentation.
func writeCSSStyleRule(buf *bytes.Buffer, indent string, tags []string, props StyleMap) {
	buf.WriteString(indent)
	buf.WriteString(escapeCSS(formatStringArray(tags)))
	buf.WriteString(" {\n")
	writeCSSDeclarations(buf, indent+tab, props)
	buf.WriteString(indent)
	buf.WriteString("}\n")
}

// writeCSSDeclarations writes non-empty CSS declarations in stable key order.
func writeCSSDeclarations(buf *bytes.Buffer, indent string, props StyleMap) {
	for _, prop := range sortedStyleKeys(props) {
		value := props[prop]
		if prop == "" || value == "" {
			continue
		}
		buf.WriteString(indent)
		buf.WriteString(escapeCSS(prop))
		buf.WriteString(": ")
		buf.WriteString(escapeCSS(value))
		buf.WriteString(";\n")
	}
}

// sortedStyleKeys returns StyleMap keys in stable order.
func sortedStyleKeys(styles StyleMap) []string {
	keys := make([]string, 0, len(styles))
	for k := range styles {
		keys = append(keys, k)
	}
	slices.Sort(keys)
	return keys
}

// formatStringArray joins CSS selectors for a rule block.
func formatStringArray(sarr []string) string {
	return strings.Join(sarr, ", ")
}

// writeIndentedCSS writes CSS text with one indentation prefix per non-empty line.
func writeIndentedCSS(buf *bytes.Buffer, indent, css string) {
	css = strings.Trim(css, "\n")
	if css == "" {
		return
	}
	for _, line := range strings.Split(css, "\n") {
		if strings.TrimSpace(line) == "" {
			buf.WriteByte('\n')
			continue
		}
		buf.WriteString(indent)
		buf.WriteString(line)
		buf.WriteByte('\n')
	}
}

// quoteCSSString returns a double-quoted CSS string.
func quoteCSSString(value string) string {
	return escapeCSS(strconv.Quote(value))
}

// formatImportHref formats an import href as a CSS url() unless already CSS-formatted.
func formatImportHref(href string) string {
	trimmed := strings.TrimSpace(href)
	if strings.HasPrefix(trimmed, "url(") || strings.HasPrefix(trimmed, "\"") || strings.HasPrefix(trimmed, "'") {
		return escapeCSS(trimmed)
	}
	return "url(" + quoteCSSString(href) + ")"
}

// IsHeadElement implements the marker interface.
func (s *Style) IsHeadElement() {}
