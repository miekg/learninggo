{.exercise data-difficulty="2"}
### Processes

Write a program that takes a list of all running processes and prints how many
child processes each parent has spawned. The output should look like:

    Pid 0 has 2 children: [1 2]
    Pid 490 has 2 children: [1199 26524]
    Pid 1824 has 1 child: [7293]

* For acquiring the process list, you'll need to capture the output of `ps -e
  -opid,ppid,comm`. This output looks like:

          PID  PPID COMMAND
         9024  9023 zsh
        19560  9024 ps

* If a parent has one child you must print `child`, if there is more than one
  print `children`.

* The process list must be numerically sorted, so you start with pid 0 and work
  your way up.

Here is a Perl version to help you on your way (or to create complete and utter confusion).
<{{ex/communication/src/proc.pl}}

### Answer

 There is lots of stuff to do here. We can divide our program
up in the following sections:

* Starting \verb|ps| and capturing the output.
* Parsing the output and saving the child PIDs for each PPID.
* Sorting the PPID list.
* Printing the sorted list to the screen.

In the solution presented below, we've used a `map[int][]int`, i.e. a map
indexed with integers, pointing to a slice of ints -- which holds the PIDs. The
builtin `append` is used to grow the integer slice. 

A possible program is: 
<{{ex/communication/src/proc.go}}
