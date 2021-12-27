package main

import "testing"

func TestDay19a(t *testing.T) {
	replacements := [][2]string{
		{"H", "HO"},
		{"H", "OH"},
		{"O", "HH"},
	}
	for _, tc := range []struct {
		molecule string
		expect   int
	}{{"HOH", 4}, {"HOHOHO", 7}} {
		result := day19a(replacements, tc.molecule)
		TestEqual(t, tc.expect, result, tc.molecule)
	}
}

func TestDay19a2(t *testing.T) {
	TestEqual(t, 535, day19a(day19input()))
}

func TestDay19b(t *testing.T) {
	replacements := [][2]string{
		{"e", "H"},
		{"e", "O"},
		{"H", "HO"},
		{"H", "OH"},
		{"O", "HH"},
	}
	for _, tc := range []struct {
		molecule string
		expect   int
	}{{"HOH", 3}, {"HOHOHO", 6}} {
		result := day19b(replacements, tc.molecule)
		TestEqual(t, tc.expect, result, tc.molecule)
	}
}

func TestDay19b2(t *testing.T) {
	TestEqual(t, 212, day19b(day19input()))
}
