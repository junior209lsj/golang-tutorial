package main

import "fmt"

func writeStar(in1 int, in2 int) {
	for i := 0; i < in2; i++ {
		for j := 0; j < in1; j++ {
			fmt.Printf("*")
		}
		fmt.Printf("\n")
	}
	return
}

func main() {
	var in1, in2 int
	_, _ = fmt.Scanln(&in1, &in2)
	writeStar(in1, in2)
}
