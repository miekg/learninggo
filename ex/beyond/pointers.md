{.exercise data-difficulty="1"}
### Pointers

1.  Suppose we have defined the following structure:

    ~~~go
    type Person struct {
        name string
        age	 int
    }
    ~~~

    What is the difference between the following two lines?

    ~~~go
    var p1 Person
    p2 := new(Person)
    ~~~

2.  What is the difference between the following two allocations?

    ~~~go
    func Set(t *T) {
        x = t
    }
    ~~~

    and

    ~~~go
    func Set(t T) {
        x= &t
    }
    ~~~


{.answer}
### Answer
1. The expression, `var p1 Person` allocates a `Person`-*value* to `p1`. The type of `p1` is `Person`.
The second line: `p2 := new(Person)` allocates memory and assigns a *pointer* to `p2`. The type of `p2` is
`*Person`.

2. In the first function, `x` points to the same thing that `t` does, which is the same thing that the
actual argument points to. So in the second function, we have an "extra" variable containing a copy of the
interesting value. In the second function, `x` points to a new (heap-allocated) variable `t` which contains
a copy of whatever the actual argument value is.
