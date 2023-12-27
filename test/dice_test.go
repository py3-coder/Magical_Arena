package test

import (
	"magical-arena/pkg/dice"
	"math/rand"
	"testing"
	"time"
)

func TestRoll(t *testing.T) {
	rand.Seed(time.Now().UnixNano())
	customRand := rand.New(rand.NewSource(time.Now().UnixNano()))
	diceCheck := dice.NewDice(customRand)
	rollResult := diceCheck.Roll()

	if rollResult < 1 || rollResult > 6 {
		t.Errorf("Expected roll result between 1 and 6, got %d", rollResult)
	}
}
