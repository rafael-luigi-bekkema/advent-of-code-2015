package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"path/filepath"
	"strconv"
	"testing"

	"golang.org/x/exp/constraints"
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
	var postfix string
	if day > 25 {
		n := day / 25
		day = (day-1)%25 + 1
		var s string
		if n > 1 {
			s = fmt.Sprint(n)
		}
		postfix = fmt.Sprintf("_test%s", s)
	}
	return Must(filepath.Abs(fmt.Sprintf("%s/day%02d%s.txt", inputDir, day, postfix)))
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

func In[T comparable](item T, items ...T) bool {
	for _, i := range items {
		if i == item {
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

func Mul[T constraints.Integer | constraints.Float](a, b T) T {
	return a * b
}

func Sum[T constraints.Ordered](list []T) T {
	var s T
	return Reduce(Add[T], s, list)
}

func Prod[T constraints.Integer | constraints.Float](list []T) T {
	return Reduce(Mul[T], list[0], list[1:])
}

func Map[T, U any](f func(T) U, list []T) []U {
	result := make([]U, len(list))
	for idx, i := range list {
		result[idx] = f(i)
	}
	return result
}

func Reduce[A, T any](f func(A, T) A, acc A, list []T) A {
	for _, l := range list {
		acc = f(acc, l)
	}
	return acc
}

func atoi(v string) int {
	return Must(strconv.Atoi(v))
}

func Combos[T comparable](items []T, f func([]T)) {
	var combos func(combo []T)
	combos = func(combo []T) {
		if len(combo) == len(items) {
			f(combo)
			return
		}
		for _, i := range items {
			if !In(i, combo...) {
				combos(append(combo, i))
			}
		}
	}
	combos(nil)
}

func Range[T constraints.Integer](min, max T) (result []T) {
	for i := min; i <= max; i++ {
		result = append(result, i)
	}
	return
}

type Set[T comparable] struct {
	keys   map[T]struct{}
	values []T
}

func (s *Set[T]) String() string {
	return fmt.Sprint(s.values)
}

func NewSet[T comparable](items ...T) *Set[T] {
	var s Set[T]
	for _, item := range items {
		s.Add(item)
	}
	return &s
}

func (s *Set[T]) Add(items ...T) {
	for _, item := range items {
		if _, ok := s.keys[item]; ok {
			continue
		}
		if s.keys == nil {
			s.keys = make(map[T]struct{})
		}
		s.keys[item] = struct{}{}
		s.values = append(s.values, item)
	}
}

func (s *Set[T]) Len() int {
	return len(s.values)
}

func (s *Set[T]) Values() []T {
	return s.values
}

func TestEqual[T comparable](t *testing.T, expect, result T, message ...string) {
	t.Helper()
	var msg string
	if len(message) > 0 {
		msg = message[0]
	}
	t.Run(msg, func(t *testing.T) {
		t.Helper()
		if result != expect {
			t.Fatalf("expected %v, got %v", expect, result)
		}
	})
}
