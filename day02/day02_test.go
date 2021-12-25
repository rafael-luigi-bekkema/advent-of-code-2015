package day02

import (
	"strings"
	"testing"

	"github.com/rafael-luigi-bekkema/advent-of-code-2015/utils"
)

func TestDay2a(t *testing.T) {
	tt := []struct {
		input  string
		expect int
	}{
		{input: "2x3x4", expect: 58},
		{input: "1x1x10", expect: 43},
	}
	for _, tc := range tt {
		t.Run(tc.input, func(t *testing.T) {
			result := day2a(strings.NewReader(tc.input))
			if result != tc.expect {
				t.Fatalf("expected %d, got %d", tc.expect, result)
			}
		})
	}
}

func TestDay2aInput(t *testing.T) {
	f := utils.InputReader(2)
	defer f.Close()
	result := day2a(f)
	expect := 1598415
	if result != expect {
		t.Fatalf("expected %d, got %d", expect, result)
	}
}

func TestDay2b(t *testing.T) {
	tt := []struct {
		input  string
		expect int
	}{
		{input: "2x3x4", expect: 34},
		{input: "1x1x10", expect: 14},
	}
	for _, tc := range tt {
		t.Run(tc.input, func(t *testing.T) {
			result := day2b(strings.NewReader(tc.input))
			if result != tc.expect {
				t.Fatalf("expected %d, got %d", tc.expect, result)
			}
		})
	}
}

func TestDay2bInput(t *testing.T) {
	f := utils.InputReader(2)
	defer f.Close()
	result := day2b(f)
	expect := 3812909
	if result != expect {
		t.Fatalf("expected %d, got %d", expect, result)
	}
}
