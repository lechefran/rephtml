package rephtml

import "bytes"

// Form represents the HTML <form> element, used for user input.
type Form struct {
	bodyElement
	contentNode[*Form]
	action        string
	method        string
	enctype       string
	name          string
	target        string
	autocomplete  string
	novalidate    bool
	acceptcharset string
}

// NewForm returns a new, empty <form> element.
func NewForm() *Form {
	v := &Form{}
	v.init(v)
	return v
}

// renderTo writes the form element's HTML to buf.
func (f *Form) renderTo(buf *bytes.Buffer) {
	tg := startTag(buf, "form")
	tg.urlAttr("action", f.action)
	tg.attr("method", f.method)
	tg.attr("enctype", f.enctype)
	tg.attr("name", f.name)
	tg.attr("target", f.target)
	tg.attr("autocomplete", f.autocomplete)
	tg.attr("accept-charset", f.acceptcharset)
	tg.boolAttr("novalidate", f.novalidate)
	f.attrs(tg)
	tg.children(f.contents)
}

// Action sets the action attribute and returns the element.
//
// An unsafe URL is replaced with [BlockedURL] when rendered; see [IsSafeURL].
func (f *Form) Action(action string) *Form {
	f.action = action
	return f
}

// Method sets the method attribute and returns the element.
func (f *Form) Method(method string) *Form {
	f.method = method
	return f
}

// Enctype sets the enctype attribute and returns the element.
func (f *Form) Enctype(enctype string) *Form {
	f.enctype = enctype
	return f
}

// Name sets the name attribute and returns the element.
func (f *Form) Name(name string) *Form {
	f.name = name
	return f
}

// Target sets the target attribute and returns the element.
func (f *Form) Target(target string) *Form {
	f.target = target
	return f
}

// AutoComplete sets the autocomplete attribute and returns the element.
func (f *Form) AutoComplete(autocomplete string) *Form {
	f.autocomplete = autocomplete
	return f
}

// AcceptCharset sets the accept-charset attribute and returns the element.
func (f *Form) AcceptCharset(acceptcharset string) *Form {
	f.acceptcharset = acceptcharset
	return f
}

// NoValidate sets the novalidate attribute and returns the element.
func (f *Form) NoValidate(novalidate bool) *Form {
	f.novalidate = novalidate
	return f
}

// Label represents the HTML <label> element, used for form controls.
type Label struct {
	bodyElement
	textContent[*Label]
	forattr string
	form    string
}

// NewLabel returns a new, empty <label> element.
func NewLabel() *Label {
	v := &Label{}
	v.init(v)
	return v
}

// renderTo writes the label element's HTML to buf.
func (l *Label) renderTo(buf *bytes.Buffer) {
	tg := startTag(buf, "label")
	tg.attr("for", l.forattr)
	tg.attr("form", l.form)
	l.attrs(tg)
	tg.children(l.contents)
}

// For sets the for attribute and returns the element.
func (l *Label) For(forattr string) *Label {
	l.forattr = forattr
	return l
}

// Form sets the form attribute and returns the element.
func (l *Label) Form(form string) *Label {
	l.form = form
	return l
}

// Input represents the HTML <input> element, used for user input.
type Input struct {
	bodyElement
	node[*Input]
	inputtype    string
	name         string
	value        string
	placeholder  string
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
	size         int
	maxlength    optionalInt
	minlength    optionalInt
	multiple     bool
	accept       string
	checked      bool
}

// NewInput returns a new, empty <input> element.
func NewInput() *Input {
	v := &Input{}
	v.init(v)
	return v
}

// renderTo writes the input element's HTML to buf.
func (i *Input) renderTo(buf *bytes.Buffer) {
	tg := startTag(buf, "input")
	tg.attr("type", i.inputtype)
	tg.attr("name", i.name)
	tg.attr("value", i.value)
	tg.attr("placeholder", i.placeholder)
	tg.attr("id", i.id)
	tg.attr("form", i.form)
	tg.attr("autocomplete", i.autocomplete)
	tg.attr("min", i.min)
	tg.attr("max", i.max)
	tg.attr("step", i.step)
	tg.attr("pattern", i.pattern)
	tg.intAttr("size", i.size)
	tg.optionalIntAttr("maxlength", i.maxlength)
	tg.optionalIntAttr("minlength", i.minlength)
	tg.attr("accept", i.accept)
	tg.boolAttr("required", i.required)
	tg.boolAttr("disabled", i.disabled)
	tg.boolAttr("readonly", i.readonly)
	tg.boolAttr("autofocus", i.autofocus)
	tg.boolAttr("multiple", i.multiple)
	tg.boolAttr("checked", i.checked)
	i.classStyle(tg)
	tg.void()
}

// Type sets the type attribute and returns the element.
func (i *Input) Type(inputtype string) *Input {
	i.inputtype = inputtype
	return i
}

// Name sets the name attribute and returns the element.
func (i *Input) Name(name string) *Input {
	i.name = name
	return i
}

// Value sets the value attribute and returns the element.
func (i *Input) Value(value string) *Input {
	i.value = value
	return i
}

// Placeholder sets the placeholder attribute and returns the element.
func (i *Input) Placeholder(placeholder string) *Input {
	i.placeholder = placeholder
	return i
}

// Form sets the form attribute and returns the element.
func (i *Input) Form(form string) *Input {
	i.form = form
	return i
}

// Required sets the required attribute and returns the element.
func (i *Input) Required(required bool) *Input {
	i.required = required
	return i
}

// Disabled sets the disabled attribute and returns the element.
func (i *Input) Disabled(disabled bool) *Input {
	i.disabled = disabled
	return i
}

// Readonly sets the readonly attribute and returns the element.
func (i *Input) Readonly(readonly bool) *Input {
	i.readonly = readonly
	return i
}

// Autofocus sets the autofocus attribute and returns the element.
func (i *Input) Autofocus(autofocus bool) *Input {
	i.autofocus = autofocus
	return i
}

// AutoComplete sets the autocomplete attribute and returns the element.
func (i *Input) AutoComplete(autocomplete string) *Input {
	i.autocomplete = autocomplete
	return i
}

// Min sets the min attribute and returns the element.
func (i *Input) Min(min string) *Input {
	i.min = min
	return i
}

// Max sets the max attribute and returns the element.
func (i *Input) Max(max string) *Input {
	i.max = max
	return i
}

// Step sets the step attribute and returns the element.
func (i *Input) Step(step string) *Input {
	i.step = step
	return i
}

// Pattern sets the pattern attribute and returns the element.
func (i *Input) Pattern(pattern string) *Input {
	i.pattern = pattern
	return i
}

// Size sets the size attribute, the visible width in characters. Values below
// one are omitted.
func (i *Input) Size(size int) *Input {
	i.size = size
	return i
}

// MaxLength sets the maxlength attribute. Zero is emitted; negative values
// are omitted.
func (i *Input) MaxLength(maxlength int) *Input {
	i.maxlength = setInt(maxlength)
	return i
}

// MinLength sets the minlength attribute. Zero is emitted; negative values
// are omitted.
func (i *Input) MinLength(minlength int) *Input {
	i.minlength = setInt(minlength)
	return i
}

// Multiple sets the multiple attribute and returns the element.
func (i *Input) Multiple(multiple bool) *Input {
	i.multiple = multiple
	return i
}

// Accept sets the accept attribute and returns the element.
func (i *Input) Accept(accept string) *Input {
	i.accept = accept
	return i
}

// Checked sets the checked attribute and returns the element.
func (i *Input) Checked(checked bool) *Input {
	i.checked = checked
	return i
}

// Output represents the HTML <output> element, used for calculation results.
type Output struct {
	bodyElement
	textContent[*Output]
	forattr string
	name    string
	form    string
}

// NewOutput returns a new, empty <output> element.
func NewOutput() *Output {
	v := &Output{}
	v.init(v)
	return v
}

// renderTo writes the output element's HTML to buf.
func (o *Output) renderTo(buf *bytes.Buffer) {
	tg := startTag(buf, "output")
	tg.attr("for", o.forattr)
	tg.attr("name", o.name)
	tg.attr("form", o.form)
	o.attrs(tg)
	tg.children(o.contents)
}

// For sets the for attribute and returns the element.
func (o *Output) For(forattr string) *Output {
	o.forattr = forattr
	return o
}

// Name sets the name attribute and returns the element.
func (o *Output) Name(name string) *Output {
	o.name = name
	return o
}

// Form sets the form attribute and returns the element.
func (o *Output) Form(form string) *Output {
	o.form = form
	return o
}

// Fieldset represents the HTML <fieldset> element, used for grouping form controls.
type Fieldset struct {
	bodyElement
	contentNode[*Fieldset]
	form     string
	name     string
	disabled bool
}

// NewFieldset returns a new, empty <fieldset> element.
func NewFieldset() *Fieldset {
	v := &Fieldset{}
	v.init(v)
	return v
}

// renderTo writes the fieldset element's HTML to buf.
func (f *Fieldset) renderTo(buf *bytes.Buffer) {
	tg := startTag(buf, "fieldset")
	tg.attr("form", f.form)
	tg.attr("name", f.name)
	tg.boolAttr("disabled", f.disabled)
	f.attrs(tg)
	tg.children(f.contents)
}

// Form sets the form attribute and returns the element.
func (f *Fieldset) Form(form string) *Fieldset {
	f.form = form
	return f
}

// Name sets the name attribute and returns the element.
func (f *Fieldset) Name(name string) *Fieldset {
	f.name = name
	return f
}

// Disabled sets the disabled attribute and returns the element.
func (f *Fieldset) Disabled(disabled bool) *Fieldset {
	f.disabled = disabled
	return f
}

// Button represents the HTML <button> element, used for clickable buttons.
type Button struct {
	bodyElement
	textContent[*Button]
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

// NewButton returns a new, empty <button> element.
func NewButton() *Button {
	v := &Button{}
	v.init(v)
	return v
}

// renderTo writes the button element's HTML to buf.
func (b *Button) renderTo(buf *bytes.Buffer) {
	tg := startTag(buf, "button")
	tg.attr("type", b.buttonType)
	tg.attr("name", b.name)
	tg.attr("value", b.value)
	tg.attr("form", b.form)
	tg.urlAttr("formaction", b.formAction)
	tg.attr("formenctype", b.formEnctype)
	tg.attr("formmethod", b.formMethod)
	tg.attr("formtarget", b.formTarget)
	tg.boolAttr("formnovalidate", b.formNovalidate)
	tg.boolAttr("disabled", b.disabled)
	tg.boolAttr("autofocus", b.autofocus)
	b.attrs(tg)
	tg.children(b.contents)
}

// Type sets the type attribute and returns the element.
func (b *Button) Type(buttonType string) *Button {
	b.buttonType = buttonType
	return b
}

// Name sets the name attribute and returns the element.
func (b *Button) Name(name string) *Button {
	b.name = name
	return b
}

// Value sets the value attribute and returns the element.
func (b *Button) Value(value string) *Button {
	b.value = value
	return b
}

// Form sets the form attribute and returns the element.
func (b *Button) Form(form string) *Button {
	b.form = form
	return b
}

// FormAction sets the formaction attribute and returns the element.
//
// An unsafe URL is replaced with [BlockedURL] when rendered; see [IsSafeURL].
func (b *Button) FormAction(formAction string) *Button {
	b.formAction = formAction
	return b
}

// FormEnctype sets the formenctype attribute and returns the element.
func (b *Button) FormEnctype(formEnctype string) *Button {
	b.formEnctype = formEnctype
	return b
}

// FormMethod sets the formmethod attribute and returns the element.
func (b *Button) FormMethod(formMethod string) *Button {
	b.formMethod = formMethod
	return b
}

// FormTarget sets the formtarget attribute and returns the element.
func (b *Button) FormTarget(formTarget string) *Button {
	b.formTarget = formTarget
	return b
}

// FormNovalidate sets the formnovalidate attribute and returns the element.
func (b *Button) FormNovalidate(formNovalidate bool) *Button {
	b.formNovalidate = formNovalidate
	return b
}

// Disabled sets the disabled attribute and returns the element.
func (b *Button) Disabled(disabled bool) *Button {
	b.disabled = disabled
	return b
}

// Autofocus sets the autofocus attribute and returns the element.
func (b *Button) Autofocus(autofocus bool) *Button {
	b.autofocus = autofocus
	return b
}

// Select represents the HTML <select> element, used for dropdown lists.
type Select struct {
	bodyElement
	contentNode[*Select]
	name         string
	form         string
	size         int
	multiple     bool
	required     bool
	disabled     bool
	autofocus    bool
	autoComplete string
}

// NewSelect returns a new, empty <select> element.
func NewSelect() *Select {
	v := &Select{}
	v.init(v)
	return v
}

// renderTo writes the select element's HTML to buf.
func (s *Select) renderTo(buf *bytes.Buffer) {
	tg := startTag(buf, "select")
	tg.attr("name", s.name)
	tg.attr("form", s.form)
	tg.intAttr("size", s.size)
	tg.attr("autocomplete", s.autoComplete)
	tg.boolAttr("multiple", s.multiple)
	tg.boolAttr("required", s.required)
	tg.boolAttr("disabled", s.disabled)
	tg.boolAttr("autofocus", s.autofocus)
	s.attrs(tg)
	tg.children(s.contents)
}

// Name sets the name attribute and returns the element.
func (s *Select) Name(name string) *Select {
	s.name = name
	return s
}

// Form sets the form attribute and returns the element.
func (s *Select) Form(form string) *Select {
	s.form = form
	return s
}

// Size sets the size attribute, the number of visible options. Values below
// one are omitted.
func (s *Select) Size(size int) *Select {
	s.size = size
	return s
}

// Multiple sets the multiple attribute and returns the element.
func (s *Select) Multiple(multiple bool) *Select {
	s.multiple = multiple
	return s
}

// Required sets the required attribute and returns the element.
func (s *Select) Required(required bool) *Select {
	s.required = required
	return s
}

// Disabled sets the disabled attribute and returns the element.
func (s *Select) Disabled(disabled bool) *Select {
	s.disabled = disabled
	return s
}

// Autofocus sets the autofocus attribute and returns the element.
func (s *Select) Autofocus(autofocus bool) *Select {
	s.autofocus = autofocus
	return s
}

// AutoComplete sets the autocomplete attribute and returns the element.
func (s *Select) AutoComplete(autoComplete string) *Select {
	s.autoComplete = autoComplete
	return s
}

// Datalist represents the HTML <datalist> element, used for predefined options.
type Datalist struct {
	bodyElement
	contentNode[*Datalist]
}

// NewDatalist returns a new, empty <datalist> element.
func NewDatalist() *Datalist {
	v := &Datalist{}
	v.init(v)
	return v
}

// renderTo writes the datalist element's HTML to buf.
func (d *Datalist) renderTo(buf *bytes.Buffer) {
	tg := startTag(buf, "datalist")
	d.attrs(tg)
	tg.children(d.contents)
}

// Optgroup represents the HTML <optgroup> element, used for grouping options.
type Optgroup struct {
	bodyElement
	contentNode[*Optgroup]
	label    string
	disabled bool
}

// NewOptgroup returns a new, empty <optgroup> element.
func NewOptgroup() *Optgroup {
	v := &Optgroup{}
	v.init(v)
	return v
}

// renderTo writes the optgroup element's HTML to buf.
func (o *Optgroup) renderTo(buf *bytes.Buffer) {
	tg := startTag(buf, "optgroup")
	tg.attr("label", o.label)
	tg.boolAttr("disabled", o.disabled)
	o.attrs(tg)
	tg.children(o.contents)
}

// Label sets the label attribute and returns the element.
func (o *Optgroup) Label(label string) *Optgroup {
	o.label = label
	return o
}

// Disabled sets the disabled attribute and returns the element.
func (o *Optgroup) Disabled(disabled bool) *Optgroup {
	o.disabled = disabled
	return o
}

// Option represents the HTML <option> element, used for select options.
type Option struct {
	bodyElement
	textContent[*Option]
	value    string
	label    string
	selected bool
	disabled bool
}

// NewOption returns a new, empty <option> element.
func NewOption() *Option {
	v := &Option{}
	v.init(v)
	return v
}

// renderTo writes the option element's HTML to buf.
func (o *Option) renderTo(buf *bytes.Buffer) {
	tg := startTag(buf, "option")
	tg.attr("value", o.value)
	tg.attr("label", o.label)
	tg.boolAttr("selected", o.selected)
	tg.boolAttr("disabled", o.disabled)
	o.attrs(tg)
	tg.children(o.contents)
}

// Value sets the value attribute and returns the element.
func (o *Option) Value(value string) *Option {
	o.value = value
	return o
}

// Label sets the label attribute and returns the element.
func (o *Option) Label(label string) *Option {
	o.label = label
	return o
}

// Selected sets the selected attribute and returns the element.
func (o *Option) Selected(selected bool) *Option {
	o.selected = selected
	return o
}

// Disabled sets the disabled attribute and returns the element.
func (o *Option) Disabled(disabled bool) *Option {
	o.disabled = disabled
	return o
}

// Textarea represents the HTML <textarea> element, used for multi-line text input.
type Textarea struct {
	bodyElement
	textContent[*Textarea]
	name         string
	form         string
	rows         int
	cols         int
	placeholder  string
	maxLength    optionalInt
	minLength    optionalInt
	wrap         string
	required     bool
	disabled     bool
	readonly     bool
	autofocus    bool
	autoComplete string
	spellcheck   string
}

// NewTextarea returns a new, empty <textarea> element.
func NewTextarea() *Textarea {
	v := &Textarea{}
	v.init(v)
	return v
}

// renderTo writes the textarea element's HTML to buf.
func (t *Textarea) renderTo(buf *bytes.Buffer) {
	tg := startTag(buf, "textarea")
	tg.attr("name", t.name)
	tg.attr("form", t.form)
	tg.intAttr("rows", t.rows)
	tg.intAttr("cols", t.cols)
	tg.attr("placeholder", t.placeholder)
	tg.optionalIntAttr("maxlength", t.maxLength)
	tg.optionalIntAttr("minlength", t.minLength)
	tg.attr("wrap", t.wrap)
	tg.attr("autocomplete", t.autoComplete)
	tg.attr("spellcheck", t.spellcheck)
	tg.boolAttr("required", t.required)
	tg.boolAttr("disabled", t.disabled)
	tg.boolAttr("readonly", t.readonly)
	tg.boolAttr("autofocus", t.autofocus)
	t.attrs(tg)
	tg.children(t.contents)
}

// Name sets the name attribute and returns the element.
func (t *Textarea) Name(name string) *Textarea {
	t.name = name
	return t
}

// Form sets the form attribute and returns the element.
func (t *Textarea) Form(form string) *Textarea {
	t.form = form
	return t
}

// Rows sets the rows attribute, the visible number of lines. Values below one
// are omitted.
func (t *Textarea) Rows(rows int) *Textarea {
	t.rows = rows
	return t
}

// Cols sets the cols attribute, the visible width in characters. Values below
// one are omitted.
func (t *Textarea) Cols(cols int) *Textarea {
	t.cols = cols
	return t
}

// Placeholder sets the placeholder attribute and returns the element.
func (t *Textarea) Placeholder(placeholder string) *Textarea {
	t.placeholder = placeholder
	return t
}

// MaxLength sets the maxlength attribute. Zero is emitted; negative values
// are omitted.
func (t *Textarea) MaxLength(maxLength int) *Textarea {
	t.maxLength = setInt(maxLength)
	return t
}

// MinLength sets the minlength attribute. Zero is emitted; negative values
// are omitted.
func (t *Textarea) MinLength(minLength int) *Textarea {
	t.minLength = setInt(minLength)
	return t
}

// Wrap sets the wrap attribute and returns the element.
func (t *Textarea) Wrap(wrap string) *Textarea {
	t.wrap = wrap
	return t
}

// Required sets the required attribute and returns the element.
func (t *Textarea) Required(required bool) *Textarea {
	t.required = required
	return t
}

// Disabled sets the disabled attribute and returns the element.
func (t *Textarea) Disabled(disabled bool) *Textarea {
	t.disabled = disabled
	return t
}

// Readonly sets the readonly attribute and returns the element.
func (t *Textarea) Readonly(readonly bool) *Textarea {
	t.readonly = readonly
	return t
}

// Autofocus sets the autofocus attribute and returns the element.
func (t *Textarea) Autofocus(autofocus bool) *Textarea {
	t.autofocus = autofocus
	return t
}

// AutoComplete sets the autocomplete attribute and returns the element.
func (t *Textarea) AutoComplete(autoComplete string) *Textarea {
	t.autoComplete = autoComplete
	return t
}

// Spellcheck sets the spellcheck attribute and returns the element.
func (t *Textarea) Spellcheck(spellcheck string) *Textarea {
	t.spellcheck = spellcheck
	return t
}

// Progress represents the HTML <progress> element, used for showing completion progress.
type Progress struct {
	bodyElement
	textContent[*Progress]
	value string
	max   string
	form  string
}

// NewProgress returns a new, empty <progress> element.
func NewProgress() *Progress {
	v := &Progress{}
	v.init(v)
	return v
}

// renderTo writes the progress element's HTML to buf.
func (p *Progress) renderTo(buf *bytes.Buffer) {
	tg := startTag(buf, "progress")
	tg.attr("value", p.value)
	tg.attr("max", p.max)
	tg.attr("form", p.form)
	p.attrs(tg)
	tg.children(p.contents)
}

// Value sets the value attribute and returns the element.
func (p *Progress) Value(value string) *Progress {
	p.value = value
	return p
}

// Max sets the max attribute and returns the element.
func (p *Progress) Max(max string) *Progress {
	p.max = max
	return p
}

// Form sets the form attribute and returns the element.
func (p *Progress) Form(form string) *Progress {
	p.form = form
	return p
}

// Meter represents the HTML <meter> element, used for displaying scalar measurements.
type Meter struct {
	bodyElement
	textContent[*Meter]
	value   string
	min     string
	max     string
	low     string
	high    string
	optimum string
	form    string
}

// NewMeter returns a new, empty <meter> element.
func NewMeter() *Meter {
	v := &Meter{}
	v.init(v)
	return v
}

// renderTo writes the meter element's HTML to buf.
func (m *Meter) renderTo(buf *bytes.Buffer) {
	tg := startTag(buf, "meter")
	tg.attr("value", m.value)
	tg.attr("min", m.min)
	tg.attr("max", m.max)
	tg.attr("low", m.low)
	tg.attr("high", m.high)
	tg.attr("optimum", m.optimum)
	tg.attr("form", m.form)
	m.attrs(tg)
	tg.children(m.contents)
}

// Value sets the value attribute and returns the element.
func (m *Meter) Value(value string) *Meter {
	m.value = value
	return m
}

// Min sets the min attribute and returns the element.
func (m *Meter) Min(min string) *Meter {
	m.min = min
	return m
}

// Max sets the max attribute and returns the element.
func (m *Meter) Max(max string) *Meter {
	m.max = max
	return m
}

// Low sets the low attribute and returns the element.
func (m *Meter) Low(low string) *Meter {
	m.low = low
	return m
}

// High sets the high attribute and returns the element.
func (m *Meter) High(high string) *Meter {
	m.high = high
	return m
}

// Optimum sets the optimum attribute and returns the element.
func (m *Meter) Optimum(optimum string) *Meter {
	m.optimum = optimum
	return m
}

// Form sets the form attribute and returns the element.
func (m *Meter) Form(form string) *Meter {
	m.form = form
	return m
}

// Legend represents the HTML <legend> element, used for fieldset captions.
type Legend struct {
	bodyElement
	textContent[*Legend]
}

// NewLegend returns a new, empty <legend> element.
func NewLegend() *Legend {
	v := &Legend{}
	v.init(v)
	return v
}

// renderTo writes the legend element's HTML to buf.
func (l *Legend) renderTo(buf *bytes.Buffer) {
	tg := startTag(buf, "legend")
	l.attrs(tg)
	tg.children(l.contents)
}
