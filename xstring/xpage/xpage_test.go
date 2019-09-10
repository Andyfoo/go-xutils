package xpage

import (
	"fmt"
	"testing"
)

func TestString1(t *testing.T) {
	fmt.Println("Page:", Page(1, 10, 101, "a.do"))
	fmt.Println("Page:", Page(1, 10, 101, "a.do", "aa=bb"))

	parms := make(map[string]string)
	parms["aa"] = "11"
	parms["bb"] = "22"
	parms["cc"] = "33"
	fmt.Println("Page:", Page(1, 10, 101, "a.do", parms))
}
