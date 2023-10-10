package main

import "strings"

func solveNQueens(n int) [][]string {
	ans := [][]string{}
	board := make([]string, n)
	for i := 0; i < n; i++ {
		board[i] = strings.Repeat(".", n)
	}
	var dfs func([]string, int)
	dfs = func(board []string, row int) {
		if row == len(board) {
			ans = append(ans, append([]string{}, board...))
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

func isQueenPositionValid(board []string, row, col int) bool {
	// 检查列
	for i := 0; i < len(board); i++ {
		if board[i][col] == 'Q' {
			return false
		}
	}
	// 检查右上方
	for i, j := row-1, col+1; i >= 0 && j < len(board); i, j = i-1, j+1 {
		if board[i][j] == 'Q' {
			return false
		}
	}
	// 检查左上方
	for i, j := row-1, col-1; i >= 0 && j >= 0; i, j = i-1, j-1 {
		if board[i][j] == 'Q' {
			return false
		}
	}
	return true
}
