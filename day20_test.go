package main

import "testing"

func TestDay20one(t *testing.T) {
	TestEqual(t, 4, day20a(70))
	TestEqual(t, 6, day20a(120))
	TestEqual(t, 6, day20a(80))
	TestEqual(t, 8, day20a(130))
}

func TestDay20two(t *testing.T) {
	if testing.Short() {
		t.Skip("skip slow-ass test")
	}
	TestEqual(t, 786240, day20a(34_000_000))
	TestEqual(t, 831600, day20b(34_000_000))
}
