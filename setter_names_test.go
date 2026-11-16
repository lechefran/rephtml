package rephtml

import (
	"reflect"
	"testing"
)

// TestRenamedSettersOutput pins the attributes written by the setters that
// replaced the names removed in v0.5.0.
func TestRenamedSettersOutput(t *testing.T) {
	tests := []struct {
		name string
		got  Element
		want string
	}{
		{"Form.AutoComplete", NewForm().AutoComplete("off"), `<form autocomplete="off"></form>`},
		{"Form.NoValidate", NewForm().NoValidate(true), `<form novalidate></form>`},
		{"Input.AutoComplete", NewInput().AutoComplete("email"), `<input autocomplete="email">`},
		{"Input.MaxLength", NewInput().MaxLength(10), `<input maxlength="10">`},
		{"Input.MinLength", NewInput().MinLength(2), `<input minlength="2">`},
		{"Iframe.AllowFullscreen", NewIframe().AllowFullscreen(true), `<iframe allowfullscreen></iframe>`},
		{"Iframe.ReferrerPolicy", NewIframe().ReferrerPolicy("no-referrer"), `<iframe referrerpolicy="no-referrer"></iframe>`},
		{"Portal.ReferrerPolicy", NewPortal().ReferrerPolicy("no-referrer"), `<portal referrerpolicy="no-referrer"></portal>`},
		{"Script.CrossOrigin", NewScript().CrossOrigin("anonymous"), `<script crossorigin="anonymous"></script>`},
		{"Script.ReferrerPolicy", NewScript().ReferrerPolicy("origin"), `<script referrerpolicy="origin"></script>`},
		{"Td.Style", NewTd().Style(StyleMap{"color": "red"}), `<td style="color: red;"></td>`},
		{"Table.Id", NewTable().Id("x"), `<table id="x"></table>`},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.got.HTML(); got != tt.want {
				t.Fatalf("HTML() = %q, want %q", got, tt.want)
			}
		})
	}
}

// TestDeprecatedNamesRemoved checks that the names deprecated in v0.2.0 and
// v0.3.0 are gone, so they are not brought back by accident.
func TestDeprecatedNamesRemoved(t *testing.T) {
	removed := []struct {
		element Element
		names   []string
	}{
		{NewForm(), []string{"Autocomplete", "Novalidate"}},
		{NewInput(), []string{"Autocomplete", "Maxlength", "Minlength"}},
		{NewIframe(), []string{"Allowfullscreen", "Referrerpolicy"}},
		{NewPortal(), []string{"Referrerpolicy"}},
		{NewScript(), []string{"Crossorigin", "Referrerpolicy"}},
		{NewTable(), []string{"AddId", "Styles"}},
		{NewCaption(), []string{"AddId", "Styles"}},
		{NewCol(), []string{"AddId", "Styles"}},
		{NewColgroup(), []string{"AddId", "Styles"}},
		{NewThead(), []string{"AddId", "Styles"}},
		{NewTbody(), []string{"AddId", "Styles"}},
		{NewTfoot(), []string{"AddId", "Styles"}},
		{NewTr(), []string{"AddId", "Styles"}},
		{NewTd(), []string{"AddId", "Styles"}},
		{NewTh(), []string{"AddId", "Styles"}},
	}
	for _, r := range removed {
		v := reflect.ValueOf(r.element)
		for _, name := range r.names {
			if v.MethodByName(name).IsValid() {
				t.Errorf("%s still has %s", v.Type(), name)
			}
		}
	}
}
