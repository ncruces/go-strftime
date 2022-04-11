package strftime

// https://strftime.org/
func goLayout(spec, pad byte) string {
	if pad == 0 {
		switch spec {
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
		case 'j':
			return ""
		case 'H':
			return ""
		}
	}

	switch spec {
	default:
		return ""

	case 'B':
		return "January"
	case 'b', 'h':
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
	case 'I':
		return "03"
	case 'H':
		return "15"
	case 'M':
		return "04"
	case 'S':
		return "05"
	case 'y':
		return "06"
	case 'Y':
		return "2006"
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

// https://nsdateformatter.com/
func uts35Pattern(spec, pad byte) string {
	if pad == 0 {
		switch spec {
		case 'm':
			return "M"
		case 'd':
			return "d"
		case 'j':
			return "D"
		case 'I':
			return "h"
		case 'H':
			return "H"
		case 'M':
			return "m"
		case 'S':
			return "s"
		}
	}

	switch spec {
	default:
		return ""

	case 'B':
		return "MMMM"
	case 'b', 'h':
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
	case 'I':
		return "hh"
	case 'H':
		return "HH"
	case 'M':
		return "mm"
	case 'S':
		return "ss"
	case 'y':
		return "yy"
	case 'Y':
		return "yyyy"
	case 'g':
		return "YY"
	case 'G':
		return "YYYY"
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
