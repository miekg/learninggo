{.exercise data-difficulty="1"}
### Interfaces and compilation

The code in listing (#src:interface fail) on page
\pageref{src:interface fail} compiles OK --- as stated
in the text. But when you run it you'll get a runtime error, so
something *is* wrong. Why does the code compile cleanly then?

### Answer

The code compiles because an integer type implements the empty interface and
that is the check that happens at compile time.

A proper way to fix this is to test if such an empty interface can
be converted and, if so, call the appropriate method. The Go code
that defines the function `g` in listing (#src:interface empty)
-- repeated here:

    func g(any interface{}) int { return any.(I).Get() }

Should be changed to become:

{callout="//"}
    func g(any interface{}) int {
        if v, ok := any.(I); ok { //<1>
            return v.Get()		  //<2>
        }
        return -1			|\coderemark{Just so we return anything}|
    }

At <1> we check if the conversion is allowed, and if set we invoke `Get()` at
<2>. The idiom used is called "comma ok" in Go. At <3> we return -1 to ratify
the function signature. A better way would be to return a tuple `(int, error)`
and explicitly return an error.

If `g()` is called now there are no run-time errors anymore.
