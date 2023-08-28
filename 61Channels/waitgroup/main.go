package main

import (
	"fmt"
	"sync"
	"time"
)

func process(num int, wg *sync.WaitGroup) {
	fmt.Println("Started:", num)
	time.Sleep(2 * time.Second)
	fmt.Println("Ended:", num)
	wg.Done()
}
func main() {
	var wg sync.WaitGroup
	for i := 0; i < 3; i++ {
		wg.Add(1)
		go process(i, &wg)
	}
	wg.Wait()
	fmt.Println("Program is ended")
}
