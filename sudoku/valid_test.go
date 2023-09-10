package sudoku

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestIsValid(t *testing.T) {
	t.Run("is valid", func(t *testing.T) {
		s := Sudoku{
			scale:     4,
			scaleRoot: 2,
			Rows: []Row{
				{1, 2, 3, 4},
				{3, 4, 1, 2},
				{2, 3, 4, 1},
				{4, 1, 2, 3},
			},
		}

		assert.True(t, s.IsValid())
	})

	t.Run("invalid row", func(t *testing.T) {
		s := Sudoku{
			scale:     4,
			scaleRoot: 2,
			Rows: []Row{
				{1, 2, 2, 4},
				{3, 4, 1, 3},
				{2, 3, 4, 1},
				{4, 1, 3, 2},
			},
		}

		assert.False(t, s.IsValid())
	})

	t.Run("invalid column", func(t *testing.T) {
		s := Sudoku{
			scale:     4,
			scaleRoot: 2,
			Rows: []Row{
				{1, 2, 3, 4},
				{3, 4, 2, 1},
				{3, 1, 4, 2},
				{4, 2, 1, 3},
			},
		}

		assert.False(t, s.IsValid())
	})

	t.Run("invalid box", func(t *testing.T) {
		s := Sudoku{
			scale:     4,
			scaleRoot: 2,
			Rows: []Row{
				{1, 2, 3, 4},
				{2, 4, 1, 3},
				{3, 1, 4, 2},
				{4, 3, 2, 1},
			},
		}

		assert.False(t, s.IsValid())
	})

	t.Run("invalid number (too big)", func(t *testing.T) {
		s := Sudoku{
			scale:     4,
			scaleRoot: 2,
			Rows: []Row{
				{1, 2, 3, 4},
				{3, 4, 1, 2},
				{2, 3, 5, 1},
				{4, 1, 2, 3},
			},
		}

		assert.False(t, s.IsValid())
	})

	t.Run("invalid number (too small)", func(t *testing.T) {
		s := Sudoku{
			scale:     4,
			scaleRoot: 2,
			Rows: []Row{
				{1, 2, 3, 4},
				{3, 4, 1, 0},
				{2, 3, 4, 1},
				{4, 1, 2, 3},
			},
		}

		assert.False(t, s.IsValid())
	})
}
