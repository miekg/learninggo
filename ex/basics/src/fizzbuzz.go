package main

import "fmt"

func main() {
	const (
		FIZZ = 3 //<1>
		BUZZ = 5
	)
	var p bool                 //<2>
	for i := 1; i < 100; i++ { //<3>
		p = false
		if i%FIZZ == 0 { //<4>
			fmt.Printf("Fizz")
			p = true
		}
		if i%BUZZ == 0 { //<5>
			fmt.Printf("Buzz")
			p = true
		}
		if !p { //<6>
			fmt.Printf("%v", i)
		}
		fmt.Println()
	}
}
