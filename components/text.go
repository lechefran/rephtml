package rephtml

import "bytes"

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
