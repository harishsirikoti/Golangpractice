package main

import "fmt"

type Salary interface {
	SalaryCal()
}
type Leaves interface {
	LeavesRemain() int
}
type EmplyeeDetails interface {
	Salary
	Leaves
}
type Employee struct {
	name        string
	pf          int
	salary      int
	totalLeaves int
	LeavesTaken int
}

func (s Employee) SalaryCal() {
	fmt.Printf("Name %s Salary %d\n", s.name, s.pf+s.salary)
}
func (l Employee) LeavesRemain() int {
	fmt.Printf("Leaves remaining %d \n", l.totalLeaves-l.LeavesTaken)
	return l.totalLeaves - l.LeavesTaken
}
func main() {
	Emp1 := Employee{"Harish", 2000, 50000, 10, 3}
	// var e Salary = Emp1
	var e EmplyeeDetails = Emp1
	e.SalaryCal()
	// var l Leaves = Emp1
	// fmt.Println("Leaves:", l.LeavesRemain())
	fmt.Println("Leaves:", e.LeavesRemain())
}
