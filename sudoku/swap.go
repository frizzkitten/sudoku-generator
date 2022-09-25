package sudoku

func (sudoku Sudoku) randomSwap() Sudoku {
	swapFunctions := [](func() Sudoku){
		sudoku.swapRandomRows,
		sudoku.swapRandomColumns,
		sudoku.swapRandomNumbers, // add swap mega-rows and mega-colums
		sudoku.swapRandomMegaRows,
	}

	randomIndex := randomInt(int8(len(swapFunctions)))
	swapFunction := swapFunctions[randomIndex]

	return swapFunction()
}

func (sudoku Sudoku) swapRandomRows() Sudoku {
	rowIndex1, rowIndex2 := sudoku.getSwappableIndexes()
	return sudoku.swapRows(rowIndex1, rowIndex2)
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

func (sudoku Sudoku) getSwappableIndexes() (int8, int8) {
	possibleIndexes := getIndexesFromZeroTo(sudoku.scaleRoot - 1)

	indexInPossibleIndexes, index1 := getRandomFromSlice(possibleIndexes)
	possibleIndexes = removeIndex(possibleIndexes, indexInPossibleIndexes)
	_, index2 := getRandomFromSlice(possibleIndexes)

	box := randomInt(sudoku.scaleRoot)
	return box*sudoku.scaleRoot + index1, box*sudoku.scaleRoot + index2
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

func (sudoku Sudoku) swapRandomMegaRows() Sudoku {
	if sudoku.scale < 2 {
		return sudoku
	}

	indexes := shuffle(getIndexesFromZeroTo(sudoku.scaleRoot - 1))

	return sudoku.swapMegaRows(indexes[0], indexes[1])
}

func (sudoku Sudoku) swapMegaRows(rowIndex1, rowIndex2 int8) Sudoku {
	var i int8
	for ; i < sudoku.scaleRoot; i++ {
		sudoku = sudoku.swapRows(rowIndex1*sudoku.scaleRoot+i, rowIndex2*sudoku.scaleRoot+i)
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
