package main

import "fmt"

func day8a(input []string) int {
	var code, mem int
	for _, line := range input {
		code += len(line)
		for i := 0; i < len(line); i++ {
			if line[i] == '\\' {
				switch line[i+1] {
				case '"', '\\':
					i++
				case 'x':
					i += 3
				default:
					panic("unknown escape sequence")
				}
				mem++
				continue
			}
			if i == 0 || i == len(line)-1 {
				if line[i] != '"' {
					panic(`expected "`)
				}
				continue
			}
			mem++
		}
	}
	return code - mem
}

func day8b(input []string) int {
	var code, total int
	for _, line := range input {
		code += len(line)
		total += 2
		for i := 0; i < len(line); i++ {
			switch line[i] {
			case '"', '\\':
				total += 2
			default:
				total += 1
			}
		}
	}
	return total - code
}

func Day8() {
	input := Lines(8)
	fmt.Println("day 8a:", day8a(input))
	fmt.Println("day 8b:", day8b(input))
}
