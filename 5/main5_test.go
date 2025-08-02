package main

import (
	"slices"
	"testing"
)

func TestGetIntersection(t *testing.T) {
	tests := []struct {
		name      string
		slice1    []int
		slice2    []int
		wantBool  bool
		wantSlice []int
	}{
		{
			name:      "пересекаются",
			slice1:    []int{65, 3, 58, 678, 64},
			slice2:    []int{64, 2, 3, 43, 3},
			wantSlice: []int{64, 3},
			wantBool:  true,
		},
		{
			name:      "не пересекаются",
			slice1:    []int{65, 3, 58, 678, 64},
			slice2:    []int{2, 43},
			wantSlice: []int{},
			wantBool:  false,
		},
		{
			name:      "не пересекаются",
			slice1:    []int{},
			slice2:    []int{2, 43},
			wantSlice: []int{},
			wantBool:  false,
		},
		{
			name:      "не пересекаются",
			slice1:    []int{65, 3, 58, 678, 64},
			slice2:    []int{},
			wantSlice: []int{},
			wantBool:  false,
		},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			got, ok := getIntersection(test.slice1, test.slice2)

			if ok != test.wantBool {
				t.Errorf("got: %v; want: %v", ok, test.wantBool)
			}

			if !slices.Equal(got, test.wantSlice) {
				t.Errorf("got: %v; want: %v", got, test.wantSlice)
			}
		})
	}
}
