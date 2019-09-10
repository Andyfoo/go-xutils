package xencode

import (
	"fmt"
	"html"
	"testing"
)

func TestEncode(t *testing.T) {
	gbkBytes := []byte{0xC4, 0xE3, 0xBA, 0xC3, 0xA3, 0xAC, 0xCA, 0xC0, 0xBD, 0xE7, 0xA3, 0xA1}
	gbkStr := string(gbkBytes)
	utfStr := "你好，世界！"
	fmt.Println("gbkStr", gbkStr)
	fmt.Println("GbkToUtf8Str", GbkToUtf8Str(gbkStr), len(GbkToUtf8Str(gbkStr)))
	fmt.Println("Utf8ToGbk", Utf8ToGbkStr(utfStr), len(Utf8ToGbkStr(utfStr)))
}
func TestBase64(t *testing.T) {

	fmt.Println("Base64Encode", Base64Encode([]byte("中国abc")))
	fmt.Println("Base64EncodeStr", Base64EncodeStr("中国abc"))
	//fmt.Println("Base64EncodeStrFile", Base64EncodeStrFile("E:/_tmp/test.data"))

	fmt.Println("Base64Decode", Base64Decode(Base64Encode([]byte("中国abc"))))
	fmt.Println("Base64Decode", Base64DecodeStr(Base64EncodeStr("中国abc")))
}

func TestUrlEncode(t *testing.T) {

	fmt.Println("UrlEncode", UrlEncodeStr("中国abc"))

	fmt.Println("UrlDecode", UrlDecodeStr(UrlEncodeStr("中国abc")))
}

func TestHexEncode(t *testing.T) {

	fmt.Println("HexEncodeStr", HexEncodeStr("中国abc"))

	fmt.Println("HexDecodeStr", HexDecodeStr(HexEncodeStr("中国abc")))
}

func TestHtmlEncode(t *testing.T) {

	fmt.Println("HtmlSpecialChars", HtmlSpecialChars("中国<>\"'&abc"))
	fmt.Println("EscapeString", html.EscapeString("中国<>\"'&abc"))
	fmt.Println("HtmlSpecialCharsDecode", HtmlSpecialCharsDecode(HtmlSpecialChars("中国<>\"'&abc")))
	fmt.Println("UnescapeString", html.UnescapeString(html.EscapeString("中国<>\"'&abc")))
}
