package strftime

import (
	"testing"
	"time"
)

func TestCenturyNegativeYears(t *testing.T) {
	for _, tt := range []struct {
		year int
		want string
	}{
		{-101, "-2"}, {-100, "-1"}, {-99, "-1"}, {-1, "-1"},
		{0, "00"}, {99, "00"}, {100, "01"}, {9999, "99"}, {10000, "100"},
	} {
		date := time.Date(tt.year, 1, 1, 0, 0, 0, 0, time.UTC)
		if got := Format("%C", date); got != tt.want {
			t.Errorf("year %d: got %q, want %q", tt.year, got, tt.want)
		}
	}
}
