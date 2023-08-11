package main

import (
	"fmt"
)

func main() {
	harish := &User{"Sirikoti Harish", "harish@gmail.com", true, 23}
	fmt.Println(harish)
	fmt.Printf("Details of Harish: %+v\n", harish)
	fmt.Printf("Name : %v, Mail: %v", harish.Name, harish.Mail)
	harish.NamePt()
	harish.MailPt()
}

type User struct {
	Name   string
	Mail   string
	Status bool
	Age    int
}

func (u User) NamePt() {
	fmt.Println("Name is :", u.Name)
}
func (u User) MailPt() {
	u.Mail = "hari@gmail.com"
	fmt.Println("New Mail is ", u.Mail)
}
