package main

import (
	"fmt" 
	"errors"
	"math"
)


func main() {
	// res := summ(1, 2)
	// fmt.Println("Res:", res)

	res, err := sqrt(81)

	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(res)
}

func summ(x, y int) int {
	return x + y
}

func sqrt(x float64) (float64, error) {
	if x < 0 {
		return 0, errors.New("error message")
	}
	return math.Sqrt(x), nil
}
