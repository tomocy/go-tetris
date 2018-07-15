package main

import (
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
	t.tetr = t.lotTetromino()
	p := t.getStartPoint()
	t.tetr.p = p
	t.setTetromino()
}

func (t *Tetris) ProcessTetromino(diff Diff) {
	p := t.getNextPoint(diff)
	if p.y < 0 || len(t.wrld) <= p.y+t.tetr.h-1 {
		return
	}
	if p.x < 0 || len(t.wrld[p.y]) <= p.x+t.tetr.w-1 {
		p.x = t.tetr.p.x
	}

	t.tetr.p = p
	t.resetWorld()
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

func (t *Tetris) resetWorld() {
	for i := 0; i < len(t.wrld); i++ {
		for j := 0; j < len(t.wrld[i]); j++ {
			t.wrld[i][j] = Space
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
