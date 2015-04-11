{.exercise data-difficulty="1"}
### Strings

1. Create a Go program that prints the following (up to 100 characters):

        A
        AA
        AAA
        AAAA
        AAAAA
        AAAAAA
        AAAAAAA
        ...

2. Create a program that counts the number of characters in this string:

        I am a Gő programmer.

In addition, make it output the number of bytes in that string.
*Hint*: Check out the `utf8` package.

3. Extend/change the program from the previous question to replace the three runes at
position 4 with 'abc'.

4. Write a Go program that reverses a string, so "foobar" is printed as "raboof".
*Hint*: You will need to know about conversion; skip ahead to (#conversions).


{.answer}
### Answer

1. This program is a solution:

   <{{ex/basics/src/string1.go}}

2. To answer this question we need some help from
the `unicode/utf8` package. First we check the documentation
with `godoc unicode/utf8 | less`, with has the
`func RuneCount(p []byte) int` function. Secondly
we can convert *string* to a `byte` slice with:

    str := "hello"
    b   := []byte(str)

On line 2, we use a conversion. Here we convert a `string`
to a slice of `byte`s. Putting this together leads to the following program.

<{{ex/basics/src/string2.go}}

3. Something along the lines of:

<{{ex/basics/src/string3.go}}

4. Reversing a string can be done as follows. We start from the left (`i`) and
the right (`j`) and swap the characters as we see them:

{callout="//"}
<{{ex/basics/src/stringrev.go}}

At <1> we have a conversion.
At <2> we use parallel assignment.
And at <3> we convert it back.
