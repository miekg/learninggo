package main

import "fmt"

func fibonacci(value int) []int {
	x := make([]int, value) //<1>
	x[0], x[1] = 1, 1 //<2>
	for n := 2; n < value; n++ {
		x[n] = x[n-1] + x[n-2] //<3>
	}
	return x //<3>
}

func main() {
	for _, term := range fibonacci(10) { //<5>
		fmt.Printf("%v ", term)
	}
}

At <1> we create an \key{array} to hold the integers up to the value given in
the function call.  At <2> we start the Fibonacci calculation. Then <3>: $$x_n
= x_{n-1} + x_{n-2}$$.  At <4> \citem{} we return the \emph{entire} array.}|
And at <5> we use the `range` keyword to  "walk" the numbers returned by the
Fibonacci function. Here up to 10. Finally, we print the numbers.
