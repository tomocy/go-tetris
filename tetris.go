package main

import (
	"fmt"
	"math/rand"
	"time"

	termbox "github.com/nsf/termbox-go"
)

// Command is ...
type Command string

// Left is ...
// Right is ...
// Down is ...
const (
	Left  Command = "j"
	Right Command = "l"
	Down  Command = "k"
)

// Tetris is ...
type Tetris struct {
	diffs       map[Command]Diff
	tetrs       []Tetromino
	w           World
	currentTetr Tetromino
	stopCh      chan bool
	cmdCh       chan Command
	rotateCh    chan bool
}

func NewTetris() *Tetris {
	diffs := map[Command]Diff{
		Left:  Diff{x: -1, y: 0},
		Right: Diff{x: 1, y: 0},
		Down:  Diff{x: 0, y: 1},
	}
	tetrs := []Tetromino{
		MakeI(Head), MakeO(Head),
		MakeZ(Head), MakeT(Head),
		MakeL(Head),
	}
	w := MakeWorld(10, 10)
	stopCh := make(chan bool)
	cmdCh := make(chan Command)
	rotateCh := make(chan bool)

	return &Tetris{diffs: diffs, tetrs: tetrs, w: w, stopCh: stopCh, cmdCh: cmdCh, rotateCh: rotateCh}
}

func (t *Tetris) Start() {
	t.currentTetr = t.GenerateTetromino(time.Now().UnixNano())
	t.w.add(t.currentTetr)
	Render(t.w)

	go t.Update()
	go t.WaitCommand()
}

func (t Tetris) GenerateTetromino(seed int64) Tetromino {
	tetr := t.PickTetromino(seed)
	tetr.frame.p = t.GetStartPoint(tetr, seed)

	return tetr
}

func (t Tetris) PickTetromino(seed int64) Tetromino {
	rand.Seed(seed)
	n := rand.Intn(len(t.tetrs))

	return t.tetrs[n]
}

func (t Tetris) GetStartPoint(tetr Tetromino, seed int64) Point {
	rand.Seed(seed)
	x := rand.Intn(len(t.w[0]) - tetr.frame.w)

	return Point{x: x, y: 0}
}

func (t *Tetris) Update() {
	ticker := time.NewTicker(1 * time.Second)
	defer ticker.Stop()
	for {
		select {
		case cmd := <-t.cmdCh:
			t.Process(cmd)
		case <-t.rotateCh:
			t.w.clear(t.currentTetr)
			t.currentTetr.Rotate()
		case <-ticker.C:
			t.Process(Down)
		}

		Render(t.w)
	}
}

func (t *Tetris) Process(cmd Command) {
	prevTetr := t.currentTetr
	diff := t.diffs[cmd]
	t.w.clear(t.currentTetr)
	t.currentTetr.moveFor(diff)
	if t.currentTetr.frame.p.x < 0 || len(t.w[0]) <= t.currentTetr.frame.p.x+t.currentTetr.frame.w-1 {
		t.currentTetr.frame.p.x = prevTetr.frame.p.x
	}
	if err := t.w.add(t.currentTetr); err != nil {
		t.currentTetr.moveFor(Diff{x: 0, y: -1})
		t.w.add(t.currentTetr)

		t.w.deleteLines()

		t.currentTetr = t.GenerateTetromino(time.Now().UnixNano())
		if err := t.w.add(t.currentTetr); err != nil {
			fmt.Println("Game Over!!!!")
			t.stopCh <- true
		}
	}
}

func (t *Tetris) WaitCommand() {
	for {
		switch ev := termbox.PollEvent(); ev.Type {
		case termbox.EventKey:
			switch ev.Key {
			case termbox.KeyArrowLeft:
				t.cmdCh <- Left
			case termbox.KeyArrowRight:
				t.cmdCh <- Right
			case termbox.KeyArrowDown:
				t.cmdCh <- Down
			case termbox.KeySpace:
				t.rotateCh <- true
			case termbox.KeyCtrlC:
				t.stopCh <- true
			}
		}
	}
}
