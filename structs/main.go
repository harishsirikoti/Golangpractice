package main

import (
	"05mytime/structs/computer"
	"fmt"
)

type Person struct {
	name    string
	age     int
	address Address
} //nested structs
type Address struct {
	state string
	city  string
}

func main() {
	p1 := Person{name: "Harish", age: 23, address: Address{state: "Telangana", city: "Hyderbad"}}
	fmt.Println("p1", p1)
	fmt.Printf("p1 details %+v\n", p1)
	spec := computer.Spec{
		Maker: "apple",
		Price: 50000,
	}
	fmt.Println("Maker:", spec.Maker)
}
