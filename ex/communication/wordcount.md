{.exercise data-difficulty="0"}
### Word and Letter Count

Write a small program that reads text from standard input and performs the
following actions:

* Count the number of characters (including spaces).
* Count the number of words.
* Count the numbers of lines

In other words implement wc(1) (check you local manual page), however you only
have to read from standard input.


### Answer

The following program is an implementation of wc(1).
{callout="//"}
<{{ex/communication/src/wc.go}}

At <1> we create a new reader that reads from standard input, we then read from
the input at <2>. And at <3> we check the value of `ok` and if we received an
error, we assume it was because of a EOF, So we print the current values;.
Otherwise <4> we count the charaters, words and increment the number lines.
