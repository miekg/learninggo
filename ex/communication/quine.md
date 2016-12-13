{.exercise difficulty="2"}
### Quine
A *Quine* is a program that prints itself. Write a Quine in Go.

### Answer
A> This solution is from Russ Cox. It was posted to the Go Nuts mailing list.

~~~go
/* Go quine */
package main
import "fmt"
func main() {
 fmt.Printf("%s%c%s%c\n", q, 0x60, q, 0x60)
}
var q = `/* Go quine */
package main
import "fmt"
func main() {
 fmt.Printf("%s%c%s%c\n", q, 0x60, q, 0x60)
}
var q = `
~~~
