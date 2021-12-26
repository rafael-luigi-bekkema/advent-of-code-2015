package main

import (
	"fmt"
	"strings"
)

func day7a(input []string, override map[string]uint16) map[string]uint16 {
	wires := map[string]uint16{}
	iOrVal := func(name string) (uint16, bool) {
		if '0' <= name[0] && name[0] <= '9' {
			return uint16(atoi(name)), true
		}
		v, ok := wires[name]
		return v, ok
	}
	for i := len(input) - 1; i >= 0; i-- {
		parts := strings.Split(input[i], " ")
		if v, ok := override[parts[len(parts)-1]]; ok {
			parts = []string{fmt.Sprint(v), "->", parts[len(parts)-1]}
		}
		switch parts[1] {
		case "->":
			v, ok := iOrVal(parts[0])
			if !ok {
				continue
			}
			wires[parts[2]] = v
		case "AND":
			v, ok := iOrVal(parts[0])
			v2, ok2 := iOrVal(parts[2])
			if !ok || !ok2 {
				continue
			}
			wires[parts[4]] = v & v2
		case "OR":
			v, ok := iOrVal(parts[0])
			v2, ok2 := iOrVal(parts[2])
			if !ok || !ok2 {
				continue
			}
			wires[parts[4]] = v | v2
		case "LSHIFT":
			v, ok := iOrVal(parts[0])
			v2, ok2 := iOrVal(parts[2])
			if !ok || !ok2 {
				continue
			}
			wires[parts[4]] = v << v2
		case "RSHIFT":
			v, ok := iOrVal(parts[0])
			v2, ok2 := iOrVal(parts[2])
			if !ok || !ok2 {
				continue
			}
			wires[parts[4]] = v >> v2
		default:
			if parts[0] == "NOT" {
				v, ok := wires[parts[1]]
				if !ok {
					continue
				}
				wires[parts[3]] = ^v
				break
			}
			panic("unknown instruction: " + input[i])

		}

		input[i], input[len(input)-1] = input[len(input)-1], input[i]
		input = input[:len(input)-1]
		if len(input) == 0 {
			break
		}
		i = len(input)
	}
	return wires
}

func Day7() {
	input := Lines(7)
	result := day7a(input, nil)
	fmt.Println("day 7a: wire a:", result["a"])
	result = day7a(input, map[string]uint16{"b": result["a"]})
	fmt.Println("day 7b: wire a:", result["a"])
}
