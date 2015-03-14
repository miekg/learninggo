package main

func main() {
	var array [100]int   //<1>
	slice := array[0:99] //<2>

	slice[98] = 1 //<3>
	slice[99] = 2 //<4>
}
