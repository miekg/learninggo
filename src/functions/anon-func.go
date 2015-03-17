package main

import "fmt"

func main() {
	a := func() { //<1>
		fmt.Println("Hello")
	} //<2>
	a() //<3>
}
