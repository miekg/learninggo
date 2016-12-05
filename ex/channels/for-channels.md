{.exercise data-difficulty="1"}
### Channels

1. Modify the program you created in exercise (#forloop) to use
   channels, in other words, the function called in the body should now be
   a goroutine and communication should happen via channels. You should not
   worry yourself on how the goroutine terminates.

2. There are a few annoying issues left if you resolve question 1 above. One of
   the problems is that the goroutine isn't neatly cleaned up when `main.main()`
   exits. And worse, due to a race condition between the exit of `main.main()`
   and `main.shower()` not all numbers are printed. It should print up until 9,
   but sometimes it prints only to 8. Adding a second quit-channel you can
   remedy both issues. Do this.

### Answer
1. A possible program is:

   <{{ex/channels/src/for-chan.go}}

   We start in the usual way, then at line 6 we create a new channel of
   ints. In the next line we fire off the function `shower` with
   the `ch` variable as it argument, so that we may communicate with
   it. Next we start our for-loop (lines 8-10) and in the loop
   we send (with `<-`) our number to the function (now a goroutine) `shower`.

   In the function `shower` we wait (as this blocks) until we receive a number
   (line 15). Any received number is printed (line 16) and then continue the
   endless loop started on line 14.

2. An answer is

   <{{ex/channels/src/for-quit-chan.go}}

   On line 20 we read from the quit channel and we discard the value we read. We
   could have used `q := <-quit`, but then we would have used the variable only
   once --- which is illegal in Go. Another trick you might have pulled out of
   your hat may be: `_ = <-quit`. This is valid in Go, but idomatic Go is the
   one given on line 20.
