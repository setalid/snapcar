package utils

func Map[T any, K any](items []T, fn func(item T) K) []K {
	newitems := make([]K, 0, len(items))
	for _, x := range items {
		newitems = append(newitems, fn(x))
	}
	return newitems
}
