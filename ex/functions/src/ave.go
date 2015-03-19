package main

func average(xs []float64) (avg float64) { |\longremark{At \citem{} We use a named return parameter.}|
	sum := 0.0
	switch len(xs) {
	case 0:                 |\longremark{If the length of \ttt{xs} is zero \citem{}, we return 0.}|
		avg = 0
	default:                |\longremark{Otherwise \citem, we calculate the average.}|
		for _, v := range xs {
			sum += v
		}
		avg = sum / float64(len(xs)) |\longremark{At \citem{} we % 
convert the value to a \key{float64} to make the division work as \ttt{len} returns an \ttt{int}.}|
	}
	return  |\longremark{Finally, at \citem{} we reutrn our avarage.}|
}
