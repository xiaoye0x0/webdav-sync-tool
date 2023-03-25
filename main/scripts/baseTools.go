package scripts

import "os"

// 判断路径是否为文件夹 是为True 不是为False
func IsDir(path string) bool {
	s, err := os.Stat(path)
	if err != nil {
		return false
	}
	return s.IsDir()
}

func MakeDir(path string) {
	os.MkdirAll(path, os.ModePerm)
}
