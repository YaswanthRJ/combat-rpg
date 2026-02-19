package domain

// Current balance:
// Player Fast → 0.8
// Player Heavy → 1.3
// Player Block → 75% reduction
// Enemy → 1.0 baseline

func ResolveRound(state *FightState, action Action) ActionResult {

	result := ActionResult{}

	// --- Player Phase ---
	playerDamage, defended := ResolvePlayerAction(state, action)
	result.PlayerDamageDealt = playerDamage

	if playerDamage > 0 {
		if applyDamage(&state.Enemy, playerDamage) {
			state.FightStatus = PlayerWon
			result.FightEnded = true
			return result
		}
	}

	// --- Enemy Phase ---
	enemyDamage := ResolveEnemyAction(state, defended)
	result.EnemyDamageDealt = enemyDamage

	if enemyDamage > 0 {
		if applyDamage(&state.Player, enemyDamage) {
			state.FightStatus = PlayerLost
			result.FightEnded = true
		}
	}

	return result
}

func ResolvePlayerAction(state *FightState, action Action) (damage int, defended bool) {

	data, ok := ActionPool[action]
	if !ok {
		panic("unknown action")
	}

	if data.Kind == ActionDefend {
		return 0, true
	}

	damage = calculateDamage(state.Player.Attack, state.Enemy.Defense, data.Multiplier)

	return damage, false
}

func ResolveEnemyAction(state *FightState, playerBlocked bool) int {

	damage := calculateDamage(state.Enemy.Attack, state.Player.Defense, 1.0)

	if playerBlocked {
		damage = applyMultiplier(damage, 0.25)
	}

	return damage
}

func applyDamage(creature *Creature, damage int) bool {
	creature.HP -= damage
	return creature.HP <= 0
}

func calculateDamage(attack, defense int, multiplier float64) int {
	damage := int(float64(attack)*multiplier - float64(defense))
	if damage < 0 {
		return 0
	}
	return damage
}

func applyMultiplier(value int, multiplier float64) int {
	return int(float64(value) * multiplier)
}
