package main

import "testing"

func TestDay18a(t *testing.T) {
	testInput := `.#.#.#
...##.
#....#
..#...
#.#..#
####..`
	input := Input(18)
	TestEqual(t, 4, day18a(testInput, 4))
	TestEqual(t, 1061, day18a(input, 100))
	TestEqual(t, 17, day18b(testInput, 5))
	TestEqual(t, 1006, day18b(input, 100))
}
