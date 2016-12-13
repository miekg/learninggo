{.exercise data-difficulty="1"}
### Stack
1. Create a simple stack which can hold a
fixed number of ints. It does not have to grow beyond this limit.
Define `push` -- put something on the stack -- and `pop`
-- retrieve something from the stack -- functions. The stack should be
a LIFO (last in, first out) stack.

![A stack](fig/stack.png "A stack.")

2. Write a `String` method which
converts the stack to a string representation. 
The stack in the figure could be represented as: `[0:m] [1:l] [2:k]` .


{.answer}
### Answer

1.  First we define a new type that represents a stack; we need an
 array (to hold the keys) and an index, which points to the last element.
 Our small stack can only hold 10 elements.

    ~~~go
    type stack struct {
        i    int
        data [10]int
    }
    ~~~

Next we need the `push` and `pop` functions to actually
use the thing. First we show the *wrong* solution!

In Go, data passed to functions is *passed-by-value* meaning a copy
is created and given to the function. The first stab for the function
`push` could be:

~~~go
func (s stack) push(k int) {
    if s.i+1 > 9 {
            return
    }
    s.data[s.i] = k
    s.i++
}
~~~

The function works on the `s` which is of the type `stack`. To
use this we just call `s.push(50)`, to push the integer 50 on
the stack. But the push function gets a copy of `s`, so it is
*not* working on the *real* thing. Nothing gets pushed to our
stack. For example the following code:

~~~go
var s stack
s.push(25)
fmt.Printf("stack %v\n", s);
s.push(14)
fmt.Printf("stack %v\n", s);
~~~

prints:

    stack [0:0]
    stack [0:0]

To solve this we need to give the function `push` a pointer
to the stack. This means we need to change `push` from

~~~go
func (s stack) push(k int)
~~~

to

~~~go
func (s *stack) push(k int).
~~~

We should now use `new()` (see (#allocation-with-new)).
in (#beyond-the-basics) to create a *pointer* to a newly
allocated `stack`, so line 1 from the example above needs to be
`s := new(stack)` .

And our two functions become:

~~~go
func (s *stack) push(k int) {
    s.data[s.i] = k
    s.i++
}

func (s *stack) pop() int {
    s.i--
    ret := s.data[s.i]
    s.data[s.i] = 0
    return ret
}
~~~

Which we then use as follows:

~~~go
func main() {
    var s stack
    s.push(25)
    s.push(14)
    fmt.Printf("stack %v\n", s)
}
~~~

2. `fmt.Printf("%v")` can
print any value (`%v`) that satisfies the `Stringer` interface
(see (#interfaces)).
For this to work we only need to define a `String()` function for
our type:

    ~~~go
    func (s stack) String() string {
        var str string
        for i := 0; i <= s.i; i++ {
            str = str + "[" +
                strconv.Itoa(i) + ":" + strconv.Itoa(s.data[i]) + "]"
        }
        return str
    }
    ~~~
