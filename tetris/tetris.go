package tetris

import (
	"github.com/tomocy/go-tetris/setting"
)

type Tetris struct {
	field  Field
	quitCh chan bool
}

func NewTetris() *Tetris {
	t := new(Tetris)
	t.setUp()

	return t
}

func (t *Tetris) setUp() {
	t.field = MakeField(setting.Field.Width(), setting.Field.Height())
	t.quitCh = make(chan bool)
}

func (t *Tetris) Start() {
	for i := 0; i < len(t.field); i++ {
		for j := 0; j < len(t.field[i]); j++ {
			t.field[i][j] = block
		}
	}

	t.quit()
}

func (t *Tetris) quit() {
	t.quitCh <- true
}

func (t *Tetris) Waiting() {
	<-t.quitCh
}
func (t Tetris) String() string {
	fieldBytes := t.field.Bytes()

	return string(fieldBytes)
}
