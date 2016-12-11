{.epigraph}
> "^"
Quote: Answer to whether there is a bit wise negation operator -- Ken Thompson

A package is a collection of functions and data.

You declare a package with the `package`(((keywords, package))) keyword. The
filename does not have to match the package name. The convention for package
names is to use lowercase characters. Go packages may consist of multiple files,
but they share the `package <name>` line. Let's define a package `even` in the
file `even.go`.

(((functions, exported)))
(((functions, private)))
(((functions, public)))
{callout="//"}
<{{src/packages/even.go}}

Here <1> we start a new namespace: "even". The function `Even` <2> starts with
a capital letter. This means the function is *exported*, and may be used outside
our package (more on that later). The function `odd` <3> does not start with
a capital letter, so it is a *private* function.

Now we just need to build the package. We create a directory under `$GOPATH`,
and copy `even.go` there (see (#compiling-and-running-code) in (#basics)).

    % mkdir $GOPATH/src/even
    % cp even.go $GOPATH/src/even
    % go build
    % go install

Now we can use the package in our own program `myeven.go`:

{callout="//"}
<{{src/packages/myeven.go}}

Import <1> the following packages. The *local* package `even` is imported here
<2>. This <3> imports the official `fmt` package. And now we use <4> the
function from the `even` package. The syntax for accessing a function from
a package is `<package>.FunctionName()`. And finally we can build our program.

    % go build myeven.go
    % ./myeven
    Is 5 even? false

If we change our `myeven.go` at <4> to use the unexported function `even.odd`:
`fmt.Printf("Is %d even? %v\n", i, even.odd(i))` We get an error when compiling,
because we are trying to use a
*private* function:

    myeven.go: cannot refer to unexported name even.odd

Note that the "starts with capital $$\rightarrow$$ exported", "starts with
lower\-case $$\rightarrow$$ private" rule also extends to other names (new
types, global variables) defined in the package. Note that the term "capital" is
not limited to US-ASCII -- it extends to all bicameral alphabets (Latin, Greek,
Cyrillic, Armenian and Coptic).


## Identifiers
The Go standard library names some function with the old (Unix) names while
others are in CamelCase. The convention is to leave well-known legacy
not-quite-words alone rather than try to figure out where the capital letters
go:  `Atoi`, `Getwd`, `Chmod`. CamelCasing works best when you have whole words
to work with: `ReadFile`, `NewWriter`, `MakeSlice`. The convention in Go is to
use CamelCase rather than underscores to write multi-word names.

As we did above in our `myeven` program, accessing content from an imported
(with `import` (((keywords, import)))) package is done with using the package's
name and then a dot.  After (((package, bytes))) `import "bytes"` the importing
program can talk about `bytes.Buffer`. A package name should be good, short,
concise and evocative. The convention in Go is that package names are lowercase,
single word names.

The package name used in the `import` statement is the default name used. But if
the need arises (two different packages with the same name for instance), you
can override this default: `import bar "bytes"` The function `Buffer` is now
accessed as `bar.Buffer`.

Another convention is that the package name is the base name of its source
directory; the package in `src/compress/gzip` is imported as `compress/gzip` but
has name `gzip`, not `compress/gzip`.

It is important to avoid stuttering when naming things. For instance, the
buffered reader type in the `bufio` (((package, bufio))) package is called
`Reader`, not `BufReader`, because users see it as `bufio.Reader`, which is
a clear, concise name.

Similarly, the function to make new instances of `ring.Ring` (package
`container/ring`), would normally be called `NewRing`, but since `Ring` is the
only type exported by the package, and since the package is called
`ring`(((package, ring))), it's called just `New`. Clients of the package see
that as `ring.New`. Use the package structure to help you choose good names.

Another short example is `once.Do` (see package `sync`); `once.Do(setup)` reads
well and would not be improved by writing `once.DoOrWaitUntilDone(setup)`. Long
names don't automatically make things more readable.


## Documenting packages
When we created our `even` package, we skipped over an important item:
documentation. Each package should have a *package comment*, a block comment
preceding the `package` clause. In our case we should extend the beginning of
the package, with:

~~~go
// The even package implements a fast function for detecting if an integer
// is even or not.
package even
~~~

When running `go doc` this will show up at the top of the page. When a package
consists of multiple files the package comment should only appear in one
file. A common convention (in really big packages) is to have a separate
`doc.go` that only holds the package comment. Here is a snippet from the
official `regexp` package:

~~~go
/*
    The regexp package implements a simple library for
    regular expressions.

    The syntax of the regular expressions accepted is:

    regexp:
        concatenation { '|' concatenation }
*/
package regexp
~~~


Each defined (and exported) function should have a small line of text
documenting the behavior of the function. Again to extend our `even` package:

~~~go
// Even returns true of i is even. Otherwise false is returned.
func Even(i int) bool {
~~~

And even though `odd` is not exported, it's good form to document it as well.

~~~go
// odd is the opposite of Even.
func odd(i int) bool {
~~~


## Testing packages
In Go it is customary to write (unit) tests for your package. Writing tests
involves the `testing` package and the program `go test`(((tooling, go test))).
Both have excellent documentation.

The `go test` program runs all the test functions. Without any defined tests for
our `even` package, `go test` yields:

    % go test
    ?       even    [no test files]

Let us fix this by defining a test in a test file. Test files reside in the
package directory and are named `*_test.go`. Those test files are just like
other Go programs, but `go test` will only execute the test functions. Each test
function has the same signature and its name should start with
`Test`: `func TestXxx(t *testing.T)` .

When writing test you will need to tell `go test` whether a test was
successful or not. A successful test function just returns. When
the test fails you can signal this with the following
functions. These are the most important ones (see `go doc testing` or `go help testfunc` for more):


* `func (t *T) Fail()`, `Fail` marks the test function as having failed but
  continues execution.

* `func (t *T) FailNow()`, `FailNow` marks the test function as having failed
  and stops its execution. Any remaining tests in this file are skipped, and
  execution continues with the next test.

* `func (t *T) Log(args ...interface{})`, `Log` formats its arguments using
  default formatting, analogous to `Print()`, and records the text in the error
  log.

* `func (t *T) Fatal(args ...interface{})`, `Fatal` is equivalent to `Log()`
  followed by `FailNow()`.

Putting all this together we can write our test. First we pick a name:
`even_test.go`. Then we add the following contents:

{callout="//"}
<{{src/packages/even_test.go}}

A test file belongs to the current <1> package. This is not only convenient, but
also allows tests of unexported functions and structures. We then <2> import the
`testing` package. And finally the test we want to execute. The code here <3>
should hold no surprises: we check if the `Even` function works OK. And now, the
moment we have been waiting form executing the test.

    % go test
    ok      even    0.001s

Our test ran and reported `ok`. Success! If we redefine our test function, we
can see the result of a failed test:

~~~go
// Entering the twilight zone
func TestEven(t *testing.T) {
    if Even(2) {
        t.Log("2 should be odd!")
        t.Fail()
    }
}
~~~

We now get:

    FAIL    even    0.004s
    --- FAIL: TestEven (0.00 seconds)
        2 should be odd!
    FAIL

And you can act accordingly (by fixing the test for instance).

Writing new packages should go hand in hand with writing (some)
documentation and test functions. It will make your code better and it
shows that you really put in the effort.

The Go test suite also allows you to incorporate example functions which serve
as documentation *and* as tests. These functions need to start with `Example`.

{callout="//"}
~~~go
func ExampleEven() {
    if Even(2) {
        fmt.Printf("Is even\n")
    }
    // Output: //<1>
    // Is even
}
~~~

Those last two comments lines <1> are part of the example, `go test` uses those
to check the *generated* output with the text in the comments. If there is
a mismatch the test fails.


## Useful packages
The standard libary of Go includes a huge number of packages. It is very
enlightening to browse the `$GOROOT/src/pkg` directory and look at the
packages. We cannot comment on each package, but the following are worth
a mention: ^[The descriptions are copied from the packages' `go doc`.]

`fmt`
:   (((package, fmt)))
    Package `fmt` implements formatted I/O with functions analogous
    to C's `printf` and `scanf`. The format verbs are derived
    from C's but are simpler. Some verbs (%-sequences) that can be used:


    %v
    :   The value in a default format. when printing structs, the plus flag (%+v) adds field names.

    %#v
    :   a Go-syntax representation of the value.

    %T
    :   a Go-syntax representation of the type of the value.

`io`
:   (((package, io)))
    This package provides basic interfaces to I/O primitives.
    Its primary job is to wrap existing implementations of such primitives,
    such as those in package os, into shared public interfaces that
    abstract the functionality, plus some other related primitives.

`bufio`
:   (((package, bufio)))
    This package implements buffered I/O.  It wraps an
    `io.Reader`
    or
    `io.Writer`
    object, creating another object (Reader or Writer) that also implements
    the interface but provides buffering and some help for textual I/O.

`sort`
:   (((package, sort)))
    The `sort` package provides primitives for sorting arrays
    and user-defined collections.

`strconv`
:   (((package, strconv)))
    The `strconv` package implements conversions to and from
    string representations of basic data types.

`os`
:   (((package, os)))
    The `os` package provides a platform-independent interface to operating
    system functionality.  The design is Unix-like.

`sync`
:   (((package, sync)))
    The package `sync` provides basic synchronization primitives such as mutual
    exclusion locks.

`flag`
:   (((package, flag)))
    The `flag` package implements command-line flag parsing.

`encoding/json`
:   (((package, encoding/json)))
    The `encoding/json` package implements encoding and decoding of JSON objects as
    defined in RFC 4627 [@RFC4627].

`html/template`
:   (((package, html/template)))
    Data-driven templates for generating textual output such as HTML.

    Templates are executed by applying them to a data structure.  Annotations in
    the template refer to elements of the data structure (typically a field of
    a struct or a key in a map) to control execution and derive values to be
    displayed.  The template walks the structure as it executes and the "cursor"
    @ represents the value at the current location in the structure.

`net/http`
:   (((package, net/http)))
    The `net/http` package implements parsing of HTTP requests, replies,
    and URLs and provides an extensible HTTP server and a basic
    HTTP client.

`unsafe`
:   (((package, unsafe)))
    The `unsafe` package contains operations that step around the type safety of Go programs.
    Normally you don't need this package, but it is worth mentioning that *unsafe* Go programs
    are possible.

`reflect`
:   (((package, reflect)))
    The `reflect` package implements run-time reflection, allowing a program to
    manipulate objects with arbitrary types.  The typical use is to take a
    value with static type `interface{}` and extract its dynamic type
    information by calling `TypeOf`, which returns an object with interface
    type `Type`. See (#interfaces), Section (#introspection-and-reflection).

`os/exec`
:   (((package, os/exec))) The `os/exec` package runs external commands.


## Exercises
{{ex/packages/ex.md}}
