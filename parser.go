package strftime

type parser struct {
	fmt      string
	basic    func(byte) string
	unpadded func(byte) string
	writeLit func(byte) error
	writeFmt func(string) error
	fallback func(spec byte, pad bool) error
}

func (p *parser) parse() error {
	const (
		initial = iota
		specifier
		nopadding
	)

	var err error
	state := initial
	for _, b := range []byte(p.fmt) {
		switch state {
		default:
			if b == '%' {
				state = specifier
			} else {
				err = p.writeLit(b)
			}

		case specifier:
			if b == '-' {
				state = nopadding
				continue
			}
			if fmt := p.basic(b); fmt != "" {
				err = p.writeFmt(fmt)
			} else {
				err = p.fallback(b, true)
			}
			state = initial

		case nopadding:
			if fmt := p.unpadded(b); fmt != "" {
				err = p.writeFmt(fmt)
			} else if fmt := p.basic(b); fmt != "" {
				err = p.writeFmt(fmt)
			} else {
				err = p.fallback(b, false)
			}
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
		return p.writeLit('%')
	case nopadding:
		err := p.writeLit('%')
		if err != nil {
			return err
		}
		return p.writeLit('-')
	}
}
