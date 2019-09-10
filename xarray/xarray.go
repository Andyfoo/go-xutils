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

package xarray

import (
	"bytes"
	"math/rand"
	"sort"
	"strconv"
	"strings"
	"time"

	"github.com/Andyfoo/go-xutils/xvar"
)

//拼接字符串
func Concat(strArr ...string) string {
	var buffer bytes.Buffer
	for _, v := range strArr {
		buffer.WriteString(v)
	}
	return buffer.String()
}

//拼接字符串数组，类型推断
func ConcatObj(objArr ...interface{}) string {
	var buffer bytes.Buffer
	for _, v := range objArr {
		buffer.WriteString(xvar.ToStr(v))
	}
	return buffer.String()
}

//int arr 转 string
func IntArr2Str(i []int) string {
	return JoinInt(i, ",")
}

//int arr 转 string
func JoinInt(i []int, sep string) string {
	s := make([]string, 0, len(i))
	for _, o := range i {
		s = append(s, strconv.Itoa(o))
	}
	return strings.Join(s, sep)
}

//int arr 转 string arr
func IntArr2StrArr(i []int) []string {
	s := make([]string, 0, len(i))
	for _, o := range i {
		s = append(s, strconv.Itoa(o))
	}
	return s
}

//int64 arr 转 string
func Int64Arr2Str(i64 []int64) string {
	return JoinInt64(i64, ",")
}

//int arr 转 string
func JoinInt64(i []int64, sep string) string {
	s := make([]string, 0, len(i))
	for _, o := range i {
		s = append(s, strconv.FormatInt(o, 10))
	}
	return strings.Join(s, sep)
}

//int64 arr 转 string arr
func Int64Arr2StrArr(i64 []int64) []string {
	s := make([]string, 0, len(i64))
	for _, o := range i64 {
		s = append(s, strconv.FormatInt(o, 10))
	}
	return s
}

//string arr 转 string
func JoinStr(s []string, sep string) string {
	return strings.Join(s, sep)
}
func JoinIntArgs(sep string, args ...int) string {
	return JoinInt(args, sep)
}
func JoinInt64Args(sep string, args ...int64) string {
	return JoinInt64(args, sep)
}
func JoinStrArgs(sep string, args ...string) string {
	return JoinStr(args, sep)
}

//string arr 转 string
func StrArr2Str(s []string) string {
	return `"` + strings.Join(s, `","`) + `"`
}

//判断元素是否包含
func InStrArray(s string, arr []string) bool {
	for _, val := range arr {
		if s == val {
			return true
		}
	}
	return false
}
func InIntArray(s int, arr []int) bool {
	for _, val := range arr {
		if s == val {
			return true
		}
	}
	return false
}
func InInt64Array(s int64, arr []int64) bool {
	for _, val := range arr {
		if s == val {
			return true
		}
	}
	return false
}

func SortIntAsc(arr []int) []int {
	sort.Ints(arr)
	return arr
}
func SortIntDesc(arr []int) []int {
	sort.Sort(sort.Reverse(sort.IntSlice(arr)))
	return arr
}

func SortStrAsc(arr []string) []string {
	sort.Strings(arr)
	return arr
}
func SortStrDesc(arr []string) []string {
	sort.Sort(sort.Reverse(sort.StringSlice(arr)))
	return arr
}

//随机打乱数组
func Shuffle(slice []interface{}) {
	r := rand.New(rand.NewSource(time.Now().Unix()))
	for len(slice) > 0 {
		n := len(slice)
		randIndex := r.Intn(n)
		slice[n-1], slice[randIndex] = slice[randIndex], slice[n-1]
		slice = slice[:n-1]
	}
}
func ShuffleStr(slice []string) {
	r := rand.New(rand.NewSource(time.Now().Unix()))
	for len(slice) > 0 {
		n := len(slice)
		randIndex := r.Intn(n)
		slice[n-1], slice[randIndex] = slice[randIndex], slice[n-1]
		slice = slice[:n-1]
	}
}
func ShuffleInt(slice []int) {
	r := rand.New(rand.NewSource(time.Now().Unix()))
	for len(slice) > 0 {
		n := len(slice)
		randIndex := r.Intn(n)
		slice[n-1], slice[randIndex] = slice[randIndex], slice[n-1]
		slice = slice[:n-1]
	}
}

func ShuffleInt64(slice []int64) {
	r := rand.New(rand.NewSource(time.Now().Unix()))
	for len(slice) > 0 {
		n := len(slice)
		randIndex := r.Intn(n)
		slice[n-1], slice[randIndex] = slice[randIndex], slice[n-1]
		slice = slice[:n-1]
	}
}

func ShuffleByte(slice []byte) {
	r := rand.New(rand.NewSource(time.Now().Unix()))
	for len(slice) > 0 {
		n := len(slice)
		randIndex := r.Intn(n)
		slice[n-1], slice[randIndex] = slice[randIndex], slice[n-1]
		slice = slice[:n-1]
	}
}
