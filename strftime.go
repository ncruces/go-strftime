package strftime

import (
	"bytes"
	"errors"
	"strconv"
	"time"
	"unicode/utf8"
)

// Format returns a textual representation of the time value
// formatted according to the strftime format specification.
func Format(fmt string, t time.Time) string {
	buf := buffer(fmt)
	return string(AppendFormat(buf, fmt, t))
}

// AppendFormat is like Format, but appends the textual representation
// to dst and returns the extended buffer.
func AppendFormat(dst []byte, fmt string, t time.Time) []byte {
	var parser parser

	parser.literal = func(i int, lit byte) error {
		dst = append(dst, lit)
		return nil
	}

	parser.format = func(i int, spec, pad byte) error {
		if layout := goLayout(spec, pad); layout != "" {
			dst = t.AppendFormat(dst, layout)
			return nil
		}

		switch spec {
		default:
			dst = append(dst, '%')
			if pad == 0 {
				dst = append(dst, '-')
			}
			dst = append(dst, spec)
		case 'L':
			dst = append(dst, t.Format(".000")[1:]...)
		case 'f':
			dst = append(dst, t.Format(".000000")[1:]...)
		case 'N':
			dst = append(dst, t.Format(".000000000")[1:]...)
		case 'C':
			dst = t.AppendFormat(dst, "2006")
			dst = dst[:len(dst)-2]
		case 'g':
			y, _ := t.ISOWeek()
			dst = time.Date(y, 1, 1, 0, 0, 0, 0, time.UTC).AppendFormat(dst, "06")
		case 'G':
			y, _ := t.ISOWeek()
			dst = time.Date(y, 1, 1, 0, 0, 0, 0, time.UTC).AppendFormat(dst, "2006")
		case 'V':
			_, w := t.ISOWeek()
			dst = appendInt2(dst, w, pad)
		case 'W':
			dst = appendWeekNumber(dst, t, pad, false)
		case 'U':
			dst = appendWeekNumber(dst, t, pad, true)
		case 'w':
			w := t.Weekday()
			dst = appendInt1(dst, int(w))
		case 'u':
			if w := t.Weekday(); w == 0 {
				dst = append(dst, '7')
			} else {
				dst = appendInt1(dst, int(w))
			}
		case 'H':
			dst = appendInt2(dst, t.Hour(), pad)
		case 'k':
			dst = appendInt2(dst, t.Hour(), ' ')
		case 'l':
			h := t.Hour()
			if h == 0 {
				h = 12
			} else if h > 12 {
				h -= 12
			}
			dst = appendInt2(dst, h, ' ')
		case 'j':
			dst = strconv.AppendInt(dst, int64(t.YearDay()), 10)
		case 's':
			dst = strconv.AppendInt(dst, t.Unix(), 10)
		case 'Q':
			dst = strconv.AppendInt(dst, t.UnixMilli(), 10)
		}
		return nil
	}

	parser.parse(fmt)
	return dst
}

// Parse converts a textual representation of time to the time value it represents
// according to the strftime format specification.
func Parse(fmt, value string) (time.Time, error) {
	pattern, err := Layout(fmt)
	if err != nil {
		return time.Time{}, err
	}
	return time.Parse(pattern, value)
}

// Layout converts a strftime format specification
// to a Go time pattern specification.
func Layout(fmt string) (string, error) {
	dst := buffer(fmt)
	var parser parser

	parser.literal = func(i int, lit byte) error {
		if '0' <= lit && lit <= '9' {
			return errors.New("strftime: unsupported literal: '" + string(lit) + "'")
		}
		dst = append(dst, lit)
		if lit == 'M' || lit == 'T' || lit == 'm' || lit == 'n' {
			switch {
			case bytes.HasSuffix(dst, []byte("Jan")):
				return errors.New("strftime: unsupported literal: 'Jan'")
			case bytes.HasSuffix(dst, []byte("Mon")):
				return errors.New("strftime: unsupported literal: 'Mon'")
			case bytes.HasSuffix(dst, []byte("MST")):
				return errors.New("strftime: unsupported literal: 'MST'")
			case bytes.HasSuffix(dst, []byte("PM")):
				return errors.New("strftime: unsupported literal: 'PM'")
			case bytes.HasSuffix(dst, []byte("pm")):
				return errors.New("strftime: unsupported literal: 'pm'")
			}
		}
		return nil
	}

	parser.format = func(i int, spec, pad byte) error {
		if layout := goLayout(spec, pad); layout != "" {
			dst = append(dst, layout...)
			return nil
		}

		switch spec {
		default:
			return errors.New("strftime: unsupported specifier: %" + runeSubstr(fmt, i))

		case 'L', 'f', 'N':
			if bytes.HasSuffix(dst, []byte(".")) || bytes.HasSuffix(dst, []byte(",")) {
				switch spec {
				default:
					dst = append(dst, "000"...)
				case 'f':
					dst = append(dst, "000000"...)
				case 'N':
					dst = append(dst, "000000000"...)
				}
				return nil
			}
			return errors.New("strftime: unsupported specifier: %" + string(spec) + " must follow '.' or ','")
		}
	}

	if err := parser.parse(fmt); err != nil {
		return "", err
	}
	return string(dst), nil
}

// UTS35 converts a strftime format specification
// to a Unicode Technical Standard #35 Date Format Pattern.
func UTS35(fmt string) (string, error) {
	const quote = '\''
	var quoted bool
	dst := buffer(fmt)

	var parser parser

	parser.literal = func(i int, lit byte) error {
		if lit == quote {
			dst = append(dst, quote, quote)
			return nil
		}
		if !quoted && ('a' <= lit && lit <= 'z' || 'A' <= lit && lit <= 'Z') {
			dst = append(dst, quote)
			quoted = true
		}
		dst = append(dst, lit)
		return nil
	}

	parser.format = func(i int, spec, pad byte) error {
		if quoted {
			dst = append(dst, quote)
			quoted = false
		}
		if pattern := uts35Pattern(spec, pad); pattern != "" {
			dst = append(dst, pattern...)
			return nil
		}
		return errors.New("strftime: unsupported specifier: %" + runeSubstr(fmt, i))
	}

	if err := parser.parse(fmt); err != nil {
		return "", err
	}
	if quoted {
		dst = append(dst, quote)
	}
	return string(dst), nil
}

func buffer(format string) (buf []byte) {
	const bufSize = 64
	max := len(format) + 10
	if max < bufSize {
		var b [bufSize]byte
		buf = b[:0]
	} else {
		buf = make([]byte, 0, max)
	}
	return
}

func appendWeekNumber(dst []byte, t time.Time, pad byte, sunday bool) []byte {
	offset := int(t.Weekday())
	if sunday {
		offset = 6 - offset
	} else if offset != 0 {
		offset = 7 - offset
	}
	return appendInt2(dst, (t.YearDay()+offset)/7, pad)
}

func appendInt1(dst []byte, i int) []byte {
	return append(dst, byte('0'+i))
}

func appendInt2(dst []byte, i int, pad byte) []byte {
	if pad == '0' || i >= 10 {
		return append(dst, smallsString[i*2:i*2+2]...)
	}
	if pad != 0 {
		dst = append(dst, pad)
	}
	return appendInt1(dst, i)
}

const smallsString = "00010203040506070809" +
	"10111213141516171819" +
	"20212223242526272829" +
	"30313233343536373839" +
	"40414243444546474849" +
	"50515253545556575859" +
	"60616263646566676869" +
	"70717273747576777879" +
	"80818283848586878889" +
	"90919293949596979899"

func runeSubstr(str string, i int) string {
	_, n := utf8.DecodeLastRuneInString(str[i:])
	return str[i : i+n]
}
