package rephtml

import "testing"

func TestStripRemovesWhitespace(t *testing.T) {
	input := []byte(" <tr>\n\t<td>Alpha</td> \r\n <td>Beta</td>\t</tr> ")
	got := string(strip(input))
	want := "<tr><td>Alpha</td><td>Beta</td></tr>"

	if got != want {
		t.Fatalf("strip() = %q, want %q", got, want)
	}
}

func TestParseStyleSortsKeys(t *testing.T) {
	paragraph := NewP().Text("Sorted").Style(map[string]string{
		"z-index": "1",
		"color":   "red",
		"border":  `1px solid "black"`,
	})
	paragraph.Prepare()

	want := `<p style="border: 1px solid &#34;black&#34;; color: red; z-index: 1;">Sorted</p>`
	if got := string(paragraph.Bytes()); got != want {
		t.Fatalf("unexpected style output:\ngot  %q\nwant %q", got, want)
	}
}
