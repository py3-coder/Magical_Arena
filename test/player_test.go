package test

import (
	"magical-arena/pkg/player"
	"testing"
)

// Test Reduece Health :
func TestReduceHealth(t *testing.T) {
	player1 := player.NewPlayer(50, 5, 10)
	player1.ReduceHealth(20)

	if player1.Health != 30 {
		t.Errorf("Expected health to be 30, got %d", player1.Health)
	}
}

// Test IsAlive  :
func TestIsAlive(t *testing.T) {
	alivePlayer := player.NewPlayer(30, 5, 10)
	deadPlayer := player.NewPlayer(0, 5, 10)

	if !alivePlayer.IsAlive() {
		t.Error("Expected alivePlayer to be alive")
	}

	if deadPlayer.IsAlive() {
		t.Error("Expected deadPlayer to be dead")
	}
}
