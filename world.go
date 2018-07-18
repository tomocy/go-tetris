package main

import "fmt"

// Level is ...
type Level int

// Space is ...
const (
	Space Level = iota
	Block
)

const (
	SpaceStr = " "
	BlockStr = "●"
	WallStr  = "×"
)

func (l Level) String() string {
	switch l {
	case Space:
		return SpaceStr
	case Block:
		return BlockStr
	default:
		return "?"
	}
}

// World is ...
type World [][]Level

// MakeWorld is ...
func MakeWorld(w, h int) World {
	wrld := make([][]Level, h)
	for i := 0; i < h; i++ {
		wrld[i] = make([]Level, w)
		for j := 0; j < w; j++ {
			wrld[i][j] = Space
		}
	}

	return wrld
}

func (w World) clear(tetr Tetromino) {
	for i := 0; i < tetr.frame.h; i++ {
		for j := 0; j < tetr.frame.w; j++ {
			if tetr.figure[i][j] != Block {
				continue
			}

			w[tetr.frame.p.y+i][tetr.frame.p.x+j] = Space
		}
	}
}

func (w World) add(tetr Tetromino) error {
	if !w.canAdd(tetr) {
		return fmt.Errorf("invalid point")
	}

	for i := 0; i < tetr.frame.h; i++ {
		for j := 0; j < tetr.frame.w; j++ {
			if tetr.figure[i][j] != Block {
				continue
			}

			w[tetr.frame.p.y+i][tetr.frame.p.x+j] = Block
		}
	}

	return nil
}

func (w World) canAdd(tetr Tetromino) bool {
	if tetr.frame.p.y < 0 || len(w) <= tetr.frame.p.y+tetr.frame.h-1 {
		fmt.Println(111111)
		return false
	}
	if tetr.frame.p.x < 0 || len(w[0]) <= tetr.frame.p.x+tetr.frame.w-1 {
		fmt.Println(22222)
		return false
	}
	for i := 0; i < tetr.frame.h; i++ {
		for j := 0; j < tetr.frame.w; j++ {
			if w[tetr.frame.p.y+i][tetr.frame.p.x+j] == Block && tetr.figure[i][j] == Block {
				return false
			}
		}
	}

	return true
}

func (w World) deleteLines() {
	for i := 0; i < len(w); i++ {
		isFilled := true
		for j := 0; j < len(w[i]); j++ {
			if w[i][j] == Space {
				isFilled = false
				break
			}
		}

		if isFilled {
			for k := i; 0 < k; k-- {
				if k == 0 {
					for l := 0; l < len(w[k]); l++ {
						w[0][l] = Space
					}
					continue
				}
				w[k] = w[k-1]
			}
		}
	}
}

func (w World) String() string {
	b := make([]byte, 0, 10)
	for i := 0; i < len(w); i++ {
		for j := 0; j < len(w[i]); j++ {
			if j == 0 || j == len(w[i])-1 {
				b = append(b, BlockStr...)
			}
			b = append(b, w[i][j].String()...)
		}
		b = append(b, "\n"...)
	}

	for i := 0; i < len(w[0])+2; i++ {
		b = append(b, BlockStr...)
	}

	b = append(b, "\n"...)

	return string(b)
}
