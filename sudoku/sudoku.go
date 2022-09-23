package sudoku

// Create returns a Sudoku with some numbers
// filled in and some intentionally missing.
func Create(scaleRoot int8) Sudoku {
	return makeEmptySudoku(scaleRoot).fillInNumbers()
}

func makeEmptySudoku(scaleRoot int8) Sudoku {
	scale := scaleRoot * scaleRoot
	sudoku := Sudoku{Rows: make([]Row, scale), scale: scale, scaleRoot: scaleRoot}

	for rowIndex := int8(0); rowIndex < scale; rowIndex++ {
		row := make([]int8, scale)
		sudoku.Rows[rowIndex] = row
	}

	return sudoku
}

func (sudoku Sudoku) fillInNumbers() Sudoku {
	return sudoku.fillInDefaultNumbers().randomize()
}

func (sudoku Sudoku) fillInDefaultNumbers() Sudoku {
	scale := sudoku.scale
	scaleRoot := sudoku.scaleRoot

	for rowIndex := int8(0); rowIndex < scale; rowIndex++ {
		shift := scaleRoot*(rowIndex%scaleRoot) + (rowIndex / scaleRoot)
		sudoku.Rows[rowIndex] = sudoku.getShiftedValues(shift)
	}

	return sudoku
}

func (sudoku Sudoku) randomize() Sudoku {
	if sudoku.scale == 1 {
		return sudoku
	}

	var i int8
	for ; i < swaps; i++ {
		sudoku = sudoku.randomSwap()
	}

	return sudoku
}

func (sudoku Sudoku) getShiftedValues(shift int8) []int8 {
	scale := sudoku.scale
	values := make([]int8, scale)
	var index int8
	for ; index < scale; index++ {
		values[index] = (index+shift)%scale + 1
	}
	return values
}
