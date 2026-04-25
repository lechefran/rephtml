package tests

import (
	"testing"

	rephtml "github.com/lechefran/rephtml/components"
)

func TestContainerAddRendersUnpreparedChildren(t *testing.T) {
	header := rephtml.NewHeader().Add(rephtml.NewH1().Text("Dashboard"))
	header.Prepare()

	want := "<header><h1>Dashboard</h1></header>"
	if got := string(header.Bytes()); got != want {
		t.Fatalf("unexpected header output:\ngot  %q\nwant %q", got, want)
	}
}

func TestContainerRendersCurrentChildState(t *testing.T) {
	paragraph := rephtml.NewP().Text("Draft")
	section := rephtml.NewSection().Add(paragraph)

	paragraph.Text("Final")
	section.Prepare()

	want := "<section><p>Final</p></section>"
	if got := string(section.Bytes()); got != want {
		t.Fatalf("unexpected section output:\ngot  %q\nwant %q", got, want)
	}
}

func TestTableRendersCurrentNestedChildState(t *testing.T) {
	span := rephtml.NewSpan().Text("Draft")
	cell := rephtml.NewTd().Add(span)
	row := rephtml.NewTr().AddTd(cell)
	body := rephtml.NewTbody().AddTr(row)
	table := rephtml.NewTable().AddTbody(body)

	span.Text("Final")
	table.Prepare()

	want := "<table><tbody><tr><td><span>Final</span></td></tr></tbody></table>"
	if got := string(table.Bytes()); got != want {
		t.Fatalf("unexpected table output:\ngot  %q\nwant %q", got, want)
	}
}

func TestBytesReturnsSnapshot(t *testing.T) {
	paragraph := rephtml.NewP().Text("Stable")
	paragraph.Prepare()

	snapshot := paragraph.Bytes()
	snapshot[3] = 'x'

	want := "<p>Stable</p>"
	if got := string(paragraph.Bytes()); got != want {
		t.Fatalf("Bytes returned mutable backing storage:\ngot  %q\nwant %q", got, want)
	}
}
