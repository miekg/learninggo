{class="epigraph"}
> I am interested in this and hope to do something.
Quote: On adding complex numbers to Go -- Ken Thompson


In this chapter we will look at the basic building blocks of the Go programming
language.

## Hello World
In the Go tutorial, you get started with Go in the typical manner: printing
"Hello World" (Ken Thompson and Dennis Ritchie started this when they presented
the C language in the 1970s). That's a great way to start, so here it is, "Hello
World" in Go.

{callout="//"}
<{{src/basics/helloworld.go}}

Lets look at the program line by line.
This first line is just required <1>. All Go files start with
`package <something>`, and `package main` is required for a standalone executable.

`import "fmt"` says we need `fmt` in
addition to `main` <2>. A package other than `main` is commonly called a
library, a familiar concept in many programming languages (see chapter (#chap:packages).
The line ends with a comment that begins with `//`.

Next we another comment, but this one is enclosed in `/*` `*/` <3>.
When your Go program is executed, the first function called will be
`main.main()`, which mimics the behavior from C. Here we declare that function <4>.

Finally we call a function from the package `fmt` to print a
string to the screen. The string is enclosed with `"` and may
contain non-ASCII characters <5>.

## Compiling and Running Code
To build a Go program, use the `go` tool.(((tooling,go)))
To build `helloworld` we just enter:

    % go build helloworld.go

(((tooling,go build)))
This results in an executable called `helloworld`. (((tooling,go run)))

    % ./helloworld
    Hello, world.

You can combine the above and just call `go run helloworld.go`.

<!--
\section{Variables, Types and Keywords}
\label{sec:vars}
In the next few sections we will look at the variables, basic types,
keywords, and control structures of our new language.

Go is different from (most) other languages in that the type of a variable
is specified \emph{after} the variable name. So not:
\lstinline{int a}, but \lstinline{a int}. When you declare a variable it
is assigned the ``natural'' null value for the type. This means that after
\lstinline{var a int}, \lstinline{a} has a value of 0. With
\lstinline{var s string}, \lstinline{s} is assigned the zero string,
which is \lstinline{""}.
Declaring and assigning in Go is a two step process, but they may
be combined. Compare the following pieces of code which have
the same effect.
(((variables,declaring)))
(((variables,assigning)))

\begin{minipage}{.5\textwidth}
\begin{lstlisting}[linewidth=.5\textwidth,numbers=none]
var a int
var b bool
a = 15
b = false
\end{lstlisting}
\hfill
\end{minipage}
\begin{minipage}{.5\textwidth}
\begin{lstlisting}[linewidth=.5\textwidth,numbers=none]


a := 15
b := false
\end{lstlisting}
\hfill
\end{minipage}

On the left we use the
\key{var} keyword to declare a variable and \emph{then} assign a value to
it. The code on the right uses \mbox{\key{:=}{ }} to do this in one
step (this form may only be used \emph{inside} functions).
In that case the variable
type is \emph{deduced} from the value. A value of 15 indicates an \type{int}.
A value of \texttt{false} tells Go that the type should be \type{bool}.
Multiple \key{var} declarations may also be grouped; \key{const} (see ``\nameref{sec:constants}'')
and \key{import} also allow this. Note the use of parentheses instead of braces:
\begin{lstlisting}[numbers=none]
var (
    x int
    b bool
)
\end{lstlisting}
Multiple variables of the same type can also be declared on a
single line: \lstinline{var x, y int} makes \var{x} and \var{y} both
\type{int} variables. You can also make use of \first{parallel
assignment}{parallel assignment}: \lstinline{a, b := 20, 16}.
This makes \var{a} and \var{b} both integer variables and assigns
20 to \var{a} and 16 to \var{b}.

A special name for a variable is \var{\textbf{\_}} \index{variables!\_}
(underscore)\index{variables!underscore}. Any value
assigned to it is discarded (it's similar to \file{/dev/null} on Unix). In this example we only assign the integer
value of 35 to \var{b} and discard the value 34: \lstinline{_, b := 34, 35}.
Declared but otherwise \emph{unused} variables are a compiler error in Go.

\subsection{Boolean Types}
A boolean type represents the set of boolean truth values denoted by the
predeclared constants \emph{true} and \emph{false}. The boolean type is \type{bool}.

\subsection{Numerical Types}
Go has most of the well-known types such as \lstinline{int}. The \lstinline{int} type
has the appropriate length for your machine,
meaning that on a 32-bit machine it is 32 bits and on
a 64-bit machine it is 64 bits. Note: an \lstinline{int} is
either 32 or 64 bits, no other values are defined. Same goes
for \lstinline{uint}, the unsigned int.

If you want to be explicit about the length, you can have that too,
with \type{int32}, or \type{uint32}. The full
list for (signed and unsigned) integers is
\type{int8}, \type{int16}, \type{int32}, \type{int64} and
\type{byte}, \type{uint8}, \type{uint16}, \type{uint32}, \type{uint64},
with \lstinline{byte} being an
alias for \lstinline{uint8}. For floating point values there is
\type{float32} and \type{float64} (there is no \lstinline{float} type).
A 64 bit integer or floating point value is \emph{always} 64 bit, also on 32 bit
architectures.

Note
that these types are all distinct and assigning variables which mix
these types is a compiler error, like in the following code:
\lstinputlisting[label=src:types,numbers=none]{src/basics/types.go}

We declare two different integers, a and b where a is an \type{int} and
b is an \type{int32}. We want to set b to the sum of a and a. This
fails and gives the error:
\error{cannot use a + a (type int)  as type int32 in assignment}.
Adding the constant 5 to b \emph{does} succeed, because constants are
not typed.

\subsection{Constants}
\label{sec:constants}
Constants in Go are just that --- constant. They are created at compile
time, and can only be numbers, strings, or booleans;
\lstinline{const x = 42} makes \var{x} a constant. You can use
\first{\key{iota}}{keyword!iota} \footnote{The word [iota] is used in a common English phrase,
'not one iota', meaning 'not the slightest difference', in reference to
a phrase in the New Testament: ``\emph{until heaven and earth pass away, not an
iota, not a dot, will pass from the Law}.'' \cite{iota}}
to enumerate values.
\begin{lstlisting}[numbers=none]
const (
	a = iota
	b
)
\end{lstlisting}
The first use of \key{iota} will yield 0, so \var{a} is equal to 0. Whenever
\key{iota} is used again on a new line its value is incremented with 1, so \var{b}
has a value of 1. Or, as shown here, you can even let Go repeat the use of \key{iota}.
You may also explicitly type a constant: \lstinline{const b string = "0"}. Now
\var{b} is a \type{string} type constant.

\subsection{Strings}
Another important built-in type is \lstinline{string}. Assigning a
string is as simple as:
\begin{lstlisting}[numbers=none]
s := "Hello World!"
\end{lstlisting}
Strings in Go are a sequence of UTF-8 characters enclosed in double
quotes ("). If you use the single quote (') you mean one character
(encoded in UTF-8) --- which is \emph{not} a \lstinline{string} in Go.

Once assigned to a variable, the string cannot be changed: strings in Go are
immutable. If you are coming from C, not that the following is not legal in Go:
\begin{lstlisting}[numbers=none]
var s string = "hello"
s[0] = 'c'
\end{lstlisting}
To do this in Go you will need the following:
\begin{lstlisting}[numbers=none]
s := "hello"
c := []rune(s)	    |\longremark{Here we convert \var{s} to an array of runes \citem.}|
c[0] = 'c'	    |\longremark{We change the first element of this array \citem.}|
s2 := string(c)     |\longremark{Then we create a \emph{new} string \var{s2} with the alteration \citem.}|
fmt.Printf("%s\n", s2) |\longremark{Finally, we print the string with \func{fmt.Printf} \citem.}|
\end{lstlisting}
\showremarks

%% remove this section.
%%Due to the insertion of semicolons (see \cite{effective_go} section
%%``Semicolons''), you need to be careful with using multi line strings. If
%%you write:
%%\begin{lstlisting}[numbers=none]
%%s := "Starting part"
%%    + "Ending part"
%%\end{lstlisting}
%%This is transformed into:
%%\begin{lstlisting}[numbers=none]
%%s := "Starting part";
%%    + "Ending part";
%%\end{lstlisting}
%%Which is not valid syntax, you need to write:
%%\begin{lstlisting}[numbers=none]
%%s := "Starting part" +
%%     "Ending part"
%%\end{lstlisting}
%%Then Go will not insert the semicolons in the wrong places. Another way
%%would be to use \emph{raw} string literals\index{string literal!raw} by using backquotes (\key{`}):
%%\begin{lstlisting}[numbers=none]
%%s := `Starting part
%%     Ending part`
%%\end{lstlisting}
%%Be aware that in this last example \var{s} now also contains the newline.
%%Unlike \emph{interpreted} string literals \index{string literal!interpreted} the value of a raw string literal
%%is composed of the \emph{uninterpreted} characters between the quotes.

\subsection{Runes}
\lstinline{Rune} is an alias for \type{int32}. It is an UTF-8 encoded code point. When is this type useful?
One example is when you're
iterating over characters in a string. You could loop over each byte (which is only equivalent to a character
when strings are encoded in 8-bit ASCII, which they are \emph{not} in Go!). But to get the actual characters you
should use the \type{rune} type.

\subsection{Complex Numbers}
Go has native support for complex numbers. To
use them you need a variable of type \type{complex128} (64
bit real and imaginary parts) or \type{complex64} (32 bit
real and imaginary parts).
Complex numbers are written as
\var{re + im$i$}, where \var{re} is the real part,
\var{im} is the imaginary part and $i$ is the literal '$i$' ($\sqrt{-1}$).

\subsection{Errors}
Any non-trivial program will have the need for error reporting sooner or later. Because of this
Go has a builtin type specially for errors, called \type{error}.
\lstinline{var e error} creates a variable \var{e} of type \type{error} with the
value \lstinline{nil}. This error type is an interface -- we'll look more at interfaces in Chapter
``\ref{chap:interfaces}''. For now you can just assume that \type{error} is a type just like all other types.

\section{Operators and Built-in Functions}
\label{sec:builtins}
Go supports the normal set of numerical operators.
Table \ref{tab:op-precedence}
lists the current ones and their relative precedence. They
all associate from left to right.

\begin{table}[Hh!]
\begin{center}
\caption{Operator precedence}
\label{tab:op-precedence}
\input{tab/precedence.tex}
\end{center}
\end{table}
\verb|+ - * /| and \verb|%| all do what you would expect,
\verb!& | ^!
and \verb!&^! are bit operators for
\first{bitwise \emph{and}}{operator!bitwise!and},
\first{bitwise \emph{or}}{operator!bitwise!or}, \first{bitwise \emph{xor}}{operator!bit
wise xor}, and \first{bit clear}{operator!bitwise!clear} respectively.
The \verb|&&| and \verb/||/ operators are
logical \first{\emph{and}}{operator!and} and
logical \first{\emph{or}}{operator!or}. Not listed in the table
is the logical \first{not}{operator!not}: \verb/!/

Although Go does not support operator overloading (or method
overloading for that matter), some of the built-in
operators \emph{are} overloaded. For instance, \texttt{+} can be used for integers,
floats, complex numbers and strings (adding strings is concatenating them).

\section{Go Keywords}
Let's start looking at keywords. Table \ref{tab:keywords} lists all the keywords in Go.
\begin{table}[Hh!]
\begin{center}
\caption{Keywords in Go}
\label{tab:keywords}
\input{tab/keywords.tex}
\end{center}
\end{table}
We've seen some of these already. We used \key{var} and \key{const} in the
``\nameref{sec:vars}'' section on page \pageref{sec:vars}, and we briefly looked at \key{package} and \key{import} in our "Hello World" program at the start of the chapter.
Others need more attention and have their own chapter or section:
\begin{itemize}
\item \key{func} is used to declare functions and methods.
\item \key{return} is used to return from functions. We'll look at both \key{func} and \key{return} in detail in Chapter \ref{chap:functions}.
\item \key{go} is used for concurrency. We'll look at this in Chapter \ref{chap:channels}.
\item \key{select} used to choose from different types of communication, We'll work with \key{select} in Chapter \ref{chap:channels}.
\item \key{interface} is covered in Chapter \ref{chap:interfaces}.
\item \key{struct} is used for abstract data types. We'll work with \key{struct} in Chapter \ref{chap:beyond}.
\item \key{type} is also covered in Chapter \ref{chap:beyond}.
\end{itemize}

\section{Control Structures}
There are only a few control structures in Go. To write loops we use the \key{for} keyword, and there is a
\key{switch} and of course an \key{if}. When working with channels \key{select} will be used (see Chapter \ref{chap:channels}).
Parentheses are are not required around the condition, and the body must \emph{always} be brace-delimited.

\subsection{If-Else}
In Go an \first{\key{if}}{keyword!if} looks like this:
\begin{lstlisting}
if x > 0 {
    return y
} else {
    return x
}
\end{lstlisting}

\index{keyword!return}

\index{keyword!if} \index{keyword!else}
Since \key{if} and \key{switch} accept an initialization statement, it's common to
see one used to set up a (local) variable.
\begin{lstlisting}[numbers=none]
if err := SomethingFunction(); err == nil {
    // do something
} else {
    return err
}
\end{lstlisting}

It is idomatic in Go to omit the \key{else} when the \key{if} statement's body has a \key{break}, \key{continue}, \key{return} or,
\key{goto}, so the above code would be better written as:
\begin{lstlisting}[numbers=none]
if err := SomethingFunction(); err != nil {
    return err
}
// do something
\end{lstlisting}
The opening brace on the first line must be positioned on the same line as the \key{if} statement. There is no
arguing about this, because this is what \prog{gofmt} outputs.

\subsection{Goto}
Go has a \first{\key{goto}}{keyword!goto} statement --- use it wisely. With \key{goto}
you jump to a \index{label} label which must be defined within the current function.
For instance, a loop in disguise:
\begin{lstlisting}[numbers=none]
func myfunc() {
        i := 0
Here:
        fmt.Println(i)
        i++
        goto Here
}
\end{lstlisting}
The string \var{Here:} indicates a label. A label does not need to start with a capital letter and is case sensitive.

\subsection{For}
\label{sec:for}
The Go \first{\key{for}}{keyword!for} loop has three forms, only one of
which has semicolons:
\begin{itemize}
    \item \lstinline|for init; condition; post { }| -- a loop using the syntax borrowed from C;
    \item \lstinline|for condition { }| -- a while loop, and;
    \item \lstinline|for { }| -- an endless loop.
\end{itemize}
Short declarations make it easy to declare the index variable right in the loop.
\begin{lstlisting}[numbers=none]
sum := 0
for i := 0; i < 10; i++ {
    sum = sum + i
}
\end{lstlisting}
Note that the variable \var{i} ceases to exist after the loop.

\subsection{Break and Continue}
With \first{\key{break}}{keyword!break} you can quit loops early.  By itself, \key{break} breaks
the current loop.
\begin{lstlisting}[numbers=none]
for i := 0; i < 10; i++ {
    if i > 5 {
	break|\longremark{Here we \key{break} the current loop \citem, and don't continue with the \prog{fmt.Println(i)} statement \citemnext.}|
    }
    fmt.Println(i)|\longremark{So we only print 0 to 5. With loops within loop you can specify a label after \key{break} to identify \emph{which} loop to stop:}|
}
\end{lstlisting}
\showremarks

\begin{lstlisting}[numbers=none]
J:  for j := 0; j < 5; j++ { |\longremark{Here we define a label "J" \citem, preceding the \key{for}-loop there.}|
        for i := 0; i < 10; i++ {
            if i > 5 {
                break J |\longremark{When we use \key{break J} \citem, we don't break the inner loop but the "J" loop.}|
            }
            fmt.Println(i)
        }
    }
\end{lstlisting}
\showremarks

With \first{\key{continue}}{keyword!continue} you begin the next iteration of the
loop, skipping any remaining code. In the same way as \key{break},
\key{continue} also accepts a label.

\subsection{Range}
The keyword \first{\key{range}}{keyword!range} can be used for loops. It
can loop over slices, arrays, strings, maps and channels (see Chapter
\ref{chap:channels}). \key{range} is
an iterator that, when called, returns the next key-value pair from the "thing" it
loops over. Depending on what that is, \key{range} returns different things.

When looping over a slice or array, \key{range} returns the index in the
slice as the key and value belonging to that index.
Consider this code: \index{keyword!range!on slices}
\begin{lstlisting}[numbers=none]
list := []string{"a", "b", "c", "d", "e", "f"}
for k, v := range list {
    // do something with k and v
}
\end{lstlisting}
First we create a slice of strings. Then we use \key{range} to loop over them. With each iteration, \key{range} will return the index as an \type{int} and the key as a \type{string}.
It will start with 0 and "a", so \var{k} will be 0 through 5, and v will be "a" through "f".

You can also use \key{range} on strings directly. Then it
will break out the individual Unicode characters
\footnote{In the UTF-8 world characters are sometimes called \first{runes}{runes}.
Mostly, when people talk about
characters, they mean 8 bit characters. As UTF-8 characters may be up to 32 bits the word
rune is used. In this case the type of \var{char} is \type{rune}.} and their start position, by parsing the UTF-8.
The loop: \index{keyword!range!on maps}
\begin{lstlisting}[numbers=none]
for pos, char := range "a|$\Phi{}$|x" {
    fmt.Printf("character '%c' starts at byte position %d\n", char, pos)
}
\end{lstlisting}
prints
\begin{alltt}
character 'a' starts at byte position 0
character '\begin{math}\Phi\end{math}' starts at byte position 1
character 'x' starts at byte position 3
\end{alltt}
Note that '\begin{math}\Phi\end{math}' took 2 bytes, so 'x' starts at byte 3.

\subsection{Switch}
Go's \first{\key{switch}}{keyword!switch} is very flexible; you can match on much more than just
integers.
The cases are evaluated top to bottom until
a match is found, and if the \key{switch} has no expression it switches on
\type{true}. It's therefore possible -- and idiomatic -- to write an
\key{if-else-if-else} chain as a \key{switch}.
\begin{lstlisting}[numbers=none]
// Convert hexadecimal character to an int value
switch { |\longremark{A \key{switch} without a condition is the same as \key{switch true} \citem.}|
case '0' <= c && c <= '9':|\longremark{We list the different cases. Each \key{case} statement has a condition that is either %
true of false. Here \citem{} we check if \var{c} is a number.}|
    return c - '0'|\longremark{If \var{c} is a number we return its value \citem.}|
case 'a' <= c && c <= 'f':|\longremark{Check if \var{c} falls between ``a'' and ``f'' \citem. For an ``a'' we return 10, for ``b'' we return 11, etc. We also do the same \citemnext{} thing for ``A'' to ``F''.}|
    return c - 'a' + 10
case 'A' <= c && c <= 'F':|\longremarkempty|
    return c - 'A' + 10
}
return 0
\end{lstlisting}
\showremarks

There is no automatic fall through, you you can use
\first{\key{fallthrough}}{keyword!fallthrough} for that.
\begin{lstlisting}[numbers=none]
switch i {
    case 0:  fallthrough
    case 1: |\longremark{\func{f()} can be called when \code{i == 0} \citem.%
With \first{\key{default}}{keyword!default} you can specify an action%
when none of the other cases match.}|
        f()
    default:
        g() |\longremark{Here \func{g()} is called when \var{i} is not 0 or 1 \citem.}|

\end{lstlisting}
\showremarks
We could rewrite the above example as:
\begin{lstlisting}[numbers=none]
switch i {
    case 0, 1:|\longremark{You can list cases on one line \citem, separated by commas.}|
        f()
    default:
        g()
\end{lstlisting}
\showremarks

\section{Built-in Functions}
A few functions are predefined, meaning
you \emph{don't} have to include any package to get
access to them. Table \ref{tab:predef-functions} lists them all.\footnote{You can use the
command \prog{godoc builtin} to read the online documentation about the built-in types and functions.}

\begin{table}[Hh!]
\begin{center}
\caption{Pre--defined functions in Go}
\label{tab:predef-functions}
\input{tab/functions.tex}
\end{center}
\end{table}

These built-in functions are documented in the \package{builtin} \index{package!builtin}
pseudo package that is included in recent Go releases. Let's go over these functions briefly.

\begin{description}
\item[\func{close}] is used in
channel communication. It closes a channel. We'll learn more about this in Chapter \ref{chap:channels}.
\index{built-in!close}

\item[\func{delete}] is used for deleting entries in maps.
\index{built-in!delete}

\item[\func{len} and \func{cap}] are used on a number of different
types, \func{len} is
used to return the lengths of strings, slices, and
arrays. In the next section \nref{sec:arrays} we'll look at slices,
arrays and the function
\func{cap}.\index{built-in!len}\index{built-in!cap}

\item[\func{new}] is used for allocating memory for user defined
data types. See \nref{sec:allocation with new} on page
\pageref{sec:allocation with new}.
\index{built-in!new}

\item[\func{make}] is used for allocating memory for built-in
types (maps, slices, and channels). See \nref{sec:allocation with make} on page
\pageref{sec:allocation with make}.
\index{built-in!make}

\item[\func{copy}] is for copying slices. See \nref{sec:slices} section in this chapter.
\index{built-in!copy}

\item[\func{append}] is for concatenating slices.
See \nref{sec:slices} in this chapter.
\index{built-in!append}

\item[\func{panic}, \func{recover}] are used for an
\emph{exception} mechanism. See \nref{sec:panic} on page \pageref{sec:panic} for more.
\index{built-in!panic}
\index{built-in!recover}

\item[\func{print}, \func{println}] are low level printing
functions that can be used without reverting to the
\package{fmt}\index{package!fmt}
package. These are mainly used for debugging.
\index{built-in!print}\index{built-in!println}

\item[\func{complex}, \func{real}, \func{imag}] all deal with
\first{complex numbers}{complex numbers}. We will not use complex numbers in this book.
(((built-in,complex)))
(((built-in,real)))
(((built-in,imag)))
\end{description}

\section{Arrays, Slices, and Maps}
\label{sec:arrays}
To store multiple values in a list, you can use arrays, or
their more flexible cousin: slices. A dictionary or hash type is also
available. It is called a \type{map} in Go.

\subsection{Arrays}
An array is defined by: \verb|[n]<type>|, where $n$ is the length
of the array and \verb|<type>| is the stuff you want to store.
To assign or index an element in the array, you use square brackets:
\begin{lstlisting}[numbers=none]
var arr [10]int
arr[0] = 42
arr[1] = 13
fmt.Printf("The first element is %d\n", arr[0])
\end{lstlisting}
Array types like \lstinline{var arr [10]int} have a fixed size. The
size is \emph{part} of the type.
They can't grow, because then they would have a different type. Also arrays
are values: Assigning one array to another \emph{copies} all the elements.
In particular, if you pass an array to a function it will receive a
copy of the array, not a pointer to it.

\index{array!multidimensional}
To declare an array you can use the following: \lstinline{var a [3]int}.
To initialize it to something other than zero, use a
\first{composite literal}{literal!composite}: \lstinline|a := [3]int{1, 2, 3}|.
This can be shortened to \lstinline|a := [...]int{1, 2, 3}|, where Go counts
the elements automatically.

\gomarginpar{A composite literal allows you
to assign a value directly to an array, slice, or map.
See \nref{sec:constructors and composite literals} on
page \pageref{sec:constructors and composite literals} for more information.}
When declaring arrays you \emph{always} have to type something in
between the square brackets, either a number or three dots (\verb|...|),
when using a composite literal.
When using multidimensional arrays, you can use the following syntax:
\lstinline|a := [2][2]int{ {1,2}, {3,4} }|. Now that you know about arrays you will
be delighted to learn that you will almost never use them in Go, because there is something
much more flexible: slices.
-->

\subsection{Slices}
\label{sec:slices}
A slice is similar to an array, but it can grow when new elements
are added.
A slice always refers to an underlying array. What makes slices different
from
arrays is that a slice is a pointer \emph{to} an array;
slices are \first{reference types}{reference types}.
\gomarginpar{Reference types are created with \lstinline{make}. We detail this further
in Chapter \ref{chap:beyond}.}
That means that if you assign one slice to
another, both refer to the \emph{same} underlying array. For instance, if a
function takes a slice argument, changes it makes to the elements of the
slice will be visible to the caller, analogous to passing a pointer to
the underlying array. With: \code{slice := make([]int, 10)},
you create a slice which can hold ten elements. Note that the
underlying array isn't specified.
A slice is always coupled to an array that has
a fixed size. For slices we define a \first{capacity}{slice!capacity} and a
\first{length}{slice!length}. \index{array!length}\index{array!capacity}
Figure \ref{fig:array-vs-slice} shows the creation of an array, then the creation of a slice.
First we create an array of $m$ elements of the type \lstinline{int}:
\lstinline{var array[m]int} .
Next, we create a slice from this array:
\lstinline{slice := array[:n]} .
And now we have:
\begin{itemize}
\item{\lstinline{len(slice) == n}{} ;}
\item{\lstinline{cap(slice) == m}{} ;}
\item{\lstinline{len(array) == cap(array) == m}{} .}
\end{itemize}
\begin{figure}[Hh]
\caption{Array versus slice}
\label{fig:array-vs-slice}
\begin{center}
\includegraphics[scale=0.65]{fig/array-vs-slice.pdf}
\end{center}
\end{figure}

Given an array, or another slice, a new slice is created via
\lstinline{a[n:m]}. This creates a new slice which refers to
the variable \lstinline{a}, starts at index \var{n}, and ends
before index \var{m}. It has length \lstinline{n - m}.

\begin{lstlisting}[numbers=none]
a := [...]int{1, 2, 3, 4, 5} |\longremark{First We define \citem{} an array with five elements, from index 0 to 4.}|
s1 := a[2:4] |\longremark{From this we create \citem{} a slice with the elements from index 2 to 3, this slices contains: \texttt{3, 4}.}|
s2 := a[1:5] |\longremark{Then we we create another slice \citem{} from \var{a}: with the elements from index 1 to 4, this contains: \texttt{2, 3, 4, 5}.}|
s3 := a[:]   |\longremark{With \var{a:[:]} \citem{} we create a slice with of all the elements in the array. This is a shorthand for: \texttt{a[0:len(a)]}.}|
s4 := a[:4]  |\longremark{And with \var{a[:4]} \citem{} we create a slice with the elements from index %
0 to 3, this is short for: \texttt{a[0:4]}, and gives us a slices that contains: \texttt{1, 2, 3, 4}.}|
s5 := s2[:]  |\longremark{With \var{s2[:]} we create a slice from the slice \var{s2} \citem, note that \texttt{s5} still refers to the array \texttt{a}.}|
s6 := a[2:4:5]  |\longremark{Finally, we create a slice with the elements from index 3 to 3 \emph{and} also set the cap to 4 \citem.}|
\end{lstlisting}
\showremarks

When working with slices you can overrun the bounds, consider this code.

\lstinputlisting[label=src:arrays]{src/basics/array-and-slices.go}
\showremarks

If you want to extend a slice, there are a couple of built-in functions
that make life easier:
\lstinline{append} and \lstinline{copy}. 
The append function appends zero or more values to a slice and returns the result: a slice with the same type as the original. If the original slice isn't big enough to fit the added values, append will allocate a new slice that is big enough. So the slice returned by append may refer to a different underlying array than the original slice does.
Here's an example:
\index{built-in!append}
\begin{lstlisting}[numbers=none]
s0 := []int{0, 0}
s1 := append(s0, 2)|\longremark{At \citem{} we append a single element, making \ttt{s1} equal to \texttt{[]int\{0, 0, 2\}}.}|
s2 := append(s1, 3, 5, 7)|\longremark{At \citem{} we append multiple elements, %%
making \ttt{s2} equal to \texttt{[]int\{0, 0, 2, 3, 5, 7\}}.}|
s3 := append(s2, s0...)|\longremark{And at \citem{} we append a slice, giving us \ttt{s3} equal to \texttt{[]int\{0, 0, 2, 3, 5, 7, 0, 0\}}. %%
Note the three dots used after \var{s0...}! This is needed make it clear explicit that you're appending another slice, instead of a single value.}|
\end{lstlisting}
\showremarks

The copy function copies slice elements from a source to a destination, and returns the number of elements it copied. This number is the minimum of the length of the source and the length of the destination.
For example:
\index{built-in!copy}
\begin{lstlisting}[numbers=none]
var a = [...]int{0, 1, 2, 3, 4, 5, 6, 7}
var s = make([]int, 6)
n1 := copy(s, a[0:])|\longremark{After \citem{}, now \texttt{n1} is 6, and \texttt{s} is \ttt{[]int\{0, 1, 2, 3, 4, 5\}}.}|
n2 := copy(s, s[2:])|\longremark{After \citem{}, \texttt{n2} is 4, and \texttt{s} is \ttt{[]int\{2, 3, 4, 5, 4, 5\}}.}|
\end{lstlisting}
\showremarks

<!--
\subsection{Maps}
\label{sec:maps}
Many other languages have a type similar to maps built-in. For instance, Perl has hashes,
Python has its dictionaries, and C++ also has maps (as part of the libraries).
In Go we have the
\first{\key{map}}{keyword!map} type. A \type{map} can be thought of as an array indexed by
strings (in its most simple form).

\begin{lstlisting}[numbers=none]
monthdays := map[string]int{
	"Jan": 31, "Feb": 28, "Mar": 31,
	"Apr": 30, "May": 31, "Jun": 30,    |\longremark{The general syntax for defining a map is {\tt map[<from type>]<to type>}. %
Here, we define a map that converts from a \key{string} (month abbreviation) to an \type{int} (number of days in that month). Note that the trailing comma at %
\citem{} is \emph{required}.}|
        you need to
	"Jul": 31, "Aug": 31, "Sep": 30,
	"Oct": 31, "Nov": 30, "Dec": 31,
}
\end{lstlisting}
\showremarks

Use \key{make} when only declaring a map:
\lstinline|monthdays := make(map[string]int)|. A map is a reference type.

For indexing ("searching") the map, we use square brackets. For example,
suppose we want to print the
number of days in December:\newline %% the code will overflow otherwise
\noindent\lstinline{fmt.Printf("%d\n", monthdays["Dec"])}

If you are looping over an array, slice, string, or map a,
\first{\key{range}}{keyword!range}
clause will help you again, it returns the key and corresponding value
with each invocation.\index{keyword!range!on maps}
\begin{lstlisting}
year := 0
for _, days := range monthdays |\longremark{At \citem{} we use the underscore to ignore (assign to nothing) the key returned by \key{range}. %
We are only interested in the values from \var{monthdays}.}|
    year += days
}
fmt.Printf("Numbers of days in a year: %d\n", year)
\end{lstlisting}
\showremarks

\index{keyword!map!add elements}
To add elements to the map, you would add new month with: \lstinline|monthdays["Undecim"] = 30|. If you use a key that
already exists, the value will be silently overwritten: \lstinline|monthdays["Feb"] = 29|.
To test for existence \index{keyword!map!existence}, you would use the
following: \lstinline{value, present := monthdays["Jan"]}. If the key "Jan" exists, \var{present}
will be true. It's more Go like to name \var{present} "ok", and use:
\lstinline{v, ok := monthdays["Jan"]}. In Go we call this the "comma ok" form.

You can remove elements \index{keyword!map!remove elements} from the \type{map}:
\lstinline{delete(monthdays, "Mar")}\footnote{Always rainy in March anyway.}.
In general the syntax \lstinline{delete(m, x)} will delete the map entry
retrieved by the expression \lstinline{m[x]}.

\section{Exercises}
\input{ex/basics/ex.tex}
\cleardoublepage
\section{Answers}
\shipoutAnswer
-->
