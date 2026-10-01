// struct

package main

import "fmt"

type Car struct {
	Make, Model, Color  string
	Year, Weight        int
	Engine              engine
}

type engine struct {
	Name   string
	Hp     int
}

func main() {
	// var myCar Car

	myCar := Car{
		Make: "Volve",
		Model: "XS87",
		Color: "White",
		Year: 2017,
		Weight: 2300,
		Engine: engine{Name: "T8", Hp: 100},
	}

	fmt.Println(myCar)
}
