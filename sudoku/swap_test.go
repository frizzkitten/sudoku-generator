package sudoku

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSwapRows(t *testing.T) {
	// TODO
}

func TestSwapColumns(t *testing.T) {
	//TODO
}

func TestSwapNumbers(t *testing.T) {
	// TODO
}

func TestSwapMegaRows(t *testing.T) {
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
	}, s.swapMegaRows(0, 1).Rows)
}

func TestSwapMegaColumns(t *testing.T) {
	// TODO
}
