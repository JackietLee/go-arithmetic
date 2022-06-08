package main

import (
	"fmt"
	"testing"
)

type para struct {
	board  [][]int
	target int
}

func Test_findInArray(t *testing.T) {
	paras := []para{
		{
			board: [][]int{
				[]int{1, 2, 8, 9},
				[]int{2, 4, 9, 12},
				[]int{4, 7, 10, 13},
				[]int{6, 8, 11, 15},
			},
			target: 7,
		},
		{
			board: [][]int{
				[]int{1, 2, 8, 9},
				[]int{2, 4, 9, 12},
				[]int{4, 6, 10, 13},
				[]int{6, 8, 11, 15},
			},
			target: 7,
		},
	}
	for _, p := range paras {
		b, t := p.board, p.target
		array := findInArray(b, t)
		fmt.Println(array)
	}
}
