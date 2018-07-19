package tetris

type Command int

const (
	Left Command = iota
	Right
	Down
	Rotate
)
