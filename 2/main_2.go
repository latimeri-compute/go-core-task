package main

import (
	"fmt"
	"log/slog"
	"math/rand/v2"
	"os"

	"golang.org/x/exp/constraints"
)

func main() {
	log := slog.New(slog.NewTextHandler(os.Stdout, nil))
	sl := newOriginalSlice()
	log.Info(fmt.Sprintf("%v", sl))
	sl = sliceExample(sl)
	log.Info(fmt.Sprintf("%v", sl))
	sl = addElements(sl, 0)
	log.Info(fmt.Sprintf("%v", sl))
	copysl := copySlice(sl)
	log.Info(fmt.Sprintf("%v", copysl))
	log.Info(fmt.Sprintf("%v", removeElement(copysl, 2)))

}

func newOriginalSlice() []int {
	return rand.Perm(10)
}

func sliceExample[S ~[]E, E constraints.Integer](slice S) S {
	newSlice := make([]E, 0)
	for _, num := range slice {
		if num%2 == 0 {
			newSlice = append(newSlice, num)
		}
	}
	return newSlice
}

func addElements[S ~[]E, E comparable](slice S, num E) S {
	return append(slice, num)
}

func copySlice[S ~[]E, E comparable](slice S) S {
	newSlice := make(S, len(slice))
	copy(newSlice, slice)
	return newSlice
}

func removeElement[S ~[]E, E comparable, T constraints.Integer](slice S, index T) S {
	if int(index) >= len(slice) || index < 0 {
		return slice
	}
	newSlice := append(slice[:index], slice[index+1:]...)
	return newSlice
}
