package main

import "sort"

func sortMap[T ~int, J any](input map[T]J) []T {
	var keys = make([]T, len(input))
	for k := range input {
		keys[k] = k
	}
	sort.Slice(keys, func(i, j int) bool {
		return keys[i] < keys[j]
	})
	return keys
}
