package rephtml

import (
	"bytes"
	"log"
	"regexp"
	"time"
)

type Anchor struct {
	buf        bytes.Buffer
	style      map[string]string
	link, text string
}

func NewAnchor() *Anchor {
	return &Anchor{
		style: make(map[string]string),
	}
}

func (a *Anchor) AddStyle(k, v string) *Anchor {
	a.style[k] = v
	return a
}

func (a *Anchor) Style(m map[string]string) *Anchor {
	a.style = m
	return a
}

func (a *Anchor) Text(s string) *Anchor {
	a.text = s
	return a
}

func (a *Anchor) Link(s string) *Anchor {
	a.link = s
	return a
}

func (a *Anchor) Bytes() []byte {
	return cloneBytes(a.buf.Bytes())
}

// IsBodyElement implements BodyElement interface
func (a *Anchor) IsBodyElement() {}

func (a *Anchor) Prepare() {
	a.buf.Reset()
	a.buf.WriteString("<a")
	if len(a.style) != 0 {
		parseStyle(&a.buf, a.style)
	}
	if a.link != "" {
		writeAttr(&a.buf, "href", a.link)
	}
	a.buf.WriteString(">" + escapeText(a.text) + "</a>")
}

type Abbr struct {
	buf   bytes.Buffer
	style map[string]string
	text  string
	title string
}

func NewAbbr() *Abbr {
	return &Abbr{
		style: make(map[string]string),
	}
}

func (a *Abbr) AddStyle(k, v string) *Abbr {
	a.style[k] = v
	return a
}

func (a *Abbr) AddStyles(m map[string]string) *Abbr {
	for k, v := range m {
		a.style[k] = v
	}
	return a
}

func (a *Abbr) Style(m map[string]string) *Abbr {
	a.style = m
	return a
}

func (a *Abbr) Text(s string) *Abbr {
	a.text = s
	return a
}

func (a *Abbr) Title(s string) *Abbr {
	a.title = s
	return a
}

func (a *Abbr) Bytes() []byte {
	return cloneBytes(a.buf.Bytes())
}

// IsBodyElement implements BodyElement interface
func (a *Abbr) IsBodyElement() {}

func (a *Abbr) Prepare() {
	a.buf.Reset()
	a.buf.WriteString("<abbr")
	if len(a.style) != 0 {
		parseStyle(&a.buf, a.style)
	}
	if a.title != "" {
		writeAttr(&a.buf, "title", a.title)
	}
	a.buf.WriteString(">" + escapeText(a.text) + "</abbr>")
}

type B struct {
	buf   bytes.Buffer
	style map[string]string
	text  string
}

func NewB() *B {
	return &B{
		style: make(map[string]string),
	}
}

func (b *B) AddStyle(k, v string) *B {
	b.style[k] = v
	return b
}

func (b *B) AddStyles(m map[string]string) *B {
	for k, v := range m {
		b.style[k] = v
	}
	return b
}

func (b *B) Style(m map[string]string) *B {
	b.style = m
	return b
}

func (b *B) Text(s string) *B {
	b.text = s
	return b
}

func (b *B) Bytes() []byte {
	return cloneBytes(b.buf.Bytes())
}

// IsBodyElement implements BodyElement interface
func (b *B) IsBodyElement() {}

func (b *B) Prepare() {
	b.buf.Reset()
	b.buf.WriteString("<b")
	if len(b.style) != 0 {
		parseStyle(&b.buf, b.style)
	}
	b.buf.WriteString(">" + escapeText(b.text) + "</b>")
}

type I struct {
	buf   bytes.Buffer
	style map[string]string
	text  string
}

func NewI() *I {
	return &I{
		style: make(map[string]string),
	}
}

func (i *I) AddStyle(k, v string) *I {
	i.style[k] = v
	return i
}

func (i *I) AddStyles(m map[string]string) *I {
	for k, v := range m {
		i.style[k] = v
	}
	return i
}

func (i *I) Style(m map[string]string) *I {
	i.style = m
	return i
}

func (i *I) Text(s string) *I {
	i.text = s
	return i
}

func (i *I) Bytes() []byte {
	return cloneBytes(i.buf.Bytes())
}

// IsBodyElement implements BodyElement interface
func (i *I) IsBodyElement() {}

func (i *I) Prepare() {
	i.buf.Reset()
	i.buf.WriteString("<i")
	if len(i.style) != 0 {
		parseStyle(&i.buf, i.style)
	}
	i.buf.WriteString(">" + escapeText(i.text) + "</i>")
}

type Q struct {
	buf   bytes.Buffer
	style map[string]string
	text  string
	cite  string
}

func NewQ() *Q {
	return &Q{
		style: make(map[string]string),
	}
}

func (q *Q) AddStyle(k, v string) *Q {
	q.style[k] = v
	return q
}

func (q *Q) AddStyles(m map[string]string) *Q {
	for k, v := range m {
		q.style[k] = v
	}
	return q
}

func (q *Q) Style(m map[string]string) *Q {
	q.style = m
	return q
}

func (q *Q) Text(s string) *Q {
	q.text = s
	return q
}

func (q *Q) Cite(s string) *Q {
	q.cite = s
	return q
}

func (q *Q) Bytes() []byte {
	return cloneBytes(q.buf.Bytes())
}

// IsBodyElement implements BodyElement interface
func (q *Q) IsBodyElement() {}

func (q *Q) Prepare() {
	q.buf.Reset()
	q.buf.WriteString("<q")
	if len(q.style) != 0 {
		parseStyle(&q.buf, q.style)
	}
	if q.cite != "" {
		writeAttr(&q.buf, "cite", q.cite)
	}
	q.buf.WriteString(">" + escapeText(q.text) + "</q>")
}

type S struct {
	buf   bytes.Buffer
	style map[string]string
	text  string
}

func NewS() *S {
	return &S{
		style: make(map[string]string),
	}
}

func (s *S) AddStyle(k, v string) *S {
	s.style[k] = v
	return s
}

func (s *S) AddStyles(m map[string]string) *S {
	for k, v := range m {
		s.style[k] = v
	}
	return s
}

func (s *S) Style(m map[string]string) *S {
	s.style = m
	return s
}

func (s *S) Text(str string) *S {
	s.text = str
	return s
}

func (s *S) Bytes() []byte {
	return cloneBytes(s.buf.Bytes())
}

// IsBodyElement implements BodyElement interface
func (s *S) IsBodyElement() {}

func (s *S) Prepare() {
	s.buf.Reset()
	s.buf.WriteString("<s")
	if len(s.style) != 0 {
		parseStyle(&s.buf, s.style)
	}
	s.buf.WriteString(">" + escapeText(s.text) + "</s>")
}

type U struct {
	buf   bytes.Buffer
	style map[string]string
	text  string
}

func NewU() *U {
	return &U{
		style: make(map[string]string),
	}
}

func (u *U) AddStyle(k, v string) *U {
	u.style[k] = v
	return u
}

func (u *U) AddStyles(m map[string]string) *U {
	for k, v := range m {
		u.style[k] = v
	}
	return u
}

func (u *U) Style(m map[string]string) *U {
	u.style = m
	return u
}

func (u *U) Text(s string) *U {
	u.text = s
	return u
}

func (u *U) Bytes() []byte {
	return cloneBytes(u.buf.Bytes())
}

// IsBodyElement implements BodyElement interface
func (u *U) IsBodyElement() {}

func (u *U) Prepare() {
	u.buf.Reset()
	u.buf.WriteString("<u")
	if len(u.style) != 0 {
		parseStyle(&u.buf, u.style)
	}
	u.buf.WriteString(">" + escapeText(u.text) + "</u>")
}

type Dbi struct {
	buf   bytes.Buffer
	style map[string]string
	text  string
}

func NewDbi() *Dbi {
	return &Dbi{
		style: make(map[string]string),
	}
}

func (d *Dbi) AddStyle(k, v string) *Dbi {
	d.style[k] = v
	return d
}

func (d *Dbi) AddStyles(m map[string]string) *Dbi {
	for k, v := range m {
		d.style[k] = v
	}
	return d
}

func (d *Dbi) Style(m map[string]string) *Dbi {
	d.style = m
	return d
}

func (d *Dbi) Text(s string) *Dbi {
	d.text = s
	return d
}

func (d *Dbi) Bytes() []byte {
	return cloneBytes(d.buf.Bytes())
}

// IsBodyElement implements BodyElement interface
func (d *Dbi) IsBodyElement() {}

func (d *Dbi) Prepare() {
	d.buf.Reset()
	d.buf.WriteString("<dbi")
	if len(d.style) != 0 {
		parseStyle(&d.buf, d.style)
	}
	d.buf.WriteString(">" + escapeText(d.text) + "</dbi>")
}

type Dbo struct {
	buf   bytes.Buffer
	style map[string]string
	text  string
}

func NewDbo() *Dbo {
	return &Dbo{
		style: make(map[string]string),
	}
}

func (d *Dbo) AddStyle(k, v string) *Dbo {
	d.style[k] = v
	return d
}

func (d *Dbo) AddStyles(m map[string]string) *Dbo {
	for k, v := range m {
		d.style[k] = v
	}
	return d
}

func (d *Dbo) Style(m map[string]string) *Dbo {
	d.style = m
	return d
}

func (d *Dbo) Text(s string) *Dbo {
	d.text = s
	return d
}

func (d *Dbo) Bytes() []byte {
	return cloneBytes(d.buf.Bytes())
}

// IsBodyElement implements BodyElement interface
func (d *Dbo) IsBodyElement() {}

func (d *Dbo) Prepare() {
	d.buf.Reset()
	d.buf.WriteString("<dbo")
	if len(d.style) != 0 {
		parseStyle(&d.buf, d.style)
	}
	d.buf.WriteString(">" + escapeText(d.text) + "</dbo>")
}

type Br struct {
	buf   bytes.Buffer
	style map[string]string
}

func NewBr() *Br {
	return &Br{
		style: make(map[string]string),
	}
}

func (br *Br) AddStyle(k, v string) *Br {
	br.style[k] = v
	return br
}

func (br *Br) AddStyles(m map[string]string) *Br {
	for k, v := range m {
		br.style[k] = v
	}
	return br
}

func (br *Br) Style(m map[string]string) *Br {
	br.style = m
	return br
}

func (br *Br) Bytes() []byte {
	return cloneBytes(br.buf.Bytes())
}

// IsBodyElement implements BodyElement interface
func (br *Br) IsBodyElement() {}

func (br *Br) Prepare() {
	br.buf.Reset()
	br.buf.WriteString("<br")
	if len(br.style) != 0 {
		parseStyle(&br.buf, br.style)
	}
	br.buf.WriteByte('>')
}

type Cite struct {
	buf   bytes.Buffer
	style map[string]string
	text  string
}

func NewCite() *Cite {
	return &Cite{
		style: make(map[string]string),
	}
}

func (c *Cite) AddStyle(k, v string) *Cite {
	c.style[k] = v
	return c
}

func (c *Cite) AddStyles(m map[string]string) *Cite {
	for k, v := range m {
		c.style[k] = v
	}
	return c
}

func (c *Cite) Style(m map[string]string) *Cite {
	c.style = m
	return c
}

func (c *Cite) Text(s string) *Cite {
	c.text = s
	return c
}

func (c *Cite) Bytes() []byte {
	return cloneBytes(c.buf.Bytes())
}

// IsBodyElement implements BodyElement interface
func (c *Cite) IsBodyElement() {}

func (c *Cite) Prepare() {
	c.buf.Reset()
	c.buf.WriteString("<cite")
	if len(c.style) != 0 {
		parseStyle(&c.buf, c.style)
	}
	c.buf.WriteString(">" + escapeText(c.text) + "</cite>")
}

type Code struct {
	buf   bytes.Buffer
	style map[string]string
	text  string
}

func NewCode() *Code {
	return &Code{
		style: make(map[string]string),
	}
}

func (c *Code) AddStyle(k, v string) *Code {
	c.style[k] = v
	return c
}

func (c *Code) AddStyles(m map[string]string) *Code {
	for k, v := range m {
		c.style[k] = v
	}
	return c
}

func (c *Code) Style(m map[string]string) *Code {
	c.style = m
	return c
}

func (c *Code) Text(s string) *Code {
	c.text = s
	return c
}

func (c *Code) Bytes() []byte {
	return cloneBytes(c.buf.Bytes())
}

// IsBodyElement implements BodyElement interface
func (c *Code) IsBodyElement() {}

func (c *Code) Prepare() {
	c.buf.Reset()
	c.buf.WriteString("<code")
	if len(c.style) != 0 {
		parseStyle(&c.buf, c.style)
	}
	c.buf.WriteString(">" + escapeText(c.text) + "</code>")
}

type Data struct {
	buf   bytes.Buffer
	style map[string]string
	text  string
	value string
}

func NewData() *Data {
	return &Data{
		style: make(map[string]string),
	}
}

func (d *Data) AddStyle(k, v string) *Data {
	d.style[k] = v
	return d
}

func (d *Data) AddStyles(m map[string]string) *Data {
	for k, v := range m {
		d.style[k] = v
	}
	return d
}

func (d *Data) Style(m map[string]string) *Data {
	d.style = m
	return d
}

func (d *Data) Text(s string) *Data {
	d.text = s
	return d
}

func (d *Data) Value(s string) *Data {
	d.value = s
	return d
}

func (d *Data) Bytes() []byte {
	return cloneBytes(d.buf.Bytes())
}

// IsBodyElement implements BodyElement interface
func (d *Data) IsBodyElement() {}

func (d *Data) Prepare() {
	d.buf.Reset()
	d.buf.WriteString("<data")
	if len(d.style) != 0 {
		parseStyle(&d.buf, d.style)
	}
	if d.value != "" {
		writeAttr(&d.buf, "value", d.value)
	}
	d.buf.WriteString(">" + escapeText(d.text) + "</data>")
}

type Dfn struct {
	buf   bytes.Buffer
	style map[string]string
	text  string
	title string
}

func NewDfn() *Dfn {
	return &Dfn{
		style: make(map[string]string),
	}
}

func (d *Dfn) AddStyle(k, v string) *Dfn {
	d.style[k] = v
	return d
}

func (d *Dfn) AddStyles(m map[string]string) *Dfn {
	for k, v := range m {
		d.style[k] = v
	}
	return d
}

func (d *Dfn) Style(m map[string]string) *Dfn {
	d.style = m
	return d
}

func (d *Dfn) Text(s string) *Dfn {
	d.text = s
	return d
}

func (d *Dfn) Title(s string) *Dfn {
	d.title = s
	return d
}

func (d *Dfn) Bytes() []byte {
	return cloneBytes(d.buf.Bytes())
}

// IsBodyElement implements BodyElement interface
func (d *Dfn) IsBodyElement() {}

func (d *Dfn) Prepare() {
	d.buf.Reset()
	d.buf.WriteString("<dfn")
	if len(d.style) != 0 {
		parseStyle(&d.buf, d.style)
	}
	if d.title != "" {
		writeAttr(&d.buf, "title", d.title)
	}
	d.buf.WriteString(">" + escapeText(d.text) + "</dfn>")
}

type Elem struct {
	buf   bytes.Buffer
	style map[string]string
	text  string
}

func NewElem() *Elem {
	return &Elem{
		style: make(map[string]string),
	}
}

func (e *Elem) AddStyle(k, v string) *Elem {
	e.style[k] = v
	return e
}

func (e *Elem) AddStyles(m map[string]string) *Elem {
	for k, v := range m {
		e.style[k] = v
	}
	return e
}

func (e *Elem) Style(m map[string]string) *Elem {
	e.style = m
	return e
}

func (e *Elem) Text(s string) *Elem {
	e.text = s
	return e
}

func (e *Elem) Bytes() []byte {
	return cloneBytes(e.buf.Bytes())
}

// IsBodyElement implements BodyElement interface
func (e *Elem) IsBodyElement() {}

func (e *Elem) Prepare() {
	e.buf.Reset()
	e.buf.WriteString("<elem")
	if len(e.style) != 0 {
		parseStyle(&e.buf, e.style)
	}
	e.buf.WriteString(">" + escapeText(e.text) + "</elem>")
}

type Mark struct {
	buf   bytes.Buffer
	style map[string]string
	text  string
}

func NewMark() *Mark {
	return &Mark{
		style: make(map[string]string),
	}
}

func (m *Mark) AddStyle(k, v string) *Mark {
	m.style[k] = v
	return m
}

func (m *Mark) AddStyles(ms map[string]string) *Mark {
	for k, v := range ms {
		m.style[k] = v
	}
	return m
}

func (m *Mark) Style(ms map[string]string) *Mark {
	m.style = ms
	return m
}

func (m *Mark) Text(s string) *Mark {
	m.text = s
	return m
}

func (m *Mark) Bytes() []byte {
	return cloneBytes(m.buf.Bytes())
}

// IsBodyElement implements BodyElement interface
func (m *Mark) IsBodyElement() {}

func (m *Mark) Prepare() {
	m.buf.Reset()
	m.buf.WriteString("<mark")
	if len(m.style) != 0 {
		parseStyle(&m.buf, m.style)
	}
	m.buf.WriteString(">" + escapeText(m.text) + "</mark>")
}

type Ruby struct {
	buf      bytes.Buffer
	style    map[string]string
	contents []Element
}

func NewRuby() *Ruby {
	return &Ruby{
		style: make(map[string]string),
	}
}

func (r *Ruby) AddStyle(k, v string) *Ruby {
	r.style[k] = v
	return r
}

func (r *Ruby) AddStyles(m map[string]string) *Ruby {
	for k, v := range m {
		r.style[k] = v
	}
	return r
}

func (r *Ruby) Style(m map[string]string) *Ruby {
	r.style = m
	return r
}

func (r *Ruby) Add(e Element) *Ruby {
	r.contents = appendElement(r.contents, e)
	return r
}

func (r *Ruby) Bytes() []byte {
	return cloneBytes(r.buf.Bytes())
}

// IsBodyElement implements BodyElement interface
func (r *Ruby) IsBodyElement() {}

func (r *Ruby) Prepare() {
	r.buf.Reset()
	r.buf.WriteString("<ruby")
	if len(r.style) != 0 {
		parseStyle(&r.buf, r.style)
	}
	r.buf.WriteByte('>')

	writeElements(&r.buf, r.contents)
	r.buf.WriteString("</ruby>")
}

type Rb struct {
	buf   bytes.Buffer
	style map[string]string
	text  string
}

func NewRb() *Rb {
	return &Rb{
		style: make(map[string]string),
	}
}

func (rb *Rb) AddStyle(k, v string) *Rb {
	rb.style[k] = v
	return rb
}

func (rb *Rb) AddStyles(m map[string]string) *Rb {
	for k, v := range m {
		rb.style[k] = v
	}
	return rb
}

func (rb *Rb) Style(m map[string]string) *Rb {
	rb.style = m
	return rb
}

func (rb *Rb) Text(s string) *Rb {
	rb.text = s
	return rb
}

func (rb *Rb) Bytes() []byte {
	return cloneBytes(rb.buf.Bytes())
}

// IsBodyElement implements BodyElement interface
func (rb *Rb) IsBodyElement() {}

func (rb *Rb) Prepare() {
	rb.buf.Reset()
	rb.buf.WriteString("<rb")
	if len(rb.style) != 0 {
		parseStyle(&rb.buf, rb.style)
	}
	rb.buf.WriteString(">" + escapeText(rb.text) + "</rb>")
}

type Rt struct {
	buf   bytes.Buffer
	style map[string]string
	text  string
}

func NewRt() *Rt {
	return &Rt{
		style: make(map[string]string),
	}
}

func (rt *Rt) AddStyle(k, v string) *Rt {
	rt.style[k] = v
	return rt
}

func (rt *Rt) AddStyles(m map[string]string) *Rt {
	for k, v := range m {
		rt.style[k] = v
	}
	return rt
}

func (rt *Rt) Style(m map[string]string) *Rt {
	rt.style = m
	return rt
}

func (rt *Rt) Text(s string) *Rt {
	rt.text = s
	return rt
}

func (rt *Rt) Bytes() []byte {
	return cloneBytes(rt.buf.Bytes())
}

// IsBodyElement implements BodyElement interface
func (rt *Rt) IsBodyElement() {}

func (rt *Rt) Prepare() {
	rt.buf.Reset()
	rt.buf.WriteString("<rt")
	if len(rt.style) != 0 {
		parseStyle(&rt.buf, rt.style)
	}
	rt.buf.WriteString(">" + escapeText(rt.text) + "</rt>")
}

type Rtc struct {
	buf      bytes.Buffer
	style    map[string]string
	contents []Element
}

func NewRtc() *Rtc {
	return &Rtc{
		style: make(map[string]string),
	}
}

func (rtc *Rtc) AddStyle(k, v string) *Rtc {
	rtc.style[k] = v
	return rtc
}

func (rtc *Rtc) AddStyles(m map[string]string) *Rtc {
	for k, v := range m {
		rtc.style[k] = v
	}
	return rtc
}

func (rtc *Rtc) Style(m map[string]string) *Rtc {
	rtc.style = m
	return rtc
}

func (rtc *Rtc) Add(e Element) *Rtc {
	rtc.contents = appendElement(rtc.contents, e)
	return rtc
}

func (rtc *Rtc) Bytes() []byte {
	return cloneBytes(rtc.buf.Bytes())
}

// IsBodyElement implements BodyElement interface
func (rtc *Rtc) IsBodyElement() {}

func (rtc *Rtc) Prepare() {
	rtc.buf.Reset()
	rtc.buf.WriteString("<rtc")
	if len(rtc.style) != 0 {
		parseStyle(&rtc.buf, rtc.style)
	}
	rtc.buf.WriteByte('>')

	writeElements(&rtc.buf, rtc.contents)
	rtc.buf.WriteString("</rtc>")
}

type Rp struct {
	buf   bytes.Buffer
	style map[string]string
	text  string
}

func NewRp() *Rp {
	return &Rp{
		style: make(map[string]string),
	}
}

func (rp *Rp) AddStyle(k, v string) *Rp {
	rp.style[k] = v
	return rp
}

func (rp *Rp) AddStyles(m map[string]string) *Rp {
	for k, v := range m {
		rp.style[k] = v
	}
	return rp
}

func (rp *Rp) Style(m map[string]string) *Rp {
	rp.style = m
	return rp
}

func (rp *Rp) Text(s string) *Rp {
	rp.text = s
	return rp
}

func (rp *Rp) Bytes() []byte {
	return cloneBytes(rp.buf.Bytes())
}

// IsBodyElement implements BodyElement interface
func (rp *Rp) IsBodyElement() {}

func (rp *Rp) Prepare() {
	rp.buf.Reset()
	rp.buf.WriteString("<rp")
	if len(rp.style) != 0 {
		parseStyle(&rp.buf, rp.style)
	}
	rp.buf.WriteString(">" + escapeText(rp.text) + "</rp>")
}

type Kbd struct {
	buf   bytes.Buffer
	style map[string]string
	text  string
}

func NewKbd() *Kbd {
	return &Kbd{
		style: make(map[string]string),
	}
}

func (k *Kbd) AddStyle(key, v string) *Kbd {
	k.style[key] = v
	return k
}

func (k *Kbd) AddStyles(m map[string]string) *Kbd {
	for key, v := range m {
		k.style[key] = v
	}
	return k
}

func (k *Kbd) Style(m map[string]string) *Kbd {
	k.style = m
	return k
}

func (k *Kbd) Text(s string) *Kbd {
	k.text = s
	return k
}

func (k *Kbd) Bytes() []byte {
	return cloneBytes(k.buf.Bytes())
}

// IsBodyElement implements BodyElement interface
func (k *Kbd) IsBodyElement() {}

func (k *Kbd) Prepare() {
	k.buf.Reset()
	k.buf.WriteString("<kbd")
	if len(k.style) != 0 {
		parseStyle(&k.buf, k.style)
	}
	k.buf.WriteString(">" + escapeText(k.text) + "</kbd>")
}

type Sub struct {
	buf   bytes.Buffer
	style map[string]string
	text  string
}

func NewSub() *Sub {
	return &Sub{
		style: make(map[string]string),
	}
}

func (s *Sub) AddStyle(k, v string) *Sub {
	s.style[k] = v
	return s
}

func (s *Sub) AddStyles(m map[string]string) *Sub {
	for k, v := range m {
		s.style[k] = v
	}
	return s
}

func (s *Sub) Style(m map[string]string) *Sub {
	s.style = m
	return s
}

func (s *Sub) Text(str string) *Sub {
	s.text = str
	return s
}

func (s *Sub) Bytes() []byte {
	return cloneBytes(s.buf.Bytes())
}

// IsBodyElement implements BodyElement interface
func (s *Sub) IsBodyElement() {}

func (s *Sub) Prepare() {
	s.buf.Reset()
	s.buf.WriteString("<sub")
	if len(s.style) != 0 {
		parseStyle(&s.buf, s.style)
	}
	s.buf.WriteString(">" + escapeText(s.text) + "</sub>")
}

type Sup struct {
	buf   bytes.Buffer
	style map[string]string
	text  string
}

func NewSup() *Sup {
	return &Sup{
		style: make(map[string]string),
	}
}

func (s *Sup) AddStyle(k, v string) *Sup {
	s.style[k] = v
	return s
}

func (s *Sup) AddStyles(m map[string]string) *Sup {
	for k, v := range m {
		s.style[k] = v
	}
	return s
}

func (s *Sup) Style(m map[string]string) *Sup {
	s.style = m
	return s
}

func (s *Sup) Text(str string) *Sup {
	s.text = str
	return s
}

func (s *Sup) Bytes() []byte {
	return cloneBytes(s.buf.Bytes())
}

// IsBodyElement implements BodyElement interface
func (s *Sup) IsBodyElement() {}

func (s *Sup) Prepare() {
	s.buf.Reset()
	s.buf.WriteString("<sup")
	if len(s.style) != 0 {
		parseStyle(&s.buf, s.style)
	}
	s.buf.WriteString(">" + escapeText(s.text) + "</sup>")
}

type Samp struct {
	buf   bytes.Buffer
	style map[string]string
	text  string
}

func NewSamp() *Samp {
	return &Samp{
		style: make(map[string]string),
	}
}

func (s *Samp) AddStyle(k, v string) *Samp {
	s.style[k] = v
	return s
}

func (s *Samp) AddStyles(m map[string]string) *Samp {
	for k, v := range m {
		s.style[k] = v
	}
	return s
}

func (s *Samp) Style(m map[string]string) *Samp {
	s.style = m
	return s
}

func (s *Samp) Text(str string) *Samp {
	s.text = str
	return s
}

func (s *Samp) Bytes() []byte {
	return cloneBytes(s.buf.Bytes())
}

// IsBodyElement implements BodyElement interface
func (s *Samp) IsBodyElement() {}

func (s *Samp) Prepare() {
	s.buf.Reset()
	s.buf.WriteString("<samp")
	if len(s.style) != 0 {
		parseStyle(&s.buf, s.style)
	}
	s.buf.WriteString(">" + escapeText(s.text) + "</samp>")
}

type Small struct {
	buf   bytes.Buffer
	style map[string]string
	text  string
}

func NewSmall() *Small {
	return &Small{
		style: make(map[string]string),
	}
}

func (s *Small) AddStyle(k, v string) *Small {
	s.style[k] = v
	return s
}

func (s *Small) AddStyles(m map[string]string) *Small {
	for k, v := range m {
		s.style[k] = v
	}
	return s
}

func (s *Small) Style(m map[string]string) *Small {
	s.style = m
	return s
}

func (s *Small) Text(str string) *Small {
	s.text = str
	return s
}

func (s *Small) Bytes() []byte {
	return cloneBytes(s.buf.Bytes())
}

// IsBodyElement implements BodyElement interface
func (s *Small) IsBodyElement() {}

func (s *Small) Prepare() {
	s.buf.Reset()
	s.buf.WriteString("<small")
	if len(s.style) != 0 {
		parseStyle(&s.buf, s.style)
	}
	s.buf.WriteString(">" + escapeText(s.text) + "</small>")
}

type Span struct {
	buf   bytes.Buffer
	style map[string]string
	text  string
}

func NewSpan() *Span {
	return &Span{
		style: make(map[string]string),
	}
}

func (s *Span) AddStyle(k, v string) *Span {
	s.style[k] = v
	return s
}

func (s *Span) AddStyles(m map[string]string) *Span {
	for k, v := range m {
		s.style[k] = v
	}
	return s
}

func (s *Span) Style(m map[string]string) *Span {
	s.style = m
	return s
}

func (s *Span) Text(str string) *Span {
	s.text = str
	return s
}

func (s *Span) Bytes() []byte {
	return cloneBytes(s.buf.Bytes())
}

// IsBodyElement implements BodyElement interface
func (s *Span) IsBodyElement() {}

func (s *Span) Prepare() {
	s.buf.Reset()
	s.buf.WriteString("<span")
	if len(s.style) != 0 {
		parseStyle(&s.buf, s.style)
	}
	s.buf.WriteString(">" + escapeText(s.text) + "</span>")
}

type Strong struct {
	buf   bytes.Buffer
	style map[string]string
	text  string
}

func NewStrong() *Strong {
	return &Strong{
		style: make(map[string]string),
	}
}

func (s *Strong) AddStyle(k, v string) *Strong {
	s.style[k] = v
	return s
}

func (s *Strong) AddStyles(m map[string]string) *Strong {
	for k, v := range m {
		s.style[k] = v
	}
	return s
}

func (s *Strong) Style(m map[string]string) *Strong {
	s.style = m
	return s
}

func (s *Strong) Text(str string) *Strong {
	s.text = str
	return s
}

func (s *Strong) Bytes() []byte {
	return cloneBytes(s.buf.Bytes())
}

// IsBodyElement implements BodyElement interface
func (s *Strong) IsBodyElement() {}

func (s *Strong) Prepare() {
	s.buf.Reset()
	s.buf.WriteString("<strong")
	if len(s.style) != 0 {
		parseStyle(&s.buf, s.style)
	}
	s.buf.WriteString(">" + escapeText(s.text) + "</strong>")
}

type Time struct {
	buf      bytes.Buffer
	style    map[string]string
	text     string
	datetime string
}

func NewTime() *Time {
	return &Time{
		style: make(map[string]string),
	}
}

func (t *Time) AddStyle(k, v string) *Time {
	t.style[k] = v
	return t
}

func (t *Time) AddStyles(m map[string]string) *Time {
	for k, v := range m {
		t.style[k] = v
	}
	return t
}

func (t *Time) Style(m map[string]string) *Time {
	t.style = m
	return t
}

func (t *Time) Text(str string) *Time {
	t.text = str
	return t
}

func (t *Time) Datetime(dt string) *Time {
	if !isValidDatetime(dt) {
		log.Fatal("Invalid datetime value: " + dt)
	}
	t.datetime = dt
	return t
}

func isValidDatetime(dt string) bool {
	if dt == "" {
		return true
	}

	datetimeFormats := []string{
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

	weekPattern := regexp.MustCompile(`^2006-W\d{2}$`)
	if weekPattern.MatchString(dt) {
		return true
	}

	monthPattern := regexp.MustCompile(`^2006-\d{2}$`)
	if monthPattern.MatchString(dt) {
		return true
	}

	durationPattern := regexp.MustCompile(`^P(?:\d+Y)?(?:\d+M)?(?:\d+D)?(?:T(?:\d+H)?(?:\d+M)?(?:\d+(?:\.\d+)?S)?)?$`)
	if durationPattern.MatchString(dt) {
		return true
	}

	for _, format := range datetimeFormats {
		if _, err := time.Parse(format, dt); err == nil {
			return true
		}
	}

	return false
}

func (t *Time) Bytes() []byte {
	return cloneBytes(t.buf.Bytes())
}

// IsBodyElement implements BodyElement interface
func (t *Time) IsBodyElement() {}

func (t *Time) Prepare() {
	t.buf.Reset()
	t.buf.WriteString("<time")
	if t.datetime != "" {
		writeAttr(&t.buf, "datetime", t.datetime)
	}
	if len(t.style) != 0 {
		parseStyle(&t.buf, t.style)
	}
	t.buf.WriteString(">" + escapeText(t.text) + "</time>")
}

type Var struct {
	buf   bytes.Buffer
	style map[string]string
	text  string
}

func NewVar() *Var {
	return &Var{
		style: make(map[string]string),
	}
}

func (v *Var) AddStyle(k, val string) *Var {
	v.style[k] = val
	return v
}

func (v *Var) AddStyles(m map[string]string) *Var {
	for k, val := range m {
		v.style[k] = val
	}
	return v
}

func (v *Var) Style(m map[string]string) *Var {
	v.style = m
	return v
}

func (v *Var) Text(str string) *Var {
	v.text = str
	return v
}

func (v *Var) Bytes() []byte {
	return cloneBytes(v.buf.Bytes())
}

// IsBodyElement implements BodyElement interface
func (v *Var) IsBodyElement() {}

func (v *Var) Prepare() {
	v.buf.Reset()
	v.buf.WriteString("<var")
	if len(v.style) != 0 {
		parseStyle(&v.buf, v.style)
	}
	v.buf.WriteString(">" + escapeText(v.text) + "</var>")
}

type Wbr struct {
	buf   bytes.Buffer
	style map[string]string
}

func NewWbr() *Wbr {
	return &Wbr{
		style: make(map[string]string),
	}
}

func (w *Wbr) AddStyle(k, v string) *Wbr {
	w.style[k] = v
	return w
}

func (w *Wbr) AddStyles(m map[string]string) *Wbr {
	for k, v := range m {
		w.style[k] = v
	}
	return w
}

func (w *Wbr) Style(m map[string]string) *Wbr {
	w.style = m
	return w
}

func (w *Wbr) Bytes() []byte {
	return cloneBytes(w.buf.Bytes())
}

// IsBodyElement implements BodyElement interface
func (w *Wbr) IsBodyElement() {}

func (w *Wbr) Prepare() {
	w.buf.Reset()
	w.buf.WriteString("<wbr")
	if len(w.style) != 0 {
		parseStyle(&w.buf, w.style)
	}
	w.buf.WriteByte('>')
}
