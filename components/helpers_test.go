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
