package main

import (
	"fmt"
)

func main() {
	// := in functions

	age := 30

	fmt.Println("Age:", age)

	age, name := 35, "John"

	fmt.Println("Age:", age)
	fmt.Println("Name:", name)

	// constants

	const pi = 3.14

	fmt.Println("Pi:", pi)

	const pi2 float64 = 3.14159

	fmt.Println("Pi2:", pi2)

	const (
		e  float64 = 2.71828
		g float64 = 9.81
		s  = "Hello"
		b1 = true
		b2 = false
	)

	fmt.Println("e:", e)
	fmt.Println("g:", g)
	fmt.Println("s:", s)
	fmt.Println("b1:", b1)
	fmt.Println("b2:", b2)

	const age3, name2 = 40, "Alice"

	fmt.Println("Age3:", age3)
	fmt.Println("Name3:", name2)

	const name10, name11 string = "Eve", "Frank"

	fmt.Println("Name10:", name10)
	fmt.Println("Name11:", name11)

	// var

	var age4 int = 50

	fmt.Println("Age4:", age4)

	var age5 = 60

	fmt.Println("Age5:", age5)

	var active bool

	fmt.Println("Active:", active)

	// var multiple in 1 line

	var age6, age7 int = 70, 80

	fmt.Println("Age6:", age6)
	fmt.Println("Age7:", age7)

	// type interfere

	var age8, name3 = 90, "Bob"

	fmt.Println("Age8:", age8)
	fmt.Println("Name3:", name3)

	// var grouped

	var (
		age9  = 100
		name4 = "Charlie"
		name5 string = "David" // full type declaration
		// zero
		name6 string
	)

	fmt.Println("Age9:", age9)
	fmt.Println("Name4:", name4)
	fmt.Println("Name5:", name5)
	fmt.Println("Name6:", name6)
}
