package main

import (
	"fmt"
	"math/rand"
	"time"
	"github.com/incherre/sudokuproj/sudokugen"
)

func main() {
	rand.Seed(time.Now().UTC().UnixNano())

	board := sudokugen.NewBoard()
	sudokugen.FindBoard(board, 0, 6)
	sudokugen.PrintBoard(board)

	var input string
	fmt.Scanln(&input)
}