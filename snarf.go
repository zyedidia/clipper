package clipper

import (
	"io/ioutil"
	"os"
)

// snarf for plan9
type Snarf struct{}

func (s *Snarf) Init() error {
	_, err := os.Stat("/dev/snarf")
	return err
}

func (s *Snarf) ReadAll(reg string) ([]byte, error) {
	if reg != RegClipboard {
		return nil, &ErrInvalidReg{
			Reg: reg,
		}
	}
	return ioutil.ReadFile("/dev/snarf")
}

func (s *Snarf) WriteAll(reg string, p []byte) error {
	if reg != RegClipboard {
		return &ErrInvalidReg{
			Reg: reg,
		}
	}
	return ioutil.WriteFile("/dev/snarf", p, 0666)
}
