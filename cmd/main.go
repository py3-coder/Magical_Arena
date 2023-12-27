package main

import (
	"magical-arena/pkg/dice"
	"magical-arena/pkg/magical_arena"
	"magical-arena/pkg/player"
	"math/rand"
	"time"
)

func main() {
	customRand := rand.New(rand.NewSource(time.Now().UnixNano()))
	playerA := player.NewPlayer(150, 5, 10)
	playerB := player.NewPlayer(100, 10, 5)

	attackingDice := dice.NewDice(customRand)
	defendingDice := dice.NewDice(customRand)

	arena := magical_arena.MagicalArena{
		PlayerA:       playerA,
		PlayerB:       playerB,
		AttackingDice: attackingDice,
		DefendingDice: defendingDice,
	}
	arena.PlayMatch()
}
