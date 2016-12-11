{.epigraph}
> I'm always delighted by the light touch and stillness of
> early programming languages.  Not much text; a lot gets
> done. Old programs read like quiet conversations
> between a well-spoken research worker and a well-
> studied mechanical colleague, not as a debate with a
> compiler.  Who'd have guessed sophistication bought
> such noise?
Quote: -- Richard P. Gabriel

Functions are the basic building blocks of Go programs; all interesting stuff
happens in them.

Here is an example of how you can declare a function:

{callout="//"}
~~~go
type mytype int
func (p mytype) funcname(q int) (r,s int) { return 0,0 }
// <1>        <2>        <3>      <4>        <5>         <6>
~~~

To declare a function, you use the `func` keyword <1>. You can optionally bind
<2> to a specific type called receiver (((functions, receiver))) (a function
with a receiver is usually called a method(((functions, method)))). This will
be explored in (#interfaces). Next <3> you write the name of your
function. Here <4> we define that the variable `q` of type `int` is the input
parameter. Parameters are passed *pass-by-value*.(((functions, pass-by-value)))
The variables `r` and `s` <5> are the *named return parameters* (((functions,
named return parameters))) for this function. Functions in Go can have multiple
return values. This is very useful to return a value *and* error. This
removes the need for in-band error returns (such as -1 for `EOF`) and modifying
an argument. If you want the return parameters not to be named you only give the
types: `(int, int)`. If you have only one value to return you may omit the
parentheses. If your function is a subroutine and does not have anything to
return you may omit this entirely. Finally, we have the body <6> of the
function. Note that `return` is a statement so the braces around the
parameter(s) are optional.

As said the return or result parameters of a Go function can be given names and
used as regular variables, just like the incoming parameters. When named, they
are initialized to the zero values for their types when the function begins. If
the function executes a `return` statement with no arguments, the current values
of the result parameters are returned. Using these features enables you (again)
to do more with less code.^[This is a motto of Go; "Do *more* with *less*
code".]

The names are not mandatory but they can make code shorter and clearer:
*they are documentation*. However don't overuse this feature, especially in
 longer functions where it might not be immediately apparent what is returned.

Functions can be declared in any order you wish. The compiler scans the entire
file before execution, so function prototyping is a thing of the past in Go. Go
does not allow nested functions, but you can work around this with anonymous
functions. See the Section (#functions-as-values) in this chapter. Recursive
functions work just as in other languages:

{callout="//"}
~~~go
func rec(i int) {
    if i == 10 { //<1>
        return
    }
    rec(i+1) //<2>
    fmt.Printf("%d ", i)
}
~~~

Here <2> we call the same function again, `rec` returns when `i` has the value
10, this is checked on the second line <1>. This function prints: `9
8 7 6 5 4 3 2 1 0`, when called as `rec(0)`.


## Scope

Variables declared outside any functions are *global* (((scope, local))) in Go,
those defined in functions are *local* (((scope, local))) to those functions. If
names overlap - a local variable is declared with the same name as a global one
- the local variable hides the global one when the current function is executed.

In the following example we call `g()` from `f()`:

{callout="//"}
~~~go
package main

var a int // <1>

func main() {
    a = 5
    print(a)
    f()
}

func f() {
    a := 6 // <2>
    print(a)
    g()
}

func g() {
    print(a)
}
~~~

Here <1>, we declare `a` to be a global variable of type `int`. Then in the
`main` function we give the *global* `a` the value of 5, after printing it we
call the function `f`. Then here <2>, `a := 6`, we create a *new, local*
variable also called `a`. This new `a` gets the value of 6, which we then print.
Then we call `g`, which uses the *global* `a` again and prints `a`'s value set
in `main`. Thus the output will be: `565`. A *local* variable is *only* valid
when we are executing the function in which it is defined. Note that the `:=`
used in line 12 is sometimes hard to spot so it is generally advised *not* to
use the same name for global and local variables.


## Functions as values
(((functions, as values))) (((functions, literals))) As with almost everything in
Go, functions are also *just* values. They can be assigned to variables as
follows:

{callout="//"}
<{{src/functions/anon-func.go}}[3,]

`a` is defined as an anonymous (nameless) function <1>.
Note the lack of parentheses `()` after `a`. If there were, that would be to *call*
some function with the name `a` before we have defined what `a` is. Once `a` is 
defined, then we can *call* it, <3>.

Functions--as--values may be used in other places, for example maps. Here we
convert from integers to functions:

~~~go
var xs = map[int]func() int{
    1: func() int { return 10 },
    2: func() int { return 20 },
    3: func() int { return 30 },
}
~~~

Note that the final comma on second to last line is *mandatory*.

Or you can write a function that takes a function as its parameter, for example
a `Map` function that works on `int` slices. This is left as an exercise for the
reader; see the exercise (#map-function).


## Callbacks

Because functions are values they are easy to pass to functions, from where they
can be used as callbacks. First define a function that does "something" with an
integer value:

~~~go
func printit(x int) {
    fmt.Printf("%v\n", x)
}
~~~

This function does not return a value and just prints its argument. The
*signature* (((functions, signature))) of this function is: `func printit(int)`,
or without the function name: `func(int)`. To create a new function that uses
this one as a callback we need to use this signature:

~~~go
func callback(y int, f func(int)) {
    f(y)
}
~~~

Here we create a new function that takes two parameters: `y int`, i.e. just an
`int` and `f func(int)`, i.e. a function that takes an int and returns nothing.
The parameter `f` is the variable holding that function. It can be used as any
other function, and we execute the function on line 2 with the parameter `y`:
`f(y)`


## Deferred Code

Suppose you have a function in which you open a file and perform various writes
and reads on it. In such a function there are often spots where you want to
return early. If you do that, you will need to close the file descriptor you are
working on. This often leads to the following code:

{callout="//"}
~~~go
func ReadWrite() bool {
    file.Open("file")
    // Do your thing
    if failureX {
        file.Close() //<1>
        return false
    }

    if failureY {
        file.Close() //<1>
        return false
    }
    file.Close() //<1>
    return true  //<2>
}
~~~

Note that we repeat a lot of code here; you can see the that `file.Close()` is
called at <1>. To overcome this, Go has the `defer` (((keywords, defer)))
keyword. After `defer` you specify a function which is called just *before* <2>
the current function exits.

With `defer` we can rewrite the above code as follows. It makes the function
more readable and it puts the `Close` *right next* to the `Open`.

{callout="//"}
~~~go
func ReadWrite() bool {
    file.Open("filename")
    defer file.Close() //<1>
    // Do your thing
    if failureX {
        return false //<2>
    }
    if failureY {
        return false //<2>
    }
    return true //<2>
}
~~~

At <1> `file.Close()` is added to the defer list. (((keywords, defer list)))
`Close` is now done automatically at <2>. This makes the function shorter and
more readable. It puts the `Close` right next to the `Open`.

You can put multiple functions on the "defer list", like this example from
[@effective_go]:

~~~go
for i := 0; i < 5; i++ {
    defer fmt.Printf("%d ", i)
}
~~~

Deferred functions are executed in LIFO order, so the above code prints: `4
3 2 1 0`.

With `defer` you can even change return values, provided that you are using
named result parameters and a function literal (((functions, literal)))^[A
function literal is sometimes called a closure (((closure))).], i.e:

~~~go
defer func() {/* ... */}()
~~~

Here we use a function without a name and specify the body of the function
inline, basically we're creating a nameless function on the spot. The final
braces are needed because `defer` needs a function call, not a function value.
If our anonymous function would take an parameter it would be easier to see why
we need the braces:

~~~go
defer func(x int) {/* ... */}(5)
~~~

In this (unnamed) function you can access any named return parameter:

{callout="//"}
~~~go
func f() (ret int)
    defer func() { //<1>
        ret++
    }()
    return 0
}
~~~

Here <1> we specify our function, the named return value `ret` is initialized
with zero. The nameless function in the defer increments the value of `ret` 
with 1. The `return 0` on line
5 *will not be the returned value*, because of `defer`. The function `f` will
return 1!


## Variadic Parameter
Functions that take a variable number of parameters are known as variadic
functions. (((functions, variadic))) To declare a function as variadic, do
something like this:

~~~go
func myfunc(arg ...int) {}
~~~

The `arg ...int` instructs Go to see this as a function that takes a variable
number of arguments. Note that these arguments all have to have the type `int`.
In the body of your function the variable `arg` is a slice of ints:

~~~go
for _, n := range arg {
    fmt.Printf("And the number is: %d\n", n)
}
~~~

We range over the arguments on the first line. We are not interested in the
index as returned by `range`, hence the use of the underscore there. In the body
of the `range` we just print the parameters we were given.

If you don't specify the type of the variadic argument it defaults to the empty
interface `interface{}` (see Chapter (#interfaces)).

Suppose we have another variadic function called `myfunc2`, the following
example shows how to pass variadic arguments to it:

~~~go
func myfunc(arg ...int) {
    myfunc2(arg...)
    myfunc2(arg[:2]...)
}
~~~

With `myfunc2(arg...)` we pass all the parameters to `myfunc2`, but because the
variadic parameters is just a slice, we can use some slice tricks as well.


## Panic and recovering

Go does not have an exception mechanism: you cannot throw exceptions. Instead it
uses a panic-and-recover mechanism. It is worth remembering that you should use
this as a last resort, your code will not look, or be, better if it is littered
with panics. It's a powerful tool: use it wisely. So, how do you use it? In the
words of the Go Authors [@go_blog_panic]:

Panic
:   is a built-in function that stops the ordinary flow of control and begins
    panicking. When the function `F` calls `panic`, execution of `F` stops, any
    deferred functions in `F` are executed normally, and then `F` returns to its
    caller. To the caller, `F` then behaves like a call to `panic`. The process
    continues up the stack until all functions in the current goroutine have
    returned, at which point the program crashes. Panics can be initiated by
    invoking `panic` directly. They can also be caused by *runtime errors*, such as
    out-of-bounds array accesses.

Recover
:   is a built-in function that regains control of a panicking goroutine.
    Recover is *only* useful inside *deferred* functions. During normal execution,
    a call to `recover` will return `nil` and have no other effect. If the current
    goroutine is panicking, a call to `recover` will capture the value given to
    `panic` and resume normal execution.

This function checks if the function it gets as argument will panic when it is
executed^[Modified from a presentation of Eleanor McHugh.]:

{callout="//"}
~~~go
func Panic(f func()) (b bool) { //<1>
    defer func() { //<2>
        if x := recover(); x != nil {
            b = true
        }
    }()
    f() //<3>
    return //<4>
}
~~~

We define a new function `Panic` <1> that takes a function as an argument (see
(#functions-as-values)). It returns true if `f` panics when run, else false. We
then <2> define a `defer` function that utilizes `recover`. If the current
goroutine panics, this defer function will notice that. If `recover()` returns
non-`nil` we set `b` to true. At <3> Execute the function we received as the
argument. And finally <4> we return the value of `b`. Because `b` is a named
return parameter.

The following code fragment, shows how we can use this function:

~~~go
func panicy() {
    var a []int
    a[3] = 5
}

func main() {
    fmt.Println(Panic(panicy))
}
~~~

On line 3 the `a[3] = 5` triggers a *runtime* out of bounds error which results
in a panic. Thus this program will print `true`. If we change line 2: `var
a []int` to `var a [4]int` the function `panicy` does not panic anymore. Why?

## Exercises

{{ex/functions/ex.md}}
