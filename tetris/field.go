package tetris

type Field interface {
	width() int
	height() int
	haveConfliction(frame frame) bool
	put(frame frame)
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

func (f field) haveConfliction(frame frame) bool {
	for i := 0; i < frame.h; i++ {
		for j := 0; j < frame.w; j++ {
			if frame.figure[i][j] != Block {
				continue
			}

			y := frame.p.y + i
			x := frame.p.x + j
			if f.figure[y][x] == Block {
				return true
			}
		}
	}

	return false
}

func (f *field) put(frame frame) {
	for i := 0; i < frame.h; i++ {
		for j := 0; j < frame.w; j++ {
			if frame.figure[i][j] != Block {
				continue
			}

			y := frame.p.y + i
			x := frame.p.x + j
			f.figure[y][x] = Block
		}
	}
}

func (f *field) clear(figure figure) {
}
