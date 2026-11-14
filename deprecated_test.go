package rephtml

import "testing"

// The deprecated setters are kept as forwarders so existing callers keep
// compiling. These tests pin that the old and new names stay equivalent, and
// the new names produce the documented attributes. The old names are called on
// purpose, so staticcheck's SA1019 is silenced for the file.

//lint:file-ignore SA1019 these tests exist to exercise the deprecated names.

func TestDeprecatedSettersMatchCanonicalNames(t *testing.T) {
	tests := []struct {
		name       string
		canonical  Element
		deprecated Element
		want       string
	}{
		{
			"Form.AutoComplete",
			NewForm().AutoComplete("off"),
			NewForm().Autocomplete("off"),
			`<form autocomplete="off"></form>`,
		},
		{
			"Form.NoValidate",
			NewForm().NoValidate(true),
			NewForm().Novalidate(true),
			`<form novalidate></form>`,
		},
		{
			"Input.AutoComplete",
			NewInput().AutoComplete("email"),
			NewInput().Autocomplete("email"),
			`<input autocomplete="email">`,
		},
		{
			"Input.MaxLength",
			NewInput().MaxLength("10"),
			NewInput().Maxlength("10"),
			`<input maxlength="10">`,
		},
		{
			"Input.MinLength",
			NewInput().MinLength("2"),
			NewInput().Minlength("2"),
			`<input minlength="2">`,
		},
		{
			"Iframe.AllowFullscreen",
			NewIframe().AllowFullscreen(true),
			NewIframe().Allowfullscreen(true),
			`<iframe allowfullscreen></iframe>`,
		},
		{
			"Iframe.ReferrerPolicy",
			NewIframe().ReferrerPolicy("no-referrer"),
			NewIframe().Referrerpolicy("no-referrer"),
			`<iframe referrerpolicy="no-referrer"></iframe>`,
		},
		{
			"Portal.ReferrerPolicy",
			NewPortal().ReferrerPolicy("no-referrer"),
			NewPortal().Referrerpolicy("no-referrer"),
			`<portal referrerpolicy="no-referrer"></portal>`,
		},
		{
			"Script.CrossOrigin",
			NewScript().CrossOrigin("anonymous"),
			NewScript().Crossorigin("anonymous"),
			`<script crossorigin="anonymous"></script>`,
		},
		{
			"Script.ReferrerPolicy",
			NewScript().ReferrerPolicy("origin"),
			NewScript().Referrerpolicy("origin"),
			`<script referrerpolicy="origin"></script>`,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.canonical.HTML(); got != tt.want {
				t.Errorf("canonical HTML() = %q, want %q", got, tt.want)
			}
			if got := tt.deprecated.HTML(); got != tt.want {
				t.Errorf("deprecated HTML() = %q, want %q", got, tt.want)
			}
		})
	}
}
