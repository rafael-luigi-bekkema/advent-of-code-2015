package main

import (
	"fmt"
	"strings"
)

type Reindeers []Reindeer

func (r Reindeers) String() string {
	var s []string
	for _, r := range r {
		s = append(s, r.String())
	}
	return strings.Join(s, "\n")
}

type Reindeer struct {
	name                string
	speed, active, rest int
	score               int
}

func (r Reindeer) String() string {
	return fmt.Sprintf("%s has score %d", r.name, r.score)
}

func (r *Reindeer) distanceAfter(seconds int) int {
	cycle := r.active + r.rest
	nth := seconds / cycle
	s := seconds % cycle
	current := s
	if s > r.active {
		current = r.active
	}
	total := (nth*r.active + current) * r.speed
	return total
}

func day14a(input []string, seconds int) int {
	var max int
	for _, line := range input {
		var r Reindeer
		fmt.Sscanf(strings.TrimRight(line, "."),
			"%s can fly %d km/s for %d seconds, but then must rest for %d seconds",
			&r.name, &r.speed, &r.active, &r.rest)

		total := r.distanceAfter(seconds)
		if total > max {
			max = total
		}
	}
	return max
}

func day14b(input []string, seconds int) int {
	var reindeers Reindeers
	for _, line := range input {
		var r Reindeer
		fmt.Sscanf(strings.TrimRight(line, "."),
			"%s can fly %d km/s for %d seconds, but then must rest for %d seconds",
			&r.name, &r.speed, &r.active, &r.rest)
		reindeers = append(reindeers, r)
	}
	for second := 1; second <= seconds; second++ {
		var max int
		var topdogs []int
		for i := range reindeers {
			r := &reindeers[i]
			dist := r.distanceAfter(second)
			if dist > max {
				max = dist
				topdogs = []int{i}
			} else if dist == max {
				topdogs = append(topdogs, i)
			}
		}
		for _, td := range topdogs {
			reindeers[td].score++
		}
	}
	var top int
	for _, r := range reindeers {
		if r.score > top {
			top = r.score
		}
	}
	return top
}
