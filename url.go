package rephtml

import "strings"

// Attribute escaping keeps a URL from breaking out of its quotes, but it does
// nothing about what the URL then does. A javascript: URL is perfectly valid
// markup and runs code when followed, so URL-bearing attributes are filtered by
// scheme as well as escaped.

// BlockedURL replaces a URL whose scheme could execute script.
//
// A blocked URL is replaced rather than dropped so the problem is visible in
// the output instead of silently changing a link's meaning. Use IsSafeURL to
// detect the case up front and handle it yourself.
const BlockedURL = "#rephtml-blocked-url"

// scriptURLSchemes are the schemes a browser will execute as code.
//
// This is a blocklist rather than an allowlist on purpose. A static HTML
// generator has good reason to emit tel:, sms:, ftp:, blob: and application
// deep links such as zoommtg: or intent:, and an allowlist would break all of
// them. The set of schemes that actually execute script is small, well known,
// and has not grown in many years.
var scriptURLSchemes = map[string]bool{
	"javascript": true,
	"vbscript":   true,
	"livescript": true,
	"mocha":      true,
}

// IsSafeURL reports whether raw may be used as a URL attribute value.
//
// Relative URLs, fragments and queries are always safe. An absolute URL is safe
// unless its scheme executes script, or it is a data: URL whose media type a
// browser would treat as a document. Values that are not safe are replaced with
// BlockedURL when rendered.
func IsSafeURL(raw string) bool {
	normalized := normalizeURL(raw)
	switch scheme := urlScheme(normalized); scheme {
	case "":
		// No scheme: relative path, fragment, query, or protocol-relative.
		return true
	case "data":
		return isSafeDataURL(normalized)
	default:
		return !scriptURLSchemes[scheme]
	}
}

// safeURLValue returns raw, or BlockedURL when raw is not safe to emit.
func safeURLValue(raw string) string {
	if IsSafeURL(raw) {
		return raw
	}
	return BlockedURL
}

// ignoredURLBytes are the bytes dropped anywhere in a URL before its scheme is
// read. Browsers strip tab, LF and CR while parsing, and a NUL in an attribute
// value is replaced during HTML parsing, so none of them can be relied on to
// keep "java<TAB>script:" from meaning javascript:.
const ignoredURLBytes = "\t\n\r\x00"

// normalizeURL removes the characters a browser ignores when deciding what a
// URL means: leading C0 controls and spaces, plus ignoredURLBytes anywhere.
//
// Doing this before reading the scheme is what stops "java&#9;script:" and
// " JavaScript:" from being read as harmless. Only the copy used for the
// decision is normalized; the value emitted is the caller's original.
func normalizeURL(raw string) string {
	start := 0
	for start < len(raw) && (raw[start] <= ' ' || raw[start] == 0x7f) {
		start++
	}
	raw = raw[start:]

	if !strings.ContainsAny(raw, ignoredURLBytes) {
		return raw
	}

	var b strings.Builder
	b.Grow(len(raw))
	for i := 0; i < len(raw); i++ {
		if strings.IndexByte(ignoredURLBytes, raw[i]) >= 0 {
			continue
		}
		b.WriteByte(raw[i])
	}
	return b.String()
}

// urlScheme returns the lower-cased scheme of an already normalized URL, or ""
// when it has none.
func urlScheme(url string) string {
	for i := 0; i < len(url); i++ {
		c := url[i]
		switch {
		case c == ':':
			if i == 0 {
				return ""
			}
			return strings.ToLower(url[:i])
		case isSchemeByte(c, i == 0):
			// keep scanning
		default:
			// Anything else before a colon means this is not a scheme, so the
			// URL is relative.
			return ""
		}
	}
	return ""
}

// isSchemeByte reports whether c may appear in a URL scheme. A scheme has to
// start with an ASCII letter.
func isSchemeByte(c byte, first bool) bool {
	switch {
	case c >= 'a' && c <= 'z', c >= 'A' && c <= 'Z':
		return true
	case first:
		return false
	default:
		return c >= '0' && c <= '9' || c == '+' || c == '-' || c == '.'
	}
}

// isSafeDataURL reports whether a data: URL carries media a browser renders
// rather than a document it would parse.
//
// data:image/png is an ordinary inline image. data:text/html is a document in
// an inherited origin, and image/svg+xml can carry script once it reaches a
// context that runs it, so both are refused.
func isSafeDataURL(url string) bool {
	rest := url[len("data:"):]

	// The media type runs up to the first parameter or the payload separator.
	if i := strings.IndexAny(rest, ";,"); i >= 0 {
		rest = rest[:i]
	}
	mediaType := strings.ToLower(strings.TrimSpace(rest))

	switch {
	case mediaType == "image/svg+xml":
		return false
	case strings.HasPrefix(mediaType, "image/"),
		strings.HasPrefix(mediaType, "audio/"),
		strings.HasPrefix(mediaType, "video/"),
		strings.HasPrefix(mediaType, "font/"):
		return true
	default:
		return false
	}
}

// filterSrcset filters each candidate in a srcset attribute.
//
// A srcset is a comma-separated list of candidates, each a URL optionally
// followed by a width or density descriptor, so the URL has to be picked out of
// every entry rather than filtering the attribute as a whole.
func filterSrcset(value string) string {
	candidates := strings.Split(value, ",")
	changed := false

	for i, candidate := range candidates {
		trimmed := strings.TrimSpace(candidate)
		if trimmed == "" {
			continue
		}

		url, descriptor := trimmed, ""
		if cut := strings.IndexAny(trimmed, " \t\n\r\f"); cut >= 0 {
			url, descriptor = trimmed[:cut], trimmed[cut:]
		}

		safe := safeURLValue(url)
		if safe == url {
			continue
		}
		candidates[i] = safe + descriptor
		changed = true
	}

	if !changed {
		return value
	}
	return strings.Join(candidates, ", ")
}
