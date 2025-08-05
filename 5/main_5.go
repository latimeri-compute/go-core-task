package main

func getIntersection[S ~[]T, T comparable](s1, s2 S) (S, bool) {
	uniqueElems := make(map[T]bool)
	for _, item := range s1 {
		uniqueElems[item] = true
	}

	intersectionSlice := make(S, 0)
	for _, item := range s2 {
		if uniqueElems[item] {
			intersectionSlice = append(intersectionSlice, item)
			delete(uniqueElems, item)
		}
	}

	return intersectionSlice, len(intersectionSlice) > 0
}
