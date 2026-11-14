package rephtml

import (
	"os"
	"path/filepath"
	"runtime"
	"testing"
)

// dirEntries lists the names in dir, so a test can assert that nothing besides
// the expected files is left there.
func dirEntries(t *testing.T, dir string) []string {
	t.Helper()
	entries, err := os.ReadDir(dir)
	if err != nil {
		t.Fatalf("ReadDir(%q): %v", dir, err)
	}
	names := make([]string, 0, len(entries))
	for _, e := range entries {
		names = append(names, e.Name())
	}
	return names
}

func TestWriteToFileReplacesExistingFile(t *testing.T) {
	dir := t.TempDir()
	path := filepath.Join(dir, "report.html")
	if err := os.WriteFile(path, []byte("old content that is much longer than the new document body"), 0o600); err != nil {
		t.Fatal(err)
	}

	doc := NewHtmlFile().AddToBody(NewP().Text("new"))
	if err := doc.WriteToFile(path); err != nil {
		t.Fatalf("WriteToFile() error = %v", err)
	}

	got, err := os.ReadFile(path)
	if err != nil {
		t.Fatal(err)
	}
	want, err := doc.RenderFormatted()
	if err != nil {
		t.Fatal(err)
	}
	if string(got) != string(want) {
		t.Errorf("file = %q, want %q", got, want)
	}
	if names := dirEntries(t, dir); len(names) != 1 || names[0] != "report.html" {
		t.Errorf("directory entries = %v, want only report.html", names)
	}
}

func TestWriteToFileLeavesNoTemporaryFileOnError(t *testing.T) {
	dir := t.TempDir()

	// Writing into a missing subdirectory fails before any file exists.
	doc := NewHtmlFile().AddToBody(NewP().Text("x"))
	if err := doc.WriteToFile(filepath.Join(dir, "missing", "report.html")); err == nil {
		t.Fatal("expected an error for a missing directory")
	}

	// A destination that cannot be replaced (a non-empty directory) fails at the
	// rename, after the temporary file was written, so cleanup has to run.
	blocked := filepath.Join(dir, "blocked")
	if err := os.MkdirAll(filepath.Join(blocked, "child"), 0o755); err != nil {
		t.Fatal(err)
	}
	if err := doc.WriteToFile(blocked); err == nil {
		t.Fatal("expected an error when the destination is a directory")
	}

	names := dirEntries(t, dir)
	if len(names) != 1 || names[0] != "blocked" {
		t.Errorf("directory entries = %v, want only the blocked directory", names)
	}
}

func TestWriteToFileKeepsExistingFileWhenDocumentHasErrors(t *testing.T) {
	path := filepath.Join(t.TempDir(), "report.html")
	if err := os.WriteFile(path, []byte("keep me"), 0o644); err != nil {
		t.Fatal(err)
	}

	doc := NewHtmlFile()
	doc.AddToHead(NewP().Text("not a head element"))
	if err := doc.WriteToFile(path); err == nil {
		t.Fatal("expected a document error")
	}

	got, err := os.ReadFile(path)
	if err != nil {
		t.Fatal(err)
	}
	if string(got) != "keep me" {
		t.Errorf("file = %q, want it untouched", got)
	}
}

func TestWriteToFileMode(t *testing.T) {
	if runtime.GOOS == "windows" {
		t.Skip("file modes are not meaningful on Windows")
	}

	path := filepath.Join(t.TempDir(), "report.html")
	if err := NewHtmlFile().AddToBody(NewP().Text("x")).WriteToFile(path); err != nil {
		t.Fatal(err)
	}
	info, err := os.Stat(path)
	if err != nil {
		t.Fatal(err)
	}
	if got := info.Mode().Perm(); got != 0o644 {
		t.Errorf("mode = %o, want 644", got)
	}
}
