package main

import "fmt"

type Employee struct {
	FirstName, LastName string
	Salary              float64
}

// Create the GetSalary() method below for the 'Employee' type that returns the 'Salary' field.
func (e Employee) GetSalary() float64 {
	return e.Salary
}

func main() {
	// Do not change the code below!
	var e Employee
	fmt.Scanln(&e.FirstName, &e.LastName, &e.Salary)

	// call the GetSalary() method on the '?' sign!
	fmt.Printf("%s %s earns: 💲 %.2f USD\n", e.FirstName, e.LastName, e.GetSalary())
}
