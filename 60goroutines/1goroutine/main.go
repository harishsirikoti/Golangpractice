package main

import (
	"fmt"
	"time"
)

func Numbers() {
	for i := 1; i <= 5; i++ {
		time.Sleep(250 * time.Millisecond)
		fmt.Printf("%d", i)
	}
}
func Alphabets() {
	for i := 'a'; i <= 'e'; i++ {
		time.Sleep(400 * time.Millisecond)
		fmt.Printf("%c", i)
	}
}
func main() {
	go Numbers()
	go Alphabets()
	time.Sleep(1 * time.Second) // telling to wait for 1sec
	fmt.Println("main terminated")
}
