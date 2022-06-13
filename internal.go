package clipper

type Internal struct {
	regs map[string][]byte
}

func (i *Internal) Init() error {
	i.regs = map[string][]byte{
		RegClipboard: []byte{},
		RegPrimary:   []byte{},
	}
	return nil
}

func (i *Internal) ReadAll(reg string) ([]byte, error) {
	if p, ok := i.regs[reg]; ok {
		b := make([]byte, len(p))
		copy(b, p)
		return b, nil
	} else {
		return nil, &ErrInvalidReg{
			Reg: reg,
		}
	}
}

func (i *Internal) WriteAll(reg string, p []byte) error {
	i.regs[reg] = make([]byte, len(p))
	copy(i.regs[reg], p)
	return nil
}
