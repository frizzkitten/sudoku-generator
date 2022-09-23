package sudoku 

import (
	"math/rand"
	"time"
)

func getIndexesFromZeroTo(max int8) []int8 {
	indexes := make([]int8, max+1)
	for i := int8(0); i <= max; i++ {
		indexes[i] = i
	}
	return indexes
}

func getRandomFromSlice[T any](values []T) (int8, T) {
	index := randomInt(int8(len(values)))
	value := values[index]
	return index, value
}

func removeIndex[T any](slice []T, index int8) []T {
	return append(slice[:index], slice[index+1:]...)
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