package utils

import (
	"io/ioutil"
	"os"
	"path"
)

func PathExists(path string) (bool, error) {
	_, err := os.Stat(path)
	if err == nil {
		return true, nil
	}
	if os.IsNotExist(err) {
		return false, nil
	}
	return false, err
}

func RemoveDir(dirName string) (err error) {
	dir, err := ioutil.ReadDir(dirName)
	if err != nil {
		return
	}
	for _, d := range dir {
		err = os.RemoveAll(path.Join([]string{dirName, d.Name()}...))
		if err != nil {
			return
		}
	}
	return nil
}

//判断文件是否存在  存在返回 true 不存在返回false
func CheckFileIsExist(filename string) bool {
	var exist = true
	if _, err := os.Stat(filename); os.IsNotExist(err) {
		exist = false
	}
	return exist
}
