package rephtml

import (
	"bytes"
	"fmt"
	"strings"
	"time"
)

var (
	timeDatetimeLayouts = []string{
		"2006-01-02T15:04:05Z07:00",
		"2006-01-02T15:04:05Z",
		"2006-01-02T15:04:05",
		"2006-01-02T15:04Z07:00",
		"2006-01-02T15:04Z",
		"2006-01-02T15:04",
		"2006-01-02",
		"2006-01",
		"2006",
		"15:04:05",
		"15:04",
	}
)

// Anchor represents the HTML <a> element, a hyperlink to another page, a location on the page or another URL.
type Anchor struct {
	bodyElement
	textNode[*Anchor]
	link string
}

// NewAnchor creates a new Anchor component.
func NewAnchor() *Anchor {
	v := &Anchor{}
	v.init(v)
	return v
}

// Link sets the link attribute.
func (a *Anchor) Link(s string) *Anchor {
	a.link = s
	return a
}

// renderTo writes the Anchor component's HTML to buf.
func (a *Anchor) renderTo(buf *bytes.Buffer) {
	tg := startTag(buf, "a")
	tg.styleAttr(a.style)
	tg.urlAttr("href", a.link)
	tg.text(a.text)
}

// Abbr represents the HTML <abbr> element, an abbreviation or acronym.
type Abbr struct {
	bodyElement
	textNode[*Abbr]
	title string
}

// NewAbbr creates a new Abbr component.
func NewAbbr() *Abbr {
	v := &Abbr{}
	v.init(v)
	return v
}

// Title sets the title attribute.
func (a *Abbr) Title(s string) *Abbr {
	a.title = s
	return a
}

// renderTo writes the Abbr component's HTML to buf.
func (a *Abbr) renderTo(buf *bytes.Buffer) {
	tg := startTag(buf, "abbr")
	tg.styleAttr(a.style)
	tg.attr("title", a.title)
	tg.text(a.text)
}

// B represents the HTML <b> element, text drawn to attention without conveying extra importance.
type B struct {
	bodyElement
	textNode[*B]
}

// NewB creates a new B component.
func NewB() *B {
	v := &B{}
	v.init(v)
	return v
}

// renderTo writes the B component's HTML to buf.
func (b *B) renderTo(buf *bytes.Buffer) {
	tg := startTag(buf, "b")
	tg.styleAttr(b.style)
	tg.text(b.text)
}

// I represents the HTML <i> element, text in an alternate voice or mood, such as a technical term.
type I struct {
	bodyElement
	textNode[*I]
}

// NewI creates a new I component.
func NewI() *I {
	v := &I{}
	v.init(v)
	return v
}

// renderTo writes the I component's HTML to buf.
func (i *I) renderTo(buf *bytes.Buffer) {
	tg := startTag(buf, "i")
	tg.styleAttr(i.style)
	tg.text(i.text)
}

// Q represents the HTML <q> element, a short inline quotation.
type Q struct {
	bodyElement
	textNode[*Q]
	cite string
}

// NewQ creates a new Q component.
func NewQ() *Q {
	v := &Q{}
	v.init(v)
	return v
}

// Cite sets the cite attribute.
func (q *Q) Cite(s string) *Q {
	q.cite = s
	return q
}

// renderTo writes the Q component's HTML to buf.
func (q *Q) renderTo(buf *bytes.Buffer) {
	tg := startTag(buf, "q")
	tg.styleAttr(q.style)
	tg.urlAttr("cite", q.cite)
	tg.text(q.text)
}

// S represents the HTML <s> element, text that is no longer accurate or relevant, rendered struck through.
type S struct {
	bodyElement
	textNode[*S]
}

// NewS creates a new S component.
func NewS() *S {
	v := &S{}
	v.init(v)
	return v
}

// renderTo writes the S component's HTML to buf.
func (s *S) renderTo(buf *bytes.Buffer) {
	tg := startTag(buf, "s")
	tg.styleAttr(s.style)
	tg.text(s.text)
}

// U represents the HTML <u> element, text with an unarticulated annotation, rendered underlined.
type U struct {
	bodyElement
	textNode[*U]
}

// NewU creates a new U component.
func NewU() *U {
	v := &U{}
	v.init(v)
	return v
}

// renderTo writes the U component's HTML to buf.
func (u *U) renderTo(buf *bytes.Buffer) {
	tg := startTag(buf, "u")
	tg.styleAttr(u.style)
	tg.text(u.text)
}

// Bdi represents the HTML <bdi> element, text isolated from its surroundings for bidirectional text formatting.
type Bdi struct {
	bodyElement
	textNode[*Bdi]
}

// NewBdi creates a new bdi component.
func NewBdi() *Bdi {
	v := &Bdi{}
	v.init(v)
	return v
}

// renderTo writes the Bdi component's HTML to buf.
func (d *Bdi) renderTo(buf *bytes.Buffer) {
	tg := startTag(buf, "bdi")
	tg.styleAttr(d.style)
	tg.text(d.text)
}

// Bdo represents the HTML <bdo> element, text whose bidirectional direction is overridden.
type Bdo struct {
	bodyElement
	textNode[*Bdo]
	dir string
}

// NewBdo creates a new bdo component.
func NewBdo() *Bdo {
	v := &Bdo{}
	v.init(v)
	return v
}

// Dir sets the dir attribute, the text direction to force: "ltr" or "rtl". The
// bdo element has no effect without it.
func (d *Bdo) Dir(dir string) *Bdo {
	d.dir = dir
	return d
}

// renderTo writes the Bdo component's HTML to buf.
func (d *Bdo) renderTo(buf *bytes.Buffer) {
	tg := startTag(buf, "bdo")
	tg.attr("dir", d.dir)
	tg.styleAttr(d.style)
	tg.text(d.text)
}

// Br represents the HTML <br> element, a line break.
type Br struct {
	bodyElement
	node[*Br]
}

// NewBr creates a new Br component.
func NewBr() *Br {
	v := &Br{}
	v.init(v)
	return v
}

// renderTo writes the Br component's HTML to buf.
func (br *Br) renderTo(buf *bytes.Buffer) {
	tg := startTag(buf, "br")
	tg.styleAttr(br.style)
	tg.void()
}

// Cite represents the HTML <cite> element, the title of a creative work.
type Cite struct {
	bodyElement
	textNode[*Cite]
}

// NewCite creates a new Cite component.
func NewCite() *Cite {
	v := &Cite{}
	v.init(v)
	return v
}

// renderTo writes the Cite component's HTML to buf.
func (c *Cite) renderTo(buf *bytes.Buffer) {
	tg := startTag(buf, "cite")
	tg.styleAttr(c.style)
	tg.text(c.text)
}

// Code represents the HTML <code> element, a fragment of computer code.
type Code struct {
	bodyElement
	textNode[*Code]
}

// NewCode creates a new Code component.
func NewCode() *Code {
	v := &Code{}
	v.init(v)
	return v
}

// renderTo writes the Code component's HTML to buf.
func (c *Code) renderTo(buf *bytes.Buffer) {
	tg := startTag(buf, "code")
	tg.styleAttr(c.style)
	tg.text(c.text)
}

// Data represents the HTML <data> element, content together with a machine-readable translation of it.
type Data struct {
	bodyElement
	textNode[*Data]
	value string
}

// NewData creates a new Data component.
func NewData() *Data {
	v := &Data{}
	v.init(v)
	return v
}

// Value sets the value attribute.
func (d *Data) Value(s string) *Data {
	d.value = s
	return d
}

// renderTo writes the Data component's HTML to buf.
func (d *Data) renderTo(buf *bytes.Buffer) {
	tg := startTag(buf, "data")
	tg.styleAttr(d.style)
	tg.attr("value", d.value)
	tg.text(d.text)
}

// Dfn represents the HTML <dfn> element, the term being defined.
type Dfn struct {
	bodyElement
	textNode[*Dfn]
	title string
}

// NewDfn creates a new Dfn component.
func NewDfn() *Dfn {
	v := &Dfn{}
	v.init(v)
	return v
}

// Title sets the title attribute.
func (d *Dfn) Title(s string) *Dfn {
	d.title = s
	return d
}

// renderTo writes the Dfn component's HTML to buf.
func (d *Dfn) renderTo(buf *bytes.Buffer) {
	tg := startTag(buf, "dfn")
	tg.styleAttr(d.style)
	tg.attr("title", d.title)
	tg.text(d.text)
}

// Em represents the HTML <em> element, text with stress emphasis.
type Em struct {
	bodyElement
	textNode[*Em]
}

// NewEm creates a new em component.
func NewEm() *Em {
	v := &Em{}
	v.init(v)
	return v
}

// renderTo writes the Em component's HTML to buf.
func (e *Em) renderTo(buf *bytes.Buffer) {
	tg := startTag(buf, "em")
	tg.styleAttr(e.style)
	tg.text(e.text)
}

// Mark represents the HTML <mark> element, text highlighted for reference.
type Mark struct {
	bodyElement
	textNode[*Mark]
}

// NewMark creates a new Mark component.
func NewMark() *Mark {
	v := &Mark{}
	v.init(v)
	return v
}

// renderTo writes the Mark component's HTML to buf.
func (m *Mark) renderTo(buf *bytes.Buffer) {
	tg := startTag(buf, "mark")
	tg.styleAttr(m.style)
	tg.text(m.text)
}

// Ruby represents the HTML <ruby> element, a ruby annotation, used for East Asian pronunciation guides.
type Ruby struct {
	bodyElement
	contentNode[*Ruby]
}

// NewRuby creates a new Ruby component.
func NewRuby() *Ruby {
	v := &Ruby{}
	v.init(v)
	return v
}

// renderTo writes the Ruby component's HTML to buf.
func (r *Ruby) renderTo(buf *bytes.Buffer) {
	tg := startTag(buf, "ruby")
	tg.styleAttr(r.style)
	tg.children(r.contents)
}

// Rb represents the HTML <rb> element, the base text of a ruby annotation. It is obsolete.
type Rb struct {
	bodyElement
	textNode[*Rb]
}

// NewRb creates a new Rb component.
func NewRb() *Rb {
	v := &Rb{}
	v.init(v)
	return v
}

// renderTo writes the Rb component's HTML to buf.
func (rb *Rb) renderTo(buf *bytes.Buffer) {
	tg := startTag(buf, "rb")
	tg.styleAttr(rb.style)
	tg.text(rb.text)
}

// Rt represents the HTML <rt> element, the ruby text of a ruby annotation, such as a pronunciation.
type Rt struct {
	bodyElement
	textNode[*Rt]
}

// NewRt creates a new Rt component.
func NewRt() *Rt {
	v := &Rt{}
	v.init(v)
	return v
}

// renderTo writes the Rt component's HTML to buf.
func (rt *Rt) renderTo(buf *bytes.Buffer) {
	tg := startTag(buf, "rt")
	tg.styleAttr(rt.style)
	tg.text(rt.text)
}

// Rtc represents the HTML <rtc> element, a container for ruby text components. It is obsolete.
type Rtc struct {
	bodyElement
	contentNode[*Rtc]
}

// NewRtc creates a new Rtc component.
func NewRtc() *Rtc {
	v := &Rtc{}
	v.init(v)
	return v
}

// renderTo writes the Rtc component's HTML to buf.
func (rtc *Rtc) renderTo(buf *bytes.Buffer) {
	tg := startTag(buf, "rtc")
	tg.styleAttr(rtc.style)
	tg.children(rtc.contents)
}

// Rp represents the HTML <rp> element, fallback parentheses for browsers that do not support ruby annotations.
type Rp struct {
	bodyElement
	textNode[*Rp]
}

// NewRp creates a new Rp component.
func NewRp() *Rp {
	v := &Rp{}
	v.init(v)
	return v
}

// renderTo writes the Rp component's HTML to buf.
func (rp *Rp) renderTo(buf *bytes.Buffer) {
	tg := startTag(buf, "rp")
	tg.styleAttr(rp.style)
	tg.text(rp.text)
}

// Kbd represents the HTML <kbd> element, keyboard or other text input.
type Kbd struct {
	bodyElement
	textNode[*Kbd]
}

// NewKbd creates a new Kbd component.
func NewKbd() *Kbd {
	v := &Kbd{}
	v.init(v)
	return v
}

// renderTo writes the Kbd component's HTML to buf.
func (k *Kbd) renderTo(buf *bytes.Buffer) {
	tg := startTag(buf, "kbd")
	tg.styleAttr(k.style)
	tg.text(k.text)
}

// Sub represents the HTML <sub> element, subscript text.
type Sub struct {
	bodyElement
	textNode[*Sub]
}

// NewSub creates a new Sub component.
func NewSub() *Sub {
	v := &Sub{}
	v.init(v)
	return v
}

// renderTo writes the Sub component's HTML to buf.
func (s *Sub) renderTo(buf *bytes.Buffer) {
	tg := startTag(buf, "sub")
	tg.styleAttr(s.style)
	tg.text(s.text)
}

// Sup represents the HTML <sup> element, superscript text.
type Sup struct {
	bodyElement
	textNode[*Sup]
}

// NewSup creates a new Sup component.
func NewSup() *Sup {
	v := &Sup{}
	v.init(v)
	return v
}

// renderTo writes the Sup component's HTML to buf.
func (s *Sup) renderTo(buf *bytes.Buffer) {
	tg := startTag(buf, "sup")
	tg.styleAttr(s.style)
	tg.text(s.text)
}

// Samp represents the HTML <samp> element, sample output from a computer program.
type Samp struct {
	bodyElement
	textNode[*Samp]
}

// NewSamp creates a new Samp component.
func NewSamp() *Samp {
	v := &Samp{}
	v.init(v)
	return v
}

// renderTo writes the Samp component's HTML to buf.
func (s *Samp) renderTo(buf *bytes.Buffer) {
	tg := startTag(buf, "samp")
	tg.styleAttr(s.style)
	tg.text(s.text)
}

// Small represents the HTML <small> element, side comments and small print.
type Small struct {
	bodyElement
	textNode[*Small]
}

// NewSmall creates a new Small component.
func NewSmall() *Small {
	v := &Small{}
	v.init(v)
	return v
}

// renderTo writes the Small component's HTML to buf.
func (s *Small) renderTo(buf *bytes.Buffer) {
	tg := startTag(buf, "small")
	tg.styleAttr(s.style)
	tg.text(s.text)
}

// Span represents the HTML <span> element, a generic inline container with no semantic meaning.
type Span struct {
	bodyElement
	textNode[*Span]
}

// NewSpan creates a new Span component.
func NewSpan() *Span {
	v := &Span{}
	v.init(v)
	return v
}

// renderTo writes the Span component's HTML to buf.
func (s *Span) renderTo(buf *bytes.Buffer) {
	tg := startTag(buf, "span")
	tg.styleAttr(s.style)
	tg.text(s.text)
}

// Strong represents the HTML <strong> element, text of strong importance.
type Strong struct {
	bodyElement
	textNode[*Strong]
}

// NewStrong creates a new Strong component.
func NewStrong() *Strong {
	v := &Strong{}
	v.init(v)
	return v
}

// renderTo writes the Strong component's HTML to buf.
func (s *Strong) renderTo(buf *bytes.Buffer) {
	tg := startTag(buf, "strong")
	tg.styleAttr(s.style)
	tg.text(s.text)
}

// Time represents the HTML time element and its machine-readable datetime value.
type Time struct {
	bodyElement
	textNode[*Time]
	datetime string
	err      error
}

// NewTime creates a time element.
func NewTime() *Time {
	v := &Time{}
	v.init(v)
	return v
}

// Datetime sets the machine-readable datetime attribute when dt is valid.
func (t *Time) Datetime(dt string) *Time {
	if !isValidDatetime(dt) {
		t.datetime = ""
		t.err = fmt.Errorf("invalid datetime value: %q", dt)
		return t
	}
	t.datetime = dt
	t.err = nil
	return t
}

// Err returns the last datetime validation error recorded on the time element.
func (t *Time) Err() error {
	return t.err
}

// isValidDatetime reports whether a value matches common HTML datetime forms.
func isValidDatetime(dt string) bool {
	if dt == "" {
		return true
	}

	if isValidWeek(dt) {
		return true
	}

	if isValidDuration(dt) {
		return true
	}

	for _, format := range timeDatetimeLayouts {
		if _, err := time.Parse(format, dt); err == nil {
			return true
		}
	}

	return false
}

// isValidWeek reports whether a value is a year and week such as 2024-W07,
// with a week from 01 to 53.
func isValidWeek(dt string) bool {
	if len(dt) != 8 || dt[4] != '-' || dt[5] != 'W' || !isDigits(dt[:4]) || !isDigits(dt[6:]) {
		return false
	}
	week := int(dt[6]-'0')*10 + int(dt[7]-'0')
	return week >= 1 && week <= 53
}

// isValidDuration reports whether a value is a non-empty ISO-like duration
// such as P1DT2H30M: date units Y, M, W, D in order, then an optional T followed
// by time units H, M, S in order. Only seconds may carry a fraction.
func isValidDuration(dt string) bool {
	if dt == "" || dt[0] != 'P' || dt == "P" || dt == "PT" || strings.HasSuffix(dt, "T") {
		return false
	}
	i, ok := scanDurationUnits(dt, 1, "YMWD", false)
	if !ok {
		return false
	}
	if i < len(dt) && dt[i] == 'T' {
		i, ok = scanDurationUnits(dt, i+1, "HMS", true)
		if !ok {
			return false
		}
	}
	return i == len(dt)
}

// scanDurationUnits consumes number-and-unit pairs from s starting at i. The
// units must appear in the order given. When fractionLast is set, the final
// unit may have a decimal fraction. It returns the index after the last pair,
// and false if a number is not followed by a valid unit.
func scanDurationUnits(s string, i int, units string, fractionLast bool) (int, bool) {
	next := 0
	for {
		start := i
		for i < len(s) && isDigit(s[i]) {
			i++
		}
		if i == start {
			return i, true
		}
		fraction := false
		if fractionLast && i+1 < len(s) && s[i] == '.' && isDigit(s[i+1]) {
			fraction = true
			i++
			for i < len(s) && isDigit(s[i]) {
				i++
			}
		}
		if i >= len(s) {
			return i, false
		}
		unit := strings.IndexByte(units[next:], s[i])
		if unit < 0 || (fraction && next+unit != len(units)-1) {
			return i, false
		}
		next += unit + 1
		i++
	}
}

// isDigit reports whether c is an ASCII digit.
func isDigit(c byte) bool {
	return c >= '0' && c <= '9'
}

// isDigits reports whether s is made up only of ASCII digits.
func isDigits(s string) bool {
	for i := 0; i < len(s); i++ {
		if !isDigit(s[i]) {
			return false
		}
	}
	return true
}

// renderTo writes the Time component's HTML to buf.
func (t *Time) renderTo(buf *bytes.Buffer) {
	tg := startTag(buf, "time")
	tg.attr("datetime", t.datetime)
	tg.styleAttr(t.style)
	tg.text(t.text)
}

// Var represents the HTML <var> element, a variable in a mathematical expression or program.
type Var struct {
	bodyElement
	textNode[*Var]
}

// NewVar creates a new Var component.
func NewVar() *Var {
	v := &Var{}
	v.init(v)
	return v
}

// renderTo writes the Var component's HTML to buf.
func (v *Var) renderTo(buf *bytes.Buffer) {
	tg := startTag(buf, "var")
	tg.styleAttr(v.style)
	tg.text(v.text)
}

// Wbr represents the HTML <wbr> element, a position where the browser may break a line.
type Wbr struct {
	bodyElement
	node[*Wbr]
}

// NewWbr creates a new Wbr component.
func NewWbr() *Wbr {
	v := &Wbr{}
	v.init(v)
	return v
}

// renderTo writes the Wbr component's HTML to buf.
func (w *Wbr) renderTo(buf *bytes.Buffer) {
	tg := startTag(buf, "wbr")
	tg.styleAttr(w.style)
	tg.void()
}
