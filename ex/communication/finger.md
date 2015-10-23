{.exercise data-difficulty="1"}
### Finger daemon

Write a finger daemon that works with the finger(1) command. 

From the [Debian](https://www.debian.org) package description:

> Fingerd is a simple daemon based on RFC 1196 [@RFC1196] that provides an interface to the
> "finger" program at most network sites.  The program is supposed to return a
> friendly, human-oriented status report on either the system at the moment or a
> particular person in depth.


Stick to the basics and only support a username argument. If the user has a `.plan` file
show the contents of that file. So your program needs to be able to figure out:

* Does the user exist?
* If the user exists, show the contents of the `.plan` file.

### Answer
A> This solution is from Fabian Becker.

<{{ex/communication/src/finger.go}}
