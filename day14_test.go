package main

import "testing"

func TestDay14a(t *testing.T) {
	input := []string{
		"Comet can fly 14 km/s for 10 seconds, but then must rest for 127 seconds.",
		"Dancer can fly 16 km/s for 11 seconds, but then must rest for 162 seconds.",
	}
	expect := 1120
	result := day14a(input, 1000)
	if result != expect {
		t.Fatalf("expected %v, got %v", expect, result)
	}
}

func TestDay14a2(t *testing.T) {
	input := Lines(14)
	expect := 2696
	result := day14a(input, 2503)
	if result != expect {
		t.Fatalf("expected %v, got %v", expect, result)
	}
}

func TestDay14b(t *testing.T) {
	input := []string{
		"Comet can fly 14 km/s for 10 seconds, but then must rest for 127 seconds.",
		"Dancer can fly 16 km/s for 11 seconds, but then must rest for 162 seconds.",
	}
	expect := 689
	result := day14b(input, 1000)
	if result != expect {
		t.Fatalf("expected %v, got %v", expect, result)
	}
}

func TestDay14b2(t *testing.T) {
	input := Lines(14)
	expect := 1084
	result := day14b(input, 2503)
	if result != expect {
		t.Fatalf("expected %v, got %v", expect, result)
	}
}
