package main

import "fmt"

func main() {
	var in1, in2, in3 int
	// _, _ = fmt.Scanf("%d %d %d", &in1, &in2, &in3)

	fmt.Printf("%d %d %d\n", in1, in2, in3)

	fmt.Scanln(&in1, &in2, &in3)

	fmt.Println(in1, in2, in3)

	return
}
