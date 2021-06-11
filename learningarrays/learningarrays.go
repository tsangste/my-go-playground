package learningarrays

import "fmt"

func PrintArrayA() {
	// array declaration
	a := [3]int{12, 78, 50}
	fmt.Println(a)
}

func PrintArrayB() {
	// array declaration auto length
	b := [...]int{12, 78, 50}
	fmt.Println(b)
	fmt.Println("length of a is", len(b))
}

func PrintArrayC() {
	c := [...]float64{67.7, 89.8, 21, 78}
	sum := float64(0)
	for i, v := range c { //range returns both the index and value
		fmt.Printf("%d the element of a is %.2f\n", i, v)
		sum += v
	}
	fmt.Println("\nsum of all elements of a", sum)
}

func PrintSpliceArray() {
	a := [5]int{76, 77, 78, 79, 80}
	var b []int = a[1:4] //creates a slice from a[1] to a[3]
	fmt.Println(b)
}
