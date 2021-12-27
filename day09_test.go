package main

import "testing"

func TestDay9a(t *testing.T) {
	tt := []struct {
		input                []string
		expectMin, expectMax int
	}{
		{[]string{"London to Dublin = 464",
			"London to Belfast = 518",
			"Dublin to Belfast = 141"}, 605, 982},
		{Lines(9), 141, 736},
	}
	for _, tc := range tt {
		t.Run("", func(t *testing.T) {
			min, max := day9a(tc.input)
			if min != tc.expectMin {
				t.Fatalf("expected min %d, got %d", tc.expectMin, min)
			}
			if max != tc.expectMax {
				t.Fatalf("expected max %d, got %d", tc.expectMax, max)
			}
		})
	}
}
