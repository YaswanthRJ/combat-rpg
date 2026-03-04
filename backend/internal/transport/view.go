package transport

import "combat-sim/internal/domain"

type FightView struct {
	Player PlayerView `json:"player"`
	Enemy  EnemyView  `json:"enemy"`
	Status string     `json:"status"`
}

type PlayerView struct {
	HP      int          `json:"hp"`
	MaxHP   int          `json:"maxHP"`
	Actions []ActionView `json:"actions"`
}

type EnemyView struct {
	Name        string `json:"name"`
	HP          int    `json:"hp"`
	MaxHP       int    `json:"maxHP"`
	Description string `json:"description"`
}

type ActionView struct {
	ID   domain.Action `json:"id"`
	Name string        `json:"name"`
}

func ToFightView(state *domain.FightState, template domain.CreatureTemplate) FightView {

	actions := make([]ActionView, 0, len(state.Player.Actions))

	for _, a := range state.Player.Actions {
		data, ok := domain.ActionPool[a]
		if !ok {
			continue
		}
		actions = append(actions, ActionView{
			ID:   a,
			Name: data.Name,
		})
	}

	return FightView{
		Player: PlayerView{
			HP:      state.Player.HP,
			MaxHP:   state.Player.MaxHP,
			Actions: actions,
		},
		Enemy: EnemyView{
			Name:        template.Name,
			HP:          state.Enemy.HP,
			MaxHP:       state.Enemy.MaxHP,
			Description: template.Description,
		},
		Status: mapStatus(state.FightStatus),
	}
}

func mapStatus(s domain.FightStatus) string {
	switch s {
	case domain.Ongoing:
		return "ongoing"
	case domain.PlayerWon:
		return "won"
	case domain.PlayerLost:
		return "lost"
	default:
		return "unknown"
	}
}
