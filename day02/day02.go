package day02

import (
	"bufio"
	"fmt"
	"io"
	"sort"

	"github.com/rafael-luigi-bekkema/advent-of-code-2015/utils"
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
		total += 2*s1 + 2*s2 + 2*s3 + utils.Min(s1, s2, s3)
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
		total += dim[0] * 2 + dim[1] * 2 + dim[0] * dim[1] * dim[2]
	}
	return total
}
