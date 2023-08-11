package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	rand.Seed(time.Now().UnixNano())
	randomnum := rand.Intn(6) + 1
	fmt.Println(randomnum)
	switch randomnum {
	case 1:
		fmt.Println("Die value 1")
	case 2:
		fmt.Println("Die value 2")
	case 3:
		fmt.Println("Die value 3")
		fallthrough // it will allow you to go for next case
	case 4:
		fmt.Println("Die value 4")
	case 5:
		fmt.Println("Die value 5")
	case 6:
		fmt.Println("Die value 6")
		fallthrough
	default:
		fmt.Println("Error in random gen")
	}
}
