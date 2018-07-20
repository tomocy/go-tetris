package tetris

import (
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
}

func newTetris(w, h int) *tetris {
	t := new(tetris)
	t.field = newField(w, h)
	t.tetrominos = []Tetromino{
		newI(), newO(), newZ(),
		newT(), newL(),
	}

	return t
}

func (t *tetris) Start() {
	t.dropRandomTetromino()
	// move tetromino down per sec
	// if command is sent, follow the command
	t.update()
}

func (t *tetris) dropRandomTetromino() {
	tetromino := t.pickRandomTetromino()

	if !tetromino.doesExistInColumn(t.field.height()) {
		panic("over y index")
	}
	if !tetromino.doesExistInRow(t.field.width()) {
		panic("over x index")
	}
	if t.field.haveConfliction(tetromino.get()) {
		panic("conflict happen")
	}

	t.field.put(tetromino.get())
	t.currentTetromino = tetromino
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
			// get next tetromino
			next := t.currentTetromino.getMoving()
			// validate the point
			if !next.doesExistInColumn(t.field.height()) {

			}
			if !next.doesExistInRow(t.field.width()) {

			}
			// check if collisions happen
			if t.field.haveConfliction(next.get()) {

			}
			// if not happen
			// move in direction
			t.currentTetromino.move(Down)
			// put it in field
			t.field.put(t.currentTetromino.get())
		}
	}
}
