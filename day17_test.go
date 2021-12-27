package main

import (
	"testing"
)

func intLines(day int) []int {
	return Map(atoi, Lines(day))
}

func TestDay17a(t *testing.T) {
	input := intLines(17)
	TestEqual(t, 4, day17a([]int{20, 15, 10, 5, 5}, 25), "25 l")
	TestEqual(t, 1304, day17a(input, 150), "150 l")
	TestEqual(t, 3, day17b([]int{20, 15, 10, 5, 5}, 25), "25 l (b)")
	TestEqual(t, 18, day17b(input, 150), "150 l (b)")
}
