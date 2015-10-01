{.exercise data-difficulty="2"}
### Fibonacci II

This is the same exercise as an earlier one (#fibonacci) in
exercise. For completeness the complete question:

> The Fibonacci sequence starts as follows: $$1, 1, 2, 3, 5, 8, 13, \ldots$$
> Or in mathematical terms: $$ x_1 = 1; x_2 = 1; x_n = x_{n-1} +
> x_{n-2}\quad\forall n > 2 $$.
>
> Write a function that takes an `int` value and gives
> that many terms of the Fibonacci sequence.

*But* now the twist: You must use channels.

### Answer

The following program calculates the Fibonacci numbers using channels.

<{{ex/channels/src/fib.go}}
