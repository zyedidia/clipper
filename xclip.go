package clipper

import (
	"os/exec"
)

// xclip for linux X
type Xclip struct{}

func (x *Xclip) Init() error {
	return verify(x, "xclip")
}

func (x *Xclip) ReadAll(reg string) ([]byte, error) {
	var cmd *exec.Cmd
	switch reg {
	case RegClipboard:
		cmd = exec.Command("xclip", "-out", "-selection", "clipboard")
	case RegPrimary:
		cmd = exec.Command("xclip", "-out")
	default:
		return nil, &ErrInvalidReg{
			Reg: reg,
		}
	}
	return cmd.Output()
}

func (x *Xclip) WriteAll(reg string, p []byte) error {
	var cmd *exec.Cmd
	switch reg {
	case RegClipboard:
		cmd = exec.Command("xclip", "-in", "-selection", "clipboard")
	case RegPrimary:
		cmd = exec.Command("xclip", "-in")
	default:
		return &ErrInvalidReg{
			Reg: reg,
		}
	}
	return write(cmd, p)
}
