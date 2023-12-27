package test

import (
	"magical-arena/pkg/player"
	"testing"
)

func TestReduceHealth(t *testing.T) {
	player1 := player.NewLocalPlayer(50, 5, 10)
	player1.ReduceHealth(20)

	if player1.Health != 30 {
		t.Errorf("Expected health to be 30, got %d", player1.Health)
	}
}

func TestIsAlive(t *testing.T) {
	alivePlayer := player.NewLocalPlayer(30, 5, 10)
	deadPlayer := player.NewLocalPlayer(0, 5, 10)

	if !alivePlayer.IsAlive() {
		t.Error("Expected alivePlayer to be alive")
	}

	if deadPlayer.IsAlive() {
		t.Error("Expected deadPlayer to be dead")
	}
}
