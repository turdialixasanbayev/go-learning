package main

import (
	"fmt"
	"time"
)


// func testIf() {
// 	i := 7

// 	if i == 7 {
// 		fmt.Println("Yetti")
// 	}
// }

// func testIfElse() {
// 	points := 8

// 	if points > 10 {
// 		fmt.Println("You have more than 10 points")
// 	} else {
// 		fmt.Println("You have less than 10 points")
// 	}
// }

// func testIfElseIf() {
// 	points := 8

// 	if points > 10 {
// 		fmt.Println("You have more than 10 points")
// 	} else if points == 10 {
// 		fmt.Println("You have exactly 10 points")
// 	} else {
// 		fmt.Println("You have less than 10 points")
// 	}
// }

// func main() {
// 	testIf()
// 	testIfElse()
// 	testIfElseIf()
// }


// switch case

func testSwitchCase() {
	weekday := time.Now().Weekday()

	fmt.Println("Today is:", weekday)

	switch weekday {
	case 1:
		fmt.Println("Dushanba")
	case 2:
		fmt.Println("Seshanba")
	case 3:
		fmt.Println("Chorshanba")
	case 4:
		fmt.Println("Payshanba")
	case 5:
		fmt.Println("Juma")
	case 6:
		fmt.Println("Shanba")
	case 0:
		fmt.Println("Yakshanba")
	default:
		fmt.Println("Noma'lum kun")
	}
}


func testSwitch() {
	var n int = 2

	switch {
		case n == 1:
			fmt.Println("n is 1")
		case n == 2:
			fmt.Println("n is 2")
		default:
			fmt.Println("n is not 1 or 2")
	}
}

// Multiple case values


func testMultipleCaseValues() {
	var userChoice string = "ikki"

	switch {
	case  userChoice == "bir", userChoice == "ikki":
		fmt.Println("Python")
	case userChoice == "ikki", userChoice == "uch":
		fmt.Println("JavaScript")
	default:
		fmt.Println("Golang")
	}
}

func main() {
	testSwitchCase()
	testSwitch()
	testMultipleCaseValues()
}
