package tests

import (
	"testing"

	rephtml "github.com/lechefran/rephtml/components"
)

func TestHgroupPrepare(t *testing.T) {
	var _ rephtml.BodyElement = (*rephtml.Hgroup)(nil)

	heading := rephtml.NewH1().Text("Frankenstein")
	heading.Prepare()
	subtitle := rephtml.NewP().Text("Or: The Modern Prometheus")
	subtitle.Prepare()

	hgroup := rephtml.NewHgroup().
		AddStyle("text-align", "right").
		Add(heading).
		Add(subtitle)
	hgroup.Prepare()

	want := `<hgroup style="text-align: right;"><h1>Frankenstein</h1><p>Or: The Modern Prometheus</p></hgroup>`
	if got := string(hgroup.Bytes()); got != want {
		t.Fatalf("unexpected hgroup output:\ngot  %q\nwant %q", got, want)
	}
}

func TestSearchPrepare(t *testing.T) {
	var _ rephtml.BodyElement = (*rephtml.Search)(nil)

	input := rephtml.NewInput().
		Type("search").
		Name("q").
		Placeholder("Search")

	form := rephtml.NewForm().
		Action("/search").
		Method("get").
		Add(input)
	form.Prepare()

	search := rephtml.NewSearch().Add(form)
	search.Prepare()

	want := `<search><form action="/search" method="get"><input type="search" name="q" placeholder="Search"></form></search>`
	if got := string(search.Bytes()); got != want {
		t.Fatalf("unexpected search output:\ngot  %q\nwant %q", got, want)
	}
}
