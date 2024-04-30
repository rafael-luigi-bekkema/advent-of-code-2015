package main

import "golang.org/x/exp/constraints"

func pow[T constraints.Integer](v, n T) (result T) {
	if n == 0 {
		return 1
	}
	result = v
	for i := T(0); i < n-1; i++ {
		result *= v
	}
	return
}

func day17a(containers []int, liters int) int {
	count := 0
	var f func(perm []bool, left int)
	f = func(perm []bool, left int) {
		if len(perm) == len(containers) {
			if left == 0 {
				count++
			}
			return
		}
		nleft := left - containers[len(perm)]
		if nleft >= 0 {
			f(append(perm, true), nleft)
		}
		f(append(perm, false), left)
	}
	f(nil, liters)

	return count
}

func day17b(containers []int, liters int) int {
	var minused int
	var f func(perm []bool, left, used int)
	counts := make(map[int]int)
	f = func(perm []bool, left, used int) {
		if len(perm) == len(containers) {
			if left == 0 {
				if minused == 0 || used < minused {
					minused = used
				}
				if minused == 0 || used == minused {
					counts[used]++
				}
			}
			return
		}
		nleft := left - containers[len(perm)]
		if nleft >= 0 {
			f(append(perm, true), nleft, used+1)
		}
		f(append(perm, false), left, used)
	}
	f(nil, liters, 0)

	return counts[minused]
}
