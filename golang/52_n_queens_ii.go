package main

import "strings"

func totalNQueens(n int) int {
	ans := 0
	board := make([]string, n)
	for i := 0; i < n; i++ {
		board[i] = strings.Repeat(".", n)
	}
	var dfs func([]string, int)
	dfs = func(board []string, row int) {
		if row == len(board) {
			ans++
			return
		}
		for col := 0; col < len(board[row]); col++ {
			if !isQueenPositionValid(board, row, col) {
				continue
			}
			newRow := []byte(board[row])
			newRow[col] = 'Q'
			board[row] = string(newRow)
			dfs(board, row+1)
			newRow[col] = '.'
			board[row] = string(newRow)
		}
	}
	dfs(board, 0)
	return ans
}
