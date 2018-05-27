package main

// TODO: Quit on esc
// TODO: Set appropriate dimensions
// TODO: Command line options for dimensions etc
// TODO: Clickable links open in other browser, which can be specified
// TODO: Keyboard shortcuts for paging down (and links?)

import (
    "github.com/zserge/webview"
    // "mime/quotedprintable"
    "os"
    "fmt"
    "path/filepath"

)

func handleRPC(w webview.WebView, data string) {
    fmt.Println("Hello")
}


func main() {
    // Open wikipedia in a 800x600 resizable window
    filename := os.Args[1]
    fp, _ := filepath.Abs(filename)
    // webview.Open("Minimal webview example",
        // "https://en.m.wikipedia.org/wiki/Main_Page", 800, 600, true)
    w := webview.New(webview.Settings{
        Width:  500,
        Height: 600,
        Title:  "Local file",
        URL:    "file:///" + fp,
        ExternalInvokeCallback: handleRPC,
	})
    defer w.Exit()
    w.Run()
    w.Eval("window.external.invoke_('hello')")
    // webview.Open("Local File", "file:///" + fp, 800, 600, true)
    // webview.ex1

    // webview.Terminate()
}
