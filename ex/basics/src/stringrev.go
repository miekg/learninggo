package main

import "fmt"

func main() {
	s := "foobar"
	a := []rune(s) //<1>
	for i, j := 0, len(a)-1; i < j; i, j = i+1, j-1 {
		a[i], a[j] = a[j], a[i] //<2>
	}
	fmt.Printf("%s\n", string(a)) //<3>
}
