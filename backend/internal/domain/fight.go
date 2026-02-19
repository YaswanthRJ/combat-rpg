package domain

type FightStatus int

const (
	Ongoing FightStatus = iota
	PlayerWon
	PlayerLost
)

type FightState struct {
	Player       Creature
	Enemy        Creature
	FightStatus  FightStatus
	ActionNumber int
}
