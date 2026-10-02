// Channel lar

package main

import (
	"fmt"
	// "math/rand"
	"time"
)

// 1-misol

// func main () {
// 	// channel := make(chan string, 2) // bufford

// 	// channel <- "anor"
// 	// channel <- "olma"

// 	// fmt.Println(<-channel)
// 	// fmt.Println(<-channel)

// 	channel := make(chan string) // unbufford

// 	go func() {
// 		channel <- "anor"
// 		channel <- "olma"
// 	}()

// 	fmt.Println(<-channel)
// 	fmt.Println(<-channel)
// }

// 2-misol

// func main() {
// 	channel := make(chan int)
// 	go getRandomNumber(channel)
// 	for randomNumber := range channel {
// 		fmt.Println("tasodifiy son: ", randomNumber)
// 	}

// }

// func getRandomNumber(channel chan int) {
// 	rand.Seed(time.Now().UnixNano())
// 	for i := 1; i <= 3; i++ {
// 		number := rand.Intn(1000)
// 		time.Sleep(time.Second * 1)
// 		channel <- number
// 	}
// 	close(channel)
// }

// 3-misol

// func main() {
// 	channel1 := make(chan string)
// 	channel2 := make(chan string)

// 	go func() {
// 		for {
// 			channel1 <- "Tez"
// 			time.Sleep(time.Millisecond * 100)
// 		}
// 	}()

// 	go func() {
// 		for {
// 			channel2 <- "Sekin"
// 			time.Sleep(time.Second * 2)
// 		}
// 	}()

// 	for {
// 		fmt.Println(<-channel1)
// 		fmt.Println(<-channel2)
// 	}
// }

// Select statement

func main() {
	channel1 := make(chan string)
	channel2 := make(chan string)

	go func() {
		for {
			channel1 <- "Tez"
			time.Sleep(time.Millisecond * 10)
		}
	}()

	go func() {
		for {
			channel2 <- "Sekin"
			time.Sleep(time.Second * 1)
		}
	}()

	for {
		select {
		case message1 := <-channel1:
			fmt.Println(message1)

		case message2 := <-channel2:
			fmt.Println(message2)

		default:
			fmt.Println("ma'lumot yo'q")
			time.Sleep(time.Millisecond * 300)
		}
	}
}
