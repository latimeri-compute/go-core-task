package main

import (
	"fmt"
	"slices"
	"testing"
)

func TestNewOriginalSlice(t *testing.T) {
	sliceOfSlices := make([][]int, 2)
	for i := range 2 {
		slice := newOriginalSlice()
		sliceOfSlices[i] = slice
		if len(slice) != 10 {
			t.Errorf("длина %v не равна 10", slice)
		}
	}
	if slices.Equal(sliceOfSlices[0], sliceOfSlices[1]) {
		t.Errorf("Сгенерированы одинаковые слайсы %v и %v", sliceOfSlices[0], sliceOfSlices[1])
	}
}

func TestSliceExample(t *testing.T) {
	tests := []struct {
		slice []int
		want  []int
	}{
		{
			slice: []int{0, -1, 23, 6, 9, 127, 4, 4},
			want:  []int{0, 6, 4, 4},
		},
		{
			slice: []int{0, 0, 0, 0, 0, 0},
			want:  []int{0, 0, 0, 0, 0, 0},
		},
		{
			slice: []int{1, 1, 1, 1, 1, 1, 1},
			want:  []int{},
		},
		{
			slice: []int{},
			want:  []int{},
		},
	}
	for ti, test := range tests {
		t.Run(fmt.Sprintf("%02d", ti), func(t *testing.T) {
			got := sliceExample(test.slice)
			EqualSlices(t, got, test.want)
		})
	}

}

func TestAddElements(t *testing.T) {
	tests := []struct {
		slice []int
		want  []int
		num   int
	}{
		{
			slice: []int{0, -1, 23, 6, 9, 127, 4, 4},
			num:   0,
			want:  []int{0, -1, 23, 6, 9, 127, 4, 4, 0},
		},
		{
			slice: []int{0, 0, 0, 0, 0, 0},
			num:   -1,
			want:  []int{0, 0, 0, 0, 0, 0, -1},
		},
		{
			slice: []int{1, 1, 1, 1, 1, 1, 1},
			num:   42,
			want:  []int{1, 1, 1, 1, 1, 1, 1, 42},
		},
	}
	for ti, test := range tests {
		t.Run(fmt.Sprintf("%02d", ti), func(t *testing.T) {
			got := addElements(test.slice, test.num)
			EqualSlices(t, got, test.want)
		})
	}
}

func TestCopySlice(t *testing.T) {
	tests := []struct {
		slice []int
	}{
		{
			slice: []int{0, -1, 23, 6, 9, 127, 4, 4},
		},
		{
			slice: []int{},
		},
	}
	for ti, test := range tests {
		t.Run(fmt.Sprintf("%02d", ti), func(t *testing.T) {
			got := copySlice(test.slice)
			EqualSlices(t, got, test.slice)

			// пока не могу найти способ достать адрес ссылки на массив
			// разве что если использовать sprintf(%p)
			// не знаю, каким способом лучше проверить различие адресов, честно говоря
			if len(got) > 0 {
				equalityCounter := 0
				for i := range 2 {
					got[0] = i
					if slices.Equal(got, test.slice) {
						equalityCounter++
					}
				}
				if equalityCounter > 1 {
					t.Error("Адрес копированного слайса совпадает с исходным")
				}
			}

		})
	}
}

func TestRemoveElement(t *testing.T) {
	tests := []struct {
		slice []int
		num   int
		want  []int
	}{
		{
			slice: []int{0, -1, 23, 6, 9, 127, 4, 4},
			num:   3,
			want:  []int{0, -1, 23, 9, 127, 4, 4},
		},
		{
			slice: []int{0, -1, 23, 6, 9, 127, 4, 4},
			num:   99,
			want:  []int{0, -1, 23, 6, 9, 127, 4, 4},
		},
		{
			slice: []int{0, -1, 23, 6, 9, 127, 4, 4},
			num:   -1,
			want:  []int{0, -1, 23, 6, 9, 127, 4, 4},
		},
	}
	for ti, test := range tests {
		t.Run(fmt.Sprintf("%02d", ti), func(t *testing.T) {
			got := removeElement(test.slice, test.num)
			EqualSlices(t, got, test.want)
		})
	}
}

func EqualSlices[S ~[]E, E comparable](t *testing.T, got, want S) {
	t.Helper()
	if !slices.Equal(got, want) {
		t.Errorf("got: %v; want: %v", got, want)
	}
}
