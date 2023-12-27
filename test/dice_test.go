package test

import (
	"magical-arena/pkg/dice"
	"testing"
)

func TestRoll(t *testing.T) {
	diceCheck := dice.NewLocalDice()
	rollResult := diceCheck.Roll()

	if rollResult < 1 || rollResult > 6 {
		t.Errorf("Expected roll result between 1 and 6, got %d", rollResult)
	}
}
