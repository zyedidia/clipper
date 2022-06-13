//go:build !windows

package clipper

var Clipboards = []Clipboard{
	&Pb{},
	&Xclip{},
	&Xsel{},
	&Wayland{},
	&Wsl{},
	&Termux{},
	&Internal{},
}
