package tetris

type command int

type direction command

const (
	Left direction = iota
	Right
	Down
)

type rotation command

const (
	Head rotation = iota
	Rightward
	Tail
	Leftward
)
