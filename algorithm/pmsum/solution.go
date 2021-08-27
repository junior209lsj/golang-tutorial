package main

import "fmt"

func solution(absolutes []int, signs []bool) int {
	sum := 0
	ind := len(absolutes)
	for i := 0; i < ind; i++ {
		if signs[i] == true {
			sum += absolutes[i]
		} else {
			sum += -1 * absolutes[i]
		}
	}
	return sum
}

func main() {
	var in1 = []int{1, 2, 3}
	var in2 = []bool{false, false, true}
	out := solution(in1, in2)
	fmt.Println(out)
}
