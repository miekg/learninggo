package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	var chars, words, lines int
	r := bufio.NewReader(os.Stdin) //<1>
	for {
		switch s, ok := r.ReadString('\n'); true { //<2>
		case ok != nil: //<3>
			fmt.Printf("%d %d %d\n", chars, words, lines)
			return
		default: //<4>
			chars += len(s)
			words += len(strings.Fields(s))
			lines++
		}
	}
}
