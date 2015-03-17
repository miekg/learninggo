package main

import "fmt"

func main() {
	var arr [10]int //<1>
	for i := 0; i < 10; i++ {
		arr[i] = i //<2>
	}
	fmt.Printf("%v", arr) //<3>
}
