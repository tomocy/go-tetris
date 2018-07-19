package main

import (
	"fmt"

	"github.com/tomocy/go-tetris/tetris"
)

type TetrisController struct {
	tetris *tetris.Tetris
}

func NewTetrisController() *TetrisController {
	tc := new(TetrisController)
	tc.setUp()

	return tc
}

func (tc *TetrisController) setUp() {
	tc.tetris = tetris.NewTetris()
}

func main() {
	tetrisController := NewTetrisController()
	tetrisController.startTetris()
	tetrisController.waitForTetrisToEnd()
	tetrisController.Render()
}

func (tc *TetrisController) startTetris() {
	go tc.tetris.Start()
}

func (tc *TetrisController) waitForTetrisToEnd() {
	tc.tetris.Waiting()
}

func (tc *TetrisController) Render() {
	res := tc.String()

	fmt.Print(res)
}

func (tc *TetrisController) String() string {
	tetrisStr := tc.tetris.String()

	return tetrisStr
}
