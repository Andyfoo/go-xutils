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

package xregex

import (
	"regexp"

	"github.com/Andyfoo/go-xutils/xlog"
)

//https://gamelife1314.github.io/2018/01/01/golang%E6%A0%87%E5%87%86%E5%BA%93%E4%B9%8Bregexp/

func GetMatchStr(str, regex string) string {
	arr := GetMatchArray(str, regex)
	if arr == nil || len(arr) == 0 {
		return ""
	}
	return arr[0]
}

func GetMatchArray(str, regex string) []string {
	reg, err := regexp.Compile(regex)
	if err != nil {
		xlog.Error(err)
		return nil
	}
	arr := reg.FindStringSubmatch(str)
	if arr == nil || len(arr) == 0 {
		return nil
	}
	resultArr := make([]string, 0)
	for k, v := range arr {
		if k == 0 {
			continue
		}
		resultArr = append(resultArr, v)
	}
	return resultArr
}

func GetMatchList(str, regex string) [][]string {
	reg, err := regexp.Compile(regex)
	if err != nil {
		xlog.Error(err)
		return nil
	}
	return reg.FindAllStringSubmatch(str, -1)
}

func ReplaceAll(srcStr, regex, repStr string) string {
	reg, err := regexp.Compile(regex)
	if err != nil {
		xlog.Error(err)
		return srcStr
	}
	return reg.ReplaceAllString(srcStr, repStr)
}

func ReplaceAllFunc(srcStr string, regex string, repFun func(string) string) string {
	reg, err := regexp.Compile(regex)
	if err != nil {
		xlog.Error(err)
		return srcStr
	}
	return reg.ReplaceAllStringFunc(srcStr, repFun)
}

func ReplaceFirst(srcStr, regex, repStr string) string {
	reg, err := regexp.Compile(regex)
	if err != nil {
		xlog.Error(err)
		return srcStr
	}
	isRep := false
	return reg.ReplaceAllStringFunc(srcStr, func(str string) string {
		if !isRep {
			isRep = true
			return repStr
		}
		return str
	})
}
