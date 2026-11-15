package rephtml

import (
	"bytes"
	"strings"
	"sync"
	"testing"
)

// The tests in this file cover Head and Body wrappers added to an HtmlFile:
// the document reads them when it renders, not when they are added.

func renderDoc(t *testing.T, doc *HtmlFile) string {
	t.Helper()
	got, err := doc.RenderString()
	if err != nil {
		t.Fatalf("RenderString() error = %v", err)
	}
	return got
}

func TestWrapperContentAddedAfterMergeIsRendered(t *testing.T) {
	head := NewHead()
	body := NewBody()
	doc := NewHtmlFile().AddToHead(head).AddToBody(body)

	head.Add(NewTitle().Text("T"))
	body.Add(NewP().Text("x"))

	want := `<!DOCTYPE html><html><head><title>T</title></head><body><p>x</p></body></html>`
	if got := renderDoc(t, doc); got != want {
		t.Fatalf("got %q, want %q", got, want)
	}
}

func TestWrapperContentKeepsItsPlaceAmongDirectContent(t *testing.T) {
	body := NewBody()
	doc := NewHtmlFile().
		AddToBody(NewP().Text("a")).
		AddToBody(body).
		AddToBody(NewP().Text("c"))
	body.Add(NewP().Text("b"))

	want := `<!DOCTYPE html><html><body><p>a</p><p>b</p><p>c</p></body></html>`
	if got := renderDoc(t, doc); got != want {
		t.Fatalf("got %q, want %q", got, want)
	}
}

func TestWrapperAttributesSetAfterMergeAreRendered(t *testing.T) {
	head := NewHead()
	body := NewBody()
	doc := NewHtmlFile().AddToHead(head).AddToBody(body)

	head.Id("h")
	body.OnLoad("init()").OnUnload("bye()").Id("b").AddClass("page").AddStyle("margin", "0")

	want := `<!DOCTYPE html><html><head id="h"></head>` +
		`<body onload="init()" onunload="bye()" id="b" class="page" style="margin: 0;"></body></html>`
	if got := renderDoc(t, doc); got != want {
		t.Fatalf("got %q, want %q", got, want)
	}
}

func TestWrapperErrorRecordedAfterMergeIsReported(t *testing.T) {
	tests := []struct {
		name  string
		build func() (*HtmlFile, func())
	}{
		{
			name: "head",
			build: func() (*HtmlFile, func()) {
				head := NewHead()
				return NewHtmlFile().AddToHead(head), func() { head.Add(NewDiv()) }
			},
		},
		{
			name: "body",
			build: func() (*HtmlFile, func()) {
				body := NewBody()
				return NewHtmlFile().AddToBody(body), func() { body.Add(NewTitle()) }
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			doc, addInvalid := tt.build()
			if err := doc.Err(); err != nil {
				t.Fatalf("Err() before the invalid add = %v, want nil", err)
			}
			addInvalid()

			if doc.Err() == nil {
				t.Fatal("Err() = nil after the wrapper recorded an error")
			}
			if _, err := doc.Render(); err == nil {
				t.Fatal("Render() error = nil, want the wrapper's error")
			}
			var buf bytes.Buffer
			if _, err := doc.WriteTo(&buf); err == nil || buf.Len() != 0 {
				t.Fatalf("WriteTo() = %q, %v; want no output and the wrapper's error", buf.String(), err)
			}
		})
	}
}

func TestWrapperErrorRecordedBeforeMergeIsReported(t *testing.T) {
	head := NewHead().Add(NewDiv())
	doc := NewHtmlFile().AddToHead(head)

	if err := doc.Err(); err == nil || !strings.Contains(err.Error(), "cannot add Div to head element") {
		t.Fatalf("Err() = %v, want the head's error", err)
	}
}

func TestLaterWrapperStyleOverridesEarlier(t *testing.T) {
	first := NewBody().AddStyle("color", "red").AddStyle("margin", "0")
	second := NewBody().AddStyle("color", "blue")
	doc := NewHtmlFile().AddToBody(first).AddToBody(second)

	want := `<!DOCTYPE html><html><body style="color: blue; margin: 0;"></body></html>`
	if got := renderDoc(t, doc); got != want {
		t.Fatalf("got %q, want %q", got, want)
	}
	// Layering the second wrapper's style must not write into the first's map.
	if got, want := first.HTML(), `<body style="color: red; margin: 0;"></body>`; got != want {
		t.Fatalf("first wrapper after render = %q, want %q", got, want)
	}
}

func TestEmptyWrappersWriteNoSection(t *testing.T) {
	doc := NewHtmlFile().AddToHead(NewHead()).AddToBody(NewBody())

	if got, want := renderDoc(t, doc), `<!DOCTYPE html><html></html>`; got != want {
		t.Fatalf("got %q, want %q", got, want)
	}
}

func TestConcurrentRenderWithLayeredWrappers(t *testing.T) {
	doc := NewHtmlFile().
		AddToBody(NewBody().AddStyle("color", "red").AddClass("a")).
		AddToBody(NewBody().AddStyle("margin", "0").AddClass("b").Add(NewP().Text("x")))
	want := renderDoc(t, doc)

	var wg sync.WaitGroup
	for range 8 {
		wg.Go(func() {
			for range 50 {
				if got, err := doc.RenderString(); err != nil || got != want {
					t.Errorf("RenderString() = %q, %v; want %q", got, err, want)
					return
				}
			}
		})
	}
	wg.Wait()
}
