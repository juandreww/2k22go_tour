package main

import (
	"fmt"
)

func main() {
	//defer LIFO
	defer fmt.Println("Rupert")
	fmt.Println("Grint")
	defer fmt.Println("Ron")
}
