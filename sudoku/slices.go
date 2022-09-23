package sudoku

type sliceEditor struct {
	alterationFunc func(int8) int8
}

func (editor sliceEditor) toEach(slice []int8) []int8 {
	newSlice := make([]int8, len(slice))
	var i int8
	sliceLength := int8(len(slice))

	for ; i < sliceLength; i++ {
		newSlice[i] = editor.alterationFunc(slice[i])
	}

	return newSlice
}

func add(toAdd int8) sliceEditor {
	return sliceEditor{
		alterationFunc: func(value int8) int8 {
			return value + toAdd
		},
	}
}
