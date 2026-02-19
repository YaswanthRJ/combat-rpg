package domain

import "math"

// Current balance:
// Player Fast → 0.8
// Player Heavy → 1.3
// Player Block → 75% reduction
// Enemy → 1.0 baseline
func ResolveAction(state *FightState, action Action) ActionResult {

	result := ActionResult{}

	//if fight isnt ongoing, return
	if state.FightStatus != Ongoing {
		return result
	}

	state.ActionNumber++

	playerMultiplier := 0.0
	playerBlocking := false

	switch action {
	case FastAttack:
		playerMultiplier = 0.8
	case HeavyAttack:
		playerMultiplier = 1.3
	case Block:
		playerBlocking = true
	}

	playerDamage := 0

	if !playerBlocking {
		//Raw power in the attack
		rawAttack := (float64(state.Player.Attack) * playerMultiplier)
		//Damage the attak dealt after encounteruing enemy defense
		playerDamage = int(math.Floor(rawAttack - float64(state.Enemy.Defense)))
		if playerDamage < 0 {
			playerDamage = 0
		}
		//Apply the damage
		state.Enemy.HP = state.Enemy.HP - playerDamage
		if state.Enemy.HP <= 0 {
			state.FightStatus = PlayerWon
		}
	}
	result.PlayerDamageDealt = playerDamage
	if state.FightStatus == PlayerWon {
		result.FightEnded = true
		return result
	}

	enemyMultiplier := 1.0
	enemyDamage := 0

	rawAttack := float64(state.Enemy.Attack) * enemyMultiplier
	enemyDamage = int(math.Floor(rawAttack - float64(state.Player.Defense)))
	if enemyDamage < 0 {
		enemyDamage = 0
	}
	if playerBlocking {
		enemyDamage = enemyDamage / 4
	}

	state.Player.HP = state.Player.HP - enemyDamage
	if state.Player.HP <= 0 {
		state.FightStatus = PlayerLost
	}
	result.EnemyDamageDealt = enemyDamage
	if state.FightStatus == PlayerLost {
		result.FightEnded = true
		return result
	}

	return result
}

//TODO:
// Extract calculateDamage(attack, defense, multiplier)
// Extract resolvePlayerPhase
// Extract resolveEnemyPhase
// Keep ResolveAction small and readable
// Then write tests
