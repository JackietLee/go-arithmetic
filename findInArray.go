package main

func findInArray(board [][]int, target int) bool {
	rlen := len(board)
	clen := len(board[0])

	for r, c := 0, clen-1; r < rlen && c >= 0; {
		if board[r][c] == target {
			return true
		}
		if board[r][c] > target {
			c--
		} else {
			r++
		}
	}
	return false
}
