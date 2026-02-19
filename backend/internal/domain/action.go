package domain

type Action int

const (
	FastAttack Action = iota
	HeavyAttack
	Block
)
