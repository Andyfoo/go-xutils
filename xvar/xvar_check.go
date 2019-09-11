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

package xvar

import (
	"regexp"
)

//判断手机号
func IsTel(str string) bool {
	pattern := regexp.MustCompile("^[0-9]{8,15}$")
	if pattern.MatchString(str) {
		return true
	} else {
		return false
	}
}

//标准手机号判断
func IsMobile(str string) bool {
	pattern := regexp.MustCompile("(?:^1[3456789]|^9[28])\\d{9}$")
	if pattern.MatchString(str) {
		return true
	} else {
		return false
	}
}

//判断是否数字
func IsNumber(str string) bool {
	pattern := regexp.MustCompile("^[0-9]+$")
	if pattern.MatchString(str) {
		return true
	} else {
		return false
	}
}

//是否邮箱
func IsEmail(b []byte) bool {
	var emailPattern = regexp.MustCompile("[\\w!#$%&'*+/=?^_`{|}~-]+(?:\\.[\\w!#$%&'*+/=?^_`{|}~-]+)*@(?:[\\w](?:[\\w-]*[\\w])?\\.)+[a-zA-Z0-9](?:[\\w-]*[\\w])?")
	return emailPattern.Match(b)
}

//正则校验
func RegMatch(str string, reg string) bool {
	r, err := regexp.Compile(reg)
	if err != nil {
		return false
	}
	return r.MatchString(str)
}

func IsEmpty(str string) bool {
	return len(str) == 0
}
func IsNotEmpty(str string) bool {
	return len(str) > 0
}
func IsAllEmpty(strs ...string) bool {
	for _, v := range strs {
		if !IsEmpty(v) {
			return false
		}
	}
	return true
}
func IsBlank(str string) bool {
	return len(str) == 0 || ContainsOnlyWhitespaces(str)
}
func IsNotBlank(str string) bool {
	return len(str) > 0 && !ContainsOnlyWhitespaces(str)
}
func ContainsOnlyWhitespaces(str string) bool {
	strBytes := []byte(str)
	for _, v := range strBytes {
		if !IsWhitespace(v) {
			return false
		}
	}
	return true
}

func IsWhitespace(c byte) bool {
	return c <= ' '
}
