package main

import "fmt"

func main() {
	harish := User{"Sirikoti Harish", "harish@gmail.com", true, 23}
	fmt.Println(harish)
	fmt.Printf("Details of Harish: %+v\n", harish)
	fmt.Printf("Name : %v, Mail: %v", harish.Name, harish.Mail)
}

type User struct {
	Name   string
	Mail   string
	Status bool
	Age    int
}
