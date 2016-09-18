package main

import "fmt"

func fibonacci(value int) []int {
	x := make([]int, value) //<1>
	x[0], x[1] = 1, 1       //<2>
	for n := 2; n < value; n++ {
		x[n] = x[n-1] + x[n-2] //<3>
	}
	return x //<4>
}

func main() {
	for _, term := range fibonacci(10) { //<5>
		fmt.Printf("%v ", term)
	}
}
