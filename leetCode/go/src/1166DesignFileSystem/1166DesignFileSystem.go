package designfilesystem

import (
	libPath "path"
	"unicode/utf8"
)

type FileSystem struct {
	pathMap map[string]int
}

func Constructor() FileSystem {
	fs := FileSystem{}
	fs.pathMap = make(map[string]int, 0)
	return fs
}

func (this *FileSystem) Create(path string, value int) bool {
	if len(path) < 2 {
		return false
	}
	r, _ := utf8.DecodeRuneInString(path)
	if r != '/' {
		return false
	}
	// check parent path
	parentDir := libPath.Dir(path)
	if _, exist := this.pathMap[parentDir]; !exist && parentDir != "/" {
		return false
	}
	// check current path
	if _, exist := this.pathMap[path]; !exist {
		this.pathMap[path] = value
		return true
	}
	return false
}

func (this *FileSystem) Get(path string) int {
	if _, exist := this.pathMap[path]; exist {
		return this.pathMap[path]
	}
	return -1
}

/**
 * Your FileSystem object will be instantiated and called as such:
 * obj := Constructor();
 * param_1 := obj.Create(path,value);
 * param_2 := obj.Get(path);
 */
