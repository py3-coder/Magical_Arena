package dice

import (
	"magical-arena/models"
	"math/rand"
	"time"
)

type LocalDiceModel struct {
	*models.Dice
}

// NewDice creates a new six-sided die.
func NewDice() *models.Dice {
	return &models.Dice{Random: rand.New(rand.NewSource(time.Now().UnixNano()))}
}

// NewLocalDice creates a new LocalDice by embedding.
func NewLocalDice() *LocalDiceModel {
	return &LocalDiceModel{NewDice()}
}

// Roll simulates a die roll and returns the result.
func (d *LocalDiceModel) Roll() int {
	return d.Dice.Random.Intn(6) + 1
}
