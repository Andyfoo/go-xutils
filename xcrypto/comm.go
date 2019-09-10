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
	"bytes"

	"github.com/Andyfoo/go-xutils/xlog"
)

func Pkcs5Padding(src []byte, blocksize int) []byte {
	padnum := blocksize - len(src)%blocksize
	pad := bytes.Repeat([]byte{byte(padnum)}, padnum)
	return append(src, pad...)
}

func UnPkcs5Padding(src []byte) []byte {
	n := len(src)
	unpadnum := int(src[n-1])
	if unpadnum > n {
		xlog.Error("UnPkcs5Padding error")
		return []byte{}
	}
	return src[:n-unpadnum]
}
