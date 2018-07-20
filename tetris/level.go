package tetris

type Level interface {
	String() string
}

const (
	Space space = iota
	Block block = iota
)

type space int

func (s space) String() string {
	return " "
}

type block int

func (b block) String() string {
	return "●"
}
