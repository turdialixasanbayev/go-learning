package main

import (
	"fmt"
	// "strconv" // to str
	// "time" // time
)

func main() {
	var name string = "Turdiali"

	fmt.Printf("Salom %s\n", name)

	// For loop

	// for i := 0; i < 8; i++ { // qavs yoq
	// 	fmt.Println("i:", i, "Salom dunyo")
	// }

	// Infinite loop

	// for {
	// 	fmt.Println("Alhamdulillah " + strconv.Itoa(time.Now().Second())) // time ni olish u int bolgani uchun str ga convert qilinadu
	// 	time.Sleep(1 * time.Second)
	// }

	// break

	// for i := 1; i <= 10; i++ {
	// 	if i == 5 {
	// 		break
	// 	}

	// 	fmt.Println(i)
	// }

	// continue

	// for i := 1; i <= 5; i++ {
	// 	if i == 3 {
	// 		continue
	// 	}

	// 	fmt.Println(i)
	// }

	// juft sonlarni chiqarish

	// for i := 1; i <= 10; i++ {
	// 	if i % 2 != 0 {
	// 		continue
	// 	}

	// 	fmt.Println(i)
	// }

	// search

	// for i := 1; i <= 100; i++ {
	// 	if i == 42 {
	// 		fmt.Println("Topildi:", i)
	// 		break
	// 	}
	// }

	// for dan while kabi foydalanish

	// i := 0

	// for i < 10 {
	// 	fmt.Println("Salom " + strconv.Itoa(i))
	// 	i++
	// }

	// array

	// myarr := [3]string{"olma", "anor", "nok"}

	// for index, value := range(myarr) {
	// 	fmt.Println("index:", index, "value:", value)
	// }

	// map

	mymap := map[int]string{
		1: "Python",
		2: "Go",
		3: "Django",
	}

	for key, value := range(mymap) {
		fmt.Println("key:", key, "value:", value)
	}
}
