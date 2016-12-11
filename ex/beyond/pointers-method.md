{.exercise data-difficulty="2"}
### Method calls
1. Suppose we have the following
program. Note the package `container/vector` was once part
of Go, but was removed when the `append` built-in was introduced.
However, for this question this isn't important. The package implemented
a stack-like structure, with push and pop methods.

    ~~~go
    package main

    import "container/vector"

    func main() {
        k1 := vector.IntVector{}
        k2 := &vector.IntVector{}
        k3 := new(vector.IntVector)
        k1.Push(2)
        k2.Push(3)
        k3.Push(4)
    }
    ~~~

    What are the types of `k1`, `k2` and `k3`?

2. Now, this program compiles and runs OK. All the `Push`
operations work even though the variables are of a different type. The
documentation for `Push` says:

    > `func (p *IntVector) Push(x int)`
    > Push appends x to the end of the vector.

    So the receiver has to be of type `*IntVector`, why does the code
    above (the Push statements) work correctly then?


{.answer}
### Answer
1. The type of `k1` is `vector.IntVector`. Why? We use
a composite literal (the `{}`), so we get a value of that type
back. The variable `k2` is of `*vector.IntVector`, because we
take the address (`&`) of the composite literal. And finally
`k3` has also the type `*vector.IntVector`, because `new`
returns a pointer to the type.

2. The answer is given in [@go_spec] in the section "Calls",
where among other things it says:

> A method call `x.m()` is valid if the method set of (the type of)
> `x`
> contains `m` and the argument list can be assigned to the parameter list
> of `m`. If `x` is addressable and `&x`'s method set
> contains `m`, `x.m()` is shorthand for `(&x).m()`.

In other words because `k1` is addressable and
`*vector.IntVector` *does* have the `Push` method, the
call `k1.Push(2)` is translated by Go into
`(&k1).Push(2)` which makes the type system happy again (and
you too -- now you know this).^[Also see (#methods) in this chapter.]
