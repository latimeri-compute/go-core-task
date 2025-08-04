package main

import (
	"slices"
	"testing"
)

func TestRandInt(t *testing.T) {
	amountToGenerate := 1024
	maxN := 512

	mapOfInt := make(map[int]bool, amountToGenerate)
	var counterSim int
	var counterDif int

	for range amountToGenerate {
		newInt := randomInt(maxN)
		_, ok := mapOfInt[newInt]
		if ok {
			counterSim++
		} else {
			mapOfInt[newInt] = true
			counterDif++
		}
	}

	if counterSim == amountToGenerate {
		t.Errorf("Были сгенерированы исключительно одинаковые числа в количестве %d ", counterSim)
	}
	t.Logf("Были сгенерированы %d разных чисел, %d одинаковых", counterDif, counterSim)
}

func TestRandomSlice(t *testing.T) {
	sliceOfSlices := make([][]int, 2)

	for i := range 2 {
		slice := randomSlice(20, 512)
		sliceOfSlices[i] = slice
	}

	if slices.Equal(sliceOfSlices[0], sliceOfSlices[1]) {
		t.Errorf("Сгенерированы одинаковые слайсы %v и %v", sliceOfSlices[0], sliceOfSlices[1])
	}
}
