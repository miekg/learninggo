{.exercise data-difficulty="0"}
### Average
1. Write a function that calculates the average of a `float64` slice.


{.answer}
### Answer
1. The following function calculates the average:

 {callout="//"}
 <{{ex/functions/src/ave.go}}

 At <1> we use a named return parameter.  If the length of `xs` is zero <2>, we
 return 0.  Otherwise <3>, we calculate the average.  At <4>  we convert the
 value to a `float64` to make the division work as `len` returns an `int`.
 Finally, at <5> we reutrn our avarage.
