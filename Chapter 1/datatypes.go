package main

import (
	"fmt"
)

type Person struct {
	Name string
	Age  int
}

type Dog struct {
	Name string
	Age  int
}

type Friend interface {
	SayHello()
}

func (p *Person) SayHello() {
	fmt.Println("Hello,", p.Name)
}

func (p *Dog) SayHello() {
	fmt.Println(p.Name, "says \"Woof woof!\"")
}

func Greet(f Friend) {
	f.SayHello()
}

func main() {
	var guy = new(Person)
	guy.Name = "Abdul"
	var pupper = new(Dog)
	pupper.Name = "Rocko"
	guy.SayHello()
	Greet(pupper)
}
