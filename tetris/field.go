package tetris

type Field interface {
	width() int
	height() int
	haveConfliction(tetromino tetromino) bool
	put(tetromino tetromino)
	clear(figure figure)
}

type field struct {
	figure figure
	w      int
	h      int
}

func newField(w, h int) *field {
	f := new(field)
	f.w = w
	f.h = h
	f.figure = make(figure, h)
	for i := 0; i < h; i++ {
		f.figure[i] = make([]Level, w)
		for j := 0; j < w; j++ {
			f.figure[i][j] = Space
		}
	}

	return f
}

func (f field) width() int {
	return f.w
}

func (f field) height() int {
	return f.h
}

func (f field) haveConfliction(tetromino tetromino) bool {
	for i := 0; i < tetromino.h; i++ {
		for j := 0; j < tetromino.w; j++ {
			if tetromino.figure[i][j] != Block {
				continue
			}

			y := tetromino.p.y + i
			x := tetromino.p.x + j
			if f.figure[y][x] == Block {
				return true
			}
		}
	}

	return false
}

func (f *field) put(tetromino tetromino) {
	for i := 0; i < tetromino.h; i++ {
		for j := 0; j < tetromino.w; j++ {
			if tetromino.figure[i][j] != Block {
				continue
			}

			y := tetromino.p.y + i
			x := tetromino.p.x + j
			f.figure[y][x] = Block
		}
	}
}

func (f *field) clear(figure figure) {
}
