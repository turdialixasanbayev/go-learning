// Goroutine'lar haqida

package main

import (
	"fmt"
	"time"
)

func main() {
	go display("salom berdik") // goroutine yaratish (thread ni yangi ko'rinishi, yengilroq variant)
	go display("alik oldik") // main ichidagi barcha kodlar goroutine bolsa yani (asinxron) bolsa console ga hech narsa chiqmaydi

	// time.Sleep(2 * time.Second) // 2 sekund kutish (goroutine'lar ishlashini kutish)

	fmt.Scanln() // Enter bosilishini kutish uchun (barcha goroutine lar ishlab turaveradi)
}

func display(input string) { // infinite loop (hech qachon to'xtamaydi)
	for i := 1; true; i++ {
		fmt.Println(i, input)
		time.Sleep(1 * time.Second)
	}
}
