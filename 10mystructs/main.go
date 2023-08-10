package main

import (
	"fmt"
	"strconv"
)

func main() {
	harish := &User{"Sirikoti Harish", "harish@gmail.com", true, 23}
	fmt.Println(harish.describe())
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

type Hi string

func Hello(name string) string {
	return "hello" + name
}

func (u *User) describe() string {
	return "I'm " + u.Name + " and i'm " + strconv.Itoa(u.Age) + "old"
}
