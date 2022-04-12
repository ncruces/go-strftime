package strftime

type parser struct {
	format  func(spec, flag byte, i int) error
	literal func(byte) error
}

func (p *parser) parse(fmt string) error {
	const (
		initial = iota
		percent
		flagged
	)

	var flag byte
	var err error
	state := initial
	for i, b := range []byte(fmt) {
		switch state {
		default:
			if b == '%' {
				state = percent
				continue
			}
			err = p.literal(b)

		case percent:
			if b == '-' || b == ':' {
				state = flagged
				flag = b
				continue
			}
			err = p.format(b, 0, i)
			state = initial

		case flagged:
			err = p.format(b, flag, i)
			state = initial
		}

		if err != nil {
			return err
		}
	}

	switch state {
	case percent:
		p.literal('%')
	case flagged:
		p.literal('%')
		p.literal(flag)
	}
	return nil
}
