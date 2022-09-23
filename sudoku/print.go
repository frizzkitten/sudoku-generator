package sudoku

import "fmt"

// Print prints the sudoku
func (sudoku Sudoku) Print() {
	for _, row := range sudoku.Rows {
		for _, value := range row {
			fmt.Printf("%v ", value)
		}
		fmt.Println()
	}
}
