package main

import (
	"bufio"
	"log"
	"os"
)

func main() {
	buf := make([]byte, 1024)
	f, e := os.Open("/etc/passwd") //<1>
	if e != nil {
		log.Fatalf(e)
	}
	defer f.Close()
	r := bufio.NewReader(f) //<2>
	w := bufio.NewWriter(os.Stdout)
	defer w.Flush() //<3>
	for {
		n, e := r.Read(buf) //<4>
		if e != nil {
			log.Fatalf(e)
		}
		if n == 0 {
			break
		}
		w.Write(buf[0:n]) //<5>
	}
}
