package main

import (
	"fmt"
)

func day6a(input []string) int {
	var lights [1_000_000]bool
	width := 1000
	for _, line := range input {
		var minx, miny, maxx, maxy int
		var action string
		fmt.Sscanf(line[5:], "%s %d,%d through %d,%d", &action, &minx, &miny, &maxx, &maxy)

		var f func(v *bool)
		switch action {
		case "on":
			f = func(v *bool) { *v = true }
		case "off":
			f = func(v *bool) { *v = false }
		default: // toggle
			f = func(v *bool) { *v = !*v }
		}
		for y := miny; y <= maxy; y++ {
			for x := minx; x <= maxx; x++ {
				f(&lights[y*width+x])
			}
		}
	}
	return Reduce(func(acc int, v bool) int {
		return acc + BoolToInt(v)
	}, 0, lights[:])
}

func day6b(input []string) int {
	var lights [1_000_000]int
	width := 1000
	for _, line := range input {
		var minx, miny, maxx, maxy int
		var action string
		fmt.Sscanf(line[5:], "%s %d,%d through %d,%d", &action, &minx, &miny, &maxx, &maxy)

		var f func(v *int)
		switch action {
		case "on":
			f = func(v *int) { *v++ }
		case "off":
			f = func(v *int) {
				if *v > 0 {
					*v--
				}
			}
		default: // toggle
			f = func(v *int) { *v += 2 }
		}
		for y := miny; y <= maxy; y++ {
			for x := minx; x <= maxx; x++ {
				f(&lights[y*width+x])
			}
		}
	}
	return Sum(lights[:])
}
