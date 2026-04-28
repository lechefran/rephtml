package rephtml

import (
	"bytes"
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
	buf        bytes.Buffer
	style      StyleMap
	link, text string
}

// NewAnchor creates a new Anchor component.
func NewAnchor() *Anchor {
	return &Anchor{
		style: make(StyleMap),
	}
}

// AddStyle adds one inline CSS declaration to the Anchor component.
func (a *Anchor) AddStyle(k, v string) *Anchor {
	a.style[k] = v
	return a
}

// Style replaces the inline CSS declarations on the Anchor component.
func (a *Anchor) Style(m StyleMap) *Anchor {
	a.style = cloneStyleMap(m)
	return a
}

// Text sets or appends text content on the Anchor component.
func (a *Anchor) Text(s string) *Anchor {
	a.text = s
	return a
}

// Link sets the link value on the Anchor component.
func (a *Anchor) Link(s string) *Anchor {
	a.link = s
	return a
}

// Bytes returns a defensive copy of the rendered Anchor bytes.
func (a *Anchor) Bytes() []byte {
	return cloneBytes(a.buf.Bytes())
}

// Render returns freshly prepared Anchor HTML bytes.
func (a *Anchor) Render() []byte {
	return renderPrepared(a)
}

// HTML returns freshly prepared Anchor HTML as a string.
func (a *Anchor) HTML() string {
	return htmlPrepared(a)
}

// String returns freshly prepared Anchor HTML as a string.
func (a *Anchor) String() string {
	return a.HTML()
}

// IsBodyElement implements BodyElement interface
func (a *Anchor) IsBodyElement() {}

// Prepare renders the Anchor component into its internal buffer.
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

// Abbr represents the Abbr component or supporting type.
type Abbr struct {
	buf   bytes.Buffer
	style StyleMap
	text  string
	title string
}

// NewAbbr creates a new Abbr component.
func NewAbbr() *Abbr {
	return &Abbr{
		style: make(StyleMap),
	}
}

// AddStyle adds one inline CSS declaration to the Abbr component.
func (a *Abbr) AddStyle(k, v string) *Abbr {
	a.style[k] = v
	return a
}

// AddStyles adds multiple inline CSS declarations to the Abbr component.
func (a *Abbr) AddStyles(m StyleMap) *Abbr {
	for k, v := range m {
		a.style[k] = v
	}
	return a
}

// Style replaces the inline CSS declarations on the Abbr component.
func (a *Abbr) Style(m StyleMap) *Abbr {
	a.style = cloneStyleMap(m)
	return a
}

// Text sets or appends text content on the Abbr component.
func (a *Abbr) Text(s string) *Abbr {
	a.text = s
	return a
}

// Title sets the title value on the Abbr component.
func (a *Abbr) Title(s string) *Abbr {
	a.title = s
	return a
}

// Bytes returns a defensive copy of the rendered Abbr bytes.
func (a *Abbr) Bytes() []byte {
	return cloneBytes(a.buf.Bytes())
}

// Render returns freshly prepared Abbr HTML bytes.
func (a *Abbr) Render() []byte {
	return renderPrepared(a)
}

// HTML returns freshly prepared Abbr HTML as a string.
func (a *Abbr) HTML() string {
	return htmlPrepared(a)
}

// String returns freshly prepared Abbr HTML as a string.
func (a *Abbr) String() string {
	return a.HTML()
}

// IsBodyElement implements BodyElement interface
func (a *Abbr) IsBodyElement() {}

// Prepare renders the Abbr component into its internal buffer.
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

// B represents the B component or supporting type.
type B struct {
	buf   bytes.Buffer
	style StyleMap
	text  string
}

// NewB creates a new B component.
func NewB() *B {
	return &B{
		style: make(StyleMap),
	}
}

// AddStyle adds one inline CSS declaration to the B component.
func (b *B) AddStyle(k, v string) *B {
	b.style[k] = v
	return b
}

// AddStyles adds multiple inline CSS declarations to the B component.
func (b *B) AddStyles(m StyleMap) *B {
	for k, v := range m {
		b.style[k] = v
	}
	return b
}

// Style replaces the inline CSS declarations on the B component.
func (b *B) Style(m StyleMap) *B {
	b.style = cloneStyleMap(m)
	return b
}

// Text sets or appends text content on the B component.
func (b *B) Text(s string) *B {
	b.text = s
	return b
}

// Bytes returns a defensive copy of the rendered B bytes.
func (b *B) Bytes() []byte {
	return cloneBytes(b.buf.Bytes())
}

// Render returns freshly prepared B HTML bytes.
func (b *B) Render() []byte {
	return renderPrepared(b)
}

// HTML returns freshly prepared B HTML as a string.
func (b *B) HTML() string {
	return htmlPrepared(b)
}

// String returns freshly prepared B HTML as a string.
func (b *B) String() string {
	return b.HTML()
}

// IsBodyElement implements BodyElement interface
func (b *B) IsBodyElement() {}

// Prepare renders the B component into its internal buffer.
func (b *B) Prepare() {
	b.buf.Reset()
	b.buf.WriteString("<b")
	if len(b.style) != 0 {
		parseStyle(&b.buf, b.style)
	}
	b.buf.WriteString(">" + escapeText(b.text) + "</b>")
}

// I represents the I component or supporting type.
type I struct {
	buf   bytes.Buffer
	style StyleMap
	text  string
}

// NewI creates a new I component.
func NewI() *I {
	return &I{
		style: make(StyleMap),
	}
}

// AddStyle adds one inline CSS declaration to the I component.
func (i *I) AddStyle(k, v string) *I {
	i.style[k] = v
	return i
}

// AddStyles adds multiple inline CSS declarations to the I component.
func (i *I) AddStyles(m StyleMap) *I {
	for k, v := range m {
		i.style[k] = v
	}
	return i
}

// Style replaces the inline CSS declarations on the I component.
func (i *I) Style(m StyleMap) *I {
	i.style = cloneStyleMap(m)
	return i
}

// Text sets or appends text content on the I component.
func (i *I) Text(s string) *I {
	i.text = s
	return i
}

// Bytes returns a defensive copy of the rendered I bytes.
func (i *I) Bytes() []byte {
	return cloneBytes(i.buf.Bytes())
}

// Render returns freshly prepared I HTML bytes.
func (i *I) Render() []byte {
	return renderPrepared(i)
}

// HTML returns freshly prepared I HTML as a string.
func (i *I) HTML() string {
	return htmlPrepared(i)
}

// String returns freshly prepared I HTML as a string.
func (i *I) String() string {
	return i.HTML()
}

// IsBodyElement implements BodyElement interface
func (i *I) IsBodyElement() {}

// Prepare renders the I component into its internal buffer.
func (i *I) Prepare() {
	i.buf.Reset()
	i.buf.WriteString("<i")
	if len(i.style) != 0 {
		parseStyle(&i.buf, i.style)
	}
	i.buf.WriteString(">" + escapeText(i.text) + "</i>")
}

// Q represents the Q component or supporting type.
type Q struct {
	buf   bytes.Buffer
	style StyleMap
	text  string
	cite  string
}

// NewQ creates a new Q component.
func NewQ() *Q {
	return &Q{
		style: make(StyleMap),
	}
}

// AddStyle adds one inline CSS declaration to the Q component.
func (q *Q) AddStyle(k, v string) *Q {
	q.style[k] = v
	return q
}

// AddStyles adds multiple inline CSS declarations to the Q component.
func (q *Q) AddStyles(m StyleMap) *Q {
	for k, v := range m {
		q.style[k] = v
	}
	return q
}

// Style replaces the inline CSS declarations on the Q component.
func (q *Q) Style(m StyleMap) *Q {
	q.style = cloneStyleMap(m)
	return q
}

// Text sets or appends text content on the Q component.
func (q *Q) Text(s string) *Q {
	q.text = s
	return q
}

// Cite sets the cite value on the Q component.
func (q *Q) Cite(s string) *Q {
	q.cite = s
	return q
}

// Bytes returns a defensive copy of the rendered Q bytes.
func (q *Q) Bytes() []byte {
	return cloneBytes(q.buf.Bytes())
}

// Render returns freshly prepared Q HTML bytes.
func (q *Q) Render() []byte {
	return renderPrepared(q)
}

// HTML returns freshly prepared Q HTML as a string.
func (q *Q) HTML() string {
	return htmlPrepared(q)
}

// String returns freshly prepared Q HTML as a string.
func (q *Q) String() string {
	return q.HTML()
}

// IsBodyElement implements BodyElement interface
func (q *Q) IsBodyElement() {}

// Prepare renders the Q component into its internal buffer.
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

// S represents the S component or supporting type.
type S struct {
	buf   bytes.Buffer
	style StyleMap
	text  string
}

// NewS creates a new S component.
func NewS() *S {
	return &S{
		style: make(StyleMap),
	}
}

// AddStyle adds one inline CSS declaration to the S component.
func (s *S) AddStyle(k, v string) *S {
	s.style[k] = v
	return s
}

// AddStyles adds multiple inline CSS declarations to the S component.
func (s *S) AddStyles(m StyleMap) *S {
	for k, v := range m {
		s.style[k] = v
	}
	return s
}

// Style replaces the inline CSS declarations on the S component.
func (s *S) Style(m StyleMap) *S {
	s.style = cloneStyleMap(m)
	return s
}

// Text sets or appends text content on the S component.
func (s *S) Text(str string) *S {
	s.text = str
	return s
}

// Bytes returns a defensive copy of the rendered S bytes.
func (s *S) Bytes() []byte {
	return cloneBytes(s.buf.Bytes())
}

// Render returns freshly prepared S HTML bytes.
func (s *S) Render() []byte {
	return renderPrepared(s)
}

// HTML returns freshly prepared S HTML as a string.
func (s *S) HTML() string {
	return htmlPrepared(s)
}

// String returns freshly prepared S HTML as a string.
func (s *S) String() string {
	return s.HTML()
}

// IsBodyElement implements BodyElement interface
func (s *S) IsBodyElement() {}

// Prepare renders the S component into its internal buffer.
func (s *S) Prepare() {
	s.buf.Reset()
	s.buf.WriteString("<s")
	if len(s.style) != 0 {
		parseStyle(&s.buf, s.style)
	}
	s.buf.WriteString(">" + escapeText(s.text) + "</s>")
}

// U represents the U component or supporting type.
type U struct {
	buf   bytes.Buffer
	style StyleMap
	text  string
}

// NewU creates a new U component.
func NewU() *U {
	return &U{
		style: make(StyleMap),
	}
}

// AddStyle adds one inline CSS declaration to the U component.
func (u *U) AddStyle(k, v string) *U {
	u.style[k] = v
	return u
}

// AddStyles adds multiple inline CSS declarations to the U component.
func (u *U) AddStyles(m StyleMap) *U {
	for k, v := range m {
		u.style[k] = v
	}
	return u
}

// Style replaces the inline CSS declarations on the U component.
func (u *U) Style(m StyleMap) *U {
	u.style = cloneStyleMap(m)
	return u
}

// Text sets or appends text content on the U component.
func (u *U) Text(s string) *U {
	u.text = s
	return u
}

// Bytes returns a defensive copy of the rendered U bytes.
func (u *U) Bytes() []byte {
	return cloneBytes(u.buf.Bytes())
}

// Render returns freshly prepared U HTML bytes.
func (u *U) Render() []byte {
	return renderPrepared(u)
}

// HTML returns freshly prepared U HTML as a string.
func (u *U) HTML() string {
	return htmlPrepared(u)
}

// String returns freshly prepared U HTML as a string.
func (u *U) String() string {
	return u.HTML()
}

// IsBodyElement implements BodyElement interface
func (u *U) IsBodyElement() {}

// Prepare renders the U component into its internal buffer.
func (u *U) Prepare() {
	u.buf.Reset()
	u.buf.WriteString("<u")
	if len(u.style) != 0 {
		parseStyle(&u.buf, u.style)
	}
	u.buf.WriteString(">" + escapeText(u.text) + "</u>")
}

// Bdi represents the bdi component or supporting type.
type Bdi struct {
	buf   bytes.Buffer
	style StyleMap
	text  string
}

// NewBdi creates a new bdi component.
func NewBdi() *Bdi {
	return &Bdi{
		style: make(StyleMap),
	}
}

// AddStyle adds one inline CSS declaration to the Bdi component.
func (d *Bdi) AddStyle(k, v string) *Bdi {
	d.style[k] = v
	return d
}

// AddStyles adds multiple inline CSS declarations to the Bdi component.
func (d *Bdi) AddStyles(m StyleMap) *Bdi {
	for k, v := range m {
		d.style[k] = v
	}
	return d
}

// Style replaces the inline CSS declarations on the Bdi component.
func (d *Bdi) Style(m StyleMap) *Bdi {
	d.style = cloneStyleMap(m)
	return d
}

// Text sets or appends text content on the Bdi component.
func (d *Bdi) Text(s string) *Bdi {
	d.text = s
	return d
}

// Bytes returns a defensive copy of the rendered Bdi bytes.
func (d *Bdi) Bytes() []byte {
	return cloneBytes(d.buf.Bytes())
}

// Render returns freshly prepared Bdi HTML bytes.
func (b *Bdi) Render() []byte {
	return renderPrepared(b)
}

// HTML returns freshly prepared Bdi HTML as a string.
func (b *Bdi) HTML() string {
	return htmlPrepared(b)
}

// String returns freshly prepared Bdi HTML as a string.
func (b *Bdi) String() string {
	return b.HTML()
}

// IsBodyElement implements BodyElement interface
func (d *Bdi) IsBodyElement() {}

// Prepare renders the Bdi component into its internal buffer.
func (d *Bdi) Prepare() {
	d.buf.Reset()
	d.buf.WriteString("<bdi")
	if len(d.style) != 0 {
		parseStyle(&d.buf, d.style)
	}
	d.buf.WriteString(">" + escapeText(d.text) + "</bdi>")
}

// Bdo represents the bdo component or supporting type.
type Bdo struct {
	buf   bytes.Buffer
	style StyleMap
	text  string
}

// NewBdo creates a new bdo component.
func NewBdo() *Bdo {
	return &Bdo{
		style: make(StyleMap),
	}
}

// AddStyle adds one inline CSS declaration to the Bdo component.
func (d *Bdo) AddStyle(k, v string) *Bdo {
	d.style[k] = v
	return d
}

// AddStyles adds multiple inline CSS declarations to the Bdo component.
func (d *Bdo) AddStyles(m StyleMap) *Bdo {
	for k, v := range m {
		d.style[k] = v
	}
	return d
}

// Style replaces the inline CSS declarations on the Bdo component.
func (d *Bdo) Style(m StyleMap) *Bdo {
	d.style = cloneStyleMap(m)
	return d
}

// Text sets or appends text content on the Bdo component.
func (d *Bdo) Text(s string) *Bdo {
	d.text = s
	return d
}

// Bytes returns a defensive copy of the rendered Bdo bytes.
func (d *Bdo) Bytes() []byte {
	return cloneBytes(d.buf.Bytes())
}

// Render returns freshly prepared Bdo HTML bytes.
func (b *Bdo) Render() []byte {
	return renderPrepared(b)
}

// HTML returns freshly prepared Bdo HTML as a string.
func (b *Bdo) HTML() string {
	return htmlPrepared(b)
}

// String returns freshly prepared Bdo HTML as a string.
func (b *Bdo) String() string {
	return b.HTML()
}

// IsBodyElement implements BodyElement interface
func (d *Bdo) IsBodyElement() {}

// Prepare renders the Bdo component into its internal buffer.
func (d *Bdo) Prepare() {
	d.buf.Reset()
	d.buf.WriteString("<bdo")
	if len(d.style) != 0 {
		parseStyle(&d.buf, d.style)
	}
	d.buf.WriteString(">" + escapeText(d.text) + "</bdo>")
}

// Br represents the Br component or supporting type.
type Br struct {
	buf   bytes.Buffer
	style StyleMap
}

// NewBr creates a new Br component.
func NewBr() *Br {
	return &Br{
		style: make(StyleMap),
	}
}

// AddStyle adds one inline CSS declaration to the Br component.
func (br *Br) AddStyle(k, v string) *Br {
	br.style[k] = v
	return br
}

// AddStyles adds multiple inline CSS declarations to the Br component.
func (br *Br) AddStyles(m StyleMap) *Br {
	for k, v := range m {
		br.style[k] = v
	}
	return br
}

// Style replaces the inline CSS declarations on the Br component.
func (br *Br) Style(m StyleMap) *Br {
	br.style = cloneStyleMap(m)
	return br
}

// Bytes returns a defensive copy of the rendered Br bytes.
func (br *Br) Bytes() []byte {
	return cloneBytes(br.buf.Bytes())
}

// Render returns freshly prepared Br HTML bytes.
func (b *Br) Render() []byte {
	return renderPrepared(b)
}

// HTML returns freshly prepared Br HTML as a string.
func (b *Br) HTML() string {
	return htmlPrepared(b)
}

// String returns freshly prepared Br HTML as a string.
func (b *Br) String() string {
	return b.HTML()
}

// IsBodyElement implements BodyElement interface
func (br *Br) IsBodyElement() {}

// Prepare renders the Br component into its internal buffer.
func (br *Br) Prepare() {
	br.buf.Reset()
	br.buf.WriteString("<br")
	if len(br.style) != 0 {
		parseStyle(&br.buf, br.style)
	}
	br.buf.WriteByte('>')
}

// Cite represents the Cite component or supporting type.
type Cite struct {
	buf   bytes.Buffer
	style StyleMap
	text  string
}

// NewCite creates a new Cite component.
func NewCite() *Cite {
	return &Cite{
		style: make(StyleMap),
	}
}

// AddStyle adds one inline CSS declaration to the Cite component.
func (c *Cite) AddStyle(k, v string) *Cite {
	c.style[k] = v
	return c
}

// AddStyles adds multiple inline CSS declarations to the Cite component.
func (c *Cite) AddStyles(m StyleMap) *Cite {
	for k, v := range m {
		c.style[k] = v
	}
	return c
}

// Style replaces the inline CSS declarations on the Cite component.
func (c *Cite) Style(m StyleMap) *Cite {
	c.style = cloneStyleMap(m)
	return c
}

// Text sets or appends text content on the Cite component.
func (c *Cite) Text(s string) *Cite {
	c.text = s
	return c
}

// Bytes returns a defensive copy of the rendered Cite bytes.
func (c *Cite) Bytes() []byte {
	return cloneBytes(c.buf.Bytes())
}

// Render returns freshly prepared Cite HTML bytes.
func (c *Cite) Render() []byte {
	return renderPrepared(c)
}

// HTML returns freshly prepared Cite HTML as a string.
func (c *Cite) HTML() string {
	return htmlPrepared(c)
}

// String returns freshly prepared Cite HTML as a string.
func (c *Cite) String() string {
	return c.HTML()
}

// IsBodyElement implements BodyElement interface
func (c *Cite) IsBodyElement() {}

// Prepare renders the Cite component into its internal buffer.
func (c *Cite) Prepare() {
	c.buf.Reset()
	c.buf.WriteString("<cite")
	if len(c.style) != 0 {
		parseStyle(&c.buf, c.style)
	}
	c.buf.WriteString(">" + escapeText(c.text) + "</cite>")
}

// Code represents the Code component or supporting type.
type Code struct {
	buf   bytes.Buffer
	style StyleMap
	text  string
}

// NewCode creates a new Code component.
func NewCode() *Code {
	return &Code{
		style: make(StyleMap),
	}
}

// AddStyle adds one inline CSS declaration to the Code component.
func (c *Code) AddStyle(k, v string) *Code {
	c.style[k] = v
	return c
}

// AddStyles adds multiple inline CSS declarations to the Code component.
func (c *Code) AddStyles(m StyleMap) *Code {
	for k, v := range m {
		c.style[k] = v
	}
	return c
}

// Style replaces the inline CSS declarations on the Code component.
func (c *Code) Style(m StyleMap) *Code {
	c.style = cloneStyleMap(m)
	return c
}

// Text sets or appends text content on the Code component.
func (c *Code) Text(s string) *Code {
	c.text = s
	return c
}

// Bytes returns a defensive copy of the rendered Code bytes.
func (c *Code) Bytes() []byte {
	return cloneBytes(c.buf.Bytes())
}

// Render returns freshly prepared Code HTML bytes.
func (c *Code) Render() []byte {
	return renderPrepared(c)
}

// HTML returns freshly prepared Code HTML as a string.
func (c *Code) HTML() string {
	return htmlPrepared(c)
}

// String returns freshly prepared Code HTML as a string.
func (c *Code) String() string {
	return c.HTML()
}

// IsBodyElement implements BodyElement interface
func (c *Code) IsBodyElement() {}

// Prepare renders the Code component into its internal buffer.
func (c *Code) Prepare() {
	c.buf.Reset()
	c.buf.WriteString("<code")
	if len(c.style) != 0 {
		parseStyle(&c.buf, c.style)
	}
	c.buf.WriteString(">" + escapeText(c.text) + "</code>")
}

// Data represents the Data component or supporting type.
type Data struct {
	buf   bytes.Buffer
	style StyleMap
	text  string
	value string
}

// NewData creates a new Data component.
func NewData() *Data {
	return &Data{
		style: make(StyleMap),
	}
}

// AddStyle adds one inline CSS declaration to the Data component.
func (d *Data) AddStyle(k, v string) *Data {
	d.style[k] = v
	return d
}

// AddStyles adds multiple inline CSS declarations to the Data component.
func (d *Data) AddStyles(m StyleMap) *Data {
	for k, v := range m {
		d.style[k] = v
	}
	return d
}

// Style replaces the inline CSS declarations on the Data component.
func (d *Data) Style(m StyleMap) *Data {
	d.style = cloneStyleMap(m)
	return d
}

// Text sets or appends text content on the Data component.
func (d *Data) Text(s string) *Data {
	d.text = s
	return d
}

// Value sets the value value on the Data component.
func (d *Data) Value(s string) *Data {
	d.value = s
	return d
}

// Bytes returns a defensive copy of the rendered Data bytes.
func (d *Data) Bytes() []byte {
	return cloneBytes(d.buf.Bytes())
}

// Render returns freshly prepared Data HTML bytes.
func (d *Data) Render() []byte {
	return renderPrepared(d)
}

// HTML returns freshly prepared Data HTML as a string.
func (d *Data) HTML() string {
	return htmlPrepared(d)
}

// String returns freshly prepared Data HTML as a string.
func (d *Data) String() string {
	return d.HTML()
}

// IsBodyElement implements BodyElement interface
func (d *Data) IsBodyElement() {}

// Prepare renders the Data component into its internal buffer.
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

// Dfn represents the Dfn component or supporting type.
type Dfn struct {
	buf   bytes.Buffer
	style StyleMap
	text  string
	title string
}

// NewDfn creates a new Dfn component.
func NewDfn() *Dfn {
	return &Dfn{
		style: make(StyleMap),
	}
}

// AddStyle adds one inline CSS declaration to the Dfn component.
func (d *Dfn) AddStyle(k, v string) *Dfn {
	d.style[k] = v
	return d
}

// AddStyles adds multiple inline CSS declarations to the Dfn component.
func (d *Dfn) AddStyles(m StyleMap) *Dfn {
	for k, v := range m {
		d.style[k] = v
	}
	return d
}

// Style replaces the inline CSS declarations on the Dfn component.
func (d *Dfn) Style(m StyleMap) *Dfn {
	d.style = cloneStyleMap(m)
	return d
}

// Text sets or appends text content on the Dfn component.
func (d *Dfn) Text(s string) *Dfn {
	d.text = s
	return d
}

// Title sets the title value on the Dfn component.
func (d *Dfn) Title(s string) *Dfn {
	d.title = s
	return d
}

// Bytes returns a defensive copy of the rendered Dfn bytes.
func (d *Dfn) Bytes() []byte {
	return cloneBytes(d.buf.Bytes())
}

// Render returns freshly prepared Dfn HTML bytes.
func (d *Dfn) Render() []byte {
	return renderPrepared(d)
}

// HTML returns freshly prepared Dfn HTML as a string.
func (d *Dfn) HTML() string {
	return htmlPrepared(d)
}

// String returns freshly prepared Dfn HTML as a string.
func (d *Dfn) String() string {
	return d.HTML()
}

// IsBodyElement implements BodyElement interface
func (d *Dfn) IsBodyElement() {}

// Prepare renders the Dfn component into its internal buffer.
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

// Em represents the em component or supporting type.
type Em struct {
	buf   bytes.Buffer
	style StyleMap
	text  string
}

// NewEm creates a new em component.
func NewEm() *Em {
	return &Em{
		style: make(StyleMap),
	}
}

// AddStyle adds one inline CSS declaration to the Em component.
func (e *Em) AddStyle(k, v string) *Em {
	e.style[k] = v
	return e
}

// AddStyles adds multiple inline CSS declarations to the Em component.
func (e *Em) AddStyles(m StyleMap) *Em {
	for k, v := range m {
		e.style[k] = v
	}
	return e
}

// Style replaces the inline CSS declarations on the Em component.
func (e *Em) Style(m StyleMap) *Em {
	e.style = cloneStyleMap(m)
	return e
}

// Text sets or appends text content on the Em component.
func (e *Em) Text(s string) *Em {
	e.text = s
	return e
}

// Bytes returns a defensive copy of the rendered Em bytes.
func (e *Em) Bytes() []byte {
	return cloneBytes(e.buf.Bytes())
}

// Render returns freshly prepared Em HTML bytes.
func (e *Em) Render() []byte {
	return renderPrepared(e)
}

// HTML returns freshly prepared Em HTML as a string.
func (e *Em) HTML() string {
	return htmlPrepared(e)
}

// String returns freshly prepared Em HTML as a string.
func (e *Em) String() string {
	return e.HTML()
}

// IsBodyElement implements BodyElement interface
func (e *Em) IsBodyElement() {}

// Prepare renders the Em component into its internal buffer.
func (e *Em) Prepare() {
	e.buf.Reset()
	e.buf.WriteString("<em")
	if len(e.style) != 0 {
		parseStyle(&e.buf, e.style)
	}
	e.buf.WriteString(">" + escapeText(e.text) + "</em>")
}

// Mark represents the Mark component or supporting type.
type Mark struct {
	buf   bytes.Buffer
	style StyleMap
	text  string
}

// NewMark creates a new Mark component.
func NewMark() *Mark {
	return &Mark{
		style: make(StyleMap),
	}
}

// AddStyle adds one inline CSS declaration to the Mark component.
func (m *Mark) AddStyle(k, v string) *Mark {
	m.style[k] = v
	return m
}

// AddStyles adds multiple inline CSS declarations to the Mark component.
func (m *Mark) AddStyles(ms StyleMap) *Mark {
	for k, v := range ms {
		m.style[k] = v
	}
	return m
}

// Style replaces the inline CSS declarations on the Mark component.
func (m *Mark) Style(ms StyleMap) *Mark {
	m.style = cloneStyleMap(ms)
	return m
}

// Text sets or appends text content on the Mark component.
func (m *Mark) Text(s string) *Mark {
	m.text = s
	return m
}

// Bytes returns a defensive copy of the rendered Mark bytes.
func (m *Mark) Bytes() []byte {
	return cloneBytes(m.buf.Bytes())
}

// Render returns freshly prepared Mark HTML bytes.
func (m *Mark) Render() []byte {
	return renderPrepared(m)
}

// HTML returns freshly prepared Mark HTML as a string.
func (m *Mark) HTML() string {
	return htmlPrepared(m)
}

// String returns freshly prepared Mark HTML as a string.
func (m *Mark) String() string {
	return m.HTML()
}

// IsBodyElement implements BodyElement interface
func (m *Mark) IsBodyElement() {}

// Prepare renders the Mark component into its internal buffer.
func (m *Mark) Prepare() {
	m.buf.Reset()
	m.buf.WriteString("<mark")
	if len(m.style) != 0 {
		parseStyle(&m.buf, m.style)
	}
	m.buf.WriteString(">" + escapeText(m.text) + "</mark>")
}

// Ruby represents the Ruby component or supporting type.
type Ruby struct {
	buf      bytes.Buffer
	style    StyleMap
	contents []Element
}

// NewRuby creates a new Ruby component.
func NewRuby() *Ruby {
	return &Ruby{
		style: make(StyleMap),
	}
}

// AddStyle adds one inline CSS declaration to the Ruby component.
func (r *Ruby) AddStyle(k, v string) *Ruby {
	r.style[k] = v
	return r
}

// AddStyles adds multiple inline CSS declarations to the Ruby component.
func (r *Ruby) AddStyles(m StyleMap) *Ruby {
	for k, v := range m {
		r.style[k] = v
	}
	return r
}

// Style replaces the inline CSS declarations on the Ruby component.
func (r *Ruby) Style(m StyleMap) *Ruby {
	r.style = cloneStyleMap(m)
	return r
}

// Add appends child content to the Ruby component.
func (r *Ruby) Add(e Element) *Ruby {
	r.contents = appendElement(r.contents, e)
	return r
}

// Bytes returns a defensive copy of the rendered Ruby bytes.
func (r *Ruby) Bytes() []byte {
	return cloneBytes(r.buf.Bytes())
}

// Render returns freshly prepared Ruby HTML bytes.
func (r *Ruby) Render() []byte {
	return renderPrepared(r)
}

// HTML returns freshly prepared Ruby HTML as a string.
func (r *Ruby) HTML() string {
	return htmlPrepared(r)
}

// String returns freshly prepared Ruby HTML as a string.
func (r *Ruby) String() string {
	return r.HTML()
}

// IsBodyElement implements BodyElement interface
func (r *Ruby) IsBodyElement() {}

// Prepare renders the Ruby component into its internal buffer.
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

// Rb represents the Rb component or supporting type.
type Rb struct {
	buf   bytes.Buffer
	style StyleMap
	text  string
}

// NewRb creates a new Rb component.
func NewRb() *Rb {
	return &Rb{
		style: make(StyleMap),
	}
}

// AddStyle adds one inline CSS declaration to the Rb component.
func (rb *Rb) AddStyle(k, v string) *Rb {
	rb.style[k] = v
	return rb
}

// AddStyles adds multiple inline CSS declarations to the Rb component.
func (rb *Rb) AddStyles(m StyleMap) *Rb {
	for k, v := range m {
		rb.style[k] = v
	}
	return rb
}

// Style replaces the inline CSS declarations on the Rb component.
func (rb *Rb) Style(m StyleMap) *Rb {
	rb.style = cloneStyleMap(m)
	return rb
}

// Text sets or appends text content on the Rb component.
func (rb *Rb) Text(s string) *Rb {
	rb.text = s
	return rb
}

// Bytes returns a defensive copy of the rendered Rb bytes.
func (rb *Rb) Bytes() []byte {
	return cloneBytes(rb.buf.Bytes())
}

// Render returns freshly prepared Rb HTML bytes.
func (r *Rb) Render() []byte {
	return renderPrepared(r)
}

// HTML returns freshly prepared Rb HTML as a string.
func (r *Rb) HTML() string {
	return htmlPrepared(r)
}

// String returns freshly prepared Rb HTML as a string.
func (r *Rb) String() string {
	return r.HTML()
}

// IsBodyElement implements BodyElement interface
func (rb *Rb) IsBodyElement() {}

// Prepare renders the Rb component into its internal buffer.
func (rb *Rb) Prepare() {
	rb.buf.Reset()
	rb.buf.WriteString("<rb")
	if len(rb.style) != 0 {
		parseStyle(&rb.buf, rb.style)
	}
	rb.buf.WriteString(">" + escapeText(rb.text) + "</rb>")
}

// Rt represents the Rt component or supporting type.
type Rt struct {
	buf   bytes.Buffer
	style StyleMap
	text  string
}

// NewRt creates a new Rt component.
func NewRt() *Rt {
	return &Rt{
		style: make(StyleMap),
	}
}

// AddStyle adds one inline CSS declaration to the Rt component.
func (rt *Rt) AddStyle(k, v string) *Rt {
	rt.style[k] = v
	return rt
}

// AddStyles adds multiple inline CSS declarations to the Rt component.
func (rt *Rt) AddStyles(m StyleMap) *Rt {
	for k, v := range m {
		rt.style[k] = v
	}
	return rt
}

// Style replaces the inline CSS declarations on the Rt component.
func (rt *Rt) Style(m StyleMap) *Rt {
	rt.style = cloneStyleMap(m)
	return rt
}

// Text sets or appends text content on the Rt component.
func (rt *Rt) Text(s string) *Rt {
	rt.text = s
	return rt
}

// Bytes returns a defensive copy of the rendered Rt bytes.
func (rt *Rt) Bytes() []byte {
	return cloneBytes(rt.buf.Bytes())
}

// Render returns freshly prepared Rt HTML bytes.
func (r *Rt) Render() []byte {
	return renderPrepared(r)
}

// HTML returns freshly prepared Rt HTML as a string.
func (r *Rt) HTML() string {
	return htmlPrepared(r)
}

// String returns freshly prepared Rt HTML as a string.
func (r *Rt) String() string {
	return r.HTML()
}

// IsBodyElement implements BodyElement interface
func (rt *Rt) IsBodyElement() {}

// Prepare renders the Rt component into its internal buffer.
func (rt *Rt) Prepare() {
	rt.buf.Reset()
	rt.buf.WriteString("<rt")
	if len(rt.style) != 0 {
		parseStyle(&rt.buf, rt.style)
	}
	rt.buf.WriteString(">" + escapeText(rt.text) + "</rt>")
}

// Rtc represents the Rtc component or supporting type.
type Rtc struct {
	buf      bytes.Buffer
	style    StyleMap
	contents []Element
}

// NewRtc creates a new Rtc component.
func NewRtc() *Rtc {
	return &Rtc{
		style: make(StyleMap),
	}
}

// AddStyle adds one inline CSS declaration to the Rtc component.
func (rtc *Rtc) AddStyle(k, v string) *Rtc {
	rtc.style[k] = v
	return rtc
}

// AddStyles adds multiple inline CSS declarations to the Rtc component.
func (rtc *Rtc) AddStyles(m StyleMap) *Rtc {
	for k, v := range m {
		rtc.style[k] = v
	}
	return rtc
}

// Style replaces the inline CSS declarations on the Rtc component.
func (rtc *Rtc) Style(m StyleMap) *Rtc {
	rtc.style = cloneStyleMap(m)
	return rtc
}

// Add appends child content to the Rtc component.
func (rtc *Rtc) Add(e Element) *Rtc {
	rtc.contents = appendElement(rtc.contents, e)
	return rtc
}

// Bytes returns a defensive copy of the rendered Rtc bytes.
func (rtc *Rtc) Bytes() []byte {
	return cloneBytes(rtc.buf.Bytes())
}

// Render returns freshly prepared Rtc HTML bytes.
func (r *Rtc) Render() []byte {
	return renderPrepared(r)
}

// HTML returns freshly prepared Rtc HTML as a string.
func (r *Rtc) HTML() string {
	return htmlPrepared(r)
}

// String returns freshly prepared Rtc HTML as a string.
func (r *Rtc) String() string {
	return r.HTML()
}

// IsBodyElement implements BodyElement interface
func (rtc *Rtc) IsBodyElement() {}

// Prepare renders the Rtc component into its internal buffer.
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

// Rp represents the Rp component or supporting type.
type Rp struct {
	buf   bytes.Buffer
	style StyleMap
	text  string
}

// NewRp creates a new Rp component.
func NewRp() *Rp {
	return &Rp{
		style: make(StyleMap),
	}
}

// AddStyle adds one inline CSS declaration to the Rp component.
func (rp *Rp) AddStyle(k, v string) *Rp {
	rp.style[k] = v
	return rp
}

// AddStyles adds multiple inline CSS declarations to the Rp component.
func (rp *Rp) AddStyles(m StyleMap) *Rp {
	for k, v := range m {
		rp.style[k] = v
	}
	return rp
}

// Style replaces the inline CSS declarations on the Rp component.
func (rp *Rp) Style(m StyleMap) *Rp {
	rp.style = cloneStyleMap(m)
	return rp
}

// Text sets or appends text content on the Rp component.
func (rp *Rp) Text(s string) *Rp {
	rp.text = s
	return rp
}

// Bytes returns a defensive copy of the rendered Rp bytes.
func (rp *Rp) Bytes() []byte {
	return cloneBytes(rp.buf.Bytes())
}

// Render returns freshly prepared Rp HTML bytes.
func (r *Rp) Render() []byte {
	return renderPrepared(r)
}

// HTML returns freshly prepared Rp HTML as a string.
func (r *Rp) HTML() string {
	return htmlPrepared(r)
}

// String returns freshly prepared Rp HTML as a string.
func (r *Rp) String() string {
	return r.HTML()
}

// IsBodyElement implements BodyElement interface
func (rp *Rp) IsBodyElement() {}

// Prepare renders the Rp component into its internal buffer.
func (rp *Rp) Prepare() {
	rp.buf.Reset()
	rp.buf.WriteString("<rp")
	if len(rp.style) != 0 {
		parseStyle(&rp.buf, rp.style)
	}
	rp.buf.WriteString(">" + escapeText(rp.text) + "</rp>")
}

// Kbd represents the Kbd component or supporting type.
type Kbd struct {
	buf   bytes.Buffer
	style StyleMap
	text  string
}

// NewKbd creates a new Kbd component.
func NewKbd() *Kbd {
	return &Kbd{
		style: make(StyleMap),
	}
}

// AddStyle adds one inline CSS declaration to the Kbd component.
func (k *Kbd) AddStyle(key, v string) *Kbd {
	k.style[key] = v
	return k
}

// AddStyles adds multiple inline CSS declarations to the Kbd component.
func (k *Kbd) AddStyles(m StyleMap) *Kbd {
	for key, v := range m {
		k.style[key] = v
	}
	return k
}

// Style replaces the inline CSS declarations on the Kbd component.
func (k *Kbd) Style(m StyleMap) *Kbd {
	k.style = cloneStyleMap(m)
	return k
}

// Text sets or appends text content on the Kbd component.
func (k *Kbd) Text(s string) *Kbd {
	k.text = s
	return k
}

// Bytes returns a defensive copy of the rendered Kbd bytes.
func (k *Kbd) Bytes() []byte {
	return cloneBytes(k.buf.Bytes())
}

// Render returns freshly prepared Kbd HTML bytes.
func (k *Kbd) Render() []byte {
	return renderPrepared(k)
}

// HTML returns freshly prepared Kbd HTML as a string.
func (k *Kbd) HTML() string {
	return htmlPrepared(k)
}

// String returns freshly prepared Kbd HTML as a string.
func (k *Kbd) String() string {
	return k.HTML()
}

// IsBodyElement implements BodyElement interface
func (k *Kbd) IsBodyElement() {}

// Prepare renders the Kbd component into its internal buffer.
func (k *Kbd) Prepare() {
	k.buf.Reset()
	k.buf.WriteString("<kbd")
	if len(k.style) != 0 {
		parseStyle(&k.buf, k.style)
	}
	k.buf.WriteString(">" + escapeText(k.text) + "</kbd>")
}

// Sub represents the Sub component or supporting type.
type Sub struct {
	buf   bytes.Buffer
	style StyleMap
	text  string
}

// NewSub creates a new Sub component.
func NewSub() *Sub {
	return &Sub{
		style: make(StyleMap),
	}
}

// AddStyle adds one inline CSS declaration to the Sub component.
func (s *Sub) AddStyle(k, v string) *Sub {
	s.style[k] = v
	return s
}

// AddStyles adds multiple inline CSS declarations to the Sub component.
func (s *Sub) AddStyles(m StyleMap) *Sub {
	for k, v := range m {
		s.style[k] = v
	}
	return s
}

// Style replaces the inline CSS declarations on the Sub component.
func (s *Sub) Style(m StyleMap) *Sub {
	s.style = cloneStyleMap(m)
	return s
}

// Text sets or appends text content on the Sub component.
func (s *Sub) Text(str string) *Sub {
	s.text = str
	return s
}

// Bytes returns a defensive copy of the rendered Sub bytes.
func (s *Sub) Bytes() []byte {
	return cloneBytes(s.buf.Bytes())
}

// Render returns freshly prepared Sub HTML bytes.
func (s *Sub) Render() []byte {
	return renderPrepared(s)
}

// HTML returns freshly prepared Sub HTML as a string.
func (s *Sub) HTML() string {
	return htmlPrepared(s)
}

// String returns freshly prepared Sub HTML as a string.
func (s *Sub) String() string {
	return s.HTML()
}

// IsBodyElement implements BodyElement interface
func (s *Sub) IsBodyElement() {}

// Prepare renders the Sub component into its internal buffer.
func (s *Sub) Prepare() {
	s.buf.Reset()
	s.buf.WriteString("<sub")
	if len(s.style) != 0 {
		parseStyle(&s.buf, s.style)
	}
	s.buf.WriteString(">" + escapeText(s.text) + "</sub>")
}

// Sup represents the Sup component or supporting type.
type Sup struct {
	buf   bytes.Buffer
	style StyleMap
	text  string
}

// NewSup creates a new Sup component.
func NewSup() *Sup {
	return &Sup{
		style: make(StyleMap),
	}
}

// AddStyle adds one inline CSS declaration to the Sup component.
func (s *Sup) AddStyle(k, v string) *Sup {
	s.style[k] = v
	return s
}

// AddStyles adds multiple inline CSS declarations to the Sup component.
func (s *Sup) AddStyles(m StyleMap) *Sup {
	for k, v := range m {
		s.style[k] = v
	}
	return s
}

// Style replaces the inline CSS declarations on the Sup component.
func (s *Sup) Style(m StyleMap) *Sup {
	s.style = cloneStyleMap(m)
	return s
}

// Text sets or appends text content on the Sup component.
func (s *Sup) Text(str string) *Sup {
	s.text = str
	return s
}

// Bytes returns a defensive copy of the rendered Sup bytes.
func (s *Sup) Bytes() []byte {
	return cloneBytes(s.buf.Bytes())
}

// Render returns freshly prepared Sup HTML bytes.
func (s *Sup) Render() []byte {
	return renderPrepared(s)
}

// HTML returns freshly prepared Sup HTML as a string.
func (s *Sup) HTML() string {
	return htmlPrepared(s)
}

// String returns freshly prepared Sup HTML as a string.
func (s *Sup) String() string {
	return s.HTML()
}

// IsBodyElement implements BodyElement interface
func (s *Sup) IsBodyElement() {}

// Prepare renders the Sup component into its internal buffer.
func (s *Sup) Prepare() {
	s.buf.Reset()
	s.buf.WriteString("<sup")
	if len(s.style) != 0 {
		parseStyle(&s.buf, s.style)
	}
	s.buf.WriteString(">" + escapeText(s.text) + "</sup>")
}

// Samp represents the Samp component or supporting type.
type Samp struct {
	buf   bytes.Buffer
	style StyleMap
	text  string
}

// NewSamp creates a new Samp component.
func NewSamp() *Samp {
	return &Samp{
		style: make(StyleMap),
	}
}

// AddStyle adds one inline CSS declaration to the Samp component.
func (s *Samp) AddStyle(k, v string) *Samp {
	s.style[k] = v
	return s
}

// AddStyles adds multiple inline CSS declarations to the Samp component.
func (s *Samp) AddStyles(m StyleMap) *Samp {
	for k, v := range m {
		s.style[k] = v
	}
	return s
}

// Style replaces the inline CSS declarations on the Samp component.
func (s *Samp) Style(m StyleMap) *Samp {
	s.style = cloneStyleMap(m)
	return s
}

// Text sets or appends text content on the Samp component.
func (s *Samp) Text(str string) *Samp {
	s.text = str
	return s
}

// Bytes returns a defensive copy of the rendered Samp bytes.
func (s *Samp) Bytes() []byte {
	return cloneBytes(s.buf.Bytes())
}

// Render returns freshly prepared Samp HTML bytes.
func (s *Samp) Render() []byte {
	return renderPrepared(s)
}

// HTML returns freshly prepared Samp HTML as a string.
func (s *Samp) HTML() string {
	return htmlPrepared(s)
}

// String returns freshly prepared Samp HTML as a string.
func (s *Samp) String() string {
	return s.HTML()
}

// IsBodyElement implements BodyElement interface
func (s *Samp) IsBodyElement() {}

// Prepare renders the Samp component into its internal buffer.
func (s *Samp) Prepare() {
	s.buf.Reset()
	s.buf.WriteString("<samp")
	if len(s.style) != 0 {
		parseStyle(&s.buf, s.style)
	}
	s.buf.WriteString(">" + escapeText(s.text) + "</samp>")
}

// Small represents the Small component or supporting type.
type Small struct {
	buf   bytes.Buffer
	style StyleMap
	text  string
}

// NewSmall creates a new Small component.
func NewSmall() *Small {
	return &Small{
		style: make(StyleMap),
	}
}

// AddStyle adds one inline CSS declaration to the Small component.
func (s *Small) AddStyle(k, v string) *Small {
	s.style[k] = v
	return s
}

// AddStyles adds multiple inline CSS declarations to the Small component.
func (s *Small) AddStyles(m StyleMap) *Small {
	for k, v := range m {
		s.style[k] = v
	}
	return s
}

// Style replaces the inline CSS declarations on the Small component.
func (s *Small) Style(m StyleMap) *Small {
	s.style = cloneStyleMap(m)
	return s
}

// Text sets or appends text content on the Small component.
func (s *Small) Text(str string) *Small {
	s.text = str
	return s
}

// Bytes returns a defensive copy of the rendered Small bytes.
func (s *Small) Bytes() []byte {
	return cloneBytes(s.buf.Bytes())
}

// Render returns freshly prepared Small HTML bytes.
func (s *Small) Render() []byte {
	return renderPrepared(s)
}

// HTML returns freshly prepared Small HTML as a string.
func (s *Small) HTML() string {
	return htmlPrepared(s)
}

// String returns freshly prepared Small HTML as a string.
func (s *Small) String() string {
	return s.HTML()
}

// IsBodyElement implements BodyElement interface
func (s *Small) IsBodyElement() {}

// Prepare renders the Small component into its internal buffer.
func (s *Small) Prepare() {
	s.buf.Reset()
	s.buf.WriteString("<small")
	if len(s.style) != 0 {
		parseStyle(&s.buf, s.style)
	}
	s.buf.WriteString(">" + escapeText(s.text) + "</small>")
}

// Span represents the Span component or supporting type.
type Span struct {
	buf   bytes.Buffer
	style StyleMap
	text  string
}

// NewSpan creates a new Span component.
func NewSpan() *Span {
	return &Span{
		style: make(StyleMap),
	}
}

// AddStyle adds one inline CSS declaration to the Span component.
func (s *Span) AddStyle(k, v string) *Span {
	s.style[k] = v
	return s
}

// AddStyles adds multiple inline CSS declarations to the Span component.
func (s *Span) AddStyles(m StyleMap) *Span {
	for k, v := range m {
		s.style[k] = v
	}
	return s
}

// Style replaces the inline CSS declarations on the Span component.
func (s *Span) Style(m StyleMap) *Span {
	s.style = cloneStyleMap(m)
	return s
}

// Text sets or appends text content on the Span component.
func (s *Span) Text(str string) *Span {
	s.text = str
	return s
}

// Bytes returns a defensive copy of the rendered Span bytes.
func (s *Span) Bytes() []byte {
	return cloneBytes(s.buf.Bytes())
}

// Render returns freshly prepared Span HTML bytes.
func (s *Span) Render() []byte {
	return renderPrepared(s)
}

// HTML returns freshly prepared Span HTML as a string.
func (s *Span) HTML() string {
	return htmlPrepared(s)
}

// String returns freshly prepared Span HTML as a string.
func (s *Span) String() string {
	return s.HTML()
}

// IsBodyElement implements BodyElement interface
func (s *Span) IsBodyElement() {}

// Prepare renders the Span component into its internal buffer.
func (s *Span) Prepare() {
	s.buf.Reset()
	s.buf.WriteString("<span")
	if len(s.style) != 0 {
		parseStyle(&s.buf, s.style)
	}
	s.buf.WriteString(">" + escapeText(s.text) + "</span>")
}

// Strong represents the Strong component or supporting type.
type Strong struct {
	buf   bytes.Buffer
	style StyleMap
	text  string
}

// NewStrong creates a new Strong component.
func NewStrong() *Strong {
	return &Strong{
		style: make(StyleMap),
	}
}

// AddStyle adds one inline CSS declaration to the Strong component.
func (s *Strong) AddStyle(k, v string) *Strong {
	s.style[k] = v
	return s
}

// AddStyles adds multiple inline CSS declarations to the Strong component.
func (s *Strong) AddStyles(m StyleMap) *Strong {
	for k, v := range m {
		s.style[k] = v
	}
	return s
}

// Style replaces the inline CSS declarations on the Strong component.
func (s *Strong) Style(m StyleMap) *Strong {
	s.style = cloneStyleMap(m)
	return s
}

// Text sets or appends text content on the Strong component.
func (s *Strong) Text(str string) *Strong {
	s.text = str
	return s
}

// Bytes returns a defensive copy of the rendered Strong bytes.
func (s *Strong) Bytes() []byte {
	return cloneBytes(s.buf.Bytes())
}

// Render returns freshly prepared Strong HTML bytes.
func (s *Strong) Render() []byte {
	return renderPrepared(s)
}

// HTML returns freshly prepared Strong HTML as a string.
func (s *Strong) HTML() string {
	return htmlPrepared(s)
}

// String returns freshly prepared Strong HTML as a string.
func (s *Strong) String() string {
	return s.HTML()
}

// IsBodyElement implements BodyElement interface
func (s *Strong) IsBodyElement() {}

// Prepare renders the Strong component into its internal buffer.
func (s *Strong) Prepare() {
	s.buf.Reset()
	s.buf.WriteString("<strong")
	if len(s.style) != 0 {
		parseStyle(&s.buf, s.style)
	}
	s.buf.WriteString(">" + escapeText(s.text) + "</strong>")
}

// Time represents the HTML time element and its machine-readable datetime value.
type Time struct {
	buf      bytes.Buffer
	style    StyleMap
	text     string
	datetime string
	err      error
}

// NewTime creates a time element.
func NewTime() *Time {
	return &Time{
		style: make(StyleMap),
	}
}

// AddStyle adds one inline CSS declaration to the time element.
func (t *Time) AddStyle(k, v string) *Time {
	t.style[k] = v
	return t
}

// AddStyles adds multiple inline CSS declarations to the time element.
func (t *Time) AddStyles(m StyleMap) *Time {
	for k, v := range m {
		t.style[k] = v
	}
	return t
}

// Style replaces the inline CSS declarations on the time element.
func (t *Time) Style(m StyleMap) *Time {
	t.style = cloneStyleMap(m)
	return t
}

// Text sets the human-readable time text.
func (t *Time) Text(str string) *Time {
	t.text = str
	return t
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

// Bytes returns a defensive copy of the rendered Time bytes.
func (t *Time) Bytes() []byte {
	return cloneBytes(t.buf.Bytes())
}

// Render returns freshly prepared Time HTML bytes.
func (t *Time) Render() []byte {
	return renderPrepared(t)
}

// HTML returns freshly prepared Time HTML as a string.
func (t *Time) HTML() string {
	return htmlPrepared(t)
}

// String returns freshly prepared Time HTML as a string.
func (t *Time) String() string {
	return t.HTML()
}

// IsBodyElement implements BodyElement interface
func (t *Time) IsBodyElement() {}

// Prepare renders the Time component into its internal buffer.
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

// Var represents the Var component or supporting type.
type Var struct {
	buf   bytes.Buffer
	style StyleMap
	text  string
}

// NewVar creates a new Var component.
func NewVar() *Var {
	return &Var{
		style: make(StyleMap),
	}
}

// AddStyle adds one inline CSS declaration to the Var component.
func (v *Var) AddStyle(k, val string) *Var {
	v.style[k] = val
	return v
}

// AddStyles adds multiple inline CSS declarations to the Var component.
func (v *Var) AddStyles(m StyleMap) *Var {
	for k, val := range m {
		v.style[k] = val
	}
	return v
}

// Style replaces the inline CSS declarations on the Var component.
func (v *Var) Style(m StyleMap) *Var {
	v.style = cloneStyleMap(m)
	return v
}

// Text sets or appends text content on the Var component.
func (v *Var) Text(str string) *Var {
	v.text = str
	return v
}

// Bytes returns a defensive copy of the rendered Var bytes.
func (v *Var) Bytes() []byte {
	return cloneBytes(v.buf.Bytes())
}

// Render returns freshly prepared Var HTML bytes.
func (v *Var) Render() []byte {
	return renderPrepared(v)
}

// HTML returns freshly prepared Var HTML as a string.
func (v *Var) HTML() string {
	return htmlPrepared(v)
}

// String returns freshly prepared Var HTML as a string.
func (v *Var) String() string {
	return v.HTML()
}

// IsBodyElement implements BodyElement interface
func (v *Var) IsBodyElement() {}

// Prepare renders the Var component into its internal buffer.
func (v *Var) Prepare() {
	v.buf.Reset()
	v.buf.WriteString("<var")
	if len(v.style) != 0 {
		parseStyle(&v.buf, v.style)
	}
	v.buf.WriteString(">" + escapeText(v.text) + "</var>")
}

// Wbr represents the Wbr component or supporting type.
type Wbr struct {
	buf   bytes.Buffer
	style StyleMap
}

// NewWbr creates a new Wbr component.
func NewWbr() *Wbr {
	return &Wbr{
		style: make(StyleMap),
	}
}

// AddStyle adds one inline CSS declaration to the Wbr component.
func (w *Wbr) AddStyle(k, v string) *Wbr {
	w.style[k] = v
	return w
}

// AddStyles adds multiple inline CSS declarations to the Wbr component.
func (w *Wbr) AddStyles(m StyleMap) *Wbr {
	for k, v := range m {
		w.style[k] = v
	}
	return w
}

// Style replaces the inline CSS declarations on the Wbr component.
func (w *Wbr) Style(m StyleMap) *Wbr {
	w.style = cloneStyleMap(m)
	return w
}

// Bytes returns a defensive copy of the rendered Wbr bytes.
func (w *Wbr) Bytes() []byte {
	return cloneBytes(w.buf.Bytes())
}

// Render returns freshly prepared Wbr HTML bytes.
func (w *Wbr) Render() []byte {
	return renderPrepared(w)
}

// HTML returns freshly prepared Wbr HTML as a string.
func (w *Wbr) HTML() string {
	return htmlPrepared(w)
}

// String returns freshly prepared Wbr HTML as a string.
func (w *Wbr) String() string {
	return w.HTML()
}

// IsBodyElement implements BodyElement interface
func (w *Wbr) IsBodyElement() {}

// Prepare renders the Wbr component into its internal buffer.
func (w *Wbr) Prepare() {
	w.buf.Reset()
	w.buf.WriteString("<wbr")
	if len(w.style) != 0 {
		parseStyle(&w.buf, w.style)
	}
	w.buf.WriteByte('>')
}
