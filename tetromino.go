package main

// Figure is ...
type Figure [][]Level

// Point is ...
type Point struct {
	x int
	y int
}

// Frame is ...
type Frame struct {
	p Point
	w int
	h int
}

type TetrominoType int

const (
	I TetrominoType = iota
	O
	Z
	T
	L
)

// RotationType is ...
type RotationType int

// Head is ...
// Lefty is ...
// Righty is ...
// Tail is ...
const (
	Head RotationType = iota
	Righty
	Tail
	Lefty
)

const (
	RotationTypeKinds = 4
)

// Tetromino is ...
type Tetromino struct {
	figure Figure
	frame  Frame
	t      TetrominoType
	rt     RotationType
}

func MakeI(rt RotationType) Tetromino {
	var figure Figure
	var frame Frame
	switch rt {
	case Head, Tail:
		figure = Figure{
			[]Level{Block, Block, Block, Block},
		}
		frame = Frame{
			w: 4,
			h: 1,
		}
	case Lefty, Righty:
		figure = Figure{
			[]Level{Block},
			[]Level{Block},
			[]Level{Block},
			[]Level{Block},
		}
		frame = Frame{
			w: 1,
			h: 4,
		}
	}

	return Tetromino{figure: figure, frame: frame, t: I, rt: rt}
}

func MakeO(rt RotationType) Tetromino {
	var figure Figure
	var frame Frame
	switch rt {
	case Head, Lefty, Righty, Tail:
		figure = Figure{
			[]Level{Block, Block},
			[]Level{Block, Block},
		}
		frame = Frame{
			w: 2,
			h: 2,
		}
	}

	return Tetromino{figure: figure, frame: frame, t: O, rt: rt}
}

func MakeZ(rt RotationType) Tetromino {
	var figure Figure
	var frame Frame
	switch rt {
	case Head:
		figure = Figure{
			[]Level{Block, Block, Space},
			[]Level{Space, Block, Block},
		}
		frame = Frame{
			w: 3,
			h: 2,
		}
	case Lefty:
		figure = Figure{
			[]Level{Space, Block},
			[]Level{Block, Block},
			[]Level{Block, Space},
		}
		frame = Frame{
			w: 2,
			h: 3,
		}
	case Righty:
		figure = Figure{
			[]Level{Block, Space},
			[]Level{Block, Block},
			[]Level{Space, Block},
		}
		frame = Frame{
			w: 2,
			h: 3,
		}
	case Tail:
		figure = Figure{
			[]Level{Space, Block, Block},
			[]Level{Block, Block, Space},
		}
		frame = Frame{
			w: 3,
			h: 2,
		}
	}

	return Tetromino{figure: figure, frame: frame, t: Z, rt: rt}
}

func MakeT(rt RotationType) Tetromino {
	var figure Figure
	var frame Frame
	switch rt {
	case Head:
		figure = Figure{
			[]Level{Block, Block, Block},
			[]Level{Space, Block, Space},
		}
		frame = Frame{
			w: 3,
			h: 2,
		}
	case Lefty:
		figure = Figure{
			[]Level{Block, Space},
			[]Level{Block, Block},
			[]Level{Block, Space},
		}
		frame = Frame{
			w: 2,
			h: 3,
		}
	case Righty:
		figure = Figure{
			[]Level{Space, Block},
			[]Level{Block, Block},
			[]Level{Space, Block},
		}
		frame = Frame{
			w: 2,
			h: 3,
		}
	case Tail:
		figure = Figure{
			[]Level{Space, Block, Space},
			[]Level{Block, Block, Block},
		}
		frame = Frame{
			w: 3,
			h: 2,
		}
	}

	return Tetromino{figure: figure, frame: frame, t: T, rt: rt}
}

func MakeL(rt RotationType) Tetromino {
	var figure Figure
	var frame Frame
	switch rt {
	case Head:
		figure = Figure{
			[]Level{Block, Block, Block},
			[]Level{Block, Space, Space},
		}
		frame = Frame{
			w: 3,
			h: 2,
		}
	case Lefty:
		figure = Figure{
			[]Level{Block, Space},
			[]Level{Block, Space},
			[]Level{Block, Block},
		}
		frame = Frame{
			w: 2,
			h: 3,
		}
	case Righty:
		figure = Figure{
			[]Level{Space, Block},
			[]Level{Space, Block},
			[]Level{Block, Block},
		}
		frame = Frame{
			w: 2,
			h: 3,
		}
	case Tail:
		figure = Figure{
			[]Level{Space, Space, Block},
			[]Level{Block, Block, Block},
		}
		frame = Frame{
			w: 3,
			h: 2,
		}
	}

	return Tetromino{figure: figure, frame: frame, t: L, rt: rt}
}

// Diff is ...
type Diff Point

func (tetr *Tetromino) moveFor(diff Diff) {
	tetr.frame.p = Point{
		x: tetr.frame.p.x + diff.x,
		y: tetr.frame.p.y + diff.y,
	}
}

func (tetr *Tetromino) Rotate() {
	var newTetr Tetromino
	rt := RotationType((tetr.rt + 1) % RotationTypeKinds)
	switch tetr.t {
	case I:
		newTetr = MakeI(rt)
	case O:
		newTetr = MakeO(rt)
	case Z:
		newTetr = MakeZ(rt)
	case T:
		newTetr = MakeT(rt)
	case L:
		newTetr = MakeL(rt)
	}

	tetr.figure = newTetr.figure
	tetr.frame.w = newTetr.frame.w
	tetr.frame.h = newTetr.frame.h
	tetr.t = newTetr.t
	tetr.rt = newTetr.rt
}
