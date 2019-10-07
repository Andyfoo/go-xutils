package xhttp

import (
	"fmt"
	"net/url"
	"testing"
)

func TestHttp(t *testing.T) {
	//testHttpGet()
	//testHttpPostStream()
	//testHttpPut()
	//testHttpPost()
	//testGetBytes()
	//testDownFile()
	//testHttpHead()
	testHttpPostMap()
}

func testHttpGet() {
	hUtil := NewDefaultHttpUtil()
	hUtil.UseCookie = true
	var a = hUtil.Get("https://getman.cn/echo")
	fmt.Println(len(a), a)
	var b = hUtil.Get("https://getman.cn/echo")
	fmt.Println(len(b), b)
	var c = hUtil.Get("https://getman.cn/echo")
	fmt.Println(len(c), b)
}
func testHttpPut() {
	hUtil := NewDefaultHttpUtil()
	var data = `{"aaa":111,"bbb":"汉字"}`
	reqParm := ReqParm{
		MediaType: "application/json",
		Charset:   "utf-8",
	}
	var a = hUtil.Req(Method_PUT, "https://getman.cn/echo", []byte(data), reqParm)

	fmt.Println(string(a))
}
func testHttpPostStream() {
	hUtil := NewDefaultHttpUtil()
	var data = `{"aaa":111,"bbb":222}`
	reqParm := ReqParm{
		MediaType: "application/json",
		Charset:   "utf-8",
	}
	var a = hUtil.Req(Method_POST, "https://getman.cn/echo", []byte(data), reqParm)

	fmt.Println(string(a))
}
func testHttpHead() {
	hUtil := NewDefaultHttpUtil()
	var data = `{"aaa":111,"bbb":222}`
	reqParm := ReqParm{
		MediaType: "application/json",
		Charset:   "utf-8",
	}
	var a = hUtil.Head("https://getman.cn/echo", []byte(data), reqParm)

	fmt.Println(a)
}
func testHttpPost() {
	hUtil := NewDefaultHttpUtil()
	var data = url.Values{}
	data.Set("aaaa", "1234")
	var a = hUtil.Post("https://getman.cn/echo", data)

	fmt.Println(a)
}
func testHttpPostMap() {
	hUtil := NewDefaultHttpUtil()
	var data = make(map[string]string)
	data["aaa"] = "weew"
	var a = hUtil.PostMap("https://getman.cn/echo", data)

	fmt.Println(a)
}

func testGetBytes() {
	hUtil := NewDefaultHttpUtil()
	var a = hUtil.GetBytes("https://getman.cn/echo")

	fmt.Println(a)
}
func testDownFile() {
	hUtil := NewDefaultHttpUtil()
	var a, filename = hUtil.DownFile("https://github.com/google/go-cmdtest/archive/v0.1.0.zip")

	fmt.Println(filename)
	fmt.Println(a)
}
