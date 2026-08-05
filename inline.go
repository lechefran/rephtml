package rephtml

import (
	"fmt"
	"regexp"
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

	timeWeekPattern     = regexp.MustCompile(`^\d{4}-W(0[1-9]|[1-4]\d|5[0-3])$`)
	timeDurationPattern = regexp.MustCompile(`^P(?:\d+Y)?(?:\d+M)?(?:\d+W)?(?:\d+D)?(?:T(?:\d+H)?(?:\d+M)?(?:\d+(?:\.\d+)?S)?)?$`)
)

// Anchor represents the Anchor component or supporting type.
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

// Link sets the link value on the Anchor component.
func (a *Anchor) Link(s string) *Anchor {
	a.link = s
	return a
}

// prepare renders the Anchor component into its internal buffer.
func (a *Anchor) prepare() {
	tg := openTag(&a.buf, "a")
	tg.styleAttr(a.style)
	tg.attr("href", a.link)
	tg.text(a.text)
}

// Abbr represents the Abbr component or supporting type.
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

// Title sets the title value on the Abbr component.
func (a *Abbr) Title(s string) *Abbr {
	a.title = s
	return a
}

// prepare renders the Abbr component into its internal buffer.
func (a *Abbr) prepare() {
	tg := openTag(&a.buf, "abbr")
	tg.styleAttr(a.style)
	tg.attr("title", a.title)
	tg.text(a.text)
}

// B represents the B component or supporting type.
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

// prepare renders the B component into its internal buffer.
func (b *B) prepare() {
	tg := openTag(&b.buf, "b")
	tg.styleAttr(b.style)
	tg.text(b.text)
}

// I represents the I component or supporting type.
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

// prepare renders the I component into its internal buffer.
func (i *I) prepare() {
	tg := openTag(&i.buf, "i")
	tg.styleAttr(i.style)
	tg.text(i.text)
}

// Q represents the Q component or supporting type.
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

// Cite sets the cite value on the Q component.
func (q *Q) Cite(s string) *Q {
	q.cite = s
	return q
}

// prepare renders the Q component into its internal buffer.
func (q *Q) prepare() {
	tg := openTag(&q.buf, "q")
	tg.styleAttr(q.style)
	tg.attr("cite", q.cite)
	tg.text(q.text)
}

// S represents the S component or supporting type.
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

// prepare renders the S component into its internal buffer.
func (s *S) prepare() {
	tg := openTag(&s.buf, "s")
	tg.styleAttr(s.style)
	tg.text(s.text)
}

// U represents the U component or supporting type.
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

// prepare renders the U component into its internal buffer.
func (u *U) prepare() {
	tg := openTag(&u.buf, "u")
	tg.styleAttr(u.style)
	tg.text(u.text)
}

// Bdi represents the bdi component or supporting type.
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

// prepare renders the Bdi component into its internal buffer.
func (d *Bdi) prepare() {
	tg := openTag(&d.buf, "bdi")
	tg.styleAttr(d.style)
	tg.text(d.text)
}

// Bdo represents the bdo component or supporting type.
type Bdo struct {
	bodyElement
	textNode[*Bdo]
}

// NewBdo creates a new bdo component.
func NewBdo() *Bdo {
	v := &Bdo{}
	v.init(v)
	return v
}

// prepare renders the Bdo component into its internal buffer.
func (d *Bdo) prepare() {
	tg := openTag(&d.buf, "bdo")
	tg.styleAttr(d.style)
	tg.text(d.text)
}

// Br represents the Br component or supporting type.
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

// prepare renders the Br component into its internal buffer.
func (br *Br) prepare() {
	tg := openTag(&br.buf, "br")
	tg.styleAttr(br.style)
	tg.void()
}

// Cite represents the Cite component or supporting type.
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

// prepare renders the Cite component into its internal buffer.
func (c *Cite) prepare() {
	tg := openTag(&c.buf, "cite")
	tg.styleAttr(c.style)
	tg.text(c.text)
}

// Code represents the Code component or supporting type.
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

// prepare renders the Code component into its internal buffer.
func (c *Code) prepare() {
	tg := openTag(&c.buf, "code")
	tg.styleAttr(c.style)
	tg.text(c.text)
}

// Data represents the Data component or supporting type.
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

// Value sets the value value on the Data component.
func (d *Data) Value(s string) *Data {
	d.value = s
	return d
}

// prepare renders the Data component into its internal buffer.
func (d *Data) prepare() {
	tg := openTag(&d.buf, "data")
	tg.styleAttr(d.style)
	tg.attr("value", d.value)
	tg.text(d.text)
}

// Dfn represents the Dfn component or supporting type.
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

// Title sets the title value on the Dfn component.
func (d *Dfn) Title(s string) *Dfn {
	d.title = s
	return d
}

// prepare renders the Dfn component into its internal buffer.
func (d *Dfn) prepare() {
	tg := openTag(&d.buf, "dfn")
	tg.styleAttr(d.style)
	tg.attr("title", d.title)
	tg.text(d.text)
}

// Em represents the em component or supporting type.
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

// prepare renders the Em component into its internal buffer.
func (e *Em) prepare() {
	tg := openTag(&e.buf, "em")
	tg.styleAttr(e.style)
	tg.text(e.text)
}

// Mark represents the Mark component or supporting type.
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

// prepare renders the Mark component into its internal buffer.
func (m *Mark) prepare() {
	tg := openTag(&m.buf, "mark")
	tg.styleAttr(m.style)
	tg.text(m.text)
}

// Ruby represents the Ruby component or supporting type.
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

// prepare renders the Ruby component into its internal buffer.
func (r *Ruby) prepare() {
	tg := openTag(&r.buf, "ruby")
	tg.styleAttr(r.style)
	tg.children(r.contents)
}

// Rb represents the Rb component or supporting type.
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

// prepare renders the Rb component into its internal buffer.
func (rb *Rb) prepare() {
	tg := openTag(&rb.buf, "rb")
	tg.styleAttr(rb.style)
	tg.text(rb.text)
}

// Rt represents the Rt component or supporting type.
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

// prepare renders the Rt component into its internal buffer.
func (rt *Rt) prepare() {
	tg := openTag(&rt.buf, "rt")
	tg.styleAttr(rt.style)
	tg.text(rt.text)
}

// Rtc represents the Rtc component or supporting type.
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

// prepare renders the Rtc component into its internal buffer.
func (rtc *Rtc) prepare() {
	tg := openTag(&rtc.buf, "rtc")
	tg.styleAttr(rtc.style)
	tg.children(rtc.contents)
}

// Rp represents the Rp component or supporting type.
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

// prepare renders the Rp component into its internal buffer.
func (rp *Rp) prepare() {
	tg := openTag(&rp.buf, "rp")
	tg.styleAttr(rp.style)
	tg.text(rp.text)
}

// Kbd represents the Kbd component or supporting type.
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

// prepare renders the Kbd component into its internal buffer.
func (k *Kbd) prepare() {
	tg := openTag(&k.buf, "kbd")
	tg.styleAttr(k.style)
	tg.text(k.text)
}

// Sub represents the Sub component or supporting type.
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

// prepare renders the Sub component into its internal buffer.
func (s *Sub) prepare() {
	tg := openTag(&s.buf, "sub")
	tg.styleAttr(s.style)
	tg.text(s.text)
}

// Sup represents the Sup component or supporting type.
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

// prepare renders the Sup component into its internal buffer.
func (s *Sup) prepare() {
	tg := openTag(&s.buf, "sup")
	tg.styleAttr(s.style)
	tg.text(s.text)
}

// Samp represents the Samp component or supporting type.
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

// prepare renders the Samp component into its internal buffer.
func (s *Samp) prepare() {
	tg := openTag(&s.buf, "samp")
	tg.styleAttr(s.style)
	tg.text(s.text)
}

// Small represents the Small component or supporting type.
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

// prepare renders the Small component into its internal buffer.
func (s *Small) prepare() {
	tg := openTag(&s.buf, "small")
	tg.styleAttr(s.style)
	tg.text(s.text)
}

// Span represents the Span component or supporting type.
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

// prepare renders the Span component into its internal buffer.
func (s *Span) prepare() {
	tg := openTag(&s.buf, "span")
	tg.styleAttr(s.style)
	tg.text(s.text)
}

// Strong represents the Strong component or supporting type.
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

// prepare renders the Strong component into its internal buffer.
func (s *Strong) prepare() {
	tg := openTag(&s.buf, "strong")
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
		t.err = fmt.Errorf("invalid datetime value: %s", dt)
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

	if timeWeekPattern.MatchString(dt) {
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

// isValidDuration reports whether a value matches a non-empty ISO-like duration.
func isValidDuration(dt string) bool {
	if !timeDurationPattern.MatchString(dt) {
		return false
	}
	if dt == "P" || dt == "PT" || strings.HasSuffix(dt, "T") {
		return false
	}
	return true
}

// prepare renders the Time component into its internal buffer.
func (t *Time) prepare() {
	tg := openTag(&t.buf, "time")
	tg.attr("datetime", t.datetime)
	tg.styleAttr(t.style)
	tg.text(t.text)
}

// Var represents the Var component or supporting type.
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

// prepare renders the Var component into its internal buffer.
func (v *Var) prepare() {
	tg := openTag(&v.buf, "var")
	tg.styleAttr(v.style)
	tg.text(v.text)
}

// Wbr represents the Wbr component or supporting type.
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

// prepare renders the Wbr component into its internal buffer.
func (w *Wbr) prepare() {
	tg := openTag(&w.buf, "wbr")
	tg.styleAttr(w.style)
	tg.void()
}
