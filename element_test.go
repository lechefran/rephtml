package rephtml

import "testing"

// The generic bases in element.go carry the type of the concrete element, which
// is what keeps fluent setters returning that concrete type. These tests fail to
// compile rather than fail at run time if that property is lost, which is the
// point: chaining is the package's whole calling convention.

func TestFluentSettersReturnConcreteTypes(t *testing.T) {
	// Each assignment is the real check: it only compiles if the promoted
	// setter returned the element's own type rather than the base's.
	var div *Div = NewDiv().AddStyle("color", "red").AddStyles(StyleMap{"margin": "0"}).Add(NewP())
	var anchor *Anchor = NewAnchor().Style(StyleMap{"color": "red"}).Text("go").Link("/x")
	var td *Td = NewTd().AddClass("c").Id("i").Style(StyleMap{"color": "red"}).Colspan(2)

	if got, want := div.HTML(), `<div style="color: red; margin: 0;"><p></p></div>`; got != want {
		t.Errorf("Div HTML() = %q, want %q", got, want)
	}
	if got, want := anchor.HTML(), `<a style="color: red;" href="/x">go</a>`; got != want {
		t.Errorf("Anchor HTML() = %q, want %q", got, want)
	}
	if got, want := td.HTML(), `<td id="i" class="c" style="color: red;" colspan="2"></td>`; got != want {
		t.Errorf("Td HTML() = %q, want %q", got, want)
	}
}

// TestEveryStyledElementHasFullStyleAPI guards the inconsistency the shared base
// removed: Div used to have only AddStyles, and Anchor and H1-H6 only AddStyle
// and Style. All three setters now come from one place, so every element has all
// of them.
func TestEveryStyledElementHasFullStyleAPI(t *testing.T) {
	tests := []struct {
		name string
		got  string
	}{
		{"Div.AddStyle", NewDiv().AddStyle("color", "red").HTML()},
		{"Div.Style", NewDiv().Style(StyleMap{"color": "red"}).HTML()},
		{"Anchor.AddStyles", NewAnchor().AddStyles(StyleMap{"color": "red"}).HTML()},
		{"H1.AddStyles", NewH1().AddStyles(StyleMap{"color": "red"}).HTML()},
		{"H6.AddStyles", NewH6().AddStyles(StyleMap{"color": "red"}).HTML()},
	}
	for _, tt := range tests {
		if tt.got == "" {
			t.Errorf("%s produced no output", tt.name)
		}
	}

	// The table elements historically spelled this Styles; Style now comes from
	// the shared base and Styles remains as a deprecated alias.
	style := StyleMap{"color": "red"}
	if a, b := NewTd().Style(style).HTML(), NewTd().Styles(style).HTML(); a != b {
		t.Errorf("Td.Style = %q but Td.Styles = %q, want both the same", a, b)
	}
	if a, b := NewTable().Id("x").HTML(), NewTable().AddId("x").HTML(); a != b {
		t.Errorf("Table.Id = %q but Table.AddId = %q, want both the same", a, b)
	}
}

// TestStyleSettersCopyCallerMaps pins that Style takes a copy, so a caller can
// reuse or mutate the map it passed in.
func TestStyleSettersCopyCallerMaps(t *testing.T) {
	shared := StyleMap{"color": "red"}
	div := NewDiv().Style(shared)
	shared["color"] = "blue"

	if got, want := div.HTML(), `<div style="color: red;"></div>`; got != want {
		t.Errorf("Style did not copy the caller's map: got %q, want %q", got, want)
	}
}

// TestZeroValueElementReportsMissingConstructor pins the failure mode for an
// element built as a bare composite literal. It cannot render, because nothing
// bound it to its base, and a named panic beats a nil dereference from inside
// the render path.
func TestZeroValueElementReportsMissingConstructor(t *testing.T) {
	defer func() {
		r := recover()
		if r == nil {
			t.Fatal("expected a panic from an unconstructed element")
		}
		want := "rephtml: element was not created with its New* constructor"
		if got, ok := r.(string); !ok || got != want {
			t.Fatalf("panicked with %v, want %q", r, want)
		}
	}()

	(&Div{}).HTML()
}

// TestElementsSatisfyInterfaces confirms the public interfaces are still
// satisfied now that Render and HTML are promoted from the base rather than
// declared per element.
func TestElementsSatisfyInterfaces(t *testing.T) {
	var (
		_ Element     = NewDiv()
		_ Element     = NewAnchor()
		_ BodyElement = NewDiv()
		_ HeadElement = NewTitle()
		// Script and Noscript are valid in both the head and the body, which is
		// why the markers are separate mixins rather than part of the tier chain.
		_ HeadElement = NewScript()
		_ BodyElement = NewScript()
		_ HeadElement = NewNoscript()
		_ BodyElement = NewNoscript()
	)
}
