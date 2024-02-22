package main

import "fmt"

const SIZE = 10

func putQueen(board []int, row int, pos *int) {
	if row == SIZE {
		*pos++
		printPosition(board)
	} else {
		for col := 0; col < SIZE; col++ {
			if isValid(board, row, col) {
				board[row] = col
				putQueen(board, row+1, pos)
			}
		}
	}
}

func isValid(board []int, row, col int) bool {
	for i := 0; i < row; i++ {
		if board[i] == col || board[i] == col-row+i || board[i] == col+row-i {
			return false
		}
	}
	return true
}

func printPosition(board []int) {
	for _, val := range board {
		fmt.Printf("%d", val)
	}
	fmt.Println()
}

func tenQueensPuzzle() int {
	board := make([]int, SIZE)
	pos := 0
	putQueen(board, 0, &pos)
	return pos
}

func main() {
	fmt.Printf("Number of solutions: %d\n", tenQueensPuzzle())
}
