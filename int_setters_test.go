package rephtml

import "testing"

// TestWholeNumberSetters checks the int setters for attributes that are always
// whole numbers: size, rows and cols are omitted below one, while maxlength
// and minlength allow zero once set and omit negative values.
func TestWholeNumberSetters(t *testing.T) {
	tests := []struct {
		name string
		got  string
		want string
	}{
		{
			name: "Input positive values keep their position",
			got:  NewInput().Type("text").Size(20).MaxLength(10).MinLength(2).AddClass("c").HTML(),
			want: `<input type="text" size="20" maxlength="10" minlength="2" class="c">`,
		},
		{
			name: "Input zero size omitted, zero lengths emitted",
			got:  NewInput().Size(0).MaxLength(0).MinLength(0).HTML(),
			want: `<input maxlength="0" minlength="0">`,
		},
		{
			name: "Input negative values omitted",
			got:  NewInput().Size(-1).MaxLength(-1).MinLength(-5).HTML(),
			want: `<input>`,
		},
		{
			name: "Input unset lengths omitted",
			got:  NewInput().HTML(),
			want: `<input>`,
		},
		{
			name: "Select size",
			got:  NewSelect().Size(4).HTML(),
			want: `<select size="4"></select>`,
		},
		{
			name: "Select zero size omitted",
			got:  NewSelect().Size(0).HTML(),
			want: `<select></select>`,
		},
		{
			name: "Textarea values",
			got:  NewTextarea().Rows(4).Cols(40).MaxLength(200).MinLength(0).HTML(),
			want: `<textarea rows="4" cols="40" maxlength="200" minlength="0"></textarea>`,
		},
		{
			name: "Textarea non-positive rows and cols omitted",
			got:  NewTextarea().Rows(0).Cols(-3).MaxLength(-1).HTML(),
			want: `<textarea></textarea>`,
		},
		{
			name: "last call wins",
			got:  NewTextarea().MaxLength(5).MaxLength(-1).HTML(),
			want: `<textarea></textarea>`,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if tt.got != tt.want {
				t.Fatalf("got %q, want %q", tt.got, tt.want)
			}
		})
	}
}
