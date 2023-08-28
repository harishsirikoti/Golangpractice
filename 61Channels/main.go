package main

import "fmt"

func printHello(done chan bool) {
	fmt.Println("Hellow World")
	done <- true
}
func main() {
	done := make(chan bool)
	go printHello(done)
	<-done //blocked till data is set for done variable
	fmt.Println("Main print")
}
