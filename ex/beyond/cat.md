{.exercise data-difficulty="1"}
### Cat

1. Write a program which mimics the Unix program `cat`.

2. Make it support the `-n` flag, where each line is numbered.

3. The solution to the above question given in contains a bug. Can you spot and fix it?


{.answer}
### Answer
1. The following is implemention of `cat` which also supports a \-n flag to number each line.

 {callout="//"}
 <{{ex/beyond/src/cat.go}}

 At <1> we include all the packages we need.
 Here <2> we define a new flag "n", which defaults to off. Note that we get the help (-h) for free.
 Start the function <3> that actually reads the file's contents and displays it;
 Read one line at the time at <4>. And stop <5> if we hit the end.
 If we should number each line, print the line number and then the line itself <6>.
 Otherwise <7> we could just print the line.
 
2. The bug show itself when the last line of the input does not
  contain a newline. Or worse, when the input contains one line without a
  closing newline nothing is shown at all. A better solution is the following 
  program.
  <{{ex/beyond/src/cat2.go}}
