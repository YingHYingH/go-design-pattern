package _5_proxy

import (
	"fmt"
	"testing"
)

type People interface {
	Eat()
	Say()
}

type Person struct {
	Age int
	age int
}

func (p *Person) Eat() {
	fmt.Printf("The address of x is: %p\n", p)
	p.Age = 10
}

func (p Person) Say() {
	fmt.Printf("The address of x is: %p\n", &p)
	p.age = 10
}

func TestPeople(t *testing.T) {
	//var p People
	////p := &Person{}
	////p.Say()
	////p.Eat()
	////
	//p = Person{}
	//
	////p := Person{}
	////p.Say()
	////p.Eat()
	////
	////pp := &Person{}
	////pp.Say()
	////pp.Eat()

	p := Person{}
	fmt.Println(p.Age)
	fmt.Printf("The address of x is: %p\n", &p)
	p.Say()
	fmt.Println(p.Age)
	p.Eat()
	fmt.Println(p.Age)
}
