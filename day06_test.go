package main

import (
	"testing"
)

func TestDay6a(t *testing.T) {
	tt := []struct {
		input  []string
		expect int
	}{
		{[]string{"turn on 0,0 through 999,999"}, 1_000_000},
		{[]string{"turn on 0,0 through 999,999", "toggle 0,0 through 999,0"}, 1_000_000 - 1_000},
		{[]string{"turn on 0,0 through 999,999", "turn off 499,499 through 500,500"}, 1_000_000 - 4},
		{Lines(6), 569999},
	}
	for _, tc := range tt {
		t.Run("", func(t *testing.T) {
			result := day6a(tc.input)
			if result != tc.expect {
				t.Fatalf("expected %v, got %v", tc.expect, result)
			}
		})
	}
}

func TestDay6b(t *testing.T) {
	tt := []struct {
		input  []string
		expect int
	}{
		{[]string{"turn on 0,0 through 0,0"}, 1},
		{[]string{"toggle 0,0 through 999,999"}, 2_000_000},
		{Lines(6), 17836115},
	}
	for _, tc := range tt {
		t.Run("", func(t *testing.T) {
			result := day6b(tc.input)
			if result != tc.expect {
				t.Fatalf("expected %v, got %v", tc.expect, result)
			}
		})
	}
}
