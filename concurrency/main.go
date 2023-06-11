package main

import (
	"fmt"
	"time"
)

func printNumbers() {
	for i := 0; i < 5; i++ {
		time.Sleep(time.Millisecond * 500)
		fmt.Println(i)
	}
}

func main() {
	go printNumbers()

	time.Sleep(time.Millisecond * 10 * 1000)
	fmt.Println("Waiting for channel...")

	// for n := range ch {
	// 	fmt.Println(n)
	// }
}
