package file

import (
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"

	"github.com/fengyoutian/holingo-util/str"
)

// GetRunPath: 获取 go run 可执行文件的绝对路径
func GetRunPath() string {
	path, _ := filepath.Abs(os.Args[0])
	return path
}

// GetRunDir: 获取 go run 可执行文件的目录
func GetRunDir() string {
	return filepath.Dir(GetRunPath())
}

// GetBuildDir: 获取 go build 执行的可执行文件目录
func GetBuildDir() string {
	dir, _ := os.Getwd()
	return dir
}

// GetParentDir: 获取 dir 的父目录
func GetParentDir(dir string) string {
	return str.Sub(dir, 0, strings.LastIndex(dir, string(os.PathSeparator)))
}

// Exists: 判断文件/目录是否存在
func Exists(path string) bool {
	if _, err := os.Stat(path); err != nil {
		return os.IsExist(err)
	}
	return true
}

// IsDir: 判断是否目录
func IsDir(path string) bool {
	info, _ := os.Stat(path)
	return info.IsDir()
}

// GetSize: 获取文件大小
// path 为目录时返回文件总大小
func GetSize(path string) (result int64) {

	var isDir bool
	filepath.Walk(path, func(filePath string, info os.FileInfo, err error) error {
		if path == filePath {
			if isDir = info.IsDir(); !isDir {
				result = info.Size()
			}
		} else if isDir && !info.IsDir() {
			result += info.Size()
		}
		return nil
	})
	return result
}

// ReadFile: 读取文件, 封装iotuil.ReadFile()
// path 文件所在路径
func ReadFile(path string) (file []byte, err error) {
	file, err = ioutil.ReadFile(path)
	return
}
