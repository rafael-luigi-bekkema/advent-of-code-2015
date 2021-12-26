package main

import (
	"bufio"
	"fmt"
	"io"
	"sort"
)

func day2a(input io.Reader) int {
	s := bufio.NewScanner(input)
	var total int
	for s.Scan() {
		var l, w, h int
		fmt.Sscanf(s.Text(), "%dx%dx%d", &l, &w, &h)
		s1 := l * w
		s2 := w * h
		s3 := h * l
		total += 2*s1 + 2*s2 + 2*s3 + Min(s1, s2, s3)
	}
	return total
}

func day2b(input io.Reader) int {
	s := bufio.NewScanner(input)
	var total int
	for s.Scan() {
		var dim [3]int
		fmt.Sscanf(s.Text(), "%dx%dx%d", &dim[0], &dim[1], &dim[2])
		sort.Ints(dim[:])
		total += dim[0]*2 + dim[1]*2 + dim[0]*dim[1]*dim[2]
	}
	return total
}

func Day2() {
	f := InputReader(2)
	defer f.Close()
	fmt.Println("day 2a:", day2a(f))
	f.Seek(0, 0)
	fmt.Println("day 2b:", day2b(f))
}
