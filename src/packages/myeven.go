package main

import ( //<1>
	"even"	//<2>
	"fmt"	//<3>
)

func main() {
	i := 5
	fmt.Printf("Is %d even? %v\n", i, even.Even(i)) //<4>
}
