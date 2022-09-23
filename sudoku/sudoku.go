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
	return sudoku.generateDefault().randomize()
}

func (sudoku Sudoku) generateDefault() Sudoku {
	scale := sudoku.scale
	scaleRoot := sudoku.scaleRoot

	for rowIndex := int8(0); rowIndex < scale; rowIndex++ {
		shift := scaleRoot*(rowIndex%scaleRoot) + (rowIndex / scaleRoot)
		sudoku.Rows[rowIndex] = sudoku.getShiftedValues(shift)
	}

	return sudoku
}

func (sudoku Sudoku) randomize() Sudoku {
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

func (sudoku Sudoku) randomSwap() Sudoku {
	swapFunctions := [](func() Sudoku){sudoku.swapRandomRows, sudoku.swapRandomColumns, sudoku.swapRandomNumbers}
	randomIndex := randomInt(int8(len(swapFunctions)))
	swapFunction := swapFunctions[randomIndex]
	return swapFunction()
}

func (sudoku Sudoku) swapRandomRows() Sudoku {
	rowIndex1, rowIndex2 := sudoku.getSwappableIndexes()
	return sudoku.swapRows(rowIndex1, rowIndex2)
}

func (sudoku Sudoku) getSwappableIndexes() (int8, int8) {
	possibleIndexes := getIndexesFromZeroTo(sudoku.scaleRoot - 1)

	indexInPossibleIndexes, index1 := getRandomFromSlice(possibleIndexes)
	possibleIndexes = removeIndex(possibleIndexes, indexInPossibleIndexes)
	_, index2 := getRandomFromSlice(possibleIndexes)

	box := randomInt(sudoku.scaleRoot)
	return box*sudoku.scaleRoot + index1, box*sudoku.scaleRoot + index2
}

func (sudoku Sudoku) swapRows(index1, index2 int8) Sudoku {
	sudoku.Rows[index1], sudoku.Rows[index2] = sudoku.Rows[index2], sudoku.Rows[index1]
	return sudoku
}

func (sudoku Sudoku) swapRandomColumns() Sudoku {
	columnIndex1, columnIndex2 := sudoku.getSwappableIndexes()
	return sudoku.swapColumns(columnIndex1, columnIndex2)
}

func (sudoku Sudoku) swapColumns(columnIndex1, columnIndex2 int8) Sudoku {
	for rowIndex := int8(0); rowIndex < sudoku.scale; rowIndex++ {
		sudoku.Rows[rowIndex][columnIndex1], sudoku.Rows[rowIndex][columnIndex2] = sudoku.Rows[rowIndex][columnIndex2], sudoku.Rows[rowIndex][columnIndex1]
	}
	return sudoku
}

func (sudoku Sudoku) swapRandomNumbers() Sudoku {
	if sudoku.scale < 2 {
		return sudoku
	}

	values := sudoku.getShuffledAllValues()
	number1, number2 := values[0], values[1]

	return sudoku.swapNumbers(number1, number2)
}

func (sudoku Sudoku) swapNumbers(number1, number2 int8) Sudoku {
	for rowIndex, row := range sudoku.Rows {
		for columnIndex, value := range row {
			if value == number1 {
				sudoku.Rows[rowIndex][columnIndex] = number2
			} else if value == number2 {
				sudoku.Rows[rowIndex][columnIndex] = number1
			}
		}
	}
	return sudoku
}

func (sudoku Sudoku) getShuffledAllValues() []int8 {
	values := make([]int8, sudoku.scale)
	for i := int8(0); i < sudoku.scale; i++ {
		values[i] = i + 1
	}
	return shuffle(values)
}

func (sudoku Sudoku) sliceWithLengthOfScale() []bool {
	return make([]bool, sudoku.scale)
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
