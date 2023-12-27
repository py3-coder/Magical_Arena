package dice

import (
	"math/rand"
)

// Dice Struct Represents a six-sided die
type Dice struct {
	Random *rand.Rand
}

// NewDice creates a new six-sided die.
func NewDice(r *rand.Rand) *Dice {
	return &Dice{Random: r}
}

// Roll simulates a die roll and returns the result.
func (d *Dice) Roll() int {
	return d.Random.Intn(6) + 1
}
