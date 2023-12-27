package player

import (
	"magical-arena/models"
)

type LocalPlayer struct {
	*models.Player
}

// Creating New Player with given attributes :
func NewPlayer(health, strength, attack int) *models.Player {
	return &models.Player{health, strength, attack}
}

// NewLocalPlayer creates a new LocalPlayer by embedding a Player with the given attributes.
func NewLocalPlayer(health, strength, attack int) *LocalPlayer {
	return &LocalPlayer{NewPlayer(health, strength, attack)}
}

// IsAlive Check Player is still alive :
func (lp *LocalPlayer) IsAlive() bool {
	return lp.Health > 0
}

// ReduceHealth reduces the player's health by the specified amount.
func (lp *LocalPlayer) ReduceHealth(damage int) {
	lp.Health -= damage
	if lp.Health < 0 {
		lp.Health = 0
	}
}
