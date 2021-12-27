package main

import (
	"fmt"
	"strings"
)

func day13a(input []string, includeMe bool) int {
	var s Set[string]
	scores := map[[2]string]int{}
	for _, line := range input {
		var p1, p2, gain string
		var amount int
		fmt.Sscanf(strings.TrimRight(line, "."),
			"%s would %s %d happiness units by sitting next to %s",
			&p1, &gain, &amount, &p2)
		if gain == "lose" {
			amount *= -1
		}
		s.Add(p1, p2)
		scores[[2]string{p1, p2}] = amount
	}
	calc := func(combo []string) int {
		var total int
		for i, p := range combo {
			var pp, np string
			if i == 0 {
				pp = combo[len(combo)-1]
			} else {
				pp = combo[i-1]
			}
			if i == len(combo)-1 {
				np = combo[0]
			} else {
				np = combo[i+1]
			}
			total += scores[[2]string{p, pp}]
			total += scores[[2]string{p, np}]
		}
		return total
	}
	var max int
	var first string
	vals := s.Values()
	if includeMe {
		first = "me"
	} else {
		first = vals[0]
		vals = vals[1:]
	}
	Combos(vals, func(combo []string) {
		c := calc(append(combo, first))
		if c > max {
			max = c
		}
	})
	return max
}
