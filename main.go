package main

import (
	"github.com/zserge/webview"
)

const (
	windowWidth  = 800
	windowHeight = 600
)


func HandleRPC(w webview.WebView, convCfgStr string) {

	saveCfgToDisk(convCfgStr)

	// Passing global cmd (*exec.Cmd) argument
	cmdGlo = restartCliApp(cmdGlo, convCfgStr)
}

func main() {

	ifDae := false
	convCfgStr := loadCfgFromDisk()

	// Can't use ":=" here
	cmdGlo, ifDae = startCliApp(convCfgStr)
	
	url := StartHTTPServ()
	
	w := webview.New(webview.Settings{
		Width:     windowWidth,
		Height:    windowHeight,
		Title:     "Poetry desktop",
		Resizable: true,
		URL:       url,
		ExternalInvokeCallback: HandleRPC,
	})


	defer w.Exit()
	w.Run()

	// If daemon
	if ifDae == false {
		
		// Use this no argument function is because
		// what passed to function in a defer scope won't be changed anymore
		// but when we restart the cli app, the cmdGlo will be changed
		defer StopCliAppWithoutArgv()
	}
}
