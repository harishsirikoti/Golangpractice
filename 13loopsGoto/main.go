package main

import "fmt"

func main() {
	weekdays := []string{"Mon", "Tue", "Wed", "Thu", "Fri", "Sat"}

	for i, day := range weekdays {
		fmt.Printf("Index %v Day %v\n", i, day)
	}
	num := 1
	for num < 10 {
		if num == 5 {
			goto somelink
		}
		fmt.Println(num)
		num++
	}
somelink:
	fmt.Println("this is goto")
}
