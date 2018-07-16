package main

import (
	"fmt"
	"math/rand"
	"time"
)

type Tetris struct {
	wrld  World
	tetrs []Tetromino
	tetr  Tetromino
}

func NewTetris() *Tetris {
	return &Tetris{
		wrld:  MakeWorld(5, 7),
		tetrs: []Tetromino{MakeI(), MakeO(), MakeZ(), MakeT(), MakeL()},
	}
}

func (t *Tetris) GenerateTetromino() {
	tetr := t.tetr
	t.tetr = t.lotTetromino()
	t.tetr.p = t.getStartPoint()
	if t.doesConflictHappen(t.tetr) {
		t.tetr = tetr
		fmt.Println("Game Over!!!!!!")
		return
	}

	t.setTetromino()
}

func (t *Tetris) ProcessTetromino(diff Diff) {
	tetr := t.tetr
	t.tetr.p = t.getNextPoint(diff)
	if t.tetr.p.y < 0 || len(t.wrld) <= t.tetr.p.y+t.tetr.h-1 {
		t.GenerateTetromino()
		return
	}
	if t.tetr.p.x < 0 || len(t.wrld[t.tetr.p.y]) <= t.tetr.p.x+t.tetr.w-1 {
		t.tetr.p.x = tetr.p.x
	}

	t.clear(t.tetr)
	if t.doesConflictHappen(t.tetr) {
		t.tetr = tetr
		t.GenerateTetromino()
	}

	t.setTetromino()
}

func (t *Tetris) setTetromino() {
	for i := 0; i < len(t.tetr.frame); i++ {
		for j := 0; j < len(t.tetr.frame[i]); j++ {
			if t.tetr.frame[i][j] == Space {
				continue
			}

			t.wrld[t.tetr.p.y+i][t.tetr.p.x+j] = Block
		}
	}

}

func (t *Tetris) clear(tetr Tetromino) {
	for i := 0; i < tetr.h; i++ {
		for j := 0; j < tetr.w; j++ {
			t.wrld[tetr.p.y+i][tetr.p.x+j] = Space
		}
	}
}

func (t *Tetris) lotTetromino() Tetromino {
	rand.Seed(time.Now().UnixNano())
	n := rand.Intn(len(t.tetrs))

	return t.tetrs[n]
}

func (t *Tetris) getStartPoint() Point {
	wrldLen := len(t.wrld[0])
	if wrldLen == t.tetr.w {
		return Point{x: 0, y: 0}
	}

	rand.Seed(time.Now().UnixNano())
	x := rand.Intn(wrldLen - t.tetr.w + 1)

	return Point{x: x, y: 0}
}

func (t *Tetris) getNextPoint(diff Diff) Point {
	return Point{
		x: t.tetr.p.x + diff.x,
		y: t.tetr.p.y + diff.y,
	}
}

func (t *Tetris) doesConflictHappen(tetr Tetromino) bool {
	for i := 0; i < tetr.h; i++ {
		for j := 0; j < tetr.w; j++ {
			if tetr.frame[i][j] == Block && t.wrld[tetr.p.y+i][tetr.p.x+j] == Block {
				return true
			}
		}
	}

	return false
}
