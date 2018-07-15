package main

func main() {
	tc := NewTetrisController()
	quitCh := make(chan bool)

	tc.Start()
	<-quitCh
}

func Render(t *Tetris) {
	for _, y := range t.wrld {
		for _, x := range y {
			print(x)
		}
		println()
	}
}
