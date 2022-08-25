package main

import (
	"testing"
)

func TestGetShiftedValues(t *testing.T) {
	MakeEmptySudoku().GenerateFromEmpty().Print()
	t.Fail()
}
