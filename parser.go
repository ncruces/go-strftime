package strftime

type parser struct {
	format  func(spec, flag byte) error
	literal func(byte) error
}

func (p *parser) parse(fmt string) error {
	const (
		initial = iota
		specifier
		nopadding
	)

	var err error
	state := initial
	for _, b := range []byte(fmt) {
		switch state {
		default:
			if b == '%' {
				state = specifier
				continue
			}
			err = p.literal(b)

		case specifier:
			if b == '-' {
				state = nopadding
				continue
			}
			err = p.format(b, '0')
			state = initial

		case nopadding:
			err = p.format(b, 0)
			state = initial
		}

		if err != nil {
			return err
		}
	}

	switch state {
	default:
		return nil
	case specifier:
		return p.literal('%')
	case nopadding:
		err := p.literal('%')
		if err != nil {
			return err
		}
		return p.literal('-')
	}
}
