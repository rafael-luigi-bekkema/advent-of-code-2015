package main

import "testing"

func TestDay16a(t *testing.T) {
	TestEqual(t, 103, day16(Lines(16), false))
	TestEqual(t, 405, day16(Lines(16), true))
}
