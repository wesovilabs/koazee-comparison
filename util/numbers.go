package util

import "math/rand"

func ArrayOfInt(min, max, cap int) []int {
	array := make([]int, cap)
	for i := 0; i < cap; i++ {
		array[i] = rand.Intn(max-min) + min
	}
	return array
}
