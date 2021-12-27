package main

import (
	"fmt"
	"strings"
)

func day16(input []string, b bool) int {
	compounds := map[string]int{
		"children":    3,
		"cats":        7,
		"samoyeds":    2,
		"pomeranians": 3,
		"akitas":      0,
		"vizslas":     0,
		"goldfish":    5,
		"trees":       3,
		"cars":        2,
		"perfumes":    1,
	}
	var gt, lt []string
	if b {
		gt = []string{"cats", "trees"}
		lt = []string{"pomeranians", "goldfish"}
	}
	var total int
	var theSue int
	for _, line := range input {
		sueS, stuff, _ := strings.Cut(line, ": ")
		var sue int
		fmt.Sscanf(sueS, "Sue %d", &sue)
		var score int
		for _, s := range strings.Split(stuff, ", ") {
			var name string
			var amount int
			fmt.Sscanf(strings.Replace(s, ":", "", 1), "%s %d", &name, &amount)
			if In(name, gt...) {
				if amount > compounds[name] {
					score++
				}
			} else if In(name, lt...) {
				if amount < compounds[name] {
					score++
				}
			} else {
				if compounds[name] == amount {
					score++
				}
			}
		}
		if score > total {
			total = score
			theSue = sue
		}
	}
	return theSue
}
