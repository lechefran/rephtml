package rephtml

// preparedElement is the internal rendering contract used by generated elements.
type preparedElement interface {
	Prepare()
	Bytes() []byte
}

// renderPrepared prepares an element and returns a snapshot of its rendered bytes.
func renderPrepared(e preparedElement) []byte {
	if e == nil {
		return nil
	}
	e.Prepare()
	return e.Bytes()
}

// htmlPrepared prepares an element and returns its rendered HTML.
func htmlPrepared(e preparedElement) string {
	return string(renderPrepared(e))
}

// Render returns raw text bytes.
func (r rawText) Render() []byte {
	return renderPrepared(r)
}

// HTML returns raw text as a string.
func (r rawText) HTML() string {
	return htmlPrepared(r)
}

// String returns raw text as a string.
func (r rawText) String() string {
	return r.HTML()
}

// Render returns escaped text bytes.
func (e escapedText) Render() []byte {
	return renderPrepared(e)
}

// HTML returns escaped text as a string.
func (e escapedText) HTML() string {
	return htmlPrepared(e)
}

// String returns escaped text as a string.
func (e escapedText) String() string {
	return e.HTML()
}

// Render returns freshly prepared Math HTML bytes.
func (m *Math) Render() []byte {
	return renderPrepared(m)
}

// HTML returns freshly prepared Math HTML as a string.
func (m *Math) HTML() string {
	return htmlPrepared(m)
}

// String returns freshly prepared Math HTML as a string.
func (m *Math) String() string {
	return m.HTML()
}

// Render returns freshly prepared Svg HTML bytes.
func (s *Svg) Render() []byte {
	return renderPrepared(s)
}

// HTML returns freshly prepared Svg HTML as a string.
func (s *Svg) HTML() string {
	return htmlPrepared(s)
}

// String returns freshly prepared Svg HTML as a string.
func (s *Svg) String() string {
	return s.HTML()
}

// Render returns freshly prepared Embed HTML bytes.
func (e *Embed) Render() []byte {
	return renderPrepared(e)
}

// HTML returns freshly prepared Embed HTML as a string.
func (e *Embed) HTML() string {
	return htmlPrepared(e)
}

// String returns freshly prepared Embed HTML as a string.
func (e *Embed) String() string {
	return e.HTML()
}

// Render returns freshly prepared Iframe HTML bytes.
func (i *Iframe) Render() []byte {
	return renderPrepared(i)
}

// HTML returns freshly prepared Iframe HTML as a string.
func (i *Iframe) HTML() string {
	return htmlPrepared(i)
}

// String returns freshly prepared Iframe HTML as a string.
func (i *Iframe) String() string {
	return i.HTML()
}

// Render returns freshly prepared Object HTML bytes.
func (o *Object) Render() []byte {
	return renderPrepared(o)
}

// HTML returns freshly prepared Object HTML as a string.
func (o *Object) HTML() string {
	return htmlPrepared(o)
}

// String returns freshly prepared Object HTML as a string.
func (o *Object) String() string {
	return o.HTML()
}

// Render returns freshly prepared Picture HTML bytes.
func (p *Picture) Render() []byte {
	return renderPrepared(p)
}

// HTML returns freshly prepared Picture HTML as a string.
func (p *Picture) HTML() string {
	return htmlPrepared(p)
}

// String returns freshly prepared Picture HTML as a string.
func (p *Picture) String() string {
	return p.HTML()
}

// Render returns freshly prepared Portal HTML bytes.
func (p *Portal) Render() []byte {
	return renderPrepared(p)
}

// HTML returns freshly prepared Portal HTML as a string.
func (p *Portal) HTML() string {
	return htmlPrepared(p)
}

// String returns freshly prepared Portal HTML as a string.
func (p *Portal) String() string {
	return p.HTML()
}

// Render returns freshly prepared Source HTML bytes.
func (s *Source) Render() []byte {
	return renderPrepared(s)
}

// HTML returns freshly prepared Source HTML as a string.
func (s *Source) HTML() string {
	return htmlPrepared(s)
}

// String returns freshly prepared Source HTML as a string.
func (s *Source) String() string {
	return s.HTML()
}

// Render returns freshly prepared Canvas HTML bytes.
func (c *Canvas) Render() []byte {
	return renderPrepared(c)
}

// HTML returns freshly prepared Canvas HTML as a string.
func (c *Canvas) HTML() string {
	return htmlPrepared(c)
}

// String returns freshly prepared Canvas HTML as a string.
func (c *Canvas) String() string {
	return c.HTML()
}

// Render returns freshly prepared Noscript HTML bytes.
func (n *Noscript) Render() []byte {
	return renderPrepared(n)
}

// HTML returns freshly prepared Noscript HTML as a string.
func (n *Noscript) HTML() string {
	return htmlPrepared(n)
}

// String returns freshly prepared Noscript HTML as a string.
func (n *Noscript) String() string {
	return n.HTML()
}

// Render returns freshly prepared Script HTML bytes.
func (s *Script) Render() []byte {
	return renderPrepared(s)
}

// HTML returns freshly prepared Script HTML as a string.
func (s *Script) HTML() string {
	return htmlPrepared(s)
}

// String returns freshly prepared Script HTML as a string.
func (s *Script) String() string {
	return s.HTML()
}

// Render returns freshly prepared Area HTML bytes.
func (a *Area) Render() []byte {
	return renderPrepared(a)
}

// HTML returns freshly prepared Area HTML as a string.
func (a *Area) HTML() string {
	return htmlPrepared(a)
}

// String returns freshly prepared Area HTML as a string.
func (a *Area) String() string {
	return a.HTML()
}

// Render returns freshly prepared Img HTML bytes.
func (i *Img) Render() []byte {
	return renderPrepared(i)
}

// HTML returns freshly prepared Img HTML as a string.
func (i *Img) HTML() string {
	return htmlPrepared(i)
}

// String returns freshly prepared Img HTML as a string.
func (i *Img) String() string {
	return i.HTML()
}

// Render returns freshly prepared Audio HTML bytes.
func (a *Audio) Render() []byte {
	return renderPrepared(a)
}

// HTML returns freshly prepared Audio HTML as a string.
func (a *Audio) HTML() string {
	return htmlPrepared(a)
}

// String returns freshly prepared Audio HTML as a string.
func (a *Audio) String() string {
	return a.HTML()
}

// Render returns freshly prepared Track HTML bytes.
func (t *Track) Render() []byte {
	return renderPrepared(t)
}

// HTML returns freshly prepared Track HTML as a string.
func (t *Track) HTML() string {
	return htmlPrepared(t)
}

// String returns freshly prepared Track HTML as a string.
func (t *Track) String() string {
	return t.HTML()
}

// Render returns freshly prepared Map HTML bytes.
func (m *Map) Render() []byte {
	return renderPrepared(m)
}

// HTML returns freshly prepared Map HTML as a string.
func (m *Map) HTML() string {
	return htmlPrepared(m)
}

// String returns freshly prepared Map HTML as a string.
func (m *Map) String() string {
	return m.HTML()
}

// Render returns freshly prepared Video HTML bytes.
func (v *Video) Render() []byte {
	return renderPrepared(v)
}

// HTML returns freshly prepared Video HTML as a string.
func (v *Video) HTML() string {
	return htmlPrepared(v)
}

// String returns freshly prepared Video HTML as a string.
func (v *Video) String() string {
	return v.HTML()
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

// Render returns freshly prepared P HTML bytes.
func (p *P) Render() []byte {
	return renderPrepared(p)
}

// HTML returns freshly prepared P HTML as a string.
func (p *P) HTML() string {
	return htmlPrepared(p)
}

// String returns freshly prepared P HTML as a string.
func (p *P) String() string {
	return p.HTML()
}

// Render returns freshly prepared Comment HTML bytes.
func (c *Comment) Render() []byte {
	return renderPrepared(c)
}

// HTML returns freshly prepared Comment HTML as a string.
func (c *Comment) HTML() string {
	return htmlPrepared(c)
}

// String returns freshly prepared Comment HTML as a string.
func (c *Comment) String() string {
	return c.HTML()
}

// Render returns freshly prepared Hr HTML bytes.
func (h *Hr) Render() []byte {
	return renderPrepared(h)
}

// HTML returns freshly prepared Hr HTML as a string.
func (h *Hr) HTML() string {
	return htmlPrepared(h)
}

// String returns freshly prepared Hr HTML as a string.
func (h *Hr) String() string {
	return h.HTML()
}

// Render returns freshly prepared Pre HTML bytes.
func (p *Pre) Render() []byte {
	return renderPrepared(p)
}

// HTML returns freshly prepared Pre HTML as a string.
func (p *Pre) HTML() string {
	return htmlPrepared(p)
}

// String returns freshly prepared Pre HTML as a string.
func (p *Pre) String() string {
	return p.HTML()
}

// Render returns freshly prepared Blockquote HTML bytes.
func (b *Blockquote) Render() []byte {
	return renderPrepared(b)
}

// HTML returns freshly prepared Blockquote HTML as a string.
func (b *Blockquote) HTML() string {
	return htmlPrepared(b)
}

// String returns freshly prepared Blockquote HTML as a string.
func (b *Blockquote) String() string {
	return b.HTML()
}

// Render returns freshly prepared Menu HTML bytes.
func (m *Menu) Render() []byte {
	return renderPrepared(m)
}

// HTML returns freshly prepared Menu HTML as a string.
func (m *Menu) HTML() string {
	return htmlPrepared(m)
}

// String returns freshly prepared Menu HTML as a string.
func (m *Menu) String() string {
	return m.HTML()
}

// Render returns freshly prepared Ol HTML bytes.
func (o *Ol) Render() []byte {
	return renderPrepared(o)
}

// HTML returns freshly prepared Ol HTML as a string.
func (o *Ol) HTML() string {
	return htmlPrepared(o)
}

// String returns freshly prepared Ol HTML as a string.
func (o *Ol) String() string {
	return o.HTML()
}

// Render returns freshly prepared Ul HTML bytes.
func (u *Ul) Render() []byte {
	return renderPrepared(u)
}

// HTML returns freshly prepared Ul HTML as a string.
func (u *Ul) HTML() string {
	return htmlPrepared(u)
}

// String returns freshly prepared Ul HTML as a string.
func (u *Ul) String() string {
	return u.HTML()
}

// Render returns freshly prepared Li HTML bytes.
func (l *Li) Render() []byte {
	return renderPrepared(l)
}

// HTML returns freshly prepared Li HTML as a string.
func (l *Li) HTML() string {
	return htmlPrepared(l)
}

// String returns freshly prepared Li HTML as a string.
func (l *Li) String() string {
	return l.HTML()
}

// Render returns freshly prepared Dl HTML bytes.
func (d *Dl) Render() []byte {
	return renderPrepared(d)
}

// HTML returns freshly prepared Dl HTML as a string.
func (d *Dl) HTML() string {
	return htmlPrepared(d)
}

// String returns freshly prepared Dl HTML as a string.
func (d *Dl) String() string {
	return d.HTML()
}

// Render returns freshly prepared Dt HTML bytes.
func (d *Dt) Render() []byte {
	return renderPrepared(d)
}

// HTML returns freshly prepared Dt HTML as a string.
func (d *Dt) HTML() string {
	return htmlPrepared(d)
}

// String returns freshly prepared Dt HTML as a string.
func (d *Dt) String() string {
	return d.HTML()
}

// Render returns freshly prepared Dd HTML bytes.
func (d *Dd) Render() []byte {
	return renderPrepared(d)
}

// HTML returns freshly prepared Dd HTML as a string.
func (d *Dd) HTML() string {
	return htmlPrepared(d)
}

// String returns freshly prepared Dd HTML as a string.
func (d *Dd) String() string {
	return d.HTML()
}

// Render returns freshly prepared Figure HTML bytes.
func (f *Figure) Render() []byte {
	return renderPrepared(f)
}

// HTML returns freshly prepared Figure HTML as a string.
func (f *Figure) HTML() string {
	return htmlPrepared(f)
}

// String returns freshly prepared Figure HTML as a string.
func (f *Figure) String() string {
	return f.HTML()
}

// Render returns freshly prepared Figcaption HTML bytes.
func (f *Figcaption) Render() []byte {
	return renderPrepared(f)
}

// HTML returns freshly prepared Figcaption HTML as a string.
func (f *Figcaption) HTML() string {
	return htmlPrepared(f)
}

// String returns freshly prepared Figcaption HTML as a string.
func (f *Figcaption) String() string {
	return f.HTML()
}

// Render returns freshly prepared Search HTML bytes.
func (s *Search) Render() []byte {
	return renderPrepared(s)
}

// HTML returns freshly prepared Search HTML as a string.
func (s *Search) HTML() string {
	return htmlPrepared(s)
}

// String returns freshly prepared Search HTML as a string.
func (s *Search) String() string {
	return s.HTML()
}

// Render returns freshly prepared Table HTML bytes.
func (t *Table) Render() []byte {
	return renderPrepared(t)
}

// HTML returns freshly prepared Table HTML as a string.
func (t *Table) HTML() string {
	return htmlPrepared(t)
}

// String returns freshly prepared Table HTML as a string.
func (t *Table) String() string {
	return t.HTML()
}

// Render returns freshly prepared Thead HTML bytes.
func (t *Thead) Render() []byte {
	return renderPrepared(t)
}

// HTML returns freshly prepared Thead HTML as a string.
func (t *Thead) HTML() string {
	return htmlPrepared(t)
}

// String returns freshly prepared Thead HTML as a string.
func (t *Thead) String() string {
	return t.HTML()
}

// Render returns freshly prepared Tbody HTML bytes.
func (t *Tbody) Render() []byte {
	return renderPrepared(t)
}

// HTML returns freshly prepared Tbody HTML as a string.
func (t *Tbody) HTML() string {
	return htmlPrepared(t)
}

// String returns freshly prepared Tbody HTML as a string.
func (t *Tbody) String() string {
	return t.HTML()
}

// Render returns freshly prepared Tfoot HTML bytes.
func (t *Tfoot) Render() []byte {
	return renderPrepared(t)
}

// HTML returns freshly prepared Tfoot HTML as a string.
func (t *Tfoot) HTML() string {
	return htmlPrepared(t)
}

// String returns freshly prepared Tfoot HTML as a string.
func (t *Tfoot) String() string {
	return t.HTML()
}

// Render returns freshly prepared Caption HTML bytes.
func (c *Caption) Render() []byte {
	return renderPrepared(c)
}

// HTML returns freshly prepared Caption HTML as a string.
func (c *Caption) HTML() string {
	return htmlPrepared(c)
}

// String returns freshly prepared Caption HTML as a string.
func (c *Caption) String() string {
	return c.HTML()
}

// Render returns freshly prepared Col HTML bytes.
func (c *Col) Render() []byte {
	return renderPrepared(c)
}

// HTML returns freshly prepared Col HTML as a string.
func (c *Col) HTML() string {
	return htmlPrepared(c)
}

// String returns freshly prepared Col HTML as a string.
func (c *Col) String() string {
	return c.HTML()
}

// Render returns freshly prepared Colgroup HTML bytes.
func (c *Colgroup) Render() []byte {
	return renderPrepared(c)
}

// HTML returns freshly prepared Colgroup HTML as a string.
func (c *Colgroup) HTML() string {
	return htmlPrepared(c)
}

// String returns freshly prepared Colgroup HTML as a string.
func (c *Colgroup) String() string {
	return c.HTML()
}

// Render returns freshly prepared Tr HTML bytes.
func (t *Tr) Render() []byte {
	return renderPrepared(t)
}

// HTML returns freshly prepared Tr HTML as a string.
func (t *Tr) HTML() string {
	return htmlPrepared(t)
}

// String returns freshly prepared Tr HTML as a string.
func (t *Tr) String() string {
	return t.HTML()
}

// Render returns freshly prepared Td HTML bytes.
func (t *Td) Render() []byte {
	return renderPrepared(t)
}

// HTML returns freshly prepared Td HTML as a string.
func (t *Td) HTML() string {
	return htmlPrepared(t)
}

// String returns freshly prepared Td HTML as a string.
func (t *Td) String() string {
	return t.HTML()
}

// Render returns freshly prepared Th HTML bytes.
func (t *Th) Render() []byte {
	return renderPrepared(t)
}

// HTML returns freshly prepared Th HTML as a string.
func (t *Th) HTML() string {
	return htmlPrepared(t)
}

// String returns freshly prepared Th HTML as a string.
func (t *Th) String() string {
	return t.HTML()
}

// Render returns freshly prepared Head HTML bytes.
func (h *Head) Render() []byte {
	return renderPrepared(h)
}

// HTML returns freshly prepared Head HTML as a string.
func (h *Head) HTML() string {
	return htmlPrepared(h)
}

// String returns freshly prepared Head HTML as a string.
func (h *Head) String() string {
	return h.HTML()
}

// Render returns freshly prepared Body HTML bytes.
func (b *Body) Render() []byte {
	return renderPrepared(b)
}

// HTML returns freshly prepared Body HTML as a string.
func (b *Body) HTML() string {
	return htmlPrepared(b)
}

// String returns freshly prepared Body HTML as a string.
func (b *Body) String() string {
	return b.HTML()
}

// Render returns freshly prepared Title HTML bytes.
func (t *Title) Render() []byte {
	return renderPrepared(t)
}

// HTML returns freshly prepared Title HTML as a string.
func (t *Title) HTML() string {
	return htmlPrepared(t)
}

// String returns freshly prepared Title HTML as a string.
func (t *Title) String() string {
	return t.HTML()
}

// Render returns freshly prepared Base HTML bytes.
func (b *Base) Render() []byte {
	return renderPrepared(b)
}

// HTML returns freshly prepared Base HTML as a string.
func (b *Base) HTML() string {
	return htmlPrepared(b)
}

// String returns freshly prepared Base HTML as a string.
func (b *Base) String() string {
	return b.HTML()
}

// Render returns freshly prepared Link HTML bytes.
func (l *Link) Render() []byte {
	return renderPrepared(l)
}

// HTML returns freshly prepared Link HTML as a string.
func (l *Link) HTML() string {
	return htmlPrepared(l)
}

// String returns freshly prepared Link HTML as a string.
func (l *Link) String() string {
	return l.HTML()
}

// Render returns freshly prepared Meta HTML bytes.
func (m *Meta) Render() []byte {
	return renderPrepared(m)
}

// HTML returns freshly prepared Meta HTML as a string.
func (m *Meta) HTML() string {
	return htmlPrepared(m)
}

// String returns freshly prepared Meta HTML as a string.
func (m *Meta) String() string {
	return m.HTML()
}

// Render returns freshly prepared StyleElement HTML bytes.
func (s *StyleElement) Render() []byte {
	return renderPrepared(s)
}

// HTML returns freshly prepared StyleElement HTML as a string.
func (s *StyleElement) HTML() string {
	return htmlPrepared(s)
}

// String returns freshly prepared StyleElement HTML as a string.
func (s *StyleElement) String() string {
	return s.HTML()
}

// Render returns freshly prepared Hgroup HTML bytes.
func (h *Hgroup) Render() []byte {
	return renderPrepared(h)
}

// HTML returns freshly prepared Hgroup HTML as a string.
func (h *Hgroup) HTML() string {
	return htmlPrepared(h)
}

// String returns freshly prepared Hgroup HTML as a string.
func (h *Hgroup) String() string {
	return h.HTML()
}

// Render returns freshly prepared H1 HTML bytes.
func (h *H1) Render() []byte {
	return renderPrepared(h)
}

// HTML returns freshly prepared H1 HTML as a string.
func (h *H1) HTML() string {
	return htmlPrepared(h)
}

// String returns freshly prepared H1 HTML as a string.
func (h *H1) String() string {
	return h.HTML()
}

// Render returns freshly prepared H2 HTML bytes.
func (h *H2) Render() []byte {
	return renderPrepared(h)
}

// HTML returns freshly prepared H2 HTML as a string.
func (h *H2) HTML() string {
	return htmlPrepared(h)
}

// String returns freshly prepared H2 HTML as a string.
func (h *H2) String() string {
	return h.HTML()
}

// Render returns freshly prepared H3 HTML bytes.
func (h *H3) Render() []byte {
	return renderPrepared(h)
}

// HTML returns freshly prepared H3 HTML as a string.
func (h *H3) HTML() string {
	return htmlPrepared(h)
}

// String returns freshly prepared H3 HTML as a string.
func (h *H3) String() string {
	return h.HTML()
}

// Render returns freshly prepared H4 HTML bytes.
func (h *H4) Render() []byte {
	return renderPrepared(h)
}

// HTML returns freshly prepared H4 HTML as a string.
func (h *H4) HTML() string {
	return htmlPrepared(h)
}

// String returns freshly prepared H4 HTML as a string.
func (h *H4) String() string {
	return h.HTML()
}

// Render returns freshly prepared H5 HTML bytes.
func (h *H5) Render() []byte {
	return renderPrepared(h)
}

// HTML returns freshly prepared H5 HTML as a string.
func (h *H5) HTML() string {
	return htmlPrepared(h)
}

// String returns freshly prepared H5 HTML as a string.
func (h *H5) String() string {
	return h.HTML()
}

// Render returns freshly prepared H6 HTML bytes.
func (h *H6) Render() []byte {
	return renderPrepared(h)
}

// HTML returns freshly prepared H6 HTML as a string.
func (h *H6) HTML() string {
	return htmlPrepared(h)
}

// String returns freshly prepared H6 HTML as a string.
func (h *H6) String() string {
	return h.HTML()
}

// Render returns freshly prepared Div HTML bytes.
func (d *Div) Render() []byte {
	return renderPrepared(d)
}

// HTML returns freshly prepared Div HTML as a string.
func (d *Div) HTML() string {
	return htmlPrepared(d)
}

// String returns freshly prepared Div HTML as a string.
func (d *Div) String() string {
	return d.HTML()
}

// Render returns freshly prepared Style HTML bytes.
func (s *Style) Render() []byte {
	return renderPrepared(s)
}

// HTML returns freshly prepared Style HTML as a string.
func (s *Style) HTML() string {
	return htmlPrepared(s)
}

// String returns freshly prepared Style HTML as a string.
func (s *Style) String() string {
	return s.HTML()
}

// Render returns freshly prepared StyleRule HTML bytes.
func (s *StyleRule) Render() []byte {
	return renderPrepared(s)
}

// HTML returns freshly prepared StyleRule HTML as a string.
func (s *StyleRule) HTML() string {
	return htmlPrepared(s)
}

// String returns freshly prepared StyleRule HTML as a string.
func (s *StyleRule) String() string {
	return s.HTML()
}

// Render returns freshly prepared RawCSSRule HTML bytes.
func (r *RawCSSRule) Render() []byte {
	return renderPrepared(r)
}

// HTML returns freshly prepared RawCSSRule HTML as a string.
func (r *RawCSSRule) HTML() string {
	return htmlPrepared(r)
}

// String returns freshly prepared RawCSSRule HTML as a string.
func (r *RawCSSRule) String() string {
	return r.HTML()
}

// Render returns freshly prepared CharsetRule HTML bytes.
func (c *CharsetRule) Render() []byte {
	return renderPrepared(c)
}

// HTML returns freshly prepared CharsetRule HTML as a string.
func (c *CharsetRule) HTML() string {
	return htmlPrepared(c)
}

// String returns freshly prepared CharsetRule HTML as a string.
func (c *CharsetRule) String() string {
	return c.HTML()
}

// Render returns freshly prepared ImportRule HTML bytes.
func (i *ImportRule) Render() []byte {
	return renderPrepared(i)
}

// HTML returns freshly prepared ImportRule HTML as a string.
func (i *ImportRule) HTML() string {
	return htmlPrepared(i)
}

// String returns freshly prepared ImportRule HTML as a string.
func (i *ImportRule) String() string {
	return i.HTML()
}

// Render returns freshly prepared FontFaceRule HTML bytes.
func (f *FontFaceRule) Render() []byte {
	return renderPrepared(f)
}

// HTML returns freshly prepared FontFaceRule HTML as a string.
func (f *FontFaceRule) HTML() string {
	return htmlPrepared(f)
}

// String returns freshly prepared FontFaceRule HTML as a string.
func (f *FontFaceRule) String() string {
	return f.HTML()
}

// Render returns freshly prepared FontFeatureValuesRule HTML bytes.
func (f *FontFeatureValuesRule) Render() []byte {
	return renderPrepared(f)
}

// HTML returns freshly prepared FontFeatureValuesRule HTML as a string.
func (f *FontFeatureValuesRule) HTML() string {
	return htmlPrepared(f)
}

// String returns freshly prepared FontFeatureValuesRule HTML as a string.
func (f *FontFeatureValuesRule) String() string {
	return f.HTML()
}

// Render returns freshly prepared MediaRule HTML bytes.
func (m *MediaRule) Render() []byte {
	return renderPrepared(m)
}

// HTML returns freshly prepared MediaRule HTML as a string.
func (m *MediaRule) HTML() string {
	return htmlPrepared(m)
}

// String returns freshly prepared MediaRule HTML as a string.
func (m *MediaRule) String() string {
	return m.HTML()
}

// Render returns freshly prepared KeyframeBlock HTML bytes.
func (k *KeyframeBlock) Render() []byte {
	return renderPrepared(k)
}

// HTML returns freshly prepared KeyframeBlock HTML as a string.
func (k *KeyframeBlock) HTML() string {
	return htmlPrepared(k)
}

// String returns freshly prepared KeyframeBlock HTML as a string.
func (k *KeyframeBlock) String() string {
	return k.HTML()
}

// Render returns freshly prepared KeyframesRule HTML bytes.
func (k *KeyframesRule) Render() []byte {
	return renderPrepared(k)
}

// HTML returns freshly prepared KeyframesRule HTML as a string.
func (k *KeyframesRule) HTML() string {
	return htmlPrepared(k)
}

// String returns freshly prepared KeyframesRule HTML as a string.
func (k *KeyframesRule) String() string {
	return k.HTML()
}

// Render returns freshly prepared Del HTML bytes.
func (d *Del) Render() []byte {
	return renderPrepared(d)
}

// HTML returns freshly prepared Del HTML as a string.
func (d *Del) HTML() string {
	return htmlPrepared(d)
}

// String returns freshly prepared Del HTML as a string.
func (d *Del) String() string {
	return d.HTML()
}

// Render returns freshly prepared Ins HTML bytes.
func (i *Ins) Render() []byte {
	return renderPrepared(i)
}

// HTML returns freshly prepared Ins HTML as a string.
func (i *Ins) HTML() string {
	return htmlPrepared(i)
}

// String returns freshly prepared Ins HTML as a string.
func (i *Ins) String() string {
	return i.HTML()
}

// Render returns freshly prepared Slot HTML bytes.
func (s *Slot) Render() []byte {
	return renderPrepared(s)
}

// HTML returns freshly prepared Slot HTML as a string.
func (s *Slot) HTML() string {
	return htmlPrepared(s)
}

// String returns freshly prepared Slot HTML as a string.
func (s *Slot) String() string {
	return s.HTML()
}

// Render returns freshly prepared Template HTML bytes.
func (t *Template) Render() []byte {
	return renderPrepared(t)
}

// HTML returns freshly prepared Template HTML as a string.
func (t *Template) HTML() string {
	return htmlPrepared(t)
}

// String returns freshly prepared Template HTML as a string.
func (t *Template) String() string {
	return t.HTML()
}

// Render returns freshly prepared Header HTML bytes.
func (h *Header) Render() []byte {
	return renderPrepared(h)
}

// HTML returns freshly prepared Header HTML as a string.
func (h *Header) HTML() string {
	return htmlPrepared(h)
}

// String returns freshly prepared Header HTML as a string.
func (h *Header) String() string {
	return h.HTML()
}

// Render returns freshly prepared Nav HTML bytes.
func (n *Nav) Render() []byte {
	return renderPrepared(n)
}

// HTML returns freshly prepared Nav HTML as a string.
func (n *Nav) HTML() string {
	return htmlPrepared(n)
}

// String returns freshly prepared Nav HTML as a string.
func (n *Nav) String() string {
	return n.HTML()
}

// Render returns freshly prepared Main HTML bytes.
func (m *Main) Render() []byte {
	return renderPrepared(m)
}

// HTML returns freshly prepared Main HTML as a string.
func (m *Main) HTML() string {
	return htmlPrepared(m)
}

// String returns freshly prepared Main HTML as a string.
func (m *Main) String() string {
	return m.HTML()
}

// Render returns freshly prepared Section HTML bytes.
func (s *Section) Render() []byte {
	return renderPrepared(s)
}

// HTML returns freshly prepared Section HTML as a string.
func (s *Section) HTML() string {
	return htmlPrepared(s)
}

// String returns freshly prepared Section HTML as a string.
func (s *Section) String() string {
	return s.HTML()
}

// Render returns freshly prepared Article HTML bytes.
func (a *Article) Render() []byte {
	return renderPrepared(a)
}

// HTML returns freshly prepared Article HTML as a string.
func (a *Article) HTML() string {
	return htmlPrepared(a)
}

// String returns freshly prepared Article HTML as a string.
func (a *Article) String() string {
	return a.HTML()
}

// Render returns freshly prepared Aside HTML bytes.
func (a *Aside) Render() []byte {
	return renderPrepared(a)
}

// HTML returns freshly prepared Aside HTML as a string.
func (a *Aside) HTML() string {
	return htmlPrepared(a)
}

// String returns freshly prepared Aside HTML as a string.
func (a *Aside) String() string {
	return a.HTML()
}

// Render returns freshly prepared Footer HTML bytes.
func (f *Footer) Render() []byte {
	return renderPrepared(f)
}

// HTML returns freshly prepared Footer HTML as a string.
func (f *Footer) HTML() string {
	return htmlPrepared(f)
}

// String returns freshly prepared Footer HTML as a string.
func (f *Footer) String() string {
	return f.HTML()
}

// Render returns freshly prepared Address HTML bytes.
func (a *Address) Render() []byte {
	return renderPrepared(a)
}

// HTML returns freshly prepared Address HTML as a string.
func (a *Address) HTML() string {
	return htmlPrepared(a)
}

// String returns freshly prepared Address HTML as a string.
func (a *Address) String() string {
	return a.HTML()
}

// Render returns freshly prepared Details HTML bytes.
func (d *Details) Render() []byte {
	return renderPrepared(d)
}

// HTML returns freshly prepared Details HTML as a string.
func (d *Details) HTML() string {
	return htmlPrepared(d)
}

// String returns freshly prepared Details HTML as a string.
func (d *Details) String() string {
	return d.HTML()
}

// Render returns freshly prepared Dialog HTML bytes.
func (d *Dialog) Render() []byte {
	return renderPrepared(d)
}

// HTML returns freshly prepared Dialog HTML as a string.
func (d *Dialog) HTML() string {
	return htmlPrepared(d)
}

// String returns freshly prepared Dialog HTML as a string.
func (d *Dialog) String() string {
	return d.HTML()
}

// Render returns freshly prepared Summary HTML bytes.
func (s *Summary) Render() []byte {
	return renderPrepared(s)
}

// HTML returns freshly prepared Summary HTML as a string.
func (s *Summary) HTML() string {
	return htmlPrepared(s)
}

// String returns freshly prepared Summary HTML as a string.
func (s *Summary) String() string {
	return s.HTML()
}

// Render returns freshly prepared Form HTML bytes.
func (f *Form) Render() []byte {
	return renderPrepared(f)
}

// HTML returns freshly prepared Form HTML as a string.
func (f *Form) HTML() string {
	return htmlPrepared(f)
}

// String returns freshly prepared Form HTML as a string.
func (f *Form) String() string {
	return f.HTML()
}

// Render returns freshly prepared Label HTML bytes.
func (l *Label) Render() []byte {
	return renderPrepared(l)
}

// HTML returns freshly prepared Label HTML as a string.
func (l *Label) HTML() string {
	return htmlPrepared(l)
}

// String returns freshly prepared Label HTML as a string.
func (l *Label) String() string {
	return l.HTML()
}

// Render returns freshly prepared Input HTML bytes.
func (i *Input) Render() []byte {
	return renderPrepared(i)
}

// HTML returns freshly prepared Input HTML as a string.
func (i *Input) HTML() string {
	return htmlPrepared(i)
}

// String returns freshly prepared Input HTML as a string.
func (i *Input) String() string {
	return i.HTML()
}

// Render returns freshly prepared Output HTML bytes.
func (o *Output) Render() []byte {
	return renderPrepared(o)
}

// HTML returns freshly prepared Output HTML as a string.
func (o *Output) HTML() string {
	return htmlPrepared(o)
}

// String returns freshly prepared Output HTML as a string.
func (o *Output) String() string {
	return o.HTML()
}

// Render returns freshly prepared Fieldset HTML bytes.
func (f *Fieldset) Render() []byte {
	return renderPrepared(f)
}

// HTML returns freshly prepared Fieldset HTML as a string.
func (f *Fieldset) HTML() string {
	return htmlPrepared(f)
}

// String returns freshly prepared Fieldset HTML as a string.
func (f *Fieldset) String() string {
	return f.HTML()
}

// Render returns freshly prepared Button HTML bytes.
func (b *Button) Render() []byte {
	return renderPrepared(b)
}

// HTML returns freshly prepared Button HTML as a string.
func (b *Button) HTML() string {
	return htmlPrepared(b)
}

// String returns freshly prepared Button HTML as a string.
func (b *Button) String() string {
	return b.HTML()
}

// Render returns freshly prepared Select HTML bytes.
func (s *Select) Render() []byte {
	return renderPrepared(s)
}

// HTML returns freshly prepared Select HTML as a string.
func (s *Select) HTML() string {
	return htmlPrepared(s)
}

// String returns freshly prepared Select HTML as a string.
func (s *Select) String() string {
	return s.HTML()
}

// Render returns freshly prepared Datalist HTML bytes.
func (d *Datalist) Render() []byte {
	return renderPrepared(d)
}

// HTML returns freshly prepared Datalist HTML as a string.
func (d *Datalist) HTML() string {
	return htmlPrepared(d)
}

// String returns freshly prepared Datalist HTML as a string.
func (d *Datalist) String() string {
	return d.HTML()
}

// Render returns freshly prepared Optgroup HTML bytes.
func (o *Optgroup) Render() []byte {
	return renderPrepared(o)
}

// HTML returns freshly prepared Optgroup HTML as a string.
func (o *Optgroup) HTML() string {
	return htmlPrepared(o)
}

// String returns freshly prepared Optgroup HTML as a string.
func (o *Optgroup) String() string {
	return o.HTML()
}

// Render returns freshly prepared Option HTML bytes.
func (o *Option) Render() []byte {
	return renderPrepared(o)
}

// HTML returns freshly prepared Option HTML as a string.
func (o *Option) HTML() string {
	return htmlPrepared(o)
}

// String returns freshly prepared Option HTML as a string.
func (o *Option) String() string {
	return o.HTML()
}

// Render returns freshly prepared Textarea HTML bytes.
func (t *Textarea) Render() []byte {
	return renderPrepared(t)
}

// HTML returns freshly prepared Textarea HTML as a string.
func (t *Textarea) HTML() string {
	return htmlPrepared(t)
}

// String returns freshly prepared Textarea HTML as a string.
func (t *Textarea) String() string {
	return t.HTML()
}

// Render returns freshly prepared Progress HTML bytes.
func (p *Progress) Render() []byte {
	return renderPrepared(p)
}

// HTML returns freshly prepared Progress HTML as a string.
func (p *Progress) HTML() string {
	return htmlPrepared(p)
}

// String returns freshly prepared Progress HTML as a string.
func (p *Progress) String() string {
	return p.HTML()
}

// Render returns freshly prepared Meter HTML bytes.
func (m *Meter) Render() []byte {
	return renderPrepared(m)
}

// HTML returns freshly prepared Meter HTML as a string.
func (m *Meter) HTML() string {
	return htmlPrepared(m)
}

// String returns freshly prepared Meter HTML as a string.
func (m *Meter) String() string {
	return m.HTML()
}

// Render returns freshly prepared Legend HTML bytes.
func (l *Legend) Render() []byte {
	return renderPrepared(l)
}

// HTML returns freshly prepared Legend HTML as a string.
func (l *Legend) HTML() string {
	return htmlPrepared(l)
}

// String returns freshly prepared Legend HTML as a string.
func (l *Legend) String() string {
	return l.HTML()
}
