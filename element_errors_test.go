package rephtml

import (
	"bytes"
	"errors"
	"path/filepath"
	"strings"
	"testing"
)

// The tests in this file cover errors recorded by elements anywhere in a
// document reaching HtmlFile.Err, which stops every output method.

const invalidDatetime = `invalid datetime value: "tomorrow"`

// assertDocumentFails checks that Err reports want and that no output method
// produces anything.
func assertDocumentFails(t *testing.T, doc *HtmlFile, want string) {
	t.Helper()
	err := doc.Err()
	if err == nil || !strings.Contains(err.Error(), want) {
		t.Fatalf("Err() = %v, want an error containing %q", err, want)
	}
	if out, err := doc.Render(); err == nil || out != nil {
		t.Fatalf("Render() = %q, %v; want no output and an error", out, err)
	}
	if out, err := doc.RenderFormatted(); err == nil || out != nil {
		t.Fatalf("RenderFormatted() = %q, %v; want no output and an error", out, err)
	}
	var buf bytes.Buffer
	if _, err := doc.WriteTo(&buf); err == nil || buf.Len() != 0 {
		t.Fatalf("WriteTo() wrote %q, %v; want no output and an error", buf.String(), err)
	}
	if err := doc.WriteToFile(filepath.Join(t.TempDir(), "out.html")); err == nil {
		t.Fatal("WriteToFile() error = nil, want an error")
	}
}

func TestElementErrorsReachDocument(t *testing.T) {
	badTime := func() *Time { return NewTime().Datetime("tomorrow").Text("soon") }

	tests := []struct {
		name string
		doc  func() *HtmlFile
		want string
	}{
		{
			name: "Time in the body",
			doc:  func() *HtmlFile { return NewHtmlFile().AddToBody(NewP().Text("x")).AddToBody(badTime()) },
			want: invalidDatetime,
		},
		{
			name: "Time nested in containers",
			doc: func() *HtmlFile {
				return NewHtmlFile().AddToBody(NewDiv().Add(NewUl().Add(NewLi().Add(badTime()))))
			},
			want: invalidDatetime,
		},
		{
			name: "Time in a table cell",
			doc: func() *HtmlFile {
				table := NewTable().AddTbody(NewTbody().AddTr(NewTr().AddTd(NewTd().Add(badTime()))))
				return NewHtmlFile().AddToBody(table)
			},
			want: invalidDatetime,
		},
		{
			name: "Time in a row added directly to a table",
			doc: func() *HtmlFile {
				return NewHtmlFile().AddToBody(NewTable().AddTr(NewTr().AddTd(NewTd().Add(badTime()))))
			},
			want: invalidDatetime,
		},
		{
			name: "Time in a merged Body wrapper",
			doc:  func() *HtmlFile { return NewHtmlFile().AddToBody(NewBody().Add(badTime())) },
			want: invalidDatetime,
		},
		{
			name: "Time added to a Body wrapper after merging",
			doc: func() *HtmlFile {
				body := NewBody()
				doc := NewHtmlFile().AddToBody(body)
				body.Add(badTime())
				return doc
			},
			want: invalidDatetime,
		},
		{
			name: "Time in html content",
			doc:  func() *HtmlFile { return NewHtmlFile().Add(NewDiv().Add(badTime())) },
			want: invalidDatetime,
		},
		{
			name: "StyleElement given an element",
			doc:  func() *HtmlFile { return NewHtmlFile().AddToHead(NewStyleElement().Add(NewDiv())) },
			want: "cannot add Div to style element",
		},
		{
			name: "Body wrapper nested in content",
			doc:  func() *HtmlFile { return NewHtmlFile().Add(NewBody().Add(NewTitle())) },
			want: "cannot add Title to body element",
		},
		{
			name: "Table mixing legacy rows and sections",
			doc: func() *HtmlFile {
				table := NewTable().AddRow([]string{"lost"}).AddTbody(NewTbody())
				return NewHtmlFile().AddToBody(NewDiv().Add(table))
			},
			want: "table headers and rows are ignored",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assertDocumentFails(t, tt.doc(), tt.want)
		})
	}
}

func TestElementErrorIsReadWhenTheDocumentIs(t *testing.T) {
	when := NewTime().Datetime("tomorrow")
	doc := NewHtmlFile().AddToBody(when)
	if doc.Err() == nil {
		t.Fatal("Err() = nil with an invalid datetime")
	}

	when.Datetime("2026-11-14")
	if err := doc.Err(); err != nil {
		t.Fatalf("Err() after correcting the datetime = %v, want nil", err)
	}
	want := `<!DOCTYPE html><html><body><time datetime="2026-11-14"></time></body></html>`
	if got, _ := doc.RenderString(); got != want {
		t.Fatalf("got %q, want %q", got, want)
	}
}

func TestDocumentErrorsAreAllReported(t *testing.T) {
	doc := NewHtmlFile().
		AddToHead(NewP()).
		AddToBody(NewTime().Datetime("tomorrow")).
		AddToBody(NewTable().Headers([]string{"h"}).AddTr(NewTr()))

	err := doc.Err()
	for _, want := range []string{"cannot add P to head element", invalidDatetime, "table headers and rows are ignored"} {
		if err == nil || !strings.Contains(err.Error(), want) {
			t.Errorf("Err() = %v, want it to contain %q", err, want)
		}
	}
}

func TestDocumentErrorIdentityIsKept(t *testing.T) {
	doc := NewHtmlFile().AddToHead(NewP())
	own := doc.err
	if own == nil {
		t.Fatal("expected a document error")
	}
	// With nothing else to report, Err returns the document's own error as is.
	if err := doc.Err(); err != own {
		t.Fatalf("Err() = %v, want the document's own error value", err)
	}

	doc.AddToBody(NewTime().Datetime("tomorrow"))
	if err := doc.Err(); !errors.Is(err, own) {
		t.Fatalf("Err() = %v, want it to wrap the document's own error", err)
	}
}

func TestValidDocumentHasNoError(t *testing.T) {
	doc := NewHtmlFile().
		AddToHead(NewStyleElement().Add(NewStyle())).
		AddToBody(NewBody().Add(NewTime().Datetime("2026-11-14"))).
		AddToBody(NewTable().AddCaption(NewCaption().Text("c")).AddRow([]string{"r"}))

	if err := doc.Err(); err != nil {
		t.Fatalf("Err() = %v, want nil", err)
	}
}

func TestTableErr(t *testing.T) {
	tests := []struct {
		name    string
		table   *Table
		wantErr bool
	}{
		{"empty", NewTable(), false},
		{"legacy only", NewTable().Headers([]string{"h"}).AddRow([]string{"r"}), false},
		{"sections only", NewTable().AddThead(NewThead()).AddTbody(NewTbody()), false},
		{"caption with legacy rows", NewTable().AddCaption(NewCaption()).AddRow([]string{"r"}), false},
		{"headers with thead", NewTable().Headers([]string{"h"}).AddThead(NewThead()), true},
		{"rows with tbody", NewTable().AddRow([]string{"r"}).AddTbody(NewTbody()), true},
		{"rows with tfoot", NewTable().AddRow([]string{"r"}).AddTfoot(NewTfoot()), true},
		{"rows with a direct tr", NewTable().AddRow([]string{"r"}).AddTr(NewTr()), true},
		{"rows with a nil tbody", NewTable().AddRow([]string{"r"}).AddTbody(nil), false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := tt.table.Err(); (err != nil) != tt.wantErr {
				t.Fatalf("Err() = %v, want error: %v", err, tt.wantErr)
			}
		})
	}
}
