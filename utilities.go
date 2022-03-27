package main

import (
	"math/rand"
	"sort"
	"time"
)

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

func sleepRandomTime(factor int) {
	rand.Seed(time.Now().UnixNano())
	time.Sleep(time.Duration(rand.Intn(factor*10)) * time.Millisecond)
}
