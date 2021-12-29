package main

import (
	"io"
	"log"
	"testing"
)

func TestDay22Battle1(t *testing.T) {
	log.SetOutput(io.Discard)
	c := make(chan spellType, 2)
	c <- poison
	c <- magicMissile
	caster := func() spellType {
		return <-c
	}

	player := Player{
		hp:    10,
		mana:  250,
		casts: caster,
	}
	boss := Player{
		hp:  13,
		dmg: 8,
	}
	player, boss = battle(player, boss, false)
	TestEqual(t, 0, boss.hp)
	TestEqual(t, 2, player.hp)
}

func TestDay22Battle2(t *testing.T) {
	log.SetOutput(io.Discard)
	c := make(chan spellType, 5)
	c <- recharge
	c <- shield
	c <- drain
	c <- poison
	c <- magicMissile
	caster := func() spellType {
		return <-c
	}
	player := Player{
		hp:    10,
		mana:  250,
		casts: caster,
	}
	boss := Player{
		hp:  14,
		dmg: 8,
	}
	player, boss = battle(player, boss, false)
	TestEqual(t, -1, boss.hp)
	TestEqual(t, 1, player.hp)
}

func TestDay22a(t *testing.T) {
	log.SetOutput(io.Discard)
	TestEqual(t, 900, day22(false))
}

func TestDay22b(t *testing.T) {
	log.SetOutput(io.Discard)
	TestEqual(t, 1216, day22(true))
}
