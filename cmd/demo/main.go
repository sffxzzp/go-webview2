package main

import (
	"embed"
	"io/fs"
	"net/http"
	"os"
	"path/filepath"

	"github.com/sffxzzp/go-webview2"
)

//go:embed static
var embedFS embed.FS

func curPath() string {
	ex, err := os.Executable()
	if err != nil {
		panic(err)
	}
	return filepath.Dir(ex)
}

func main() {
	dataPath := curPath() + "/data"
	iconPath := dataPath + "/favicon.ico"
	name := "Aria2 WebUI"
	ebd, _ := fs.Sub(embedFS, "static")
	fServer := http.FileServer(http.FS(ebd))
	go http.ListenAndServe("127.0.0.1:65533", fServer)
	icon, _ := fs.ReadFile(ebd, "favicon.ico")
	os.Mkdir(curPath()+"/data", 0777)
	os.WriteFile(iconPath, icon, 0777)
	w := webview2.NewWithOptions(webview2.WebViewOptions{
		Debug:    false,
		DataPath: dataPath,
		WindowOptions: webview2.WindowOptions{
			Title:  name,
			Width:  1440,
			Height: 900,
			Icon:   iconPath,
		},
	})
	defer w.Destroy()
	w.Navigate("http://127.0.0.1:65533/")
	w.Run()
}
