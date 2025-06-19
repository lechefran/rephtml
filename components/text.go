package rephtml

import (
	"bytes"
	"log"
	"regexp"
	"time"
)

type P struct {
	buf   bytes.Buffer
	style map[string]string
	text  string
}

func NewP() *P {
	return &P{
		style: make(map[string]string),
	}
}

func (p *P) AddStyle(k, v string) *P {
	p.style[k] = v
	return p
}

func (p *P) AddStyles(m map[string]string) *P {
	for k, v := range m {
		p.style[k] = v
	}
	return p
}

func (p *P) Style(m map[string]string) *P {
	p.style = m
	return p
}

func (p *P) Text(s string) *P {
	p.text = s
	return p
}

func (p *P) Bytes() []byte {
	return p.buf.Bytes()
}

func (p *P) Prepare() {
	if len(p.style) != 0 {
		idx := 0
		p.buf.WriteString("<p style=\"")
		for k, v := range p.style {
			p.buf.WriteString(k + ": " + v + ";")
			if idx != len(p.style)-1 {
				p.buf.WriteByte(' ')
			}
			idx++
		}
		p.buf.WriteString("\">" + p.text + "</p>")
	} else {
		p.buf.WriteString("<p>" + p.text + "</p>")
	}
}

type Comment struct {
	buf  bytes.Buffer
	text string
}

func NewComment() *Comment {
	return &Comment{}
}

func (c *Comment) Text(s string) *Comment {
	c.text = s
	return c
}

func (c *Comment) Bytes() []byte {
	return c.buf.Bytes()
}

func (c *Comment) Prepare() {
	c.buf.WriteString("<!--" + c.text + "-->")
}

type H1 struct {
	buf   bytes.Buffer
	style map[string]string
	text  string
}

func NewH1() *H1 {
	return &H1{
		style: make(map[string]string),
	}
}

func (h *H1) AddStyle(k, v string) *H1 {
	h.style[k] = v
	return h
}

func (h *H1) Style(m map[string]string) *H1 {
	h.style = m
	return h
}

func (h *H1) Text(s string) *H1 {
	h.text = s
	return h
}

func (h *H1) Bytes() []byte {
	return h.buf.Bytes()
}

func (h *H1) Prepare() {
	if len(h.style) != 0 {
		idx := 0
		h.buf.WriteString("<h1 style=\"")
		for k, v := range h.style {
			h.buf.WriteString(k + ": " + v + ";")
			if idx != len(h.style)-1 {
				h.buf.WriteByte(' ')
			}
			idx++
		}
		h.buf.WriteString("\">" + h.text + "</h1>")
	} else {
		h.buf.WriteString("<h1>" + h.text + "</h1>")
	}
}

type H2 struct {
	buf   bytes.Buffer
	style map[string]string
	text  string
}

func NewH2() *H2 {
	return &H2{
		style: make(map[string]string),
	}
}

func (h *H2) AddStyle(k, v string) *H2 {
	h.style[k] = v
	return h
}

func (h *H2) Style(m map[string]string) *H2 {
	h.style = m
	return h
}

func (h *H2) Text(s string) *H2 {
	h.text = s
	return h
}

func (h *H2) Bytes() []byte {
	return h.buf.Bytes()
}

func (h *H2) Prepare() {
	if len(h.style) != 0 {
		idx := 0
		h.buf.WriteString("<h2 style=\"")
		for k, v := range h.style {
			h.buf.WriteString(k + ": " + v + ";")
			if idx != len(h.style)-1 {
				h.buf.WriteByte(' ')
			}
			idx++
		}
		h.buf.WriteString("\">" + h.text + "</h2>")
	} else {
		h.buf.WriteString("<h2>" + h.text + "</h2>")
	}
}

type H3 struct {
	buf   bytes.Buffer
	style map[string]string
	text  string
}

func NewH3() *H3 {
	return &H3{
		style: make(map[string]string),
	}
}

func (h *H3) AddStyle(k, v string) *H3 {
	h.style[k] = v
	return h
}

func (h *H3) Style(m map[string]string) *H3 {
	h.style = m
	return h
}

func (h *H3) Text(s string) *H3 {
	h.text = s
	return h
}

func (h *H3) Bytes() []byte {
	return h.buf.Bytes()
}

func (h *H3) Prepare() {
	if len(h.style) != 0 {
		idx := 0
		h.buf.WriteString("<h3 style=\"")
		for k, v := range h.style {
			h.buf.WriteString(k + ": " + v + ";")
			if idx != len(h.style)-1 {
				h.buf.WriteByte(' ')
			}
			idx++
		}
		h.buf.WriteString("\">" + h.text + "</h3>")
	} else {
		h.buf.WriteString("<h3>" + h.text + "</h3>")
	}
}

type H4 struct {
	buf   bytes.Buffer
	style map[string]string
	text  string
}

func NewH4() *H4 {
	return &H4{
		style: make(map[string]string),
	}
}

func (h *H4) AddStyle(k, v string) *H4 {
	h.style[k] = v
	return h
}

func (h *H4) Style(m map[string]string) *H4 {
	h.style = m
	return h
}

func (h *H4) Text(s string) *H4 {
	h.text = s
	return h
}

func (h *H4) Bytes() []byte {
	return h.buf.Bytes()
}

func (h *H4) Prepare() {
	if len(h.style) != 0 {
		idx := 0
		h.buf.WriteString("<h4 style=\"")
		for k, v := range h.style {
			h.buf.WriteString(k + ": " + v + ";")
			if idx != len(h.style)-1 {
				h.buf.WriteByte(' ')
			}
			idx++
		}
		h.buf.WriteString("\">" + h.text + "</h4>")
	} else {
		h.buf.WriteString("<h4>" + h.text + "</h4>")
	}
}

type H5 struct {
	buf   bytes.Buffer
	style map[string]string
	text  string
}

func NewH5() *H5 {
	return &H5{
		style: make(map[string]string),
	}
}

func (h *H5) AddStyle(k, v string) *H5 {
	h.style[k] = v
	return h
}

func (h *H5) Style(m map[string]string) *H5 {
	h.style = m
	return h
}

func (h *H5) Text(s string) *H5 {
	h.text = s
	return h
}

func (h *H5) Bytes() []byte {
	return h.buf.Bytes()
}

func (h *H5) Prepare() {
	if len(h.style) != 0 {
		idx := 0
		h.buf.WriteString("<h5 style=\"")
		for k, v := range h.style {
			h.buf.WriteString(k + ": " + v + ";")
			if idx != len(h.style)-1 {
				h.buf.WriteByte(' ')
			}
			idx++
		}
		h.buf.WriteString("\">" + h.text + "</h5>")
	} else {
		h.buf.WriteString("<h5>" + h.text + "</h5>")
	}
}

type H6 struct {
	buf   bytes.Buffer
	style map[string]string
	text  string
}

func NewH6() *H6 {
	return &H6{
		style: make(map[string]string),
	}
}

func (h *H6) AddStyle(k, v string) *H6 {
	h.style[k] = v
	return h
}

func (h *H6) Style(m map[string]string) *H6 {
	h.style = m
	return h
}

func (h *H6) Text(s string) *H6 {
	h.text = s
	return h
}

func (h *H6) Bytes() []byte {
	return h.buf.Bytes()
}

func (h *H6) Prepare() {
	if len(h.style) != 0 {
		idx := 0
		h.buf.WriteString("<h6 style=\"")
		for k, v := range h.style {
			h.buf.WriteString(k + ": " + v + ";")
			if idx != len(h.style)-1 {
				h.buf.WriteByte(' ')
			}
			idx++
		}
		h.buf.WriteString("\">" + h.text + "</h6>")
	} else {
		h.buf.WriteString("<h6>" + h.text + "</h6>")
	}
}

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
	return a.buf.Bytes()
}

func (a *Anchor) Prepare() {
	if len(a.style) != 0 {
		idx := 0
		a.buf.WriteString("<a style=\"")
		for k, v := range a.style {
			a.buf.WriteString(k + ": " + v + ";")
			if idx != len(a.style)-1 {
				a.buf.WriteByte(' ')
			}
			idx++
		}
		if a.link != "" {
			a.buf.WriteString(" href=\"" + a.link + "\">" + a.text + "</a>")
		} else {
			a.buf.WriteString(">" + a.text + "</a>")
		}
	} else {
		if a.link != "" {
			a.buf.WriteString("<a href=\"" + a.link + "\">" + a.text + "</a>")
		} else {
			a.buf.WriteString("<a>" + a.text + "</a>")
		}
	}
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
	return a.buf.Bytes()
}

func (a *Abbr) Prepare() {
	if len(a.style) != 0 {
		idx := 0
		a.buf.WriteString("<abbr style=\"")
		for k, v := range a.style {
			a.buf.WriteString(k + ": " + v + ";")
			if idx != len(a.style)-1 {
				a.buf.WriteByte(' ')
			}
			idx++
		}
		a.buf.WriteString("\"")
		if a.title != "" {
			a.buf.WriteString(" title=\"" + a.title + "\"")
		}
		a.buf.WriteString(">" + a.text + "</abbr>")
	} else {
		a.buf.WriteString("<abbr")
		if a.title != "" {
			a.buf.WriteString(" title=\"" + a.title + "\"")
		}
		a.buf.WriteString(">" + a.text + "</abbr>")
	}
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
	return b.buf.Bytes()
}

func (b *B) Prepare() {
	if len(b.style) != 0 {
		idx := 0
		b.buf.WriteString("<b style=\"")
		for k, v := range b.style {
			b.buf.WriteString(k + ": " + v + ";")
			if idx != len(b.style)-1 {
				b.buf.WriteByte(' ')
			}
			idx++
		}
		b.buf.WriteString("\">" + b.text + "</b>")
	} else {
		b.buf.WriteString("<b>" + b.text + "</b>")
	}
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
	return i.buf.Bytes()
}

func (i *I) Prepare() {
	if len(i.style) != 0 {
		idx := 0
		i.buf.WriteString("<i style=\"")
		for k, v := range i.style {
			i.buf.WriteString(k + ": " + v + ";")
			if idx != len(i.style)-1 {
				i.buf.WriteByte(' ')
			}
			idx++
		}
		i.buf.WriteString("\">" + i.text + "</i>")
	} else {
		i.buf.WriteString("<i>" + i.text + "</i>")
	}
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
	return q.buf.Bytes()
}

func (q *Q) Prepare() {
	if len(q.style) != 0 {
		idx := 0
		q.buf.WriteString("<q style=\"")
		for k, v := range q.style {
			q.buf.WriteString(k + ": " + v + ";")
			if idx != len(q.style)-1 {
				q.buf.WriteByte(' ')
			}
			idx++
		}
		q.buf.WriteString("\"")
		if q.cite != "" {
			q.buf.WriteString(" cite=\"" + q.cite + "\"")
		}
		q.buf.WriteString(">" + q.text + "</q>")
	} else {
		q.buf.WriteString("<q")
		if q.cite != "" {
			q.buf.WriteString(" cite=\"" + q.cite + "\"")
		}
		q.buf.WriteString(">" + q.text + "</q>")
	}
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
	return s.buf.Bytes()
}

func (s *S) Prepare() {
	if len(s.style) != 0 {
		idx := 0
		s.buf.WriteString("<s style=\"")
		for k, v := range s.style {
			s.buf.WriteString(k + ": " + v + ";")
			if idx != len(s.style)-1 {
				s.buf.WriteByte(' ')
			}
			idx++
		}
		s.buf.WriteString("\">" + s.text + "</s>")
	} else {
		s.buf.WriteString("<s>" + s.text + "</s>")
	}
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
	return u.buf.Bytes()
}

func (u *U) Prepare() {
	if len(u.style) != 0 {
		idx := 0
		u.buf.WriteString("<u style=\"")
		for k, v := range u.style {
			u.buf.WriteString(k + ": " + v + ";")
			if idx != len(u.style)-1 {
				u.buf.WriteByte(' ')
			}
			idx++
		}
		u.buf.WriteString("\">" + u.text + "</u>")
	} else {
		u.buf.WriteString("<u>" + u.text + "</u>")
	}
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
	return d.buf.Bytes()
}

func (d *Dbi) Prepare() {
	if len(d.style) != 0 {
		idx := 0
		d.buf.WriteString("<dbi style=\"")
		for k, v := range d.style {
			d.buf.WriteString(k + ": " + v + ";")
			if idx != len(d.style)-1 {
				d.buf.WriteByte(' ')
			}
			idx++
		}
		d.buf.WriteString("\">" + d.text + "</dbi>")
	} else {
		d.buf.WriteString("<dbi>" + d.text + "</dbi>")
	}
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
	return d.buf.Bytes()
}

func (d *Dbo) Prepare() {
	if len(d.style) != 0 {
		idx := 0
		d.buf.WriteString("<dbo style=\"")
		for k, v := range d.style {
			d.buf.WriteString(k + ": " + v + ";")
			if idx != len(d.style)-1 {
				d.buf.WriteByte(' ')
			}
			idx++
		}
		d.buf.WriteString("\">" + d.text + "</dbo>")
	} else {
		d.buf.WriteString("<dbo>" + d.text + "</dbo>")
	}
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
	return br.buf.Bytes()
}

func (br *Br) Prepare() {
	if len(br.style) != 0 {
		idx := 0
		br.buf.WriteString("<br style=\"")
		for k, v := range br.style {
			br.buf.WriteString(k + ": " + v + ";")
			if idx != len(br.style)-1 {
				br.buf.WriteByte(' ')
			}
			idx++
		}
		br.buf.WriteString("\">")
	} else {
		br.buf.WriteString("<br>")
	}
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
	return c.buf.Bytes()
}

func (c *Cite) Prepare() {
	if len(c.style) != 0 {
		idx := 0
		c.buf.WriteString("<cite style=\"")
		for k, v := range c.style {
			c.buf.WriteString(k + ": " + v + ";")
			if idx != len(c.style)-1 {
				c.buf.WriteByte(' ')
			}
			idx++
		}
		c.buf.WriteString("\">" + c.text + "</cite>")
	} else {
		c.buf.WriteString("<cite>" + c.text + "</cite>")
	}
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
	return c.buf.Bytes()
}

func (c *Code) Prepare() {
	if len(c.style) != 0 {
		idx := 0
		c.buf.WriteString("<code style=\"")
		for k, v := range c.style {
			c.buf.WriteString(k + ": " + v + ";")
			if idx != len(c.style)-1 {
				c.buf.WriteByte(' ')
			}
			idx++
		}
		c.buf.WriteString("\">" + c.text + "</code>")
	} else {
		c.buf.WriteString("<code>" + c.text + "</code>")
	}
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
	return d.buf.Bytes()
}

func (d *Data) Prepare() {
	if len(d.style) != 0 {
		idx := 0
		d.buf.WriteString("<data style=\"")
		for k, v := range d.style {
			d.buf.WriteString(k + ": " + v + ";")
			if idx != len(d.style)-1 {
				d.buf.WriteByte(' ')
			}
			idx++
		}
		d.buf.WriteString("\"")
		if d.value != "" {
			d.buf.WriteString(" value=\"" + d.value + "\"")
		}
		d.buf.WriteString(">" + d.text + "</data>")
	} else {
		d.buf.WriteString("<data")
		if d.value != "" {
			d.buf.WriteString(" value=\"" + d.value + "\"")
		}
		d.buf.WriteString(">" + d.text + "</data>")
	}
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
	return d.buf.Bytes()
}

func (d *Dfn) Prepare() {
	if len(d.style) != 0 {
		idx := 0
		d.buf.WriteString("<dfn style=\"")
		for k, v := range d.style {
			d.buf.WriteString(k + ": " + v + ";")
			if idx != len(d.style)-1 {
				d.buf.WriteByte(' ')
			}
			idx++
		}
		d.buf.WriteString("\"")
		if d.title != "" {
			d.buf.WriteString(" title=\"" + d.title + "\"")
		}
		d.buf.WriteString(">" + d.text + "</dfn>")
	} else {
		d.buf.WriteString("<dfn")
		if d.title != "" {
			d.buf.WriteString(" title=\"" + d.title + "\"")
		}
		d.buf.WriteString(">" + d.text + "</dfn>")
	}
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
	return e.buf.Bytes()
}

func (e *Elem) Prepare() {
	if len(e.style) != 0 {
		idx := 0
		e.buf.WriteString("<elem style=\"")
		for k, v := range e.style {
			e.buf.WriteString(k + ": " + v + ";")
			if idx != len(e.style)-1 {
				e.buf.WriteByte(' ')
			}
			idx++
		}
		e.buf.WriteString("\">" + e.text + "</elem>")
	} else {
		e.buf.WriteString("<elem>" + e.text + "</elem>")
	}
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
	return m.buf.Bytes()
}

func (m *Mark) Prepare() {
	if len(m.style) != 0 {
		idx := 0
		m.buf.WriteString("<mark style=\"")
		for k, v := range m.style {
			m.buf.WriteString(k + ": " + v + ";")
			if idx != len(m.style)-1 {
				m.buf.WriteByte(' ')
			}
			idx++
		}
		m.buf.WriteString("\">" + m.text + "</mark>")
	} else {
		m.buf.WriteString("<mark>" + m.text + "</mark>")
	}
}

type Ruby struct {
	buf      bytes.Buffer
	style    map[string]string
	contents [][]byte
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

func (r *Ruby) Add(e Elements) *Ruby {
	r.contents = append(r.contents, e.Bytes())
	return r
}

func (r *Ruby) Bytes() []byte {
	return r.buf.Bytes()
}

func (r *Ruby) Prepare() {
	if len(r.style) != 0 {
		idx := 0
		r.buf.WriteString("<ruby style=\"")
		for k, v := range r.style {
			r.buf.WriteString(k + ": " + v + ";")
			if idx != len(r.style)-1 {
				r.buf.WriteByte(' ')
			}
			idx++
		}
		r.buf.WriteString("\">")
	} else {
		r.buf.WriteString("<ruby>")
	}

	for _, content := range r.contents {
		r.buf.Write(content)
	}
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
	return rb.buf.Bytes()
}

func (rb *Rb) Prepare() {
	if len(rb.style) != 0 {
		idx := 0
		rb.buf.WriteString("<rb style=\"")
		for k, v := range rb.style {
			rb.buf.WriteString(k + ": " + v + ";")
			if idx != len(rb.style)-1 {
				rb.buf.WriteByte(' ')
			}
			idx++
		}
		rb.buf.WriteString("\">" + rb.text + "</rb>")
	} else {
		rb.buf.WriteString("<rb>" + rb.text + "</rb>")
	}
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
	return rt.buf.Bytes()
}

func (rt *Rt) Prepare() {
	if len(rt.style) != 0 {
		idx := 0
		rt.buf.WriteString("<rt style=\"")
		for k, v := range rt.style {
			rt.buf.WriteString(k + ": " + v + ";")
			if idx != len(rt.style)-1 {
				rt.buf.WriteByte(' ')
			}
			idx++
		}
		rt.buf.WriteString("\">" + rt.text + "</rt>")
	} else {
		rt.buf.WriteString("<rt>" + rt.text + "</rt>")
	}
}

type Rtc struct {
	buf      bytes.Buffer
	style    map[string]string
	contents [][]byte
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

func (rtc *Rtc) Add(e Elements) *Rtc {
	rtc.contents = append(rtc.contents, e.Bytes())
	return rtc
}

func (rtc *Rtc) Bytes() []byte {
	return rtc.buf.Bytes()
}

func (rtc *Rtc) Prepare() {
	if len(rtc.style) != 0 {
		idx := 0
		rtc.buf.WriteString("<rtc style=\"")
		for k, v := range rtc.style {
			rtc.buf.WriteString(k + ": " + v + ";")
			if idx != len(rtc.style)-1 {
				rtc.buf.WriteByte(' ')
			}
			idx++
		}
		rtc.buf.WriteString("\">")
	} else {
		rtc.buf.WriteString("<rtc>")
	}

	for _, content := range rtc.contents {
		rtc.buf.Write(content)
	}
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
	return rp.buf.Bytes()
}

func (rp *Rp) Prepare() {
	if len(rp.style) != 0 {
		idx := 0
		rp.buf.WriteString("<rp style=\"")
		for k, v := range rp.style {
			rp.buf.WriteString(k + ": " + v + ";")
			if idx != len(rp.style)-1 {
				rp.buf.WriteByte(' ')
			}
			idx++
		}
		rp.buf.WriteString("\">" + rp.text + "</rp>")
	} else {
		rp.buf.WriteString("<rp>" + rp.text + "</rp>")
	}
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
	return k.buf.Bytes()
}

func (k *Kbd) Prepare() {
	if len(k.style) != 0 {
		idx := 0
		k.buf.WriteString("<kbd style=\"")
		for key, v := range k.style {
			k.buf.WriteString(key + ": " + v + ";")
			if idx != len(k.style)-1 {
				k.buf.WriteByte(' ')
			}
			idx++
		}
		k.buf.WriteString("\">" + k.text + "</kbd>")
	} else {
		k.buf.WriteString("<kbd>" + k.text + "</kbd>")
	}
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
	return s.buf.Bytes()
}

func (s *Sub) Prepare() {
	if len(s.style) != 0 {
		idx := 0
		s.buf.WriteString("<sub style=\"")
		for k, v := range s.style {
			s.buf.WriteString(k + ": " + v + ";")
			if idx != len(s.style)-1 {
				s.buf.WriteByte(' ')
			}
			idx++
		}
		s.buf.WriteString("\">" + s.text + "</sub>")
	} else {
		s.buf.WriteString("<sub>" + s.text + "</sub>")
	}
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
	return s.buf.Bytes()
}

func (s *Sup) Prepare() {
	if len(s.style) != 0 {
		idx := 0
		s.buf.WriteString("<sup style=\"")
		for k, v := range s.style {
			s.buf.WriteString(k + ": " + v + ";")
			if idx != len(s.style)-1 {
				s.buf.WriteByte(' ')
			}
			idx++
		}
		s.buf.WriteString("\">" + s.text + "</sup>")
	} else {
		s.buf.WriteString("<sup>" + s.text + "</sup>")
	}
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
	return s.buf.Bytes()
}

func (s *Samp) Prepare() {
	if len(s.style) != 0 {
		idx := 0
		s.buf.WriteString("<samp style=\"")
		for k, v := range s.style {
			s.buf.WriteString(k + ": " + v + ";")
			if idx != len(s.style)-1 {
				s.buf.WriteByte(' ')
			}
			idx++
		}
		s.buf.WriteString("\">" + s.text + "</samp>")
	} else {
		s.buf.WriteString("<samp>" + s.text + "</samp>")
	}
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
	return s.buf.Bytes()
}

func (s *Small) Prepare() {
	if len(s.style) != 0 {
		idx := 0
		s.buf.WriteString("<small style=\"")
		for k, v := range s.style {
			s.buf.WriteString(k + ": " + v + ";")
			if idx != len(s.style)-1 {
				s.buf.WriteByte(' ')
			}
			idx++
		}
		s.buf.WriteString("\">" + s.text + "</small>")
	} else {
		s.buf.WriteString("<small>" + s.text + "</small>")
	}
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
	return s.buf.Bytes()
}

func (s *Span) Prepare() {
	if len(s.style) != 0 {
		idx := 0
		s.buf.WriteString("<span style=\"")
		for k, v := range s.style {
			s.buf.WriteString(k + ": " + v + ";")
			if idx != len(s.style)-1 {
				s.buf.WriteByte(' ')
			}
			idx++
		}
		s.buf.WriteString("\">" + s.text + "</span>")
	} else {
		s.buf.WriteString("<span>" + s.text + "</span>")
	}
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
	return s.buf.Bytes()
}

func (s *Strong) Prepare() {
	if len(s.style) != 0 {
		idx := 0
		s.buf.WriteString("<strong style=\"")
		for k, v := range s.style {
			s.buf.WriteString(k + ": " + v + ";")
			if idx != len(s.style)-1 {
				s.buf.WriteByte(' ')
			}
			idx++
		}
		s.buf.WriteString("\">" + s.text + "</strong>")
	} else {
		s.buf.WriteString("<strong>" + s.text + "</strong>")
	}
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
	return t.buf.Bytes()
}

func (t *Time) Prepare() {
	t.buf.WriteString("<time")
	if t.datetime != "" {
		t.buf.WriteString(" datetime=\"" + t.datetime + "\"")
	}
	if len(t.style) != 0 {
		idx := 0
		t.buf.WriteString(" style=\"")
		for k, v := range t.style {
			t.buf.WriteString(k + ": " + v + ";")
			if idx != len(t.style)-1 {
				t.buf.WriteByte(' ')
			}
			idx++
		}
		t.buf.WriteString("\"")
	}
	t.buf.WriteString(">" + t.text + "</time>")
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
	return v.buf.Bytes()
}

func (v *Var) Prepare() {
	if len(v.style) != 0 {
		idx := 0
		v.buf.WriteString("<var style=\"")
		for k, val := range v.style {
			v.buf.WriteString(k + ": " + val + ";")
			if idx != len(v.style)-1 {
				v.buf.WriteByte(' ')
			}
			idx++
		}
		v.buf.WriteString("\">" + v.text + "</var>")
	} else {
		v.buf.WriteString("<var>" + v.text + "</var>")
	}
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
	return w.buf.Bytes()
}

func (w *Wbr) Prepare() {
	if len(w.style) != 0 {
		idx := 0
		w.buf.WriteString("<wbr style=\"")
		for k, v := range w.style {
			w.buf.WriteString(k + ": " + v + ";")
			if idx != len(w.style)-1 {
				w.buf.WriteByte(' ')
			}
			idx++
		}
		w.buf.WriteString("\">")
	} else {
		w.buf.WriteString("<wbr>")
	}
}

type Hr struct {
	buf   bytes.Buffer
	style map[string]string
}

func NewHr() *Hr {
	return &Hr{
		style: make(map[string]string),
	}
}

func (h *Hr) AddStyle(k, v string) *Hr {
	h.style[k] = v
	return h
}

func (h *Hr) AddStyles(m map[string]string) *Hr {
	for k, v := range m {
		h.style[k] = v
	}
	return h
}

func (h *Hr) Style(m map[string]string) *Hr {
	h.style = m
	return h
}

func (h *Hr) Bytes() []byte {
	return h.buf.Bytes()
}

func (h *Hr) Prepare() {
	if len(h.style) != 0 {
		idx := 0
		h.buf.WriteString("<hr style=\"")
		for k, v := range h.style {
			h.buf.WriteString(k + ": " + v + ";")
			if idx != len(h.style)-1 {
				h.buf.WriteByte(' ')
			}
			idx++
		}
		h.buf.WriteString("\">")
	} else {
		h.buf.WriteString("<hr>")
	}
}

type Pre struct {
	buf   bytes.Buffer
	style map[string]string
	text  string
}

func NewPre() *Pre {
	return &Pre{
		style: make(map[string]string),
	}
}

func (p *Pre) AddStyle(k, v string) *Pre {
	p.style[k] = v
	return p
}

func (p *Pre) AddStyles(m map[string]string) *Pre {
	for k, v := range m {
		p.style[k] = v
	}
	return p
}

func (p *Pre) Style(m map[string]string) *Pre {
	p.style = m
	return p
}

func (p *Pre) Text(str string) *Pre {
	p.text = str
	return p
}

func (p *Pre) Bytes() []byte {
	return p.buf.Bytes()
}

func (p *Pre) Prepare() {
	if len(p.style) != 0 {
		idx := 0
		p.buf.WriteString("<pre style=\"")
		for k, v := range p.style {
			p.buf.WriteString(k + ": " + v + ";")
			if idx != len(p.style)-1 {
				p.buf.WriteByte(' ')
			}
			idx++
		}
		p.buf.WriteString("\">" + p.text + "</pre>")
	} else {
		p.buf.WriteString("<pre>" + p.text + "</pre>")
	}
}

type Blockquote struct {
	buf   bytes.Buffer
	style map[string]string
	text  string
	cite  string
}

func NewBlockquote() *Blockquote {
	return &Blockquote{
		style: make(map[string]string),
	}
}

func (b *Blockquote) AddStyle(k, v string) *Blockquote {
	b.style[k] = v
	return b
}

func (b *Blockquote) AddStyles(m map[string]string) *Blockquote {
	for k, v := range m {
		b.style[k] = v
	}
	return b
}

func (b *Blockquote) Style(m map[string]string) *Blockquote {
	b.style = m
	return b
}

func (b *Blockquote) Text(str string) *Blockquote {
	b.text = str
	return b
}

func (b *Blockquote) Cite(c string) *Blockquote {
	b.cite = c
	return b
}

func (b *Blockquote) Bytes() []byte {
	return b.buf.Bytes()
}

func (b *Blockquote) Prepare() {
	b.buf.WriteString("<blockquote")
	if b.cite != "" {
		b.buf.WriteString(" cite=\"" + b.cite + "\"")
	}
	if len(b.style) != 0 {
		idx := 0
		b.buf.WriteString(" style=\"")
		for k, v := range b.style {
			b.buf.WriteString(k + ": " + v + ";")
			if idx != len(b.style)-1 {
				b.buf.WriteByte(' ')
			}
			idx++
		}
		b.buf.WriteString("\"")
	}
	b.buf.WriteString(">" + b.text + "</blockquote>")
}

type Menu struct {
	buf      bytes.Buffer
	style    map[string]string
	contents [][]byte
	menuType string
	label    string
}

func NewMenu() *Menu {
	return &Menu{
		style: make(map[string]string),
	}
}

func (m *Menu) AddStyle(k, v string) *Menu {
	m.style[k] = v
	return m
}

func (m *Menu) AddStyles(ms map[string]string) *Menu {
	for k, v := range ms {
		m.style[k] = v
	}
	return m
}

func (m *Menu) Style(ms map[string]string) *Menu {
	m.style = ms
	return m
}

func (m *Menu) Add(e Elements) *Menu {
	m.contents = append(m.contents, e.Bytes())
	return m
}

func (m *Menu) Type(t string) *Menu {
	m.menuType = t
	return m
}

func (m *Menu) Label(l string) *Menu {
	m.label = l
	return m
}

func (m *Menu) Bytes() []byte {
	return m.buf.Bytes()
}

func (m *Menu) Prepare() {
	m.buf.WriteString("<menu")
	if m.menuType != "" {
		m.buf.WriteString(" type=\"" + m.menuType + "\"")
	}
	if m.label != "" {
		m.buf.WriteString(" label=\"" + m.label + "\"")
	}
	if len(m.style) != 0 {
		idx := 0
		m.buf.WriteString(" style=\"")
		for k, v := range m.style {
			m.buf.WriteString(k + ": " + v + ";")
			if idx != len(m.style)-1 {
				m.buf.WriteByte(' ')
			}
			idx++
		}
		m.buf.WriteString("\"")
	}
	m.buf.WriteByte('>')

	for _, content := range m.contents {
		m.buf.Write(content)
	}
	m.buf.WriteString("</menu>")
}

type Ol struct {
	buf      bytes.Buffer
	style    map[string]string
	contents [][]byte
	start    int
	listType string
	reversed bool
}

func NewOl() *Ol {
	return &Ol{
		style: make(map[string]string),
	}
}

func (o *Ol) AddStyle(k, v string) *Ol {
	o.style[k] = v
	return o
}

func (o *Ol) AddStyles(m map[string]string) *Ol {
	for k, v := range m {
		o.style[k] = v
	}
	return o
}

func (o *Ol) Style(m map[string]string) *Ol {
	o.style = m
	return o
}

func (o *Ol) Add(e Elements) *Ol {
	o.contents = append(o.contents, e.Bytes())
	return o
}

func (o *Ol) Start(s int) *Ol {
	o.start = s
	return o
}

func (o *Ol) Type(t string) *Ol {
	o.listType = t
	return o
}

func (o *Ol) Reversed(r bool) *Ol {
	o.reversed = r
	return o
}

func (o *Ol) Bytes() []byte {
	return o.buf.Bytes()
}

func (o *Ol) Prepare() {
	o.buf.WriteString("<ol")
	if o.start > 0 {
		o.buf.WriteString(" start=\"")
		o.buf.WriteString(string(rune(o.start + '0')))
		o.buf.WriteString("\"")
	}
	if o.listType != "" {
		o.buf.WriteString(" type=\"" + o.listType + "\"")
	}
	if o.reversed {
		o.buf.WriteString(" reversed")
	}
	if len(o.style) != 0 {
		idx := 0
		o.buf.WriteString(" style=\"")
		for k, v := range o.style {
			o.buf.WriteString(k + ": " + v + ";")
			if idx != len(o.style)-1 {
				o.buf.WriteByte(' ')
			}
			idx++
		}
		o.buf.WriteString("\"")
	}
	o.buf.WriteByte('>')

	for _, content := range o.contents {
		o.buf.Write(content)
	}
	o.buf.WriteString("</ol>")
}

type Ul struct {
	buf      bytes.Buffer
	style    map[string]string
	contents [][]byte
}

func NewUl() *Ul {
	return &Ul{
		style: make(map[string]string),
	}
}

func (u *Ul) AddStyle(k, v string) *Ul {
	u.style[k] = v
	return u
}

func (u *Ul) AddStyles(m map[string]string) *Ul {
	for k, v := range m {
		u.style[k] = v
	}
	return u
}

func (u *Ul) Style(m map[string]string) *Ul {
	u.style = m
	return u
}

func (u *Ul) Add(e Elements) *Ul {
	u.contents = append(u.contents, e.Bytes())
	return u
}

func (u *Ul) Bytes() []byte {
	return u.buf.Bytes()
}

func (u *Ul) Prepare() {
	if len(u.style) != 0 {
		idx := 0
		u.buf.WriteString("<ul style=\"")
		for k, v := range u.style {
			u.buf.WriteString(k + ": " + v + ";")
			if idx != len(u.style)-1 {
				u.buf.WriteByte(' ')
			}
			idx++
		}
		u.buf.WriteString("\">")
	} else {
		u.buf.WriteString("<ul>")
	}

	for _, content := range u.contents {
		u.buf.Write(content)
	}
	u.buf.WriteString("</ul>")
}

type Li struct {
	buf      bytes.Buffer
	style    map[string]string
	contents [][]byte
	value    int
}

func NewLi() *Li {
	return &Li{
		style: make(map[string]string),
	}
}

func (l *Li) AddStyle(k, v string) *Li {
	l.style[k] = v
	return l
}

func (l *Li) AddStyles(m map[string]string) *Li {
	for k, v := range m {
		l.style[k] = v
	}
	return l
}

func (l *Li) Style(m map[string]string) *Li {
	l.style = m
	return l
}

func (l *Li) Add(e Elements) *Li {
	l.contents = append(l.contents, e.Bytes())
	return l
}

func (l *Li) Value(v int) *Li {
	l.value = v
	return l
}

func (l *Li) Bytes() []byte {
	return l.buf.Bytes()
}

func (l *Li) Prepare() {
	l.buf.WriteString("<li")
	if l.value > 0 {
		l.buf.WriteString(" value=\"")
		l.buf.WriteString(string(rune(l.value + '0')))
		l.buf.WriteString("\"")
	}
	if len(l.style) != 0 {
		idx := 0
		l.buf.WriteString(" style=\"")
		for k, v := range l.style {
			l.buf.WriteString(k + ": " + v + ";")
			if idx != len(l.style)-1 {
				l.buf.WriteByte(' ')
			}
			idx++
		}
		l.buf.WriteString("\"")
	}
	l.buf.WriteByte('>')

	for _, content := range l.contents {
		l.buf.Write(content)
	}
	l.buf.WriteString("</li>")
}

type Dl struct {
	buf      bytes.Buffer
	style    map[string]string
	contents [][]byte
}

func NewDl() *Dl {
	return &Dl{
		style: make(map[string]string),
	}
}

func (d *Dl) AddStyle(k, v string) *Dl {
	d.style[k] = v
	return d
}

func (d *Dl) AddStyles(m map[string]string) *Dl {
	for k, v := range m {
		d.style[k] = v
	}
	return d
}

func (d *Dl) Style(m map[string]string) *Dl {
	d.style = m
	return d
}

func (d *Dl) Add(e Elements) *Dl {
	d.contents = append(d.contents, e.Bytes())
	return d
}

func (d *Dl) Bytes() []byte {
	return d.buf.Bytes()
}

func (d *Dl) Prepare() {
	if len(d.style) != 0 {
		idx := 0
		d.buf.WriteString("<dl style=\"")
		for k, v := range d.style {
			d.buf.WriteString(k + ": " + v + ";")
			if idx != len(d.style)-1 {
				d.buf.WriteByte(' ')
			}
			idx++
		}
		d.buf.WriteString("\">")
	} else {
		d.buf.WriteString("<dl>")
	}

	for _, content := range d.contents {
		d.buf.Write(content)
	}
	d.buf.WriteString("</dl>")
}

type Dt struct {
	buf      bytes.Buffer
	style    map[string]string
	contents [][]byte
}

func NewDt() *Dt {
	return &Dt{
		style: make(map[string]string),
	}
}

func (d *Dt) AddStyle(k, v string) *Dt {
	d.style[k] = v
	return d
}

func (d *Dt) AddStyles(m map[string]string) *Dt {
	for k, v := range m {
		d.style[k] = v
	}
	return d
}

func (d *Dt) Style(m map[string]string) *Dt {
	d.style = m
	return d
}

func (d *Dt) Add(e Elements) *Dt {
	d.contents = append(d.contents, e.Bytes())
	return d
}

func (d *Dt) Bytes() []byte {
	return d.buf.Bytes()
}

func (d *Dt) Prepare() {
	if len(d.style) != 0 {
		idx := 0
		d.buf.WriteString("<dt style=\"")
		for k, v := range d.style {
			d.buf.WriteString(k + ": " + v + ";")
			if idx != len(d.style)-1 {
				d.buf.WriteByte(' ')
			}
			idx++
		}
		d.buf.WriteString("\">")
	} else {
		d.buf.WriteString("<dt>")
	}

	for _, content := range d.contents {
		d.buf.Write(content)
	}
	d.buf.WriteString("</dt>")
}

type Dd struct {
	buf      bytes.Buffer
	style    map[string]string
	contents [][]byte
}

func NewDd() *Dd {
	return &Dd{
		style: make(map[string]string),
	}
}

func (d *Dd) AddStyle(k, v string) *Dd {
	d.style[k] = v
	return d
}

func (d *Dd) AddStyles(m map[string]string) *Dd {
	for k, v := range m {
		d.style[k] = v
	}
	return d
}

func (d *Dd) Style(m map[string]string) *Dd {
	d.style = m
	return d
}

func (d *Dd) Add(e Elements) *Dd {
	d.contents = append(d.contents, e.Bytes())
	return d
}

func (d *Dd) Bytes() []byte {
	return d.buf.Bytes()
}

func (d *Dd) Prepare() {
	if len(d.style) != 0 {
		idx := 0
		d.buf.WriteString("<dd style=\"")
		for k, v := range d.style {
			d.buf.WriteString(k + ": " + v + ";")
			if idx != len(d.style)-1 {
				d.buf.WriteByte(' ')
			}
			idx++
		}
		d.buf.WriteString("\">")
	} else {
		d.buf.WriteString("<dd>")
	}

	for _, content := range d.contents {
		d.buf.Write(content)
	}
	d.buf.WriteString("</dd>")
}

type Figure struct {
	buf      bytes.Buffer
	style    map[string]string
	contents [][]byte
}

func NewFigure() *Figure {
	return &Figure{
		style: make(map[string]string),
	}
}

func (f *Figure) AddStyle(k, v string) *Figure {
	f.style[k] = v
	return f
}

func (f *Figure) AddStyles(m map[string]string) *Figure {
	for k, v := range m {
		f.style[k] = v
	}
	return f
}

func (f *Figure) Style(m map[string]string) *Figure {
	f.style = m
	return f
}

func (f *Figure) Add(e Elements) *Figure {
	f.contents = append(f.contents, e.Bytes())
	return f
}

func (f *Figure) Bytes() []byte {
	return f.buf.Bytes()
}

func (f *Figure) Prepare() {
	if len(f.style) != 0 {
		idx := 0
		f.buf.WriteString("<figure style=\"")
		for k, v := range f.style {
			f.buf.WriteString(k + ": " + v + ";")
			if idx != len(f.style)-1 {
				f.buf.WriteByte(' ')
			}
			idx++
		}
		f.buf.WriteString("\">")
	} else {
		f.buf.WriteString("<figure>")
	}

	for _, content := range f.contents {
		f.buf.Write(content)
	}
	f.buf.WriteString("</figure>")
}

type Figcaption struct {
	buf      bytes.Buffer
	style    map[string]string
	contents [][]byte
}

func NewFigcaption() *Figcaption {
	return &Figcaption{
		style: make(map[string]string),
	}
}

func (f *Figcaption) AddStyle(k, v string) *Figcaption {
	f.style[k] = v
	return f
}

func (f *Figcaption) AddStyles(m map[string]string) *Figcaption {
	for k, v := range m {
		f.style[k] = v
	}
	return f
}

func (f *Figcaption) Style(m map[string]string) *Figcaption {
	f.style = m
	return f
}

func (f *Figcaption) Add(e Elements) *Figcaption {
	f.contents = append(f.contents, e.Bytes())
	return f
}

func (f *Figcaption) Bytes() []byte {
	return f.buf.Bytes()
}

func (f *Figcaption) Prepare() {
	if len(f.style) != 0 {
		idx := 0
		f.buf.WriteString("<figcaption style=\"")
		for k, v := range f.style {
			f.buf.WriteString(k + ": " + v + ";")
			if idx != len(f.style)-1 {
				f.buf.WriteByte(' ')
			}
			idx++
		}
		f.buf.WriteString("\">")
	} else {
		f.buf.WriteString("<figcaption>")
	}

	for _, content := range f.contents {
		f.buf.Write(content)
	}
	f.buf.WriteString("</figcaption>")
}
