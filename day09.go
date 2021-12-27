package main

import (
	"fmt"
)

func day9a(input []string) (int, int) {
	var nodes []string
	edges := map[string]map[string]float64{}
	for _, line := range input {
		var from, to string
		var dist float64
		fmt.Sscanf(line, "%s to %s = %f", &from, &to, &dist)
		if !In(from, nodes...) {
			nodes = append(nodes, from)
		}
		if !In(to, nodes...) {
			nodes = append(nodes, to)
		}
		if edges[from] == nil {
			edges[from] = make(map[string]float64)
		}
		if edges[to] == nil {
			edges[to] = make(map[string]float64)
		}
		edges[from][to] = dist
		edges[to][from] = dist
	}

	route := func(combi []int, v bool) float64 {
		var dist float64
		for idx, i := range combi {
			if idx > 0 {
				source := nodes[combi[idx-1]]
				d := edges[source][nodes[i]]
				dist += d
				if v {
					fmt.Print(" (", d, ") -> ")
				}
			}
		}
		if v {
			fmt.Println(" =", dist)
		}
		return dist
	}
	var mindist, maxdist float64
	Combos(Range(0, len(nodes)-1), func (c []int) {
		dist := route(c, false)
		if mindist == 0 || dist < mindist {
			mindist = dist
		}
		if dist > maxdist {
			maxdist = dist
		}
	})
	return int(mindist), int(maxdist)
}
