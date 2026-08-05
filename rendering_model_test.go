package rephtml

import (
	"testing"
)

func TestElementRenderMethodsPrepareInternally(t *testing.T) {
	paragraph := NewP().Text("Hello")

	want := "<p>Hello</p>"
	if got := paragraph.HTML(); got != want {
		t.Fatalf("HTML rendered unexpected output:\ngot  %q\nwant %q", got, want)
	}
	if got := string(paragraph.Render()); got != want {
		t.Fatalf("Render rendered unexpected output:\ngot  %q\nwant %q", got, want)
	}
	// Rendering twice must not double the output, which is what a prepare that
	// failed to reset its buffer would produce.
	if got := paragraph.HTML(); got != want {
		t.Fatalf("second HTML rendered unexpected output:\ngot  %q\nwant %q", got, want)
	}
}

func TestContainerAddRendersUnpreparedChildren(t *testing.T) {
	header := NewHeader().Add(NewH1().Text("Dashboard"))

	want := "<header><h1>Dashboard</h1></header>"
	if got := header.HTML(); got != want {
		t.Fatalf("unexpected header output:\ngot  %q\nwant %q", got, want)
	}
}

func TestContainerRendersCurrentChildState(t *testing.T) {
	paragraph := NewP().Text("Draft")
	section := NewSection().Add(paragraph)

	paragraph.Text("Final")

	want := "<section><p>Final</p></section>"
	if got := section.HTML(); got != want {
		t.Fatalf("unexpected section output:\ngot  %q\nwant %q", got, want)
	}
}

func TestTableRendersCurrentNestedChildState(t *testing.T) {
	span := NewSpan().Text("Draft")
	cell := NewTd().Add(span)
	row := NewTr().AddTd(cell)
	body := NewTbody().AddTr(row)
	table := NewTable().AddTbody(body)

	span.Text("Final")

	want := "<table><tbody><tr><td><span>Final</span></td></tr></tbody></table>"
	if got := table.HTML(); got != want {
		t.Fatalf("unexpected table output:\ngot  %q\nwant %q", got, want)
	}
}

func TestRenderReturnsSnapshot(t *testing.T) {
	paragraph := NewP().Text("Stable")

	snapshot := paragraph.Render()
	snapshot[3] = 'x'

	want := "<p>Stable</p>"
	if got := paragraph.HTML(); got != want {
		t.Fatalf("Render returned mutable backing storage:\ngot  %q\nwant %q", got, want)
	}
}
