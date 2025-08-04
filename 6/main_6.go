package main

import "math/rand"

func randomInt(maxN int) int {
	return rand.Intn(maxN)
}

func randomSlice(size, maxN int) []int {
	sl := rand.Perm(maxN)
	res := make([]int, size)

	copy(res, sl)
	return res
}
