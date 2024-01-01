package main

import (
	"fmt"

	"tutorial/greeting"
)

func main() {
	message := greeting.Hello("Raquib")
	fmt.Println("Hello World")
	fmt.Println(message)
}
