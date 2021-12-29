package main

import (
	"log"
)

type bufType uint

const (
	bShield bufType = iota
	bPoison
	bRecharge
)

type spellType string

const (
	magicMissile spellType = "Magic Missile"
	drain                  = "Drain"
	shield                 = "Shield"
	poison                 = "Poison"
	recharge               = "Recharge"
	none                   = "None"
)

type Buf struct {
	typ   bufType
	turns int
}

type Player struct {
	hp, armor, mana, dmg int
	spentMana            int
	extraArmor           int
	bufs                 []Buf
	casts                func() spellType
}

func printf(s string, extra ...any) {
	log.Printf(s, extra...)
}

func println(extra ...any) {
	log.Println(extra...)
}

func applyBufs(p *Player) {
	p.extraArmor = 0
	for i := len(p.bufs) - 1; i >= 0; i-- {
		b := &p.bufs[i]
		switch b.typ {
		case bShield:
			armor := 7
			// printf("Applying Shield. Add %d player armor.", armor)
			p.extraArmor = armor
		case bPoison:
			dmg := 3
			p.hp -= dmg
			// printf("Applying Poison. Does %d damage. Boss has %d hit points.", dmg, p.hp)
		case bRecharge:
			mana := 101
			p.mana += mana
			// printf("Applying Recharge. Gain %d mana. Player has %d mana.", mana, p.mana)
		}
		b.turns--
		if b.turns <= 0 {
			p.bufs[len(p.bufs)-1], p.bufs[i] = p.bufs[i], p.bufs[len(p.bufs)-1]
			p.bufs = p.bufs[:len(p.bufs)-1]
		}
		// printf(" Turns left %d.\n", b.turns)
	}
}

func cast(spell spellType, player, boss *Player) {
	mana := 0
	switch spell {
	case magicMissile:
		mana = 53
		dmg := 4
		// printf("Player casts %s. Mana cost %d. Does %d damage.\n",
		// 	spell, mana, dmg)

		player.mana -= mana
		boss.hp -= dmg
	case drain:
		mana = 73
		dmg := 2
		heals := 2
		// printf("Player casts %s. Mana cost %d. Does %d damage. Heals %d.\n",
		// 	spell, mana, dmg, heals)

		player.mana -= mana
		boss.hp -= dmg
		player.hp += heals
	case shield:
		mana = 113
		turns := 6
		// printf("Player casts %s. Mana cost %d. Lasts %d turns.\n",
		// 	spell, mana, turns)

		for _, b := range player.bufs {
			if b.typ == bShield {
				panic("already has shield")
			}
		}

		player.mana -= mana
		player.bufs = append(player.bufs, Buf{bShield, turns})
	case poison:
		mana = 173
		turns := 6
		// printf("Player casts %s. Mana cost %d. Lasts %d turns.\n",
		// 	spell, mana, turns)

		for _, b := range boss.bufs {
			if b.typ == bPoison {
				panic("already has poison")
			}
		}
		player.mana -= mana
		boss.bufs = append(boss.bufs, Buf{bPoison, turns})
	case recharge:
		mana = 229
		turns := 5
		// printf("Player casts %s. Mana cost %d. Lasts %d turns.\n",
		// 	spell, mana, turns)

		for _, b := range player.bufs {
			if b.typ == bRecharge {
				panic("already has recharge")
			}
		}
		player.mana -= mana
		player.bufs = append(player.bufs, Buf{bRecharge, turns})
	default:
		panic("unknown spell")
	}
	player.spentMana += mana
}

func spellPicker(seed int) (nextSpell func() spellType, closer func()) {
	c := make(chan spellType)
	next := make(chan bool)
	go func() {
		defer close(c)
		count := 1
		for range next {
			switch seed {
			case 0:
				switch count {
				case 1, 7:
					c <- poison
				case 3:
					c <- recharge
				case 11:
					c <- shield
				default:
					c <- magicMissile
				}
			case 1:
				switch count {
				case 1, 7, 13:
					c <- poison
				case 3, 9:
					c <- recharge
				case 5:
					c <- shield
				case 11:
					c <- drain
				default:
					c <- magicMissile
				}
			default:
				c <- magicMissile
			}
			count += 2
		}
	}()
	return func() spellType {
		next <- true
		a := <-c
		return a
	}, func() { close(next) }
}

func battle(player, boss Player, hard bool) (Player, Player) {
	playerTurn := true
	turn := 1
	for {
		if playerTurn {
			// println("-- Player turn --")
			// printf("- Player has %d hit points, %d armor, %d mana\n", player.hp, player.armor, player.mana)

			nextCast := player.casts()

			if hard {
				player.hp--
				if player.hp <= 0 {
					return player, boss
				}
			}

			applyBufs(&player)
			applyBufs(&boss)
			cast(nextCast, &player, &boss)
			printf("turn %d: hp player %d, boss %d, mana %d, %s\n", turn, player.hp, boss.hp, player.mana, nextCast)

			if boss.hp <= 0 || player.mana <= 0 {
				// printf("Boss has %d hit points. Boss is dead.\n", boss.hp)
				break
			}
		}
		if !playerTurn {
			// println("-- Boss turn --")
			// printf("- Player has %d hit points, %d armor, %d mana\n",
			// 	player.hp, player.armor, player.mana)
			// printf("- Boss has %d hit points\n", boss.hp)

			applyBufs(&player)
			applyBufs(&boss)
			// printf("turn %d: hp player %d, boss %d\n", turn, player.hp, boss.hp)
			if boss.hp <= 0 {
				// printf("Boss has %d hit points. Boss is dead.\n", boss.hp)
				break
			}

			dmg := boss.dmg - (player.armor + player.extraArmor)
			if dmg < 1 {
				dmg = 1
			}
			// printf("Boss attacks for %d - %d = %d damage.\n",
			// 	boss.dmg, player.armor+player.extraArmor, dmg)

			player.hp -= dmg
			// printf("turn %d: hp player %d, boss %d\n", turn, player.hp, boss.hp)
			if player.hp <= 0 {
				// printf("Player has %d hit points. Player is dead.\n", player.hp)
				break
			}
		}
		// println()

		if turn > 1000 {
			panic("Too many turns!")
		}

		playerTurn = !playerTurn
		turn++
	}
	// println()
	return player, boss
}

func day22(b bool) int {
	seed := 0
	if b {
		seed = 1
	}
	casts, close := spellPicker(seed)
	defer close()
	player := Player{hp: 50, mana: 500, casts: casts}
	boss := Player{hp: 51, dmg: 9}
	player, boss = battle(player, boss, b)

	if player.mana < 0 {
		printf("seed %d. player no mana (hp %d, %d) mana %d.", seed, player.hp, boss.hp, player.mana)
	} else if player.hp > 0 {
		printf("seed %d. player wins (hp %d, %d). spent mana: %d", seed, player.hp, boss.hp, player.spentMana)
		return player.spentMana
	} else {
		printf("seed %d. boss wins. (hp %d, %d).", seed, player.hp, boss.hp)
	}

	return 0
}
