package main

import "testing"

func TestDay4(t *testing.T) {
	    if testing.Short() {
		t.Skip("skipping test in short mode.")
	    }
	tt := []struct {
		input  string
		zeroes int
		expect int
	}{
		{input: "abcdef", zeroes: 5, expect: 609043},
		{input: "pqrstuv", zeroes: 5, expect: 1048970},
		{input: "bgvyzdsv", zeroes: 5, expect: 254575},
		{input: "bgvyzdsv", zeroes: 6, expect: 1038736},
	}
	for _, tc := range tt {

		t.Run("", func(t *testing.T) {
			result := day4a(tc.input, tc.zeroes)
			if result != tc.expect {
				t.Fatalf("expected %d, got %d", tc.expect, result)
			}
		})
	}
}
