package main

import (
	"fmt"
	"os"
)

type World struct {
	Height int
	Width  int
	Cells  [][]bool
}

func (w *World) SaveState(filename string) error {
	file, err := os.Create(filename)
	if err != nil {
		return err
	}
	defer file.Close()

	for i, row := range w.Cells {
		for _, cell := range row {
			var value uint8
			if cell {
				value = 1
			}
			fmt.Fprintf(file, "%d", value)
			if err != nil {
				return err
			}
		}
		if i < len(w.Cells)-1 {
			fmt.Fprintln(file)
		}
	}

	return nil
}
