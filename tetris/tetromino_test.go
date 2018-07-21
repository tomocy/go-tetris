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

func TestGetAsMoved(t *testing.T) {
	given := getTetromino()
	thens := getGetAsMovedTestCases()
	for _, then := range thens {
		have := given.getAsMoved(then.in)
		if !reflect.DeepEqual(have.p, then.want) {
			t.Errorf("have %v, but want %v", have.p, then.want)
			t.Errorf("when in is %v", then.in)
		}
	}
}

func getGetAsMovedTestCases() []struct {
	in   command
	want point
} {
	return []struct {
		in   command
		want point
	}{
		{
			in: command(Left),
			want: point{
				x: 0,
				y: 0,
			},
		},
		{
			in: command(Right),
			want: point{
				x: 2,
				y: 0,
			},
		},
		{
			in: command(Down),
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
		have := given
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

func getTetromino() tetromino {
	tetromino := tetromino{
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

	return tetromino
}

func TestNewI(t *testing.T) {
	want := &i{
		tetromino: tetromino{
			figure: figure{
				[]Level{Block, Block, Block, Block},
			},
			w: 4,
			h: 1,
		},
	}
	have := newI()
	if !reflect.DeepEqual(have, want) {
		t.Errorf("have %#v, but want %#v", have, want)
		t.Error("set it properly in newI method\n")
	}
}

func TestNewO(t *testing.T) {
	want := &o{
		tetromino: tetromino{
			figure: figure{
				[]Level{Block, Block},
				[]Level{Block, Block},
			},
			w: 2,
			h: 2,
		},
	}
	have := newO()
	if !reflect.DeepEqual(have, want) {
		t.Errorf("have %#v, but want %#v", have, want)
		t.Error("set it properly in newO method\n")
	}
}

func TestNewZ(t *testing.T) {
	want := &z{
		tetromino: tetromino{
			figure: figure{
				[]Level{Block, Block, Space},
				[]Level{Space, Block, Block},
			},
			w: 3,
			h: 2,
		},
	}
	have := newZ()
	if !reflect.DeepEqual(have, want) {
		t.Errorf("have %#v, but want %#v", have, want)
		t.Error("set it properly in newZ method\n")
	}
}

func TestNewT(tt *testing.T) {
	want := &t{
		tetromino: tetromino{
			figure: figure{
				[]Level{Block, Block, Block},
				[]Level{Space, Block, Space},
			},
			w: 3,
			h: 2,
		},
	}
	have := newT()
	if !reflect.DeepEqual(have, want) {
		tt.Errorf("have %#v, but want %#v", have, want)
		tt.Error("set it properly in newT method\n")
	}
}

func TestNewL(t *testing.T) {
	want := &l{
		tetromino: tetromino{
			figure: figure{
				[]Level{Block, Block, Block},
				[]Level{Block, Space, Space},
			},
			w: 3,
			h: 2,
		},
	}
	have := newL()
	if !reflect.DeepEqual(have, want) {
		t.Errorf("have %#v, but want %#v", have, want)
		t.Error("set it properly in newL method\n")
	}
}
