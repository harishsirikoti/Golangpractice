package main

import "fmt"

type employee struct {
	name   string
	Salary int
}

func main() {
	emp1 := employee{"Harish", 20000}
	emp2 := employee{"Ujji", 30000}
	emp3 := employee{"Chiti", 40000}
	employeinfo := make(map[int]employee)
	employeinfo[01] = emp1
	employeinfo[02] = emp2
	employeinfo[03] = emp3
	for id, info := range employeinfo {
		fmt.Printf("EmpID: ZLRSE-%v,EmpName: %v,EmpSalary: %v \n", id, info.name, info.Salary)
	}
}
