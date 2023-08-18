package main

import "fmt"

type Describer interface {
	Describe()
}
type Person struct {
	name string
	age  int
}
type Address struct {
	State string
	city  string
}

func (p Person) Describe() {
	fmt.Printf("%s is %d years old\n", p.name, p.age)
}
func (a Address) Describe() {
	fmt.Printf("State %s City %s \n", a.State, a.city)
}
func main() {
	var d1 Describer
	p1 := Person{"Harish", 23}
	d1 = p1
	d1.Describe()
	p2 := Person{"Chitti", 16}
	d1 = &p2
	d1.Describe()
	var d2 Describer
	a1 := Address{"Telangana", "Ghatkesar"}
	d2 = &a1
	d2.Describe()
}
