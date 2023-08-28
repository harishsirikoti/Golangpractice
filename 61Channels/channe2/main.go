package main

import "fmt"

func square(number int, squareoc chan int) {
	sum := 0
	for number != 0 {
		digit := number % 10
		sum += digit * digit
		number /= 10
	}
	squareoc <- sum
}
func cube(number int, cubeoc chan int) {
	sum := 0
	for number != 0 {
		digit := number % 10
		sum += digit * digit * digit
		number /= 10
	}
	cubeoc <- sum
}
func main() {
	number := 589
	squarech := make(chan int)
	cubech := make(chan int)
	go cube(number, cubech)
	go square(number, squarech)
	sqrest, cuberest := <-squarech, <-cubech
	fmt.Println("Total: ", sqrest+cuberest)
}
