package main

import "fmt"

type NameAge struct {
	name string // Both non exported fields.
	age  int
}

func main() {
	a := new(NameAge)
	a.name = "Pete"
	a.age = 42
	fmt.Printf("%v\n", a)
}
