package main

import (
	"bufio"
	"errors"
	"os"
	"strconv"
)

type World struct {
	Height int
	Width  int
	Cells  [][]bool
}

func (w *World) LoadState(filename string) error {
	file, err := os.Open(filename)
	if err != nil {
		return err
	}
	defer file.Close()

	cells := make([][]bool, w.Height)

	scanner := bufio.NewScanner(file)
	var firstLineLength int

	for scanner.Scan() {
		cells_row := make([]bool, w.Width)
		line := scanner.Text()

		if len(line) != firstLineLength && firstLineLength != 0 {
			return errors.New("Строки в файле имеют разную длину.")
		} else {
			firstLineLength = len(line)
		}

		for _, n := range line {
			cell, err := strconv.ParseBool(string(n))
			if err != nil {
				return err
			}
			cells_row = append(cells_row, cell)
		}
		cells = append(cells, cells_row)
	}

	if err := scanner.Err(); err != nil {
		return err
	}

	w.Cells = cells

	return nil
}
