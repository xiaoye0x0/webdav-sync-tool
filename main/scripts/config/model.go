package config

import "github.com/studio-b12/gowebdav"

type WebDavConf struct {
	ServerPath   string //webdav url
	Username     string
	Password     string
	FilePath     string //远端服务器文件夹路径
	WebdavClient *gowebdav.Client
}

type LocalConf struct {
	LocalPath string
}
