package sudoku

// Sudoku is a sudoku
type Sudoku struct {
	Rows      []Row
	scale     int8
	scaleRoot int8
}

// Row is a row
type Row []int8
