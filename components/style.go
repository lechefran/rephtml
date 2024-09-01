package rephtml

import (
	"bytes"
	"reflect"
)

/*
* Implementation of all CSS properties
* This may be overkill since a bulk majority of these
* properties may not be relevant currently
* or may never be relevant for this project
 */
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
	Charset                  string // parsed as @charset
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
	FontFace                 string // parsed as @font-face
	FontFamily               string
	FontFeatureSettings      string
	FontFeatureValues        string // parsed as @font-feature-values
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
	Import                   string // parsed as @import
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
	Keyframes                string // parsed as @keyframes
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
	Media                    string // parsed as @media
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
	OverflowAnchor           string // parsed as @overflow-anchor
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

func (p *CssProps) NewCssProps() *CssProps {
	return &CssProps{}
}

type Style struct {
	buf   bytes.Buffer
	pmap  *PropMap
	Props CssProps
	Tags  []string
}

func (s *Style) Bytes() []byte {
	return s.buf.Bytes()
}

// todo: use byte array over string
func (s *Style) Prepare() {
	res := ""
	val := reflect.ValueOf(s.Props)
	t := val.Type()
	for i := 0; i < val.NumField(); i++ {
		k, v := t.Field(i).Name, val.Field(i)
		if k != "" && v.String() != "" {
			res += s.pmap.pmap[k] + ":" + v.String() + ";"
		}
	}
	s.buf.WriteString(formatStringArray(s.Tags) + "{" + res + "}")
}

func formatStringArray(sarr []string) string {
	res := ""
	for i := 0; i < len(sarr); i++ {
		if i != len(sarr)-1 {
			res += sarr[i] + ","
		} else {
			res += sarr[i]
		}
	}
	return res
}

func (s *Style) PropMap(p *PropMap) *Style {
	s.pmap = p
	return s
}

type PropMap struct {
	pmap  map[string]string
	count int
}

func NewPropMap() *PropMap {
	pmap := map[string]string{}
	pmap["@charset"] = "@charset"
	pmap["@fontFace"] = "@font-face"
	pmap["@fontFeatureValues"] = "@font-feature-values"
	pmap["@import"] = "@import"
	pmap["@keyframes"] = "@keyframes"
	pmap["@media"] = "@media"
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
		pmap:  pmap,
		count: len(pmap),
	}
}
