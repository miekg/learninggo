{.exercise data-difficulty="1"}
### Bubble sort
1. Write a function that performs a bubble sort on a slice of ints. From [@bubblesort]:

  > It works by repeatedly stepping through the list to be sorted, comparing each
  > pair of adjacent items and swapping them if they are in the wrong order. The
  > pass through the list is repeated until no swaps are needed, which indicates
  > that the list is sorted. The algorithm gets its name from the way smaller
  > elements "bubble" to the top of the list. 

It also gives an example in pseudo code:

    procedure bubbleSort( A : list of sortable items )
      do
        swapped = false
        for each i in 1 to length(A) - 1 inclusive do:
          if A[i-1] > A[i] then
            swap( A[i-1], A[i] )
            swapped = true
          end if
        end for
      while swapped
    end procedure

### Answer

1.  Bubble sort isn't terribly efficient. For $$n$$ elements it scales $$O(n^2)$$.
    But bubble sort is easy to implement:

    <{{ex/functions/src/bubblesort.go}}[4,18]

    Because a slice is a reference type, the `bubblesort` function works and
    does not need to return a sorted slice.
