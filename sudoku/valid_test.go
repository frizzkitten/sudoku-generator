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
				Row{1, 2, 3, 4},
				Row{3, 4, 1, 2},
				Row{2, 3, 4, 1},
				Row{4, 1, 2, 3},
			},
		}

		assert.True(t, s.IsValid())
	})
}
