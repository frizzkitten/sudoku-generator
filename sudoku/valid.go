package sudoku

// IsValid tells you if this is a valid sudoku
func (sudoku Sudoku) IsValid() bool {
	// TODO
	return false
}
func (sudoku Sudoku) isValidNumberInColumnAndBox(number, rowIndex, columnIndex int8) bool {
	return sudoku.isValidNumberInColumn(number, rowIndex, columnIndex) && sudoku.isValidNumberInBox(number, rowIndex, columnIndex)
}

func (sudoku Sudoku) isValidNumberInColumn(number int8, rowIndex int8, columnIndex int8) bool {
	for rowIndexToCheck := int8(0); rowIndexToCheck < sudoku.scale; rowIndexToCheck++ {
		if rowIndex == rowIndexToCheck {
			continue
		}

		numberInColumn := sudoku.Rows[rowIndexToCheck][columnIndex]
		if numberInColumn == number {
			return false
		}
	}

	return true
}

func (sudoku Sudoku) isValidNumberInBox(number int8, rowIndex int8, columnIndex int8) bool {
	scaleRoot := sudoku.scaleRoot
	boxRow := rowIndex / scaleRoot
	boxColumn := columnIndex / scaleRoot
	for rowIndexToCheck := boxRow * scaleRoot; rowIndexToCheck < boxRow*scaleRoot+1; rowIndexToCheck++ {
		for columnIndexToCheck := boxColumn * scaleRoot; columnIndexToCheck < boxColumn*scaleRoot+scaleRoot; columnIndexToCheck++ {
			if rowIndexToCheck == rowIndex && columnIndexToCheck == columnIndex {
				continue
			}
			if sudoku.Rows[rowIndexToCheck][columnIndexToCheck] == number {
				return false
			}
		}
	}

	return true
}
