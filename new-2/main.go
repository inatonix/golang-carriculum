package main

import (
	"fmt"

	"github.com/oklog/ulid/v2"
)

func main() {
	// fmt.Println("hello, world")
	id := ulid.Make()
	fmt.Println(id)
}
