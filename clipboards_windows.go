//go:build windows

package clipper

var Clipboards = []Clipboard{
	&WinApi{},
	&Internal{},
}
