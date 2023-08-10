package main

import "fmt"

func Changelocal(subjects [5]string) {
	subjects[0] = "hindi"
	fmt.Println("inside function:", subjects)

}

func main() {
	subjects := []string{"eng", "telugu", "maths", "science"} //using ... no need to define length of array
	subjects = append(subjects, "social")                     //to use this subjects should be slice
	Changelocal([5]string(subjects))                          //imp line
	fmt.Println(subjects)
	fmt.Println("lenght of subjects: ", len(subjects))

	//looping from 0 to the length of the array
	a := [...]float64{67.7, 89.8, 21, 78}
	for i := 0; i < len(a); i++ {
		fmt.Printf("%d th element of a is %.2f\n", i, a[i])
	}

	sum2 := a[0] //doubt
	sum := float64(0)
	for i, v := range a { //range returns both the index and value
		fmt.Printf("%d the element of a is %.2f\n", i, v)
		// fmt.Println(sum)
		sum += v //imp line
		sum2 += v
		// fmt.Println(sum)
	}
	fmt.Println("\nsum of all elements of a", sum)
	fmt.Println("\nsum of all elements of a", sum2)

}
