package domain

import "testing"

func TestHeavyAttackDealsDamage(t *testing.T) {
	player := Creature{HP: 100, MaxHP: 100, Attack: 20, Defense: 5}
	enemy := Creature{HP: 100, MaxHP: 100, Attack: 0, Defense: 5}

	state := &FightState{
		Player:      player,
		Enemy:       enemy,
		FightStatus: Ongoing,
	}
	result, err := ResolveRound(state, HeavyAttack)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	expected := int(float64(20)*(ActionPool[HeavyAttack].Multiplier) - 5)
	if result.PlayerDamageDealt != expected {
		t.Errorf("expected %d damage, got %d", expected, result.PlayerDamageDealt)
	}
	if state.Enemy.HP != 100-expected {
		t.Errorf("expected enemy HP %d, got %d", 100-expected, state.Enemy.HP)
	}
}

func TestPlayerWin(t *testing.T) {
	player := Creature{HP: 100, MaxHP: 100, Attack: 50, Defense: 5}
	enemy := Creature{HP: 10, MaxHP: 100, Attack: 0, Defense: 5}

	state := &FightState{
		Player:      player,
		Enemy:       enemy,
		FightStatus: Ongoing,
	}
	result, err := ResolveRound(state, HeavyAttack)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	if state.FightStatus != PlayerWon {
		t.Errorf("expected PlayerWon, got %v", state.FightStatus)
	}

	if !result.FightEnded {
		t.Errorf("expected fight to end")
	}
}

func TestPlayerLosesFight(t *testing.T) {
	player := Creature{HP: 10, MaxHP: 10, Attack: 5, Defense: 0}
	enemy := Creature{HP: 100, MaxHP: 100, Attack: 50, Defense: 0}

	state := &FightState{
		Player:      player,
		Enemy:       enemy,
		FightStatus: Ongoing,
	}

	result, err := ResolveRound(state, Block) // block won't save here
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	if state.FightStatus != PlayerLost {
		t.Errorf("expected PlayerLost, got %v", state.FightStatus)
	}

	if !result.FightEnded {
		t.Errorf("expected fight to end")
	}
}

func TestInvalidAction(t *testing.T) {
	player := Creature{HP: 100, MaxHP: 100, Attack: 20, Defense: 5}
	enemy := Creature{HP: 100, MaxHP: 100, Attack: 20, Defense: 5}

	state := &FightState{
		Player:      player,
		Enemy:       enemy,
		FightStatus: Ongoing,
	}

	invalid := Action(999)

	_, err := ResolveRound(state, invalid)
	if err == nil {
		t.Errorf("expected error for invalid action")
	}
}
