package main

import (
	"fmt"
	"sort"
	"strings"
)

func day19input() (replacements [][2]string, molecule string) {
	b, close := InputScanner(19)
	defer close()

	for b.Scan() {
		line := b.Text()
		if line == "" {
			b.Scan()
			molecule = b.Text()
			break
		}
		f, t, _ := strings.Cut(line, " => ")
		replacements = append(replacements, [2]string{f, t})
	}
	return
}

func day19a(replacements [][2]string, molecule string) int {
	set := NewSet[string]()
	for i := 0; i < len(molecule); i++ {
		for _, r := range replacements {
			l := len(r[0])
			if i+l > len(molecule) {
				continue
			}
			if r[0] == molecule[i:i+l] {
				s := fmt.Sprint(molecule[:i], r[1], molecule[i+l:])
				set.Add(s)
			}
		}
	}
	return set.Len()
}

func day19b(replacements [][2]string, molecule string) int {
	var minsteps int

	sort.Slice(replacements, func(i, j int) bool {
		return len(replacements[i][1]) > len(replacements[j][1])
	})

	mem := map[string]struct{}{}

	var f func(string, int, []string)
	f = func(molecule string, steps int, v []string) {
		if molecule == "e" {
			if minsteps == 0 || steps < minsteps {
				minsteps = steps
			}
			mem[molecule] = struct{}{}
			return
		}
		if minsteps != 0 && steps >= minsteps {
			return
		}
		if _, ok := mem[molecule]; ok {
			return
		}
		mem[molecule] = struct{}{}
	outer:
		for _, r := range replacements {
			for i := 0; i < len(molecule); i++ {
				l := len(r[1])
				if i+l > len(molecule) {
					continue
				}
				if r[1] == molecule[i:i+l] {
					// fmt.Println("r", r[1], r[0])
					s := fmt.Sprint(molecule[:i], r[0], molecule[i+l:])
					f(s, steps+1, append(v, molecule))
					break outer // ?
					// fmt.Println(molecule, "->", s, r, steps, len(molecule), len(s), "mem", ok)
				}
			}
		}
	}
	f(molecule, 0, nil)
	return minsteps
}
