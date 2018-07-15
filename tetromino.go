package main

// Tetromino is a figure in tetris
type Tetromino struct {
	frame tetromino
	w     int
	h     int
	p     Point
}

type tetromino [][]Level

func MakeI() Tetromino {
	frame := tetromino{
		[]Level{Block, Block, Block, Block},
	}
	return Tetromino{
		frame: frame,
		w:     4,
		h:     1,
	}
}

func MakeO() Tetromino {
	frame := tetromino{
		[]Level{Block, Block},
		[]Level{Block, Block},
	}
	return Tetromino{
		frame: frame,
		w:     2,
		h:     2,
	}
}

func MakeZ() Tetromino {
	frame := tetromino{
		[]Level{Block, Block, Space},
		[]Level{Space, Block, Block},
	}
	return Tetromino{
		frame: frame,
		w:     3,
		h:     2,
	}
}

func MakeT() Tetromino {
	frame := tetromino{
		[]Level{Block, Block, Block},
		[]Level{Space, Block, Space},
	}
	return Tetromino{
		frame: frame,
		w:     3,
		h:     2,
	}
}

func MakeL() Tetromino {
	frame := tetromino{
		[]Level{Block, Block, Block},
		[]Level{Block, Space, Space},
	}
	return Tetromino{
		frame: frame,
		w:     3,
		h:     2,
	}
}
