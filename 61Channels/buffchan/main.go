package main

import (
	"fmt"
	"time"
)

func putvalue(ch chan int) {
	for i := 0; i < 5; i++ {
		ch <- i
		fmt.Println("value put into channel :", i)
	}
	close(ch)
}
func main() {
	ch := make(chan int, 2)
	go putvalue(ch)
	time.Sleep(2 * time.Second)
	for v := range ch {
		fmt.Println("Value received:", v)
		time.Sleep(2 * time.Second)
	}
}
