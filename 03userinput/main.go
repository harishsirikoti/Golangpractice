package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	fmt.Println("Welcom to user Input")
	reader := bufio.NewReader(os.Stdin)
	fmt.Println("Enter your rating:")
	input, _ := reader.ReadString('\n')
	fmt.Println("Your rating:", input)
	fmt.Printf("Type of rating: %T", input)
}
