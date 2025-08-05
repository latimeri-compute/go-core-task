package main

import (
	"math/rand/v2"
	"slices"
	"testing"
)

func TestMergeChannels(t *testing.T) {
	t.Run("слияние каналов", func(t *testing.T) {
		channels := make([]<-chan int, 0)

		want := make([]int, 0)
		for range 3 {
			arrRand := randomSlice(128, 10)
			want = append(want, arrRand...)

			ch := make(chan int, 10)
			go fillChannelRand(ch, arrRand)

			channels = append(channels, ch)
		}

		got := dumpChansToArray(mergeChannels(channels...))

		slices.Sort(got)
		slices.Sort(want)
		if !slices.Equal(got, want) {
			t.Errorf("got: %v; want: %v", got, want)
		}
	})
}

func randomSlice(maxN, length int) []int {
	sl := rand.Perm(maxN)
	res := make([]int, length)
	copy(res, sl)
	return res
}

func fillChannelRand[T any](ch chan<- T, arr []T) {
	for _, n := range arr {
		ch <- n
	}
	close(ch)
}

func dumpChansToArray[T any](channels ...<-chan T) []T {
	arr := make([]T, 0)
	for _, ch := range channels {
		for val := range ch {
			arr = append(arr, val)
		}
	}
	return arr
}
