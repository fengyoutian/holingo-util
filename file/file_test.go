package file

import (
	"fmt"
	"os"
	"path/filepath"
	"testing"
)

func TestGetAppPath(t *testing.T) {
	fmt.Println(os.Getenv("GOPATH"))
	fmt.Println(GetRunPath())
	fmt.Println(GetRunDir())
	fmt.Println(GetBuildDir())
	fmt.Println(Exists(filepath.Join(GetRunDir(), "file.go")))
	fmt.Println(Exists(filepath.Join(GetBuildDir(), "file.go")))
}

func TestExists(t *testing.T) {
	type RunOrBuild string
	const (
		RUN RunOrBuild = "RUN"
		BUILD RunOrBuild = "BUILD"
	)

	tests := []struct{
		runOrBuild RunOrBuild
		fileName string
		exists bool
	} {
		{RUN, "./file.go", false},
		{BUILD, "./file.go", true},
		{RUN, "./file_test.go", false},
		{BUILD, "./file_test.go", true},
		{RUN, "./test", false},
		{BUILD, "./test", false},
		{RUN, "./", true}, // 判断目录
		{BUILD, "./", true},
	}

	for _, v := range tests {
		var dir string
		switch v.runOrBuild {
		case RUN:
			dir = GetRunDir()
		case BUILD:
			dir = GetBuildDir()
		}
		if exists := Exists(filepath.Join(dir, v.fileName)); exists != v.exists {
			t.Errorf("%s was %t, but it %t", v.fileName, exists, v.exists)
		}
	}
}

func TestGetSize(t *testing.T) {
	// 直接测试本代码，所以文件大小的测试数据结果这里就不给了
	datas := []struct{
		fileName string
	} {
		{"./file.go"},
		{"file.go"},
		{"file_test.go"},
		{""},
		{"."},
		{"./"},
		{"../"},
		{"../log"},
		{"../log/log.go"},
		{"../log/log_test.go"},
		{"../../"},
	}

	dir := GetBuildDir()
	for _, v := range datas {
		t.Logf("[%s] size: %d", v.fileName, GetSize(filepath.Join(dir, v.fileName)))
	}
}

func TestIsDir(t *testing.T) {
	datas := []struct{
		fileName string
		result bool
	} {
		{"./", true},
		{"./file.go", false},
		{"", true},
		{"../", true},
	}

	for _, v := range datas {
		if IsDir(filepath.Join(GetBuildDir(), v.fileName)) != v.result {
			t.Errorf("[%s] not dir", v.fileName)
		}
	}
}