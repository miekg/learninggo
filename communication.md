\epi{"Good communication is as stimulating as black coffee, and just as hard
to sleep after."}{\textsc{ANNE MORROW LINDBERGH}}
In this chapter we are going to look at the building blocks in Go for 
communicating with the outside world. We will look at files, directories, networking
and executing other programs. Central to Go's I/O are the interfaces `io.Reader`
and `io.Writer`.

Reading from (and writing to) files is easy in Go. This program
only uses the `os` package to read data from the file \file{/etc/passwd}.
\lstinputlisting[caption=Reading from a file (unbuffered),label=src:read]{src/file.go}
The following is happening here:
\showremarks
If you want to use \first{buffered}{buffered} IO there is the
`bufio`(((package!bufio))) package:
\lstinputlisting[caption=Reading from a file (buffered),label=src:bufread]{src/buffile.go}
\showremarks

## io.Reader
As mentioned above the \first{io.Reader}{io.Reader} is an important interface in the language Go. A lot
(if not all) functions that need to read from something take an `io.Reader`(((package!io)))
as input. To fulfill the interface a type needs to implement only one method: \func{Read(p []byte) (n
int, err error)}. The writing side is (you may have guessed) an `io.Writer`, which has
the `Write` method.

If you think of a new type in your program or package and you make it fulfill the `io.Reader`
or `io.Writer` interface, *the whole standard Go library can be used* on that type!

## Some examples
The previous program reads a file in its entirety, but a more common scenario is that
you want to read a file on a line-by-line basis. The following snippet shows a way
to do just that:

\begin{lstlisting}
f, _ := os.Open("/etc/passwd"); defer f.Close()
r := bufio.NewReader(f) |\coderemark{Make it a bufio to access the ReadString method}|
s, ok := r.ReadString('\n') |\coderemark{Read a line from the input}|
// ... \coderemark{`s` holds the string, with the `strings` package you can parse it}
\end{lstlisting}

A more robust method (but slightly more complicated) is `ReadLine`, see the documentation
of the `bufio` package.

A common scenario in shell scripting is that you want to check if a directory
exists and if not, create one. 

\begin{minipage}{.5\textwidth}
\begin{lstlisting}[language=sh,caption={Create a directory with the shell}]
if [ ! -e name ]; then
    mkdir name
else
    # error
fi
\end{lstlisting}
\end{minipage}
\hspace{1em}
\begin{minipage}{.5\textwidth}
\begin{lstlisting}[caption={Create a directory with Go}]
if f, e := os.Stat("name"); e != nil {
    os.Mkdir("name", 0755)
} else {
    // error
}
\end{lstlisting}
\end{minipage}
The similarity between these two examples have prompted comments that Go has a
"script"-like feel to it, i.e. programming in Go can be compared to programming in
an interpreted language (Python, Ruby, Perl or PHP).

## Command line arguments

Arguments from the command line are available inside your program via
the string slice `os.Args`, provided you have imported the package
`os`. The `flag` package has a more sophisticated
interface, and also provides a way to parse flags. Take this example
from a DNS query tool:
\begin{lstlisting}
dnssec := flag.Bool("dnssec", false, "Request DNSSEC records") |\longremark{Define a `bool` flag, %%
`-dnssec`. The variable must be a pointer otherwise the package can not set its value;}|
port := flag.String("port", "53", "Set the query port")      |\longremark{Idem, but for a `port` option;}|
flag.Usage = func() {   |\longremark{Slightly redefine the `Usage` function, to be a little more verbose;}|
    fmt.Fprintf(os.Stderr, "Usage: %s [OPTIONS] [name ...]\n", os.Args[0])
    flag.PrintDefaults() |\longremark{For every flag given, `PrintDefaults` will output the help string;}|
}
flag.Parse()   |\longremark{Parse the flags and fill the variables.}|
\end{lstlisting}
\showremarks
After the flags have been parsed you can used them:
\begin{lstlisting}
if *dnssec {    |\coderemark{Dereference the `dnssec` flag variable}|
    // do something
}
\end{lstlisting}

## Executing commands
The `os/exec`(((package!os/exec))) package has functions to run external commands, and is the premier way to
execute commands from within a Go program. It works by defining a `*exec.Cmd` structure for which it
defines a number of methods.
Let's execute \verb|ls -l|:
\begin{lstlisting}
import "os/exec"

cmd := exec.Command("/bin/ls", "-l")    |\coderemark{Create a `*cmd`}|
err := cmd.Run()                        |\coderemark{`Run()` it}|
\end{lstlisting}
The above example just runs "ls -l" without doing anything with the returned data,
capturing the standard output from a command is done as follows:
\begin{lstlisting}
import "os/exec"

cmd := exec.Command("/bin/ls", "-l")
buf, err := cmd.Output()                 |\coderemark{`buf` is a (`[]byte`)}|
\end{lstlisting}

## Networking
All network related types and functions can be found in the package `net`. One of the
most important functions in there is `Dial`(((networking!Dial))). When you `Dial`
into a remote system the function returns a `Conn` interface type, which can be used
to send and receive information. The function `Dial` neatly abstracts away the network
family and transport. So IPv4 or IPv6, TCP or UDP can all share a common interface. 

Dialing a remote system (port 80) over TCP, then UDP and lastly TCP over IPv6 looks
like this\footnote{In case
you are wondering, 192.0.32.10 and 2620:0:2d0:200::10 are \url{www.example.org}.}:
\begin{lstlisting}
conn, e := Dial("tcp", "192.0.32.10:80")
conn, e := Dial("udp", "192.0.32.10:80")
conn, e := Dial("tcp", "[2620:0:2d0:200::10]:80") |\coderemark{Mandatory brackets}|
\end{lstlisting}

If there were no errors (returned in `e`), you can use `conn` to read and write.
The primitives defined in the package `net` are:
\begin{quote}
// `Read` reads data from the connection.\\
`Read(b []byte) (n int, err error)`
\end{quote}
This makes `conn` an `io.Reader`.

\begin{quote}
// `Write` writes data to the connection.\\
`Write(b []byte) (n int, err error)`
\end{quote}
This makes `conn` also an `io.Writer`, in fact `conn` is an
`io.ReadWriter`.\footnote{The variable `conn` also implements a `close` method, this really makes
it an `io.ReadWriteCloser`.}

But these are the low level nooks and crannies\footnote{Exercise Q(#ex:echo) is about using
these.}, you will almost always use higher level packages.
Such as the `http` package. For instance a simple Get for http:
\begin{lstlisting}
package main
import ( "io/ioutil"; "http"; "fmt" ) |\longremark{The imports needed;}|

func main() {
        r, err := http.Get("http://www.google.com/robots.txt") |\longremark{Use http's `Get` to retrieve the html;}|
        if err != nil { fmt.Printf("%s\n", err.String()); return } |\longremark{Error handling;}|
        b, err := ioutil.ReadAll(r.Body)    |\longremark{Read the entire document into `b`;}|
        r.Body.Close()  
        if err == nil { fmt.Printf("%s", string(b)) } |\longremark{If everything was OK, print the document.}|
}
\end{lstlisting}
\showremarks

## Exercises
\input{ex-communication/ex-processes.tex}

\input{ex-communication/ex-wordcount.tex}

\input{ex-communication/ex-uniq.tex}

\input{ex-communication/ex-quine.tex}

\input{ex-communication/ex-echo.tex}

\input{ex-communication/ex-numbercruncher.tex}

\input{ex-communication/ex-finger.tex}

\cleardoublepage
## Answers
\shipoutAnswer
