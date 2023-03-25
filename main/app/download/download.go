package download

import (
	"WebDavSyncTool/scripts"
	"WebDavSyncTool/scripts/config"
	"fmt"
	"io/ioutil"
	basePath "path"
	"path/filepath"

	"github.com/studio-b12/gowebdav"
)

func findAllFiles(c *gowebdav.Client, path string) (totalFiles []string, err error) {
	res, err := c.ReadDir(path)
	if err != nil {
		return
	}
	for _, file := range res {
		if file.IsDir() {
			childFiles, childErr := findAllFiles(c, basePath.Join(path, file.Name()))
			if childErr != nil {
				err = childErr
				return
			}
			totalFiles = append(totalFiles, childFiles...)
		} else {
			totalFiles = append(totalFiles, basePath.Join(path, file.Name()))
		}

	}
	return
}

func getWebDavClient() {
	webdavConf := config.GetWebDavConfig()
	if webdavConf.WebdavClient != nil {
		return
	}

	webdavConf.WebdavClient = gowebdav.NewClient(
		webdavConf.ServerPath,
		webdavConf.Username,
		webdavConf.Password,
	)
}

func FindFiles() (targetFiles []string, err error) {
	webdavConf := config.GetWebDavConfig()
	getWebDavClient()

	targetFiles, err = findAllFiles(webdavConf.WebdavClient, webdavConf.FilePath)

	if len(targetFiles) == 0 {
		err = fmt.Errorf("can't find target files (%s)", webdavConf.FilePath)
	}
	return
}

func Download(files []string) (err error) {
	var totalErr string
	webdavConf := config.GetWebDavConfig()
	localConf := config.GetLocalConfig()
	getWebDavClient()

	if !scripts.IsDir(localConf.LocalPath) {
		scripts.MakeDir(localConf.LocalPath)
	}

	for _, filePath := range files {
		bytes, oneErr := webdavConf.WebdavClient.Read(filePath)
		if oneErr != nil {
			totalErr += fmt.Sprint("\n", oneErr.Error())
			continue
		}

		fileDir, fileName := filepath.Split(filePath)
		fileDir = filepath.Join(localConf.LocalPath, fileDir)
		if !scripts.IsDir(fileDir) {
			scripts.MakeDir(fileDir)
		}
		writeErr := ioutil.WriteFile(filepath.Join(fileDir, fileName), bytes, 0644)
		if writeErr != nil {
			totalErr += fmt.Sprint("\n", writeErr.Error())
			continue
		}
	}
	if totalErr != "" {
		err = fmt.Errorf(totalErr)
	}
	return
}
