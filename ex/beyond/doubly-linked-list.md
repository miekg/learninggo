{.exercise data-difficulty="1"}
### Linked List

1. Make use of the package `container/list` to create
 a (doubly) linked list. Push the values 1, 2 and 4 to the list and then
 print it.

2.  Create your own linked list implementation. And perform the same actions
as above.

{.answer}
### Answer
1. The following is the implementation of a program using doubly
   linked lists from `container/list`.
   <{{ex/beyond/src/doubly-linked-list-container.go}}

2. The following is a program implementing a simple doubly
 linked list supporting `int` values.
 {callout="//"}
  <{{ex/beyond/src/doubly-linked-list.go}}

Import <1> the packages we will need. At <2> we declare a type for the value our list will contain,
this is not strictly neccesary. And at <3> we declare a type for the each node in our list.
At <4> we define the `Front` method for our list.
When pushing, create a new Node <5> with the provided value. If the list is empty <6>,
put the new node at the head. Otherwise <7> put it at the tail and make sure <8>
the new node points back to the previously existing one. At <9> we re-adjust tail
to the newly inserted node.

In the Pop <10> method, we return an error if the list is empty. If it is not empty <11>
we save the last value. And then <12> discard the last node from the list. Finally at <13>
we make sure the list is consistent if it becomes empty.
