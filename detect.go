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

// GetClipboard iterates through the given clipboards and returns the first one that works.
func GetClipboard(clips ...Clipboard) (clip Clipboard, err error) {
	var errs MultiErr

	for _, clip := range clips {
		if err = clip.Init(); err == nil {
			return clip, nil
		}
		errs = append(errs, err)
	}
	return nil, errs
}
