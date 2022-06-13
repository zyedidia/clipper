//go:build windows

package clipper

func GetClipboard() (clip Clipboard, err error) {
	return &WinApi{}, nil
}
