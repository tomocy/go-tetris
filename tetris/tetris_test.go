package tetris

import (
	"reflect"
	"testing"
)

func TestNewTetris(t *testing.T) {
	want := getNewTetrisExpectation()
	have := newTetris(5, 2)
	if !reflect.DeepEqual(have.field, want.field) {
		t.Errorf("have %#v, but want %#v", have.field, want.field)
	}
	if !reflect.DeepEqual(have.tetrominos, want.tetrominos) {
		t.Errorf("have %#v, but want %#v", have.tetrominos, want.tetrominos)
	}
}

func getNewTetrisExpectation() *tetris {
	return &tetris{
		field: &field{
			figure: figure{
				[]Level{Space, Space, Space, Space, Space},
				[]Level{Space, Space, Space, Space, Space},
			},
			w: 5,
			h: 2,
		},
		tetrominos: []Tetromino{
			getI(), getO(), getZ(),
			getT(), getL(),
		},
	}
}
