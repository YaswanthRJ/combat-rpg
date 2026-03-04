package domain

import "fmt"

type Creature struct {
	HP      int
	MaxHP   int
	Attack  int
	Defense int
	Actions []Action
}

type CreatureTemplate struct {
	Name        string
	MaxHP       int
	Attack      int
	Defense     int
	Actions     []Action
	Description string
}

var CreaturePool = map[string]CreatureTemplate{
	"Bandit":  {MaxHP: 20, Attack: 15, Defense: 10, Actions: []Action{FastAttack, HeavyAttack}, Description: "A common bandit"},
	"Soldier": {MaxHP: 30, Attack: 20, Defense: 25, Actions: []Action{FastAttack, HeavyAttack, Block}, Description: "A trained and well equipped soldier"},
}

func GenerateCreature(Name string) (Creature, error) {
	t, ok := CreaturePool[Name]

	if !ok {
		return Creature{}, fmt.Errorf("creature not found: %s", Name)
	}
	return Creature{
		HP:      t.MaxHP,
		MaxHP:   t.MaxHP,
		Attack:  t.Attack,
		Defense: t.Defense,
		Actions: append([]Action{}, t.Actions...),
	}, nil

}
