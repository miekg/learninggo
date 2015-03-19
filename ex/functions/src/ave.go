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
