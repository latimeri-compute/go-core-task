package main

func getUniqueSliceFromFirst[S ~[]T, T comparable](s1, s2 S) S {
	uniqueElems := make(map[T]bool)
	for _, item := range s2 {
		uniqueElems[item] = true
	}

	intersectionSlice := make(S, 0)
	for _, item := range s1 {
		if !uniqueElems[item] {
			intersectionSlice = append(intersectionSlice, item)
		}
	}

	return intersectionSlice
}
