package sudoku

import (
	"bufio"
	"fmt"
	"io"
	"strconv"
	"strings"
)

// N defines the Sudoku's board size
const N = 9

// Board represents the state of a Sudoku.
type Board struct {
	Cells [N][N]int
}

// GetBoardFrom takes an io.Reader and returns a Board with the input data.
// Invalid or malformed inputs will be rejected.
func GetBoardFrom(source io.Reader) (*Board, error) {
	scanner := bufio.NewScanner(source)
	board := new(Board)
	row := 0

	for ; scanner.Scan() && row < N; row++ {
		rowList := strings.Split(strings.TrimSpace(scanner.Text()), " ")
		if len(rowList) != N {
			return nil, fmt.Errorf("row %d needs to contain %d fields", row, N)
		}
		for column, val := range rowList {
			n, err := validateInputCell(val)
			if err != nil {
				return nil, err
			}
			board.Cells[row][column] = n
		}
	}
	if row < N {
		return nil, fmt.Errorf("the board contains only %d rows", row)
	}
	return board, nil
}

func validateInputCell(val string) (int, error) {
	if val == "_" {
		return 0, nil
	}
	digit, err := strconv.Atoi(val)
	if err != nil || digit < 1 || digit > N {
		return 0, fmt.Errorf("only digits from 1 to %d and _ as placeholder are allowed values", N)
	}
	return digit, nil
}

func (b *Board) String() string {
	var output string
	for row, rowList := range b.Cells {
		str := fmt.Sprint(rowList)
		output += str[1 : len(str)-1]
		if row+1 < N {
			output += "\n"
		}
	}
	return output
}
