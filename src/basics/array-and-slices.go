package main

func main() {
var array [100]int |\longremark{At \citem{} we create an array with a 100 elements, indexed from 0 to 99.}|
slice := array[0:99] |\longremark{Then at \citem{} we create a slice that has index 0 to 98.}|

slice[98] = 1	|\longremark{We assign $1$ to the 99th element \citem{} of the slice. This \citem{} works as expected.}|
slice[99] = 2 |\longremark{But at \citem{} we dare to do the impossible, and and try to allocate something %
beyond the length of the slice and we are greeted with a \emph{runtime} error: Error: "throw: index out of range".}|
}
