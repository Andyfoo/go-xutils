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

package xencode

import (
	"bytes"
	"encoding/base64"
	"encoding/hex"
	"io/ioutil"
	"net/url"
	"strings"

	"github.com/Andyfoo/go-xutils/xlog"

	"golang.org/x/text/encoding/simplifiedchinese"
	"golang.org/x/text/transform"
)

func GbkToUtf8(srcBytes []byte) []byte {
	I := bytes.NewReader(srcBytes)
	O := transform.NewReader(I, simplifiedchinese.GBK.NewDecoder())
	data, err := ioutil.ReadAll(O)
	if err != nil {
		xlog.Error(err)
		return []byte{}
	}
	return data
}
func GbkToUtf8Str(src string) string {
	return string(GbkToUtf8([]byte(src)))
}
func Utf8ToGbk(srcBytes []byte) []byte {

	I := bytes.NewReader(srcBytes)
	O := transform.NewReader(I, simplifiedchinese.GBK.NewEncoder())
	data, err := ioutil.ReadAll(O)
	if err != nil {
		xlog.Error(err)
		return []byte{}
	}
	return data
}
func Utf8ToGbkStr(src string) string {
	return string(Utf8ToGbk([]byte(src)))
}

//将特殊字符转换为 HTML 实体,也可使用 html.EscapeString
func HtmlSpecialChars(str string) string {
	result := strings.Replace(str, "&", "&amp;", -1)
	result = strings.Replace(result, "<", "&lt;", -1)
	result = strings.Replace(result, ">", "&gt;", -1)
	result = strings.Replace(result, "\"", "&quot;", -1)
	result = strings.Replace(result, "'", "&#39;", -1)
	return result
}

//将特殊的 HTML 实体转换回普通字符，也可使用 html.UnescapeString
func HtmlSpecialCharsDecode(str string) string {
	result := strings.Replace(str, "&amp;", "&", -1)
	result = strings.Replace(result, "&lt;", "<", -1)
	result = strings.Replace(result, "&gt;", ">", -1)
	result = strings.Replace(result, "&quot;", "\"", -1)
	result = strings.Replace(result, "&#39;", "'", -1)
	return result
}

//计算Base64Encode
func Base64EncodeStr(src string) string {
	return string(Base64Encode([]byte(src)))
}
func Base64Encode(src []byte) []byte {
	encoding := base64.StdEncoding
	dst := make([]byte, encoding.EncodedLen(len(src)))
	base64.StdEncoding.Encode(dst, src)
	return dst
}
func Base64EncodeFile(file string) []byte {
	content, err := ioutil.ReadFile(file)
	if err != nil {
		xlog.Fatal(err)
		return []byte{}
	}
	return Base64Encode(content)
}
func Base64EncodeFileStr(file string) string {
	return string(Base64EncodeFile(file))
}

//计算Base64Decode
func Base64DecodeStr(src string) string {
	return string(Base64Decode([]byte(src)))
}

func Base64Decode(src []byte) []byte {
	encoding := base64.StdEncoding
	dst := make([]byte, encoding.DecodedLen(len(src)))
	n, err := base64.StdEncoding.Decode(dst, src)
	if err != nil {
		xlog.Fatal(err)
		return []byte{}
	}
	return dst[:n]
}

//计算UrlEncode
func UrlEncodeStr(buf string) string {
	return url.QueryEscape(buf)
}

//计算UrlDecode
func UrlDecodeStr(buf string) string {
	content, err := url.QueryUnescape(buf)
	if err != nil {
		xlog.Fatal(err)
		return ""
	}
	return string(content)
}

//计算HexEncode
func HexEncode(buf []byte) []byte {
	dst := make([]byte, hex.EncodedLen(len(buf)))
	n := hex.Encode(dst, buf)
	return dst[:n]
}
func HexEncodeStr(src string) string {
	return string(HexEncode([]byte(src)))
}

//计算HexDecode
func HexDecode(buf []byte) []byte {
	dst := make([]byte, hex.DecodedLen(len(buf)))
	n, err := hex.Decode(dst, buf)
	if err != nil {
		xlog.Fatal(err)
		return []byte{}
	}
	return dst[:n]
}
func HexDecodeStr(src string) string {
	return string(HexDecode([]byte(src)))
}
