package main

import (
	"os"

	"github.com/01-edu/z01"
)

func main() {
	arguments := os.Args[0]
	for _, ch := range arguments {
		if ch != rune(47) && ch != rune(46) {
			z01.PrintRune(ch)
		}
	}
}
