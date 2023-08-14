package main

import "fmt"

//Pointer Receivers vs Value Receivers
type Employee struct {
	Name string
	Age  int
}

func (n Employee) changename(name string) {
	n.Name = name
	fmt.Println(n.Name)
}
func (a *Employee) changeage(age int) { //using pointers chage of value is observed
	a.Age = age
}
func main() {
	Emp1 := Employee{Name: "Harish", Age: 23}
	fmt.Println("Emp name befor:", Emp1.Name)
	Emp1.changename("Sirikoti Harish")
	fmt.Println("\n Emp name befor:", Emp1.Name)
	fmt.Println("Emp age befor:", Emp1.Age)
	Emp1.changeage(25)
	fmt.Println("\n Emp age befor:", Emp1.Age)

}
