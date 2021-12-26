package main

import "testing"

func TestDay3a(t *testing.T) {
	tt := []struct {
		input  string
		expect int
	}{
		{"^v^v^v^v^v", 2},
		{">", 2},
		{"^>v<", 4},
		{Input(3), 2565},
	}
	for _, tc := range tt {
		t.Run("", func(t *testing.T) {
			result := day3a(tc.input)
			if result != tc.expect {
				t.Fatalf("expected %v, got %v", tc.expect, result)
			}
		})
	}
}

func TestDay3b(t *testing.T) {
	tt := []struct {
		input  string
		expect int
	}{
		{"^v", 3},
		{"^>v<", 3},
		{"^v^v^v^v^v", 11},
		{Input(3), 2639},
	}
	for _, tc := range tt {
		t.Run("", func(t *testing.T) {
			result := day3b(tc.input)
			if result != tc.expect {
				t.Fatalf("expected %v, got %v", tc.expect, result)
			}
		})
	}
}
