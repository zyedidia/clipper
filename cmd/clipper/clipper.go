package main

import (
	"flag"
	"fmt"
	"io"
	"os"

	"github.com/zyedidia/clipper"
)

func must(err error) {
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}

var paste = flag.Bool("paste", false, "paste clipboard contents")
var cpy = flag.Bool("copy", false, "copy into clipboard")
var reg = flag.String("reg", "clipboard", "clipboard register to use")

func prepend[T any](c T, cs []T) []T {
	return append([]T{c}, cs...)
}

func main() {
	flag.Parse()

	clip, err := clipper.GetClipboard(prepend[clipper.Clipboard](&clipper.Custom{
		Name: "clipper-clip",
	}, clipper.Clipboards)...)
	must(err)

	if *cpy {
		data, err := io.ReadAll(os.Stdin)
		must(err)
		err = clip.WriteAll(*reg, data)
		must(err)
	}

	if *paste {
		data, err := clip.ReadAll(*reg)
		must(err)
		fmt.Print(string(data))
	}
}
