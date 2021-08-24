package main

import "fmt"

func solution(a []int, b []int) int {
	ret := 0
	for i := 0; i < len(a) && i < len(b); i++ {
		ret += a[i] * b[i]
	}
	return ret
}

func main() {
	in1 := []int{1, 2, 3, 4, 5}
	in2 := []int{2, 2, 2, 2, 2}

	out := solution(in1, in2)

	fmt.Println(out)
}
