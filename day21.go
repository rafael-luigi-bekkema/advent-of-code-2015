package main

import "fmt"

type Boss struct {
	hp, dmg, armor int
}

type Playa struct {
	items []Item
	hp    int
}

func (p *Playa) Armor() int {
	var total int
	for _, item := range p.items {
		total += item.armor
	}
	return total
}

func (p *Playa) Dmg() int {
	var total int
	for _, item := range p.items {
		total += item.dmg
	}
	return total
}

func (p *Playa) Cost() int {
	var total int
	for _, item := range p.items {
		total += item.cost
	}
	return total
}

type Item struct {
	name             string
	cost, dmg, armor int
}

func play(playa Playa, boss Boss, verbose bool) bool {
	playaDmg := playa.Dmg() - boss.armor
	if playaDmg <= 0 {
		playaDmg = 1
	}
	bossDmg := boss.dmg - playa.Armor()
	if bossDmg <= 0 {
		bossDmg = 1
	}
	for {
		boss.hp -= playaDmg
		if verbose {
			fmt.Printf("Player deals %d damage; boss goes down to %d hp\n", playaDmg, boss.hp)
		}
		if boss.hp <= 0 {
			if verbose {
				fmt.Println("boss is dead, player wins")
			}
			return true
		}

		playa.hp -= bossDmg
		if verbose {
			fmt.Printf("Boss deals %d damage; player goes down to %d hp\n", bossDmg, playa.hp)
		}
		if playa.hp <= 0 {
			if verbose {
				fmt.Println("player is dead, boss wins")
			}
			return false
		}
	}
}

func day21a(b bool) int {
	boss := Boss{104, 8, 1}
	playa := Playa{
		hp: 100,
		items: []Item{
			{"Shortsword", 10, 5, 0},
			{"Bandedmail", 75, 0, 4},
		},
	}
	weapons := []Item{
		{"Dagger", 8, 4, 0},
		{"Shortsword", 10, 5, 0},
		{"Warhammer", 25, 6, 0},
		{"Longsword", 40, 7, 0},
		{"Greataxe", 74, 8, 0},
	}
	armor := []Item{
		{"Naked", 0, 0, 0},
		{"Leather", 13, 0, 1},
		{"Chainmail", 31, 0, 2},
		{"Splintmail", 53, 0, 3},
		{"Bandedmail", 75, 0, 4},
		{"Platemail", 102, 0, 5},
	}
	rings := []Item{
		{"Damage +1", 25, 1, 0},
		{"Damage +2", 50, 2, 0},
		{"Damage +3", 100, 3, 0},
		{"Defense +1", 20, 0, 1},
		{"Defense +2", 40, 0, 2},
		{"Defense +3", 80, 0, 3},
	}

	ringf := func(n int) (items [][]Item) {
		if n == 0 {
			return [][]Item{nil}
		}
		for i, r := range rings {
			if n == 1 {
				items = append(items, []Item{r})
				continue
			}
			for _, r2 := range rings[i+1:] {
				items = append(items, []Item{r, r2})
			}
		}
		return
	}

	var mincost, maxcost int
	for _, w := range weapons {
		for _, a := range armor {
			for i := 0; i <= 2; i++ {
				for _, rs := range ringf(i) {
					playa.items = append(rs, a, w)
					cost := playa.Cost()
					if play(playa, boss, false) {
						if mincost == 0 || cost < mincost {
							mincost = cost
						}
					} else if cost > maxcost {
						maxcost = cost
					}
				}
			}
		}
	}
	if b {
		return maxcost
	}
	return mincost
}
