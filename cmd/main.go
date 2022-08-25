package main

import (
	"errors"
	"fmt"
	"math/rand"
	"time"
)

const (
	maxRetries = int8(30)
	swaps      = 1
)

func main() {
	// sudoku := Sudoku{
	// 	Rows: []Row{
	// 		{1, 2, 3, 4, 5, 6, 7, 8, 9},
	// 		{4, 5, 6, 7, 8, 9, 1, 2, 3},
	// 		{7, 8, 9, 1, 2, 3, 4, 5, 6},
	// 		{2, 3, 4, 5, 6, 7, 8, 9, 1},
	// 		{1, 2, 3, 4, 5, 6, 7, 8, 9},
	// 		{1, 2, 3, 4, 5, 6, 7, 8, 9},
	// 		{1, 2, 3, 4, 5, 6, 7, 8, 9},
	// 		{1, 2, 3, 4, 5, 6, 7, 8, 9},
	// 		{1, 2, 3, 4, 5, 6, 7, 8, 9},
	// 	},
	// }

	sudoku := MakeEmptySudoku().GenerateFromEmpty()
	sudoku.Print()
}

// Sudoku is a sudoku
type Sudoku struct {
	Rows      []Row
	scale     int8
	scaleRoot int8
}

// Row is a row
type Row []int8

// MakeEmptySudoku does stuff
func MakeEmptySudoku() Sudoku {
	scale := int8(9)
	scaleRoot := int8(3)
	sudoku := Sudoku{Rows: make([]Row, scale), scale: scale, scaleRoot: scaleRoot}

	for rowIndex := int8(0); rowIndex < scale; rowIndex++ {
		row := make([]int8, scale)
		sudoku.Rows[rowIndex] = row
	}

	return sudoku
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

func (sudoku Sudoku) getShiftedValues(shift int8) []int8 {
	scale := sudoku.scale
	values := make([]int8, scale)
	for index := int8(0); index < scale; index++ {
		values[index] = (index+shift)%scale + 1
	}
	return values
}

func (sudoku Sudoku) randomSwap() Sudoku {
	swapFunctions := [](func() Sudoku){sudoku.swapRandomRows/*, sudoku.swapRandomColumns, sudoku.swapRandomNumbers*/}
	randomIndex := randomInt(int8(len(swapFunctions)))
	swapFunction := swapFunctions[randomIndex]
	return swapFunction()
}

func getIndexesFromZeroTo(max int8) []int8 {
	indexes := make([]int8, max+1)
	for i := int8(0); i <= max; i++ {
		indexes[i] = i
	}
	return indexes
}

func (sudoku Sudoku) swapRandomRows() Sudoku {
	possibleIndexes := getIndexesFromZeroTo(sudoku.scaleRoot - 1)

	box := randomInt(sudoku.scaleRoot)

	indexInPossibleIndexes := randomInt(int8(len(possibleIndexes)))
	rowIndex1 := possibleIndexes[indexInPossibleIndexes]

	possibleIndexes = removeIndex(possibleIndexes, indexInPossibleIndexes)

	indexInPossibleIndexes = randomInt(int8(len(possibleIndexes)))
	rowIndex2 := possibleIndexes[indexInPossibleIndexes]

	return sudoku.swapRows(box+rowIndex1, box+rowIndex2)
}

func removeIndex[T any](slice []T, index int8) []T {
	return append(slice[:index], slice[index:]...)
}

func (sudoku Sudoku) swapRows(index1, index2 int8) Sudoku {
	fmt.Println("swapping", index1, "and", index2)
	sudoku.Rows[index1], sudoku.Rows[index2] = sudoku.Rows[index2], sudoku.Rows[index1]
	return sudoku
}

func (sudoku Sudoku) swapRandomColumns() Sudoku {
	// TODO
	return sudoku
}

func (sudoku Sudoku) swapRandomNumbers() Sudoku {
	// TODO
	return sudoku
}

// GenerateFromEmpty generates from empty duh
func (sudoku Sudoku) GenerateFromEmpty() Sudoku {
	sudoku = sudoku.generateDefault()
	for i := int8(0); i < swaps; i++ {
		sudoku = sudoku.randomSwap()
	}
	return sudoku

	// retriesPerRow := sudoku.getRetriesPerRow()

	// var err error
	// for rowIndex := int8(0); rowIndex < sudoku.scale; rowIndex++ {
	// 	sudoku, err = sudoku.GenerateRow(rowIndex)
	// 	if err != nil {
	// 		fmt.Println("retrying row", rowIndex)
	// 		retriesPerRow[rowIndex]--
	// 		if retriesPerRow[rowIndex] < 1 {
	// 			return sudoku, fmt.Errorf("ran out of retries for row %v", rowIndex)
	// 		}
	// 		rowIndex--
	// 	}
	// }

	// return sudoku, nil
}

func (sudoku Sudoku) getRetriesPerRow() map[int8]int8 {
	retriesPerRow := map[int8]int8{}
	for i := int8(0); i < sudoku.scale; i++ {
		retriesPerRow[i] = maxRetries
	}
	return retriesPerRow
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

// GenerateRow returns the sudoku with the wanted row
func (sudoku Sudoku) GenerateRow(rowIndex int8) (Sudoku, error) {
	availableNumbers := sudoku.getShuffledAllValues()

	for columnIndex := range sudoku.sliceWithLengthOfScale() {
		for availableNumberIndex, availableNumber := range availableNumbers {
			if sudoku.IsValidNumberInColumnAndBox(availableNumber, rowIndex, int8(columnIndex)) {
				// add the number to the sudoku
				sudoku.Rows[rowIndex][columnIndex] = availableNumber

				// remove the number from the list of available numbers
				availableNumbers = append(availableNumbers[0:availableNumberIndex], availableNumbers[availableNumberIndex+1:]...)

				break
			}
		}

		if sudoku.Rows[rowIndex][columnIndex] == 0 {
			return sudoku, errors.New("ran out of retries :(")
		}
	}

	return sudoku, nil
}

// Print prints it u jangus
func (sudoku Sudoku) Print() {
	for _, row := range sudoku.Rows {
		for _, value := range row {
			fmt.Printf("%v ", value)
		}
		fmt.Println()
	}
}

// IsValid tells you if this is a valid sudoku
func (sudoku Sudoku) IsValid() bool {
	// TODO
	return false
}

// IsValidNumberInColumnAndBox tells you if the number is allowed in a column and box
func (sudoku Sudoku) IsValidNumberInColumnAndBox(number, rowIndex, columnIndex int8) bool {
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

func shuffle(slice []int8) []int8 {
	for i := 0; i < len(slice); i++ {
		randomIndex := randomInt(int8(len(slice)))
		slice[i], slice[randomIndex] = slice[randomIndex], slice[i]
	}

	return slice
}

func randomInt(firstAboveMax int8) int8 {
	random := rand.New(rand.NewSource(time.Now().UnixNano()))
	return int8(random.Int() % int(firstAboveMax))
}
