package zh_s2t

import (
	"fmt"
	"testing"
)

func Test1(t *testing.T) {
	fmt.Println(ToTrad("中华Aabc"))
	fmt.Println(ToSimp(ToTrad("中华Aabc")))
}
