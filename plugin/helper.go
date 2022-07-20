package plugin

import "golang.org/x/exp/constraints"

func Max[T constraints.Ordered](a T, b ...T) T {
	res := a
	for _, v := range b {
		if v > a {
			res = v
		}
	}
	return res
}

// Map like map in Python
func Map[T1, T2 any](slice []T1, f func(T1) T2) []T2 {
	result := make([]T2, len(slice))
	for i, v := range slice {
		result[i] = f(v)
	}
	return result
}
