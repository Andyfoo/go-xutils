package xmath

import (
	"fmt"
	"testing"
)

func Test1(t *testing.T) {
	fmt.Println("RandInt64", RandInt64(1, 32))
	fmt.Println("RandInt", RandInt(1, 32))
}
