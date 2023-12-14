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
			ni := (i + w.Height) % w.Height
			nj := (j + w.Width) % w.Width
			if w.Cells[ni][nj] {
				cnt_neighbors++
			}
		}
	}

	return cnt_neighbors
}
