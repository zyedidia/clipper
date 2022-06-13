# Clipper: cross-platform clipboard library

[![Go Reference](https://pkg.go.dev/badge/github.com/zyedidia/clipper.svg)](https://pkg.go.dev/github.com/zyedidia/clipper)
[![MIT License](https://img.shields.io/badge/license-MIT-blue.svg)](https://github.com/zyedidia/clipper/blob/master/LICENSE)

Platforms supported:

* Linux (via `xclip` or `xsel` or `wl-copy`/`wl-paste`)
* MacOS (via `pbcopy`/`pbpaste`)
* Windows (via the Windows clipboard API)
* WSL (via `clip.exe`/`powershell.exe`)
* Android Termux (via `termux-clipboard-set`/`termux-clipboard-get`)
* Plan9 (via `/dev/snarf`)
* Anything else (via user-defined `myclip` script)

Fallback methods:

* Internal in-memory clipboard
* File-based clipboard

# Example

```
func main() {
    clip, err := clipper.GetClipboard()
    must(err)

    // copy from stdin
    data, err := io.ReadAll(os.Stdin)
    must(err)
    err = clip.WriteAll(clipper.RegClipboard, data)
    must(err)

    // paste to stdout
    data, err := clip.ReadAll(clipper.RegClipboard)
    must(err)
    fmt.Print(string(data))
}
```

A CLI tool is provided as an example in `cmd/clipper`.
