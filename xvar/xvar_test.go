package xvar

import (
	"fmt"
	"reflect"
	"testing"
)

func TestVar1(t *testing.T) {

	fmt.Println("IntVal", IntVal(""))
	fmt.Println("IntVal", IntVal("3"))
	fmt.Println("IntVal", IntVal("d", 1))

	fmt.Println("Int64Val", Int64Val(""))
	fmt.Println("Int64Val", Int64Val("3"))
	fmt.Println("Int64Val", Int64Val("d", 1))

	fmt.Println("FloatVal", FloatVal(""))
	fmt.Println("FloatVal", FloatVal("3.1"))
	fmt.Println("FloatVal", FloatVal("d", 1))

	fmt.Println("Float64Val", Float64Val(""))
	fmt.Println("Float64Val", Float64Val("3.1"))
	fmt.Println("Float64Val", Float64Val("d", 1))

}
func TestToStr(t *testing.T) {
	fmt.Println("ToStr", ToStr(int16(122)))
	fmt.Println("ToStrArr", ToStrArr(1, 2, 3))
	fmt.Println("ToStr", ToStr([]int{1, 2, 3}))
	fmt.Println("ToStr", ToStr([]int64{1, 2, 3}))

	str := []string{"wd"}
	res_type := reflect.TypeOf(str)
	fmt.Println(res_type) //string
}
