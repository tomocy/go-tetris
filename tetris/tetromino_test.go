package tetris

import (
	"reflect"
	"testing"
)

func TestDoesExistInColumn(t *testing.T) {
	given := getTetromino()
	thens := getDoesExistInColumnTestCases()
	for _, then := range thens {
		have := given.doesExistInColumn(then.in)
		if have != then.want {
			t.Errorf("have %v, but want %v", have, then.want)
			t.Errorf("when in is %v", then.in)
		}
	}
}

func getDoesExistInColumnTestCases() []struct {
	in   int
	want bool
} {
	return []struct {
		in   int
		want bool
	}{
		{-1, false},
		{0, false},
		{1, true},
		{2, true},
	}
}
func TestDoesExistInRow(t *testing.T) {
	given := getTetromino()
	thens := getDoesExistInRowTestCases()
	for _, then := range thens {
		have := given.doesExistInRow(then.in)
		if have != then.want {
			t.Errorf("have %v, but want %v", have, then.want)
			t.Errorf("when in is %v", then.in)
		}
	}
}

func getDoesExistInRowTestCases() []struct {
	in   int
	want bool
} {
	return []struct {
		in   int
		want bool
	}{
		{3, false},
		{4, false},
		{5, true},
		{6, true},
	}
}

func TestFrame(t *testing.T) {
	given := getTetromino()
	want := getGetFrameExpectation()
	have := given.frame()
	if !reflect.DeepEqual(have, want) {
		t.Errorf("have %v, but want %v", have, want)
	}
}

func getGetFrameExpectation() frame {
	return frame{
		figure: figure{
			[]Level{Block, Block, Block, Block},
		},
		p: point{
			x: 1,
			y: 0,
		},
		w: 4,
		h: 1,
	}
}

func TestGetAsMoved(t *testing.T) {
	givens := getTetrominos()
	thens := getAsMovedTestCases()
	for _, given := range givens {
		for _, then := range thens {
			have := given.asMoved(then.in).frame()
			if !reflect.DeepEqual(have.p, then.want) {
				t.Errorf("have %v, but want %v", have.p, then.want)
				t.Errorf("when given is %v", given)
				t.Errorf("when in is %v", then.in)
			}
		}
	}
}

func getAsMovedTestCases() []struct {
	in   direction
	want point
} {
	return []struct {
		in   direction
		want point
	}{
		{
			in: Left,
			want: point{
				x: 0,
				y: 0,
			},
		},
		{
			in: Right,
			want: point{
				x: 2,
				y: 0,
			},
		},
		{
			in: Down,
			want: point{
				x: 1,
				y: 1,
			},
		},
	}
}

func TestMove(t *testing.T) {
	given := getTetromino()
	thens := getMoveTestCases()
	for _, then := range thens {
		given.move(then.in)
		have := given.frame()
		if !reflect.DeepEqual(have.p, then.want) {
			t.Errorf("have %v, but want %v", have.p, then.want)
			t.Errorf("when in is %v", then.in)
		}

		given = getTetromino()
	}
}

func getMoveTestCases() []struct {
	in   direction
	want point
} {
	return []struct {
		in   direction
		want point
	}{
		{
			in: Left,
			want: point{
				x: 0,
				y: 0,
			},
		},
		{
			in: Right,
			want: point{
				x: 2,
				y: 0,
			},
		},
		{
			in: Down,
			want: point{
				x: 1,
				y: 1,
			},
		},
	}
}

func getTetromino() Tetromino {
	i := getI()
	i.tetromino.f.p = point{
		x: 1,
		y: 0,
	}
	return i
}

func getTetrominos() []Tetromino {
	p := point{
		x: 1,
		y: 0,
	}
	i := getI()
	i.tetromino.f.p = p
	o := getO()
	o.tetromino.f.p = p
	z := getZ()
	z.tetromino.f.p = p
	t := getT()
	t.tetromino.f.p = p
	l := getL()
	l.tetromino.f.p = p
	return []Tetromino{
		i, o, z, t, l,
	}
}

func TestNewI(t *testing.T) {
	want := getI()
	have := newI()
	if !reflect.DeepEqual(have, want) {
		t.Errorf("have %#v, but want %#v", have, want)
		t.Error("set it properly in newI method\n")
	}
}

func getI() *i {
	return &i{
		tetromino: tetromino{
			f: frame{
				figure: figure{
					[]Level{Block, Block, Block, Block},
				},
				w: 4,
				h: 1,
			},
		},
	}
}

func TestNewO(t *testing.T) {
	want := getO()
	have := newO()
	if !reflect.DeepEqual(have, want) {
		t.Errorf("have %#v, but want %#v", have, want)
		t.Error("set it properly in newO method\n")
	}
}

func getO() *o {
	return &o{
		tetromino: tetromino{
			f: frame{
				figure: figure{
					[]Level{Block, Block},
					[]Level{Block, Block},
				},
				w: 2,
				h: 2,
			},
		},
	}
}

func TestNewZ(t *testing.T) {
	want := getZ()
	have := newZ()
	if !reflect.DeepEqual(have, want) {
		t.Errorf("have %#v, but want %#v", have, want)
		t.Error("set it properly in newZ method\n")
	}
}

func getZ() *z {
	return &z{
		tetromino: tetromino{
			f: frame{
				figure: figure{
					[]Level{Block, Block, Space},
					[]Level{Space, Block, Block},
				},
				w: 3,
				h: 2,
			},
		},
	}
}

func TestNewT(t *testing.T) {
	want := getT()
	have := newT()
	if !reflect.DeepEqual(have, want) {
		t.Errorf("have %#v, but want %#v", have, want)
		t.Error("set it properly in newT method\n")
	}
}

func getT() *t {
	return &t{
		tetromino: tetromino{
			f: frame{
				figure: figure{
					[]Level{Block, Block, Block},
					[]Level{Space, Block, Space},
				},
				w: 3,
				h: 2,
			},
		},
	}
}

func TestNewL(t *testing.T) {
	want := getL()
	have := newL()
	if !reflect.DeepEqual(have, want) {
		t.Errorf("have %#v, but want %#v", have, want)
		t.Error("set it properly in newL method\n")
	}
}

func getL() *l {
	return &l{
		tetromino: tetromino{
			f: frame{
				figure: figure{
					[]Level{Block, Block, Block},
					[]Level{Block, Space, Space},
				},
				w: 3,
				h: 2,
			},
		},
	}
}
