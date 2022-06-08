package main

import (
	"errors"
)

func numWays(n int) (int, error) {
	if n < 0 {
		return 0, errors.New("不能小于0")
	} else if n <= 2 {
		return n, nil
	}
	first := 1
	second := 2
	tmp := 0
	for i := 0; i < n-2; i++ {
		tmp = second
		second += first
		first = tmp
	}
	return second, nil
}
