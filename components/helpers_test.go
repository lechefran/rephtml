package rephtml

import "testing"

// TestStripRemovesWhitespace provides TestStripRemovesWhitespace behavior for the package.
func TestStripRemovesWhitespace(t *testing.T) {
	input := []byte(" <tr>\n\t<td>Alpha</td> \r\n <td>Beta</td>\t</tr> ")
	got := string(strip(input))
	want := "<tr><td>Alpha</td><td>Beta</td></tr>"

	if got != want {
		t.Fatalf("strip() = %q, want %q", got, want)
	}
}

// TestParseStyleSortsKeys provides TestParseStyleSortsKeys behavior for the package.
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

func TestStyleCopiesInputMap(t *testing.T) {
	styles := map[string]string{
		"color":       "red",
		"font-weight": "700",
	}

	paragraph := NewP().Text("Styled").Style(styles)
	styles["color"] = "blue"
	styles["margin"] = "12px"

	paragraph.Prepare()

	want := `<p style="color: red; font-weight: 700;">Styled</p>`
	if got := string(paragraph.Bytes()); got != want {
		t.Fatalf("unexpected style output after caller map mutation:\ngot  %q\nwant %q", got, want)
	}
}

func TestCloneStyleMapReturnsIndependentMap(t *testing.T) {
	styles := map[string]string{"color": "red"}
	clone := cloneStyleMap(styles)
	styles["color"] = "blue"

	if clone["color"] != "red" {
		t.Fatalf("cloneStyleMap retained caller map, got color %q", clone["color"])
	}
}
