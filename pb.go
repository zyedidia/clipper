package clipper

import (
	"os/exec"
)

// pbpaste/pbcopy for macos
type Pb struct{}

func (pb *Pb) Init() error {
	return verify(pb, "pbpaste", "pbcopy")
}

func (pb *Pb) ReadAll(reg string) ([]byte, error) {
	if reg != RegClipboard {
		return nil, &ErrInvalidReg{
			Reg: reg,
		}
	}
	cmd := exec.Command("pbpaste")
	return cmd.Output()
}

func (pb *Pb) WriteAll(reg string, p []byte) error {
	if reg != RegClipboard {
		return &ErrInvalidReg{
			Reg: reg,
		}
	}
	cmd := exec.Command("pbcopy")
	return write(cmd, p)
}
