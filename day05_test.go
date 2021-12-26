package main

import "testing"

func TestDay5a(t *testing.T) {
	tt := []struct {
		input  string
		expect bool
	}{
		{"ugknbfddgicrmopn", true},
		{"aaa", true},
		{"jchzalrnumimnmhp", false},
		{"haegwjzuvuyypxyu", false},
		{"dvszwmarrgswjxmb", false},
	}
	for _, tc := range tt {
		t.Run(tc.input, func(t *testing.T) {
			result := day5a(tc.input)
			if result != tc.expect {
				t.Fatalf("expected %v, got %v", tc.expect, result)
			}
		})
	}
}

func TestDay5count(t *testing.T) {
	tt := []struct {
		f  func(string)bool
		expect int
	}{
		{day5a, 255},
		{day5b, 55},
	}
	for _, tc := range tt {
		t.Run("", func(t *testing.T) {
			result := day5count(tc.f)
			if result != tc.expect {
				t.Fatalf("expected %v, got %v", tc.expect, result)
			}
		})
	}
}

func TestDay5b(t *testing.T) {
	tt := []struct {
		input  string
		expect bool
	}{
		{"qjhvhtzxzqqjkmpb", true},
		{"xxyxx", true},
		{"uurcxstgmygtbstg", false},
		{"ieodomkazucvgmuy", false},
	}
	for _, tc := range tt {

		t.Run(tc.input, func(t *testing.T) {
			result := day5b(tc.input)
			if result != tc.expect {
				t.Fatalf("expected %v, got %v", tc.expect, result)
			}
		})
	}
}
