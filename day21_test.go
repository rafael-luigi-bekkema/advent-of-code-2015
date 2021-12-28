package main

import "testing"

func TestDay21a(t *testing.T) {
	TestEqual(t, 78, day21a(false))
	TestEqual(t, 148, day21a(true))
}
