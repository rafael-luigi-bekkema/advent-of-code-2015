package main

import "fmt"

func day25a(row, col int) int {
	codes := map[[2]int]int{}
	v := false
	printTable := func() {
		if !v {
			return
		}
		const upto = 6
		fmt.Print("   |")
		for i := 1; i <= upto; i++ {
			fmt.Print("    ", i, "     ")
		}
		fmt.Print("\n---+")
		for i := 1; i <= upto; i++ {
			fmt.Print("---------+")
		}
		fmt.Println()
		for r := 0; r < upto; r++ {
			fmt.Print(" ", r+1, " |")
			for c := 0; c < upto; c++ {
				fmt.Printf(" %8d ", codes[[2]int{r, c}])
			}
			fmt.Println()
		}
		fmt.Println()
	}

	code := 20151125
	var srow int
outer:
	for {
		for scol := 0; scol <= srow; scol++ {
			ssrow := srow - scol
			if ssrow+1 == row && scol+1 == col {
				break outer
			}
			codes[[2]int{ssrow, scol}] = code
			code = (code * 252533) % 33554393
		}
		srow++
	}
	printTable()
	return code
}
