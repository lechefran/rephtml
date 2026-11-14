# Changelog

All notable changes to this project are documented in this file.

The format is based on [Keep a Changelog](https://keepachangelog.com/en/1.1.0/),
and this project adheres to [Semantic Versioning](https://semver.org/spec/v2.0.0.html).
While the major version is 0, a minor release may contain breaking changes;
deprecated names are kept for at least one minor release before removal.

## [Unreleased]

### Changed

- `Time.Datetime` validates week and duration values with small hand-written
  scanners instead of regular expressions. The accepted values are unchanged,
  and the package no longer links `regexp`, which makes programs that use it
  about 250 KB smaller.

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

[0.2.0]: https://github.com/lechefran/rephtml/compare/v0.1.0-beta...v0.2.0
[0.1.0-beta]: https://github.com/lechefran/rephtml/releases/tag/v0.1.0-beta
