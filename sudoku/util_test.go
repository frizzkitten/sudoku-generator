package sudoku

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGetIndexesFromZeroTo(t *testing.T) {
	t.Run("-5", func(t *testing.T) {
		assert.Equal(t, []int8{}, getIndexesFromZeroTo(-5))
	})

	t.Run("0", func(t *testing.T) {
		assert.Equal(t, []int8{0}, getIndexesFromZeroTo(0))
	})

	t.Run("5", func(t *testing.T) {
		assert.Equal(t, []int8{0, 1, 2, 3, 4, 5}, getIndexesFromZeroTo(5))
	})
}

func TestShuffle(t *testing.T) {
	original := getIndexesFromZeroTo(100)
	shuffled := shuffle(original)
	assert.NotEqual(t, original, shuffled) // extremely unlikely to be equal
}

func TestRandomInt(t *testing.T) {
	generatedNumbers := map[int8]bool{}

	for i := 0; i < 1000; i++ {
		random := randomInt(3)
		assert.Contains(t, getIndexesFromZeroTo(2), random)
		generatedNumbers[random] = true
	}

	assert.Len(t, generatedNumbers, 3) // very unlikely to not have generated a 0, 1, and 2 at least once each
}

func TestCopy(t *testing.T) {
	original := []int8{0, 1, 2}
	copy := copySlice(original)
	assert.Equal(t, original, copy)
	original[1] = 3
	assert.NotEqual(t, original, copy)
}
