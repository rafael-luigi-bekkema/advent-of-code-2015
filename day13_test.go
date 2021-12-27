package main

import "testing"

func TestDay13a(t *testing.T) {
	tt := []struct {
		input     []string
		includeMe bool
		expect    int
	}{
		{Lines(13 + 25), false, 330},
		{Lines(13), false, 733},
		{Lines(13), true, 725},
	}
	for _, tc := range tt {
		t.Run("", func(t *testing.T) {
			result := day13a(tc.input, tc.includeMe)
			if result != tc.expect {
				t.Fatalf("expected %v, got %v", tc.expect, result)
			}
		})
	}

}
