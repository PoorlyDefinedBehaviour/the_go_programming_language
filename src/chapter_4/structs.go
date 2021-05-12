package main

import "time"

// a struct of type S can't have a field of type S,
// but it can have a field of type *S
type Employee struct {
	ID        int
	Name      string
	Address   string
	DoB       time.Time
	Position  string
	Salary    int
	ManagerID int
}

func main() {
	var dilbert Employee

	var employeeOfTheMonth *Employee = &dilbert
	// this is equivalent to (*employeeOfTheMonth).Position += " (proactive team player)"
	employeeOfTheMonth.Position += " (proactive team player)"

}
