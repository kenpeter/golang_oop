package main

import "employee"


// main() {, not new line
func main() {

	// call construct straight
	e := employee.Employee {
		FirstName: "Ken",
		LastName: "Peter",
		TotalLeave: 30,
		UseLeave: 20,
	}

	e.LeaveLeft()
}

// no need to call main()
