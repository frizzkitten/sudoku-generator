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

func getRandomFromSlice[T any](values []T) (int8, T) {
	numValues := int8(len(values))
	index := randomInt(numValues)
	value := values[index]
	return index, value
}

func removeIndex[T any](slice []T, index int8) []T {
	return append(slice[:index], slice[index+1:]...)
}

func shuffle(slice []int8) []int8 {
	var i int8
	sliceLength := int8(len(slice))
	
	for ; i < sliceLength; i++ {
		randomIndex := randomInt(sliceLength)
		slice[i], slice[randomIndex] = slice[randomIndex], slice[i]
	}

	return slice
}

func randomInt(firstAboveMax int8) int8 {
	random := rand.New(rand.NewSource(time.Now().UnixNano()))
	return int8(random.Int() % int(firstAboveMax))
}