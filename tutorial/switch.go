package main

import (
	"fmt"
)

func main() {
	var in int
	fmt.Scanln(&in)

	switch in {
	case 0:
		fmt.Println("Zero")
		fallthrough
	case 1:
		fmt.Println("One")
		fallthrough
	default:
		fmt.Println("Others")
	}
}
