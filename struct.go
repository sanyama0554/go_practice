package main

import "fmt"

type Person struct {
	name string
	age int
}

func (p *Person) SetPerson(name string, age int) {
	p.name = name
	p.age = age
}

func (p *Person) GetPerson() (string, int) {
	return p.name, p.age
}

func main() {
	var p1 Person
	p1.SetPerson("Alice", 30)
	name, age := p1.GetPerson()
	fmt.Printf("Name: %s, Age: %d\n", name, age)
}