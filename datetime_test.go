package rephtml

import (
	"regexp"
	"strings"
	"testing"
)

// The package validates datetime values with hand-written scanners. These
// reference regexes are the original implementation, kept here so the scanners
// can be checked against it.
var (
	refWeekPattern     = regexp.MustCompile(`^\d{4}-W(0[1-9]|[1-4]\d|5[0-3])$`)
	refDurationPattern = regexp.MustCompile(`^P(?:\d+Y)?(?:\d+M)?(?:\d+W)?(?:\d+D)?(?:T(?:\d+H)?(?:\d+M)?(?:\d+(?:\.\d+)?S)?)?$`)
)

func refIsValidDuration(dt string) bool {
	if !refDurationPattern.MatchString(dt) {
		return false
	}
	return dt != "P" && dt != "PT" && !strings.HasSuffix(dt, "T")
}

func TestIsValidWeekTable(t *testing.T) {
	tests := []struct {
		in   string
		want bool
	}{
		{"2024-W01", true},
		{"2024-W09", true},
		{"2024-W10", true},
		{"2024-W49", true},
		{"2024-W50", true},
		{"2024-W53", true},
		{"2024-W00", false},
		{"2024-W54", false},
		{"2024-W60", false},
		{"2024-W1", false},
		{"2024-W011", false},
		{"024-W01", false},
		{"2024W01", false},
		{"2024-w01", false},
		{"٢٠٢٤-W01", false},
		{"", false},
	}
	for _, tt := range tests {
		t.Run(tt.in, func(t *testing.T) {
			if got := isValidWeek(tt.in); got != tt.want {
				t.Fatalf("isValidWeek(%q) = %v, want %v", tt.in, got, tt.want)
			}
		})
	}
}

func TestIsValidDurationTable(t *testing.T) {
	tests := []struct {
		in   string
		want bool
	}{
		{"P1Y", true},
		{"P1Y2M3W4DT5H6M7S", true},
		{"PT1.5S", true},
		{"PT36H", true},
		{"P3D", true},
		{"P", false},
		{"PT", false},
		{"P1DT", false},
		{"P1T", false},
		{"P1D1Y", false},
		{"PT1.5H", false},
		{"PT1.S", false},
		{"PT.5S", false},
		{"P1", false},
		{"1D", false},
		{"", false},
	}
	for _, tt := range tests {
		t.Run(tt.in, func(t *testing.T) {
			if got := isValidDuration(tt.in); got != tt.want {
				t.Fatalf("isValidDuration(%q) = %v, want %v", tt.in, got, tt.want)
			}
		})
	}
}

// TestDatetimeScannersMatchRegexes compares the scanners with the regexes they
// replaced over every sequence of up to five tokens.
func TestDatetimeScannersMatchRegexes(t *testing.T) {
	tokens := []string{"P", "T", "1", "12", "0", "5", "53", "54", "2024", ".", "Y", "M", "W", "D", "H", "S", "-", "W0", "٢"}
	var walk func(prefix string, depth int)
	walk = func(prefix string, depth int) {
		if got, want := isValidDuration(prefix), refIsValidDuration(prefix); got != want {
			t.Fatalf("isValidDuration(%q) = %v, regex says %v", prefix, got, want)
		}
		if got, want := isValidWeek(prefix), refWeekPattern.MatchString(prefix); got != want {
			t.Fatalf("isValidWeek(%q) = %v, regex says %v", prefix, got, want)
		}
		if depth == 0 {
			return
		}
		for _, tok := range tokens {
			walk(prefix+tok, depth-1)
		}
	}
	walk("", 5)
}
