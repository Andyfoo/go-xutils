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

	fmt.Println("IntToStr", IntToStr(32))
	fmt.Println("Int64ToStr", Int64ToStr(int64(33)))
	fmt.Println("FloatToStr", FloatToStr(32.2))
	fmt.Println("Float64ToStr", Float64ToStr(float64(64.32)))
}

func TestIsEmpty(t *testing.T) {
	fmt.Println("IsEmpty", IsEmpty("a"))
	fmt.Println("IsEmpty", IsEmpty(""))
	fmt.Println("IsNotEmpty", IsNotEmpty("a"))
	fmt.Println("IsNotEmpty", IsNotEmpty(""))
	fmt.Println("IsBlank", IsBlank("a a"))
	fmt.Println("IsBlank", IsBlank("   "))
	fmt.Println("IsNotBlank", IsNotBlank("a a"))
	fmt.Println("IsNotBlank", IsNotBlank("   "))
	fmt.Println("ContainsOnlyWhitespaces", ContainsOnlyWhitespaces("   "))
	fmt.Println("IsAllEmpty", IsAllEmpty("", ""))

}
