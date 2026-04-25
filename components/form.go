package rephtml

import "bytes"

// Form represents the HTML form element for user input
type Form struct {
	buf           bytes.Buffer
	style         map[string]string
	contents      []Element
	ttrack        int
	action        string
	method        string
	enctype       string
	name          string
	target        string
	autocomplete  string
	novalidate    bool
	acceptcharset string
}

// NewForm creates a new Form element
func NewForm() *Form {
	return &Form{
		style: make(map[string]string),
	}
}

// Bytes returns the buffer contents
func (f *Form) Bytes() []byte {
	return cloneBytes(f.buf.Bytes())
}

// IsBodyElement implements BodyElement interface
func (f *Form) IsBodyElement() {}

// Prepare builds the HTML for the form element
func (f *Form) Prepare() {
	f.buf.Reset()
	f.buf.WriteString("<form")

	if f.action != "" {
		writeAttr(&f.buf, "action", f.action)
	}

	if f.method != "" {
		writeAttr(&f.buf, "method", f.method)
	}

	if f.enctype != "" {
		writeAttr(&f.buf, "enctype", f.enctype)
	}

	if f.name != "" {
		writeAttr(&f.buf, "name", f.name)
	}

	if f.target != "" {
		writeAttr(&f.buf, "target", f.target)
	}

	if f.autocomplete != "" {
		writeAttr(&f.buf, "autocomplete", f.autocomplete)
	}

	if f.acceptcharset != "" {
		writeAttr(&f.buf, "accept-charset", f.acceptcharset)
	}

	if f.novalidate {
		f.buf.WriteString(" novalidate")
	}

	if len(f.style) != 0 {
		parseStyle(&f.buf, f.style)
	}

	f.buf.WriteByte('>')

	writeElements(&f.buf, f.contents)

	f.buf.WriteString("</form>")
}

// Add adds content to the form element
func (f *Form) Add(e Element) *Form {
	if e != nil {
		f.contents = appendElement(f.contents, e)
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
	contents []Element
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
	return cloneBytes(l.buf.Bytes())
}

// IsBodyElement implements BodyElement interface
func (l *Label) IsBodyElement() {}

// Prepare builds the HTML for the label element
func (l *Label) Prepare() {
	l.buf.Reset()
	l.buf.WriteString("<label")

	if l.forattr != "" {
		writeAttr(&l.buf, "for", l.forattr)
	}

	if l.form != "" {
		writeAttr(&l.buf, "form", l.form)
	}

	if len(l.style) != 0 {
		parseStyle(&l.buf, l.style)
	}

	l.buf.WriteByte('>')

	writeElements(&l.buf, l.contents)

	l.buf.WriteString("</label>")
}

// Add adds content to the label element
func (l *Label) Add(e Element) *Label {
	if e != nil {
		l.contents = appendElement(l.contents, e)
	}
	return l
}

// Text adds text content to the label element
func (l *Label) Text(text string) *Label {
	l.contents = appendElement(l.contents, escapedText(text))
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
	return cloneBytes(i.buf.Bytes())
}

// IsBodyElement implements BodyElement interface
func (i *Input) IsBodyElement() {}

// Prepare builds the HTML for the input element
func (i *Input) Prepare() {
	i.buf.Reset()
	i.buf.WriteString("<input")

	if i.inputtype != "" {
		writeAttr(&i.buf, "type", i.inputtype)
	}

	if i.name != "" {
		writeAttr(&i.buf, "name", i.name)
	}

	if i.value != "" {
		writeAttr(&i.buf, "value", i.value)
	}

	if i.placeholder != "" {
		writeAttr(&i.buf, "placeholder", i.placeholder)
	}

	if i.id != "" {
		writeAttr(&i.buf, "id", i.id)
	}

	if i.form != "" {
		writeAttr(&i.buf, "form", i.form)
	}

	if i.autocomplete != "" {
		writeAttr(&i.buf, "autocomplete", i.autocomplete)
	}

	if i.min != "" {
		writeAttr(&i.buf, "min", i.min)
	}

	if i.max != "" {
		writeAttr(&i.buf, "max", i.max)
	}

	if i.step != "" {
		writeAttr(&i.buf, "step", i.step)
	}

	if i.pattern != "" {
		writeAttr(&i.buf, "pattern", i.pattern)
	}

	if i.size != "" {
		writeAttr(&i.buf, "size", i.size)
	}

	if i.maxlength != "" {
		writeAttr(&i.buf, "maxlength", i.maxlength)
	}

	if i.minlength != "" {
		writeAttr(&i.buf, "minlength", i.minlength)
	}

	if i.accept != "" {
		writeAttr(&i.buf, "accept", i.accept)
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
		parseStyle(&i.buf, i.style)
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
	contents []Element
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
	return cloneBytes(o.buf.Bytes())
}

// IsBodyElement implements BodyElement interface
func (o *Output) IsBodyElement() {}

// Prepare builds the HTML for the output element
func (o *Output) Prepare() {
	o.buf.Reset()
	o.buf.WriteString("<output")

	if o.forattr != "" {
		writeAttr(&o.buf, "for", o.forattr)
	}

	if o.name != "" {
		writeAttr(&o.buf, "name", o.name)
	}

	if o.form != "" {
		writeAttr(&o.buf, "form", o.form)
	}

	if len(o.style) != 0 {
		parseStyle(&o.buf, o.style)
	}

	o.buf.WriteByte('>')

	writeElements(&o.buf, o.contents)

	o.buf.WriteString("</output>")
}

// Add adds content to the output element
func (o *Output) Add(e Element) *Output {
	if e != nil {
		o.contents = appendElement(o.contents, e)
	}
	return o
}

// Text adds text content to the output element
func (o *Output) Text(text string) *Output {
	o.contents = appendElement(o.contents, escapedText(text))
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
	contents []Element
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
	return cloneBytes(f.buf.Bytes())
}

// IsBodyElement implements BodyElement interface
func (f *Fieldset) IsBodyElement() {}

// Prepare builds the HTML for the fieldset element
func (f *Fieldset) Prepare() {
	f.buf.Reset()
	f.buf.WriteString("<fieldset")

	if f.form != "" {
		writeAttr(&f.buf, "form", f.form)
	}

	if f.name != "" {
		writeAttr(&f.buf, "name", f.name)
	}

	if f.disabled {
		f.buf.WriteString(" disabled")
	}

	if len(f.style) != 0 {
		parseStyle(&f.buf, f.style)
	}

	f.buf.WriteByte('>')

	writeElements(&f.buf, f.contents)

	f.buf.WriteString("</fieldset>")
}

// Add adds content to the fieldset element
func (f *Fieldset) Add(e Element) *Fieldset {
	if e != nil {
		f.contents = appendElement(f.contents, e)
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
	contents       []Element
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
	return cloneBytes(b.buf.Bytes())
}

// IsBodyElement implements BodyElement interface
func (b *Button) IsBodyElement() {}

// Prepare builds the HTML for the button element
func (b *Button) Prepare() {
	b.buf.Reset()
	b.buf.WriteString("<button")

	if b.buttonType != "" {
		writeAttr(&b.buf, "type", b.buttonType)
	}

	if b.name != "" {
		writeAttr(&b.buf, "name", b.name)
	}

	if b.value != "" {
		writeAttr(&b.buf, "value", b.value)
	}

	if b.form != "" {
		writeAttr(&b.buf, "form", b.form)
	}

	if b.formAction != "" {
		writeAttr(&b.buf, "formaction", b.formAction)
	}

	if b.formEnctype != "" {
		writeAttr(&b.buf, "formenctype", b.formEnctype)
	}

	if b.formMethod != "" {
		writeAttr(&b.buf, "formmethod", b.formMethod)
	}

	if b.formTarget != "" {
		writeAttr(&b.buf, "formtarget", b.formTarget)
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
		parseStyle(&b.buf, b.style)
	}

	b.buf.WriteByte('>')

	writeElements(&b.buf, b.contents)

	b.buf.WriteString("</button>")
}

// Add adds content to the button element
func (b *Button) Add(e Element) *Button {
	if e != nil {
		b.contents = appendElement(b.contents, e)
	}
	return b
}

// Text adds text content to the button element
func (b *Button) Text(text string) *Button {
	b.contents = appendElement(b.contents, escapedText(text))
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
	buf          bytes.Buffer
	style        map[string]string
	contents     []Element
	ttrack       int
	name         string
	form         string
	size         string
	multiple     bool
	required     bool
	disabled     bool
	autofocus    bool
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
	return cloneBytes(s.buf.Bytes())
}

// IsBodyElement implements BodyElement interface
func (s *Select) IsBodyElement() {}

// Prepare builds the HTML for the select element
func (s *Select) Prepare() {
	s.buf.Reset()
	s.buf.WriteString("<select")

	if s.name != "" {
		writeAttr(&s.buf, "name", s.name)
	}

	if s.form != "" {
		writeAttr(&s.buf, "form", s.form)
	}

	if s.size != "" {
		writeAttr(&s.buf, "size", s.size)
	}

	if s.autoComplete != "" {
		writeAttr(&s.buf, "autocomplete", s.autoComplete)
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
		parseStyle(&s.buf, s.style)
	}

	s.buf.WriteByte('>')

	writeElements(&s.buf, s.contents)

	s.buf.WriteString("</select>")
}

// Add adds content to the select element
func (s *Select) Add(e Element) *Select {
	if e != nil {
		s.contents = appendElement(s.contents, e)
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

// Datalist represents the HTML datalist element for predefined options
type Datalist struct {
	buf      bytes.Buffer
	style    map[string]string
	contents []Element
	ttrack   int
	id       string
}

// NewDatalist creates a new Datalist element
func NewDatalist() *Datalist {
	return &Datalist{
		style: make(map[string]string),
	}
}

// Bytes returns the buffer contents
func (d *Datalist) Bytes() []byte {
	return cloneBytes(d.buf.Bytes())
}

// IsBodyElement implements BodyElement interface
func (d *Datalist) IsBodyElement() {}

// Prepare builds the HTML for the datalist element
func (d *Datalist) Prepare() {
	d.buf.Reset()
	d.buf.WriteString("<datalist")

	if d.id != "" {
		writeAttr(&d.buf, "id", d.id)
	}

	if len(d.style) != 0 {
		parseStyle(&d.buf, d.style)
	}

	d.buf.WriteByte('>')

	writeElements(&d.buf, d.contents)

	d.buf.WriteString("</datalist>")
}

// Add adds content to the datalist element
func (d *Datalist) Add(e Element) *Datalist {
	if e != nil {
		d.contents = appendElement(d.contents, e)
	}
	return d
}

// Id sets the id attribute
func (d *Datalist) Id(id string) *Datalist {
	d.id = id
	return d
}

// AddStyle adds a single CSS property
func (d *Datalist) AddStyle(k, v string) *Datalist {
	d.style[k] = v
	return d
}

// AddStyles adds multiple CSS properties
func (d *Datalist) AddStyles(m map[string]string) *Datalist {
	for k, v := range m {
		d.style[k] = v
	}
	return d
}

// Style replaces all styles
func (d *Datalist) Style(m map[string]string) *Datalist {
	d.style = make(map[string]string)
	for k, v := range m {
		d.style[k] = v
	}
	return d
}

// Optgroup represents the HTML optgroup element for grouping options
type Optgroup struct {
	buf      bytes.Buffer
	style    map[string]string
	contents []Element
	ttrack   int
	label    string
	disabled bool
}

// NewOptgroup creates a new Optgroup element
func NewOptgroup() *Optgroup {
	return &Optgroup{
		style: make(map[string]string),
	}
}

// Bytes returns the buffer contents
func (o *Optgroup) Bytes() []byte {
	return cloneBytes(o.buf.Bytes())
}

// IsBodyElement implements BodyElement interface
func (o *Optgroup) IsBodyElement() {}

// Prepare builds the HTML for the optgroup element
func (o *Optgroup) Prepare() {
	o.buf.Reset()
	o.buf.WriteString("<optgroup")

	if o.label != "" {
		writeAttr(&o.buf, "label", o.label)
	}

	if o.disabled {
		o.buf.WriteString(" disabled")
	}

	if len(o.style) != 0 {
		parseStyle(&o.buf, o.style)
	}

	o.buf.WriteByte('>')

	writeElements(&o.buf, o.contents)

	o.buf.WriteString("</optgroup>")
}

// Add adds content to the optgroup element
func (o *Optgroup) Add(e Element) *Optgroup {
	if e != nil {
		o.contents = appendElement(o.contents, e)
	}
	return o
}

// Label sets the label attribute
func (o *Optgroup) Label(label string) *Optgroup {
	o.label = label
	return o
}

// Disabled sets the disabled attribute
func (o *Optgroup) Disabled(disabled bool) *Optgroup {
	o.disabled = disabled
	return o
}

// AddStyle adds a single CSS property
func (o *Optgroup) AddStyle(k, v string) *Optgroup {
	o.style[k] = v
	return o
}

// AddStyles adds multiple CSS properties
func (o *Optgroup) AddStyles(m map[string]string) *Optgroup {
	for k, v := range m {
		o.style[k] = v
	}
	return o
}

// Style replaces all styles
func (o *Optgroup) Style(m map[string]string) *Optgroup {
	o.style = make(map[string]string)
	for k, v := range m {
		o.style[k] = v
	}
	return o
}

// Option represents the HTML option element for select options
type Option struct {
	buf      bytes.Buffer
	style    map[string]string
	contents []Element
	ttrack   int
	value    string
	label    string
	selected bool
	disabled bool
}

// NewOption creates a new Option element
func NewOption() *Option {
	return &Option{
		style: make(map[string]string),
	}
}

// Bytes returns the buffer contents
func (o *Option) Bytes() []byte {
	return cloneBytes(o.buf.Bytes())
}

// IsBodyElement implements BodyElement interface
func (o *Option) IsBodyElement() {}

// Prepare builds the HTML for the option element
func (o *Option) Prepare() {
	o.buf.Reset()
	o.buf.WriteString("<option")

	if o.value != "" {
		writeAttr(&o.buf, "value", o.value)
	}

	if o.label != "" {
		writeAttr(&o.buf, "label", o.label)
	}

	if o.selected {
		o.buf.WriteString(" selected")
	}

	if o.disabled {
		o.buf.WriteString(" disabled")
	}

	if len(o.style) != 0 {
		parseStyle(&o.buf, o.style)
	}

	o.buf.WriteByte('>')

	writeElements(&o.buf, o.contents)

	o.buf.WriteString("</option>")
}

// Add adds content to the option element
func (o *Option) Add(e Element) *Option {
	if e != nil {
		o.contents = appendElement(o.contents, e)
	}
	return o
}

// Text adds text content to the option element
func (o *Option) Text(text string) *Option {
	o.contents = appendElement(o.contents, escapedText(text))
	return o
}

// Value sets the value attribute
func (o *Option) Value(value string) *Option {
	o.value = value
	return o
}

// Label sets the label attribute
func (o *Option) Label(label string) *Option {
	o.label = label
	return o
}

// Selected sets the selected attribute
func (o *Option) Selected(selected bool) *Option {
	o.selected = selected
	return o
}

// Disabled sets the disabled attribute
func (o *Option) Disabled(disabled bool) *Option {
	o.disabled = disabled
	return o
}

// AddStyle adds a single CSS property
func (o *Option) AddStyle(k, v string) *Option {
	o.style[k] = v
	return o
}

// AddStyles adds multiple CSS properties
func (o *Option) AddStyles(m map[string]string) *Option {
	for k, v := range m {
		o.style[k] = v
	}
	return o
}

// Style replaces all styles
func (o *Option) Style(m map[string]string) *Option {
	o.style = make(map[string]string)
	for k, v := range m {
		o.style[k] = v
	}
	return o
}

// Textarea represents the HTML textarea element for multi-line text input
type Textarea struct {
	buf          bytes.Buffer
	style        map[string]string
	contents     []Element
	ttrack       int
	name         string
	form         string
	rows         string
	cols         string
	placeholder  string
	maxLength    string
	minLength    string
	wrap         string
	required     bool
	disabled     bool
	readonly     bool
	autofocus    bool
	autoComplete string
	spellcheck   string
}

// NewTextarea creates a new Textarea element
func NewTextarea() *Textarea {
	return &Textarea{
		style: make(map[string]string),
	}
}

// Bytes returns the buffer contents
func (t *Textarea) Bytes() []byte {
	return cloneBytes(t.buf.Bytes())
}

// IsBodyElement implements BodyElement interface
func (t *Textarea) IsBodyElement() {}

// Prepare builds the HTML for the textarea element
func (t *Textarea) Prepare() {
	t.buf.Reset()
	t.buf.WriteString("<textarea")

	if t.name != "" {
		writeAttr(&t.buf, "name", t.name)
	}

	if t.form != "" {
		writeAttr(&t.buf, "form", t.form)
	}

	if t.rows != "" {
		writeAttr(&t.buf, "rows", t.rows)
	}

	if t.cols != "" {
		writeAttr(&t.buf, "cols", t.cols)
	}

	if t.placeholder != "" {
		writeAttr(&t.buf, "placeholder", t.placeholder)
	}

	if t.maxLength != "" {
		writeAttr(&t.buf, "maxlength", t.maxLength)
	}

	if t.minLength != "" {
		writeAttr(&t.buf, "minlength", t.minLength)
	}

	if t.wrap != "" {
		writeAttr(&t.buf, "wrap", t.wrap)
	}

	if t.autoComplete != "" {
		writeAttr(&t.buf, "autocomplete", t.autoComplete)
	}

	if t.spellcheck != "" {
		writeAttr(&t.buf, "spellcheck", t.spellcheck)
	}

	if t.required {
		t.buf.WriteString(" required")
	}

	if t.disabled {
		t.buf.WriteString(" disabled")
	}

	if t.readonly {
		t.buf.WriteString(" readonly")
	}

	if t.autofocus {
		t.buf.WriteString(" autofocus")
	}

	if len(t.style) != 0 {
		parseStyle(&t.buf, t.style)
	}

	t.buf.WriteByte('>')

	writeElements(&t.buf, t.contents)

	t.buf.WriteString("</textarea>")
}

// Add adds content to the textarea element
func (t *Textarea) Add(e Element) *Textarea {
	if e != nil {
		t.contents = appendElement(t.contents, e)
	}
	return t
}

// Text adds text content to the textarea element
func (t *Textarea) Text(text string) *Textarea {
	t.contents = appendElement(t.contents, escapedText(text))
	return t
}

// Name sets the name attribute
func (t *Textarea) Name(name string) *Textarea {
	t.name = name
	return t
}

// Form sets the form attribute
func (t *Textarea) Form(form string) *Textarea {
	t.form = form
	return t
}

// Rows sets the rows attribute
func (t *Textarea) Rows(rows string) *Textarea {
	t.rows = rows
	return t
}

// Cols sets the cols attribute
func (t *Textarea) Cols(cols string) *Textarea {
	t.cols = cols
	return t
}

// Placeholder sets the placeholder attribute
func (t *Textarea) Placeholder(placeholder string) *Textarea {
	t.placeholder = placeholder
	return t
}

// MaxLength sets the maxlength attribute
func (t *Textarea) MaxLength(maxLength string) *Textarea {
	t.maxLength = maxLength
	return t
}

// MinLength sets the minlength attribute
func (t *Textarea) MinLength(minLength string) *Textarea {
	t.minLength = minLength
	return t
}

// Wrap sets the wrap attribute
func (t *Textarea) Wrap(wrap string) *Textarea {
	t.wrap = wrap
	return t
}

// Required sets the required attribute
func (t *Textarea) Required(required bool) *Textarea {
	t.required = required
	return t
}

// Disabled sets the disabled attribute
func (t *Textarea) Disabled(disabled bool) *Textarea {
	t.disabled = disabled
	return t
}

// Readonly sets the readonly attribute
func (t *Textarea) Readonly(readonly bool) *Textarea {
	t.readonly = readonly
	return t
}

// Autofocus sets the autofocus attribute
func (t *Textarea) Autofocus(autofocus bool) *Textarea {
	t.autofocus = autofocus
	return t
}

// AutoComplete sets the autocomplete attribute
func (t *Textarea) AutoComplete(autoComplete string) *Textarea {
	t.autoComplete = autoComplete
	return t
}

// Spellcheck sets the spellcheck attribute
func (t *Textarea) Spellcheck(spellcheck string) *Textarea {
	t.spellcheck = spellcheck
	return t
}

// AddStyle adds a single CSS property
func (t *Textarea) AddStyle(k, v string) *Textarea {
	t.style[k] = v
	return t
}

// AddStyles adds multiple CSS properties
func (t *Textarea) AddStyles(m map[string]string) *Textarea {
	for k, v := range m {
		t.style[k] = v
	}
	return t
}

// Style replaces all styles
func (t *Textarea) Style(m map[string]string) *Textarea {
	t.style = make(map[string]string)
	for k, v := range m {
		t.style[k] = v
	}
	return t
}

// Progress represents the HTML progress element for showing completion progress
type Progress struct {
	buf      bytes.Buffer
	style    map[string]string
	contents []Element
	ttrack   int
	value    string
	max      string
	form     string
}

// NewProgress creates a new Progress element
func NewProgress() *Progress {
	return &Progress{
		style: make(map[string]string),
	}
}

// Bytes returns the buffer contents
func (p *Progress) Bytes() []byte {
	return cloneBytes(p.buf.Bytes())
}

// IsBodyElement implements BodyElement interface
func (p *Progress) IsBodyElement() {}

// Prepare builds the HTML for the progress element
func (p *Progress) Prepare() {
	p.buf.Reset()
	p.buf.WriteString("<progress")

	if p.value != "" {
		writeAttr(&p.buf, "value", p.value)
	}

	if p.max != "" {
		writeAttr(&p.buf, "max", p.max)
	}

	if p.form != "" {
		writeAttr(&p.buf, "form", p.form)
	}

	if len(p.style) != 0 {
		parseStyle(&p.buf, p.style)
	}

	p.buf.WriteByte('>')

	writeElements(&p.buf, p.contents)

	p.buf.WriteString("</progress>")
}

// Add adds content to the progress element
func (p *Progress) Add(e Element) *Progress {
	if e != nil {
		p.contents = appendElement(p.contents, e)
	}
	return p
}

// Text adds text content to the progress element (fallback for non-supporting browsers)
func (p *Progress) Text(text string) *Progress {
	p.contents = appendElement(p.contents, escapedText(text))
	return p
}

// Value sets the value attribute
func (p *Progress) Value(value string) *Progress {
	p.value = value
	return p
}

// Max sets the max attribute
func (p *Progress) Max(max string) *Progress {
	p.max = max
	return p
}

// Form sets the form attribute
func (p *Progress) Form(form string) *Progress {
	p.form = form
	return p
}

// AddStyle adds a single CSS property
func (p *Progress) AddStyle(k, v string) *Progress {
	p.style[k] = v
	return p
}

// AddStyles adds multiple CSS properties
func (p *Progress) AddStyles(m map[string]string) *Progress {
	for k, v := range m {
		p.style[k] = v
	}
	return p
}

// Style replaces all styles
func (p *Progress) Style(m map[string]string) *Progress {
	p.style = make(map[string]string)
	for k, v := range m {
		p.style[k] = v
	}
	return p
}

// Meter represents the HTML meter element for displaying scalar measurements
type Meter struct {
	buf      bytes.Buffer
	style    map[string]string
	contents []Element
	ttrack   int
	value    string
	min      string
	max      string
	low      string
	high     string
	optimum  string
	form     string
}

// NewMeter creates a new Meter element
func NewMeter() *Meter {
	return &Meter{
		style: make(map[string]string),
	}
}

// Bytes returns the buffer contents
func (m *Meter) Bytes() []byte {
	return cloneBytes(m.buf.Bytes())
}

// IsBodyElement implements BodyElement interface
func (m *Meter) IsBodyElement() {}

// Prepare builds the HTML for the meter element
func (m *Meter) Prepare() {
	m.buf.Reset()
	m.buf.WriteString("<meter")

	if m.value != "" {
		writeAttr(&m.buf, "value", m.value)
	}

	if m.min != "" {
		writeAttr(&m.buf, "min", m.min)
	}

	if m.max != "" {
		writeAttr(&m.buf, "max", m.max)
	}

	if m.low != "" {
		writeAttr(&m.buf, "low", m.low)
	}

	if m.high != "" {
		writeAttr(&m.buf, "high", m.high)
	}

	if m.optimum != "" {
		writeAttr(&m.buf, "optimum", m.optimum)
	}

	if m.form != "" {
		writeAttr(&m.buf, "form", m.form)
	}

	if len(m.style) != 0 {
		parseStyle(&m.buf, m.style)
	}

	m.buf.WriteByte('>')

	writeElements(&m.buf, m.contents)

	m.buf.WriteString("</meter>")
}

// Add adds content to the meter element
func (m *Meter) Add(e Element) *Meter {
	if e != nil {
		m.contents = appendElement(m.contents, e)
	}
	return m
}

// Text adds text content to the meter element (fallback for non-supporting browsers)
func (m *Meter) Text(text string) *Meter {
	m.contents = appendElement(m.contents, escapedText(text))
	return m
}

// Value sets the value attribute
func (m *Meter) Value(value string) *Meter {
	m.value = value
	return m
}

// Min sets the min attribute
func (m *Meter) Min(min string) *Meter {
	m.min = min
	return m
}

// Max sets the max attribute
func (m *Meter) Max(max string) *Meter {
	m.max = max
	return m
}

// Low sets the low attribute
func (m *Meter) Low(low string) *Meter {
	m.low = low
	return m
}

// High sets the high attribute
func (m *Meter) High(high string) *Meter {
	m.high = high
	return m
}

// Optimum sets the optimum attribute
func (m *Meter) Optimum(optimum string) *Meter {
	m.optimum = optimum
	return m
}

// Form sets the form attribute
func (m *Meter) Form(form string) *Meter {
	m.form = form
	return m
}

// AddStyle adds a single CSS property
func (m *Meter) AddStyle(k, v string) *Meter {
	m.style[k] = v
	return m
}

// AddStyles adds multiple CSS properties
func (m *Meter) AddStyles(ma map[string]string) *Meter {
	for k, v := range ma {
		m.style[k] = v
	}
	return m
}

// Style replaces all styles
func (m *Meter) Style(ma map[string]string) *Meter {
	m.style = make(map[string]string)
	for k, v := range ma {
		m.style[k] = v
	}
	return m
}

// Legend represents the HTML legend element for fieldset captions
type Legend struct {
	buf      bytes.Buffer
	style    map[string]string
	contents []Element
	ttrack   int
}

// NewLegend creates a new Legend element
func NewLegend() *Legend {
	return &Legend{
		style: make(map[string]string),
	}
}

// Bytes returns the buffer contents
func (l *Legend) Bytes() []byte {
	return cloneBytes(l.buf.Bytes())
}

// IsBodyElement implements BodyElement interface
func (l *Legend) IsBodyElement() {}

// Prepare builds the HTML for the legend element
func (l *Legend) Prepare() {
	l.buf.Reset()
	l.buf.WriteString("<legend")

	if len(l.style) != 0 {
		parseStyle(&l.buf, l.style)
	}

	l.buf.WriteByte('>')

	writeElements(&l.buf, l.contents)

	l.buf.WriteString("</legend>")
}

// Add adds content to the legend element
func (l *Legend) Add(e Element) *Legend {
	if e != nil {
		l.contents = appendElement(l.contents, e)
	}
	return l
}

// Text adds text content to the legend element
func (l *Legend) Text(text string) *Legend {
	l.contents = appendElement(l.contents, escapedText(text))
	return l
}

// AddStyle adds a single CSS property
func (l *Legend) AddStyle(k, v string) *Legend {
	l.style[k] = v
	return l
}

// AddStyles adds multiple CSS properties
func (l *Legend) AddStyles(m map[string]string) *Legend {
	for k, v := range m {
		l.style[k] = v
	}
	return l
}

// Style replaces all styles
func (l *Legend) Style(m map[string]string) *Legend {
	l.style = make(map[string]string)
	for k, v := range m {
		l.style[k] = v
	}
	return l
}
