// Package strftime provides strftime/strptime compatible time formatting and parsing.
package strftime

import (
	"bytes"
	"errors"
	"strconv"
	"time"
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
	parser.fmt = fmt
	parser.basic = goLayout
	parser.unpadded = goLayoutUnpadded

	parser.writeLit = func(b byte) error {
		dst = append(dst, b)
		return nil
	}

	parser.writeFmt = func(fmt string) error {
		dst = t.AppendFormat(dst, fmt)
		return nil
	}

	parser.fallback = func(spec byte, pad bool) error {
		switch spec {
		default:
			dst = append(dst, '%')
			if !pad {
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
			if w < 10 && pad {
				dst = append(dst, '0')
			}
			dst = strconv.AppendInt(dst, int64(w), 10)
		case 'W':
			dst = append(dst, weekNumber(t, pad, true)...)
		case 'U':
			dst = append(dst, weekNumber(t, pad, false)...)
		case 'w':
			dst = strconv.AppendInt(dst, int64(t.Weekday()), 10)
		case 'u':
			if w := byte(t.Weekday()); w == 0 {
				dst = append(dst, '7')
			} else {
				dst = append(dst, '0'+w)
			}
		case 'k':
			h := t.Hour()
			if h < 10 {
				dst = append(dst, ' ')
			}
			dst = strconv.AppendInt(dst, int64(h), 10)
		case 'l':
			h := t.Hour()
			if h == 0 {
				h = 12
			} else if h > 12 {
				h -= 12
			}
			if h < 10 {
				dst = append(dst, ' ')
			}
			dst = strconv.AppendInt(dst, int64(h), 10)
		case 's':
			dst = strconv.AppendInt(dst, t.Unix(), 10)
		}
		return nil
	}

	parser.parse()
	return dst
}

// Parse converts a textual representation of time to the time value it represents
// according to the strftime format specification.
func Parse(fmt, value string) (time.Time, error) {
	layout, err := Layout(fmt)
	if err != nil {
		return time.Time{}, err
	}
	return time.Parse(layout, value)
}

// Layout converts a strftime format specification
// to a Go time layout specification.
func Layout(fmt string) (string, error) {
	var parser parser
	parser.fmt = fmt
	parser.basic = goLayout
	parser.unpadded = goLayoutUnpadded

	dst := buffer(fmt)

	parser.writeLit = func(b byte) error {
		if '0' <= b && b <= '9' {
			return errors.New("strftime: unsupported literal digit: '" + string(b) + "'")
		}
		dst = append(dst, b)
		if b == 'M' || b == 'T' || b == 'm' || b == 'n' {
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

	parser.writeFmt = func(fmt string) error {
		dst = append(dst, fmt...)
		return nil
	}

	parser.fallback = func(spec byte, pad bool) error {
		switch spec {
		case 'L', 'f', 'N':
			if bytes.HasSuffix(dst, []byte(".")) || bytes.HasSuffix(dst, []byte(",")) {
				switch spec {
				case 'L':
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

		return errors.New("strftime: unsupported specifier: %" + string(spec))
	}

	if err := parser.parse(); err != nil {
		return "", err
	}

	parser.writeFmt("")
	return string(dst), nil
}

// UTS35 converts a strftime format specification
// to a Unicode Technical Standard #35 Date Format Pattern.
func UTS35(fmt string) (string, error) {
	var parser parser
	parser.fmt = fmt
	parser.basic = uts35Pattern
	parser.unpadded = uts35PatternUnpadded

	const quote = '\''
	var literal bool
	dst := buffer(fmt)

	parser.writeLit = func(b byte) error {
		if b == quote {
			dst = append(dst, quote, quote)
			return nil
		}
		if !literal && ('a' <= b && b <= 'z' || 'A' <= b && b <= 'Z') {
			literal = true
			dst = append(dst, quote)
		}
		dst = append(dst, b)
		return nil
	}

	parser.writeFmt = func(fmt string) error {
		if literal {
			literal = false
			dst = append(dst, quote)
		}
		dst = append(dst, fmt...)
		return nil
	}

	parser.fallback = func(spec byte, pad bool) error {
		return errors.New("strftime: unsupported specifier: %" + string(spec))
	}

	if err := parser.parse(); err != nil {
		return "", err
	}

	parser.writeFmt("")
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
