package main

import (
	"fmt"
	"runtime"
)

func main() {
	var num int = 10

	// fmt.Println(num)

	fmt.Println("numのアドレス：", &num)
	fmt.Println("numの値：", num)

	runtime.GC()egnum

	var ptr *int = &num

	fmt.Println("ptrが指すアドレス：", ptr)
	fmt.Println("ptrが指す値：", *ptr)
	fmt.Println("ptr自体のアドレス", &ptr)

	*ptr = 20

	fmt.Println("numの新しい値：", num)
}
