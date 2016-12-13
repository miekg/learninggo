{.exercise data-dificulty="2"}
 ### Interfaces and max()

In the maximum exercise we created a max function that works on a slice of
integers.  The question now is to create a program that shows the maximum number
and that works for both integers and floats.  Try to make your program as
generic as possible, although that is quite difficult in this case.

### Answer

The following program calculates a maximum. It is as generic as you can get with
Go.

{callout="//"}
~~~go
package main

import "fmt"

func Less(l, r interface{}) bool { //<1>
    switch l.(type) {
    case int:
        if _, ok := r.(int); ok {
            return l.(int) < r.(int) //<2>
        }
    case float32:
        if _, ok := r.(float32); ok {
            return l.(float32) < r.(float32) //<3>
        }
    }
    return false
}

func main() {
    var a, b, c int = 5, 15, 0
    var x, y, z float32 = 5.4, 29.3, 0.0

    if c = a; Less(a, b) { //<4>
        c = b
    }
    if z = x; Less(x, y) { //<4>
        z = y
    }
    fmt.Println(c, z)
}
~~~

We could have chosen to make the return type of this <1> function an
`interface{}`, but that would mean that a caller would always have to do a type
assertion to extract the actual type from the interface. At <2> we compare the
parameters. All parameters are confirmed to be integers, so this is legit. And
at <3> we do the some for floats. At <4> we get the maximum value for `a`, `b`
and `x` and `y`.
