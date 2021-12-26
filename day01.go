package main

import (
	"fmt"
)

func day1a(input string) int {
	var n, count int
	for _, r := range input {
		count++
		if r == '(' {
			n++
		} else {
			n--
		}
	}
	return n
}

func day1b(input string) int {
	var n, count int
	for _, r := range input {
		count++
		if r == '(' {
			n++
		} else {
			n--
		}
		if n < 0 {
			return count
		}
	}
	panic("should not happen")
}

func Day1() {
	fmt.Println("day 1a:", day1a(Input(1)))
	fmt.Println("day 1b:", day1a(Input(1)))
}
