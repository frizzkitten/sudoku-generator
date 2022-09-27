package sudoku

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSwapLines(t *testing.T) {
	t.Run("rows", func(t *testing.T) {
		s := Sudoku{
			scale:     4,
			scaleRoot: 2,
			Rows: []Row{
				Row{1, 2, 3, 4},
				Row{3, 4, 1, 2},
				Row{2, 3, 4, 1},
				Row{4, 1, 2, 3},
			},
		}

		assert.Equal(t, []Row{
			Row{1, 2, 3, 4},
			Row{3, 4, 1, 2},
			Row{4, 1, 2, 3},
			Row{2, 3, 4, 1},
		}, s.swapLines(ROW, 2, 3).Rows)
	})

	t.Run("columns", func(t *testing.T) {
		s := Sudoku{
			scale:     4,
			scaleRoot: 2,
			Rows: []Row{
				Row{1, 2, 3, 4},
				Row{3, 4, 1, 2},
				Row{2, 3, 4, 1},
				Row{4, 1, 2, 3},
			},
		}

		assert.Equal(t, []Row{
			Row{1, 2, 4, 3},
			Row{3, 4, 2, 1},
			Row{2, 3, 1, 4},
			Row{4, 1, 3, 2},
		}, s.swapLines(COLUMN, 2, 3).Rows)
	})
}

func TestSwapNumbers(t *testing.T) {
	s := Sudoku{
		scale:     4,
		scaleRoot: 2,
		Rows: []Row{
			Row{1, 2, 3, 4},
			Row{3, 4, 1, 2},
			Row{2, 3, 4, 1},
			Row{4, 1, 2, 3},
		},
	}

	assert.Equal(t, []Row{
		Row{2, 1, 3, 4},
		Row{3, 4, 2, 1},
		Row{1, 3, 4, 2},
		Row{4, 2, 1, 3},
	}, s.swapNumbers(1, 2).Rows)
}

func TestSwapMegaLines(t *testing.T) {
	t.Run("mega-rows", func(t *testing.T) {
		s := Sudoku{
			scale:     4,
			scaleRoot: 2,
			Rows: []Row{
				Row{1, 2, 3, 4},
				Row{3, 4, 1, 2},
				Row{2, 3, 4, 1},
				Row{4, 1, 2, 3},
			},
		}

		assert.Equal(t, []Row{
			Row{2, 3, 4, 1},
			Row{4, 1, 2, 3},
			Row{1, 2, 3, 4},
			Row{3, 4, 1, 2},
		}, s.swapMegaLines(ROW, 0, 1).Rows)
	})

	t.Run("mega-columns", func(t *testing.T) {
		s := Sudoku{
			scale:     4,
			scaleRoot: 2,
			Rows: []Row{
				Row{1, 2, 3, 4},
				Row{3, 4, 1, 2},
				Row{2, 3, 4, 1},
				Row{4, 1, 2, 3},
			},
		}

		assert.Equal(t, []Row{
			Row{3, 4, 1, 2},
			Row{1, 2, 3, 4},
			Row{4, 1, 2, 3},
			Row{2, 3, 4, 1},
		}, s.swapMegaLines(COLUMN, 0, 1).Rows)
	})
}
