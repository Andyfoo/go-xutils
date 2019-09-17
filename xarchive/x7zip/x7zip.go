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

package x7zip

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"os"
	"syscall"

	"github.com/Andyfoo/go-xutils/xtime"

	"os/exec"
	"strings"

	"github.com/Andyfoo/go-xutils/xfile"

	"github.com/Andyfoo/go-xutils/xlog"

	"golang.org/x/text/encoding/simplifiedchinese"
	"golang.org/x/text/transform"
)

var (
	BinPath = "bin/7z.exe"
)

func Clear(path string) {
	if xfile.FileIsExist(path) {
		xlog.Info("clear temp path:", path)
		os.RemoveAll(path)
	}

}

func UnRar(extName string, data []byte, objPath string) []string {
	if len(BinPath) == 0 {
		xlog.Error("BinPath is null")
		return nil
	}
	if !xfile.FileIsExist(BinPath) {
		xlog.Error("BinPath is not exists")
		return nil
	}
	objPath = fmt.Sprintf("%s/%s/", objPath, xtime.Now().PFormat("YmdHis"))
	xfile.MkdirAll(objPath)
	xlog.Info("in temp unrar:", objPath)
	if xfile.FileIsExist(objPath) {
		err := os.MkdirAll(objPath, 777)
		if err != nil {
			xlog.Error(err)
			return nil
		}
	}
	objFile := fmt.Sprintf("%s/data.%s", objPath, extName)
	err := ioutil.WriteFile(objFile, data, 0644)
	if err != nil {
		xlog.Error("write file:", objFile, ", error=", err)
		return nil
	}

	cmd := exec.Command(xfile.FormatPathSys(BinPath), "e", xfile.FormatPathSys(objFile), fmt.Sprintf("-o%s", xfile.FormatPathSys(objPath+"/data")), "-aoa")
	cmd.SysProcAttr = &syscall.SysProcAttr{HideWindow: true}
	out, err := cmd.Output()
	if err != nil {
		xlog.Error(err)
		return nil
	}
	out, err = ioutil.ReadAll(transform.NewReader(bytes.NewReader(out), simplifiedchinese.GBK.NewDecoder()))
	if err != nil {
		xlog.Error(err)
		return nil
	}
	outStr := string(out)
	xlog.Info("out=", outStr)
	if strings.Contains(outStr, "Everything is Ok") {
		xlog.Info("unrar success")
	} else {
		xlog.Error("unrar fail")
		return nil
	}
	filelist := xfile.GetFileList(objPath)
	if filelist == nil {
		xlog.Error("filelist is nil")
		return nil
	}
	return filelist
}
