package main

import (
	"encoding/json"
	"fmt"
)

type emp struct {
	Name    string `json:"Employe name"` //print to Employe name
	Age     int
	Salary  int      `json:"-"`              //Dont print
	Hobbies []string `json:"tags,omitempty"` //if nil no print
}

func main() {
	Employee := []emp{
		{"Harish", 23, 20000, []string{"Kabbadi", "Dance"}},
		{"Chitti", 25, 22000, []string{"Drawing", "Using phone"}},
		{"Harish", 26, 30000, nil},
	}
	finalJson, err := json.MarshalIndent(Employee, "", "\t")
	if err != nil {
		panic(err)
	}
	fmt.Println("Require JSON:", string(finalJson))
}
