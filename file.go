package clipper

import (
	"io/ioutil"
	"os"
	"path/filepath"
)

type File struct {
	Dir string
}

func (fb *File) Init() error {
	return os.MkdirAll(fb.Dir, os.ModePerm)
}

func (fb *File) ReadAll(reg string) ([]byte, error) {
	return ioutil.ReadFile(filepath.Join(fb.Dir, reg))
}

func (fb *File) WriteAll(reg string, p []byte) error {
	return ioutil.WriteFile(filepath.Join(fb.Dir, reg), p, 0666)
}
