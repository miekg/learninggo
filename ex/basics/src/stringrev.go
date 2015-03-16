package main

import "fmt"

func main() {
  s := "foobar"
  a := []rune(s) |\longremark{At \citem we have a conversion.}|
  for i, j := 0, len(a)-1; i < j; i, j = i+1, j-1 {
     a[i], a[j] = a[j], a[i] |\longremark{At \citem{} we use parallel assignment.}|
  }
  fmt.Printf("%s\n", string(a)) |\longremark{And at \citem we convert it back.}|
}
