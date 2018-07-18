package main

import (
	termbox "github.com/nsf/termbox-go"
)

func main() {
	err := termbox.Init()
	if err != nil {
		panic(err)
	}
	defer termbox.Close()

	t := NewTetris()
	t.Start()
	<-t.stopCh
}

func Render(w World) {
	termbox.Clear(termbox.ColorDefault, termbox.ColorDefault)
	for i := 0; i < len(w); i++ {
		for j := 1; j <= len(w[i]); j++ {
			if j == 1 {
				termbox.SetCell(0, i, []rune(WallStr)[0], termbox.ColorDefault, termbox.ColorDefault)
			}
			if j == len(w[i]) {
				termbox.SetCell(len(w[i])+1, i, []rune(WallStr)[0], termbox.ColorDefault, termbox.ColorDefault)
			}
			termbox.SetCell(j, i, []rune(w[i][j-1].String())[0], termbox.ColorDefault, termbox.ColorDefault)
		}
	}

	for j := 0; j < len(w[0])+2; j++ {
		termbox.SetCell(j, len(w), []rune(WallStr)[0], termbox.ColorDefault, termbox.ColorDefault)
	}

	termbox.SetCursor(0, len(w)+1)
	termbox.Flush()
}
