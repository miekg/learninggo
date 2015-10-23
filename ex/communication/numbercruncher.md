{.exercise data-difficulty="2"}
### Number cruncher


* Pick six (6) random numbers from this list: $$1, 2, 3, 4, 5, 6, 7, 8, 9, 10,
  25, 50, 75, 100$$ Numbers may be picked multiple times.
* Pick one (1) random number ($$i$$) in the range: $$1 \ldots 1000$$.
* Tell how, by combining the first 6 numbers (or a subset thereof)
with the operators `+,-,*` and `/`, you can make $$i$$.

An example. We have picked the numbers: 1, 6, 7, 8, 8 and 75. And $$i$$ is
977. This can be done in many different ways, one way is:
$$ ((((1 * 6) * 8) + 75) * 8) - 7 = 977$$
or
$$ (8*(75+(8*6)))-(7/1) = 977$$

Implement a number cruncher that works like that. Make it print the solution in
a similar format (i.e. output should be infix with parenthesis) as used above.

Calculate *all* possible solutions and show them (or only show how many there
are). In the example above there are 544 ways to do it.


### Answer

The following is one possibility. It uses recursion and backtracking to get
an answer. When starting `permrec` we give 977 as the first argument:


    % ./permrec 977
    1+(((6+7)*75)+(8/8)) = 977  #1
    ...                         ...
    ((75+(8*6))*8)-7 = 977      #542
    (((75+(8*6))*8)-7)*1 = 977  #543
    (((75+(8*6))*8)-7)/1 = 977  #544


<{{ex/communication/src/permrec.go}}
