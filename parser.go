package strftime

type parser struct {
	format  func(i int, spec, flag byte) error
	literal func(i int, lit byte) error
}

func (p *parser) parse(fmt string) error {
	const (
		initial = iota
		specifier
		nopadding
	)

	var err error
	state := initial
	for i, b := range []byte(fmt) {
		switch state {
		default:
			if b == '%' {
				state = specifier
				continue
			}
			err = p.literal(i, b)

		case specifier:
			if b == '-' {
				state = nopadding
				continue
			}
			err = p.format(i, b, '0')
			state = initial

		case nopadding:
			err = p.format(i, b, 0)
			state = initial
		}

		if err != nil {
			return err
		}
	}

	switch state {
	case specifier:
		p.literal(0, '%')
	case nopadding:
		p.literal(0, '%')
		p.literal(0, '-')
	}
	return nil
}
