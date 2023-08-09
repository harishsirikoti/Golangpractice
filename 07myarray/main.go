package main

import "fmt"

func main() {
	var fruits [5]string
	fruits[0] = "Apple"
	fruits[1] = "Banana"
	fruits[3] = "Mango"
	fruits[4] = "grapes"
	var vegetables = [4]string{"Carrot", "radish", "potato", "beans"}
	fmt.Println("Array of Fruits :", fruits)
	fmt.Println("Length of Fruits :", len(fruits)) //we have only 4 but it gives 5(output)
	fmt.Println("Array of Vegetables :", vegetables[2])
	fmt.Println("Length of Vegetables :", len(vegetables))

}
