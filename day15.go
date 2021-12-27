package main

import (
	"fmt"
	"strings"
)

type Ingredient struct {
	name                                            string
	capacity, durability, flavor, texture, calories int
}

func day15a(input []string) int {
	return day15(input, -1)
}

func day15b(input []string) int {
	return day15(input, 500)
}

func day15(input []string, calories int) int {
	teaspoons := 100
	ingredients := make([]Ingredient, len(input))
	for idx, line := range input {
		var i Ingredient
		fmt.Sscanf(strings.Replace(line, ":", "", 1),
			"%s capacity %d, durability %d, flavor %d, texture %d, calories %d",
			&i.name, &i.capacity, &i.durability, &i.flavor, &i.texture, &i.calories)
		ingredients[idx] = i
	}

	minzero := func(amount int) int {
		if amount < 0 {
			return 0
		}
		return amount
	}
	v := false
	score := func(amounts []int) (int, int) {
		var tcap, tdur, tfl, ttex int
		var cal int
		for i, amount := range amounts {
			ing := ingredients[i]
			tcap += ing.capacity * amount
			tdur += ing.durability * amount
			tfl += ing.flavor * amount
			ttex += ing.texture * amount
			cal += ing.calories * amount
			if v {
				fmt.Print(amount, " x ", ingredients[i].name, " | ")
			}
		}
		total := minzero(tcap) * minzero(tdur) * minzero(tfl) * minzero(ttex)
		if v {
			fmt.Println(total)
		}
		return total, cal
	}

	var maxscore int

	var f func(amounts []int, remaining int)
	f = func(amounts []int, remaining int) {
		if remaining == 0 {
			s, cal := score(amounts)
			if calories > 0 && cal != calories {
				return
			}
			if s > maxscore {
				maxscore = s
			}
			return
		}
		if len(amounts) == len(ingredients)-1 {
			f(append(amounts, remaining), 0)
			return
		}
		for ts := 0; ts <= remaining; ts++ {
			f(append(amounts, ts), remaining-ts)
		}
	}
	f(nil, teaspoons)
	return maxscore
}
