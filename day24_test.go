package main

import "testing"

func TestDay24a(t *testing.T) {
	TestEqual(t, 99, day24a([]int{1, 2, 3, 4, 5, 7, 8, 9, 10, 11}, false))
	TestEqual(t, 10439961859, day24a(intLines(24), false))
	TestEqual(t, 44, day24a([]int{1, 2, 3, 4, 5, 7, 8, 9, 10, 11}, true))
	TestEqual(t, 72050269, day24a(intLines(24), true))
}
