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
