package enum

type Operation int

const (
	Add Operation = iota + 1
	Subtract
	Multiply
)
