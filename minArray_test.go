package main

import (
	"fmt"
	"testing"
)

func Test_minArray(t *testing.T) {
	array1 := []int{3}
	min := minArray(array1)
	fmt.Println(min)
}
