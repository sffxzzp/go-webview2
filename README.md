# go-webview2
This package is a fork of jchv's [go-webview2](https://github.com/jchv/go-webview2/).

Features:
1. Width & Height init
2. Window are now create on screen center (I don't know if it's working currectly with multi-screen user)
3. Icon support (I tried to use `CreateIconFromResource` but failed, maybe someone could help me with this?)
4. `Debug` options finally work. it'll disable the context menu and devtool.