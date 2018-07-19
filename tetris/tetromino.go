package tetris

type Tetromino struct {
	frame  Frame
	figure Figure
}

func NewI() *Tetromino {
	t := new(Tetromino)
	t.setUpI()
	return t
}

func (t *Tetromino) setUpI() {
	t.figure = Figure{
		[]Level{block, block, block, block},
	}
	t.frame = Frame{
		w: 4,
		h: 1,
	}
}
