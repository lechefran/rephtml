package rephtml

import (
	"bytes"
	"reflect"
	"strconv"
	"strings"
)

// CssProps holds CSS declarations for normal selector blocks.
//
// Each non-empty field renders as a CSS property based on PropMap. At-rules such
// as @media and @font-face have their own rule types because they use different
// CSS grammar than ordinary declarations.
type CssProps struct {
	AccentColor              string
	AlignContent             string
	AlignItems               string
	AlignSelf                string
	All                      string
	Animation                string
	AnimationDelay           string
	AnimationDirection       string
	AnimationDuration        string
	AnimationFillMode        string
	AnimationIterationCount  string
	AnimationName            string
	AnimationPlayState       string
	AnimationTimingFunction  string
	AspectRatio              string
	BackdropFilter           string
	BackfaceVisibility       string
	Background               string
	BackgroundAttachment     string
	BackgroundBlendMode      string
	BackgroundClip           string
	BackgroundColor          string
	BackgroundImage          string
	BackgroundOrigin         string
	BackgroundPosition       string
	BackgroundPositionX      string
	BackgroundPositionY      string
	BackgroundRepeat         string
	BackgroundSize           string
	BlockSize                string
	Border                   string
	BorderBlock              string
	BorderBlockColor         string
	BorderBlockEnd           string
	BorderBlockEndColor      string
	BorderBlockEndStyle      string
	BorderBlockEndWidth      string
	BorderBlockStart         string
	BorderBlockStartColor    string
	BorderBlockStartStyle    string
	BorderBlockStartWidth    string
	BorderBlockStyle         string
	BorderBlockWidth         string
	BorderBottom             string
	BorderBottomColor        string
	BorderBottomLeftRadius   string
	BorderBottomRightRadius  string
	BorderBottomStyle        string
	BorderBottomWidth        string
	BorderCollapse           string
	BorderColor              string
	BorderEndEndRadius       string
	BorderEndStartRadius     string
	BorderImage              string
	BorderImageOutset        string
	BorderImageRepeat        string
	BorderImageSlice         string
	BorderImageSource        string
	BorderImageWidth         string
	BorderInline             string
	BorderInlineColor        string
	BorderInlineEnd          string
	BorderInlineEndColor     string
	BorderInlineEndStyle     string
	BorderInlineEndWidth     string
	BorderInlineStart        string
	BorderInlineStartColor   string
	BorderInlineStartStyle   string
	BorderInlineStartWidth   string
	BorderInlineStyle        string
	BorderInlineWidth        string
	BorderLeft               string
	BorderLeftColor          string
	BorderLeftStyle          string
	BorderLeftWidth          string
	BorderRadius             string
	BorderRight              string
	BorderRightColor         string
	BorderRightStyle         string
	BorderRightWidth         string
	BorderSpacing            string
	BorderStartEndRadius     string
	BorderStartStartRadius   string
	BorderStyle              string
	BorderTop                string
	BorderTopColor           string
	BorderTopLeftRadius      string
	BorderTopRightRadius     string
	BorderTopStyle           string
	BorderTopWidth           string
	BorderWidth              string
	Bottom                   string
	BoxDecorationBreak       string
	BoxReflect               string
	BoxShadow                string
	BoxSizing                string
	BreakAfter               string
	BreakBefore              string
	BreakInside              string
	CaptionSide              string
	CaretColor               string
	Clear                    string
	Clip                     string
	ClipPath                 string
	Color                    string
	ColorScheme              string
	ColumnCount              string
	ColumnFill               string
	ColumnGap                string
	ColumnRule               string
	ColumnRuleColor          string
	ColumnRuleStyle          string
	ColumnRuleWidth          string
	ColumnSpan               string
	ColumnWidth              string
	Columns                  string
	Content                  string
	CounterIncrement         string
	CounterReset             string
	CounterSet               string
	Cursor                   string
	Direction                string
	Display                  string
	EmptyCells               string
	Filter                   string
	Flex                     string
	FlexBasis                string
	FlexDirection            string
	FlexFlow                 string
	FlexGrow                 string
	FlexShrink               string
	FlexWrap                 string
	Float                    string
	Font                     string
	FontFamily               string
	FontFeatureSettings      string
	FontKerning              string
	FontLanguageOverride     string
	FontSize                 string
	FontSizeAdjust           string
	FontStretch              string
	FontStyle                string
	FontSynthesis            string
	FontVariant              string
	FontVariantAlternates    string
	FontVariantCaps          string
	FontVariantEastAsian     string
	FontVariantLigatures     string
	FontVariantNumeric       string
	FontVariantPosition      string
	FontWeight               string
	Gap                      string
	Grid                     string
	GridArea                 string
	GridAutoColumns          string
	GridAutoFlow             string
	GridAutoRows             string
	GridColumn               string
	GridColumnEnd            string
	GridColumnStart          string
	GridRow                  string
	GridRowEnd               string
	GridRowStart             string
	GridTemplate             string
	GridTemplateAreas        string
	GridTemplateColumns      string
	GridTemplateRows         string
	HangingPunctuation       string
	Height                   string
	Hyphens                  string
	HypenateCharacter        string
	ImageRendering           string
	InitialLetter            string
	InlineSize               string
	Inset                    string
	InsetBlock               string
	InsetBlockEnd            string
	InsetBlockStart          string
	InsetInline              string
	InsetInlineEnd           string
	InsetInlineStart         string
	Isolation                string
	JustifyContent           string
	JustifyItems             string
	JustifySelf              string
	Left                     string
	LetterSpacing            string
	LineBreak                string
	LineHeight               string
	ListStyle                string
	ListStyleImage           string
	ListStylePosition        string
	ListStyleType            string
	Margin                   string
	MarginBlock              string
	MarginBlockEnd           string
	MarginBlockStart         string
	MarginBottom             string
	MarginInline             string
	MarginInlineEnd          string
	MarginInlineStart        string
	MarginLeft               string
	MarginRight              string
	MarginTop                string
	Marker                   string
	MarkerEnd                string
	MarkerMid                string
	MarkerStart              string
	Mask                     string
	MaskClip                 string
	MaskComposite            string
	MaskImage                string
	MaskMode                 string
	MaskOrigin               string
	MaskPosition             string
	MaskRepeat               string
	MaskSize                 string
	MaskType                 string
	MaxHeight                string
	MaxWidth                 string
	MaxBlockSize             string
	MaxInlineSize            string
	MinBlockSize             string
	MinInlineSize            string
	MinHeight                string
	MinWidth                 string
	MixBlendMode             string
	ObjectFit                string
	ObjectPosition           string
	Offset                   string
	OffsetAnchor             string
	OffsetDistance           string
	OffsetPath               string
	OffsetPosition           string
	OffsetRotate             string
	Opacity                  string
	Order                    string
	Orphans                  string
	Outline                  string
	OutlineColor             string
	OutlineOffset            string
	OutlineStyle             string
	OutlineWidth             string
	Overflow                 string
	OverflowAnchor           string
	OverflowWrap             string
	OverflowX                string
	OverflowY                string
	OverscrollBehavior       string
	OverscrollBehaviorBlock  string
	OverscrollBehaviorInline string
	OverscrollBehaviorX      string
	OverscrollBehaviorY      string
	Padding                  string
	PaddingBlock             string
	PaddingBlockEnd          string
	PaddingBlockStart        string
	PaddingBottom            string
	PaddingInline            string
	PaddingInlineEnd         string
	PaddingInlineStart       string
	PaddingLeft              string
	PaddingRight             string
	PaddingTop               string
	PageBreakAfter           string
	PageBreakBefore          string
	PageBreakInside          string
	PaintOrder               string
	Perspective              string
	PerspectiveOrigin        string
	PlaceContent             string
	PlaceItems               string
	PlaceSelf                string
	PointerEvents            string
	Position                 string
	Quotes                   string
	Resize                   string
	Right                    string
	Rotate                   string
	RowGap                   string
	Scale                    string
	ScrollBehavior           string
	ScrollMargin             string
	ScrollMarginBlock        string
	ScrollMarginBlockEnd     string
	ScrollMarginBlockStart   string
	ScrollMarginBottom       string
	ScrollMarginInline       string
	ScrollMarginInlineEnd    string
	ScrollMarginInlineStart  string
	ScrollMarginLeft         string
	ScrollMarginRight        string
	ScrollMarginTop          string
	ScrollPadding            string
	ScrollPaddingBlock       string
	ScrollPaddingBlockEnd    string
	ScrollPaddingBlockStart  string
	ScrollPaddingBottom      string
	ScrollPaddingInline      string
	ScrollPaddingInlineEnd   string
	ScrollPaddingInlineStart string
	ScrollPaddingLeft        string
	ScrollPaddingRight       string
	ScrollPaddingTop         string
	ScrollSnapAlign          string
	ScrollSnapStop           string
	ScrollSnapType           string
	ScrollbarColor           string
	TabSize                  string
	TableLayout              string
	TextAlign                string
	TextAlignLast            string
	TextCombineUpright       string
	TextDecoration           string
	TextDecorationColor      string
	TextDecorationLine       string
	TextDecorationStyle      string
	TextDecorationThickness  string
	TextEmphasis             string
	TextEmphasisColor        string
	TextEmphasisPosition     string
	TextEmphasisStyle        string
	TextIndent               string
	TextJustify              string
	TextOrientation          string
	TextOverflow             string
	TextShadow               string
	TextTransform            string
	TextUnderlineOffset      string
	TextUnderlinePosition    string
	Top                      string
	Transform                string
	TransformOrigin          string
	TransformStyle           string
	Transition               string
	TransitionDelay          string
	TransitionDuration       string
	TransitionProperty       string
	TransitionTimingFunction string
	Translate                string
	UnicodeBidi              string
	UserSelect               string
	VerticalAlign            string
	Visibility               string
	WhiteSpace               string
	Widows                   string
	Width                    string
	WordBreak                string
	WordSpacing              string
	WordWrap                 string
	WritingMode              string
	ZIndex                   string
}

// FontFaceProps represents CSS descriptors accepted inside @font-face.
type FontFaceProps struct {
	FontDisplay           string
	FontFamily            string
	FontFeatureSettings   string
	FontStretch           string
	FontStyle             string
	FontVariationSettings string
	FontWeight            string
	Src                   string
	UnicodeRange          string
}

// NewCssProps returns an empty CssProps value.
func NewCssProps() *CssProps {
	return &CssProps{}
}

// Style renders a complete style element for one selector rule.
type Style struct {
	buf   bytes.Buffer
	pmap  *PropMap
	Props CssProps
	Tags  []string
}

// NewStyle creates a style element for one CSS selector rule.
func NewStyle(tags ...string) *Style {
	return &Style{
		pmap: NewPropMap(),
		Tags: append([]string(nil), tags...),
	}
}

// Bytes returns a defensive copy of the rendered Style bytes.
func (s *Style) Bytes() []byte {
	return cloneBytes(s.buf.Bytes())
}

// Prepare renders the Style component into its internal buffer.
func (s *Style) Prepare() {
	s.buf.Reset()
	s.buf.WriteString("<style>")
	s.buf.WriteString(formatStyleRule(s.Tags, s.Props, s.pmap))
	s.buf.WriteString("</style>")
}

// IsHeadElement implements HeadElement interface
func (s *Style) IsHeadElement() {}

// formatStringArray joins CSS selectors for a rule block.
func formatStringArray(sarr []string) string {
	res := ""
	for i := 0; i < len(sarr); i++ {
		if i != len(sarr)-1 {
			res += sarr[i] + ", "
		} else {
			res += sarr[i]
		}
	}
	return res
}

// PropMap replaces the field-to-property mapping used by the style rule.
func (s *Style) PropMap(p *PropMap) *Style {
	s.pmap = p
	return s
}

// CSSRule represents renderable CSS that can be placed inside a style element.
type CSSRule interface {
	Element
	IsCSSRule()
}

// StyleRule renders one CSS selector rule without wrapping it in a style tag.
type StyleRule struct {
	buf   bytes.Buffer
	pmap  *PropMap
	Props CssProps
	Tags  []string
}

// NewStyleRule creates one unwrapped selector rule.
func NewStyleRule(tags ...string) *StyleRule {
	return &StyleRule{
		pmap: NewPropMap(),
		Tags: append([]string(nil), tags...),
	}
}

// Bytes returns a defensive copy of the rendered StyleRule bytes.
func (s *StyleRule) Bytes() []byte {
	return cloneBytes(s.buf.Bytes())
}

// Prepare renders the StyleRule component into its internal buffer.
func (s *StyleRule) Prepare() {
	s.buf.Reset()
	s.buf.WriteString(formatStyleRule(s.Tags, s.Props, s.pmap))
}

// PropMap replaces the field-to-property mapping used by the rule.
func (s *StyleRule) PropMap(p *PropMap) *StyleRule {
	s.pmap = p
	return s
}

// IsCSSRule marks StyleRule as CSS content for style elements.
func (s *StyleRule) IsCSSRule() {}

// RawCSSRule represents raw CSS content for custom or unsupported rules.
type RawCSSRule struct {
	buf bytes.Buffer
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

// Bytes returns a defensive copy of the rendered RawCSSRule bytes.
func (r *RawCSSRule) Bytes() []byte {
	return cloneBytes(r.buf.Bytes())
}

// Prepare renders the RawCSSRule component into its internal buffer.
func (r *RawCSSRule) Prepare() {
	r.buf.Reset()
	css := strings.Trim(r.css, "\n")
	if css == "" {
		return
	}
	r.buf.WriteByte('\n')
	r.buf.WriteString(css)
	r.buf.WriteByte('\n')
}

// IsCSSRule marks RawCSSRule as CSS content for style elements.
func (r *RawCSSRule) IsCSSRule() {}

// CharsetRule represents a CSS @charset rule.
type CharsetRule struct {
	buf     bytes.Buffer
	Charset string
}

// NewCharsetRule creates a new @charset rule.
func NewCharsetRule(charset string) *CharsetRule {
	return &CharsetRule{Charset: charset}
}

// Bytes returns a defensive copy of the rendered CharsetRule bytes.
func (c *CharsetRule) Bytes() []byte {
	return cloneBytes(c.buf.Bytes())
}

// Prepare renders the CharsetRule component into its internal buffer.
func (c *CharsetRule) Prepare() {
	c.buf.Reset()
	if c.Charset == "" {
		return
	}
	c.buf.WriteString("\n@charset ")
	c.buf.WriteString(quoteCSSString(c.Charset))
	c.buf.WriteString(";\n")
}

// IsCSSRule marks CharsetRule as CSS content for style elements.
func (c *CharsetRule) IsCSSRule() {}

// ImportRule represents a CSS @import rule.
type ImportRule struct {
	buf        bytes.Buffer
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

// Bytes returns a defensive copy of the rendered ImportRule bytes.
func (i *ImportRule) Bytes() []byte {
	return cloneBytes(i.buf.Bytes())
}

// Prepare renders the ImportRule component into its internal buffer.
func (i *ImportRule) Prepare() {
	i.buf.Reset()
	if i.Href == "" {
		return
	}
	i.buf.WriteString("\n@import ")
	i.buf.WriteString(formatImportHref(i.Href))
	if len(i.Conditions) != 0 {
		i.buf.WriteByte(' ')
		i.buf.WriteString(strings.Join(i.Conditions, " "))
	}
	i.buf.WriteString(";\n")
}

// IsCSSRule marks ImportRule as CSS content for style elements.
func (i *ImportRule) IsCSSRule() {}

// FontFaceRule represents a CSS @font-face declaration block.
type FontFaceRule struct {
	buf   bytes.Buffer
	Props FontFaceProps
}

// NewFontFaceRule creates a new @font-face rule.
func NewFontFaceRule() *FontFaceRule {
	return &FontFaceRule{}
}

// Bytes returns a defensive copy of the rendered FontFaceRule bytes.
func (f *FontFaceRule) Bytes() []byte {
	return cloneBytes(f.buf.Bytes())
}

// Prepare renders the FontFaceRule component into its internal buffer.
func (f *FontFaceRule) Prepare() {
	f.buf.Reset()
	f.buf.WriteString(formatFontFaceRule(f.Props))
}

// IsCSSRule marks FontFaceRule as CSS content for style elements.
func (f *FontFaceRule) IsCSSRule() {}

// FontFeatureValuesRule represents a CSS @font-feature-values block.
type FontFeatureValuesRule struct {
	buf     bytes.Buffer
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

// Bytes returns a defensive copy of the rendered FontFeatureValuesRule bytes.
func (f *FontFeatureValuesRule) Bytes() []byte {
	return cloneBytes(f.buf.Bytes())
}

// Prepare renders the FontFeatureValuesRule component into its internal buffer.
func (f *FontFeatureValuesRule) Prepare() {
	f.buf.Reset()
	f.buf.WriteString("\n@font-feature-values")
	if f.Family != "" {
		f.buf.WriteByte(' ')
		f.buf.WriteString(f.Family)
	}
	f.buf.WriteString(" {\n")
	writeIndentedCSS(&f.buf, tab, f.Content)
	f.buf.WriteString("}\n")
}

// IsCSSRule marks FontFeatureValuesRule as CSS content for style elements.
func (f *FontFeatureValuesRule) IsCSSRule() {}

// MediaRule represents a CSS @media grouping rule.
type MediaRule struct {
	buf   bytes.Buffer
	Query string
	Rules []CSSRule
}

// NewMediaRule creates a new @media rule.
func NewMediaRule(query string) *MediaRule {
	return &MediaRule{Query: query}
}

// AddRule appends a CSS rule inside the media block.
func (m *MediaRule) AddRule(rule CSSRule) *MediaRule {
	if !isNilCSSRule(rule) {
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
			pmap:  element.pmap,
			Props: element.Props,
			Tags:  append([]string(nil), element.Tags...),
		})
	default:
		return m
	}
}

// Bytes returns a defensive copy of the rendered MediaRule bytes.
func (m *MediaRule) Bytes() []byte {
	return cloneBytes(m.buf.Bytes())
}

// Prepare renders the MediaRule component into its internal buffer.
func (m *MediaRule) Prepare() {
	m.buf.Reset()
	m.buf.WriteString(formatGroupingAtRule("@media", m.Query, m.Rules))
}

// IsCSSRule marks MediaRule as CSS content for style elements.
func (m *MediaRule) IsCSSRule() {}

// KeyframeBlock represents one keyframe selector block inside @keyframes.
type KeyframeBlock struct {
	buf      bytes.Buffer
	pmap     *PropMap
	Selector string
	Props    CssProps
}

// NewKeyframeBlock creates a new keyframe selector block.
func NewKeyframeBlock(selector string) *KeyframeBlock {
	return &KeyframeBlock{
		pmap:     NewPropMap(),
		Selector: selector,
	}
}

// PropMap sets the property map for the KeyframeBlock component.
func (k *KeyframeBlock) PropMap(p *PropMap) *KeyframeBlock {
	k.pmap = p
	return k
}

// Bytes returns a defensive copy of the rendered KeyframeBlock bytes.
func (k *KeyframeBlock) Bytes() []byte {
	return cloneBytes(k.buf.Bytes())
}

// Prepare renders the KeyframeBlock component into its internal buffer.
func (k *KeyframeBlock) Prepare() {
	k.buf.Reset()
	k.buf.WriteString(formatStyleRule([]string{k.Selector}, k.Props, k.pmap))
}

// KeyframesRule represents a CSS @keyframes rule.
type KeyframesRule struct {
	buf    bytes.Buffer
	Name   string
	Frames []*KeyframeBlock
}

// NewKeyframesRule creates a new @keyframes rule.
func NewKeyframesRule(name string) *KeyframesRule {
	return &KeyframesRule{Name: name}
}

// AddFrame appends one keyframe selector block.
func (k *KeyframesRule) AddFrame(selector string, props CssProps) *KeyframesRule {
	frame := NewKeyframeBlock(selector)
	frame.Props = props
	return k.AddBlock(frame)
}

// AddBlock appends one keyframe block.
func (k *KeyframesRule) AddBlock(frame *KeyframeBlock) *KeyframesRule {
	if frame != nil {
		k.Frames = append(k.Frames, frame)
	}
	return k
}

// Bytes returns a defensive copy of the rendered KeyframesRule bytes.
func (k *KeyframesRule) Bytes() []byte {
	return cloneBytes(k.buf.Bytes())
}

// Prepare renders the KeyframesRule component into its internal buffer.
func (k *KeyframesRule) Prepare() {
	k.buf.Reset()
	k.buf.WriteString("\n@keyframes ")
	k.buf.WriteString(k.Name)
	k.buf.WriteString(" {\n")
	for _, frame := range k.Frames {
		if frame == nil {
			continue
		}
		frame.Prepare()
		writeIndentedCSS(&k.buf, tab, string(frame.Bytes()))
	}
	k.buf.WriteString("}\n")
}

// IsCSSRule marks KeyframesRule as CSS content for style elements.
func (k *KeyframesRule) IsCSSRule() {}

// formatStyleRule renders one CSS rule from selectors and typed properties.
func formatStyleRule(tags []string, props CssProps, pmap *PropMap) string {
	if pmap == nil {
		pmap = NewPropMap()
	}

	var buf bytes.Buffer
	buf.WriteByte('\n')
	writeCSSStyleRule(&buf, "", tags, props, pmap)
	return buf.String()
}

// formatFontFaceRule renders a @font-face descriptor block.
func formatFontFaceRule(props FontFaceProps) string {
	var buf bytes.Buffer
	buf.WriteByte('\n')
	buf.WriteString("@font-face {\n")
	writeMappedDeclarations(&buf, tab, props, fontFacePropMap())
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
		buf.WriteString(prelude)
	}
	buf.WriteString(" {\n")
	for _, rule := range rules {
		if isNilCSSRule(rule) {
			continue
		}
		rule.Prepare()
		writeIndentedCSS(&buf, tab, string(rule.Bytes()))
	}
	buf.WriteString("}\n")
	return buf.String()
}

// writeCSSStyleRule writes a CSS selector rule at the given indentation.
func writeCSSStyleRule(buf *bytes.Buffer, indent string, tags []string, props CssProps, pmap *PropMap) {
	buf.WriteString(indent)
	buf.WriteString(formatStringArray(tags))
	buf.WriteString(" {\n")
	writeCSSDeclarations(buf, indent+tab, props, pmap)
	buf.WriteString(indent)
	buf.WriteString("}\n")
}

// writeCSSDeclarations writes non-empty typed CSS declarations.
func writeCSSDeclarations(buf *bytes.Buffer, indent string, props CssProps, pmap *PropMap) {
	if pmap == nil {
		pmap = NewPropMap()
	}
	writeMappedDeclarations(buf, indent, props, pmap.pmap)
}

// writeMappedDeclarations writes non-empty string fields using a field-to-property map.
func writeMappedDeclarations(buf *bytes.Buffer, indent string, props any, propMap map[string]string) {
	val := reflect.ValueOf(props)
	t := val.Type()
	for i := 0; i < val.NumField(); i++ {
		k, v := t.Field(i).Name, val.Field(i)
		if k == "" || v.String() == "" {
			continue
		}
		prop, ok := propMap[k]
		if !ok {
			continue
		}
		buf.WriteString(indent)
		buf.WriteString(prop)
		buf.WriteString(": ")
		buf.WriteString(v.String())
		buf.WriteString(";\n")
	}
}

// fontFacePropMap maps FontFaceProps field names to CSS font descriptors.
func fontFacePropMap() map[string]string {
	return map[string]string{
		"FontDisplay":           "font-display",
		"FontFamily":            "font-family",
		"FontFeatureSettings":   "font-feature-settings",
		"FontStretch":           "font-stretch",
		"FontStyle":             "font-style",
		"FontVariationSettings": "font-variation-settings",
		"FontWeight":            "font-weight",
		"Src":                   "src",
		"UnicodeRange":          "unicode-range",
	}
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
	return strconv.Quote(value)
}

// formatImportHref formats an import href as a CSS url() unless already CSS-formatted.
func formatImportHref(href string) string {
	trimmed := strings.TrimSpace(href)
	if strings.HasPrefix(trimmed, "url(") || strings.HasPrefix(trimmed, "\"") || strings.HasPrefix(trimmed, "'") {
		return trimmed
	}
	return "url(" + quoteCSSString(href) + ")"
}

// isNilCSSRule reports whether a CSSRule is nil or a nil pointer.
func isNilCSSRule(rule CSSRule) bool {
	if rule == nil {
		return true
	}
	value := reflect.ValueOf(rule)
	switch value.Kind() {
	case reflect.Chan, reflect.Func, reflect.Interface, reflect.Map, reflect.Pointer, reflect.Slice:
		return value.IsNil()
	default:
		return false
	}
}

// PropMap maps CssProps field names to emitted CSS property names.
type PropMap struct {
	pmap map[string]string
}

// NewPropMap creates a new PropMap component.
func NewPropMap() *PropMap {
	pmap := map[string]string{}
	pmap["AccentColor"] = "accent-color"
	pmap["AlignContent"] = "align-content"
	pmap["AlignItems"] = "align-items"
	pmap["AlignSelf"] = "align-self"
	pmap["All"] = "all"
	pmap["Animation"] = "animation"
	pmap["AnimationDelay"] = "animation-delay"
	pmap["AnimationDirection"] = "animation-direction"
	pmap["AnimationDuration"] = "animation-duration"
	pmap["AnimationFillMode"] = "animation-fill-mode"
	pmap["AnimationIterationCount"] = "animation-iteration-count"
	pmap["AnimationName"] = "animation-name"
	pmap["AnimationPlayState"] = "animation-play-state"
	pmap["AnimationTimingFunction"] = "animation-timing-function"
	pmap["AspectRatio"] = "aspect-ratio"
	pmap["BackdropFilter"] = "backdrop-filter"
	pmap["BackfaceVisibility"] = "backface-visibility"
	pmap["Background"] = "background"
	pmap["BackgroundAttachment"] = "background-attachment"
	pmap["BackgroundBlendMode"] = "background-blend-mode"
	pmap["BackgroundClip"] = "background-clip"
	pmap["BackgroundColor"] = "background-color"
	pmap["BackgroundImage"] = "background-image"
	pmap["BackgroundOrigin"] = "background-origin"
	pmap["BackgroundPosition"] = "background-position"
	pmap["BackgroundPositionX"] = "background-position-x"
	pmap["BackgroundPositionY"] = "background-position-y"
	pmap["BackgroundRepeat"] = "background-repeat"
	pmap["BackgroundSize"] = "background-size"
	pmap["BlockSize"] = "block-size"
	pmap["Border"] = "border"
	pmap["BorderBlock"] = "border-block"
	pmap["BorderBlockColor"] = "border-block-color"
	pmap["BorderBlockEnd"] = "border-block-end"
	pmap["BorderBlockEndColor"] = "border-block-end-color"
	pmap["BorderBlockEndStyle"] = "border-block-end-style"
	pmap["BorderBlockEndWidth"] = "border-block-end-width"
	pmap["BorderBlockStart"] = "border-block-start"
	pmap["BorderBlockStartColor"] = "border-block-start-color"
	pmap["BorderBlockStartStyle"] = "border-block-start-style"
	pmap["BorderBlockStartWidth"] = "border-block-start-width"
	pmap["BorderBlockStyle"] = "border-block-style"
	pmap["BorderBlockWidth"] = "border-block-width"
	pmap["BorderBottom"] = "border-bottom"
	pmap["BorderBottomColor"] = "border-bottom-color"
	pmap["BorderBottomLeftRadius"] = "border-bottom-left-radius"
	pmap["BorderBottomRightRadius"] = "border-bottom-right-radius"
	pmap["BorderBottomStyle"] = "border-bottom-style"
	pmap["BorderBottomWidth"] = "border-bottom-width"
	pmap["BorderCollapse"] = "border-collapse"
	pmap["BorderColor"] = "border-color"
	pmap["BorderEndEndRadius"] = "border-end-end-radius"
	pmap["BorderEndStartRadius"] = "border-end-start-radius"
	pmap["BorderImage"] = "border-image"
	pmap["BorderImageOutset"] = "border-image-outset"
	pmap["BorderImageRepeat"] = "border-image-repeat"
	pmap["BorderImageSlice"] = "border-image-slice"
	pmap["BorderImageSource"] = "border-image-source"
	pmap["BorderImageWidth"] = "border-image-width"
	pmap["BorderInline"] = "border-inline"
	pmap["BorderInlineColor"] = "border-inline-color"
	pmap["BorderInlineEnd"] = "border-inline-end"
	pmap["BorderInlineEndColor"] = "border-inline-end-color"
	pmap["BorderInlineEndStyle"] = "border-inline-end-style"
	pmap["BorderInlineEndWidth"] = "border-inline-end-width"
	pmap["BorderInlineStart"] = "border-inline-start"
	pmap["BorderInlineStartColor"] = "border-inline-start-color"
	pmap["BorderInlineStartStyle"] = "border-inline-start-style"
	pmap["BorderInlineStartWidth"] = "border-inline-start-width"
	pmap["BorderInlineStyle"] = "border-inline-style"
	pmap["BorderInlineWidth"] = "border-inline-width"
	pmap["BorderLeft"] = "border-left"
	pmap["BorderLeftColor"] = "border-left-color"
	pmap["BorderLeftStyle"] = "border-left-style"
	pmap["BorderLeftWidth"] = "border-left-width"
	pmap["BorderRadius"] = "border-radius"
	pmap["BorderRight"] = "border-right"
	pmap["BorderRightColor"] = "border-right-color"
	pmap["BorderRightStyle"] = "border-right-style"
	pmap["BorderRightWidth"] = "border-right-width"
	pmap["BorderSpacing"] = "border-spacing"
	pmap["BorderStartEndRadius"] = "border-start-end-radius"
	pmap["BorderStartStartRadius"] = "border-start-start-radius"
	pmap["BorderStyle"] = "border-style"
	pmap["BorderTop"] = "border-top"
	pmap["BorderTopColor"] = "border-top-color"
	pmap["BorderTopLeftRadius"] = "border-top-left-radius"
	pmap["BorderTopRightRadius"] = "border-top-right-radius"
	pmap["BorderTopStyle"] = "border-top-style"
	pmap["BorderTopWidth"] = "border-top-width"
	pmap["BorderWidth"] = "border-width"
	pmap["Bottom"] = "bottom"
	pmap["BoxDecorationBreak"] = "box-decoration-break"
	pmap["BoxReflect"] = "box-reflect"
	pmap["BoxShadow"] = "box-shadow"
	pmap["BoxSizing"] = "box-sizing"
	pmap["BreakAfter"] = "break-after"
	pmap["BreakBefore"] = "break-before"
	pmap["BreakInside"] = "break-inside"
	pmap["CaptionSide"] = "caption-side"
	pmap["CaretColor"] = "caret-color"
	pmap["Clear"] = "clear"
	pmap["Clip"] = "clip"
	pmap["ClipPath"] = "clip-path"
	pmap["Color"] = "color"
	pmap["ColorScheme"] = "color-scheme"
	pmap["ColumnCount"] = "column-count"
	pmap["ColumnFill"] = "column-fill"
	pmap["ColumnGap"] = "column-gap"
	pmap["ColumnRule"] = "column-rule"
	pmap["ColumnRuleColor"] = "column-rule-color"
	pmap["ColumnRuleStyle"] = "column-rule-style"
	pmap["ColumnRuleWidth"] = "column-rule-width"
	pmap["ColumnSpan"] = "column-span"
	pmap["ColumnWidth"] = "column-width"
	pmap["Columns"] = "columns"
	pmap["Content"] = "content"
	pmap["CounterIncrement"] = "counter-increment"
	pmap["CounterReset"] = "counter-reset"
	pmap["CounterSet"] = "counter-set"
	pmap["Cursor"] = "cursor"
	pmap["Direction"] = "direction"
	pmap["Display"] = "display"
	pmap["EmptyCells"] = "empty-cells"
	pmap["Filter"] = "filter"
	pmap["Flex"] = "flex"
	pmap["FlexBasis"] = "flex-basis"
	pmap["FlexDirection"] = "flex-direction"
	pmap["FlexFlow"] = "flex-flow"
	pmap["FlexGrow"] = "flex-grow"
	pmap["FlexShrink"] = "flex-shrink"
	pmap["FlexWrap"] = "flex-wrap"
	pmap["Float"] = "float"
	pmap["Font"] = "font"
	pmap["FontFamily"] = "font-family"
	pmap["FontFeatureSettings"] = "font-feature-settings"
	pmap["FontKerning"] = "font-kerning"
	pmap["FontLanguageOverride"] = "font-language-override"
	pmap["FontSize"] = "font-size"
	pmap["FontSizeAdjust"] = "font-size-adjust"
	pmap["FontStretch"] = "font-stretch"
	pmap["FontStyle"] = "font-style"
	pmap["FontSynthesis"] = "font-synthesis"
	pmap["FontVariant"] = "font-variant"
	pmap["FontVariantAlternates"] = "font-variant-alternates"
	pmap["FontVariantCaps"] = "font-variant-caps"
	pmap["FontVariantEastAsian"] = "font-variant-east-asian"
	pmap["FontVariantLigatures"] = "font-variant-ligatures"
	pmap["FontVariantNumeric"] = "font-variant-numeric"
	pmap["FontVariantPosition"] = "font-variant-position"
	pmap["FontWeight"] = "font-weight"
	pmap["Gap"] = "gap"
	pmap["Grid"] = "grid"
	pmap["GridArea"] = "grid-area"
	pmap["GridAutoColumns"] = "grid-auto-columns"
	pmap["GridAutoFlow"] = "grid-auto-flow"
	pmap["GridAutoRows"] = "grid-auto-rows"
	pmap["GridColumn"] = "grid-column"
	pmap["GridColumnEnd"] = "grid-column-end"
	pmap["GridColumnStart"] = "grid-column-start"
	pmap["GridRow"] = "grid-row"
	pmap["GridRowEnd"] = "grid-row-end"
	pmap["GridRowStart"] = "grid-row-start"
	pmap["GridTemplate"] = "grid-template"
	pmap["GridTemplateAreas"] = "grid-template-areas"
	pmap["GridTemplateColumns"] = "grid-template-columns"
	pmap["GridTemplateRows"] = "grid-template-rows"
	pmap["HangingPunctuation"] = "hanging-punctuation"
	pmap["Height"] = "height"
	pmap["HypenateCharacter"] = "hypenate-character"
	pmap["Hyphens"] = "hyphens"
	pmap["ImageRendering"] = "image-rendering"
	pmap["InitialLetter"] = "initial-letter"
	pmap["InlineSize"] = "inline-size"
	pmap["Inset"] = "inset"
	pmap["InsetBlock"] = "inset-block"
	pmap["InsetBlockEnd"] = "inset-block-end"
	pmap["InsetBlockStart"] = "inset-block-start"
	pmap["InsetInline"] = "inset-inline"
	pmap["InsetInlineEnd"] = "inset-inline-end"
	pmap["InsetInlineStart"] = "inset-inline-start"
	pmap["Isolation"] = "isolation"
	pmap["JustifyContent"] = "justify-content"
	pmap["JustifyItems"] = "justify-items"
	pmap["JustifySelf"] = "justify-self"
	pmap["Left"] = "left"
	pmap["LetterSpacing"] = "letter-spacing"
	pmap["LineBreak"] = "line-break"
	pmap["LineHeight"] = "line-height"
	pmap["ListStyle"] = "list-style"
	pmap["ListStyleImage"] = "list-style-image"
	pmap["ListStylePosition"] = "list-style-position"
	pmap["ListStyleType"] = "list-style-type"
	pmap["Margin"] = "margin"
	pmap["MarginBlock"] = "margin-block"
	pmap["MarginBlockEnd"] = "margin-block-end"
	pmap["MarginBlockStart"] = "margin-block-start"
	pmap["MarginBottom"] = "margin-bottom"
	pmap["MarginInline"] = "margin-inline"
	pmap["MarginInlineEnd"] = "margin-inline-end"
	pmap["MarginInlineStart"] = "margin-inline-start"
	pmap["MarginLeft"] = "margin-left"
	pmap["MarginRight"] = "margin-right"
	pmap["MarginTop"] = "margin-top"
	pmap["Marker"] = "marker"
	pmap["MarkerEnd"] = "marker-end"
	pmap["MarkerMid"] = "marker-mid"
	pmap["MarkerStart"] = "marker-start"
	pmap["Mask"] = "mask"
	pmap["MaskClip"] = "mask-clip"
	pmap["MaskComposite"] = "mask-composite"
	pmap["MaskImage"] = "mask-image"
	pmap["MaskMode"] = "mask-mode"
	pmap["MaskOrigin"] = "mask-origin"
	pmap["MaskPosition"] = "mask-position"
	pmap["MaskRepeat"] = "mask-repeat"
	pmap["MaskSize"] = "mask-size"
	pmap["MaskType"] = "mask-type"
	pmap["MaxBlockSize"] = "max-block-size"
	pmap["MaxHeight"] = "max-height"
	pmap["MaxInlineSize"] = "max-inline-size"
	pmap["MaxWidth"] = "max-width"
	pmap["MinBlockSize"] = "min-block-size"
	pmap["MinHeight"] = "min-height"
	pmap["MinInlineSize"] = "min-inline-size"
	pmap["MinWidth"] = "min-width"
	pmap["MixBlendMode"] = "mix-blend-mode"
	pmap["ObjectFit"] = "object-fit"
	pmap["ObjectPosition"] = "object-position"
	pmap["Offset"] = "offset"
	pmap["OffsetAnchor"] = "offset-anchor"
	pmap["OffsetDistance"] = "offset-distance"
	pmap["OffsetPath"] = "offset-path"
	pmap["OffsetPosition"] = "offset-position"
	pmap["OffsetRotate"] = "offset-rotate"
	pmap["Opacity"] = "opacity"
	pmap["Order"] = "order"
	pmap["Orphans"] = "orphans"
	pmap["Outline"] = "outline"
	pmap["OutlineColor"] = "outline-color"
	pmap["OutlineOffset"] = "outline-offset"
	pmap["OutlineStyle"] = "outline-style"
	pmap["OutlineWidth"] = "outline-width"
	pmap["Overflow"] = "overflow"
	pmap["OverflowAnchor"] = "overflow-anchor"
	pmap["OverflowWrap"] = "overflow-wrap"
	pmap["OverflowX"] = "overflow-x"
	pmap["OverflowY"] = "overflow-y"
	pmap["OverscrollBehavior"] = "overscroll-behavior"
	pmap["OverscrollBehaviorBlock"] = "overscroll-behavior-block"
	pmap["OverscrollBehaviorInline"] = "overscroll-behavior-inline"
	pmap["OverscrollBehaviorX"] = "overscroll-behavior-x"
	pmap["OverscrollBehaviorY"] = "overscroll-behavior-y"
	pmap["Padding"] = "padding"
	pmap["PaddingBlock"] = "padding-block"
	pmap["PaddingBlockEnd"] = "padding-block-end"
	pmap["PaddingBlockStart"] = "padding-block-start"
	pmap["PaddingBottom"] = "padding-bottom"
	pmap["PaddingInline"] = "padding-inline"
	pmap["PaddingInlineEnd"] = "padding-inline-end"
	pmap["PaddingInlineStart"] = "padding-inline-start"
	pmap["PaddingLeft"] = "padding-left"
	pmap["PaddingRight"] = "padding-right"
	pmap["PaddingTop"] = "padding-top"
	pmap["PageBreakAfter"] = "page-break-after"
	pmap["PageBreakBefore"] = "page-break-before"
	pmap["PageBreakInside"] = "page-break-inside"
	pmap["PaintOrder"] = "paint-order"
	pmap["Perspective"] = "perspective"
	pmap["PerspectiveOrigin"] = "perspective-origin"
	pmap["PlaceContent"] = "place-content"
	pmap["PlaceItems"] = "place-items"
	pmap["PlaceSelf"] = "place-self"
	pmap["PointerEvents"] = "pointer-events"
	pmap["Position"] = "position"
	pmap["Quotes"] = "quotes"
	pmap["Resize"] = "resize"
	pmap["Right"] = "right"
	pmap["Rotate"] = "rotate"
	pmap["RowGap"] = "row-gap"
	pmap["Scale"] = "scale"
	pmap["ScrollBehavior"] = "scroll-behavior"
	pmap["ScrollMargin"] = "scroll-margin"
	pmap["ScrollMarginBlock"] = "scroll-margin-block"
	pmap["ScrollMarginBlockEnd"] = "scroll-margin-block-end"
	pmap["ScrollMarginBlockStart"] = "scroll-margin-block-start"
	pmap["ScrollMarginBottom"] = "scroll-margin-bottom"
	pmap["ScrollMarginInline"] = "scroll-margin-inline"
	pmap["ScrollMarginInlineEnd"] = "scroll-margin-inline-end"
	pmap["ScrollMarginInlineStart"] = "scroll-margin-inline-start"
	pmap["ScrollMarginLeft"] = "scroll-margin-left"
	pmap["ScrollMarginRight"] = "scroll-margin-right"
	pmap["ScrollMarginTop"] = "scroll-margin-top"
	pmap["ScrollPadding"] = "scroll-padding"
	pmap["ScrollPaddingBlock"] = "scroll-padding-block"
	pmap["ScrollPaddingBlockEnd"] = "scroll-padding-block-end"
	pmap["ScrollPaddingBlockStart"] = "scroll-padding-block-start"
	pmap["ScrollPaddingBottom"] = "scroll-padding-bottom"
	pmap["ScrollPaddingInline"] = "scroll-padding-inline"
	pmap["ScrollPaddingInlineEnd"] = "scroll-padding-inline-end"
	pmap["ScrollPaddingInlineStart"] = "scroll-padding-inline-start"
	pmap["ScrollPaddingLeft"] = "scroll-padding-left"
	pmap["ScrollPaddingRight"] = "scroll-padding-right"
	pmap["ScrollPaddingTop"] = "scroll-padding-top"
	pmap["ScrollSnapAlign"] = "scroll-snap-align"
	pmap["ScrollSnapStop"] = "scroll-snap-stop"
	pmap["ScrollSnapType"] = "scroll-snap-type"
	pmap["ScrollbarColor"] = "scrollbar-color"
	pmap["TabSize"] = "tab-size"
	pmap["TableLayout"] = "table-layout"
	pmap["TextAlign"] = "text-align"
	pmap["TextAlignLast"] = "text-align-last"
	pmap["TextCombineUpright"] = "text-combine-upright"
	pmap["TextDecoration"] = "text-decoration"
	pmap["TextDecorationColor"] = "text-decoration-color"
	pmap["TextDecorationLine"] = "text-decoration-line"
	pmap["TextDecorationStyle"] = "text-decoration-style"
	pmap["TextDecorationThickness"] = "text-decoration-thickness"
	pmap["TextEmphasis"] = "text-emphasis"
	pmap["TextEmphasisColor"] = "text-emphasis-color"
	pmap["TextEmphasisPosition"] = "text-emphasis-position"
	pmap["TextEmphasisStyle"] = "text-emphasis-style"
	pmap["TextIndent"] = "text-indent"
	pmap["TextJustify"] = "text-justify"
	pmap["TextOrientation"] = "text-orientation"
	pmap["TextOverflow"] = "text-overflow"
	pmap["TextShadow"] = "text-shadow"
	pmap["TextTransform"] = "text-transform"
	pmap["TextUnderlineOffset"] = "text-underline-offset"
	pmap["TextUnderlinePosition"] = "text-underline-position"
	pmap["Top"] = "top"
	pmap["Transform"] = "transform"
	pmap["TransformOrigin"] = "transform-origin"
	pmap["TransformStyle"] = "transform-style"
	pmap["Transition"] = "transition"
	pmap["TransitionDelay"] = "transition-delay"
	pmap["TransitionDuration"] = "transition-duration"
	pmap["TransitionProperty"] = "transition-property"
	pmap["TransitionTimingFunction"] = "transition-timing-function"
	pmap["Translate"] = "translate"
	pmap["UnicodeBidi"] = "unicode-bidi"
	pmap["UserSelect"] = "user-select"
	pmap["VerticalAlign"] = "vertical-align"
	pmap["Visibility"] = "visibility"
	pmap["WhiteSpace"] = "white-space"
	pmap["Widows"] = "widows"
	pmap["Width"] = "width"
	pmap["WordBreak"] = "word-break"
	pmap["WordSpacing"] = "word-spacing"
	pmap["WordWrap"] = "word-wrap"
	pmap["WritingMode"] = "writing-mode"
	pmap["ZIndex"] = "z-index"
	return &PropMap{
		pmap: pmap,
	}
}
