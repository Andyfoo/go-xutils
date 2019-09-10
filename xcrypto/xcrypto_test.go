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

package xcrypto

import (
	"fmt"
	"testing"

	"github.com/Andyfoo/go-xutils/xencode"
)

func TestAes(t *testing.T) {
	x := []byte("中文123abc")
	key := []byte("VK6Qbp4tVhDCiDK8")
	iv := []byte("0503080601040702")
	x1 := AesEncrypt(x, key, iv)
	fmt.Println("AesEncrypt hex", xencode.HexEncodeStr(string(x1)))
	fmt.Println("AesEncrypt base64", xencode.Base64EncodeStr(string(x1)))

	x2 := AesDecrypt(x1, key, iv)
	fmt.Println("AesDecrypt", string(x2))

	fmt.Println("AesEncryptHex", AesEncryptHex(string(x), string(key), string(iv)))
	fmt.Println("AesEncryptBase64", AesEncryptBase64(string(x), string(key), string(iv)))

	fmt.Println("AesDecryptHex", AesDecryptHex(AesEncryptHex(string(x), string(key), string(iv)), string(key), string(iv)))
	fmt.Println("AesDecryptBase64", AesDecryptBase64(AesEncryptBase64(string(x), string(key), string(iv)), string(key), string(iv)))

}
func TestDes(t *testing.T) {
	x := []byte("中文123abc")
	key := []byte("VK6Qbp42")
	iv := []byte("12345678")
	x1 := DesEncrypt(x, key, iv)
	fmt.Println("DesEncrypt hex", xencode.HexEncodeStr(string(x1)))
	fmt.Println("DesEncrypt base64", xencode.Base64EncodeStr(string(x1)))

	x2 := DesDecrypt(x1, key, iv)
	fmt.Println("DesDecrypt", string(x2))

	fmt.Println("DesEncryptHex", DesEncryptHex(string(x), string(key), string(iv)))
	fmt.Println("DesEncryptBase64", DesEncryptBase64(string(x), string(key), string(iv)))

	fmt.Println("DesDecryptHex", DesDecryptHex(DesEncryptHex(string(x), string(key), string(iv)), string(key), string(iv)))
	fmt.Println("DesDecryptBase64", DesDecryptBase64(DesEncryptBase64(string(x), string(key), string(iv)), string(key), string(iv)))

}

func TestTripleDes(t *testing.T) {
	x := []byte("中文123abc")
	key := []byte("123456789012345678901234")
	iv := []byte("12345678")
	x1 := TripleDesEncrypt(x, key, iv)
	fmt.Println("TripleDesEncrypt hex", xencode.HexEncodeStr(string(x1)))
	fmt.Println("TripleDesEncrypt base64", xencode.Base64EncodeStr(string(x1)))

	x2 := TripleDesDecrypt(x1, key, iv)
	fmt.Println("TripleDesDecrypt", string(x2))

	fmt.Println("TripleDesEncryptHex", TripleDesEncryptHex(string(x), string(key), string(iv)))
	fmt.Println("TripleDesEncryptBase64", TripleDesEncryptBase64(string(x), string(key), string(iv)))

	fmt.Println("TripleDesDecryptHex", TripleDesDecryptHex(TripleDesEncryptHex(string(x), string(key), string(iv)), string(key), string(iv)))
	fmt.Println("TripleDesDecryptBase64", TripleDesDecryptBase64(TripleDesEncryptBase64(string(x), string(key), string(iv)), string(key), string(iv)))

}
