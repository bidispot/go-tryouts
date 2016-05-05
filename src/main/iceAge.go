package main

import "fmt"

type Animal interface {
	String() string
}

type Action interface {
	Do()
}

type AnimalAction interface {
	Animal
	Action
}

type Identity struct {
	name string
}

type Mammoth struct {
	Identity
	Action
}

type Squirrel struct {
	Identity
	Action
}

type Eat struct {
	Food string
}

type Travel struct {
	Goal string
}

func main() {
	m := Mammoth{Identity{"Manny"}, Travel{"San Francisco"}}
	s := Squirrel{Identity{"Scrat"}, Eat{"Nut"}}
	ms := []AnimalAction{&m, &s}

	for _, a := range ms {
		Talk(a)
		Do(a)
	}
}

func Talk(a Animal) {
	fmt.Println(a)
}

func Do(a Action) {
	a.Do()
}

func (e Eat) Do() {
	fmt.Println(e.Food + " is tasty!")
}

func (t Travel) Do() {
	fmt.Println("Going to " + t.Goal)
}

func (m Mammoth) String() string {
	return "Hello I'm a Mammoth and my name is " + m.name + "."
}

func (m Squirrel) String() string {
	return "Hello I'm a Squirrel and my name is " + m.name + "."
}