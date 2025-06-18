package rephtml

import "bytes"

type Header struct {
	buf      bytes.Buffer
	style    map[string]string
	contents [][]byte
}

func NewHeader() *Header {
	return &Header{
		style: make(map[string]string),
	}
}

func (h *Header) AddStyle(k, v string) *Header {
	h.style[k] = v
	return h
}

func (h *Header) AddStyles(m map[string]string) *Header {
	for k, v := range m {
		h.style[k] = v
	}
	return h
}

func (h *Header) Style(m map[string]string) *Header {
	h.style = m
	return h
}

func (h *Header) Add(e Elements) *Header {
	h.contents = append(h.contents, e.Bytes())
	return h
}

func (h *Header) Bytes() []byte {
	return h.buf.Bytes()
}

func (h *Header) Prepare() {
	h.buf.WriteString("<header")
	if len(h.style) != 0 {
		idx := 0
		h.buf.WriteString(" style=\"")
		for k, v := range h.style {
			h.buf.WriteString(k + ": " + v + ";")
			if idx != len(h.style)-1 {
				h.buf.WriteByte(' ')
			}
			idx++
		}
		h.buf.WriteString("\"")
	}
	h.buf.WriteByte('>')

	for _, content := range h.contents {
		h.buf.Write(content)
	}
	h.buf.WriteString("</header>")
}

type Nav struct {
	buf      bytes.Buffer
	style    map[string]string
	contents [][]byte
	role     string
}

func NewNav() *Nav {
	return &Nav{
		style: make(map[string]string),
	}
}

func (n *Nav) AddStyle(k, v string) *Nav {
	n.style[k] = v
	return n
}

func (n *Nav) AddStyles(m map[string]string) *Nav {
	for k, v := range m {
		n.style[k] = v
	}
	return n
}

func (n *Nav) Style(m map[string]string) *Nav {
	n.style = m
	return n
}

func (n *Nav) Add(e Elements) *Nav {
	n.contents = append(n.contents, e.Bytes())
	return n
}

func (n *Nav) Role(r string) *Nav {
	n.role = r
	return n
}

func (n *Nav) Bytes() []byte {
	return n.buf.Bytes()
}

func (n *Nav) Prepare() {
	n.buf.WriteString("<nav")
	if n.role != "" {
		n.buf.WriteString(" role=\"" + n.role + "\"")
	}
	if len(n.style) != 0 {
		idx := 0
		n.buf.WriteString(" style=\"")
		for k, v := range n.style {
			n.buf.WriteString(k + ": " + v + ";")
			if idx != len(n.style)-1 {
				n.buf.WriteByte(' ')
			}
			idx++
		}
		n.buf.WriteString("\"")
	}
	n.buf.WriteByte('>')

	for _, content := range n.contents {
		n.buf.Write(content)
	}
	n.buf.WriteString("</nav>")
}

type Section struct {
	buf      bytes.Buffer
	style    map[string]string
	contents [][]byte
	ariaLabel string
}

func NewSection() *Section {
	return &Section{
		style: make(map[string]string),
	}
}

func (s *Section) AddStyle(k, v string) *Section {
	s.style[k] = v
	return s
}

func (s *Section) AddStyles(m map[string]string) *Section {
	for k, v := range m {
		s.style[k] = v
	}
	return s
}

func (s *Section) Style(m map[string]string) *Section {
	s.style = m
	return s
}

func (s *Section) Add(e Elements) *Section {
	s.contents = append(s.contents, e.Bytes())
	return s
}

func (s *Section) AriaLabel(label string) *Section {
	s.ariaLabel = label
	return s
}

func (s *Section) Bytes() []byte {
	return s.buf.Bytes()
}

func (s *Section) Prepare() {
	s.buf.WriteString("<section")
	if s.ariaLabel != "" {
		s.buf.WriteString(" aria-label=\"" + s.ariaLabel + "\"")
	}
	if len(s.style) != 0 {
		idx := 0
		s.buf.WriteString(" style=\"")
		for k, v := range s.style {
			s.buf.WriteString(k + ": " + v + ";")
			if idx != len(s.style)-1 {
				s.buf.WriteByte(' ')
			}
			idx++
		}
		s.buf.WriteString("\"")
	}
	s.buf.WriteByte('>')

	for _, content := range s.contents {
		s.buf.Write(content)
	}
	s.buf.WriteString("</section>")
}

type Main struct {
	buf      bytes.Buffer
	style    map[string]string
	contents [][]byte
}

func NewMain() *Main {
	return &Main{
		style: make(map[string]string),
	}
}

func (m *Main) AddStyle(k, v string) *Main {
	m.style[k] = v
	return m
}

func (m *Main) AddStyles(mp map[string]string) *Main {
	for k, v := range mp {
		m.style[k] = v
	}
	return m
}

func (m *Main) Style(mp map[string]string) *Main {
	m.style = mp
	return m
}

func (m *Main) Add(e Elements) *Main {
	m.contents = append(m.contents, e.Bytes())
	return m
}

func (m *Main) Bytes() []byte {
	return m.buf.Bytes()
}

func (m *Main) Prepare() {
	m.buf.WriteString("<main")
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
	m.buf.WriteString("</main>")
}

type Article struct {
	buf      bytes.Buffer
	style    map[string]string
	contents [][]byte
}

func NewArticle() *Article {
	return &Article{
		style: make(map[string]string),
	}
}

func (a *Article) AddStyle(k, v string) *Article {
	a.style[k] = v
	return a
}

func (a *Article) AddStyles(mp map[string]string) *Article {
	for k, v := range mp {
		a.style[k] = v
	}
	return a
}

func (a *Article) Style(mp map[string]string) *Article {
	a.style = mp
	return a
}

func (a *Article) Add(e Elements) *Article {
	a.contents = append(a.contents, e.Bytes())
	return a
}

func (a *Article) Bytes() []byte {
	return a.buf.Bytes()
}

func (a *Article) Prepare() {
	a.buf.WriteString("<article")
	if len(a.style) != 0 {
		idx := 0
		a.buf.WriteString(" style=\"")
		for k, v := range a.style {
			a.buf.WriteString(k + ": " + v + ";")
			if idx != len(a.style)-1 {
				a.buf.WriteByte(' ')
			}
			idx++
		}
		a.buf.WriteString("\"")
	}
	a.buf.WriteByte('>')

	for _, content := range a.contents {
		a.buf.Write(content)
	}
	a.buf.WriteString("</article>")
}
