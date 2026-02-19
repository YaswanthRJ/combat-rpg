package domain

type Actionkind int
type Action int

const (
	HeavyAttack Action = iota
	FastAttack
	Block
)

const (
	ActionAttack Actionkind = iota
	ActionDefend
	ActionUtility
)

type ActionData struct {
	Name       string
	Kind       Actionkind
	Multiplier float64
}

var ActionPool = map[Action]ActionData{
	HeavyAttack: {Name: "Heavy Attack", Kind: ActionAttack, Multiplier: 1.3},
	FastAttack:  {Name: "Fast Attack", Kind: ActionAttack, Multiplier: 0.8},
	Block:       {Name: "Block", Kind: ActionDefend, Multiplier: 0.25},
}
