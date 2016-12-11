{.epigraph}
> Go has pointers but not pointer arithmetic. You cannot use a pointer
> variable to walk through the bytes of a string.
Quote: Go For C++ Programmers -- Go Authors

In this chapter we delve deeper in to the language.

Go has pointers. There is however no pointer arithmetic, so they act more like
references than pointers that you may know from C. Pointers are useful. Remember
that when you call a function in Go, the variables are
*pass-by-value*. So, for efficiency and the possibility to modify a passed value
 *in* functions we have pointers.

You declare a pointer by prefixing the type with an '`*`': `var p *int`. Now `p`
is a pointer to an integer value. All newly declared variables are assigned
their zero value and pointers are no different. A newly declared pointer, or
just a pointer that points to nothing, has a nil-value (((nil))). In other
languages this is often called a NULL pointer in Go it is just `nil`. To make
a pointer point to something you can use the address-of operator
(((operators, address-of))) (`&`), which we demonstrate here:

{callout="//"}
~~~go
var p *int
fmt.Printf("%v", p) //<1>

var i int	    //<2>
p = &i		    //<3>

fmt.Printf("%v", p) //<4>
~~~

This <1> Prints `nil`. Declare <2> an integer variable `i`. Make `p` point <3>
to `i`, i.e. take the address of `i`. And this <4> will print something like
`0x7ff96b81c000a`. De-referencing a pointer is done by prefixing the pointer
variable with `*`.

As said, there is no pointer arithmetic, so if you write: `*p++`, it is
interpreted as `(*p)++`: first reference and then increment the
value.(((operators, increment)))


## Allocation
Go also has garbage collection, meaning that you don't have to worry about
memory deallocation.^[The downside is that you know have to worry about garbage
collection. If you really need it garbage collection in a Go program can be
disabled by running it with the environment variable `GOGC` set to `off`:
`GOGC=off ./myprogram`.]

To allocate memory Go has two primitives, `new` and `make`. They do different
things and apply to different types, which can be confusing, but the rules are
simple. The following sections show how to handle allocation in Go and hopefully
clarifies the somewhat artificial distinction between `new` (((built-in, new)))
and `make` (((built-in,make))).


### Allocation with new
The built-in function `new` is essentially the same as its namesakes in other
languages: `new(T)` allocates zeroed storage for a new item of type `T` and
returns its address, a value of type `*T`. Or in other words, it returns
a pointer to a newly allocated zero value of type `T`. This is important to
remember.

The documentation for `bytes.Buffer` states that "the zero value for Buffer is
an empty buffer ready to use.". Similarly, `sync.Mutex` does not have an
explicit constructor or Init method. Instead, the zero value for a `sync.Mutex`
is defined to be an unlocked mutex.


### Allocation with make
The built-in function `make(T, args)` serves a purpose different from `new(T)`.
It creates slices, maps, and channels *only*, and it returns an initialized (not
zero!) value of type `T`, and not a pointer: `*T`. The reason for the
distinction is that these three types are, under the covers, references to data
structures that must be initialized before use. A slice, for example, is
a three-item descriptor containing a pointer to the data (inside an array), the
length, and the capacity; until those items are initialized, the slice is `nil`.
For slices, maps, and channels, `make` initializes the internal data structure
and prepares the value for use.

For instance, `make([]int, 10, 100)` allocates an array of 100 ints and then
creates a slice structure with length 10 and a capacity of 100 pointing at the
first 10 elements of the array. In contrast, `new([]int)` returns a pointer to
a newly allocated, zeroed slice structure, that is, a pointer to a `nil` slice
value. These examples illustrate the difference between `new` and `make`.

{callout="//"}
~~~go
var p *[]int = new([]int)       //<1>
var v  []int = make([]int, 100) //<2>

var p *[]int = new([]int)       //<3>
*p = make([]int, 100, 100)

v := make([]int, 100)           //<4>
~~~

Allocates <1> slice structure; rarely useful. `v` <2> refers to a new array of
100 ints. At <3> we make it unnecessarily complex, <4> is more idiomatic.

Remember that `make` applies only to maps, slices and channels and does not
return a pointer. To obtain an explicit pointer allocate with `new`.

A> **new** allocates; **make** initializes.
A>
A> The above two paragraphs can be summarized as:
A>
A> * `new(T)` returns `*T` pointing to a zeroed `T`
A> * `make(T)` returns an initialized `T`
A>
A> And of course `make` is only used for slices, maps and channels.


### Constructors and composite literals
Sometimes the zero value isn't good enough and an initializing constructor is
necessary, as in this example taken from the package `os`.

~~~go
func NewFile(fd int, name string) *File {
    if fd < 0 {
        return nil
    }
    f := new(File)
    f.fd = fd
    f.name = name
    f.dirinfo = nil
    f.nepipe = 0
    return f
}
~~~

There's a lot of boiler plate in there. We can simplify it using a
*composite literal* (((literal, composite))), which is an expression that
 creates a new instance each time it is evaluated.

{callout="//"}
~~~go
func NewFile(fd int, name string) *File {
    if fd < 0 {
        return nil
    }
    f := File{fd, name, nil, 0}
    return &f	//<1>
}
~~~

It is OK to return the address of a local variable <1> the storage associated
with the variable survives after the function returns.

In fact, taking the address of a composite literal allocates a fresh instance
each time it is evaluated, so we can combine these last two lines.^[Taking the
address of a composite literal tells the compiler to allocate it on the heap,
not the stack.]

~~~go
return &File{fd, name, nil, 0}
~~~

The items (called fields) of a composite literal are laid out in order and must
all be present. However, by labeling the elements explicitly as field:value
pairs, the initializers can appear in any order, with the missing ones left as
their respective zero values. Thus we could say

~~~go
return &File{fd: fd, name: name}
~~~

As a limiting case, if a composite literal contains no fields at all, it creates
a zero value for the type. The expressions `new(File)` and `&File{}` are
equivalent. In fact the use of `new` is discouraged.

Composite literals can also be created for arrays, slices, and maps, with the
field labels being indices or map keys as appropriate. In these examples, the
initializations work regardless of the values of `Enone`, and `Einval`, as long
as they are distinct:

~~~go
ar := [...]string{Enone: "no error", Einval: "invalid argument"}
sl := []string{Enone: "no error", Einval: "invalid argument"}
ma := map[int]string {Enone: "no error", Einval: "invalid argument"}
~~~


## Defining your own types
Of course Go allows you to define new types, it does this with the
`type`(((keywords, type))) keyword: `type foo int`

This creates a new type `foo` which acts like an `int`. Creating more sophisticated
types is done with the `struct` (((keywords, struct))) keyword. An example would
be when we want record somebody's name (`string`) and age (`int`) in a single
structure and make it a new type:

<{{src/beyond/struct.go}}

Apropos, the output of `fmt.Printf("%v\n", a)` is `&{Pete 42}`

That is nice! Go knows how to print your structure. If you only want to print
one, or a few, fields of the structure you'll need to use `.<field name>`. For
example to only print the name:

~~~go
fmt.Printf("%s", a.name)
~~~



### More on structure fields
As said each item in a structure is called a field(((field))). A struct with no
fields: `struct {}`. Or one with four fields:

~~~go
struct {
    x, y int
    A *[]int
    F func()
}
~~~

If you omit the name for a field, you create an anonymous field (((field,
anonymous))), for instance:

~~~go
struct {
    T1        // Field name is T1.
    *T2       // Field name is T2.
    P.T3      // Field name is T3.
    x, y int  // Field names are x and y.
}
~~~

Note that field names that start with a capital letter are exported, i.e. can be
set or read from other packages. Field names that start with a lowercase are
private to the current package. The same goes for functions defined in packages,
see (#packages) for the details.


### Methods
If you create functions that work on your newly defined type, you can take two
routes:

1. Create a function that takes the type as an argument.

~~~go
func doSomething(n1 *NameAge, n2 int) { /* */ }
~~~

2. Create a function that works on the type (see *receiver* in (#functions)):

~~~go
func (n1 *NameAge) doSomething(n2 int) { /* */ }
~~~

   This is a method call, which can be used as:

~~~go
var n *NameAge
n.doSomething(2)
~~~

Whether to use a function or method is entirely up to the programmer, but if you
want to satisfy an interface (see the next chapter) you must use methods. If no
such requirement exists it is a matter of taste whether to use functions or
methods.

But keep the following in mind, this is quoted from [@go_spec]:

> If `x` is
> addressable and `&x`'s method set contains `m`,
> `x.m()` is shorthand for `(&x).m()`.

In the above case this means that the following is *not* an error:

~~~go
var n NameAge	    // Not a pointer
n.doSomething(2)
~~~

Here Go will search the method list for `n` of type `NameAge`, come up empty and
will then *also* search the method list for the type `*NameAge` and will
translate this call to `(&n).doSomething(2)`.

There is a subtle but major difference between the following type declarations.
Also see the Section "Type Declarations" [@go_spec]. Suppose we have:

~~~go
// A Mutex is a data type with two methods, Lock and Unlock.
type Mutex struct         { /* Mutex fields */ }
func (m *Mutex) Lock()    { /* Lock impl. */ }
func (m *Mutex) Unlock()  { /* Unlock impl. */ }
~~~

We now create two types in two different manners:

* `type NewMutex Mutex`.
* `type PrintableMutex struct{Mutex}`.

`NewMutex` is equal to `Mutex`, but it *does not* have *any* of the methods of
`Mutex`. In other words its method set is empty. But `PrintableMutex` *has*
*inherited* (((methods, inherited))) the method set from `Mutex`. The Go term
 for this is *embedding* (((structures, embed))). In the words of [@go_spec]:

> The method set of `*PrintableMutex` contains the methods
> `Lock` and `Unlock` bound to its anonymous field `Mutex`.


## Conversions
Sometimes you want to convert a type to another type. This is possible in Go,
but there are some rules. For starters, converting from one value to another is
done by operators (that look like functions: `byte()`) and not all conversions
are allowed.

{#tab-conversion}
{{tab/conversion.md}}

* From a `string` to a slice of bytes or runes.

    ~~~go
    mystring := "hello this is string"
    byteslice := []byte(mystring)
    ~~~

    Converts to a `byte` slice, each `byte` contains the integer value of the
    corresponding byte in the string. Note that as strings in Go are encoded in
    UTF-8 some characters in the string may end up in 1, 2, 3 or 4 bytes.

    ~~~go
    runeslice  := []rune(mystring)
    ~~~

    Converts to an `rune` slice, each `rune` contains a Unicode code point.
    Every character from the string corresponds to one rune.

* From a slice of bytes or runes to a `string`.

    ~~~go
    b := []byte{'h','e','l','l','o'} // Composite literal.
    s := string(b)
    i := []rune{257,1024,65}
    r := string(i)
    ~~~

For numeric values the following conversions are defined:

* Convert to an integer with a specific (bit) length: `uint8(int)`
* From floating point to an integer value: `int(float32)`. This discards the
  fraction part from the floating point value.
* And the other way around: `float32(int)`.


### User defined types and conversions
How can you convert between the types you have defined yourself? We create two
types here `Foo` and `Bar`, where `Bar` is an alias for `Foo`:

~~~go
type foo struct { int }  // Anonymous struct field.
type bar foo             // bar is an alias for foo.
~~~

Then we:

~~~go
var b bar = bar{1} // Declare `b` to be a `bar`.
var f foo = b	   // Assign `b` to `f`.
~~~

Which fails on the last line with:
`cannot use b (type bar) as type foo in assignment`

This can be fixed with a conversion: `var f foo = foo(b)`

Note that converting structures that are not identical in their fields is more
difficult. Also note that converting `b` to a plain `int` also fails; an integer
is not the same as a structure containing an integer.

## Exercises
{{ex/beyond/ex.md}}
