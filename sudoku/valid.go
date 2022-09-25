package sudoku

// IsValid tells you if this is a valid sudoku
func (sudoku Sudoku) IsValid() bool {
	return sudoku.numbersAreValid() && sudoku.rowsAreValid() && sudoku.columnsAreValid() && sudoku.boxesAreValid()
}

func (sudoku Sudoku) numbersAreValid() bool {
	for _, row := range sudoku.Rows {
		for _, value := range row {
			if value < 1 || value > sudoku.scale {
				return false
			}
		}
	}

	return true
}

func (sudoku Sudoku) rowsAreValid() bool {
	for _, row := range sudoku.Rows {
		seen := map[int8]bool{}

		for _, value := range row {
			if seen[value] {
				return false
			}
			seen[value] = true
		}
	}

	return true
}

func (sudoku Sudoku) columnsAreValid() bool {
	for columnIndex := range sudoku.Rows {
		seen := map[int8]bool{}

		for rowIndex := int8(0); rowIndex < sudoku.scale; rowIndex++ {
			value := sudoku.Rows[rowIndex][columnIndex]
			if seen[value] {
				return false
			}
			seen[value] = true
		}
	}

	return true
}

func (sudoku Sudoku) boxesAreValid() bool {
	for boxRowIndex := int8(0); boxRowIndex < sudoku.scaleRoot; boxRowIndex++ {
		for boxColumnIndex := int8(0); boxColumnIndex < sudoku.scaleRoot; boxColumnIndex++ {
			if !sudoku.isValidBox(boxRowIndex, boxColumnIndex) {
				return false
			}
		}
	}

	return true
}

func (sudoku Sudoku) isValidBox(boxRowIndex, boxColumnIndex int8) bool {
	seen := map[int8]bool{}
	scaleRoot := sudoku.scaleRoot

	for rowIndex := boxRowIndex * scaleRoot; rowIndex < boxRowIndex*scaleRoot+scaleRoot; rowIndex++ {
		for columnIndex := boxColumnIndex * scaleRoot; columnIndex < boxColumnIndex*scaleRoot+scaleRoot; columnIndex++ {
			value := sudoku.Rows[rowIndex][columnIndex]
			if seen[value] {
				return false
			}
			seen[value] = true
		}
	}

	return true
}
