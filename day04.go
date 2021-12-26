package main

import (
	"crypto/md5"
	"fmt"
	"strings"
)

func day4a(input string, zeroes int) int {
	prefix := strings.Repeat("0", zeroes)

	for i := 1; i < 5000_000; i++ {
		s := fmt.Sprintf("%s%d", input, i)
		h := fmt.Sprintf("%x", md5.Sum([]byte(s)))
		if strings.HasPrefix(h, prefix) {
			return i
		}
	}
	panic("not found")
}

func Day4() {
	input := "bgvyzdsv"
	fmt.Println("day 4a:", day4a(input, 5))
	fmt.Println("day 4b:", day4a(input, 6))
}
