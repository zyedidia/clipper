package clipper

import (
	"fmt"
	"os/exec"
)

type Clipboard interface {
	Init() error
	ReadAll(reg string) ([]byte, error)
	WriteAll(reg string, p []byte) error
}

const (
	RegClipboard = "clipboard"
	RegPrimary   = "primary"
)

type ErrInvalidReg struct {
	Reg string
}

func (e *ErrInvalidReg) Error() string {
	return fmt.Sprintf("invalid register: %s", e.Reg)
}

func verify(clip Clipboard, cmds ...string) error {
	for _, cmd := range cmds {
		if _, err := exec.LookPath(cmd); err != nil {
			// command not installed
			return err
		}
	}
	_, err := clip.ReadAll(RegClipboard)
	if err == nil {
		// reading clipboard worked
		return nil
	}
	// reading could fail if the clipboard has no contents, so check if writing
	// works in that case
	return clip.WriteAll(RegClipboard, []byte{})
}

func write(cmd *exec.Cmd, p []byte) error {
	in, err := cmd.StdinPipe()
	if err != nil {
		return err
	}
	if err := cmd.Start(); err != nil {
		return err
	}
	if _, err := in.Write(p); err != nil {
		return err
	}
	in.Close()
	return cmd.Wait()
}
