package main

import (
	"strings"
)

func day7a(input []string) map[string]uint16 {
	wires := map[string]uint16{}
	iOrVal := func(name string) (uint16, bool) {
		if '0' <= name[0] && name[0] <= '9' {
			return uint16(atoi(name)), true
		}
		v, ok := wires[name]
		return v, ok
	}
	for {
		var changed bool
		for i := len(input) - 1; i >= 0; i-- {
			parts := strings.Split(input[i], " ")
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

			changed = true
			input[i], input[len(input)-1] = input[len(input)-1], input[i]
			input = input[:len(input)-1]
			if len(input) == 0 {
				break
			}
		}
		if !changed {
			panic("nothing left to do")
		}
		if len(input) == 0 {
			break
		}
	}
	return wires
}
