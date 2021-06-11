package main

import (
	"fmt"
	"github.com/tsangste/my-go-playground/learningarrays"
	"github.com/tsangste/my-go-playground/simpleinterest"

	"github.com/tsangste/my-go-playground/employee"

	"rsc.io/quote"
)

func calculateBill(price, no int) int {
	var totalPrice = price * no
	return totalPrice
}

func rectProps(length, width float64) (float64, float64) {
	var area = length * width
	var perimeter = (length + width) * 2
	return area, perimeter
}

func main() {
	// hello, world
	fmt.Println("Hello, World!")

	// calling package function
	fmt.Println(quote.Go())

	// function call
	price, no := 90, 6
	totalPrice := calculateBill(price, no)
	fmt.Println("Total price is", totalPrice)

	// function call with multiple results
	area, perimeter := rectProps(10.8, 5.6)
	fmt.Printf("Area %f Perimeter %f", area, perimeter)

	// calling module
	fmt.Println("Simple interest calculation")
	p := 5000.0
	r := 10.0
	t := 1.0
	si := simpleinterest.Calculate(p, r, t)
	fmt.Println("Simple interest is", si)

	// arrays module
	learningarrays.PrintArrayA()
	learningarrays.PrintArrayB()
	learningarrays.PrintArrayC()
	learningarrays.PrintSpliceArray()

	e := employee.Employee{
		FirstName:   "Sam",
		LastName:    "Adolf",
		TotalLeaves: 30,
		LeavesTaken: 20,
	}
	e.LeavesRemaining()
}
