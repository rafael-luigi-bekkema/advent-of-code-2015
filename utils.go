package main

import (
	"bufio"
	"constraints"
	"fmt"
	"io"
	"os"
	"path/filepath"
)

const inputDir = "input"

func Must[T any](p T, err error) T {
	if err != nil {
		panic(err)
	}
	return p
}

func InputReader(day int) io.ReadSeekCloser {
	return Must(os.Open(inputPath(day)))
}

func InputScanner(day int) (s *bufio.Scanner, close func() error) {
	f := InputReader(day)
	s = bufio.NewScanner(f)
	return s, f.Close
}

func Lines(day int) []string {
	b, close := InputScanner(day)
	defer close()
	var lines []string
	for b.Scan() {
		lines = append(lines, b.Text())
	}
	return lines
}

func inputPath(day int) string {
	return Must(filepath.Abs(fmt.Sprintf("%s/day%02d.txt", inputDir, day)))
}

func Input(day int) string {
	data := Must(os.ReadFile(inputPath(day)))
	return string(data)
}

func Min(nums ...int) (n int) {
	for i, num := range nums {
		if i == 0 || num < n {
			n = num
		}
	}
	return
}

func In[T comparable](items ...T) bool {
	for _, i := range items[1:] {
		if i == items[0] {
			return true
		}
	}
	return false
}

func BoolToInt(v bool) int {
	if v {
		return 1
	}
	return 0
}

func Add[T constraints.Ordered](a, b T) T {
	return a + b
}

func Sum[T constraints.Ordered](list []T) T {
	var s T
	return Reduce(Add[T], s, list)
}

func Map[T, U any](f func(T) U, list []T) []U {
	result := make([]U, len(list))
	for _, i := range list {
		result = append(result, f(i))
	}
	return result
}

func Reduce[A, T any](f func(A, T) A, acc A, list []T) A {
	for _, l := range list {
		acc = f(acc, l)
	}
	return acc
}
