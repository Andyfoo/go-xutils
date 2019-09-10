package xstring

import (
	"fmt"
	"testing"
)

func TestString1(t *testing.T) {
	fmt.Println("CutStr:", CutStr("我的abc", 1, "..."))
	fmt.Println("CutStr:", CutStr("我的abc", 3, "..."))
	fmt.Println("CutStr:", CutStr("我的abc", 33, "..."))

	fmt.Println("Substr:", Substr("asdfdasfsa", 0, 3))
	fmt.Println("SizeFormat:", SizeFormat(100000))
	fmt.Println("StrPad:", StrPad("a", 4, "0"))

	fmt.Printf("%e\n", 123400000.0)
	fmt.Printf("%E\n", 123400000.0)
	fmt.Printf("%.2f", 123400000.0)
}
func TestRandStr(t *testing.T) {

	fmt.Println("RandStr:", RandStr(6))
	fmt.Println("RandStrFor:", RandStrFor("num", 12))
	fmt.Println("RandStrFor:", RandStrFor("all", 12))
	fmt.Println("RandStrFor:", RandStrFor("lotter", 12))

}
func Test1NumberFormat(t *testing.T) {
	fmt.Println(NumberFormat("1234567898.55", -1, true))
	fmt.Println(NumberFormat("1234567898.55", -1, false))
	fmt.Println(NumberFormat("1234567898.5592", 2, false))

	fmt.Println(NumberFormat(1234567898.5592, 2, true))
}
func TestIsMobileUA(t *testing.T) {
	fmt.Println("######IsMobileUA:", IsMobileUA("Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/76.0.3809.132 Safari/537.36"))
	fmt.Println("######IsMobileUA:", IsMobileUA("Mozilla/5.0 (iPhone; CPU iPhone OS 12_0 like Mac OS X) AppleWebKit/605.1.15 (KHTML, like Gecko) Mobile/16A366 MicroMessenger/6.7.3(0x16070321) NetType/WIFI Language/zh_CN	"))
	fmt.Println("######IsMobileUA:", IsMobileUA("Mozilla/5.0 (Linux; U; Android 5.1.1; zh-CN; SM-J3109 Build/LMY47X) AppleWebKit/537.36 (KHTML, like Gecko) Version/4.0 Chrome/57.0.2987.108 UCBrowser/11.8.0.960 UWS/2.12.1.18 Mobile Safari/537.36 AliApp(TB/7.5.4) UCBS/2.11.1.1 WindVane/8.3.0 720X1280	"))

}
