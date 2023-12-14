package main

type World struct {
	Height int
	Width  int
	Cells  [][]bool
}

func (w *World) Neighbors(x, y int) int {
	cnt_neighbors := -1

	for i := x - 1; i <= x+1; i++ {
		for j := y - 1; j <= y+1; j++ {
			if i > -1 && j > -1 && w.Cells[i][j] {
				cnt_neighbors++
			}
		}
	}

	return cnt_neighbors
}
