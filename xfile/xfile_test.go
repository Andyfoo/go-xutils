package xfile

import (
	"fmt"
	"testing"
)

func TestFile1(t *testing.T) {
	fmt.Println(BaseName("aaa/sdafsa\\adsfsa//sdaf/abc.do"))
	fmt.Println(FormatPath("aaa/sdafsa\\adsfsa//sdaf/"))
}
func TestReadFile(t *testing.T) {
	data := ReadFile("e:/a.txt")
	fmt.Println(len(data), data)
	fmt.Println(len(string(data)), string(data))
	fmt.Println(len([]byte(string(data))), []byte(string(data)))
}
func TestWriteFilee(t *testing.T) {
	WriteFile("e:/b.txt", []byte("a1"))
	WriteFile("e:/c.txt", []byte("a1"), "a+")
}
