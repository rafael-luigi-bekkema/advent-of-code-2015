package main

import (
	"reflect"
	"strings"
	"testing"
)

func TestDay7a(t *testing.T) {
	input := `123 -> x
456 -> y
x AND y -> d
x OR y -> e
x LSHIFT 2 -> f
y RSHIFT 2 -> g
NOT x -> h
NOT y -> i`
	expect := map[string]uint16{
		"d": 72,
		"e": 507,
		"f": 492,
		"g": 114,
		"h": 65412,
		"i": 65079,
		"x": 123,
		"y": 456,
	}
	result := day7a(strings.Split(input, "\n"), nil)
	if !reflect.DeepEqual(result, expect) {
		t.Fatalf("expected %v, got %v", expect, result)
	}
}

func TestDay7aFile(t *testing.T) {
	expect := uint16(46065)
	result := day7a(Lines(7), nil)
	if result["a"] != expect {
		t.Fatalf("expected %v, got %v", expect, result["a"])
	}
}

func TestDay7bFile(t *testing.T) {
	expect := uint16(14134)
	resultA := day7a(Lines(7), nil)["a"]
	result := day7a(Lines(7), map[string]uint16{"b": resultA})
	if result["a"] != expect {
		t.Fatalf("expected %v, got %v", expect, result["a"])
	}
}
