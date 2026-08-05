package rephtml

import (
	"sync"
	"testing"
)

// Rendering is read-only with respect to the element tree: an element writes to
// the buffer it is handed and keeps no render state of its own. These tests are
// the ones that matter under -race, which is how the suite is expected to run
// in CI.

// TestConcurrentRenderOfSameElement renders one element from many goroutines.
// Building a fragment once and rendering it per request is the most natural way
// to use this package, so it has to be safe.
func TestConcurrentRenderOfSameElement(t *testing.T) {
	element := NewDiv().AddStyle("color", "red").
		Add(NewSection().Add(NewH1().Text("Title")).Add(NewP().Text("Body & <text>"))).
		Add(NewUl().Add(NewLi().Add(NewAnchor().Link("/a").Text("one"))))

	want := element.HTML()

	var wg sync.WaitGroup
	results := make([]string, 64)
	for i := range results {
		wg.Add(1)
		go func(i int) {
			defer wg.Done()
			results[i] = element.HTML()
		}(i)
	}
	wg.Wait()

	for i, got := range results {
		if got != want {
			t.Fatalf("goroutine %d rendered %q, want %q", i, got, want)
		}
	}
}

// TestConcurrentRenderOfSharedChild renders two parents that share one child at
// the same time. Under the old model each parent re-rendered the child into the
// child's own buffer, so concurrent parents corrupted each other's output
// without necessarily tripping the race detector.
func TestConcurrentRenderOfSharedChild(t *testing.T) {
	shared := NewP().Text("shared")
	left := NewDiv().Add(NewH1().Text("left")).Add(shared)
	right := NewSection().Add(NewH2().Text("right")).Add(shared)

	wantLeft := `<div><h1>left</h1><p>shared</p></div>`
	wantRight := `<section><h2>right</h2><p>shared</p></section>`

	var wg sync.WaitGroup
	for range 64 {
		wg.Add(2)
		go func() {
			defer wg.Done()
			if got := left.HTML(); got != wantLeft {
				t.Errorf("left rendered %q, want %q", got, wantLeft)
			}
		}()
		go func() {
			defer wg.Done()
			if got := right.HTML(); got != wantRight {
				t.Errorf("right rendered %q, want %q", got, wantRight)
			}
		}()
	}
	wg.Wait()
}

// TestConcurrentDocumentRender renders a whole document concurrently, covering
// the hand-written HtmlFile path and the formatter alongside the elements.
func TestConcurrentDocumentRender(t *testing.T) {
	doc := NewHtmlFile().Lang("en")
	doc.AddToHead(NewTitle().Text("Report"))
	doc.AddToHead(NewMeta().Charset("utf-8"))
	doc.AddToBody(NewMain().Add(NewH1().Text("Report")).Add(NewP().Text("Body")))

	wantCompact, err := doc.RenderString()
	if err != nil {
		t.Fatalf("RenderString: %v", err)
	}
	wantFormatted, err := doc.RenderFormattedString()
	if err != nil {
		t.Fatalf("RenderFormattedString: %v", err)
	}

	var wg sync.WaitGroup
	for range 32 {
		wg.Add(2)
		go func() {
			defer wg.Done()
			got, err := doc.RenderString()
			if err != nil || got != wantCompact {
				t.Errorf("RenderString = %q, %v; want %q, nil", got, err, wantCompact)
			}
		}()
		go func() {
			defer wg.Done()
			got, err := doc.RenderFormattedString()
			if err != nil || got != wantFormatted {
				t.Errorf("RenderFormattedString = %q, %v", got, err)
			}
		}()
	}
	wg.Wait()
}
