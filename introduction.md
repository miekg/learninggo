{.epigraph}
> Is Go an object-oriented language? Yes and no.
Quote: Frequently asked questions -- Go Authors


The Go programming language is an open source project language to make
programmers more productive.

According to the website [@go_web] "Go is expressive, concise, clean, and
efficient". And indeed it is. My initial interest was piqued when I read early
announcements about this new language that had built-in concurreny and a C-like
syntax (Erlang also has built-in concurrency, but I could never get used to its
syntax). Go is a compiled statically typed language that feels like
a dynamically typed, interpreted language. My go to (scripting!) language Perl
has taken a back seat now that Go is around.

The unique Go language is defined by these principles:

Clean and Simple
:   Go strives to keep things small and beautiful. You should
    be able to do a lot in only a few lines of code.

Concurrent
:   Go makes it easy to "fire off" functions to be
    run as *very* lightweight threads. These threads are called
    goroutines (((goroutine)))^[Yes, that sounds a lot like
    *co*routines, but goroutines are slightly different as we will
    see in (#communication).] in Go.

Channels
:   Communication with these goroutines is done, either via shared state or
    via (((channels))) channels [@csp].

Fast
:   Compilation is fast and execution is fast. The aim is
    to be as fast as C. Compilation time is measured in seconds.

Safe
:   Explicit casting and strict rules when converting one type to another.
    Go has garbage collection. No more `free()` in Go: the language takes care of this.

Standard format
:   A Go program can be formatted in (almost) any way the programmers want,
    but an official format exists. The rule is very simple:
    The output of the filter `gofmt` *is the officially endorsed
    format*.

Postfix types
:   Types are given *after* the variable name, thus `var a int`,
    instead of `int a`.

UTF-8
:   UTF-8 is everywhere, in strings
    *and* in the program code. Finally you can use $$\Phi = \Phi + 1$$ in your source code.

Open Source
:   The Go license is completely open source.

Fun
:   Programming with Go should be fun!

As I mentioned Erlang also shares some
features of Go. A notable difference between Erlang
and Go is that Erlang borders on being a functional language, while Go is imperative.
And Erlang runs in a virtual machine, while Go is compiled.


## How to Read this Book
I've written this book for people who already know some programming languages
and how to program. In order to use this book, you (of course) need Go installed
on your system, but you can easily try examples online in the Go
playground^[<http://play.golang.org>.]. All exercises in this book work with Go
1, the first stable release of Go -- if not, it's a bug.

The best way to learn Go is to create your own programs. Each chapter therefore
includes exercises (and answers to exercises) to acquaint you with the language.
Each exercise is either *easy*, *intermediate*, or *difficult*. The answers are
included after the exercises on a new page. Some exercises don't have an answer;
these are marked with an asterisk.

Here's what you can expect from each chapter:

(#basics)
:   We'll look at the basic types, variables, and control structures available in the language.

(#functions)
:   Here we look at functions, the basic building blocks of Go programs.

(#packages)
:   We'll see that functions and data can be grouped together
    in packages. We'll also see how to document and test our packages.

(#beyond-the-basics)
:   We'll create our own types. We'll also look at memory allocations in Go.

(#interfaces)
:   We'll learn how to use interfaces. Interfaces are the central concept in Go,
    as Go does not support object orientation in the traditional sense.

(#concurrency)
:   We'll learn the `go` keyword, which can be used to start function in
    separate routines (called goroutines). Communication with those goroutines is
    done via channels.

(#communication)
:   Finally we'll see how to interface with the rest of the world from within
    a Go program. We'll see how to create files and read and write to and from them.
    We'll also briefly look into networking.


## Official Documentation
There is a substantial amount of documentation written about Go. The Go Tutorial
[@go_tutorial], the Go Tour (with lots of exercises) and the Effective Go
[@effective_go] are helpful resources. The website <http://golang.org/doc/> is
a very good starting point for reading up on Go^[<http://golang.org/doc/> itself
is served by `godoc`.]. Reading these documents is certainly not required, but
it is recommended.

> When searching on the internet use the term "golang" instead of plain "go".

Go comes with its own documentation in the form of a program called
`godoc`^[When building from source it must be installed separately with `go get
golang.org/x/tools/cmd/godoc`.]. If you are interested in the documentation for
the built-ins, simply do this:

    % godoc builtin

To get the documentation of the `hash` package, just:

    % godoc hash

To read the documentation of `fnv` contained in `hash`, you'll need
to issue `godoc hash/fnv` as `fnv` is a subdirectory of `hash`.

~~~go
PACKAGE DOCUMENTATION

package fnv
    import "hash/fnv"

    Package fnv implements FNV-1 and FNV-1a, non-cryptographic hash
    ...
~~~
