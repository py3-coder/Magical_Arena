package main

import (
	"fmt"
	"magical-arena/pkg/dice"
	"magical-arena/pkg/magical_arena"
	"magical-arena/pkg/player"
	"math/rand"
	"time"
)

func main() {
	fmt.Println("Let's Start the Match ::[Player A , Player B] are Ready:")
	fmt.Println("Let's go..")
	customRand := rand.New(rand.NewSource(time.Now().UnixNano()))
	playerA := player.NewPlayer(50, 5, 10)
	playerB := player.NewPlayer(40, 10, 5)

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
