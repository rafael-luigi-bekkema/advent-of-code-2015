package main

import "testing"

func TestDay10a(t *testing.T) {
	tt := []struct {
		input  string
		expect string
	}{
		{"1", "11"},
		{"11", "21"},
		{"21", "1211"},
		{"1211", "111221"},
		{"111221", "312211"},
	}
	for _, tc := range tt {
		t.Run("", func(t *testing.T) {
			result := day10a(tc.input, 1)
			if result != tc.expect {
				t.Fatalf("expected %q, got %q", tc.expect, result)
			}
		})
	}
}

func TestDay10a2(t *testing.T) {
	input := "3113322113"
	result := day10a(input, 40)
	expect := 329356
	if len(result) != expect {
		t.Fatalf("expected %d, got %d", expect, len(result))
	}
}

func TestDay10b(t *testing.T) {
	input := "3113322113"
	result := day10a(input, 50)
	expect := 4666278
	if len(result) != expect {
		t.Fatalf("expected %d, got %d", expect, len(result))
	}
}
