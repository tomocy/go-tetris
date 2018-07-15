package main

import (
	"bufio"
	"fmt"
	"os"
	"time"
)

type TetrisController struct {
	t     *Tetris
	diffs map[Command]Diff
	cmdCh chan Command
}

func NewTetrisController() *TetrisController {
	diffs := map[Command]Diff{
		Left:  Diff{x: -1, y: 0},
		Right: Diff{x: 1, y: 0},
		Down:  Diff{x: 0, y: 1},
	}
	return &TetrisController{
		t:     NewTetris(),
		diffs: diffs,
		cmdCh: make(chan Command),
	}
}

func (tc *TetrisController) Start() {
	tc.t.GenerateTetromino()
	tc.Render()
	go tc.update()
	go tc.waitCommand()
}

func (tc *TetrisController) update() {
	ticker := time.NewTicker(1 * time.Second)
	defer ticker.Stop()
	for {
		select {
		case cmd := <-tc.cmdCh:
			tc.Process(cmd)
		case <-ticker.C:
			tc.Process(Down)
		}

		fmt.Println("--------")
		tc.Render()
	}
}

func (tc *TetrisController) waitCommand() {
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		cmd := Command(scanner.Text())
		tc.cmdCh <- cmd
	}
}

func (tc *TetrisController) Process(cmd Command) {
	diff, ok := tc.diffs[cmd]
	if !ok {
		panic("Ooops, Process: cmd not found")
	}

	tc.t.ProcessTetromino(diff)
}
func (tc *TetrisController) Render() {
	Render(tc.t)
}
