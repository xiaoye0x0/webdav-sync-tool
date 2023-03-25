package download

import (
	"WebDavSyncTool/scripts/logLevel"
)

func DownloadAll() {
	targetFiles, err := FindFiles()
	if err != nil {
		logLevel.Errorf(err.Error())
		return
	}
	err = Download(targetFiles)
	if err != nil {
		logLevel.Errorf(err.Error())
	}
	logLevel.Infof("The download is complete")
}
