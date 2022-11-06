//go:build !windows && !darwin && !plan9

package clipper

var Clipboards = []Clipboard{
	&Wayland{},
	&Xclip{},
	&Xsel{},
	&Wsl{},
	&Termux{},
}
