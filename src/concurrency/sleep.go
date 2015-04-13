package main

import (
	"fmt"
	"time"
)

func ready(w string, sec int) {
	time.Sleep(time.Duration(sec) * time.Second)
	fmt.Println(w, "is ready!")
}

func main() {
	go ready("Tea", 2) //<1>
	go ready("Coffee", 1) //<1>
	fmt.Println("I'm waiting")
	time.Sleep(5 * time.Second) //<2>
}
