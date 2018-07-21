package tetris

import (
	"reflect"
	"testing"
)

func TestNewField(t *testing.T) {
	want := getNewFieldExpectation()
	have := newField(5, 2)
	if !reflect.DeepEqual(have, want) {
		t.Errorf("have %#v, but want %#v", have, want)
	}
}

func getNewFieldExpectation() *field {
	return &field{
		figure: figure{
			[]Level{Space, Space, Space, Space, Space},
			[]Level{Space, Space, Space, Space, Space},
		},
		w: 5,
		h: 2,
	}
}

func TestHaveConflict(t *testing.T) {
	given := getField()
	thens := getHaveConflictTestCases()
	for _, then := range thens {
		have := given.haveConflict(then.in)
		if have != then.want {
			t.Errorf("have %v, but want %v", have, then.want)
			t.Errorf("when in is %v", then.in)
		}
	}
}

func getHaveConflictTestCases() []struct {
	in   frame
	want bool
} {
	return []struct {
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
}

func TestPut(t *testing.T) {
	given := getField()
	thens := getPutTestCases()
	for _, then := range thens {
		given.put(then.in)
		if !reflect.DeepEqual(given, then.want) {
			t.Errorf("have %v, but want %v", given, then.want)
			t.Errorf("when in is %v", then.in)
		}

		given = getField()
	}
}

func getPutTestCases() []struct {
	in   frame
	want field
} {
	return []struct {
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
}

func getField() field {
	return field{
		figure: figure{
			[]Level{Space, Space, Space, Space, Space},
			[]Level{Block, Block, Block, Space, Space},
		},
	}
}
