{.exersice data-difficulty="0"}
### For-loop

1. Create a loop with the `for` construct. Make it loop
   10 times and print out the loop counter with the `fmt` package.

2. Rewrite the loop from 1 to use `goto`. The keyword `for` may not be used.

3.  Rewrite the loop again so that it fills an array and then prints that array to the screen.


{.answer}
### Answer

1. There are many possibilities. One solutions could be:
<{{ex/basics/src/for.go}}
    Let's compile this and look at the output.

        % go build for.go
        % ./for
        0
        1
        .
        .
        .
        9

2. Rewriting the loop results in code that should look something
    like this (only showing the `main`-function):

    {callout="yes"}
    ~~~go
    func main() {
        i := 0	<1>
    Loop:		    <2>
        if i < 10 {
            fmt.Printf("%d\n", i)
            i++
            goto Loop <3>
        }
    }
    ~~~

    At <1> we define our loop variable. And at <2> we define a label and at <3> we jump
   to this label.

3. The following is one possible solution:
    {callout="//"}
    <{{ex/basics/src/for-arr.go}}

    Here <1> we create an array with 10 elements.
    Which we then fill <2> one by one. And finally we print it <3> with `%v` which lets
    Go to print the value for us. You could even do this in one fell swoop by using a composite literal:

~~~go
fmt.Printf("%v\n", [...]int{0,1,2,3,4,5,6,7,8,9})
~~~
