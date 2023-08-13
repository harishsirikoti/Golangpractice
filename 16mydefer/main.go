package main

import "fmt"

func main() {
	defer defertest()
	defer fmt.Println("First")
	defer fmt.Println("Second")
	defer fmt.Println("Third")
	fmt.Println("Hi all")

}
func defertest() {
	defer fmt.Println("Defer Testing function")
	for i := 0; i < 5; i++ {
		defer fmt.Println(i)
	}
}
