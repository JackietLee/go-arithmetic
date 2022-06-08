package main

import (
	"fmt"
	"testing"
)

func Test_numWays(t *testing.T) {
	ways, _ := numWays(7)
	fmt.Println(ways)
}
