package xregex

import (
	"fmt"
	"testing"
)

func TestGetMatchStr(t *testing.T) {
	fmt.Println(GetMatchStr("aa22adfaa", `aa(\d+)`))
	fmt.Println(GetMatchStr("aa22adfaa", `(?i)(AA)`))

}

func TestGetMatchArray(t *testing.T) {
	fmt.Println(GetMatchArray("aa22adfaa33", `(\d+)`))

	a2 := GetMatchArray("星际争霸II – GSL第二季", `([^\-]+) [\-–]+ (.*)`)
	fmt.Println(len(a2), a2)
}
func TestGetMatchList(t *testing.T) {
	fmt.Println(GetMatchList("aa22adfaa33####ads34asdff", `(\d+)([a-z]+)`))

}
func TestReplaceAll(t *testing.T) {
	fmt.Println(ReplaceAll("aa22adfaa33####ads34asdff", `(\d+)([a-z]+)`, "***"))

}
func TestReplaceAllFunc(t *testing.T) {
	fmt.Println(ReplaceAllFunc("aa22adfaa33####ads34asdff", `(\d+)([a-z]+)`, func(str string) string {
		return str + "***"
	}))

}

func TestReplaceFirst(t *testing.T) {
	fmt.Println(ReplaceFirst("aa22adfaa33####ads34asdff", `(\d+)`, "***"))

}
