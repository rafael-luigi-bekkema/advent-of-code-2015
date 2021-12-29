package main

import (
	"fmt"
	"strings"
)

func day23a(input []string, b bool) uint {
	reg := map[string]uint{}
	if b {
		reg["a"] = 1
	}
	for i := 0; i < len(input); i++ {
		ins, rest, _ := strings.Cut(input[i], " ")
		switch ins {
		case "hlf":
			reg[rest] /= 2
		case "tpl":
			reg[rest] *= 3
		case "inc":
			reg[rest]++
		case "jmp":
			n := atoi(rest)
			i += n - 1
		case "jie":
			r, offset, _ := strings.Cut(rest, ", ")
			if reg[r]%2 == 0 {
				i += atoi(offset) - 1
				if i < -1 {
					panic(fmt.Sprintf("jump to offset: %d", i))
				}
			}
		case "jio":
			r, offset, _ := strings.Cut(rest, ", ")
			if reg[r] == 1 {
				i += atoi(offset) - 1
				if i < -1 {
					panic(fmt.Sprintf("jump to offset: %d", i))
				}
			}
		}
	}
	return reg["b"]
}
