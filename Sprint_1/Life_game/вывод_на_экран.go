package main

type World struct {
	Height int
	Width  int
	Cells  [][]bool
}

func (w *World) String() string {
	brownSquare := "\xF0\x9F\x9F\xAB"
	greenSquare := "\xF0\x9F\x9F\xA9"
	res := ""

	for _, row := range w.Cells {
		for _, cell := range row {
			if cell {
				res += greenSquare
			} else {
				res += brownSquare
			}
		}
		res += "\n"
	}

	return res
}
