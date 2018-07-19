package tetris

import (
	"errors"
	"math/rand"
	"time"

	"github.com/tomocy/go-tetris/setting"
)

type Tetris struct {
	tetrominos       []*Tetromino
	field            *Field
	currentTetromino *Tetromino
	cmdCh            chan Command
	quitCh           chan bool
}

func NewTetris() *Tetris {
	t := new(Tetris)
	t.setUp()

	return t
}

func (t *Tetris) setUp() {
	t.tetrominos = []*Tetromino{
		NewI(),
	}
	t.field = NewField(setting.Field.Width(), setting.Field.Height())
	t.cmdCh = make(chan Command)
	t.quitCh = make(chan bool)
}

func (t *Tetris) Start() {
	t.dropTetromino()
	t.update()
}

func (t *Tetris) dropTetromino() {
	tetromino := t.newTetrominoRandomly()
	err := t.putTetromino(*tetromino)
	if err != nil {
		t.quit()
		return
	}
	t.reflectTetromino()
}

func (t Tetris) newTetrominoRandomly() *Tetromino {
	rand.Seed(time.Now().UnixNano())
	n := rand.Intn(len(t.tetrominos))
	return t.tetrominos[n]
}

func (t *Tetris) putTetromino(tetromino Tetromino) error {
	if !t.isValidY(tetromino) {
		return errors.New("over y index")
	}
	if !t.isValidX(tetromino) {
		tetromino.frame.point = t.currentTetromino.frame.point
	}

	if t.currentTetromino != nil {
		t.removeTetromino(*t.currentTetromino)
	}
	t.currentTetromino = &tetromino
	return nil
}

func (t Tetris) isValidY(tetromino Tetromino) bool {
	if tetromino.frame.point.y < 0 {
		return false
	}
	if t.field.h <= tetromino.frame.point.y+tetromino.frame.h-1 {
		return false
	}

	return true
}

func (t Tetris) isValidX(tetromino Tetromino) bool {
	if tetromino.frame.point.x < 0 {
		return false
	}
	if t.field.w <= tetromino.frame.point.x+tetromino.frame.w-1 {
		return false
	}

	return true
}

func (t *Tetris) removeTetromino(tetromino Tetromino) {
	for i := 0; i < tetromino.frame.h; i++ {
		for j := 0; j < tetromino.frame.w; j++ {
			if tetromino.figure[i][j] != block {
				continue
			}
			y := tetromino.frame.point.y + i
			x := tetromino.frame.point.x + j
			t.field.figure[y][x] = space
		}
	}
}

func (t *Tetris) reflectTetromino() {
	for i := 0; i < t.currentTetromino.frame.h; i++ {
		for j := 0; j < t.currentTetromino.frame.w; j++ {
			if t.currentTetromino.figure[i][j] != block {
				continue
			}

			y := t.currentTetromino.frame.point.y + i
			x := t.currentTetromino.frame.point.x + j
			t.field.figure[y][x] = block
		}
	}
}

func (t *Tetris) update() {
	ticker := time.NewTicker(1 * time.Second)
	defer ticker.Stop()
	for {
		select {
		case cmd := <-t.cmdCh:
			err := t.follow(cmd)
			if err != nil {
				t.quit()
			}
		case <-ticker.C:
			err := t.follow(Down)
			if err != nil {
				t.quit()
			}
		}

		t.reflectTetromino()
	}
}

func (t *Tetris) follow(cmd Command) error {
	nextTetromino := *t.currentTetromino
	switch cmd {
	case Left:
		nextTetromino.frame.point.x -= 1
	case Right:
		nextTetromino.frame.point.x += 1
	case Down:
		nextTetromino.frame.point.y += 1
	}

	err := t.putTetromino(nextTetromino)
	return err
}

func (t *Tetris) Quit() {
	t.quit()
}

func (t *Tetris) quit() {
	t.quitCh <- true
}

func (t *Tetris) Waiting() {
	<-t.quitCh
}

func (t *Tetris) Receive(cmd Command) {
	t.cmdCh <- cmd
}

func (t Tetris) String() string {
	fieldBytes := t.field.Bytes()
	return string(fieldBytes)
}

func (t Tetris) Strings() [][]string {
	res := make([][]string, t.field.h)
	for i := 0; i < t.field.h; i++ {
		res[i] = make([]string, t.field.w)
		for j := 0; j < t.field.w; j++ {
			res[i][j] = t.field.figure[i][j].String()
		}
	}

	return res
}
