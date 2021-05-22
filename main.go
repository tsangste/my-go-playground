package main

import (
	"fmt"

	"github.com/tsangste/my-go-playground/employee"

	"rsc.io/quote"
)

func main() {
	fmt.Println("Hello, World!")
	fmt.Println(quote.Go())

	e := employee.Employee {
		FirstName: "Sam",
		LastName: "Adolf",
		TotalLeaves: 30,
		LeavesTaken: 20,
	}
	e.LeavesRemaining()
}
