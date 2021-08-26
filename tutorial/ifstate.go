package main

import (
	"fmt"
	"math"
)

func square(x float64) float64 {
	return x * x
}

func main() {
	var in float64
	fmt.Scanln(&in)
	x := math.Pow(in, 2.0)
	if x < 10.0 {
		fmt.Println("Less than 10")
	}
	if v := square(in); v < 10.0 {
		fmt.Println("Less than 10")
	}
}
