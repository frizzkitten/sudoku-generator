package sudoku

import "fmt"

// Print prints the sudoku
func (sudoku Sudoku) Print() {
	for _, row := range sudoku.Rows {
		for _, value := range row {
			// fmt.Printf("%v ", translate(value))
			fmt.Printf("%v ", value)
		}
		fmt.Println()
	}
}

// func translate(value int8) string {
// 	return map[int8]string{
// 		1: "c",
// 		2: "a",
// 		3: "t",
// 		4: "s",
// 	}[value]
// }
