package main

func Filter[T any](arr []T, predicate func(T) bool) []T {
	filtered := make([]T, 0, len(arr))
	for _, v := range arr {
		if predicate(v) {
			filtered = append(filtered, v)
		}
	}

	return filtered
}
