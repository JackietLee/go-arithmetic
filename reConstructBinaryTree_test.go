package main

import (
	"fmt"
	"testing"
)

func Test_reConstructBinaryTree(t *testing.T) {
	//pre := []int{1,2,4,7,3,5,6,8}
	//in := []int{4,7,2,1,5,3,6,8}
	pre := []int{3, 9, 20, 15, 7}
	in := []int{9, 3, 15, 20, 7}
	root := reConstructBinaryTree(pre, in)
	printPre(root)
	fmt.Println()
	printIn(root)
}
