package main

import (
	"fmt"
	"unicode/utf8"
)

func main() {
	str := "I am a Gö programmer."
	fmt.Printf("String %s\nLength: %d, Runes: %d\n", str,
		len([]byte(str)), utf8.RuneCount([]byte(str)))
}
