{.exercise data-difficulty="1"}
### Fibonacci

1.  The Fibonacci sequence starts as follows: $$1, 1, 2, 3, 5, 8, 13, \ldots$$
    Or in mathematical terms: $$ x_1 = 1; x_2 = 1; x_n = x_{n-1} + x_{n-2}\quad\forall n > 2 $$.

    Write a function that takes an `int` value and gives 
    that many terms of the Fibonacci sequence.


{.answer}
### Answer
1. The following program calculates Fibonacci numbers:

 {callout="//"}
<{{ex/functions/src/fib.go}}

At <1> we create an array to hold the integers up to the value given in
the function call.  At <2> we start the Fibonacci calculation. Then <3>:
$$x_n = x_{n-1} + x_{n-2}$$.  At <4> we return the *entire* array.
And at <5> we use the `range` keyword to  "walk" the numbers returned by the
Fibonacci function. Here up to 10. Finally, we print the numbers.
