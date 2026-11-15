package rephtml

import (
	"bytes"
	"errors"
	"io"
	"sync"
	"testing"
)

func sampleDocument() *HtmlFile {
	return NewHtmlFile().Lang("en").
		AddToHead(NewTitle().Text("Report")).
		AddToBody(NewDiv().Add(NewP().Text("a < b & c")).Add(NewUl().Add(NewLi().Add(NewSpan().Text("one")))))
}

// failingWriter fails every write with a fixed error.
type failingWriter struct{ err error }

func (w failingWriter) Write([]byte) (int, error) { return 0, w.err }

// shortWriter accepts one byte fewer than it is given without reporting an
// error, which io.Writer forbids and WriteFormattedTo has to catch.
type shortWriter struct{}

func (shortWriter) Write(p []byte) (int, error) { return len(p) - 1, nil }

func TestWriteToMatchesRender(t *testing.T) {
	doc := sampleDocument()
	want, err := doc.Render()
	if err != nil {
		t.Fatal(err)
	}

	var buf bytes.Buffer
	n, err := doc.WriteTo(&buf)
	if err != nil {
		t.Fatalf("WriteTo() error = %v", err)
	}
	if !bytes.Equal(buf.Bytes(), want) {
		t.Errorf("WriteTo() wrote %q, want %q", buf.Bytes(), want)
	}
	if n != int64(len(want)) {
		t.Errorf("WriteTo() = %d bytes, want %d", n, len(want))
	}
}

func TestWriteFormattedToMatchesRenderFormatted(t *testing.T) {
	doc := sampleDocument()
	want, err := doc.RenderFormatted()
	if err != nil {
		t.Fatal(err)
	}

	var buf bytes.Buffer
	n, err := doc.WriteFormattedTo(&buf)
	if err != nil {
		t.Fatalf("WriteFormattedTo() error = %v", err)
	}
	if !bytes.Equal(buf.Bytes(), want) {
		t.Errorf("WriteFormattedTo() wrote %q, want %q", buf.Bytes(), want)
	}
	if n != int64(len(want)) {
		t.Errorf("WriteFormattedTo() = %d bytes, want %d", n, len(want))
	}
}

func TestWriteToReturnsDocumentErrorWithoutWriting(t *testing.T) {
	doc := NewHtmlFile()
	doc.AddToHead(NewP().Text("not a head element"))

	writers := map[string]func(io.Writer) (int64, error){
		"WriteTo":          doc.WriteTo,
		"WriteFormattedTo": doc.WriteFormattedTo,
	}
	for name, write := range writers {
		t.Run(name, func(t *testing.T) {
			var buf bytes.Buffer
			n, err := write(&buf)
			if err == nil {
				t.Fatal("expected the document error")
			}
			if n != 0 || buf.Len() != 0 {
				t.Errorf("wrote %d bytes (%q) despite the error", n, buf.Bytes())
			}
		})
	}
}

func TestWriteToPropagatesWriterErrors(t *testing.T) {
	boom := errors.New("boom")
	doc := sampleDocument()

	if _, err := doc.WriteTo(failingWriter{boom}); !errors.Is(err, boom) {
		t.Errorf("WriteTo() error = %v, want %v", err, boom)
	}
	if _, err := doc.WriteFormattedTo(failingWriter{boom}); !errors.Is(err, boom) {
		t.Errorf("WriteFormattedTo() error = %v, want %v", err, boom)
	}
}

func TestWriteFormattedToReportsShortWrite(t *testing.T) {
	if _, err := sampleDocument().WriteFormattedTo(shortWriter{}); !errors.Is(err, io.ErrShortWrite) {
		t.Errorf("WriteFormattedTo() error = %v, want io.ErrShortWrite", err)
	}
}

// TestConcurrentWriteTo streams one built document from many goroutines, the
// way an HTTP handler serving a prebuilt page would.
func TestConcurrentWriteTo(t *testing.T) {
	doc := sampleDocument()
	want, err := doc.Render()
	if err != nil {
		t.Fatal(err)
	}

	var wg sync.WaitGroup
	for range 64 {
		wg.Add(1)
		go func() {
			defer wg.Done()
			var buf bytes.Buffer
			if _, err := doc.WriteTo(&buf); err != nil {
				t.Errorf("WriteTo() error = %v", err)
				return
			}
			if !bytes.Equal(buf.Bytes(), want) {
				t.Errorf("WriteTo() wrote %q, want %q", buf.Bytes(), want)
			}
		}()
	}
	wg.Wait()
}

// BenchmarkRender measures rendering a document made mostly of text elements.
func BenchmarkRender(b *testing.B) {
	body := NewDiv()
	for range 50 {
		body.Add(NewP().Text("a < b & c")).
			Add(NewH2().Text("heading")).
			Add(NewSpan().AddStyle("color", "red").Text("one")).
			Add(NewStrong().Text("two")).
			Add(NewAnchor().Link("https://example.com").Text("link"))
	}
	doc := NewHtmlFile().AddToBody(body)
	b.ReportAllocs()
	for b.Loop() {
		if _, err := doc.Render(); err != nil {
			b.Fatal(err)
		}
	}
}
