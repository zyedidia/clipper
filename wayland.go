package clipper

import (
	"fmt"
	"os"
	"os/exec"
)

// wl-paste/wl-copy for linux Wayland
type Wayland struct{}

func (wl *Wayland) Init() error {
	if os.Getenv("WAYLAND_DISPLAY") != "" {
		return verify(wl, "wl-paste", "wl-copy")
	}
	return fmt.Errorf("Wayland display not found")
}

func (wl *Wayland) ReadAll(reg string) ([]byte, error) {
	var cmd *exec.Cmd
	switch reg {
	case RegClipboard:
		cmd = exec.Command("wl-paste", "--no-newline")
	case RegPrimary:
		cmd = exec.Command("wl-paste", "--no-newline", "--primary")
	default:
		return nil, &ErrInvalidReg{
			Reg: reg,
		}
	}
	return cmd.Output()
}

func (wl *Wayland) WriteAll(reg string, p []byte) error {
	var cmd *exec.Cmd
	switch reg {
	case RegClipboard:
		cmd = exec.Command("wl-copy")
	case RegPrimary:
		cmd = exec.Command("wl-copy", "--primary")
	default:
		return &ErrInvalidReg{
			Reg: reg,
		}
	}
	return write(cmd, p)
}
