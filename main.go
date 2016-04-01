package main

import (
	"fmt"
	"os"
	"time"

	"github.com/ndeslandes/sudoku-solver-golang/sudoku"
)

func main() {
	board, err := sudoku.GetBoardFrom(os.Stdin)
	if err != nil {
		fmt.Printf("Unable to process the input: %s\n", err)
		os.Exit(1)
	}

	start := time.Now()
	if board.Backtrack() {
		fmt.Println("The Sudoku was solved successfully:")
		fmt.Println(board)
	} else {
		fmt.Printf("The Sudoku can't be solved.")
	}
	elapsed := time.Since(start)
	fmt.Printf("Backtrack took %s\n", elapsed)
}
