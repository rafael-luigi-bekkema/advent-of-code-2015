package main

import "testing"

func TestDay23a(t *testing.T) {
	input := []string{"inc b", "jio b, +2", "tpl b", "inc b"}
	l := Lines(23)
	TestEqual(t, 2, day23a(input, false))
	TestEqual(t, 307, day23a(l, false))
	TestEqual(t, 160, day23a(l, true))
}
