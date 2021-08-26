package main

import (
	"fmt"
)

func main() {
	var in int
	in2 := ""
	fmt.Scanln(&in, &in2)

	switch in {
	case 0:
		fmt.Println("Zero")
		if in2 == "a" {
			break
		}
		fallthrough
	case 1:
		fmt.Println("One")
		fallthrough
	default:
		fmt.Println("Others")
	}
}
