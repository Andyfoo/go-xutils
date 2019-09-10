package xarray

import (
	"fmt"
	"testing"
)

func TestIntArr2Str(t *testing.T) {
	i := []int{1, 2, 3, 4, 5}
	rs := IntArr2Str(i)
	fmt.Println("IntArr2Str", rs)
	rs2 := IntArr2StrArr(i)
	fmt.Println("IntArr2StrArr", rs2)
}
func TestInt64Arr2Str(t *testing.T) {
	i := []int64{11, 12, 13, 14}
	rs := Int64Arr2Str(i)
	fmt.Println("Int64Arr2Str", rs)
	rs2 := Int64Arr2StrArr(i)
	fmt.Println("Int64Arr2StrArr", rs2)
}
func TestStrArr2Str(t *testing.T) {
	i := []string{"a", "b", "c"}
	rs := StrArr2Str(i)
	fmt.Println("StrArr2Str", rs)
}
func TestInStrArray(t *testing.T) {
	bool := InStrArray("1", []string{"0", "1", "12"})
	fmt.Println("InStrArray", bool)
	bool = InIntArray(22, []int{1, 2, 3})
	fmt.Println("InIntArray", bool)
}
func TestSortInt(t *testing.T) {
	arr := []int{5, 2, 3}
	fmt.Println("SortIntAsc", SortIntAsc(arr))
	fmt.Println("SortIntDesc", SortIntDesc(arr))

	arr2 := []string{"a", "c", "b"}
	fmt.Println("SortStrAsc", SortStrAsc(arr2))
	fmt.Println("SortStrDesc", SortStrDesc(arr2))
}

func Test1Join(t *testing.T) {
	var arr = []string{"0", "1", "12"}
	fmt.Println("JoinStr", JoinStr(arr, ";"))
	fmt.Println("JoinStrArgs", JoinStrArgs(";", "0", "1", "12"))

}

func Test1Concat(t *testing.T) {
	fmt.Println("Concat", Concat("a"))
	fmt.Println("Concat", Concat("a", "b"))
	var arr1 = []string{"1", "2", "3"}
	fmt.Println("ConcatObj", ConcatObj(arr1))

	fmt.Println("ConcatObj", ConcatObj("a", "b", 2, 3.2, arr1, ",", []int{1, 2, 3}))
}
