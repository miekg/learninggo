{.exercise data-difficulty="0"}
### Uniq

Write a Go program that mimics the function of the Unix `uniq` command. This
program should work as follows, given a list with the following items: 

    'a' 'b' 'a' 'a' 'a' 'c' 'd' 'e' 'f' 'g'

it should print only those items which don't have the same successor:

    'a' 'b' 'a' 'c' 'd' 'e' 'f' 'g'

The next listing is a Perl implementation of the algorithm.
<{{ex/communication/src/uniq.pl}}

### Answer

The following is a `uniq` implementation in Go.
<{{ex/communication/src/uniq.go}}
