package strftime

import (
	"strconv"
	"time"
)

// https://strftime.org/
func goLayout(s byte) string {
	switch s {
	default:
		return ""

	case 'B':
		return "January"
	case 'b':
		return "Jan"
	case 'h':
		return "Jan"
	case 'm':
		return "01"
	case 'A':
		return "Monday"
	case 'a':
		return "Mon"
	case 'e':
		return "_2"
	case 'd':
		return "02"
	case 'j':
		return "002"
	case 'H':
		return "15"
	case 'I':
		return "03"
	case 'M':
		return "04"
	case 'S':
		return "05"
	case 'Y':
		return "2006"
	case 'y':
		return "06"
	case 'p':
		return "PM"
	case 'P':
		return "pm"
	case 'Z':
		return "MST"
	case 'z':
		return "-0700"

	case '+':
		return "Mon Jan _2 15:04:05 MST 2006"
	case 'c':
		return "Mon Jan _2 15:04:05 2006"
	case 'v':
		return "_2-Jan-2006"
	case 'F':
		return "2006-01-02"
	case 'D':
		return "01/02/06"
	case 'x':
		return "01/02/06"
	case 'r':
		return "03:04:05 PM"
	case 'T':
		return "15:04:05"
	case 'X':
		return "15:04:05"
	case 'R':
		return "15:04"

	case '%':
		return "%"
	case 't':
		return "\t"
	case 'n':
		return "\n"
	}
}

// https://strftime.org/
func goLayoutUnpadded(s byte) string {
	switch s {
	default:
		return ""
	case 'm':
		return "1"
	case 'd':
		return "2"
	case 'I':
		return "3"
	case 'M':
		return "4"
	case 'S':
		return "5"
	}
}

// https://nsdateformatter.com/
func uts35Pattern(s byte) string {
	switch s {
	default:
		return ""

	case 'B':
		return "MMMM"
	case 'b':
		return "MMM"
	case 'h':
		return "MMM"
	case 'm':
		return "MM"
	case 'A':
		return "EEEE"
	case 'a':
		return "E"
	case 'd':
		return "dd"
	case 'j':
		return "DDD"
	case 'H':
		return "HH"
	case 'I':
		return "hh"
	case 'M':
		return "mm"
	case 'S':
		return "ss"
	case 'Y':
		return "yyyy"
	case 'y':
		return "yy"
	case 'G':
		return "YYYY"
	case 'g':
		return "YY"
	case 'V':
		return "ww"
	case 'p':
		return "a"
	case 'Z':
		return "zzz"
	case 'z':
		return "Z"
	case 'L':
		return "SSS"
	case 'f':
		return "SSSSSS"
	case 'N':
		return "SSSSSSSSS"

	case '+':
		return "E MMM d HH:mm:ss zzz yyyy"
	case 'c':
		return "E MMM d HH:mm:ss yyyy"
	case 'v':
		return "d-MMM-yyyy"
	case 'F':
		return "yyyy-MM-dd"
	case 'D':
		return "MM/dd/yy"
	case 'x':
		return "MM/dd/yy"
	case 'r':
		return "hh:mm:ss a"
	case 'T':
		return "HH:mm:ss"
	case 'X':
		return "HH:mm:ss"
	case 'R':
		return "HH:mm"

	case '%':
		return "%"
	case 't':
		return "\t"
	case 'n':
		return "\n"
	}
}

func uts35PatternUnpadded(s byte) string {
	switch s {
	default:
		return ""
	case 'm':
		return "M"
	case 'd':
		return "d"
	case 'j':
		return "D"
	case 'H':
		return "H"
	case 'I':
		return "h"
	case 'M':
		return "m"
	case 'S':
		return "s"
	}
}

func weekNumber(t time.Time, pad, monday bool) string {
	day := t.YearDay()
	offset := int(t.Weekday())
	if monday {
		if offset == 0 {
			offset = 6
		} else {
			offset--
		}
	}

	if day < offset {
		if pad {
			return "00"
		} else {
			return "0"
		}
	}

	n := (day-offset)/7 + 1
	if n < 10 && pad {
		return "0" + strconv.Itoa(n)
	}
	return strconv.Itoa(n)
}
