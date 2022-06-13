package clipper

import (
	"os/exec"
)

// xsel for linux X
type Xsel struct{}

func (x *Xsel) Init() error {
	return verify(x, "xsel")
}

func (x *Xsel) ReadAll(reg string) ([]byte, error) {
	var cmd *exec.Cmd
	switch reg {
	case RegClipboard:
		cmd = exec.Command("xsel", "--output", "--clipboard")
	case RegPrimary:
		cmd = exec.Command("xsel", "--output")
	default:
		return nil, &ErrInvalidReg{
			Reg: reg,
		}
	}
	return cmd.Output()
}

func (x *Xsel) WriteAll(reg string, p []byte) error {
	var cmd *exec.Cmd
	switch reg {
	case RegClipboard:
		cmd = exec.Command("xsel", "--input", "--clipboard")
	case RegPrimary:
		cmd = exec.Command("xsel", "--input")
	default:
		return &ErrInvalidReg{
			Reg: reg,
		}
	}
	return write(cmd, p)
}
