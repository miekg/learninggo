{.exercise data-difficulty="1"}
### Map function

A `map()`-function is a function that takes
a function and a list. The function is applied to
each member in the list and a new list containing
these calculated values is returned.
Thus:

$$ \mathrm{map}(f(), (a_1,a_2,\ldots,a_{n-1},a_n)) =  (f(a_1), f(a_2),\ldots,f(a_{n-1}), f(a_n)) $$

1.  Write a simple
`map()`-function in Go. It is sufficient for this function only to work for ints.


{.answer}
### Answer

1. A possible answer:

    ~~~go
    func Map(f func(int) int, l []int) []int {
        j := make([]int, len(l))
        for k, v := range l {
            j[k] = f(v)
        }
        return j
    }

    func main() {
        m := []int{1, 3, 4}
        f := func(i int) int {
            return i * i
        }
        fmt.Printf("%v", (Map(f, m)))
    }
    ~~~
