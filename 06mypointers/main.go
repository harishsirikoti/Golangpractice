package main

import "fmt"

func main() {
	mynumber := 10
	var ptr = &mynumber
	fmt.Println("Pointer address:", ptr)
	fmt.Println("Pointer Value: ", *ptr)
	*ptr = *ptr + 5
	fmt.Println("New mynumber value: ", mynumber)
}
