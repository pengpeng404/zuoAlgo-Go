package main

import "fmt"

func main() {
	board := [][]byte{
		{'5', '3', '.', '.', '7', '.', '.', '.', '.'},
		{'6', '.', '.', '1', '9', '5', '.', '.', '.'},
		{'.', '9', '8', '.', '.', '.', '.', '6', '.'},
		{'8', '.', '.', '.', '6', '.', '.', '.', '3'},
		{'4', '.', '.', '8', '.', '3', '.', '.', '1'},
		{'7', '.', '.', '.', '2', '.', '.', '.', '6'},
		{'.', '6', '.', '.', '.', '.', '2', '8', '.'},
		{'.', '.', '.', '4', '1', '9', '.', '.', '5'},
		{'.', '.', '.', '.', '8', '.', '.', '7', '9'},
	}
	// true
	fmt.Println(isValidSudoku(board))
}

// https://leetcode.com/problems/valid-sudoku/

func isValidSudoku(board [][]byte) bool {
	row := make([][]bool, 9)
	col := make([][]bool, 9)
	bucket := make([][]bool, 9)

	for i := 0; i < 9; i++ {
		row[i] = make([]bool, 10)
		col[i] = make([]bool, 10)
		bucket[i] = make([]bool, 10)
	}

	for i := 0; i < 9; i++ {
		for j := 0; j < 9; j++ {
			if board[i][j] != '.' {
				num := board[i][j] - '0'
				bid := 3*(i/3) + (j / 3)

				if row[i][num] || col[j][num] || bucket[bid][num] {
					return false
				}

				row[i][num] = true
				col[j][num] = true
				bucket[bid][num] = true
			}
		}
	}

	return true
}
