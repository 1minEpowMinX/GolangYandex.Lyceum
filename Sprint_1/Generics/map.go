package main

func Map[T, U any](arr []T, transform func(T) U) []U {
	transformed := make([]U, 0, len(arr))
	for _, v := range arr {
		transformed = append(transformed, transform(v))
	}

	return transformed
}
