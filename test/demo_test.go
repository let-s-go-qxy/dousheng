package test

import (
	"fmt"
	"testing"
)

func Sup1(num int) {
	num++
}

func Sup2(num *int) {
	*num++
}

func TestStruct(t *testing.T) {
	var num1 int
	num2 := new(int)
	Sup1(num1)
	Sup2(num2)
	fmt.Println(num1, *num2)
}

func appendArr(arr *[]int) {
	*arr = append(*arr, 1)
}
