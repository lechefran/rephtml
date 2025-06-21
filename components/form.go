package rephtml

import "bytes"

// Form represents the HTML form element for user input
type Form struct {
	buf              bytes.Buffer
	style            map[string]string
	contents         [][]byte
	ttrack           int
	action           string
	method           string
	enctype          string
	name             string
	target           string
	autocomplete     string
	novalidate       bool
	acceptcharset    string
}

// NewForm creates a new Form element
func NewForm() *Form {
	return &Form{
		style: make(map[string]string),
	}
}

// Bytes returns the buffer contents
func (f *Form) Bytes() []byte {
	return f.buf.Bytes()
}

// Prepare builds the HTML for the form element
func (f *Form) Prepare() {
	f.buf.WriteString("<form")
	
	if f.action != "" {
		f.buf.WriteString(" action=\"" + f.action + "\"")
	}
	
	if f.method != "" {
		f.buf.WriteString(" method=\"" + f.method + "\"")
	}
	
	if f.enctype != "" {
		f.buf.WriteString(" enctype=\"" + f.enctype + "\"")
	}
	
	if f.name != "" {
		f.buf.WriteString(" name=\"" + f.name + "\"")
	}
	
	if f.target != "" {
		f.buf.WriteString(" target=\"" + f.target + "\"")
	}
	
	if f.autocomplete != "" {
		f.buf.WriteString(" autocomplete=\"" + f.autocomplete + "\"")
	}
	
	if f.acceptcharset != "" {
		f.buf.WriteString(" accept-charset=\"" + f.acceptcharset + "\"")
	}
	
	if f.novalidate {
		f.buf.WriteString(" novalidate")
	}
	
	if len(f.style) != 0 {
		idx := 0
		f.buf.WriteString(" style=\"")
		for k, v := range f.style {
			f.buf.WriteString(k + ": " + v + ";")
			if idx != len(f.style)-1 {
				f.buf.WriteByte(' ')
			}
			idx++
		}
		f.buf.WriteString("\"")
	}
	
	f.buf.WriteByte('>')
	
	for _, content := range f.contents {
		f.buf.Write(content)
	}
	
	f.buf.WriteString("</form>")
}

// Add adds content to the form element
func (f *Form) Add(e Elements) *Form {
	if e != nil {
		e.Prepare()
		f.contents = append(f.contents, e.Bytes())
	}
	return f
}

// Action sets the action attribute
func (f *Form) Action(action string) *Form {
	f.action = action
	return f
}

// Method sets the method attribute
func (f *Form) Method(method string) *Form {
	f.method = method
	return f
}

// Enctype sets the enctype attribute
func (f *Form) Enctype(enctype string) *Form {
	f.enctype = enctype
	return f
}

// Name sets the name attribute
func (f *Form) Name(name string) *Form {
	f.name = name
	return f
}

// Target sets the target attribute
func (f *Form) Target(target string) *Form {
	f.target = target
	return f
}

// Autocomplete sets the autocomplete attribute
func (f *Form) Autocomplete(autocomplete string) *Form {
	f.autocomplete = autocomplete
	return f
}

// AcceptCharset sets the accept-charset attribute
func (f *Form) AcceptCharset(acceptcharset string) *Form {
	f.acceptcharset = acceptcharset
	return f
}

// Novalidate sets the novalidate attribute
func (f *Form) Novalidate(novalidate bool) *Form {
	f.novalidate = novalidate
	return f
}

// AddStyle adds a single CSS property
func (f *Form) AddStyle(k, v string) *Form {
	f.style[k] = v
	return f
}

// AddStyles adds multiple CSS properties
func (f *Form) AddStyles(m map[string]string) *Form {
	for k, v := range m {
		f.style[k] = v
	}
	return f
}

// Style replaces all styles
func (f *Form) Style(m map[string]string) *Form {
	f.style = make(map[string]string)
	for k, v := range m {
		f.style[k] = v
	}
	return f
}

// Label represents the HTML label element for form controls
type Label struct {
	buf      bytes.Buffer
	style    map[string]string
	contents [][]byte
	ttrack   int
	forattr  string
	form     string
}

// NewLabel creates a new Label element
func NewLabel() *Label {
	return &Label{
		style: make(map[string]string),
	}
}

// Bytes returns the buffer contents
func (l *Label) Bytes() []byte {
	return l.buf.Bytes()
}

// Prepare builds the HTML for the label element
func (l *Label) Prepare() {
	l.buf.WriteString("<label")
	
	if l.forattr != "" {
		l.buf.WriteString(" for=\"" + l.forattr + "\"")
	}
	
	if l.form != "" {
		l.buf.WriteString(" form=\"" + l.form + "\"")
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
	
	l.buf.WriteString("</label>")
}

// Add adds content to the label element
func (l *Label) Add(e Elements) *Label {
	if e != nil {
		e.Prepare()
		l.contents = append(l.contents, e.Bytes())
	}
	return l
}

// Text adds text content to the label element
func (l *Label) Text(text string) *Label {
	l.contents = append(l.contents, []byte(text))
	return l
}

// For sets the for attribute
func (l *Label) For(forattr string) *Label {
	l.forattr = forattr
	return l
}

// Form sets the form attribute
func (l *Label) Form(form string) *Label {
	l.form = form
	return l
}

// AddStyle adds a single CSS property
func (l *Label) AddStyle(k, v string) *Label {
	l.style[k] = v
	return l
}

// AddStyles adds multiple CSS properties
func (l *Label) AddStyles(m map[string]string) *Label {
	for k, v := range m {
		l.style[k] = v
	}
	return l
}

// Style replaces all styles
func (l *Label) Style(m map[string]string) *Label {
	l.style = make(map[string]string)
	for k, v := range m {
		l.style[k] = v
	}
	return l
}

// Input represents the HTML input element for user input
type Input struct {
	buf          bytes.Buffer
	style        map[string]string
	inputtype    string
	name         string
	value        string
	placeholder  string
	id           string
	form         string
	required     bool
	disabled     bool
	readonly     bool
	autofocus    bool
	autocomplete string
	min          string
	max          string
	step         string
	pattern      string
	size         string
	maxlength    string
	minlength    string
	multiple     bool
	accept       string
	checked      bool
}

// NewInput creates a new Input element
func NewInput() *Input {
	return &Input{
		style: make(map[string]string),
	}
}

// Bytes returns the buffer contents
func (i *Input) Bytes() []byte {
	return i.buf.Bytes()
}

// Prepare builds the HTML for the input element
func (i *Input) Prepare() {
	i.buf.WriteString("<input")
	
	if i.inputtype != "" {
		i.buf.WriteString(" type=\"" + i.inputtype + "\"")
	}
	
	if i.name != "" {
		i.buf.WriteString(" name=\"" + i.name + "\"")
	}
	
	if i.value != "" {
		i.buf.WriteString(" value=\"" + i.value + "\"")
	}
	
	if i.placeholder != "" {
		i.buf.WriteString(" placeholder=\"" + i.placeholder + "\"")
	}
	
	if i.id != "" {
		i.buf.WriteString(" id=\"" + i.id + "\"")
	}
	
	if i.form != "" {
		i.buf.WriteString(" form=\"" + i.form + "\"")
	}
	
	if i.autocomplete != "" {
		i.buf.WriteString(" autocomplete=\"" + i.autocomplete + "\"")
	}
	
	if i.min != "" {
		i.buf.WriteString(" min=\"" + i.min + "\"")
	}
	
	if i.max != "" {
		i.buf.WriteString(" max=\"" + i.max + "\"")
	}
	
	if i.step != "" {
		i.buf.WriteString(" step=\"" + i.step + "\"")
	}
	
	if i.pattern != "" {
		i.buf.WriteString(" pattern=\"" + i.pattern + "\"")
	}
	
	if i.size != "" {
		i.buf.WriteString(" size=\"" + i.size + "\"")
	}
	
	if i.maxlength != "" {
		i.buf.WriteString(" maxlength=\"" + i.maxlength + "\"")
	}
	
	if i.minlength != "" {
		i.buf.WriteString(" minlength=\"" + i.minlength + "\"")
	}
	
	if i.accept != "" {
		i.buf.WriteString(" accept=\"" + i.accept + "\"")
	}
	
	if i.required {
		i.buf.WriteString(" required")
	}
	
	if i.disabled {
		i.buf.WriteString(" disabled")
	}
	
	if i.readonly {
		i.buf.WriteString(" readonly")
	}
	
	if i.autofocus {
		i.buf.WriteString(" autofocus")
	}
	
	if i.multiple {
		i.buf.WriteString(" multiple")
	}
	
	if i.checked {
		i.buf.WriteString(" checked")
	}
	
	if len(i.style) != 0 {
		idx := 0
		i.buf.WriteString(" style=\"")
		for k, v := range i.style {
			i.buf.WriteString(k + ": " + v + ";")
			if idx != len(i.style)-1 {
				i.buf.WriteByte(' ')
			}
			idx++
		}
		i.buf.WriteString("\"")
	}
	
	i.buf.WriteString(">")
}

// Type sets the type attribute
func (i *Input) Type(inputtype string) *Input {
	i.inputtype = inputtype
	return i
}

// Name sets the name attribute
func (i *Input) Name(name string) *Input {
	i.name = name
	return i
}

// Value sets the value attribute
func (i *Input) Value(value string) *Input {
	i.value = value
	return i
}

// Placeholder sets the placeholder attribute
func (i *Input) Placeholder(placeholder string) *Input {
	i.placeholder = placeholder
	return i
}

// Id sets the id attribute
func (i *Input) Id(id string) *Input {
	i.id = id
	return i
}

// Form sets the form attribute
func (i *Input) Form(form string) *Input {
	i.form = form
	return i
}

// Required sets the required attribute
func (i *Input) Required(required bool) *Input {
	i.required = required
	return i
}

// Disabled sets the disabled attribute
func (i *Input) Disabled(disabled bool) *Input {
	i.disabled = disabled
	return i
}

// Readonly sets the readonly attribute
func (i *Input) Readonly(readonly bool) *Input {
	i.readonly = readonly
	return i
}

// Autofocus sets the autofocus attribute
func (i *Input) Autofocus(autofocus bool) *Input {
	i.autofocus = autofocus
	return i
}

// Autocomplete sets the autocomplete attribute
func (i *Input) Autocomplete(autocomplete string) *Input {
	i.autocomplete = autocomplete
	return i
}

// Min sets the min attribute
func (i *Input) Min(min string) *Input {
	i.min = min
	return i
}

// Max sets the max attribute
func (i *Input) Max(max string) *Input {
	i.max = max
	return i
}

// Step sets the step attribute
func (i *Input) Step(step string) *Input {
	i.step = step
	return i
}

// Pattern sets the pattern attribute
func (i *Input) Pattern(pattern string) *Input {
	i.pattern = pattern
	return i
}

// Size sets the size attribute
func (i *Input) Size(size string) *Input {
	i.size = size
	return i
}

// Maxlength sets the maxlength attribute
func (i *Input) Maxlength(maxlength string) *Input {
	i.maxlength = maxlength
	return i
}

// Minlength sets the minlength attribute
func (i *Input) Minlength(minlength string) *Input {
	i.minlength = minlength
	return i
}

// Multiple sets the multiple attribute
func (i *Input) Multiple(multiple bool) *Input {
	i.multiple = multiple
	return i
}

// Accept sets the accept attribute
func (i *Input) Accept(accept string) *Input {
	i.accept = accept
	return i
}

// Checked sets the checked attribute
func (i *Input) Checked(checked bool) *Input {
	i.checked = checked
	return i
}

// AddStyle adds a single CSS property
func (i *Input) AddStyle(k, v string) *Input {
	i.style[k] = v
	return i
}

// AddStyles adds multiple CSS properties
func (i *Input) AddStyles(m map[string]string) *Input {
	for k, v := range m {
		i.style[k] = v
	}
	return i
}

// Style replaces all styles
func (i *Input) Style(m map[string]string) *Input {
	i.style = make(map[string]string)
	for k, v := range m {
		i.style[k] = v
	}
	return i
}

// Output represents the HTML output element for calculation results
type Output struct {
	buf      bytes.Buffer
	style    map[string]string
	contents [][]byte
	ttrack   int
	forattr  string
	name     string
	form     string
}

// NewOutput creates a new Output element
func NewOutput() *Output {
	return &Output{
		style: make(map[string]string),
	}
}

// Bytes returns the buffer contents
func (o *Output) Bytes() []byte {
	return o.buf.Bytes()
}

// Prepare builds the HTML for the output element
func (o *Output) Prepare() {
	o.buf.WriteString("<output")
	
	if o.forattr != "" {
		o.buf.WriteString(" for=\"" + o.forattr + "\"")
	}
	
	if o.name != "" {
		o.buf.WriteString(" name=\"" + o.name + "\"")
	}
	
	if o.form != "" {
		o.buf.WriteString(" form=\"" + o.form + "\"")
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
	
	o.buf.WriteString("</output>")
}

// Add adds content to the output element
func (o *Output) Add(e Elements) *Output {
	if e != nil {
		e.Prepare()
		o.contents = append(o.contents, e.Bytes())
	}
	return o
}

// Text adds text content to the output element
func (o *Output) Text(text string) *Output {
	o.contents = append(o.contents, []byte(text))
	return o
}

// For sets the for attribute
func (o *Output) For(forattr string) *Output {
	o.forattr = forattr
	return o
}

// Name sets the name attribute
func (o *Output) Name(name string) *Output {
	o.name = name
	return o
}

// Form sets the form attribute
func (o *Output) Form(form string) *Output {
	o.form = form
	return o
}

// AddStyle adds a single CSS property
func (o *Output) AddStyle(k, v string) *Output {
	o.style[k] = v
	return o
}

// AddStyles adds multiple CSS properties
func (o *Output) AddStyles(m map[string]string) *Output {
	for k, v := range m {
		o.style[k] = v
	}
	return o
}

// Style replaces all styles
func (o *Output) Style(m map[string]string) *Output {
	o.style = make(map[string]string)
	for k, v := range m {
		o.style[k] = v
	}
	return o
}

// Fieldset represents the HTML fieldset element for grouping form controls
type Fieldset struct {
	buf      bytes.Buffer
	style    map[string]string
	contents [][]byte
	ttrack   int
	form     string
	name     string
	disabled bool
}

// NewFieldset creates a new Fieldset element
func NewFieldset() *Fieldset {
	return &Fieldset{
		style: make(map[string]string),
	}
}

// Bytes returns the buffer contents
func (f *Fieldset) Bytes() []byte {
	return f.buf.Bytes()
}

// Prepare builds the HTML for the fieldset element
func (f *Fieldset) Prepare() {
	f.buf.WriteString("<fieldset")
	
	if f.form != "" {
		f.buf.WriteString(" form=\"" + f.form + "\"")
	}
	
	if f.name != "" {
		f.buf.WriteString(" name=\"" + f.name + "\"")
	}
	
	if f.disabled {
		f.buf.WriteString(" disabled")
	}
	
	if len(f.style) != 0 {
		idx := 0
		f.buf.WriteString(" style=\"")
		for k, v := range f.style {
			f.buf.WriteString(k + ": " + v + ";")
			if idx != len(f.style)-1 {
				f.buf.WriteByte(' ')
			}
			idx++
		}
		f.buf.WriteString("\"")
	}
	
	f.buf.WriteByte('>')
	
	for _, content := range f.contents {
		f.buf.Write(content)
	}
	
	f.buf.WriteString("</fieldset>")
}

// Add adds content to the fieldset element
func (f *Fieldset) Add(e Elements) *Fieldset {
	if e != nil {
		e.Prepare()
		f.contents = append(f.contents, e.Bytes())
	}
	return f
}

// Form sets the form attribute
func (f *Fieldset) Form(form string) *Fieldset {
	f.form = form
	return f
}

// Name sets the name attribute
func (f *Fieldset) Name(name string) *Fieldset {
	f.name = name
	return f
}

// Disabled sets the disabled attribute
func (f *Fieldset) Disabled(disabled bool) *Fieldset {
	f.disabled = disabled
	return f
}

// AddStyle adds a single CSS property
func (f *Fieldset) AddStyle(k, v string) *Fieldset {
	f.style[k] = v
	return f
}

// AddStyles adds multiple CSS properties
func (f *Fieldset) AddStyles(m map[string]string) *Fieldset {
	for k, v := range m {
		f.style[k] = v
	}
	return f
}

// Style replaces all styles
func (f *Fieldset) Style(m map[string]string) *Fieldset {
	f.style = make(map[string]string)
	for k, v := range m {
		f.style[k] = v
	}
	return f
}

// Button represents the HTML button element for clickable buttons
type Button struct {
	buf            bytes.Buffer
	style          map[string]string
	contents       [][]byte
	ttrack         int
	buttonType     string
	name           string
	value          string
	form           string
	formAction     string
	formEnctype    string
	formMethod     string
	formTarget     string
	formNovalidate bool
	disabled       bool
	autofocus      bool
}

// NewButton creates a new Button element
func NewButton() *Button {
	return &Button{
		style: make(map[string]string),
	}
}

// Bytes returns the buffer contents
func (b *Button) Bytes() []byte {
	return b.buf.Bytes()
}

// Prepare builds the HTML for the button element
func (b *Button) Prepare() {
	b.buf.WriteString("<button")
	
	if b.buttonType != "" {
		b.buf.WriteString(" type=\"" + b.buttonType + "\"")
	}
	
	if b.name != "" {
		b.buf.WriteString(" name=\"" + b.name + "\"")
	}
	
	if b.value != "" {
		b.buf.WriteString(" value=\"" + b.value + "\"")
	}
	
	if b.form != "" {
		b.buf.WriteString(" form=\"" + b.form + "\"")
	}
	
	if b.formAction != "" {
		b.buf.WriteString(" formaction=\"" + b.formAction + "\"")
	}
	
	if b.formEnctype != "" {
		b.buf.WriteString(" formenctype=\"" + b.formEnctype + "\"")
	}
	
	if b.formMethod != "" {
		b.buf.WriteString(" formmethod=\"" + b.formMethod + "\"")
	}
	
	if b.formTarget != "" {
		b.buf.WriteString(" formtarget=\"" + b.formTarget + "\"")
	}
	
	if b.formNovalidate {
		b.buf.WriteString(" formnovalidate")
	}
	
	if b.disabled {
		b.buf.WriteString(" disabled")
	}
	
	if b.autofocus {
		b.buf.WriteString(" autofocus")
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
	
	b.buf.WriteByte('>')
	
	for _, content := range b.contents {
		b.buf.Write(content)
	}
	
	b.buf.WriteString("</button>")
}

// Add adds content to the button element
func (b *Button) Add(e Elements) *Button {
	if e != nil {
		e.Prepare()
		b.contents = append(b.contents, e.Bytes())
	}
	return b
}

// Text adds text content to the button element
func (b *Button) Text(text string) *Button {
	b.contents = append(b.contents, []byte(text))
	return b
}

// Type sets the type attribute
func (b *Button) Type(buttonType string) *Button {
	b.buttonType = buttonType
	return b
}

// Name sets the name attribute
func (b *Button) Name(name string) *Button {
	b.name = name
	return b
}

// Value sets the value attribute
func (b *Button) Value(value string) *Button {
	b.value = value
	return b
}

// Form sets the form attribute
func (b *Button) Form(form string) *Button {
	b.form = form
	return b
}

// FormAction sets the formaction attribute
func (b *Button) FormAction(formAction string) *Button {
	b.formAction = formAction
	return b
}

// FormEnctype sets the formenctype attribute
func (b *Button) FormEnctype(formEnctype string) *Button {
	b.formEnctype = formEnctype
	return b
}

// FormMethod sets the formmethod attribute
func (b *Button) FormMethod(formMethod string) *Button {
	b.formMethod = formMethod
	return b
}

// FormTarget sets the formtarget attribute
func (b *Button) FormTarget(formTarget string) *Button {
	b.formTarget = formTarget
	return b
}

// FormNovalidate sets the formnovalidate attribute
func (b *Button) FormNovalidate(formNovalidate bool) *Button {
	b.formNovalidate = formNovalidate
	return b
}

// Disabled sets the disabled attribute
func (b *Button) Disabled(disabled bool) *Button {
	b.disabled = disabled
	return b
}

// Autofocus sets the autofocus attribute
func (b *Button) Autofocus(autofocus bool) *Button {
	b.autofocus = autofocus
	return b
}

// AddStyle adds a single CSS property
func (b *Button) AddStyle(k, v string) *Button {
	b.style[k] = v
	return b
}

// AddStyles adds multiple CSS properties
func (b *Button) AddStyles(m map[string]string) *Button {
	for k, v := range m {
		b.style[k] = v
	}
	return b
}

// Style replaces all styles
func (b *Button) Style(m map[string]string) *Button {
	b.style = make(map[string]string)
	for k, v := range m {
		b.style[k] = v
	}
	return b
}

// Select represents the HTML select element for dropdown lists
type Select struct {
	buf        bytes.Buffer
	style      map[string]string
	contents   [][]byte
	ttrack     int
	name       string
	form       string
	size       string
	multiple   bool
	required   bool
	disabled   bool
	autofocus  bool
	autoComplete string
}

// NewSelect creates a new Select element
func NewSelect() *Select {
	return &Select{
		style: make(map[string]string),
	}
}

// Bytes returns the buffer contents
func (s *Select) Bytes() []byte {
	return s.buf.Bytes()
}

// Prepare builds the HTML for the select element
func (s *Select) Prepare() {
	s.buf.WriteString("<select")
	
	if s.name != "" {
		s.buf.WriteString(" name=\"" + s.name + "\"")
	}
	
	if s.form != "" {
		s.buf.WriteString(" form=\"" + s.form + "\"")
	}
	
	if s.size != "" {
		s.buf.WriteString(" size=\"" + s.size + "\"")
	}
	
	if s.autoComplete != "" {
		s.buf.WriteString(" autocomplete=\"" + s.autoComplete + "\"")
	}
	
	if s.multiple {
		s.buf.WriteString(" multiple")
	}
	
	if s.required {
		s.buf.WriteString(" required")
	}
	
	if s.disabled {
		s.buf.WriteString(" disabled")
	}
	
	if s.autofocus {
		s.buf.WriteString(" autofocus")
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
	
	s.buf.WriteString("</select>")
}

// Add adds content to the select element
func (s *Select) Add(e Elements) *Select {
	if e != nil {
		e.Prepare()
		s.contents = append(s.contents, e.Bytes())
	}
	return s
}

// Name sets the name attribute
func (s *Select) Name(name string) *Select {
	s.name = name
	return s
}

// Form sets the form attribute
func (s *Select) Form(form string) *Select {
	s.form = form
	return s
}

// Size sets the size attribute
func (s *Select) Size(size string) *Select {
	s.size = size
	return s
}

// Multiple sets the multiple attribute
func (s *Select) Multiple(multiple bool) *Select {
	s.multiple = multiple
	return s
}

// Required sets the required attribute
func (s *Select) Required(required bool) *Select {
	s.required = required
	return s
}

// Disabled sets the disabled attribute
func (s *Select) Disabled(disabled bool) *Select {
	s.disabled = disabled
	return s
}

// Autofocus sets the autofocus attribute
func (s *Select) Autofocus(autofocus bool) *Select {
	s.autofocus = autofocus
	return s
}

// AutoComplete sets the autocomplete attribute
func (s *Select) AutoComplete(autoComplete string) *Select {
	s.autoComplete = autoComplete
	return s
}

// AddStyle adds a single CSS property
func (s *Select) AddStyle(k, v string) *Select {
	s.style[k] = v
	return s
}

// AddStyles adds multiple CSS properties
func (s *Select) AddStyles(m map[string]string) *Select {
	for k, v := range m {
		s.style[k] = v
	}
	return s
}

// Style replaces all styles
func (s *Select) Style(m map[string]string) *Select {
	s.style = make(map[string]string)
	for k, v := range m {
		s.style[k] = v
	}
	return s
}
