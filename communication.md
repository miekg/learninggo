{.epigraph}
> Good communication is as stimulating as black coffee, and just as hard
> to sleep after.
Quote: -- Anne Morrow Lindbergh

In this chapter we are going to look at the building blocks in Go for
communicating with the outside world. We will look at files, directories,
networking and executing other programs. Central to Go's I/O are the interfaces
`io.Reader` and `io.Writer`. The `io.Reader` interface specifies one method
`Read(p []byte) (n int, err err)`.

Reading from (and writing to) files is easy in Go. This program
only uses the `os` package to read data from the file `/etc/passwd`.

{callout="//"}
<{{src/communication/file.go}}

We open the file at <1> with `os.Open` that returns a `*os.File`
`*os.File` implements `io.Reader` and `io.Writer` interface.
After the `Open` we directly put the `f.Close()` which we defer until the function
return. At <3> we call `Read` on `f` and read up to 1024 bytes at the time. If anything
fails we bail out at <4>. If the number of bytes read is 0 we've read the end of the
file <5>. And at <6> we output the buffer to standard output.

If you want to use buffered (((io, buffered))) I/O there is the
`bufio`(((package, bufio))) package:

{callout="//"}
<{{src/communication/buffile.go}}

Again, we open <1> the file. Then at <2> we
Turn `f` into a buffered `Reader`. `NewReader` expects an `io.Reader`, so you this will work.
Then at <4> we read and at <5> we write. We also call `Flush()` at <3> to flush all output.
This entire program could be optimized further by using `io.Copy`.


## io.Reader

As mentioned above the `io.Reader` (((io.Reader))) is an important interface in the language Go. A lot
(if not all) functions that need to read from something take an `io.Reader`(((package, io)))
as input. To fulfill the interface a type needs to implement that one method.
The writing side `io.Writer`, has the `Write` method.

If you think of a new type in your program or package and you make it fulfill the `io.Reader`
or `io.Writer` interface, *the whole standard Go library can be used* on that type!


## Some examples

The previous program reads a file in its entirety, but a more common scenario is that
you want to read a file on a line-by-line basis. The following snippet shows a way
to do just that (we're discarding the error returned from `os.Open` here to keep
the examples smaller -- don't ever do this in real life code).

{callout="//"}
~~~go
f, _ := os.Open("/etc/passwd"); defer f.Close()
r := bufio.NewReader(f) //<1>
s, ok := r.ReadString('\n') //<2>
~~~

At <1> make `f` a `bufio` to have access to the `ReadString` method. Then at <2> we read
a line from the input, `s`  now holds a string which we can manipulate with, for instance,
the `strings` package.

A more robust method (but slightly more complicated) is `ReadLine`, see the documentation
of the `bufio` package.

A common scenario in shell scripting is that you want to check if a directory
exists and if not, create one.

~~~go
if [ ! -e name ]; then          if f, e := os.Stat("name"); e != nil {
    mkdir name                      os.Mkdir("name", 0755)
else                            } else {
    # error                         // error
fi                              }
~~~

The similarity between these two examples (and with other scripting languages)
have prompted comments that Go has a "script"-like feel to it, i.e. programming
in Go can be compared to programming in an interpreted language (Python, Ruby,
Perl or PHP).


## Command line arguments

Arguments from the command line are available inside your program via the string
slice `os.Args`, provided you have imported the package `os`. The `flag` package
(((package, flag)))
has a more sophisticated interface, and also provides a way to parse flags. Take
this example from a DNS query tool:

{callout="//"}
~~~go
dnssec := flag.Bool("dnssec", false, "Request DNSSEC records") //<1>
port := flag.String("port", "53", "Set the query port") //<2>
flag.Usage = func() {   //<3>
    fmt.Fprintf(os.Stderr, "Usage: %s [OPTIONS] [name ...]\n", os.Args[0])
    flag.PrintDefaults() //<4>
}
flag.Parse() //<4>
~~~

At <1> we define a `bool` flag `-dnssec`. Note that this function returns
a *pointer* to the value, the `dnssec` is now a pointer to a `bool`. At <2> we
define an `strings` flag. Then at <3> we *redefine* the `Usage` variable of the
flag package so we can add some extra text. The `PrintDefaults` at <4> will
output the default help for the flags that are defined. Note even without
redefining a `flag.Usage` the flag `-h` is supported and will just output the help text
for each of the flags. Finally at <4> we call `Parse` that parses the command
line and fills the variables.

After the flags have been parsed you can used them: `if *dnssec { ... }`


## Executing commands

The `os/exec`(((package,os/exec))) package has functions to run external
commands, and is the premier way to execute commands from within a Go program.
It works by defining a `*exec.Cmd` structure for which it defines a number of
methods. Let's execute `ls -l`:

~~~go
import "os/exec"

cmd := exec.Command("/bin/ls", "-l")
err := cmd.Run()
~~~

The above example just runs "ls -l" without doing anything with the returned
data, capturing the standard output from a command is done as follows:

~~~go
cmd := exec.Command("/bin/ls", "-l")
buf, err := cmd.Output()
~~~

And `buf` is byte slice, that you can further use in your program.


## Networking

All network related types and functions can be found in the package `net`. One
of the most important functions in there is `Dial`(((networking, Dial))). When
you `Dial` into a remote system the function returns a `Conn` interface type,
which can be used to send and receive information. The function `Dial` neatly
abstracts away the network family and transport. So IPv4 or IPv6, TCP or UDP can
all share a common interface.

Dialing a remote system (port 80) over TCP, then UDP and lastly TCP over IPv6
looks like this^[In case you are wondering, 192.0.32.10 and 2620:0:2d0:200::10
are <http://www.example.org>.]:

~~~go
conn, e := Dial("tcp", "192.0.32.10:80")
conn, e := Dial("udp", "192.0.32.10:80")
conn, e := Dial("tcp", "[2620:0:2d0:200::10]:80")
~~~

If there were no errors (returned in `e`), you can use `conn` to read and write.
And `conn` implements the `io.Reader` and `io.Writer` interface. ^[The variable
`conn` also implements a `close` method, this really makes it an
`io.ReadWriteCloser`.]

But these are the low level nooks and crannies, you will almost always use
higher level packages, such as the `http` package. For instance a simple Get for
http:

~~~go
package main

import (
    "fmt"
    "http"
    "io/ioutil"
)

func main() {
    r, err := http.Get("http://www.google.com/robots.txt")
    if err != nil {
        fmt.Printf("%s\n", err.String())
        return
    }
    b, err := ioutil.ReadAll(r.Body)
    r.Body.Close()
    if err == nil {
        fmt.Printf("%s", string(b))
    }
}
~~~


## Exercises

{{ex/communication/ex.md}}
