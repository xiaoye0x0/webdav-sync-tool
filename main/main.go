package main

import (
	"WebDavSyncTool/app"
	"WebDavSyncTool/scripts/logLevel"
)

func main() {
	logLevel.Debugf("Running")
	app.DownloadAll()
}
