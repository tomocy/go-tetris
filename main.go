package main

import (
	"time"

	termbox "github.com/nsf/termbox-go"
	"github.com/tomocy/go-tetris/tetris"
)

type TetrisController struct {
	tetris *tetris.Tetris
	rendCh chan bool
}

func NewTetrisController() *TetrisController {
	tc := new(TetrisController)
	tc.setUp()

	return tc
}

func (tc *TetrisController) setUp() {
	tc.tetris = tetris.NewTetris()
	tc.rendCh = make(chan bool)
}

func main() {
	err := termbox.Init()
	if err != nil {
		panic(err)
	}
	defer termbox.Close()

	tetrisController := NewTetrisController()
	tetrisController.startTetris()
	tetrisController.waitForTetrisToEnd()
}

func (tc *TetrisController) startTetris() {
	go tc.tetris.Start()
	go tc.waitCommand()
	time.Sleep(5 * time.Millisecond)
	go tc.renderPerUpdate()
}

func (tc *TetrisController) waitCommand() {
	for {
		switch ev := termbox.PollEvent(); ev.Type {
		case termbox.EventKey:
			switch ev.Key {
			case termbox.KeyArrowLeft:
				tc.tetris.Receive(tetris.Left)
				tc.rendCh <- true
			case termbox.KeyArrowRight:
				tc.tetris.Receive(tetris.Right)
				tc.rendCh <- true
			case termbox.KeyArrowDown:
				tc.tetris.Receive(tetris.Down)
				tc.rendCh <- true
			case termbox.KeyCtrlC:
				tc.tetris.Quit()
			}
		}
	}
}

func (tc *TetrisController) renderPerUpdate() {
	tc.render()
	ticker := time.NewTicker(1 * time.Second)
	defer ticker.Stop()
	for {
		select {
		case <-tc.rendCh:
			tc.render()
		case <-ticker.C:
			tc.render()
		}
	}
}

func (tc *TetrisController) render() {
	res := tc.tetris.Strings()
	termbox.Clear(termbox.ColorDefault, termbox.ColorDefault)
	termbox.SetCursor(0, len(res)+1)
	for i := 0; i < len(res); i++ {
		for j := 0; j < len(res[i]); j++ {
			termbox.SetCell(j, i, []rune(res[i][j])[0], termbox.ColorDefault, termbox.ColorDefault)
		}
	}
	termbox.Flush()
}

func (tc *TetrisController) waitForTetrisToEnd() {
	tc.tetris.Waiting()
}
