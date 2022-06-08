package main

import (
	"fmt"
	"testing"
)

func TestReplaceSpace(t *testing.T) {
	text1 := " java"
	space := replaceSpace(text1)
	fmt.Println(space)
}
