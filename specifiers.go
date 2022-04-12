package strftime

// https://strftime.org/
func goLayout(spec, flag byte) string {
	switch spec {
	default:
		return ""

	case 'B':
		return "January"
	case 'b', 'h':
		return "Jan"
	case 'm':
		if flag == '-' {
			return "1"
		}
		return "01"
	case 'A':
		return "Monday"
	case 'a':
		return "Mon"
	case 'e':
		return "_2"
	case 'd':
		if flag == '-' {
			return "2"
		}
		return "02"
	case 'j':
		if flag == '-' {
			return ""
		}
		return "002"
	case 'I':
		if flag == '-' {
			return "3"
		}
		return "03"
	case 'H':
		if flag == '-' {
			return ""
		}
		return "15"
	case 'M':
		if flag == '-' {
			return "4"
		}
		return "04"
	case 'S':
		if flag == '-' {
			return "5"
		}
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
		if flag == ':' {
			return "-07:00"
		}
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
func uts35Pattern(spec, flag byte) string {
	switch spec {
	default:
		return ""

	case 'B':
		return "MMMM"
	case 'b', 'h':
		return "MMM"
	case 'm':
		if flag == '-' {
			return "M"
		}
		return "MM"
	case 'A':
		return "EEEE"
	case 'a':
		return "E"
	case 'd':
		if flag == '-' {
			return "d"
		}
		return "dd"
	case 'j':
		if flag == '-' {
			return "D"
		}
		return "DDD"
	case 'I':
		if flag == '-' {
			return "h"
		}
		return "hh"
	case 'H':
		if flag == '-' {
			return "H"
		}
		return "HH"
	case 'M':
		if flag == '-' {
			return "m"
		}
		return "mm"
	case 'S':
		if flag == '-' {
			return "s"
		}
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
		if flag == '-' {
			return "w"
		}
		return "ww"
	case 'p':
		return "a"
	case 'Z':
		return "zzz"
	case 'z':
		if flag == ':' {
			return "xxx"
		}
		return "xx"
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
