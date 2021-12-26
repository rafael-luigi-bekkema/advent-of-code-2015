package main

import (
	"fmt"
)

func day5a(input string) bool {
	var vowels int
	var repeat bool
	for i, r := range input {
		if In(r, 'a', 'e', 'i', 'o', 'u') {
			vowels++
		}
		if i > 0 && rune(input[i-1]) == r {
			repeat = true
		}
		if i > 0 && In(input[i-1:i+1], "ab", "cd", "pq", "xy") {
			return false
		}
	}
	return vowels >= 3 && repeat
}

func day5count(f func(string) bool) int {
	s, close := InputScanner(5)
	defer close()
	var count int
	for s.Scan() {
		if f(s.Text()) {
			count++
		}
	}
	return count
}

func day5b(input string) bool {
	var repeatD, repeatC bool
	for i := range input {
		if i > 0 {
			c := input[i-1 : i+1]
			for i2 := i + 2; i2 < len(input); i2++ {
				if input[i2-1:i2+1] == c {
					repeatD = true
					break
				}
			}
		}
		if i >= 2 {
			if input[i] == input[i-2] {
				repeatC = true
			}
		}
	}
	return repeatD && repeatC
}

func Day5() {
	fmt.Println("day 5a:", day5count(day5a))
	fmt.Println("day 5b:", day5count(day5b))
}
