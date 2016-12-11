{.exercise data-difficulty="0"}
### Maximum
1.  Write a function that finds the
maximum value in an `int` slice (`[]int`).


{.answer}
### Answer
1.  This function returns the largest int in the slice \var{l}:

	{callout="//"}
	~~~go
	func max(l []int) (max int) {   //<1>
	    max = l[0]
	    for _, v := range l {   //<2>
	        if v > max {    //<3>
	            max = v
	        }
	    }
	    return //<4>
	}
	~~~

	At <1> we use a named return parameter.
	At <2> we loop over `l`. The index of the element is not important.
	At <3>, if we find a new maximum, we remember it.
	And at <4> we have a "lone" return; the current value of `max` is now returned.
