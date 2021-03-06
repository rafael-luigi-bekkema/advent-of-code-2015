package main

import (
	"fmt"
	"testing"
)

func TestDay01a(t *testing.T) {
	tt := []struct {
		input  string
		expect int
	}{
		{"(())", 0},
		{"()()", 0},
		{"(()(()(", 3},
		{")())())", -3},
		{Input(1), 232},
	}
	for i, tc := range tt {
		t.Run(fmt.Sprintf("case %d", i+1), func(t *testing.T) {
			result := day1a(tc.input)
			if result != tc.expect {
				t.Fatalf("expected %d, got %d", tc.expect, result)
			}
		})
	}
}

func TestDay01b(t *testing.T) {
	tt := []struct {
		input  string
		expect int
	}{
		{")", 1},
		{"()())", 5},
		{Input(1), 1783},
	}
	for i, tc := range tt {
		t.Run(fmt.Sprintf("case %d", i+1), func(t *testing.T) {
			result := day1b(tc.input)
			if result != tc.expect {
				t.Fatalf("expected %d, got %d", tc.expect, result)
			}
		})
	}
}
