package strftime_test

import (
	"fmt"
	"log"
	"net/http"
	"testing"
	"time"

	"github.com/ncruces/go-strftime"
)

var reference = time.Date(2009, 8, 7, 6, 5, 4, 300000000, time.UTC)

var timeTests = []struct {
	format string
	layout string
	uts35  string
	time   string
}{
	// Date and time formats
	{"%c", time.ANSIC, "E MMM d HH:mm:ss yyyy", "Fri Aug  7 06:05:04 2009"},
	{"%+", time.UnixDate, "E MMM d HH:mm:ss zzz yyyy", "Fri Aug  7 06:05:04 UTC 2009"},
	{"%FT%TZ", time.RFC3339[:20], "yyyy-MM-dd'T'HH:mm:ss'Z'", "2009-08-07T06:05:04Z"},
	{"%a %b %e %T %Y", time.ANSIC, "", "Fri Aug  7 06:05:04 2009"},
	{"%a %b %e %T %Z %Y", time.UnixDate, "", "Fri Aug  7 06:05:04 UTC 2009"},
	{"%a %b %d %T %z %Y", time.RubyDate, "E MMM dd HH:mm:ss Z yyyy", "Fri Aug 07 06:05:04 +0000 2009"},
	{"%a %h %d %T %z %Y", time.RubyDate, "E MMM dd HH:mm:ss Z yyyy", "Fri Aug 07 06:05:04 +0000 2009"},
	{"%a, %d %b %Y %T %Z", time.RFC1123, "E, dd MMM yyyy HH:mm:ss zzz", "Fri, 07 Aug 2009 06:05:04 UTC"},
	{"%a, %d %b %Y %T GMT", http.TimeFormat, "E, dd MMM yyyy HH:mm:ss 'GMT'", "Fri, 07 Aug 2009 06:05:04 GMT"},
	{"%Y-%m-%dT%H:%M:%SZ", time.RFC3339[:20], "yyyy-MM-dd'T'HH:mm:ss'Z'", "2009-08-07T06:05:04Z"},
	// Date formats
	{"%v", "_2-Jan-2006", "d-MMM-yyyy", " 7-Aug-2009"},
	{"%F", "2006-01-02", "yyyy-MM-dd", "2009-08-07"},
	{"%D", "01/02/06", "MM/dd/yy", "08/07/09"},
	{"%x", "01/02/06", "MM/dd/yy", "08/07/09"},
	{"%e-%b-%Y", "_2-Jan-2006", "", " 7-Aug-2009"},
	{"%Y-%m-%d", "2006-01-02", "yyyy-MM-dd", "2009-08-07"},
	{"%m/%d/%y", "01/02/06", "MM/dd/yy", "08/07/09"},
	// Time formats
	{"%R", "15:04", "HH:mm", "06:05"},
	{"%T", "15:04:05", "HH:mm:ss", "06:05:04"},
	{"%X", "15:04:05", "HH:mm:ss", "06:05:04"},
	{"%r", "03:04:05 PM", "hh:mm:ss a", "06:05:04 AM"},
	{"%H:%M", "15:04", "HH:mm", "06:05"},
	{"%H:%M:%S", "15:04:05", "HH:mm:ss", "06:05:04"},
	{"%I:%M:%S %p", "03:04:05 PM", "hh:mm:ss a", "06:05:04 AM"},
	// Misc
	{"%g", "", "YY", "09"},
	{"%V/%G", "", "ww/YYYY", "32/2009"},
	{"%Cth Century Fox", "", "", "20th Century Fox"},
	{"%-d-%-m-%Y", "2-1-2006", "d-M-yyyy", "7-8-2009"},
	{"%-Hh%-Mm%-Ss", "15h4m5s", "H'h'm'm's's'", "06h5m4s"}, // FIXME: Format zero pads %-H
	{"%-I o'clock", "3 o'clock", "h 'o''clock'", "6 o'clock"},
	{"%-M past %-I %p", "4 past 3 PM", "m 'past 'h a", "5 past 6 AM"},
	{"%-A, the %uth day of the week", "", "", "Friday, the 5th day of the week"},
	{"%fμs since %T", "", "SSSSSSμ's since 'HH:mm:ss", "300000μs since 06:05:04"},
	{"%Nns since %T", "", "SSSSSSSSS'ns since 'HH:mm:ss", "300000000ns since 06:05:04"},
	{"%-S.%Ls since %R", "5.000s since 15:04", "s.SSS's since 'HH:mm", "4.300s since 06:05"},
	{"zero padded %-j is %j", "zero padded 002 is 002", "'zero padded 'D 'is 'DDD", "zero padded 219 is 219"},
	// Parsing
	{"", "", "", ""},
	{"%", "%", "%", "%"},
	{"%%", "%", "%", "%"},
	{"%-", "%-", "%-", "%-"},
	{"%n", "\n", "\n", "\n"},
	{"%t", "\t", "\t", "\t"},
	{"%q", "", "", "%q"},
	{"%-q", "", "", "%-q"},
	{"'", "'", "''", "'"},
	{"100%", "", "100%", "100%"},
	{"Monday", "", "'Monday'", "Monday"},
	{"January", "", "'January'", "January"},
	{"MST", "", "'MST'", "MST"},
	{"AM", "AM", "'AM'", "AM"},
	{"am", "am", "'am'", "am"},
	{"PM", "", "'PM'", "PM"},
	{"pm", "", "'pm'", "pm"},
}

func TestFormat(t *testing.T) {
	for _, test := range timeTests {
		if got := strftime.Format(test.format, reference); got != test.time {
			t.Errorf("Format(%q) = %q, want %q", test.format, got, test.time)
		}
	}
}

func TestFormat_Unix(t *testing.T) {
	tm := time.Unix(123456, 789*int64(time.Millisecond))

	if got := strftime.Format("%s", tm); got != "123456" {
		t.Errorf("Format(%q) = %q, want %q", "%s", got, "123456")
	}

	if got := strftime.Format("%Q", tm); got != "123456789" {
		t.Errorf("Format(%q) = %q, want %q", "%s", got, "123456789")
	}
}

func TestLayout(t *testing.T) {
	for _, test := range timeTests {
		if got, err := strftime.Layout(test.format); err != nil && test.layout != "" {
			t.Errorf("Layout(%q) = %v", test.format, err)
		} else if got != test.layout {
			t.Errorf("Layout(%q) = %q, want %q", test.format, got, test.layout)
		}
	}
}

func TestUTS35(t *testing.T) {
	for _, test := range timeTests {
		if got, err := strftime.UTS35(test.format); err != nil && test.uts35 != "" {
			t.Errorf("UTS35(%q) = %v", test.format, err)
		} else if got != test.uts35 {
			t.Errorf("UTS35(%q) = %q, want %q", test.format, got, test.uts35)
		}
	}
}

func ExampleLayout() {
	layout, err := strftime.Layout("%Y-%m-%d %H:%M:%S")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("%q", layout)
	// Output:
	// "2006-01-02 15:04:05"
}

func ExampleUTS35() {
	layout, err := strftime.UTS35("%Y-%m-%d %H:%M:%S")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("%q", layout)
	// Output:
	// "yyyy-MM-dd HH:mm:ss"
}
