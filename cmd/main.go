package main

import (
	"fmt"
	"math"
	"os"
	"strconv"

	"github.com/frizzkitten/sudoku-generator/sudoku"
)

func main() {
	base := getBase()
	doku := sudoku.Create(base)
	doku.Print()
}

func getBase() int8 {
	fmt.Println("What do you want your sudoku base to be? Default is 3.")

	var input string
	fmt.Scanln(&input)

	base, err := strconv.Atoi(input)
	if base < 1 || err != nil {
		fmt.Printf("Invalid base. Must be an integer > 0 and < %v.\n", math.MaxInt8)
		os.Exit(1)
	}

	return int8(base)
}
