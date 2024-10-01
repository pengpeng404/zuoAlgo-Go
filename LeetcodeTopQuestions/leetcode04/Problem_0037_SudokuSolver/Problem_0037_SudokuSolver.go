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
	solveSudoku(board)
	for _, row := range board {
		for _, cell := range row {
			fmt.Printf("%c ", cell)
		}
		fmt.Println()
	}
}

// https://leetcode.com/problems/sudoku-solver/

func solveSudoku(board [][]byte) {
	row := make([][]bool, 9)
	col := make([][]bool, 9)
	bucket := make([][]bool, 9)

	for i := 0; i < 9; i++ {
		row[i] = make([]bool, 10)
		col[i] = make([]bool, 10)
		bucket[i] = make([]bool, 10)
	}

	initMaps(row, col, bucket, board)
	process(0, 0, board, row, col, bucket)

}

func process(i, j int, board [][]byte, row, col, bucket [][]bool) bool {
	if i == 9 {
		return true
	}
	var nexti, nextj int
	if j == 8 {
		nexti = i + 1
		nextj = 0
	} else {
		nexti = i
		nextj = j + 1
	}
	if board[i][j] != '.' {
		return process(nexti, nextj, board, row, col, bucket)
	}
	bid := 3*(i/3) + (j / 3)
	for num := 1; num <= 9; num++ {
		if (!row[i][num]) && (!col[j][num]) && (!bucket[bid][num]) {
			row[i][num] = true
			col[j][num] = true
			bucket[bid][num] = true
			board[i][j] = byte(num + '0')
			if process(nexti, nextj, board, row, col, bucket) {
				return true
			}
			row[i][num] = false
			col[j][num] = false
			bucket[bid][num] = false
			board[i][j] = '.'
		}
	}
	return false
}

func initMaps(row, col, bucket [][]bool, board [][]byte) {
	for i := 0; i < 9; i++ {
		for j := 0; j < 9; j++ {
			if board[i][j] != '.' {
				bid := 3*(i/3) + (j / 3)
				num := int(board[i][j] - '0')
				row[i][num] = true
				col[j][num] = true
				bucket[bid][num] = true
			}
		}
	}
}
