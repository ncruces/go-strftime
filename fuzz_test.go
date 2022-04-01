//go:build go1.18

package strftime_test

import (
	"testing"

	"github.com/ncruces/go-strftime"
)

func FuzzFormat(f *testing.F) {
	for _, test := range timeTests {
		f.Add(test.format)
	}

	f.Fuzz(func(t *testing.T, fmt string) {
		s := strftime.Format(fmt, reference)
		if s == "" && fmt != "" {
			t.Errorf("Format(%q) = %q", fmt, s)
		}
	})
}

func FuzzParse(f *testing.F) {
	for _, test := range timeTests {
		f.Add(test.format, strftime.Format(test.format, reference))
	}

	f.Fuzz(func(t *testing.T, format, value string) {
		tm, err := strftime.Parse(format, value)
		if tm.IsZero() && err == nil {
			t.Errorf("Parse(%q, %q) = (%v, %v)", format, value, tm, err)
		}
	})
}

func FuzzLayout(f *testing.F) {
	for _, test := range timeTests {
		f.Add(test.format)
	}

	f.Fuzz(func(t *testing.T, format string) {
		layout, err := strftime.Layout(format)
		if format != "" && layout == "" && err == nil {
			t.Errorf("Layout(%q) = (%q, %v)", format, layout, err)
		}
	})
}

func FuzzUTS35(f *testing.F) {
	for _, test := range timeTests {
		f.Add(test.format)
	}

	f.Fuzz(func(t *testing.T, format string) {
		pattern, err := strftime.UTS35(format)
		if format != "" && pattern == "" && err == nil {
			t.Errorf("UTS35(%q) = (%q, %v)", format, pattern, err)
		}
	})
}
