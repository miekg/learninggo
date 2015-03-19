package main

func average(xs []float64) (avg float64) { //<1>
	sum := 0.0
	switch len(xs) {
	case 0:                 //<2>
		avg = 0
	default:                //<3>
		for _, v := range xs {
			sum += v
		}
		avg = sum / float64(len(xs)) //<4>
	}
	return  //<5>
}

At <1> we use a named return parameter.  If the length of `xs` is zero <2>, we
return 0.  Otherwise <3>, we calculate the average.  At <4>  we convert the
value to a `float64` to make the division work as `len` returns an `int`.
Finally, at <5> we reutrn our avarage.
