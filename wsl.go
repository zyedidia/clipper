package clipper

import (
	"os/exec"
)

// powershell.exe/clip.exe for WSL
type Wsl struct{}

func (w *Wsl) Init() error {
	return verify(w, "powershell.exe", "clip.exe")
}

func (w *Wsl) ReadAll(reg string) ([]byte, error) {
	var cmd *exec.Cmd
	switch reg {
	case RegClipboard:
		cmd = exec.Command("powershell.exe", "Get-Clipboard")
	default:
		return nil, &ErrInvalidReg{
			Reg: reg,
		}
	}
	out, err := cmd.Output()
	if err != nil {
		return nil, err
	}
	if len(out) > 1 {
		out = out[:len(out)-2]
	}
	return out, nil
}

func (w *Wsl) WriteAll(reg string, p []byte) error {
	var cmd *exec.Cmd
	switch reg {
	case RegClipboard:
		cmd = exec.Command("clip.exe")
	default:
		return &ErrInvalidReg{
			Reg: reg,
		}
	}
	return write(cmd, p)
}
