package app

import "combat-sim/internal/domain"

type Campaign struct {
	ID            string
	Player        domain.Creature
	Fights        map[string]*domain.FightState
	ActiveFightID string
}
