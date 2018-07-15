package main

type World [][]Level

func MakeWorld(w, h int) World {
	wrld := make(World, h)
	for i := 0; i < h; i++ {
		wrld[i] = make([]Level, w)
		for j := 0; j < w; j++ {
			wrld[i][j] = Space
		}
	}

	return wrld
}
