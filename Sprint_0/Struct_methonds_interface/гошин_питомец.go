package main

import "fmt"

type Animal interface {
	MakeSound()
}

type Cat struct {
}

type Dog struct {
}

func (c Cat) MakeSound() {
	fmt.Println("Мяу!")
}

func (d Dog) MakeSound() {
	fmt.Println("Гав!")
}
