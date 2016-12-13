{.exercise data-difficulty="1"}
### Average

1. Write code to calculate the average of a `float64` slice. In
a later exercise you will make it into a function.


{.answer}
### Answer

1. The following code calculates the average.

{callout="//"}
~~~go
sum := 0.0
switch len(xs) {
case 0: //<1>
    avg = 0
default: //<2>
    for _, v := range xs {
        sum += v
    }
    avg = sum / float64(len(xs)) //<3>
}
~~~

Here at <1> we check if the length is zero and if so, we return 0.
Otherwise we calculate the average at <2>.
We have to convert the value return from `len` to a `float64`
to make the division work at <3>.
