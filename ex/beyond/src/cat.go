package main

import (
	"bufio"
	"flag"
	"fmt"
	"io" //<1>
	"os"
)

var numberFlag = flag.Bool("n", false, "number each line") // <2>

func cat(r *bufio.Reader) { //<3>
	i := 1
	for {
		buf, e := r.ReadBytes('\n') //<4>
		if e == io.EOF {            //<5>
			break
		}
		if *numberFlag { //<6>
			fmt.Fprintf(os.Stdout, "%5d  %s", i, buf)
			i++
		} else { //<7>
			fmt.Fprintf(os.Stdout, "%s", buf)
		}
	}
	return
}

func main() {
	flag.Parse()
	if flag.NArg() == 0 {
		cat(bufio.NewReader(os.Stdin))
	}
	for i := 0; i < flag.NArg(); i++ {
		f, e := os.Open(flag.Arg(i))
		if e != nil {
			fmt.Fprintf(os.Stderr, "%s: error reading from %s: %s\n",
				os.Args[0], flag.Arg(i), e.Error())
			continue
		}
		cat(bufio.NewReader(f))
	}
}
