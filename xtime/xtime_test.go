package xtime

import (
	"fmt"
	"testing"
	"time"
)

func Test1(t *testing.T) {
	fmt.Println(time.Now().Format("2006-01-02 15:04:05"))

	fmt.Println(Now().PFormat("Y-m-d H:i:s"))
	fmt.Println(Now().Unix())
	fmt.Println(Now().UnixMilli())
	fmt.Println(Now().UnixNano())
	fmt.Println(Now().Weekday())
	fmt.Println(Now().WeekdayStr(1))
	fmt.Println(Str2Time("2018-04-23 23:11:23", "Y-m-d H:i:s").PFormat("Y-m-d H:i:s"))
}
