{.exercise data-difficulty="1"}
### Echo server

Write a simple echo server. Make it listen to TCP port number 8053 on localhost.
It should be able to read a line (up to the newline), echo back that line and
then close the connection. 

Make the server concurrent so that every request is taken care of in a separate
goroutine.

### Answer

A simple echo server might be:
<{{ex/communication/src/echo.go}}

When started you should see the following:

    % nc 127.0.0.1 8053
    Go is *awesome*
    Go is *awesome*


To make the connection handling concurrent we *only need to change one line* in our
echo server, the line:

~~~go
if c, err := l.Accept(); err == nil { Echo(c) }
~~~

becomes:

~~~go
if c, err := l.Accept(); err == nil { go Echo(c) }
~~~
