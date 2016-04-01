package sudoku

import (
	"reflect"
	"testing"
)

var (
	board = &Board{
		Cells: [9][9]int{
			{0, 6, 1, 0, 0, 7, 0, 0, 3},
			{0, 9, 2, 0, 0, 3, 0, 0, 0},
			{0, 0, 0, 0, 0, 0, 0, 0, 0},
			{0, 0, 8, 5, 3, 0, 0, 0, 0},
			{0, 0, 0, 0, 0, 0, 5, 0, 4},
			{5, 0, 0, 0, 0, 8, 0, 0, 0},
			{0, 4, 0, 0, 0, 0, 0, 0, 1},
			{0, 0, 0, 1, 6, 0, 8, 0, 0},
			{6, 0, 0, 0, 0, 0, 0, 0, 0},
		},
	}

	solution = &Board{
		Cells: [9][9]int{
			{4, 6, 1, 9, 8, 7, 2, 5, 3},
			{7, 9, 2, 4, 5, 3, 1, 6, 8},
			{3, 8, 5, 2, 1, 6, 4, 7, 9},
			{1, 2, 8, 5, 3, 4, 7, 9, 6},
			{9, 3, 6, 7, 2, 1, 5, 8, 4},
			{5, 7, 4, 6, 9, 8, 3, 1, 2},
			{8, 4, 9, 3, 7, 5, 6, 2, 1},
			{2, 5, 3, 1, 6, 9, 8, 4, 7},
			{6, 1, 7, 8, 4, 2, 9, 3, 5},
		},
	}
)

func copy(board *Board) *Board {
	return &Board{Cells: board.Cells}
}

func Test_Board_isValid(t *testing.T) {
	for i, check := range []struct {
		row, col, digit int
		expect          bool
	}{
		{1, 1, 1, false},
		{1, 1, 2, false},
		{1, 1, 3, false},
		{1, 1, 4, false},
		{1, 1, 5, true},
		{1, 1, 6, false},
		{1, 1, 7, true},
		{1, 1, 8, true},
		{1, 1, 9, false},
	} {
		if board.isValid(check.row, check.col, check.digit) != check.expect {
			t.Errorf("[%d] did not expect to find digit %d in column %d, row %d or in the corresponding 3x3 section",
				i, check.digit, check.col, check.row)
		}
	}
}

func Test_Board_findEmptyCell(t *testing.T) {
	for i, check := range []struct {
		board           *Board
		row, col        int
		expectEmptyCell bool
	}{
		{board, 0, 0, true},
		{solution, 0, 0, false},
	} {
		nextRow, nextCol, foundEmptyCell := check.board.findEmptyCell()

		if foundEmptyCell != check.expectEmptyCell || nextRow != check.row || nextCol != check.col {
			t.Errorf("[%d] expect to find empty cell in row %d and col %d, got row %d, col %d",
				i, check.row, check.col, nextRow, nextCol)
		}
	}
}

func Test_Board_Backtrack(t *testing.T) {
	original := copy(board)
	if !original.Backtrack() || !reflect.DeepEqual(original, solution) {
		t.Errorf("the board wasn't solved as expected")
	}

}
