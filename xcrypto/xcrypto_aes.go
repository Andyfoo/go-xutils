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
	"crypto/aes"
	"crypto/cipher"

	"github.com/Andyfoo/go-xutils/xencode"

	"github.com/Andyfoo/go-xutils/xlog"
)

func AesEncrypt(src []byte, key []byte, iv []byte) []byte {
	if len(key) != 16 {
		xlog.Error("key size must is 16")
		return []byte{}
	}
	if len(iv) != 16 {
		xlog.Error("iv size must is 16")
		return []byte{}
	}
	block, err := aes.NewCipher(key)
	if err != nil {
		xlog.Error(err)
		return []byte{}
	}
	bs := block.BlockSize()
	data := Pkcs5Padding(src, bs)
	if len(data)%bs != 0 {
		xlog.Error("need a multiple of the blocksize")
		return []byte{}
	}
	blockmode := cipher.NewCBCEncrypter(block, iv)
	dst := make([]byte, len(data))
	blockmode.CryptBlocks(dst, data)
	return dst
}
func AesEncryptHex(src string, key string, iv string) string {
	return xencode.HexEncodeStr(string(AesEncrypt([]byte(src), []byte(key), []byte(iv))))
}
func AesEncryptBase64(src string, key string, iv string) string {
	return xencode.Base64EncodeStr(string(AesEncrypt([]byte(src), []byte(key), []byte(iv))))
}

func AesDecrypt(src []byte, key []byte, iv []byte) []byte {
	if len(key) != 16 {
		xlog.Error("key size must is 16")
		return []byte{}
	}
	if len(iv) != 16 {
		xlog.Error("iv size must is 16")
		return []byte{}
	}
	block, err := aes.NewCipher(key)
	if err != nil {
		xlog.Error(err)
		return []byte{}
	}
	blockmode := cipher.NewCBCDecrypter(block, iv)

	dst := make([]byte, len(src))
	blockmode.CryptBlocks(dst, src)
	return UnPkcs5Padding(dst)
}
func AesDecryptHex(src string, key string, iv string) string {
	return string(AesDecrypt(xencode.HexDecode([]byte(src)), []byte(key), []byte(iv)))
}

func AesDecryptBase64(src string, key string, iv string) string {
	return string(AesDecrypt(xencode.Base64Decode([]byte(src)), []byte(key), []byte(iv)))
}
