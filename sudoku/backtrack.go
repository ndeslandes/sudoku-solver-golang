package sudoku

// Backtrack tries to find a valid solution for a given Board.
// A returned Boolean indicates if the board was solved successfully.
func (board *Board) Backtrack() bool {
	row, column, hasEmptyCell := board.findEmptyCell()
	if !hasEmptyCell {
		return true
	}
	for digit := 1; digit <= N; digit++ {
		if board.isValid(row, column, digit) {
			board.Cells[row][column] = digit
			if board.Backtrack() {
				return true
			}
			board.Cells[row][column] = 0
		}
	}
	return false
}

func (board *Board) findEmptyCell() (int, int, bool) {
	for row := 0; row < N; row++ {
		for column := 0; column < N; column++ {
			if board.Cells[row][column] == 0 {
				return row, column, true
			}
		}
	}
	return 0, 0, false
}

func (board *Board) isValid(row int, column int, number int) bool {
	sectionRowStart := row / 3 * 3
	sectionColumnStart := column / 3 * 3
	for i := 0; i < N; i++ {
		if board.Cells[row][i] == number || board.Cells[i][column] == number ||
			board.Cells[sectionRowStart+i/3][sectionColumnStart+i%3] == number {
			return false
		}
	}
	return true
}
