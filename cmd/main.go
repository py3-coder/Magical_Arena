package main

import (
	"fmt"
	"magical-arena/pkg/player"
)

func main() {
	fmt.Print("Main Application Entry Point")
	playerA := player.NewPlayer(23, 34, 56)
	fmt.Print(playerA.Health)
}
