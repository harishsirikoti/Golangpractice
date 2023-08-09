package main

import (
	"fmt"
	"sort"
)

func main() {
	var fruits = []string{"apple", "banana", "grapes", "Mango"}
	fmt.Printf("Type of fruits: %T\n", fruits)
	fruits = append(fruits, "apple2.0", "mango2.0")
	fruits = append(fruits[2:5])
	fmt.Println("Fruits after slice: ", fruits)
	fmt.Println("Length of fruits: ", len(fruits)) // Length will change after slice

	Scores := make([]int, 4) //another syntax
	Scores[0] = 27697
	Scores[1] = 67
	Scores[2] = 8779
	Scores[3] = 277
	Scores = append(Scores, 34, 123)
	sort.Ints(Scores) //Sorting of slicers
	fmt.Println("To know slices are sorted: ", sort.IntsAreSorted(Scores))
	fmt.Println(Scores)
	//To remove some elements in array
	courses := []string{"React", "Java", "Python", "Ruby", "C++", "Shift"}
	index := 2
	courses = append(courses[:index], courses[index+1:]...) //triple dots are imp
	fmt.Println("Removed index value from slice: ", courses)
}
