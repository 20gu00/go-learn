package main

import (
	"fmt"
)

type People struct {
	age  *int
	name string
}

func (p People) SetAge(age int) {
	p.age = &age
}

func (p People) GetAge() int {
	return *p.age
}

func (p People) SetName(name string) {
	p.name = name
}

func (p People) GetName() string {
	return p.name
}
func NewPeople(name string, age int) (p *People) {
	p = new(People)
	p.age = new(int) // int的指针
	p.SetName(name)
	p.SetAge(age)
	return
}

func main() {
	var people *People = NewPeople("John", 22)
	people.SetName("Grace")
	people.SetAge(45)
	fmt.Printf("%s,%d", people.GetName(), people.GetAge())
}
