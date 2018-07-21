package tetris

type Tetromino interface {
	doesExistInColumn(h int) bool
	doesExistInRow(w int) bool
	frame() frame
	obey(cmd command)
	asMoved(direction direction) Tetromino
	move(direction direction)
	moveFor(diff diff)
	rotate(rotation rotation)
}

type tetromino struct {
	f frame
}

type frame struct {
	figure figure
	p      point
	w      int
	h      int
}

type diff point

var (
	diffs = map[direction]diff{
		Left:  diff{x: -1, y: 0},
		Right: diff{x: 1, y: 0},
		Down:  diff{x: 0, y: 1},
	}
)

func (t tetromino) doesExistInColumn(h int) bool {
	if t.f.p.y < 0 || h <= t.f.p.y+t.f.h-1 {
		return false
	}

	return true
}

func (t tetromino) doesExistInRow(w int) bool {
	if t.f.p.x < 0 || w <= t.f.p.x+t.f.w-1 {
		return false
	}

	return true
}

func (t tetromino) frame() frame {
	return t.f
}

func (t *tetromino) obey(cmd command) {
}

func (t *tetromino) move(direction direction) {
	diff, ok := diffs[direction]
	if !ok {
		return
	}

	t.moveFor(diff)
}

func (t *tetromino) moveFor(diff diff) {
	t.f.p.x += diff.x
	t.f.p.y += diff.y
}

type i struct {
	tetromino
}

func newI() *i {
	i := new(i)
	figure := figure{
		[]Level{Block, Block, Block, Block},
	}
	i.tetromino = tetromino{
		f: frame{
			figure: figure,
			w:      4,
			h:      1,
		},
	}

	return i
}

func (i i) asMoved(direction direction) Tetromino {
	i.move(direction)
	return &i
}

func (i *i) rotate(rotation rotation) {
}

type o struct {
	tetromino
}

func newO() *o {
	o := new(o)
	figure := figure{
		[]Level{Block, Block},
		[]Level{Block, Block},
	}
	o.tetromino = tetromino{
		f: frame{
			figure: figure,
			w:      2,
			h:      2,
		},
	}

	return o
}

func (o o) asMoved(direction direction) Tetromino {
	o.move(direction)
	return &o
}

func (o *o) rotate(rotation rotation) {
}

type z struct {
	tetromino
}

func newZ() *z {
	z := new(z)
	figure := figure{
		[]Level{Block, Block, Space},
		[]Level{Space, Block, Block},
	}
	z.tetromino = tetromino{
		f: frame{
			figure: figure,
			w:      3,
			h:      2,
		},
	}

	return z
}

func (z z) asMoved(direction direction) Tetromino {
	z.move(direction)
	return &z
}

func (z *z) rotate(rotation rotation) {
}

type t struct {
	tetromino
}

func newT() *t {
	t := new(t)
	figure := figure{
		[]Level{Block, Block, Block},
		[]Level{Space, Block, Space},
	}
	t.tetromino = tetromino{
		f: frame{
			figure: figure,
			w:      3,
			h:      2,
		},
	}

	return t
}

func (t t) asMoved(direction direction) Tetromino {
	t.move(direction)
	return &t
}

func (t *t) rotate(rotation rotation) {
}

type l struct {
	tetromino
}

func newL() *l {
	l := new(l)
	figure := figure{
		[]Level{Block, Block, Block},
		[]Level{Block, Space, Space},
	}
	l.tetromino = tetromino{
		f: frame{
			figure: figure,
			w:      3,
			h:      2,
		},
	}

	return l
}

func (l l) asMoved(direction direction) Tetromino {
	l.move(direction)
	return &l
}

func (l *l) rotate(rotation rotation) {
}
