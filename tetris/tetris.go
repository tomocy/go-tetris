package tetris

import (
	"errors"
	"math/rand"
	"time"
)

type Tetris interface {
	Start()
}

func New(w, h int) Tetris {
	return newTetris(w, h)
}

type tetris struct {
	field            Field
	tetrominos       []Tetromino
	currentTetromino Tetromino
	quit             chan bool
}

func newTetris(w, h int) *tetris {
	t := new(tetris)
	t.field = newField(w, h)
	t.tetrominos = []Tetromino{
		newI(), newO(), newZ(),
		newT(), newL(),
	}
	t.quit = make(chan bool)

	return t
}

func (t *tetris) Start() {
	t.dropRandomTetromino()
	t.update()
}

func (t *tetris) dropRandomTetromino() {
	tetromino := t.pickRandomTetromino()
	if err := t.putTetrominoOrError(tetromino); err != nil {
		t.quit <- true
	}
}

func (t tetris) pickRandomTetromino() Tetromino {
	rand.Seed(time.Now().UnixNano())
	n := rand.Intn(len(t.tetrominos))

	return t.tetrominos[n]
}

func (t *tetris) update() {
	ticker := time.NewTicker(1 * time.Second)
	defer ticker.Stop()
	for {
		select {
		case <-ticker.C:
			next := t.currentTetromino.asMoved(Down)
			if err := t.putTetrominoOrError(next); err != nil {
				t.dropRandomTetromino()
			}
		}
	}
}

func (t *tetris) putTetrominoOrError(tetromino Tetromino) error {
	if !tetromino.doesExistInColumn(t.field.height()) {
		return errors.New("invalid y index of tetromnio")
	}
	if !tetromino.doesExistInRow(t.field.width()) {
		frame := tetromino.frame()
		if frame.p.x < 0 {
			tetromino.move(Right)
		} else {
			tetromino.move(Left)
		}
	}

	if t.field.haveConflict(tetromino.frame()) {
		return errors.New("conflict happen")
	}

	t.putTetromino(tetromino)
	return nil
}

func (t *tetris) putTetromino(tetromino Tetromino) {
	t.field.put(tetromino.frame())
	t.currentTetromino = tetromino
}
