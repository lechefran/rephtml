package rephtml

import "testing"

// TestAnchorEscapesTextAndAttributes provides TestAnchorEscapesTextAndAttributes behavior for the package.
func TestAnchorEscapesTextAndAttributes(t *testing.T) {
	anchor := NewAnchor().
		Link(`https://example.test/search?q="cats"&tag=<pet>`).
		Text(`A&B < C > "quote"`)

	anchor.Prepare()

	want := `<a href="https://example.test/search?q=&#34;cats&#34;&amp;tag=&lt;pet&gt;">A&amp;B &lt; C &gt; &#34;quote&#34;</a>`
	if got := string(anchor.Bytes()); got != want {
		t.Fatalf("unexpected anchor html:\ngot  %q\nwant %q", got, want)
	}
}

// TestTableEscapesLegacyHeadersRowsAndAttributes provides TestTableEscapesLegacyHeadersRowsAndAttributes behavior for the package.
func TestTableEscapesLegacyHeadersRowsAndAttributes(t *testing.T) {
	table := NewTable().
		Id(`sales"2026`).
		AddClass(`rank&lead`).
		Headers([]string{`Name <display>`, `Notes & "flags"`}).
		AddRow([]string{`A < B`, `Tom & "Jerry"`})

	table.Prepare()

	want := `<table id="sales&#34;2026" class="rank&amp;lead"><tr><th>Name &lt;display&gt;</th><th>Notes &amp; &#34;flags&#34;</th></tr><tr><td>A &lt; B</td><td>Tom &amp; &#34;Jerry&#34;</td></tr></table>`
	if got := string(table.Bytes()); got != want {
		t.Fatalf("unexpected table html:\ngot  %q\nwant %q", got, want)
	}
}

// TestContainerTextHelpersEscapeText provides TestContainerTextHelpersEscapeText behavior for the package.
func TestContainerTextHelpersEscapeText(t *testing.T) {
	button := NewButton().Text(`<Save & Close>`)
	button.Prepare()

	want := `<button>&lt;Save &amp; Close&gt;</button>`
	if got := string(button.Bytes()); got != want {
		t.Fatalf("unexpected button html:\ngot  %q\nwant %q", got, want)
	}
}

// TestStyleElementTextRemainsRaw provides TestStyleElementTextRemainsRaw behavior for the package.
func TestStyleElementTextRemainsRaw(t *testing.T) {
	style := NewStyleElement().Text(`.a > .b { content: "x & y"; }`)
	style.Prepare()

	want := `<style>.a > .b { content: "x & y"; }</style>`
	if got := string(style.Bytes()); got != want {
		t.Fatalf("unexpected style html:\ngot  %q\nwant %q", got, want)
	}
}
