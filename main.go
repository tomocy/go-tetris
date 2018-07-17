package main

import "fmt"

func main() {
	// tc := NewTetrisController()
	// quitCh := make(chan bool)

	// tc.Start()
	// <-quitCh
	w := MakeWorld(5, 7)
	w[0][1] = Block
	for _, y := range w {
		for _, x := range y {
			fmt.Print(x)
		}
		fmt.Println()
	}

	f := Frame{p: Point{x: 0, y: 0}, w: len(w[0]), h: len(w)}
	w.clear(f)
	for _, y := range w {
		for _, x := range y {
			fmt.Print(x)
		}
		fmt.Println()
	}
}

func Render(t *Tetris) {
	for _, y := range t.wrld {
		for _, x := range y {
			print(x)
		}
		println()
	}
}
