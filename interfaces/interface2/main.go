package main

import "fmt"

type SalaryCalculator interface {
	Calculate() int
}

type Permanent struct {
	empId, basicpay, pf int
}
type Contract struct {
	empId, basicpay int
}

func (p Permanent) Calculate() int {
	return p.basicpay + p.pf
}
func (c Contract) Calculate() int {
	return c.basicpay
}
func Totalexpences(t []SalaryCalculator) {
	totalamount := 0
	for _, value := range t {
		totalamount = totalamount + value.Calculate()
	}
	fmt.Println("Total expenses:", totalamount)
}
func main() {
	pemp1 := Permanent{
		empId:    1,
		basicpay: 5000,
		pf:       20,
	}
	pemp2 := Permanent{
		empId:    2,
		basicpay: 6000,
		pf:       30,
	}
	cemp1 := Contract{
		empId:    3,
		basicpay: 3000,
	}
	employee := []SalaryCalculator{pemp1, pemp2, cemp1}
	Totalexpences(employee)
}
