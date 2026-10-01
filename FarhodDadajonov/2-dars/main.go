// Variables


package main

import (
	"fmt"
)

func main() {
	fmt.Println("Hello, Go")
	
	// var x int = 3
	// var y int = 4

	// var x, y int = 3, 4 // short declaration

	// x := 3
	// y := 4

	x, y := 3, 4 // short declaration

	// var summ int = x + y

	summ := x + y

	fmt.Println("Value of x:", x)
	fmt.Println("Value of y:", y)

	fmt.Println("Sum of x and y:", summ)
}
