package main

import (
	"fmt"
	"math/rand"
)

func main() {
	// sudoku := Sudoku{
	// 	Rows: []Row{
	// 		{1, 2, 3, 4, 5, 6, 7, 8, 9},
	// 		{1, 2, 3, 4, 5, 6, 7, 8, 9},
	// 		{1, 2, 3, 4, 5, 6, 7, 8, 9},
	// 		{1, 2, 3, 4, 5, 6, 7, 8, 9},
	// 		{1, 2, 3, 4, 5, 6, 7, 8, 9},
	// 		{1, 2, 3, 4, 5, 6, 7, 8, 9},
	// 		{1, 2, 3, 4, 5, 6, 7, 8, 9},
	// 		{1, 2, 3, 4, 5, 6, 7, 8, 9},
	// 		{1, 2, 3, 4, 5, 6, 7, 8, 9},
	// 	},
	// }

	sudoku := MakeEmptySudoku()
	sudoku = sudoku.GenerateFromEmpty()
	sudoku.Print()
}

// Sudoku is a sudoku
type Sudoku struct {
	Rows []Row
}

// Row is a row
type Row []int8

// MakeEmptySudoku fuks u in the butt
func MakeEmptySudoku() Sudoku {
	sudoku := Sudoku{Rows: make([]Row, 9)}

	for rowIndex := 0; rowIndex < 9; rowIndex++ {
		row := make([]int8, 9)
		sudoku.Rows[rowIndex] = row
	}

	return sudoku
}

// GenerateFromEmpty generates from empty duh
func (sudoku Sudoku) GenerateFromEmpty() Sudoku {
	firstRow := sudoku.Rows[0]
	availableNumbers := []int8{1, 2, 3, 4, 5, 6, 7, 8, 9}

	for index := range firstRow {
		randomIndex := rand.Intn(len(availableNumbers))
		firstRow[index] = availableNumbers[randomIndex]
		availableNumbers = append(availableNumbers[0:randomIndex], availableNumbers[randomIndex+1:]...)
	}

	return sudoku
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
	return false
}

// IsValidBox tells you if the specific index is allowed
func (sudoku Sudoku) IsValidBox() bool {
	return false
}
