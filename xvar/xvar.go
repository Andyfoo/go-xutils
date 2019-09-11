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
	"bytes"
	"strconv"
	"strings"
)

//任意类型转字符串
func ToStr(obj interface{}) string {

	switch v1 := obj.(type) {
	case string:
		return v1
	case int:
		return strconv.Itoa(v1)
	case int8:
		return strconv.Itoa(int(v1))
	case int16:
		return strconv.Itoa(int(v1))
	case int32:
		return strconv.Itoa(int(v1))
	case int64:
		return strconv.FormatInt(v1, 10)
	case uint:
		return strconv.Itoa(int(v1))
	case uint8:
		return strconv.Itoa(int(v1))
	case uint16:
		return strconv.Itoa(int(v1))
	case uint32:
		return strconv.Itoa(int(v1))
	case uint64:
		return strconv.FormatInt(int64(v1), 10)
	case float32:
		return strconv.FormatFloat(float64(v1), 'f', -1, 64)
	case float64:
		return strconv.FormatFloat(v1, 'f', -1, 64)
	case []string:
		var buffer bytes.Buffer
		for _, v := range v1 {
			buffer.WriteString(v)
		}
		return buffer.String()
	case []int:
		var buffer bytes.Buffer
		for _, v := range v1 {
			buffer.WriteString(strconv.Itoa(v))
		}
		return buffer.String()
	case []int8:
		var buffer bytes.Buffer
		for _, v := range v1 {
			buffer.WriteString(strconv.Itoa(int(v)))
		}
		return buffer.String()
	case []int16:
		var buffer bytes.Buffer
		for _, v := range v1 {
			buffer.WriteString(strconv.Itoa(int(v)))
		}
		return buffer.String()
	case []int32:
		var buffer bytes.Buffer
		for _, v := range v1 {
			buffer.WriteString(strconv.Itoa(int(v)))
		}
		return buffer.String()
	case []int64:
		var buffer bytes.Buffer
		for _, v := range v1 {
			buffer.WriteString(strconv.FormatInt(v, 10))
		}
		return buffer.String()
	case []uint:
		var buffer bytes.Buffer
		for _, v := range v1 {
			buffer.WriteString(strconv.Itoa(int(v)))
		}
		return buffer.String()
	case []uint8:
		var buffer bytes.Buffer
		for _, v := range v1 {
			buffer.WriteString(strconv.Itoa(int(v)))
		}
		return buffer.String()
	case []uint16:
		var buffer bytes.Buffer
		for _, v := range v1 {
			buffer.WriteString(strconv.Itoa(int(v)))
		}
		return buffer.String()
	case []uint32:
		var buffer bytes.Buffer
		for _, v := range v1 {
			buffer.WriteString(strconv.Itoa(int(v)))
		}
		return buffer.String()
	case []uint64:
		var buffer bytes.Buffer
		for _, v := range v1 {
			buffer.WriteString(strconv.FormatInt(int64(v), 10))
		}
		return buffer.String()
	case []float32:
		var buffer bytes.Buffer
		for _, v := range v1 {
			buffer.WriteString(strconv.FormatFloat(float64(v), 'f', -1, 64))
		}
		return buffer.String()
	case []float64:
		var buffer bytes.Buffer
		for _, v := range v1 {
			buffer.WriteString(strconv.FormatFloat(v, 'f', -1, 64))
		}
		return buffer.String()
	}
	return ""
}

//任意类型多参数转字符串数组
func ToStrArr(objArr ...interface{}) []string {
	var arr = make([]string, len(objArr))
	for k, v := range objArr {
		arr[k] = ToStr(v)
	}
	return arr
}

//任意类型数组转字符串数组
func ObjArrToStrArr(objArr []interface{}) []string {
	var arr = make([]string, len(objArr))
	for k, v := range objArr {
		arr[k] = ToStr(v)
	}
	return arr
}

func IntToStr(v int) string {
	return strconv.Itoa(v)
}
func Int64ToStr(v int64) string {
	return strconv.FormatInt(int64(v), 10)
}
func FloatToStr(v float32) string {
	return strconv.Itoa(v)
}
func Float64ToStr(v float64) string {
	return strconv.FormatInt(int64(v), 10)
}

//string转int
func IntVal(s string, _def ...int) int {
	var def int
	if len(_def) > 0 {
		def = _def[0]
	}
	if len(s) == 0 {
		return def
	}
	v, err := strconv.Atoi(s)
	if err != nil {
		return def
	}
	return v
}

//string转int64
func Int64Val(s string, _def ...int64) int64 {
	var def int64
	if len(_def) > 0 {
		def = _def[0]
	}
	if len(s) == 0 {
		return def
	}
	v, err := strconv.ParseInt(s, 10, 64)
	if err != nil {
		return def
	}
	return v
}

//string转float
func FloatVal(s string, _def ...float32) float32 {
	var def float32
	if len(_def) > 0 {
		def = _def[0]
	}
	if len(s) == 0 {
		return def
	}
	v, err := strconv.ParseFloat(s, 32)
	if err != nil {
		return def
	}
	return float32(v)
}

//string转float64
func Float64Val(s string, _def ...float64) float64 {
	var def float64
	if len(_def) > 0 {
		def = _def[0]
	}
	if len(s) == 0 {
		return def
	}
	v, err := strconv.ParseFloat(s, 64)
	if err != nil {
		return def
	}
	return v
}

//处理sql字符
func FilterSql(str string) (result string) {
	result = strings.Replace(str, "'", "", -1)
	result = strings.Replace(result, "\"", "", -1)
	result = strings.Replace(result, ";", "", -1)
	result = strings.Replace(result, "%", "", -1)
	result = strings.Replace(result, "=", "", -1)
	result = strings.Replace(result, "*", "", -1)
	return result
}
