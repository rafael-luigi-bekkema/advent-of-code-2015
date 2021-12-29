package main

import "testing"

func TestDay25a(t *testing.T) {
	TestEqual(t, 18749137, day25a(1, 2))
	TestEqual(t, 27995004, day25a(6, 6))
	TestEqual(t, 28094349, day25a(5, 3))
	TestEqual(t, 2650453, day25a(2978, 3083))
}
