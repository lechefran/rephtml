package rephtml

import (
	"go/ast"
	"go/parser"
	"go/token"
	"os"
	"strings"
	"testing"
)

// TestExportedIdentifiersDocumented keeps pkg.go.dev complete: every exported
// type, function, method, constant, variable, interface method and struct
// field must carry a doc comment that starts with its own name.
//
// Methods are checked even on unexported receivers (the generic element tiers
// in element.go), because pkg.go.dev shows them on every element that embeds
// the tier.
func TestExportedIdentifiersDocumented(t *testing.T) {
	entries, err := os.ReadDir(".")
	if err != nil {
		t.Fatalf("read package directory: %v", err)
	}

	fset := token.NewFileSet()
	var problems []string

	// report records a problem unless doc is a non-empty comment that starts
	// with name. label identifies the declaration in the failure message.
	report := func(pos token.Pos, label, name string, doc *ast.CommentGroup) {
		text := ""
		if doc != nil {
			text = doc.Text()
		}
		where := fset.Position(pos).String() + ": " + label
		switch {
		case strings.TrimSpace(text) == "":
			problems = append(problems, where+" has no doc comment")
		case !strings.HasPrefix(text, name+" ") && !strings.HasPrefix(text, name+"\n"):
			problems = append(problems, where+" doc should start with "+name)
		}
	}

	for _, e := range entries {
		name := e.Name()
		if e.IsDir() || !strings.HasSuffix(name, ".go") || strings.HasSuffix(name, "_test.go") {
			continue
		}
		file, err := parser.ParseFile(fset, name, nil, parser.ParseComments)
		if err != nil {
			t.Fatalf("parse %s: %v", name, err)
		}
		for _, decl := range file.Decls {
			switch d := decl.(type) {
			case *ast.FuncDecl:
				if d.Name.IsExported() {
					report(d.Pos(), "func "+d.Name.Name, d.Name.Name, d.Doc)
				}
			case *ast.GenDecl:
				checkGenDecl(d, report)
			}
		}
	}

	for _, p := range problems {
		t.Error(p)
	}
}

// checkGenDecl checks the type, const and var declarations of one GenDecl.
func checkGenDecl(d *ast.GenDecl, report func(pos token.Pos, label, name string, doc *ast.CommentGroup)) {
	for _, spec := range d.Specs {
		switch s := spec.(type) {
		case *ast.TypeSpec:
			if !s.Name.IsExported() {
				continue
			}
			doc := s.Doc
			if doc == nil && len(d.Specs) == 1 {
				doc = d.Doc
			}
			report(s.Pos(), "type "+s.Name.Name, s.Name.Name, doc)
			switch typ := s.Type.(type) {
			case *ast.StructType:
				for _, f := range typ.Fields.List {
					for _, n := range f.Names {
						if n.IsExported() {
							report(n.Pos(), "field "+s.Name.Name+"."+n.Name, n.Name, f.Doc)
						}
					}
				}
			case *ast.InterfaceType:
				for _, m := range typ.Methods.List {
					for _, n := range m.Names {
						if n.IsExported() {
							report(n.Pos(), "interface method "+s.Name.Name+"."+n.Name, n.Name, m.Doc)
						}
					}
				}
			}
		case *ast.ValueSpec:
			for _, n := range s.Names {
				if !n.IsExported() {
					continue
				}
				// A comment on the spec names the value; a comment on a
				// single-spec declaration does too.
				doc := s.Doc
				if doc == nil && len(d.Specs) == 1 {
					doc = d.Doc
				}
				report(n.Pos(), "value "+n.Name, n.Name, doc)
			}
		}
	}
}
