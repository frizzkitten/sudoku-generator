package sudoku

func (sudoku Sudoku) randomSwap() Sudoku {
	if sudoku.scale < 2 {
		return sudoku
	}

	swapFunctions := [](func() Sudoku){
		sudoku.swapRandomNumbers,
		sudoku.swapRandomRows,
		sudoku.swapRandomColumns,
		sudoku.swapRandomMegaRows,
		sudoku.swapRandomMegaColumns,
	}

	randomIndex := randomInt(int8(len(swapFunctions)))
	swapFunction := swapFunctions[randomIndex]

	return swapFunction()
}

func (sudoku Sudoku) swapRandomNumbers() Sudoku {
	values := sudoku.getShuffledAllValues()
	number1, number2 := values[0], values[1]
	return sudoku.swapNumbers(number1, number2)
}

func (sudoku Sudoku) swapRandomRows() Sudoku {
	return sudoku.swapRandomLines(ROW)
}

func (sudoku Sudoku) swapRandomColumns() Sudoku {
	return sudoku.swapRandomLines(COLUMN)
}

func (sudoku Sudoku) swapRandomMegaRows() Sudoku {
	return sudoku.swapRandomMegaLines(ROW)
}

func (sudoku Sudoku) swapRandomMegaColumns() Sudoku {
	return sudoku.swapRandomMegaLines(COLUMN)
}

func (sudoku Sudoku) swapRandomLines(lineType int8) Sudoku {
	rowIndex1, rowIndex2 := sudoku.getSwappableIndexes()
	return sudoku.swapLines(lineType, rowIndex1, rowIndex2)
}

func (sudoku Sudoku) swapRandomMegaLines(lineType int8) Sudoku {
	indexes := shuffle(getIndexesFromZeroTo(sudoku.scaleRoot - 1))
	return sudoku.swapMegaLines(lineType, indexes[0], indexes[1])
}

func (sudoku Sudoku) swapLines(lineType, index1, index2 int8) Sudoku {
	if lineType == ROW {
		return sudoku.swapRows(index1, index2)
	}
	return sudoku.swapColumns(index1, index2)
}

func (sudoku Sudoku) swapRows(index1, index2 int8) Sudoku {
	sudoku.Rows[index1], sudoku.Rows[index2] = sudoku.Rows[index2], sudoku.Rows[index1]
	return sudoku
}

func (sudoku Sudoku) swapColumns(columnIndex1, columnIndex2 int8) Sudoku {
	for rowIndex := int8(0); rowIndex < sudoku.scale; rowIndex++ {
		sudoku.Rows[rowIndex][columnIndex1], sudoku.Rows[rowIndex][columnIndex2] = sudoku.Rows[rowIndex][columnIndex2], sudoku.Rows[rowIndex][columnIndex1]
	}
	return sudoku
}

func (sudoku Sudoku) getSwappableIndexes() (int8, int8) {
	possibleIndexes := shuffle(getIndexesFromZeroTo(sudoku.scaleRoot - 1))
	index1 := possibleIndexes[0]
	index2 := possibleIndexes[1]

	box := randomInt(sudoku.scaleRoot)
	return box*sudoku.scaleRoot + index1, box*sudoku.scaleRoot + index2
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

func (sudoku Sudoku) swapMegaLines(lineType int8, rowIndex1, rowIndex2 int8) Sudoku {
	var i int8
	for ; i < sudoku.scaleRoot; i++ {
		sudoku = sudoku.swapLines(lineType, rowIndex1*sudoku.scaleRoot+i, rowIndex2*sudoku.scaleRoot+i)
	}

	return sudoku
}

func (sudoku Sudoku) getShuffledAllValues() []int8 {
	values := sudoku.getSliceWithOneThroughScale()
	return shuffle(values)
}

func (sudoku Sudoku) getSliceWithOneThroughScale() []int8 {
	indexesFromZeroToScaleMinusOne := getIndexesFromZeroTo(sudoku.scale - 1)
	return add(1).toEach(indexesFromZeroToScaleMinusOne)
}
