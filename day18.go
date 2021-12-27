package main

import (
	"strings"
)

func day18a(input string, steps int) int {
	return day18(input, steps, false)
}

func day18b(input string, steps int) int {
	return day18(input, steps, true)
}

func day18(input string, steps int, b bool) int {
	width := strings.Index(input, "\n")
	grid := []byte(strings.ReplaceAll(input, "\n", ""))
	height := len(grid) / width
	corners := []int{0, width - 1, (height - 1) * width, len(grid) - 1}
	if b {
		for _, i := range corners {
			grid[i] = '#'
		}
	}
	newgrid := make([]byte, len(grid))
	var count int
	for step := 1; step <= steps; step++ {
		count = 0
		for i, c := range grid {
			if b && In(i, corners...) {
				count++
				newgrid[i] = '#'
				continue
			}
			row := i / width
			col := i % width
			var ncount int
			for i := 0; i < 9; i++ {
				if i == 4 {
					continue
				}
				drow := row + (i/3 - 1)
				dcol := col + (i%3 - 1)
				if drow < 0 || drow >= height || dcol < 0 || dcol >= width {
					continue
				}
				if grid[drow*width+dcol] == '#' {
					ncount++
				}
			}
			if c == '#' {
				if 2 <= ncount && ncount <= 3 {
					count++
					newgrid[i] = '#'
				} else {
					newgrid[i] = '.'
				}
			} else {
				if ncount == 3 {
					count++
					newgrid[i] = '#'
				} else {
					newgrid[i] = '.'
				}
			}
		}
		grid, newgrid = newgrid, grid
	}

	return count
}
