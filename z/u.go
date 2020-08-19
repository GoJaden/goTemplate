package z

import (
	"fmt"
	"github.com/lhlyu/logger/v3/color"
	"os"
	"path"
	"strings"
)

func CheckDir() {
	pt, _ := os.Getwd()
	gopath := os.Getenv("GOPATH")
	if !strings.Contains(pt, gopath) {
		fmt.Println(color.Red("【警告】当前目录不是在$GOPATH/src子目录下！！！\n"))
	}
	return
}

func GetWorkDir() string {
	pt, _ := os.Getwd()
	if pt == "" {
		return ""
	}
	index := strings.Index(pt, "src")
	if index == -1 {
		return ""
	}
	index += 4
	if len(pt) < index {
		pt = ""
	} else {
		pt = pt[index:]
		pt = strings.ReplaceAll(pt, "\\", "/") + "/"
	}

	return pt
}

func Exists(dir string) bool {
	_, err := os.Stat(dir)
	if err == nil {
		return true
	}
	if os.IsNotExist(err) {
		return false
	}
	return false
}

func MakeDirs(s ...string) string {
	root := path.Join(s...)
	if !Exists(root) {
		os.MkdirAll(root, os.ModePerm)
	}
	return root
}
