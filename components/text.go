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
