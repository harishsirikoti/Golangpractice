package main

import "fmt"

func main() {
	adderres, mess := adder(3, 4, 5, 6)
	fmt.Println(mess, adderres)

}
func adder(values ...int) (int, string) {
	total := 0
	for _, val := range values {

		total += val
	}
	return total, "Adder Results"
}
