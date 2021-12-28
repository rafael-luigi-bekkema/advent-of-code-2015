package main

import (
	"math"
)

func sumDivisorsB(n int) int {
	var total int
	for i := 1; i <= int(math.Sqrt(float64(n))); i++ {
		if n%i == 0 {
			if n <= 50*i {
				total += i
			}
			if div := n / i; div != i && n <= 50*div {
				total += div
			}
		}
	}
	return total
}

func sumDivisors(n int) int {
	var total int
	for i := 1; i <= int(math.Sqrt(float64(n))); i++ {
		if n%i == 0 {
			total += i
			if div := n / i; div != i {
				total += div
			}
		}
	}
	return total
}

func day20a(minPresents int) int {
	house := 0
	for {
		house++
		presents := sumDivisors(house) * 10
		if presents >= minPresents {
			return house
		}
	}
}

func day20b(minPresents int) int {
	house := 0
	for {
		house++
		presents := sumDivisorsB(house) * 11
		if presents >= minPresents {
			return house
		}
	}
}
