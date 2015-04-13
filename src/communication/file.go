package main

import (
	"log"
	"os"
)

func main() {
	buf := make([]byte, 1024)
	f, e := os.Open("/etc/passwd") //<1>
	if e != nil {
		log.Fatalf(e)
	}
	defer f.Close()                //<2>
	for {
		n, e := f.Read(buf) //<3>
		if e != nil {
			log.Fatalf(e) //<4>
		}
		if n == 0 { //<5>
			break
		}
		os.Stdout.Write(buf[:n]) //<6>
	}
}
