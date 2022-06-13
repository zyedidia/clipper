# Clipper: cross-platform clipboard library

Platforms supported:

* Linux (via `xclip` or `xsel` or `wl-copy`/`wl-paste`)
* MacOS (via `pbcopy`/`pbpaste`)
* Windows (via the Windows clipboard API)
* WSL (via `clip.exe`/`powershell.exe`)
* Android Termux (via `termux-clipboard-set`/`termux-clipboard-get`)
* Plan9 (via `/dev/snarf`)

Fallback methods:

* Internal in-memory clipboard
* File-based clipboard
* Custom user-defined `myclip` script
