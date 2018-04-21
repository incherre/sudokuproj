// Package sudokugen contains functions for generating completed sudoku boards.
// This package is for generating boards for a specific variant of sudoku.
package sudokugen

import (
	"fmt"
	"math/rand"
)

// NewBoard creates a new, empty sudoku board.
// The new board will be initialized with all zeros, representing blank spots.
// The function returns a pointer to the board.
func NewBoard() *[9][9]int {
	var temp [9][9]int
	for i := 0; i < 9; i++ {
		for j := 0; j < 9; j++ {
			temp[i][j] = 0
		}
	}
	return &temp
}

// PrintBoard displays a sudoku board.
// Spots containing zeros are printed as blank.
func PrintBoard(board *[9][9]int) {
	for i := 0; i < 9; i++ {
		for j := 0; j < 9; j++ {
			if board[i][j] > 0 {
				fmt.Printf("%d ", board[i][j])
			} else {
				fmt.Printf("  ")
			}
		}
		fmt.Printf("\n")
	}
}

// checkColumn checks to see if a particular column is valid so far.
// Any collisions present in the column will cause it to return false.
// Zeros are not considered when checking for collisions.
// This particular implementation has to check some parts of rows too, because of this specific variant of sudoku.
func checkColumn(board *[9][9]int, column int) bool {
	var contains [9]bool
	for i := 0; i < 9; i++ {
		contains[i] = false
	}
	var index, i int

	if column >= 0 && column < 3 {
		for i = 6; i < 9; i++ {
			index = board[i][column] - 1
			if index >= 0 && contains[index] {
				return false
			} else if index >= 0 {
				contains[index] = true
			}
		}

		for i = 3; i < 9; i++ {
			index = board[column+3][i] - 1
			if index >= 0 && contains[index] {
				return false
			} else if index >= 0 {
				contains[index] = true
			}
		}

		return true
	} else if column >= 3 && column < 6 {
		for i = 3; i < 9; i++ {
			index = board[i][column] - 1
			if index >= 0 && contains[index] {
				return false
			} else if index >= 0 {
				contains[index] = true
			}
		}

		for i = 6; i < 9; i++ {
			index = board[column-3][i] - 1
			if index >= 0 && contains[index] {
				return false
			} else if index >= 0 {
				contains[index] = true
			}
		}

		return true
	} else if column >= 6 && column < 9 {
		for i = 0; i < 9; i++ {
			index = board[i][column] - 1
			if index >= 0 && contains[index] {
				return false
			} else if index >= 0 {
				contains[index] = true
			}
		}

		return true
	}

	return false
}

// checkRow checks to see if a particular row is valid so far.
// Any collisions present in the row will cause it to return false.
// Zeros are not considered when checking for collisions.
// This particular implementation has to check some parts of columns too, because of this specific variant of sudoku.
func checkRow(board *[9][9]int, row int) bool {
	var contains [9]bool
	for i := 0; i < 9; i++ {
		contains[i] = false
	}
	var index, i int

	if row >= 0 && row < 3 {
		for i = 6; i < 9; i++ {
			index = board[row][i] - 1
			if index >= 0 && contains[index] {
				return false
			} else if index >= 0 {
				contains[index] = true
			}
		}

		for i = 3; i < 9; i++ {
			index = board[i][row+3] - 1
			if index >= 0 && contains[index] {
				return false
			} else if index >= 0 {
				contains[index] = true
			}
		}

		return true
	} else if row >= 3 && row < 6 {
		for i = 3; i < 9; i++ {
			index = board[row][i] - 1
			if index >= 0 && contains[index] {
				return false
			} else if index >= 0 {
				contains[index] = true
			}
		}

		for i = 6; i < 9; i++ {
			index = board[i][row-3] - 1
			if index >= 0 && contains[index] {
				return false
			} else if index >= 0 {
				contains[index] = true
			}
		}

		return true
	} else if row >= 6 && row < 9 {
		for i = 0; i < 9; i++ {
			index = board[row][i] - 1
			if index >= 0 && contains[index] {
				return false
			} else if index >= 0 {
				contains[index] = true
			}
		}

		return true
	}

	return false
}

// checkBox checks to see if a particular box is valid so far.
// The box to be checked is specified by the coordinates of its upper left member. (the member with the smallest coordinates)
// Any collisions present in the box will cause it to return false.
// Zeros are not considered when checking for collisions.
func checkBox(board *[9][9]int, row, column int) bool { 
	var contains [9]bool
	for i := 0; i < 9; i++ {
		contains[i] = false
	}
	var index int

	for i := row; i < row+3; i++ {
		for j := column; j < column+3; j++ {
			index = board[i][j] - 1
			if index >= 0 && contains[index] {
				return false
			} else if index >= 0 {
				contains[index] = true
			}
		}
	}

	return true
}

// IsValidBoard checks if the entire board is valid.
// Any collisions present will cause it to return false.
// This particular implementation checks only those parts required for this specific variant of sudoku.
func IsValidBoard(board *[9][9]int) bool {
	for i := 0; i < 9; i++ {
		if !checkColumn(board, i) {
			return false
		}
	}

	for i := 6; i < 9; i++ {
		if !checkRow(board, i) {
			return false
		}
	}

	for i := 0; i < 9; i = i + 3 {
		for j := 6 - i; j < 9; j = j + 3 {
			if !checkBox(board, i, j) {
				return false
			}
		}
	}

	return true
}

// getNextCords returns the row, column pair that should be checked next.
// This particular implementation visits only those spots to be filled for this specific variant of sudoku.
func getNextCords(row, column int) (int, int) {
	if column < 8 {
		return row, column + 1
	} else if row+1 < 3 {
		return row + 1, 6
	} else if row+1 < 6 {
		return row + 1, 3
	}
	return row + 1, 0
}

// IsFull checks that every required spot has been filled.
// This particular implementation checks only those parts required for this specific variant of sudoku.
func IsFull(board *[9][9]int) bool {
	for row, column := 0, 6; row < 9; row, column = getNextCords(row, column) {
		if board[row][column] == 0 {
			return false
		}
	}
	return true
}

// FindBoard finds a random valid completely filled sudoku board.
// This function recursively tries each numeral 1-9 in a random order. Pruning the search tree when one is invalid, stopping when the board has been filled.
// The provided board is modified in place.
func FindBoard(board *[9][9]int, row, column int) bool {
	valid := IsValidBoard(board)
	if !valid {
		return false
	} else if row > 8 {
		return valid
	}

	options := [9]int{1, 2, 3, 4, 5, 6, 7, 8, 9}
	var index, temp, i int
	for i = 0; i < 9; i++ {
		index = rand.Intn(9)
		temp = options[i]
		options[i] = options[index]
		options[index] = temp
	}

	tRow, tColumn := getNextCords(row, column)
	for i = 0; i < 9; i++ {
		board[row][column] = options[i]
		if FindBoard(board, tRow, tColumn) {
			return true
		}
	}

	board[row][column] = 0
	return false
}
