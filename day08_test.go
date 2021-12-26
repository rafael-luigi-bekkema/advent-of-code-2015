package main

import (
	"testing"
)

func TestDay8a(t *testing.T) {
	tt := []struct {
		input  []string
		expect int
	}{
		{[]string{`""`, `"abc"`, `"aaa\"aaa"`, `"\x27"`}, 12},
		{Lines(8), 1371},
	}
	for _, tc := range tt {
		t.Run("", func(t *testing.T) {
			result := day8a(tc.input)
			if result != tc.expect {
				t.Fatalf("expected %d, got %d", tc.expect, result)
			}
		})
	}
}

func TestDay8b(t *testing.T) {
	tt := []struct {
		input  []string
		expect int
	}{
		{[]string{`""`, `"abc"`, `"aaa\"aaa"`, `"\x27"`}, 19},
		{Lines(8), 2117},
	}
	for _, tc := range tt {
		t.Run("", func(t *testing.T) {
			result := day8b(tc.input)
			if result != tc.expect {
				t.Fatalf("expected %d, got %d", tc.expect, result)
			}
		})
	}
}
