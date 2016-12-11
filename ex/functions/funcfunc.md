{.exercise data-difficulty="1"}
### Functions that return functions

1. Write a function that returns a function that performs a $$+2$$ on integers. Name the function `plusTwo`.
    You should then be able do the following:

    ~~~go
    p := plusTwo()
    fmt.Printf("%v\n", p(2))
    ~~~

    Which should print 4. See (#callbacks).

2. Generalize the function from above and create a `plusX(x)` which returns functions that add `x` to an integer.


{.answer}
### Answer
1. Define a new function that returns a function: `return func(x int) int { return x + 2 }`
Function literals at work, we define the +2--function right there in the return statement.

	{callout="//"}
	~~~go
	func main() {
	   p2 := plusTwo()
	   fmt.Printf("%v\n",p2(2))
	}

	func plusTwo() func(int) int { //<1>
	    return func(x int) int { return x + 2 } //<2>
	}
	~~~

2. Here we use a closure:

	{callout="//"}
	~~~go
	func plusX(x int) func(int) int { //<1>
	   return func(y int) int { return x + y } //<2>
	}
	~~~

	Here <1>, we again define a function that returns a function.
	We use the *local* variable `x` in the function literal at <2>.
