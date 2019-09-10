// Copyright 2019 Andyfoo
//
// Licensed under the Apache License, Version 2.0 (the "License"): you may
// not use this file except in compliance with the License. You may obtain
// a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS, WITHOUT
// WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied. See the
// License for the specific language governing permissions and limitations
// under the License.

package xfile

import (
	"io"
	"io/ioutil"
	"os"
	"os/exec"
	"path/filepath"
	"regexp"
	"strings"

	"github.com/Andyfoo/go-xutils/xlog"
)

//获取基本文件名
func BaseName(filename string) string {
	return filepath.Base(filename)
}

//获取路径名
func DirName(filename string) string {
	return filepath.Dir(filename)
}

//获取扩展名,不带"."
func ExtName(filename string) string {
	ext := filepath.Ext(filename)
	pos := strings.Index(ext, ".")
	if pos > -1 {
		return ext[pos+1:]
	}
	return ext
}

//替换路径多余分隔符，统一为 "/"
func FormatPath(path string) string {
	re, _ := regexp.Compile("[\\\\/]+")
	path = re.ReplaceAllString(path, "/")
	return path
}

//替换路径多余分隔符，统一为系统分隔符
func FormatPathSys(path string) string {
	re, _ := regexp.Compile("[\\\\/]+")
	path = re.ReplaceAllString(path, string(os.PathSeparator))
	return path
}

//获取当前exe路径
func GetCurrentPath() string {
	file, err := exec.LookPath(os.Args[0])
	if err != nil {
		return ""
	}
	path, err := filepath.Abs(file)
	if err != nil {
		return ""
	}
	i := strings.LastIndex(path, "/")
	if i < 0 {
		i = strings.LastIndex(path, "\\")
	}
	if i < 0 {
		return ""
	}
	return string(path[0 : i+1])
}

// SelfPath gets compiled executable file absolute path
func SelfPath() string {
	path, _ := filepath.Abs(os.Args[0])
	return path
}

// SelfDir gets compiled executable file directory
func SelfDir() string {
	return filepath.Dir(SelfPath())
}

//判断文件是否存在
func FileIsExist(path string) bool {
	_, err := os.Stat(path)
	return err == nil || os.IsExist(err)
	// 或者
	//return err == nil || !os.IsNotExist(err)
	// 或者
	//return !os.IsNotExist(err)
}

//判断是否文件夹
func IsDir(path string) bool {
	f, err := os.Stat(path)
	if err != nil {
		xlog.Error(err)
		return false
	}
	if os.IsNotExist(err) {
		xlog.Error(err)
		return false
	}
	if f.IsDir() {
		return true
	}
	return false
}

//创建目录
func Mkdir(path string) bool {
	err := os.Mkdir(path, 777)
	if err == nil {
		return true
	}
	return false
}

//创建目录
func MkdirAll(path string) bool {
	err := os.MkdirAll(path, 777)
	if err == nil {
		return true
	}
	return false
}

//获取目录文件列表
func GetDirList(dirpath string) []string {
	var dir_list []string
	dir_err := filepath.Walk(dirpath,
		func(path string, f os.FileInfo, err error) error {
			if f == nil {
				xlog.Error("err", err)
				return nil
			}
			if f.IsDir() {
				dir_list = append(dir_list, path)
				return nil
			}

			return nil
		})
	if dir_err != nil {
		xlog.Error("dir_err", dir_err)
	}
	return dir_list
}

//获取文件列表
func GetFileList(dirpath string) []string {
	var dir_list []string
	dir_err := filepath.Walk(dirpath,
		func(path string, f os.FileInfo, err error) error {
			if f == nil {
				xlog.Error("err", err)
				return nil
			}
			if !f.IsDir() {
				dir_list = append(dir_list, path)
				return nil
			}

			return nil
		})
	if dir_err != nil {
		xlog.Error("dir_err", dir_err)
	}
	return dir_list
}

//复制文件
func CopyFile(src, dst string) int64 {
	sourceFileStat, err := os.Stat(src)
	if err != nil {
		xlog.Error("err", err)
		return 0
	}
	if !sourceFileStat.Mode().IsRegular() {
		xlog.Errorf("%s is not a regular file", src)
		return 0
	}

	source, err := os.Open(src)
	if err != nil {
		xlog.Error("err", err)
		return 0
	}
	defer source.Close()

	destination, err := os.Create(dst)
	if err != nil {
		xlog.Error("err", err)
		return 0
	}
	defer destination.Close()
	nBytes, err := io.Copy(destination, source)
	return nBytes
}

func ReadFile(filename string) []byte {
	contents, err := ioutil.ReadFile(filename)
	if err != nil {
		xlog.Error("err", err)
		return nil
	}
	return contents
}
func ReadFileStr(filename string) string {
	contents, err := ioutil.ReadFile(filename)
	if err != nil {
		xlog.Error("err", err)
		return ""
	}
	return string(contents)
}
func WriteFile(filename string, data []byte, _mode ...string) bool {
	mode := ""
	if len(_mode) > 0 {
		mode = _mode[0]
	}
	perm := os.FileMode(0644)
	flag := os.O_WRONLY | os.O_CREATE | os.O_TRUNC
	if mode == "a+" {
		flag = os.O_RDWR | os.O_CREATE | os.O_APPEND
	}
	f, err := os.OpenFile(filename, flag, perm)
	if err != nil {
		xlog.Error("err", err)
		return false
	}
	n, err := f.Write(data)
	if err == nil && n < len(data) {
		err = io.ErrShortWrite
	}
	if err1 := f.Close(); err == nil {
		err = err1
	}
	if err != nil {
		xlog.Error("err", err)
		return false
	}
	return true
}

func WriteFileStr(filename string, data string, _mode ...string) bool {
	mode := ""
	if len(_mode) > 0 {
		mode = _mode[0]
	}
	return WriteFile(filename, []byte(data), mode)
}
