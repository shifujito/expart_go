package main

import "fmt"

type Person struct{}

type Footstepper interface {
	Footsteps() string
}

type Crier interface {
	Cry() string
}

type CryFootstepper interface {
	Crier
	Footstepper
}

func (p *Person) Cry() string {
	return "Hi"
}

func (p *Person) Footsteps() string {
	return "Pitapat"
}

type PartyPeople struct {
	Person
}

func (p *PartyPeople) Cry() string {
	return "Sup?"
}

var cf CryFootstepper

func main() {
	cf = &Person{}
	fmt.Println(cf.Cry(), cf.Footsteps())

	cf = &PartyPeople{}
	fmt.Println(cf.Cry(), cf.Footsteps())
}
