package main

import (
	"fmt"
	"time"
	"sync"
)

func main() {
	wg := sync.WaitGroup{} // WaitGroup to wait for goroutines to finish
	wg.Add(2) // Add 2 goroutines to the wait group
	// go display("Salom berdik", &wg) // wg is passed as a pointer to the display function
	// go display("Alik oldik", &wg) // bu usul Single responsibility ga zid (wg ni funksiyaga berish)

	go func () {
		display("Salom berdik")
		wg.Done() // Mark this goroutine as done
	}()

	go func () {
		display("Alik oldik")
		wg.Done() // Mark this goroutine as done
	}()

	// fmt.Scanln()

	wg.Wait() // Wait for all goroutines to finish

	fmt.Println("All goroutines finished")
}

// func display(input string, wg *sync.WaitGroup) { // wg - pointer to WaitGroup
// 	for i := 0; true; i++ {
// 		fmt.Println(i, input)
// 		time.Sleep(1 * time.Second)
// 	}
// 	wg.Done() // Mark this goroutine as done
// }

// func display(input string, wg *sync.WaitGroup) { // wg - pointer to WaitGroup
// 	for i := 0; i < 7; i++ { // 7 iterations instead of infinite loop
// 		fmt.Println(i, input)
// 		time.Sleep(1 * time.Second)
// 	}
// 	wg.Done() // Mark this goroutine as done
// } // not Single responsibility

func display(input string) {
	for i := 0; i < 7; i++ {
		fmt.Println(i, input)
		time.Sleep(1 * time.Second)
	}
}
