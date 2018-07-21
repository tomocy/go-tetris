package tetris

import (
	"reflect"
	"testing"
)

func TestNewField(t *testing.T) {
	figure := figure{
		[]Level{Space, Space, Space, Space, Space},
		[]Level{Space, Space, Space, Space, Space},
		[]Level{Space, Space, Space, Space, Space},
		[]Level{Space, Space, Space, Space, Space},
		[]Level{Space, Space, Space, Space, Space},
		[]Level{Space, Space, Space, Space, Space},
		[]Level{Space, Space, Space, Space, Space},
		[]Level{Space, Space, Space, Space, Space},
		[]Level{Space, Space, Space, Space, Space},
		[]Level{Space, Space, Space, Space, Space},
	}
	w := 5
	h := 10
	field := newField(w, h)
	if field.w != w {
		t.Error("field width is incorrect\n")
		t.Errorf("have %v, but want %v", field.w, w)
		t.Error("set it properly in newField method\n")
	}
	if field.h != h {
		t.Error("field height is incorrect\n")
		t.Errorf("have %v, but want %v", field.h, h)
		t.Error("set it properly in newField method\n")
	}
	if !reflect.DeepEqual(field.figure, figure) {
		t.Error("field figure is incorrect\n")
		t.Error("field figure by newField should be filled with Space\n")
		t.Error("set it properly in newField method\n")
	}
}

func TestHaveConfliction(t *testing.T) {
	field := field{
		figure: figure{
			[]Level{Space, Space, Space, Space, Space},
			[]Level{Block, Block, Block, Space, Space},
		},
	}
	tables := []struct {
		in   frame
		want bool
	}{
		{
			in: frame{
				figure: figure{
					[]Level{Block, Block, Block, Block},
				},
				p: point{
					x: 0,
					y: 0,
				},
				w: 4,
				h: 1,
			},
			want: false,
		},
		{
			in: frame{
				figure: figure{
					[]Level{Block, Block},
					[]Level{Block, Block},
				},
				p: point{
					x: 3,
					y: 0,
				},
				w: 2,
				h: 2,
			},
			want: false,
		},
		{
			in: frame{
				figure: figure{
					[]Level{Block, Block, Block},
					[]Level{Block, Space, Space},
				},
				p: point{
					x: 0,
					y: 0,
				},
				w: 3,
				h: 2,
			},
			want: true,
		},
		{
			in: frame{
				figure: figure{
					[]Level{Space, Block, Block},
					[]Level{Block, Block, Space},
				},
				p: point{
					x: 2,
					y: 0,
				},
				w: 3,
				h: 2,
			},
			want: true,
		},
	}

	for _, table := range tables {
		have := field.haveConfliction(table.in)
		if have != table.want {
			t.Errorf("have %v, but want %v", have, table.want)
			t.Errorf("when in is %v", table.in)
		}
	}
}

func TestPut(t *testing.T) {
	f := field{
		figure: figure{
			[]Level{Space, Space, Space, Space, Space},
			[]Level{Block, Block, Block, Space, Space},
		},
	}
	tables := []struct {
		in   frame
		want field
	}{
		{
			in: frame{
				figure: figure{
					[]Level{Block, Block, Block, Block},
				},
				p: point{
					x: 0,
					y: 0,
				},
				w: 4,
				h: 1,
			},
			want: field{
				figure: figure{
					[]Level{Block, Block, Block, Block, Space},
					[]Level{Block, Block, Block, Space, Space},
				},
			},
		},
		{
			in: frame{
				figure: figure{
					[]Level{Block, Block},
					[]Level{Block, Block},
				},
				p: point{
					x: 3,
					y: 0,
				},
				w: 2,
				h: 2,
			},
			want: field{
				figure: figure{
					[]Level{Space, Space, Space, Block, Block},
					[]Level{Block, Block, Block, Block, Block},
				},
			},
		},
	}

	for _, table := range tables {
		f.put(table.in)
		if !reflect.DeepEqual(f, table.want) {
			t.Errorf("have %v, but want %v", f, table.want)
			t.Errorf("when in is %v", table.in)
		}

		// reset
		f = field{
			figure: figure{
				[]Level{Space, Space, Space, Space, Space},
				[]Level{Block, Block, Block, Space, Space},
			},
		}
	}
}
