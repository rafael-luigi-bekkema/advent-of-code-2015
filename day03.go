package main

import "fmt"

type Coord struct {
	x, y int
}

func day3a(input string) int {
	pos := Coord{0, 0}
	visited := map[Coord]bool{
		pos: true,
	}
	for _, r := range input {
		switch r {
		case '<':
			pos.x--
		case '>':
			pos.x++
		case '^':
			pos.y++
		case 'v':
			pos.y--
		}
		visited[pos] = true
	}
	return len(visited)
}

func day3b(input string) int {
	santa := Coord{0, 0}
	robo := Coord{0, 0}
	current := &santa
	visited := map[Coord]bool{
		santa: true,
	}
	for _, r := range input {
		switch r {
		case '<':
			current.x--
		case '>':
			current.x++
		case '^':
			current.y++
		case 'v':
			current.y--
		}
		visited[*current] = true
		if current == &robo {
			current = &santa
		} else {
			current = &robo
		}
	}
	return len(visited)
}

func Day3() {
	fmt.Println("day 3a:", day3a(Input(3)))
	fmt.Println("day 3b:", day3b(Input(3)))
}
