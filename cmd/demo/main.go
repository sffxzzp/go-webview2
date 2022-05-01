package main

import (
	"embed"
	"io/fs"
	"net/http"

	"github.com/sffxzzp/go-webview2"
)

//go:embed static
var embedFS embed.FS

func main() {
	dataPath := "./data"
	name := "Webview2 Example"
	ebd, _ := fs.Sub(embedFS, "static")
	fServer := http.FileServer(http.FS(ebd))
	go http.ListenAndServe("127.0.0.1:65533", fServer)
	w := webview2.NewWithOptions(webview2.WebViewOptions{
		Debug:    false,
		DataPath: dataPath,
		WindowOptions: webview2.WindowOptions{
			Title:  name,
			Width:  1440,
			Height: 900,
			IconId: 2, // using akavel/rsrc, 2 is the default icon id
		},
	})
	defer w.Destroy()
	w.Navigate("http://127.0.0.1:65533/")
	w.Run()
}
