package sudoku

import (
	"math/rand"
	"time"
)

func getIndexesFromZeroTo(max int8) []int8 {
	if max < 0 {
		max = -1
	}

	indexes := make([]int8, max+1)
	var i int8
	for ; i <= max; i++ {
		indexes[i] = i
	}
	return indexes
}

func shuffle(slice []int8) []int8 {
	sliceLength := int8(len(slice))
	newSlice := copySlice(slice)

	for i := range newSlice {
		randomIndex := randomInt(sliceLength)
		newSlice[i], newSlice[randomIndex] = newSlice[randomIndex], newSlice[i]
	}

	return newSlice
}

func randomInt(firstAboveMax int8) int8 {
	random := rand.New(rand.NewSource(time.Now().UnixNano()))
	return int8(random.Int() % int(firstAboveMax))
}

func  copySlice[T any](slice []T) []T {
	copy := make([]T, len(slice))
	for index, value := range slice {
		copy[index] = value
	}
	return copy
}
