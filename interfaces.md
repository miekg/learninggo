{.epigraph}
> I have this phobia about having my body penetrated surgically. You know what
> I mean?
Quote: eXistenZ -- Ted Pikul

In Go, the word *interface*(((interface))) is overloaded to mean several
different things. Every type has an interface, which is the *set of methods
defined* for (((interface, set of methods))) that type. This bit of code defines
a struct type `S` with one field, and defines two methods for `S`. ^[The following text is partly from [@go_interfaces].]

~~~go
type S struct { i int }
func (p *S) Get() int  { return p.i }
func (p *S) Put(v int) { p.i = v }
~~~
Figure: Defining a struct and methods on it.

You can also define an (((interface, type)))interface type, which is simply
a set of methods. This defines an interface `I` with two methods:

~~~go
type I interface {
    Get() int
    Put(int)
}
~~~

`S` is a valid *implementation* for interface `I`, because it defines the two
methods which `I` requires. Note that this is true even though there is no
explicit declaration that `S` implements `I`.

A Go program can use this fact via yet another meaning of interface, which is an
interface value: (((interface, value)))

{callout="//"}
~~~go
func f(p I) { //<1>
    fmt.Println(p.Get()) //<2>
    p.Put(1) //<3>
}
~~~

At <1> we declare a function that takes an interface type as the argument.
Because `p` implements `I`, it *must* have the `Get()` method, which we call at
<2>. And the same holds true for the `Put()` method at <3>. Because `S`
implements `I`, we can call the function `f` passing in a pointer to a value of
type `S`: `var s S; f(&s)`

The reason we need to take the address of `s`, rather than a value of type `S`,
is because we defined the methods on `s` to operate on pointers, see the
definition in the code above. This is not a requirement -- we could have defined
the methods to take values -- but then the `Put` method would not work as
expected.

The fact that you do not need to declare whether or not a type implements an
interface means that Go implements a form of duck typing (((duck, typing)))
[@duck_typing]. This is not pure duck typing, because when possible the
Go compiler will statically check whether the type implements the interface.
However, Go does have a purely dynamic aspect, in that you can convert from one
interface type to another. In the general case, that conversion is checked at
run time. If the conversion is invalid -- if the type of the value stored in
the existing interface value does not satisfy the interface to which it is being
converted -- the program will fail with a run time error.

Interfaces in Go are similar to ideas in several other programming languages:
pure abstract virtual base classes in C++, typeclasses in Haskell or duck typing
in Python. However there is no other language which combines interface values,
static type checking, dynamic run time conversion, and no requirement for
explicitly declaring that a type satisfies an interface. The result in Go is
powerful, flexible, efficient, and easy to write.


## Which is what?

Let's define another type `R` that also implements the interface `I`:

~~~go
type R struct { i int }
func (p *R) Get() int  { return p.i }
func (p *R) Put(v int) { p.i = v }
~~~

The function `f` can now accept variables of type `R` and `S`.

Suppose you need to know the actual type in the function `f`. In Go you can
figure that out by using a type switch(((type switch))).

{callout="//"}
~~~go
func f(p I) {
    switch t := p.(type) { //<1>
        case *S: //<2>
        case *R: //<2>
        default: //<3>
    }
}
~~~

At <1> we use the type switch, note that the `.(type)` syntax is *only* valid
within a `switch` statement. We store the value in the variable `t`. The
subsequent cases <2> each check for a different *actual* type. And we can even
have a `default` <3> clause. It is worth pointing out that both `case R` and
`case s` aren't possible, because `p` needs to be a pointer in order to satisfy
`i`.

A type switch isn't the only way to discover the type at *run-time*.

{callout="//"}
~~~go
if t, ok := something.(I); ok { //<1>
    // ...
}
~~~

You can also use a "comma, ok" form <1> to see if an interface type implements
a specific interface. If `ok` is true, `t` will hold the type of `something`.
When you are sure a variable implements an interface you can use: `t := something.(I)` .


## Empty interface

Since every type satisfies the empty interface: `interface{}` we can create
a generic function which has an empty interface as its argument:

~~~go
func g(something interface{}) int {
    return something.(I).Get()
}
~~~

The `return something.(I).Get()` is the tricky bit in this function. The value
`something` has type `interface{}`, meaning no guarantee of any methods at all:
it could contain any type. The `.(I)` is a type assertion (((type assertion)))
which converts `something` to an interface of type `I`. If we have that type we
can invoke the `Get()` function. So if we create a new variable of the type
`*S`, we can just call `g()`, because `*S` also implements the empty interface.

~~~go
s = new(S)
fmt.Println(g(s));
~~~

The call to `g` will work fine and will print 0. If we however invoke `g()` with
a value that does not implement `I` we have a problem:

~~~go
var i int
fmt.Println(g(i))
~~~

This compiles, but when we run this we get slammed with: "panic: interface
conversion: int is not main.I: missing method Get".

Which is completely true, the built-in type `int` does not have a `Get()`
method.


## Methods

Methods are functions that have a receiver (see (#functions)).
You can define methods on any type (except on non-local types, this includes
built-in types: the type `int` can not have methods).
You can however make a new integer type with its own methods. For example:

~~~go
type Foo int

func (self Foo) Emit() {
    fmt.Printf("%v", self)
}

type Emitter interface {
    Emit()
}
~~~

Doing this on non-local (types defined in other packages) types yields an error
"cannot define new methods on non-local type int".


## Methods on interface types

An interface defines a set of methods. A method contains the actual code. In
other words, an interface is the definition and the methods are the
implementation. So a receiver can not be an interface type, doing so results in
a "invalid receiver type ..." compiler error. The authoritative word from the
language spec [@go_spec]:

> The receiver type must be of the form `T` or `*T` where `T` is a type name. `T`
> is called the receiver base type or just base type. The base type must not be
> a pointer or interface type and must be declared in the same package as the
> method.


A> Creating a pointer to an interface value is a useless action in Go. It is in
A> fact illegal to create a pointer to an interface value. The release notes for an
A> earlier Go release that made them illegal leave no room for doubt:
A>
A> > The language change is that uses of pointers to interface values no longer
A> > automatically de-reference the pointer.  A pointer to an interface value is
A> > more often a beginner's bug than correct code.


## Interface names

By convention, one-method interfaces are named by the method name plus the *-er*
suffix: Read*er*, Writ*er*, Formatt*er* etc.

There are a number of such names and it's productive to honor them and the
function names they capture. `Read`, `Write`, `Close`, `Flush`, `String` and so
on have canonical signatures and meanings. To avoid confusion, don't give your
method one of those names unless it has the same signature and meaning.
Conversely, if your type implements a method with the same meaning as a method
on a well-known type, give it the same name and signature; call your
string-converter method `String` not `ToString`. ^[Text copied from
[@effective_go].]


## A sorting example

Recall the Bubblesort exercise, where we sorted an array of integers:

~~~go
func bubblesort(n []int) {
    for i := 0; i < len(n)-1; i++ {
        for j := i + 1; j < len(n); j++ {
            if n[j] < n[i] {
                n[i], n[j] = n[j], n[i]
            }
        }
    }
}
~~~

A version that sorts strings is identical except for the signature of the
function: `func bubblesortString(n []string) { /* ... */ }` . Using this
approach would lead to two functions, one for each type. By using interfaces we
can make this more (((generic))) generic. Let's create a new function that will
sort both strings and integers, something along the lines of this non-working
example:

{callout="//"}
~~~go
func sort(i []interface{}) {  //<1>
    switch i.(type) {         //<2>
    case string:              //<3>
        // ...
    case int:
        // ...
    }
    return /* ... */          //<4>
}
~~~

Our function will receive a slice of empty interfaces at <1>. We then <2> use a
type switch to find out what the actual type of the input is. And then <3>
then sort accordingly. And, when done, return <4> the sorted slice.

But when we call this function with `sort([]int{1, 4, 5})`, it fails with:
"cannot use i (type []int) as type []interface { } in function argument"

This is because Go can not easily convert to a *slice* of interfaces.
Just converting to an interface is easy, but to a slice is much more costly.
The full mailing list discussion on this subject can be found at
[@go_nuts_interfaces]. To keep a long story short: Go does not (implicitly) convert slices for you.

So what is the Go way of creating such a "generic" function?
Instead of doing the type inference ourselves with a type switch, we let
Go do it implicitly:
The following steps are required:

* Define an interface type (called `Sorter` here) with a number of methods
  needed for sorting. We will at least need a function to get the length of the
  slice, a function to compare two values and a swap function.

    ~~~go
    type Sorter interface {
        Len() int           // len() as a method.
        Less(i, j int) bool // p[j] < p[i] as a method.
        Swap(i, j int)      // p[i], p[j] = p[j], p[i] as a method.
    }
    ~~~

* Define new types for the slices we want to sort. Note that we declare slice types:

    ~~~go
    type Xi []int
    type Xs []string
    ~~~

* Implementation of the methods of the `Sorter` interface.
  For integers:

    ~~~go
    func (p Xi) Len() int               {return len(p)}
    func (p Xi) Less(i int, j int) bool {return p[j] < p[i]}
    func (p Xi) Swap(i int, j int)      {p[i], p[j] = p[j], p[i]}
    ~~~

    And for strings:

    ~~~go
    func (p Xs) Len() int               {return len(p)}
    func (p Xs) Less(i int, j int) bool {return p[j] < p[i]}
    func (p Xs) Swap(i int, j int)      {p[i], p[j] = p[j], p[i]}
    ~~~

* Write a *generic* Sort function that works on the `Sorter` interface.

    {callout="//"}
    ~~~go
    func Sort(x Sorter) { //<1>
        for i := 0; i < x.Len() - 1; i++ { //<2>
            for j := i + 1; j < x.Len(); j++ {
                if x.Less(i, j) {
                    x.Swap(i, j)
                }
            }
        }
    }
    ~~~

	At <1> `x` is now of the `Sorter` type and using the defined methods for this interface we implement
	Bubblesort at <2>.

	Now we can use our *generic* `Sort` function as follows:

    ~~~go
    ints := Xi{44, 67, 3, 17, 89, 10, 73, 9, 14, 8}
    strings := Xs{"nut", "ape", "elephant", "zoo", "go"}

    Sort(ints)
    fmt.Printf("%v\n", ints)
    Sort(strings)
    fmt.Printf("%v\n", strings)
    ~~~


## Listing interfaces in interfaces

Take a look at the following example of an interface definition, this one is
from the package `container/heap`:

~~~go
type Interface interface {
    sort.Interface
    Push(x interface{})
    Pop() interface{}
}
~~~

Here another interface is listed inside the definition of `heap.Interface`, this
may look odd, but is perfectly valid, remember that on the surface an interface is nothing
more than a listing of methods. `sort.Interface` is also such a listing, so it is
perfectly legal to include it in the interface.


## Introspection and reflection

In the following example we want to look at the "tag" (here named "namestr")
defined in the type definition of `Person`. To do this we need the
`reflect`(((package,reflect))) package (there is no other way in Go). Keep in
mind that looking at a tag means going back to the *type* definition. So we use
the `reflect` package to figure out the type of the variable and *then* access
the tag.

{callout="//"}
~~~go
type Person struct {
    name string "namestr"
    age  int
}

func ShowTag(i interface{}) { //<1>
    switch t := reflect.TypeOf(i); t.Kind() {
    case reflect.Ptr: //<2>
        tag := t.Elem().Field(0).Tag
    //             <3>     <4>       <5>
~~~
Figure: Introspection using reflection.

We are calling `ShowTag` at <1> with a `*Person`, so at <2> we're expecting
a `reflect.Ptr`. We are dealing with a `Type` <3> and according to the
documentation ^[`go doc reflect`]:

> Elem returns a type's element type.
> It panics if the type's Kind is not Array, Chan, Map, Ptr, or Slice.

So on `t` we use `Elem()` to get the value the pointer points to. We have now
dereferenced the pointer and are "inside" our structure. We then <4> use
`Field(0)` to access the zeroth field.

The struct `StructField` has a `Tag` member which returns the tag-name as
a string. So on the $$0^{th}$$ field we can unleash `.Tag` <5> to access this
name: `Field(0).Tag`. This gives us `namestr`.

To make the difference between types and values more clear, take a look at the
following code:

{callout="//"}
~~~go
func show(i interface{}) {
    switch t := i.(type) {
    case *Person:
        t := reflect.TypeOf(i)  //<1>
        v := reflect.ValueOf(i) //<2>
        tag := t.Elem().Field(0).Tag //<3>
        name := v.Elem().Field(0).String() //<4>
    }
}
~~~
Figure: Reflection and the type and value.

At <1> we create `t` the type data of `i`, and `v` gets the actual values at
<2>. Here at <3> we want to get to the "tag". So we need `Elem()` to redirect
the pointer, access the first field and get the tag. Note that we operate on `t`
a `reflect.Type`. Now <4> we want to get access to the *value* of one of the
members and we employ `Elem()` on `v` to do the redirection. we have "arrived"
at the structure. Then we go to the first field `Field(0)` and invoke the
`String()` method on it.

![Peeling away the layers using reflection.](fig/reflection.png "Peeling away the layers using reflection.
Going from a `*Person` via `Elem` using the methods described in `go doc reflect` to get the actual `string` contained within.")

Setting a value works similarly as getting a value, but only works on
*exported* members. Again some code:

~~~go
type Person struct {
    name string
    age  int
}

func Set(i interface{}) {
    switch i.(type) {
    case *Person:
        r := reflect.ValueOf(i)
        r.Elem(0).Field(0).SetString("Albert Einstein")
    }
}
~~~
Figure: Reflect with *private* member.

~~~go
type Person struct {
    Name string
    age  int
}

func Set(i interface{}) {
    switch i.(type) {
    case *Person:
        r := reflect.ValueOf(i)
        r.Elem().Field(0).SetString("Albert Einstein")
    }
}
~~~
Figure: Reflect with *public* member.

The first program compiles and runs, but when you run it, you are greeted with a
stack trace and a *run time* error:
"panic: reflect.Value.SetString using value obtained using unexported field".

The second program works OK and sets the member `Name` to "Albert Einstein".
Of course this only works when you call `Set()` with a pointer argument.


## Exercises

{{ex/interfaces//ex.md}}
