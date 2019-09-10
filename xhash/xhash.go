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

package xhash

import (
	"crypto/md5"
	"crypto/sha1"
	"crypto/sha256"
	"crypto/sha512"
	"fmt"
	"hash/crc32"
	"io"
	"os"

	"github.com/Andyfoo/go-xutils/xlog"
)

//计算md5
func Md5(buf []byte) []byte {
	hash := md5.New()
	hash.Write(buf)
	return hash.Sum(nil)
}
func Md5Str(src string) string {
	return fmt.Sprintf("%x", Md5([]byte(src)))
}

func Md5File(file string) string {
	f, err := os.Open(file)
	if err != nil {
		xlog.Fatal(err)
	}
	defer f.Close()

	h := md5.New()
	if _, err := io.Copy(h, f); err != nil {
		xlog.Fatal(err)
	}

	return fmt.Sprintf("%x", h.Sum(nil))
}

//计算Sha1
func Sha1(buf []byte) []byte {
	hash := sha1.New()
	hash.Write(buf)
	return hash.Sum(nil)
}
func Sha1Str(src string) string {
	return fmt.Sprintf("%x", Sha1([]byte(src)))
}

func Sha1File(file string) string {
	f, err := os.Open(file)
	if err != nil {
		xlog.Fatal(err)
	}
	defer f.Close()

	h := sha1.New()
	if _, err := io.Copy(h, f); err != nil {
		xlog.Fatal(err)
	}

	return fmt.Sprintf("%x", h.Sum(nil))
}

//计算Sha256
func Sha256(buf []byte) []byte {
	hash := sha256.New()
	hash.Write(buf)
	return hash.Sum(nil)
}
func Sha256Str(src string) string {
	return fmt.Sprintf("%x", Sha256([]byte(src)))
}

func Sha256File(file string) string {
	f, err := os.Open(file)
	if err != nil {
		xlog.Fatal(err)
	}
	defer f.Close()

	h := sha256.New()
	if _, err := io.Copy(h, f); err != nil {
		xlog.Fatal(err)
	}

	return fmt.Sprintf("%x", h.Sum(nil))
}

//计算Sha512
func Sha512(buf []byte) []byte {
	hash := sha512.New()
	hash.Write(buf)
	return hash.Sum(nil)
}
func Sha512Str(src string) string {
	return fmt.Sprintf("%x", Sha512([]byte(src)))
}

func Sha512File(file string) string {
	f, err := os.Open(file)
	if err != nil {
		xlog.Fatal(err)
	}
	defer f.Close()

	h := sha512.New()
	if _, err := io.Copy(h, f); err != nil {
		xlog.Fatal(err)
	}

	return fmt.Sprintf("%x", h.Sum(nil))
}

//计算Crc32
func Crc32(buf []byte) []byte {
	hash := crc32.NewIEEE()
	hash.Write(buf)
	return hash.Sum(nil)
}
func Crc32Str(src string) string {
	return fmt.Sprintf("%x", Crc32([]byte(src)))
}

func Crc32File(file string) string {
	f, err := os.Open(file)
	if err != nil {
		xlog.Fatal(err)
	}
	defer f.Close()

	h := crc32.NewIEEE()
	if _, err := io.Copy(h, f); err != nil {
		xlog.Fatal(err)
	}

	return fmt.Sprintf("%x", h.Sum(nil))
}
