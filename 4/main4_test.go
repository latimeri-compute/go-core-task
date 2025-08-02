package main

import (
	"slices"
	"testing"
)

func TestGetUniqueSlice(t *testing.T) {
	tests := []struct {
		name   string
		slice1 []string
		slice2 []string
		want   []string
	}{
		{
			name:   "пересекаются",
			slice1: []string{"apple", "banana", "cherry", "date", "43", "lead", "gno1"},
			slice2: []string{"banana", "date", "fig"},
			want:   []string{"apple", "cherry", "43", "lead", "gno1"},
		},
		{
			name:   "не пересекаются",
			slice1: []string{"apple", "banana", "cherry", "date", "43", "lead", "gno1"},
			slice2: []string{"fig"},
			want:   []string{"apple", "banana", "cherry", "date", "43", "lead", "gno1"},
		},
		{
			name:   "пустой слайс #1",
			slice1: []string{},
			slice2: []string{"fig"},
			want:   []string{},
		},
		{
			name:   "пустой слайс #2",
			slice1: []string{"apple", "banana", "cherry", "date", "43", "lead", "gno1"},
			slice2: []string{},
			want:   []string{"apple", "banana", "cherry", "date", "43", "lead", "gno1"},
		},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			got := getUniqueSliceFromFirst(test.slice1, test.slice2)

			if !slices.Equal(got, test.want) {
				t.Errorf("got: %v; want: %v", got, test.want)
			}
		})
	}
}
