//go:build !windows && !darwin && !plan9

package clipper

var Clipboards = []Clipboard{
	&Xclip{},
	&Xsel{},
	&Wayland{},
	&Wsl{},
	&Termux{},
}
