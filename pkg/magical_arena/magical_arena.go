package magical_arena

import (
	"fmt"
	"magical-arena/pkg/dice"
	"magical-arena/pkg/player"
)

type MagicalArena struct {
	PlayerA       *player.Player
	PlayerB       *player.Player
	AttackingDice *dice.Dice
	DefendingDice *dice.Dice
}

// PlayMatch simulates a match in the Magical Arena.
func (arena *MagicalArena) PlayMatch() {
	for arena.PlayerA.IsAlive() && arena.PlayerB.IsAlive() {

		var attacker, defender *player.Player
		if arena.PlayerA.Health < arena.PlayerB.Health {
			attacker, defender = arena.PlayerA, arena.PlayerB
		} else {
			attacker, defender = arena.PlayerB, arena.PlayerA
		}

		attackingRoll := arena.AttackingDice.Roll()
		attackDamage := attackingRoll * attacker.Attack

		defendingRoll := arena.DefendingDice.Roll()
		defendingStrength := defendingRoll * defender.Strength

		damageTaken := attackDamage - defendingStrength
		if damageTaken > 0 {
			defender.ReduceHealth(damageTaken)
		}

		fmt.Printf("%s attacks for %d (Rolled: %d), %s defends for %d (Rolled: %d), %s health: %d\n",
			getPlayerType(attacker, arena), attackDamage, attackingRoll,
			getPlayerType(defender, arena), defendingStrength, defendingRoll,
			getPlayerType(defender, arena), defender.Health)
	}

	winner := arena.PlayerA
	if arena.PlayerB.IsAlive() {
		winner = arena.PlayerB
	}

	fmt.Printf("%s wins!\n", getPlayerType(winner, arena))
}

func getPlayerType(currplayer *player.Player, arena *MagicalArena) string {
	if currplayer == nil {
		return ""
	}
	if currplayer == arena.PlayerA {
		return "Player A"
	}
	return "Player B"
}
