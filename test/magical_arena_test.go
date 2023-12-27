package test

import (
	"magical-arena/pkg/dice"
	"magical-arena/pkg/magical_arena"
	"magical-arena/pkg/player"
	"math/rand"
	"testing"
	"time"
)

func TestPlayMatch(t *testing.T) {
	customRand := rand.New(rand.NewSource(time.Now().UnixNano()))
	playerA := player.NewPlayer(50, 5, 10)
	playerB := player.NewPlayer(100, 2, 5)
	attackingDice := dice.NewDice(customRand)
	defendingDice := dice.NewDice(customRand)

	arena := magical_arena.MagicalArena{
		PlayerA:       playerA,
		PlayerB:       playerB,
		AttackingDice: attackingDice,
		DefendingDice: defendingDice,
	}
	arena.PlayMatch()

	if playerA.GetHealth() != 0 && playerB.GetHealth() != 0 {
		t.Error("Expected either playerA or playerB health to be 0")
	}
}
