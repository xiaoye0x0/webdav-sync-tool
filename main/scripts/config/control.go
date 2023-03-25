package config

import (
	"WebDavSyncTool/scripts/logLevel"
	"path/filepath"

	"os"

	"github.com/go-ini/ini"
)

var webDavConfing WebDavConf
var localConfing LocalConf

func init() {
	rootPath, _ := filepath.Abs("")
	cfgs, err := ini.Load(filepath.Join(rootPath, "conf.ini"))
	if err != nil {
		logLevel.Errorf(err.Error())
		os.Exit(1)
	}
	webDavConfing.ServerPath = getServerPath(cfgs)
	webDavConfing.Username = getUsername(cfgs)
	webDavConfing.Password = getPassword(cfgs)
	webDavConfing.FilePath = getFilePath(cfgs)
	localConfing.LocalPath = getLocalPath(cfgs)
}

func getServerPath(cfgs *ini.File) string {
	r := cfgs.Section("server").Key("url").Value()
	if r == "" {
		logLevel.Errorf("未获取到远端服务器地址")
		os.Exit(1)
	}
	return r
}

func getUsername(cfgs *ini.File) string {
	return cfgs.Section("server").Key("username").Value()
}

func getPassword(cfgs *ini.File) string {
	return cfgs.Section("server").Key("password").Value()
}

func getFilePath(cfgs *ini.File) string {
	r := cfgs.Section("server").Key("filePath").Value()
	if r == "" {
		logLevel.Errorf("未获取到远端服务器文件路径")
		os.Exit(1)
	}
	return r
}

func getLocalPath(cfgs *ini.File) string {
	r := cfgs.Section("local").Key("savePath").Value()
	if r == "" {
		logLevel.Warningf("未获取到本地文件存放路径,将存放在同级目录")
		r = "./"
	}
	return r
}

func GetWebDavConfig() *WebDavConf {
	return &webDavConfing
}

func GetLocalConfig() *LocalConf {
	return &localConfing
}
