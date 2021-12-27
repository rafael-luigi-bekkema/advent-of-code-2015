package main

import "testing"

func TestDay12a(t *testing.T) {
	tt := []struct {
		input  string
		expect float64
	}{
		{`[1,2,3]`, 6},
		{`{"a":2,"b":4}`, 6},
		{`[[[3]]]`, 3},
		{`{"a":{"b":4},"c":-1}`, 3},
		{`{"a":[-1,1]}`, 0},
		{`[-1,{"a":1}]`, 0},
		{`[]`, 0},
		{`{}`, 0},
		{Input(12), 191164},
	}
	for _, tc := range tt {
		t.Run("", func(t *testing.T) {
			result := day12a(tc.input, false)
			if result != tc.expect {
				t.Fatalf("expected %v, got %v", tc.expect, result)
			}
		})
	}
}

func TestDay12b(t *testing.T) {
	tt := []struct {
		input  string
		expect float64
	}{
		{`[1,2,3]`, 6},
		{`[1,{"c":"red","b":2},3]`, 4},
		{`{"d":"red","e":[1,2,3,4],"f":5}`, 0},
		{`[1,"red",5]`, 6},
		{Input(12), 87842},
	}
	for _, tc := range tt {
		t.Run("", func(t *testing.T) {
			result := day12a(tc.input, true)
			if result != tc.expect {
				t.Fatalf("expected %v, got %v", tc.expect, result)
			}
		})
	}
}
