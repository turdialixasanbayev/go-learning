package main

import (
	"fmt"
	// "sort"
)

func main() {

	// Arrays

	// var arrays

	// var myarr [3]string

	// myarr[0] = "Go"
	// myarr[1] = "is"
	// myarr[2] = "awesome!"

	// fmt.Println("Qatorlarning elementlari: ")

	// fmt.Println("1: ", myarr[0])
	// fmt.Println("2: ", myarr[1])
	// fmt.Println("3: ", myarr[2])

	// qiymatini berib ketish va := bilan ishlash

	// myarr := [3]int{1, 2, 3}

	// fmt.Println("Qatorlarning elementlari: ")
	// fmt.Println("1: ", myarr[0])
	// fmt.Println("2: ", myarr[1])
	// fmt.Println("3: ", myarr[2])

	// myarr := [...]string{"Go", "is", "awesome!"} // o'lchamni dinamik qilish
	// myarr2 := [3]string{"Go", "is", "awesome!"}

	// fmt.Println("Qatorlarning elementlari: ")
	// fmt.Println("1: ", myarr[0])
	// fmt.Println("2: ", myarr[1])
	// fmt.Println("3: ", myarr[2])

	// fmt.Println("4: ", myarr2[0])
	// fmt.Println("5: ", myarr2[1])
	// fmt.Println("6: ", myarr2[2])

	// // Solishtirish

	// fmt.Println(myarr == myarr2)

	//  Ko'p o'lchamli arrays

	// myarr := [2][3]string{
	// 	{"Go", "is", "awesome!"},
	// 	{"I", "love", "Go!"},
	// }

	// fmt.Println("Ko'p o'lchamli arrays:")

	// fmt.Println("1: ", myarr[0][0])
	// fmt.Println("2: ", myarr[0][1])
	// fmt.Println("3: ", myarr[0][2])

	// fmt.Println("4: ", myarr[1][0])
	// fmt.Println("5: ", myarr[1][1])
	// fmt.Println("6: ", myarr[1][2])

	// myarr := [3]string{"Go", "is", "awesome!"}

	// myarr2 := myarr // toliq nusxa

	// fmt.Println(myarr)
	// fmt.Println(myarr2)

	// myarr2[0] = "Python"

	// fmt.Println(myarr)
	// fmt.Println(myarr2)

	// myarr := [3]int{1,2,3}


	// myarr2 := &myarr
	// fmt.Println(myarr)
	// fmt.Println(*myarr2) // referal nusxa

	// myarr[2] = 100
	// fmt.Println(myarr)
	// fmt.Println(*myarr2)

	// Slices

	// myslice := []int{2, 4, 8}
	// myslice = append(myslice, 10)

	// fmt.Printf("'slice' ning uzunligi: %d", len(myslice)) // uzunligini topish

	// Array dan slice hosil qilish

	// arr := [3]int{1,2,3}

	// slice := arr[0:1]

	// fmt.Println(slice)

	// Slice ni tartiblash

	// myslice := []int{3,6,9,8}

	// sort.Ints(myslice)

	// fmt.Println(myslice)

	// Map

	status := make(map[string]int) // make

	status["pending"] = 0
	status["succes"] = 1 // add

	var pendingStatus = status["pending"]
	fmt.Println(pendingStatus) // get

	delete(status, "succes") // delete
}
