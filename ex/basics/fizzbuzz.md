{.exercise data-difficulty="0"}
### FizzBuzz

1. Solve this problem, called the Fizz-Buzz [@fizzbuzz_cite] problem:

Write a program that prints the numbers from 1 to 100. But for multiples
of three print, "Fizz" instead of the number, and for multiples of
five, print "Buzz". For numbers which are multiples of both three and
five, print "FizzBuzz".


{.answer}
### Answer
1. A possible solution to this problem is the following program.

{callout="//"}
<{{ex/basics/src/fizzbuzz.go}}

Here <1> we define two constants to make our code more readable, see (#constants).
At <2> we define a boolean that keeps track if we already printed something.
At <3> we start our for-loop, see (#for).
If the value is divisible by FIZZ - that is, 3 - , we print "Fizz" <4>.
And at <5> we check if the value is divisble by BUZZ -- that is, 5 -- if so print
"Buzz". Note that we have also taken care of the FizzBuzz case.
At <6>, if printed neither Fizz nor Buzz printed, we print the value.
