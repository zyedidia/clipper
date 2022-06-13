package clipper

import (
	"os/exec"
)

type Custom struct {
	Name string
}

func (c *Custom) Init() error {
	return verify(c, c.Name)
}

func (c *Custom) ReadAll(reg string) ([]byte, error) {
	cmd := exec.Command(c.Name, "-o", reg)
	return cmd.Output()
}

func (c *Custom) WriteAll(reg string, p []byte) error {
	cmd := exec.Command(c.Name, "-i", reg)
	return write(cmd, p)
}
