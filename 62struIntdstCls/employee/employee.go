package employee

import "fmt"

type employee struct {
	firstname, lastname      string
	totalleaves, leavestaken int
}

func New(firstname string, lastname string, totalleaves int, leavestaken int) employee {
	e := employee{firstname, lastname, totalleaves, leavestaken}
	return e
}
func (e employee) LeavesRemaining() {
	fmt.Println(e.firstname, " ", e.lastname, " has ", e.totalleaves-e.leavestaken, " Leaves\n")
}
