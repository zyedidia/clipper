package clipper

import (
	"os/exec"
)

// termux-clipboard for Android Termux
type Termux struct{}

func (t *Termux) Init() error {
	return verify(t, "termux-clipboard-get", "termux-clipboard-set")
}

func (t *Termux) ReadAll(reg string) ([]byte, error) {
	var cmd *exec.Cmd
	switch reg {
	case RegClipboard:
		cmd = exec.Command("termux-clipboard-get")
	default:
		return nil, &ErrInvalidReg{
			Reg: reg,
		}
	}
	return cmd.Output()
}

func (t *Termux) WriteAll(reg string, p []byte) error {
	var cmd *exec.Cmd
	switch reg {
	case RegClipboard:
		cmd = exec.Command("termux-clipboard-set")
	default:
		return &ErrInvalidReg{
			Reg: reg,
		}
	}
	return write(cmd, p)
}
