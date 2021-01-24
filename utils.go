package render

import "os"

//isExistingDir 判断目录是否存在
func isExistingDir(pth string) bool {
	if fi, err := os.Stat(pth); err == nil {
		return fi.Mode().IsDir()
	}
	return false
}
