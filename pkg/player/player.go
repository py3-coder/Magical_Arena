package player

// Player Struct represents a player in the Magical Arena.
type Player struct {
	Health   int
	Strength int
	Attack   int
}

// Creating New Player with given attributes :
func NewPlayer(health, strength, attack int) *Player {
	return &Player{Health: health, Strength: strength, Attack: attack}
}

// IsAlive Check Player is still alive :
func (player *Player) IsAlive() bool {
	return player.Health > 0
}

// GetHealth to Check Player Health :
func (player *Player) GetHealth() int {
	return player.Health
}

// ReduceHealth reduces the player's health by the specified amount.
func (player *Player) ReduceHealth(damage int) {
	player.Health -= damage
	if player.Health < 0 {
		player.Health = 0
	}
}
