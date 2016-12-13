{.epigraph}
> * Parallelism is about performance.
> * Concurrency is about program design.
Quote: Google I/O 2010 -- Rob Pike

In this chapter we will show off Go's ability for concurrent programming using
channels and goroutines. Goroutines are the central entity in Go's ability for
concurrency.

But what *is* a goroutine, from [@effective_go]:

> They're called goroutines because the existing terms -- threads, coroutines,
> processes, and so on -- convey inaccurate connotations. A goroutine has a simple
> model: *it is a function executing in parallel with other goroutines in the same
> address space*. It is lightweight, costing little more than the allocation of
> stack space. And the stacks start small, so they are cheap, and grow by
> allocating (and freeing) heap storage as required.

A goroutine (((goroutine))) is a normal function, except that you start
it with the keyword `go`. (((keywords, go)))

~~~go
ready("Tea", 2)	    // Normal function call.
go ready("Tea", 2)  // ... as goroutine.
~~~

{callout="//"}
<{{src/channels/sleep.go}}[8,18]
Figure: Go routines in action.

The following idea for a program was taken from [@go_course_day3]. We run
a function as two goroutines, the goroutines wait for an amount of time and then
print something to the screen. At <1> we start the goroutines. The `main`
function waits long enough at <2>, so that both goroutines will have printed
their text. Right now we wait for 5 seconds, but in fact we have no idea how
long we should wait until all goroutines have exited. This outputs:

~~~go
I'm waiting         // Right away
Coffee is ready!    // After 1 second
Tea is ready!       // After 2 seconds
~~~

If we did not wait for the goroutines (i.e. remove the last line at <2>) the
program would be terminated immediately and any running goroutines would
*die with it*.

To fix this we need some kind of mechanism which allows us to
communicate with the goroutines. This mechanism is available to us in the form
of channels (((channels))). A channel can be compared to a two-way pipe in Unix
shells: you can send to and receive values from it. Those values can only be of
a specific type: the type of the channel. If we define a channel, we must also
define the type of the values we can send on the channel. Note that we must use
`make` to create a channel:

~~~go
ci := make(chan int)
cs := make(chan string)
cf := make(chan interface{})
~~~

Makes `ci` a channel on which we can send and receive integers,
makes `cs` a channel for strings and `cf` a channel for types
that satisfy the empty interface.
Sending on a channel and receiving from it, is done with the same operator:
`<-`. (((operators, channel)))

Depending on the operands it figures out what to do:

~~~go
ci <- 1   // *Send* the integer 1 to the channel ci.
<-ci      // *Receive* an integer from the channel ci.
i := <-ci // *Receive* from the channel ci and store it in i.
~~~

Let's put this to use.

{callout="//"}
~~~go
var c chan int //<1>

func ready(w string, sec int) {
    time.Sleep(time.Duration(sec) * time.Second)
    fmt.Println(w, "is ready!")
    c <- 1	//<2>
}

func main() {
    c = make(chan int) //<3>
    go ready("Tea", 2) //<4>
    go ready("Coffee", 1) //<4>
    fmt.Println("I'm waiting, but not too long")
    <-c //<5>
    <-c //<5>
}
~~~

At <1> we declare `c` to be a variable that is a channel of ints. That is: this
channel can move integers. Note that this variable is global so that the
goroutines have access to it. At <2> in the `ready` function we send the integer
1 on the channel. In our `main` function we initialize `c` at <3> and start our
goroutines <4>. At <5> we Wait until we receive a value from the channel, the
value we receive is discarded. We have started two goroutines, so we expect two
values to receive.

There is still some remaining ugliness; we have to read twice from the channel
<5>). This is OK in this case, but what if we don't know how many goroutines we
started? This is where another Go built-in comes in: `select` (((keywords,
select))). With `select` you can (among other things) listen for incoming data
on a channel.

Using `select` in our program does not really make it shorter, because we run
too few go-routines. We remove last lines and replace them with the following:

~~~go
L: for {
    select {
    case <-c:
        i++
        if i > 1 {
            break L
        }
    }
}
~~~

We will now wait as long as it takes. Only when we have received more than one
reply on the channel `c` will we exit the loop `L`.


## Make it run in parallel

While our goroutines were running concurrently, they were not running in
parallel. When you do not tell Go anything there can only be one goroutine
running at a time. With `runtime.GOMAXPROCS(n)` you can set the number of
goroutines that can run in parallel. From the documentation:

> GOMAXPROCS sets the maximum number of CPUs that can be executing
> simultaneously and returns the previous setting. If n < 1, it does not
> change the current setting. *This call will go away when the scheduler
> improves.*

If you do not want to change any source code you can also set an environment
variable `GOMAXPROCS` to the desired value.

Note that the above discussion relates to older versions of Go. From
version 1.5 and above, `GOMAXPROCS` defaults to the number of CPU
cores[@go_1_5_release_notes].

## More on channels

When you create a channel in Go with `ch := make(chan bool)`, an unbuffered
channel (((channel, unbuffered))) for bools is created. What does this mean for
your program? For one, if you read (`value := <-ch`) it will block until there
is data to receive. Secondly anything sending (`ch <- true`) will block until there
is somebody to read it. Unbuffered channels make a perfect tool for
synchronizing multiple goroutines. (((channel, blocking read))) (((channel,
blocking write)))

But Go allows you to specify the buffer size of a channel, which is quite simply
how many elements a channel can hold. `ch := make(chan bool, 4)`, creates
a buffered channel of bools that can hold 4 elements. The first 4 elements in
this channel are written without any blocking. When you write the 5^th^ element,
your code *will* block, until another goroutine reads some elements from the
channel to make room. (((channel, non-blocking read))) (((channel, non-blocking
write)))

In conclusion, the following is true in Go:

$$
\textsf{ch := make(chan type, value)}
\left\{
\begin{array}{ll}
value == 0 & \rightarrow \textsf{unbuffered} \\
value >  0 & \rightarrow \textsf{buffer }{} value{} \textsf{ elements}
\end{array}
\right.
$$

When a channel is closed the reading side needs to know this. The following code
will check if a channel is closed.

~~~go
x, ok = <-ch
~~~

Where `ok` is set to `true` the channel is not closed
*and* we've read something. Otherwise `ok` is set to `false`. In that case the
channel was closed and the value received is a zero value of the
channel's type.


## Exercises

{{ex/channels/ex.md}}
