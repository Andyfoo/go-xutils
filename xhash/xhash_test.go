package xhash

import (
	"fmt"
	"testing"

	"github.com/Andyfoo/go-xutils/xencode"
)

func TestMd5(t *testing.T) {
	fmt.Println(Md5([]byte("a国产do")))
	fmt.Println(Md5Str(xencode.Utf8ToGbkStr("a国产do")))
	//fmt.Println(Md5File("E:/_tmp/test.data"))
}
func TestSha1(t *testing.T) {
	fmt.Println(Sha1([]byte("a国产do")))
	//fmt.Println(Sha1File("E:/_tmp/test.data"))
}
func TestSha256(t *testing.T) {
	fmt.Println(Sha256([]byte("a国产do")))
	//fmt.Println(Sha256File("E:/_tmp/test.data"))
}
func TestCrc32(t *testing.T) {
	fmt.Println(Crc32([]byte("a国产do")))
	//fmt.Println(Crc32File("E:/_tmp/test.data"))
}
