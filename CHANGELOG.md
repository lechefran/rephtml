# Changelog

All notable changes to this project are documented in this file.

The format is based on [Keep a Changelog](https://keepachangelog.com/en/1.1.0/),
and this project adheres to [Semantic Versioning](https://semver.org/spec/v2.0.0.html).
While the major version is 0, a minor release may contain breaking changes;
deprecated names are kept for at least one minor release before removal.

## [Unreleased]

### Changed

- ⚠️ **POTENTIALLY BREAKING:** `Text` now replaces the text on every element.
  Before, it appended on the elements that also hold child elements (`Title`,
  `Label`, `Button`, `Option`, `Output`, `Textarea`, `Progress`, `Meter`,
  `Legend`) and on `StyleElement`. On those elements the child elements and
  CSS rules stay where they are: the new text takes the place of the first
  earlier text, or goes after the children when there was none.

  ```go
  rephtml.NewButton().Text("Save").Text(" draft")
  ```

  Before: `<button>Save draft</button>`. After: `<button> draft</button>`.

  *Who is affected:* code that calls `Text` more than once on one of these
  elements, or after `AddText`, to build up its text. *Migration:* use
  `AddText` for every call after the first.

- ⚠️ **POTENTIALLY BREAKING:** the setters for attributes that are always whole
  numbers take an `int` instead of a `string`: `Input.Size`, `Input.MaxLength`,
  `Input.MinLength`, `Select.Size`, `Textarea.Rows`, `Textarea.Cols`,
  `Textarea.MaxLength` and `Textarea.MinLength`. Values the attribute does not
  allow are omitted: `size`, `rows` and `cols` below one, and negative
  `maxlength` / `minlength`. Zero lengths are emitted.

  ```go
  rephtml.NewTextarea().Rows("4")  // before
  rephtml.NewTextarea().Rows(4)    // after
  ```

  The output for valid numbers is unchanged. Before, any text was written as
  given, so `Rows("abc")` produced `rows="abc"`; that no longer compiles.

  *Who is affected:* every caller of these setters, since the code no longer
  compiles. *Migration:* pass the number instead of its text, using
  `strconv.Atoi` where the value starts out as a string.

### Removed

- ⚠️ **POTENTIALLY BREAKING:** the deprecated names are removed. This reverses
  the 0.3.0 note that they would be kept until v1.0. Output from the
  replacement names is unchanged.

  | Removed | Use instead |
  |---|---|
  | `Form.Autocomplete`, `Input.Autocomplete` | `AutoComplete` |
  | `Form.Novalidate` | `NoValidate` |
  | `Input.Maxlength`, `Input.Minlength` | `MaxLength`, `MinLength` |
  | `Iframe.Allowfullscreen` | `AllowFullscreen` |
  | `Iframe.Referrerpolicy`, `Portal.Referrerpolicy`, `Script.Referrerpolicy` | `ReferrerPolicy` |
  | `Script.Crossorigin` | `CrossOrigin` |
  | `AddId` on the table elements | `Id` |
  | `Styles` on the table elements | `Style` |

  *Who is affected:* code that still uses an old name, which no longer
  compiles. The compiler error names the replacement. *Migration:* rename the
  call; `go vet` and staticcheck on v0.4.0 list every use.

## [0.4.0] - 2026-11-15

Additive only: no API is removed or renamed, and no existing output changes.

### Added

- `HtmlFile.Id`, `AddClass`, `AddClasses` and `Class` set the `id` and `class`
  attributes of the `<html>` element, written after its other attributes and
  before `style`. Output is unchanged for documents that don't set them.

## [0.3.0] - 2026-11-15

No API is removed or renamed. Four fixes below are marked ⚠️ **POTENTIALLY
BREAKING** because they change behaviour or output:

- element errors anywhere in a document now reach `HtmlFile.Err()`;
- `Head` and `Body` wrappers are read when the document renders;
- `Table` rows added with `AddTr` render before the `<tfoot>`;
- formatted output keeps inline content on one line.

To upgrade, check the document's error, finish building wrappers before adding
them if you relied on later changes being ignored, and update tests that
compare exact table or formatted output. Each entry has a before / after example
and its own migration note.

### Added

- Every element now supports the global `id` and `class` attributes, through
  `Id`, `AddClass`, `AddClasses` and `Class`. Before, only the table elements,
  `Input`, `Datalist` and `Template` had `Id`, and only the table elements had
  `class`. The attributes are written with the inline style, as `id`, `class`,
  then `style`. Output is unchanged for elements that don't set them, including
  the existing `id` position on `Input`, `Datalist` and `Template`.
- `Head` and `Body` wrappers keep their `id` and `class` when merged into an
  `HtmlFile`.
- `AddText` on every text element always appends, whether `Text` replaces or
  appends on that element. See the README for which does which.
- `Table.AddColgroup` adds a `<colgroup>`, rendered after the caption and
  before the thead. It appends, since a table can have several.
- `Table.Err` reports a table that mixes legacy `Headers` / `Rows` with
  structured sections.

### Fixed

- ⚠️ **POTENTIALLY BREAKING:** errors recorded by elements anywhere in a
  document now reach `HtmlFile.Err()`, so every output method returns them
  instead of rendering. Before, they were only visible on the element itself.
  This covers:
  - an invalid `Time.Datetime` value;
  - a `StyleElement` given content that is not CSS;
  - a `Head` or `Body` wrapper nested in other content instead of merged;
  - a `Table` that mixes legacy `Headers` / `Rows` with `thead`, `tbody`,
    `tfoot` or `tr` rows. Its legacy cells were silently not rendered; the new
    `Table.Err()` reports it.

  ```go
  doc := rephtml.NewHtmlFile().
      AddToBody(rephtml.NewTime().Datetime("tomorrow").Text("soon"))
  out, err := doc.RenderString()
  ```

  Before: `out` is `<!DOCTYPE html><html><body><time>soon</time></body></html>`,
  `err` is nil. After: `out` is empty, and `err` is `invalid datetime value:
  "tomorrow"`.

  The errors are read when `Err()` is called, so correcting the element, for
  example by setting a valid datetime, clears the document's error.

  *Who is affected:* documents that contain one of these mistakes and were
  rendering without them. *Migration:* check `HtmlFile.Err()` (or the output
  method's error), and fix the reported element. For a table, use either the
  legacy `Headers` / `Rows` or the structured sections, not both.

- ⚠️ **POTENTIALLY BREAKING:** `Head` and `Body` wrappers added with
  `AddToHead` / `AddToBody` are now read when the document renders, not copied
  when they are added. Before, anything added to a wrapper afterwards was
  silently lost.

  ```go
  body := rephtml.NewBody()
  doc := rephtml.NewHtmlFile().AddToBody(body)
  body.Add(rephtml.NewP().Text("late"))
  ```

  Before: `<!DOCTYPE html><html></html>`.
  After: `<!DOCTYPE html><html><body><p>late</p></body></html>`.

  This applies to content, `id`, `class`, style, `onload` / `onunload`, and
  errors. An invalid element added to a wrapper after merging now makes
  `HtmlFile.Err()` and every output method return an error, where before it was
  dropped without one.

  *Who is affected:* code that kept changing a wrapper after adding it to a
  document and relied on those changes being ignored. *Migration:* finish
  building the wrapper before adding it, or add a fresh wrapper.

- ⚠️ **POTENTIALLY BREAKING:** rows added directly to a `Table` with `AddTr`
  now render before the `<tfoot>`, as the HTML table content model requires,
  instead of after it.

  ```go
  rephtml.NewTable().AddTfoot(foot).AddTr(row)
  ```

  Before: `<table><tfoot>…</tfoot><tr>…</tr></table>`.
  After: `<table><tr>…</tr><tfoot>…</tfoot></table>`.

  Browsers already displayed the footer last, so pages look the same; the
  markup changes. *Who is affected:* code that compares or post-processes the
  exact table markup. *Migration:* update the expected output.

- ⚠️ **POTENTIALLY BREAKING:** formatted output (`RenderFormatted`,
  `RenderFormattedString`, `WriteFormattedTo`, `WriteToFile`) keeps each run of
  inline content on one line, exactly as rendered. Before, every tag went on its
  own line, which inserted a visible space between adjacent inline elements and
  before punctuation, and dropped whitespace-only text between them.

  ```html
  <div>Hello <b>x</b>.</div>
  ```

  Before:

  ```html
  <div>
  	Hello
  	<b>x</b>
  	.
  </div>
  ```

  which displays as `Hello x .`. After:

  ```html
  <div>
  	Hello <b>x</b>.
  </div>
  ```

  which displays as `Hello x.`, the same as the compact output. Block elements
  are still indented one per line, and whitespace between block elements is
  still dropped.

  *Who is affected:* code that compares formatted output byte for byte, and
  files written with `WriteToFile`. Compact `Render` output is unchanged.
  *Migration:* update the expected formatted output.

### Changed

- `Time.Datetime` validates week and duration values with small hand-written
  scanners instead of regular expressions. The accepted values are unchanged,
  and the package no longer links `regexp`, which makes programs that use it
  about 250 KB smaller.
- Deprecated names are now kept until v1.0, instead of for at least one minor
  release. This covers the names deprecated in 0.2.0, and `AddId` and `Styles`
  on the table elements.
- The README documents which setters take an `int` and which take a `string`,
  and lists the setters that take a `string` for a whole number. They keep
  their signatures until v1.0.

## [0.2.0] - 2026-11-14

First official release, following `v0.1.0-beta`.

### Security

- `Script.Text` now rewrites `</script` as `<\/script`, so script text can no
  longer close its element and inject markup. The rewrite means the same thing
  inside a JavaScript string or regular expression literal.
- `StyleElement.Text`, `RawCSSRule` and `ImportRule` conditions now apply the
  same `</style` protection as the structured CSS rules.
- `<meta http-equiv="refresh">` filters its redirect target by URL scheme, so a
  `javascript:` target is replaced with `BlockedURL`.
- `StyleElement.Add` no longer accepts arbitrary HTML elements into a raw-text
  `<style>` element. Other elements are dropped and the problem is reported by
  the new `StyleElement.Err`.
- Documented that `Iframe.Srcdoc` content runs as a document and should be paired
  with `Sandbox` when it is not trusted.

### Added

- `HtmlFile.WriteTo` (implements `io.WriterTo`) and `HtmlFile.WriteFormattedTo`
  stream a document to any `io.Writer`, such as an `http.ResponseWriter`,
  without the extra copy `Render` makes. A document error is returned before
  anything is written.
- `Bdo.Dir`, since `bdo` has no effect without a direction.
- `StyleElement.Err`.
- Consistently named setters: `Form.AutoComplete`, `Form.NoValidate`,
  `Input.AutoComplete`, `Input.MaxLength`, `Input.MinLength`,
  `Iframe.AllowFullscreen`, `Iframe.ReferrerPolicy`, `Portal.ReferrerPolicy`,
  `Script.CrossOrigin` and `Script.ReferrerPolicy`.
- Package documentation, runnable examples for pkg.go.dev, and a GitHub Actions
  workflow that runs gofmt, `go vet`, staticcheck and the tests with the race
  detector.

### Changed

- `HtmlFile.WriteToFile` writes atomically: it writes a temporary file in the
  destination directory, syncs it and renames it over the target, so a failed
  write never leaves a truncated file. Files are created with mode 0644 (before
  the umask) instead of 0666.
- `Ol.Start` and `Li.Value` emit zero and negative numbers instead of treating
  zero as unset.
- `Table.Headers`, `Rows`, `AddRow` and `AddRows` copy the slices they are
  given, so later changes by the caller no longer alter the table.
- Inline `style` attributes skip declarations with an empty property or value
  instead of writing `color: ;`, and are omitted when nothing remains.
- `Link` writes `crossorigin` and `referrerpolicy` in lower case, matching
  `Script`.
- Every doc comment that described the old "Prepare" step or was generic
  boilerplate now describes the element or attribute it documents.

### Deprecated

These keep working and forward to the new names. They will be removed in a
future minor release.

| Deprecated | Use instead |
|---|---|
| `Form.Autocomplete`, `Input.Autocomplete` | `AutoComplete` |
| `Form.Novalidate` | `NoValidate` |
| `Input.Maxlength`, `Input.Minlength` | `MaxLength`, `MinLength` |
| `Iframe.Allowfullscreen` | `AllowFullscreen` |
| `Iframe.Referrerpolicy`, `Portal.Referrerpolicy`, `Script.Referrerpolicy` | `ReferrerPolicy` |
| `Script.Crossorigin` | `CrossOrigin` |

### Removed

- IDE settings (`.idea/`) and the unused `cssprops.csv` are no longer tracked in
  the repository.

## [0.1.0-beta] - 2026-04-28

Initial pre-release.

[Unreleased]: https://github.com/lechefran/rephtml/compare/v0.4.0...HEAD
[0.4.0]: https://github.com/lechefran/rephtml/compare/v0.3.0...v0.4.0
[0.3.0]: https://github.com/lechefran/rephtml/compare/v0.2.0...v0.3.0
[0.2.0]: https://github.com/lechefran/rephtml/compare/v0.1.0-beta...v0.2.0
[0.1.0-beta]: https://github.com/lechefran/rephtml/releases/tag/v0.1.0-beta
