//go:build !windows

package clipper

import "bytes"

type MultiErr []error

func (me MultiErr) Error() string {
	b := &bytes.Buffer{}
	for _, e := range me {
		b.WriteString(e.Error())
		b.WriteByte(';')
	}
	return b.String()
}

func GetClipboard() (clip Clipboard, err error) {
	var errs MultiErr

	clip = &Custom{
		Name: "myclip",
	}
	if err = clip.Init(); err == nil {
		return clip, nil
	}
	errs = append(errs, err)

	clip = &Xclip{}
	if err = clip.Init(); err == nil {
		return clip, nil
	}
	errs = append(errs, err)

	clip = &Xsel{}
	if err = clip.Init(); err == nil {
		return clip, nil
	}
	errs = append(errs, err)

	clip = &Wayland{}
	if err = clip.Init(); err == nil {
		return clip, nil
	}
	errs = append(errs, err)

	clip = &Wsl{}
	if err = clip.Init(); err == nil {
		return clip, nil
	}
	errs = append(errs, err)

	clip = &Pb{}
	if err = clip.Init(); err == nil {
		return clip, nil
	}
	errs = append(errs, err)

	clip = &Internal{}
	clip.Init()
	return clip, errs
}
