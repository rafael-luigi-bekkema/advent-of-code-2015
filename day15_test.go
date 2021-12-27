package main

import "testing"

func TestDay15(t *testing.T) {
	input := []string{
		"Butterscotch: capacity -1, durability -2, flavor 6, texture 3, calories 8",
		"Cinnamon: capacity 2, durability 3, flavor -2, texture -1, calories 3",
	}
	TestEqual(t, 62842880, day15a(input), "a test")
	TestEqual(t, 13882464, day15a(Lines(15)), "a answer")
	TestEqual(t, 57600000, day15b(input), "b test")
	TestEqual(t, 11171160, day15b(Lines(15)), "b answer")
}
