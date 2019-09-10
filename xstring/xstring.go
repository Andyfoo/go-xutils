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

package xstring

import (
	"fmt"
	"math/rand"
	"strconv"
	"strings"
	"time"
)

//截取字符串，并追加后缀
func CutStr(str string, toCount int, more string) string {
	orgLen := len(str)
	if orgLen <= toCount {
		return str
	}
	return Substr(str, 0, toCount) + more
}

//start：正数 - 在字符串的指定位置开始,超出字符串长度强制把start变为字符串长度
//       负数 - 在从字符串结尾的指定位置开始
//       0 - 在字符串中的第一个字符处开始
//length:正数 - 从 start 参数所在的位置返回
//       负数 - 从字符串末端返回
func Substr(str string, start, length int) string {
	if length == 0 {
		return ""
	}
	rune_str := []rune(str)
	len_str := len(rune_str)

	if start < 0 {
		start = len_str + start
	}
	if start > len_str {
		start = len_str
	}
	end := start + length
	if end > len_str {
		end = len_str
	}
	if length < 0 {
		end = len_str + length
	}
	if start > end {
		start, end = end, start
	}
	return string(rune_str[start:end])
}

//格式化大小
func SizeFormat(size float64) string {
	units := []string{"Byte", "KB", "MB", "GB", "TB"}
	n := 0
	for size > 1024 {
		size /= 1024
		n += 1
	}

	return fmt.Sprintf("%.2f %s", size, units[n])
}

//前面补字符
func StrPad(str string, length int, schar string) string {
	want := length - len(str)
	if want > 0 {
		return strings.Repeat(schar, want) + str
	}
	return str
}

//重复字符 strings.Repeat

//随机字符串
func RandStr(length int) string {
	str := "ABCDEFGHJKLMNPQRSTUVWXYZabcdefghjkmnpqrstuvwxyz0123456789" //IiloO
	bytes := []byte(str)
	strLen := len(bytes)
	result := []byte{}
	//使用math rand
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	for i := 0; i < length; i++ {
		r.Seed(time.Now().UnixNano() + int64(i))
		result = append(result, bytes[r.Intn(strLen)])
	}

	//使用crypto/rand
	// for i := 0; i < length; i++ {
	// 	randNum, _ := rand.Int(rand.Reader, big.NewInt(strLen))
	// 	result = append(result, bytes[randNum.Int64()])
	// }
	return string(result)
}

//随机字符串
func RandStrFor(t string, length int) string {
	var str string
	if t == "str" {
		str = "ABCDEFGHJKLMNPQRSTUVWXYZabcdefghjkmnpqrstuvwxyz0123456789"
	} else if t == "all" {
		str = "ABCDEFGHJKLMNPQRSTUVWXYZabcdefghjkmnpqrstuvwxyz0123456789~!@#+-_=$%^&*:;<>,./?"
	} else if t == "num" {
		str = "0123456789"
	} else if t == "en" {
		str = "ABCDEFGHJKLMNPQRSTUVWXYZabcdefghjkmnpqrstuvwxyz"
	} else if t == "lower" {
		str = "abcdefghjkmnpqrstuvwxyz"
	} else if t == "upper" {
		str = "ABCDEFGHJKLMNPQRSTUVWXYZ"
	} else if t == "lower_str" {
		str = "abcdefghjkmnpqrstuvwxyz0123456789"
	} else if t == "upper_str" {
		str = "ABCDEFGHJKLMNPQRSTUVWXYZ0123456789"
	} else if len(t) > 0 {
		str = t
	} else {
		return ""
	}
	bytes := []byte(str)
	strLen := len(bytes)
	//xarray.ShuffleByte(bytes)
	result := []byte{}
	//使用math rand
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	for i := 0; i < length; i++ {
		r.Seed(time.Now().UnixNano() + int64(i))
		result = append(result, bytes[r.Intn(strLen)])
	}
	//使用crypto/rand
	// for i := 0; i < length; i++ {
	// 	randNum, _ := rand.Int(rand.Reader, big.NewInt(strLen))
	// 	result = append(result, bytes[randNum.Int64()])
	// }
	return string(result)
}

func NumberFormat(par1 interface{}, decimals int, dot bool) string {
	var decimalsFmt = "%." + strconv.Itoa(decimals) + "f"
	var str string
	switch v := par1.(type) {
	case string:
		n_val, err := strconv.ParseFloat(v, 64)
		if err != nil {
			return ""
		}
		if decimals >= 0 {
			str = fmt.Sprintf(decimalsFmt, n_val)
		} else {
			str = v
		}
		break
	case int:
		if decimals >= 0 {
			str = fmt.Sprintf(decimalsFmt, float64(v))
		} else {
			str = strconv.Itoa(v)
		}

		break
	case int64:
		if decimals >= 0 {
			str = fmt.Sprintf(decimalsFmt, float64(v))
		} else {
			str = strconv.FormatInt(v, 10)
		}

		break
	case float32:
		if decimals >= 0 {
			str = fmt.Sprintf(decimalsFmt, v)
		} else {
			str = strconv.FormatFloat(float64(v), 'f', -1, 64)
		}

		break
	case float64:
		if decimals >= 0 {
			str = fmt.Sprintf(decimalsFmt, v)
		} else {
			str = strconv.FormatFloat(v, 'f', -1, 64)
		}

		break
	default:
		return ""
		break
	}
	if dot {
		return NumberFormatDot(str)
	}
	return str
}

//格式数值为逗号分组    1,234,567,898.55
func NumberFormatDot(str string) string {
	length := len(str)
	if length < 4 {
		return str
	}
	arr := strings.Split(str, ".") //用小数点符号分割字符串,为数组接收
	length1 := len(arr[0])
	if length1 < 4 {
		return str
	}
	count := (length1 - 1) / 3
	for i := 0; i < count; i++ {
		arr[0] = arr[0][:length1-(i+1)*3] + "," + arr[0][length1-(i+1)*3:]
	}
	return strings.Join(arr, ".") //将一系列字符串连接为一个字符串，之间用sep来分隔。
}
