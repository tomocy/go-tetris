package main

import "github.com/tomocy/go-tetris/tetris"

func main() {
	// tetris start
	t := tetris.New(5, 10)
	t.Start()
}
