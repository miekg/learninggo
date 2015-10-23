{.exercise data-difficulty="0"}
### Stack as package

1. See the Stack exercise. In this exercise we want to create a separate package
   for that code. Create a proper package for your stack implementation, `Push`,
   `Pop` and the `Stack` type need to be exported.

2. Write a simple unit test for this package.
 You should at least test that a `Pop` works after a `Push`.


{.answer}
### Answer
1. There are a few details that should be changed to make a proper package
 for our stack. First, the exported functions should begin with a capital
 letter and so should `Stack`. The package file is named `stack-as-package.go`
 and contains:
    <{{ex/packages/src/stack-as-package.go}}

2. To make the unit testing work properly you need to do some
 preparations. We'll come to those in a minute. First the actual unit test.
 Create a file with the name `pushpop_test.go`, with the following contents:
 <{{ex/packages/src/pushpop_test.go}}

For `go test` to work we need to put our package files in a directory
under `$GOPATH/src`:

    % mkdir $GOPATH/src/stack
    % cp pushpop_test.go $GOPATH/src/stack
    % cp stack-as-package.go $GOPATH/src/stack

Yields:

    % go test stack
    ok stack 0.001s
