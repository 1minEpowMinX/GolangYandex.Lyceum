package main

import (
	"errors"
	"fmt"
)

type Grasshopper struct {
	position int
} // знает своё местоположение на линейке

func (g *Grasshopper) WhereAmI() int {
	return g.position
}

func (g *Grasshopper) Jump() (int, error) {
	if target == g.position {
		return g.position, errors.New("Кузнечик уже ест зернышко")
	}

	dist := target - g.position
	if dist < 0 { // если зернышко слева
		if dist >= -5 {
			g.position = target
		} else {
			g.position -= 5
		}
	} else { // если зернышко справа
		if dist <= 5 {
			g.position = target
		} else {
			g.position += 5
		}
	}

	return g.position, nil
}

type Jumper interface {
	WhereAmI() int      // выводит текущее положение кузнечика на линейке
	Jump() (int, error) // кузнечик прыгает к зерну. Выводит новое положение кузнечика, или ошибку, если он уже ест зерно
}

func PlaceJumper(place, target int) Jumper {
	return &Grasshopper{position: place}
}

const (
	place  = 20
	target = 13
)

func main() {
	g := PlaceJumper(place, target)
	fmt.Println(g.WhereAmI())
	for {
		currPlace, err := g.Jump()
		if err != nil {
			break
		}
		fmt.Println(currPlace)
	}
}
