package main

import "fmt"

type animal interface {
	mover() string
}

type perro struct{}

type pez struct{}

type pajaro struct{}

func (perro) mover() string {
	return "Soy un perro y camino"
}

func (pez) mover() string {
	return "Soy un pez y nado"
}

func (pajaro) mover() string {
	return "Soy un pajaro y estoy volando"
}

func moverAnimal(a animal) {
	fmt.Println(a.mover())
}

func main() {
	p := perro{}

	pe := pez{}

	pa := pajaro{}

	moverAnimal(p)
	moverAnimal(pe)
	moverAnimal(pa)
}
