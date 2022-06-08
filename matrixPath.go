package main

func exist(board [][]byte, word string) bool {
	wordByte := []byte(word)
	for i := 0; i < len(board); i++ {
		for j := 0; j < len(board[0]); j++ {
			if dfs(board, wordByte, i, j, 0) {
				return true
			}
		}
	}
	return false
}

func dfs(board [][]byte, word []byte, i, j, k int) bool {
	if i < 0 || i > len(board)-1 || j < 0 || j > len(board[0])-1 || board[i][j] != word[k] {
		return false
	}
	if k == len(word)-1 {
		return true
	}

	temp := board[i][j]
	board[i][j] = '/'
	b := dfs(board, word, i+1, j, k+1) || dfs(board, word, i-1, j, k+1) || dfs(board, word, i, j-1, k+1) || dfs(board, word, i, j+1, k+1)
	board[i][j] = temp

	return b
}
